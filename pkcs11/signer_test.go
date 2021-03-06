package pkcs11

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"reflect"
	"testing"
	"time"

	"golang.org/x/crypto/ssh"

	"github.com/theparanoids/crypki"
)

const (
	defaultIdentifier = "dummy"
	badIdentifier     = "unknown"
	timeout           = 1 * time.Second
)

// enforce signer implements CertSign interface.
var _ crypki.CertSign = (*signer)(nil)

// createCAKeysAndCert generates key pairs and the corresponding x509 certificate for unit tests CA based on key type.
func createCAKeysAndCert(keyType crypki.PublicKeyAlgorithm) (priv crypto.Signer, cert *x509.Certificate, err error) {
	var pkAlgo x509.PublicKeyAlgorithm
	var sigAlgo x509.SignatureAlgorithm
	switch keyType {
	case crypki.ECDSA:
		pkAlgo = x509.ECDSA
		sigAlgo = x509.ECDSAWithSHA256
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case crypki.RSA:
		fallthrough
	default:
		pkAlgo = x509.RSA
		sigAlgo = x509.SHA256WithRSA
		priv, err = rsa.GenerateKey(rand.Reader, 2048)
	}
	if err != nil {
		return nil, nil, err
	}

	template := &x509.Certificate{
		Subject: pkix.Name{
			Country:      []string{"US"},
			Organization: []string{"Oath Inc."},
			Locality:     []string{"Sunnyvale"},
			CommonName:   "testca.cameo.ouroath.com",
		},
		SerialNumber:          big.NewInt(1),
		PublicKeyAlgorithm:    pkAlgo,
		PublicKey:             priv.Public(),
		SignatureAlgorithm:    sigAlgo,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, priv.Public(), priv)
	if err != nil {
		return nil, nil, err
	}

	cert, err = x509.ParseCertificate(certBytes)
	return
}

// initMockSigner initializes a mock signer.
func initMockSigner(keyType crypki.PublicKeyAlgorithm, priv crypto.Signer, cert *x509.Certificate, isBad bool) *signer {
	s := &signer{
		x509CACerts: make(map[string]*x509.Certificate),
		sPool:       make(map[string]sPool),
	}

	sp := newMockSignerPool(isBad, keyType, priv)
	s.sPool[defaultIdentifier] = sp
	s.x509CACerts[defaultIdentifier] = cert
	return s
}

func TestGetSSHCertSigningKey(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	testcases := map[string]struct {
		ctx         context.Context
		identifier  string
		isBadSigner bool
		expectError bool
	}{
		"good-signer":         {ctx, defaultIdentifier, false, false},
		"bad-identifier":      {ctx, badIdentifier, false, true},
		"bad-signer":          {ctx, defaultIdentifier, true, true},
		"bad-request-timeout": {timeoutCtx, defaultIdentifier, false, true},
	}
	for label, tt := range testcases {
		label, tt := label, tt
		t.Run(label, func(t *testing.T) {
			t.Parallel()
			caPriv, caCert, err := createCAKeysAndCert(crypki.RSA)
			if err != nil {
				t.Fatalf("unable to create CA keys and certificate: %v", err)
			}
			signer := initMockSigner(crypki.RSA, caPriv, caCert, tt.isBadSigner)
			_, err = signer.GetSSHCertSigningKey(tt.ctx, tt.identifier)
			if err != nil != tt.expectError {
				t.Fatalf("got err: %v, expect err: %v", err, tt.expectError)
			}
		})
	}
}

func TestSignSSHCert(t *testing.T) {
	t.Parallel()
	rsakey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatalf("unable to generate RSA key: %v", err)
	}
	rsaPubKey, err := ssh.NewPublicKey(&rsakey.PublicKey)
	if err != nil {
		t.Fatalf("unable to create ssh RSA public key: %v", err)
	}
	eckey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("unable to generate EC key: %v", err)
	}
	ecPubKey, _ := ssh.NewPublicKey(&eckey.PublicKey)

	exts := make(map[string]string)
	exts["permit-pty"] = ""
	exts["permit-X11-forwarding"] = ""
	exts["permit-agent-forwarding"] = ""
	exts["permit-port-forwarding"] = ""
	exts["permit-user-rc"] = ""
	opts := make(map[string]string)
	opts["source-address"] = "10.11.12.13/32"
	opts["force-command"] = "ls -l"

	now := time.Now().Unix()

	userCertRSA := &ssh.Certificate{
		CertType:        ssh.UserCert,
		ValidPrincipals: []string{"alice"},
		Key:             rsaPubKey,
		KeyId:           "foo",
		ValidBefore:     uint64(now + 1000),
		ValidAfter:      uint64(now - 1000),
		Permissions: ssh.Permissions{
			Extensions:      exts,
			CriticalOptions: opts,
		},
	}
	userCertEC := &ssh.Certificate{
		CertType:        ssh.UserCert,
		ValidPrincipals: []string{"alice"},
		Key:             ecPubKey,
		KeyId:           "foo",
		ValidBefore:     uint64(now + 1000),
		ValidAfter:      uint64(now - 1000),
		Permissions: ssh.Permissions{
			Extensions:      exts,
			CriticalOptions: opts,
		},
	}
	hostCertRSA := &ssh.Certificate{
		CertType:    ssh.HostCert,
		ValidBefore: uint64(now + 1000),
		ValidAfter:  uint64(now - 1000),
		Key:         rsaPubKey,
	}
	hostCertEC := &ssh.Certificate{
		CertType:    ssh.HostCert,
		ValidBefore: uint64(now + 1000),
		ValidAfter:  uint64(now - 1000),
		Key:         ecPubKey,
	}
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	testcases := map[string]struct {
		ctx         context.Context
		cert        *ssh.Certificate
		keyType     crypki.PublicKeyAlgorithm
		identifier  string
		isBadSigner bool
		expectError bool
	}{
		"host-cert-rsa":             {ctx, hostCertRSA, crypki.RSA, defaultIdentifier, false, false},
		"host-cert-ec":              {ctx, hostCertEC, crypki.ECDSA, defaultIdentifier, false, false},
		"host-cert-bad-identifier":  {ctx, hostCertRSA, crypki.RSA, badIdentifier, false, true},
		"host-cert-bad-signer":      {ctx, hostCertRSA, crypki.RSA, defaultIdentifier, true, true},
		"user-cert-rsa":             {ctx, userCertRSA, crypki.RSA, defaultIdentifier, false, false},
		"user-cert-ec":              {ctx, userCertEC, crypki.ECDSA, defaultIdentifier, false, false},
		"user-cert-bad-identifier":  {ctx, userCertRSA, crypki.RSA, badIdentifier, false, true},
		"user-cert-bad-signer":      {ctx, userCertRSA, crypki.RSA, defaultIdentifier, true, true},
		"user-cert-request-timeout": {timeoutCtx, userCertRSA, crypki.RSA, defaultIdentifier, false, true},
	}
	for label, tt := range testcases {
		label, tt := label, tt
		t.Run(label, func(t *testing.T) {
			t.Parallel()
			caPriv, caCert, err := createCAKeysAndCert(tt.keyType)
			if err != nil {
				t.Fatalf("unable to create CA keys and certificate: %v", err)
			}
			signer := initMockSigner(tt.keyType, caPriv, caCert, tt.isBadSigner)
			data, err := signer.SignSSHCert(tt.ctx, tt.cert, tt.identifier)
			if err != nil != tt.expectError {
				t.Fatalf("got err: %v, expect err: %v", err, tt.expectError)
			}
			if err != nil {
				return
			}
			cc := &ssh.CertChecker{SupportedCriticalOptions: []string{"force-command"}}
			pk, _, _, _, err := ssh.ParseAuthorizedKey(data)
			if err != nil {
				t.Fatalf("unable to parse key from signed cert: %v", err)
			}
			cert, ok := pk.(*ssh.Certificate)
			if !ok {
				t.Fatal("not an ssh.Certificate type")
			}
			if err := cc.CheckCert("alice", cert); err != nil {
				t.Fatalf("check cert failed: %v", err)
			}
		})
	}
}

func TestGetX509CACert(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	testcases := map[string]struct {
		identifier  string
		isBadSigner bool
		expectError bool
	}{
		"good-signer":    {defaultIdentifier, false, false},
		"bad-identifier": {badIdentifier, false, true},
		"bad-signer":     {defaultIdentifier, true, false},
	}
	for label, tt := range testcases {
		label, tt := label, tt
		t.Run(label, func(t *testing.T) {
			t.Parallel()
			caPriv, caCert, err := createCAKeysAndCert(crypki.RSA)
			if err != nil {
				t.Fatalf("unable to create CA keys and certificate: %v", err)
			}
			signer := initMockSigner(crypki.RSA, caPriv, caCert, tt.isBadSigner)
			_, err = signer.GetX509CACert(ctx, tt.identifier)
			if err != nil != tt.expectError {
				t.Fatalf("got err: %v, expect err: %v", err, tt.expectError)
			}
		})
	}
}

func TestSignX509RSACert(t *testing.T) {
	t.Parallel()
	subject := pkix.Name{
		Country:            []string{"US"},
		Organization:       []string{"Foo"},
		OrganizationalUnit: []string{"FooUnit"},
		Locality:           []string{"Bar"},
		Province:           []string{"Baz"},
		CommonName:         "foo.bar.com",
	}
	rsakey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("unable to generate RSA key: %v", err)
	}
	certRSA := &x509.Certificate{
		Subject:               subject,
		SerialNumber:          big.NewInt(0),
		PublicKeyAlgorithm:    x509.RSA,
		PublicKey:             &rsakey.PublicKey,
		SignatureAlgorithm:    x509.SHA256WithRSA,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24),
		DNSNames:              []string{subject.CommonName},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	caPriv, caCert, err := createCAKeysAndCert(crypki.RSA)
	if err != nil {
		t.Fatalf("unable to create CA keys and certificate: %v", err)
	}
	cp := x509.NewCertPool()
	cp.AddCert(caCert)

	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	testcases := map[string]struct {
		ctx         context.Context
		cert        *x509.Certificate
		identifier  string
		isBadSigner bool
		expectError bool
	}{
		"cert-rsa-good-signer": {ctx, certRSA, defaultIdentifier, false, false},
		"cert-bad-identifier":  {ctx, certRSA, badIdentifier, false, true},
		"cert-bad-signer":      {ctx, certRSA, defaultIdentifier, true, true},
		"cert-request-timeout": {timeoutCtx, certRSA, defaultIdentifier, false, true},
	}
	for label, tt := range testcases {
		label, tt := label, tt
		t.Run(label, func(t *testing.T) {
			t.Parallel()
			signer := initMockSigner(crypki.RSA, caPriv, caCert, tt.isBadSigner)
			data, err := signer.SignX509Cert(tt.ctx, tt.cert, tt.identifier)
			if err != nil != tt.expectError {
				t.Fatalf("%s: got err: %v, expect err: %v", label, err, tt.expectError)
			}
			if err != nil {
				return
			}
			cd, _ := pem.Decode(data)
			cert, err := x509.ParseCertificate(cd.Bytes)
			if err != nil {
				t.Fatalf("unable to parse certificate: %v", err)
			}
			if _, err := cert.Verify(x509.VerifyOptions{
				Roots:     cp,
				KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			}); err != nil {
				t.Fatalf("failed to verify certificate: %v", err)
			}
			if !reflect.DeepEqual(cert.Issuer.String(), caCert.Issuer.String()) {
				t.Fatalf("issuer mismatch: got %q, want: %q", cert.Issuer, caCert.Issuer.String())
			}
			if !reflect.DeepEqual(cert.Subject.String(), subject.String()) {
				t.Fatalf("subject mismatch: got %q, want: %q", cert.Subject, subject)
			}
		})
	}
}

func TestSignX509ECCert(t *testing.T) {
	t.Parallel()
	subject := pkix.Name{
		Country:            []string{"US"},
		Organization:       []string{"Foo"},
		OrganizationalUnit: []string{"FooUnit"},
		Locality:           []string{"Bar"},
		Province:           []string{"Baz"},
		CommonName:         "foo.bar.com",
	}
	eckey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("unable to generate EC key: %v", err)
	}
	certEC := &x509.Certificate{
		Subject:               subject,
		SerialNumber:          big.NewInt(0),
		PublicKeyAlgorithm:    x509.ECDSA,
		PublicKey:             &eckey.PublicKey,
		SignatureAlgorithm:    x509.ECDSAWithSHA256,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24),
		DNSNames:              []string{subject.CommonName},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	caPriv, caCert, err := createCAKeysAndCert(crypki.ECDSA)
	if err != nil {
		t.Fatalf("unable to create CA keys and certificate: %v", err)
	}
	cp := x509.NewCertPool()
	cp.AddCert(caCert)

	ctx := context.Background()

	testcases := map[string]struct {
		ctx         context.Context
		cert        *x509.Certificate
		identifier  string
		isBadSigner bool
		expectError bool
	}{
		"cert-ec-good-signer": {ctx, certEC, defaultIdentifier, false, false},
		"cert-ec-bad-signer":  {ctx, certEC, badIdentifier, false, true},
	}
	for label, tt := range testcases {
		label, tt := label, tt
		t.Run(label, func(t *testing.T) {
			t.Parallel()
			signer := initMockSigner(crypki.ECDSA, caPriv, caCert, tt.isBadSigner)
			data, err := signer.SignX509Cert(tt.ctx, tt.cert, tt.identifier)
			if err != nil != tt.expectError {
				t.Fatalf("%s: got err: %v, expect err: %v", label, err, tt.expectError)
			}
			if err != nil {
				return
			}
			cd, _ := pem.Decode(data)
			cert, err := x509.ParseCertificate(cd.Bytes)
			if err != nil {
				t.Fatalf("unable to parse certificate: %v", err)
			}
			if _, err := cert.Verify(x509.VerifyOptions{
				Roots:     cp,
				KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
			}); err != nil {
				t.Fatalf("failed to verify certificate: %v", err)
			}
			if !reflect.DeepEqual(cert.Issuer.String(), caCert.Issuer.String()) {
				t.Fatalf("issuer mismatch: got %q, want: %q", cert.Issuer, caCert.Issuer.String())
			}
			if !reflect.DeepEqual(cert.Subject.String(), subject.String()) {
				t.Fatalf("subject mismatch: got %q, want: %q", cert.Subject, subject)
			}
		})
	}
}

func TestGetBlobSigningPublicKey(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	testcases := map[string]struct {
		ctx         context.Context
		identifier  string
		isBadSigner bool
		expectError bool
	}{
		"good-signer":         {ctx, defaultIdentifier, false, false},
		"bad-identifier":      {ctx, badIdentifier, false, true},
		"bad-signer":          {ctx, defaultIdentifier, true, true},
		"bad-request-timeout": {timeoutCtx, defaultIdentifier, false, true},
	}
	for label, tt := range testcases {
		label, tt := label, tt
		t.Run(label, func(t *testing.T) {
			t.Parallel()
			caPriv, caCert, err := createCAKeysAndCert(crypki.RSA)
			if err != nil {
				t.Fatalf("unable to create CA keys and certificate: %v", err)
			}
			signer := initMockSigner(crypki.RSA, caPriv, caCert, tt.isBadSigner)
			_, err = signer.GetBlobSigningPublicKey(tt.ctx, tt.identifier)
			if err != nil != tt.expectError {
				t.Fatalf("got err: %v, expect err: %v", err, tt.expectError)
			}
		})
	}
}

func TestSignBlob(t *testing.T) {
	t.Parallel()
	blob := []byte("good")
	goodDigestSHA224 := sha256.Sum224(blob)
	goodDigestSHA256 := sha256.Sum256(blob)
	goodDigestSHA384 := sha512.Sum384(blob)
	goodDigestSHA512 := sha512.Sum512(blob)

	caPriv, caCert, err := createCAKeysAndCert(crypki.RSA)
	if err != nil {
		t.Fatalf("unable to create CA keys and certificate: %v", err)
	}
	key, ok := caPriv.(*rsa.PrivateKey)
	if !ok {
		t.Fatal("unable to create RSA CA keys")
	}

	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	testcases := map[string]struct {
		ctx         context.Context
		digest      []byte
		opts        crypto.SignerOpts
		identifier  string
		isBadSigner bool
		expectError bool
	}{
		"good-SHA224":         {ctx, goodDigestSHA224[:], crypto.SHA224, defaultIdentifier, false, false},
		"good-SHA256":         {ctx, goodDigestSHA256[:], crypto.SHA256, defaultIdentifier, false, false},
		"good-SHA384":         {ctx, goodDigestSHA384[:], crypto.SHA384, defaultIdentifier, false, false},
		"good-SHA512":         {ctx, goodDigestSHA512[:], crypto.SHA512, defaultIdentifier, false, false},
		"bad-digest":          {ctx, []byte("bad digest"), crypto.SHA256, defaultIdentifier, false, true},
		"bad-wrong-hash":      {ctx, goodDigestSHA224[:], crypto.SHA256, defaultIdentifier, false, true},
		"bad-identifier":      {ctx, goodDigestSHA224[:], crypto.SHA256, badIdentifier, false, true},
		"bad-signer":          {ctx, goodDigestSHA224[:], crypto.SHA256, defaultIdentifier, true, true},
		"bad-request-timeout": {timeoutCtx, goodDigestSHA512[:], crypto.SHA512, defaultIdentifier, false, true},
	}
	for label, tt := range testcases {
		label, tt := label, tt
		t.Run(label, func(t *testing.T) {
			t.Parallel()
			signer := initMockSigner(crypki.RSA, caPriv, caCert, tt.isBadSigner)
			signature, err := signer.SignBlob(tt.ctx, tt.digest, tt.opts, tt.identifier)
			if err != nil != tt.expectError {
				t.Fatalf("got err: %v, expect err: %v", err, tt.expectError)
			}
			if err != nil {
				return
			}

			err = rsa.VerifyPKCS1v15(&key.PublicKey, tt.opts.(crypto.Hash), tt.digest, signature)
			if err != nil {
				t.Fatalf("failed to verify certificate: %v", err)
			}
		})
	}
}

func TestIsValidCertRequest(t *testing.T) {
	t.Parallel()
	subject := pkix.Name{
		Country:            []string{"US"},
		Organization:       []string{"Foo"},
		OrganizationalUnit: []string{"FooUnit"},
		Locality:           []string{"Bar"},
		Province:           []string{"Baz"},
		CommonName:         "foo.bar.com",
	}
	eckey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("unable to generate EC key: %v", err)
	}
	certEC := &x509.Certificate{
		Subject:               subject,
		SerialNumber:          big.NewInt(0),
		PublicKeyAlgorithm:    x509.ECDSA,
		PublicKey:             &eckey.PublicKey,
		SignatureAlgorithm:    x509.ECDSAWithSHA256,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24),
		DNSNames:              []string{subject.CommonName},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	tests := map[string]struct {
		sa   crypki.SignatureAlgorithm
		want bool
	}{
		"happy path":               {sa: crypki.ECDSAWithSHA256, want: true},
		"rsa-public-key-algo":      {sa: crypki.SHA256WithRSA, want: false},
		"incorrect-signature-algo": {sa: crypki.ECDSAWithSHA384, want: false},
	}
	for name, tt := range tests {
		name, tt := name, tt
		t.Run(name, func(t *testing.T) {
			got := isValidCertRequest(certEC, tt.sa)
			if got != tt.want {
				t.Fatalf("%s: got %v want %v", name, got, tt.want)
			}
		})
	}
}
