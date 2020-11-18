// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// SigningClient is the client API for Signing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SigningClient interface {
	// GetX509CertificateAvailableSigningKeys returns all available keys that can sign X509 certificates.
	GetX509CertificateAvailableSigningKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KeyMetas, error)
	// GetX509CACertificate returns the CA X509 certificate self-signed by the specified key.
	GetX509CACertificate(ctx context.Context, in *KeyMeta, opts ...grpc.CallOption) (*X509Certificate, error)
	// PostX509Certificate signs the given CSR using the specified key and returns a PEM encoded X509 certificate.
	PostX509Certificate(ctx context.Context, in *X509CertificateSigningRequest, opts ...grpc.CallOption) (*X509Certificate, error)
	// GetUserSSHCertificateAvailableSigningKeys returns all available keys that can sign user SSH certificates.
	GetUserSSHCertificateAvailableSigningKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KeyMetas, error)
	// GetUserSSHCertificateSigningKey returns the public signing key of the
	// specified key that signs the user ssh certificate.
	GetUserSSHCertificateSigningKey(ctx context.Context, in *KeyMeta, opts ...grpc.CallOption) (*SSHKey, error)
	// PostUserSSHCertificate signs the SSH user certificate given request fields using the specified key.
	PostUserSSHCertificate(ctx context.Context, in *SSHCertificateSigningRequest, opts ...grpc.CallOption) (*SSHKey, error)
	// GetHostSSHCertificateAvailableSigningKeys returns all available keys that can sign host SSH certificates.
	GetHostSSHCertificateAvailableSigningKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KeyMetas, error)
	// GetHostSSHCertificateSigningKey returns the public signing key of the
	// specified key that signs the host ssh certificate.
	GetHostSSHCertificateSigningKey(ctx context.Context, in *KeyMeta, opts ...grpc.CallOption) (*SSHKey, error)
	// PostHostSSHCertificate signs the SSH host certificate given request fields using the specified key.
	PostHostSSHCertificate(ctx context.Context, in *SSHCertificateSigningRequest, opts ...grpc.CallOption) (*SSHKey, error)
	// GetBlobAvailableSigningKeys returns all available keys that can sign
	GetBlobAvailableSigningKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KeyMetas, error)
	// GetBlobSigningKey returns the public signing key of the
	// specified key that signs the user's data.
	GetBlobSigningKey(ctx context.Context, in *KeyMeta, opts ...grpc.CallOption) (*PublicKey, error)
	// PostSignBlob signs the digest using the specified key.
	PostSignBlob(ctx context.Context, in *BlobSigningRequest, opts ...grpc.CallOption) (*Signature, error)
}

type signingClient struct {
	cc grpc.ClientConnInterface
}

func NewSigningClient(cc grpc.ClientConnInterface) SigningClient {
	return &signingClient{cc}
}

func (c *signingClient) GetX509CertificateAvailableSigningKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KeyMetas, error) {
	out := new(KeyMetas)
	err := c.cc.Invoke(ctx, "/v3.Signing/GetX509CertificateAvailableSigningKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) GetX509CACertificate(ctx context.Context, in *KeyMeta, opts ...grpc.CallOption) (*X509Certificate, error) {
	out := new(X509Certificate)
	err := c.cc.Invoke(ctx, "/v3.Signing/GetX509CACertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) PostX509Certificate(ctx context.Context, in *X509CertificateSigningRequest, opts ...grpc.CallOption) (*X509Certificate, error) {
	out := new(X509Certificate)
	err := c.cc.Invoke(ctx, "/v3.Signing/PostX509Certificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) GetUserSSHCertificateAvailableSigningKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KeyMetas, error) {
	out := new(KeyMetas)
	err := c.cc.Invoke(ctx, "/v3.Signing/GetUserSSHCertificateAvailableSigningKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) GetUserSSHCertificateSigningKey(ctx context.Context, in *KeyMeta, opts ...grpc.CallOption) (*SSHKey, error) {
	out := new(SSHKey)
	err := c.cc.Invoke(ctx, "/v3.Signing/GetUserSSHCertificateSigningKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) PostUserSSHCertificate(ctx context.Context, in *SSHCertificateSigningRequest, opts ...grpc.CallOption) (*SSHKey, error) {
	out := new(SSHKey)
	err := c.cc.Invoke(ctx, "/v3.Signing/PostUserSSHCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) GetHostSSHCertificateAvailableSigningKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KeyMetas, error) {
	out := new(KeyMetas)
	err := c.cc.Invoke(ctx, "/v3.Signing/GetHostSSHCertificateAvailableSigningKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) GetHostSSHCertificateSigningKey(ctx context.Context, in *KeyMeta, opts ...grpc.CallOption) (*SSHKey, error) {
	out := new(SSHKey)
	err := c.cc.Invoke(ctx, "/v3.Signing/GetHostSSHCertificateSigningKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) PostHostSSHCertificate(ctx context.Context, in *SSHCertificateSigningRequest, opts ...grpc.CallOption) (*SSHKey, error) {
	out := new(SSHKey)
	err := c.cc.Invoke(ctx, "/v3.Signing/PostHostSSHCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) GetBlobAvailableSigningKeys(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*KeyMetas, error) {
	out := new(KeyMetas)
	err := c.cc.Invoke(ctx, "/v3.Signing/GetBlobAvailableSigningKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) GetBlobSigningKey(ctx context.Context, in *KeyMeta, opts ...grpc.CallOption) (*PublicKey, error) {
	out := new(PublicKey)
	err := c.cc.Invoke(ctx, "/v3.Signing/GetBlobSigningKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingClient) PostSignBlob(ctx context.Context, in *BlobSigningRequest, opts ...grpc.CallOption) (*Signature, error) {
	out := new(Signature)
	err := c.cc.Invoke(ctx, "/v3.Signing/PostSignBlob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SigningServer is the server API for Signing service.
// All implementations must embed UnimplementedSigningServer
// for forward compatibility
type SigningServer interface {
	// GetX509CertificateAvailableSigningKeys returns all available keys that can sign X509 certificates.
	GetX509CertificateAvailableSigningKeys(context.Context, *empty.Empty) (*KeyMetas, error)
	// GetX509CACertificate returns the CA X509 certificate self-signed by the specified key.
	GetX509CACertificate(context.Context, *KeyMeta) (*X509Certificate, error)
	// PostX509Certificate signs the given CSR using the specified key and returns a PEM encoded X509 certificate.
	PostX509Certificate(context.Context, *X509CertificateSigningRequest) (*X509Certificate, error)
	// GetUserSSHCertificateAvailableSigningKeys returns all available keys that can sign user SSH certificates.
	GetUserSSHCertificateAvailableSigningKeys(context.Context, *empty.Empty) (*KeyMetas, error)
	// GetUserSSHCertificateSigningKey returns the public signing key of the
	// specified key that signs the user ssh certificate.
	GetUserSSHCertificateSigningKey(context.Context, *KeyMeta) (*SSHKey, error)
	// PostUserSSHCertificate signs the SSH user certificate given request fields using the specified key.
	PostUserSSHCertificate(context.Context, *SSHCertificateSigningRequest) (*SSHKey, error)
	// GetHostSSHCertificateAvailableSigningKeys returns all available keys that can sign host SSH certificates.
	GetHostSSHCertificateAvailableSigningKeys(context.Context, *empty.Empty) (*KeyMetas, error)
	// GetHostSSHCertificateSigningKey returns the public signing key of the
	// specified key that signs the host ssh certificate.
	GetHostSSHCertificateSigningKey(context.Context, *KeyMeta) (*SSHKey, error)
	// PostHostSSHCertificate signs the SSH host certificate given request fields using the specified key.
	PostHostSSHCertificate(context.Context, *SSHCertificateSigningRequest) (*SSHKey, error)
	// GetBlobAvailableSigningKeys returns all available keys that can sign
	GetBlobAvailableSigningKeys(context.Context, *empty.Empty) (*KeyMetas, error)
	// GetBlobSigningKey returns the public signing key of the
	// specified key that signs the user's data.
	GetBlobSigningKey(context.Context, *KeyMeta) (*PublicKey, error)
	// PostSignBlob signs the digest using the specified key.
	PostSignBlob(context.Context, *BlobSigningRequest) (*Signature, error)
	mustEmbedUnimplementedSigningServer()
}

// UnimplementedSigningServer must be embedded to have forward compatible implementations.
type UnimplementedSigningServer struct {
}

func (UnimplementedSigningServer) GetX509CertificateAvailableSigningKeys(context.Context, *empty.Empty) (*KeyMetas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetX509CertificateAvailableSigningKeys not implemented")
}
func (UnimplementedSigningServer) GetX509CACertificate(context.Context, *KeyMeta) (*X509Certificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetX509CACertificate not implemented")
}
func (UnimplementedSigningServer) PostX509Certificate(context.Context, *X509CertificateSigningRequest) (*X509Certificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostX509Certificate not implemented")
}
func (UnimplementedSigningServer) GetUserSSHCertificateAvailableSigningKeys(context.Context, *empty.Empty) (*KeyMetas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSSHCertificateAvailableSigningKeys not implemented")
}
func (UnimplementedSigningServer) GetUserSSHCertificateSigningKey(context.Context, *KeyMeta) (*SSHKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSSHCertificateSigningKey not implemented")
}
func (UnimplementedSigningServer) PostUserSSHCertificate(context.Context, *SSHCertificateSigningRequest) (*SSHKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostUserSSHCertificate not implemented")
}
func (UnimplementedSigningServer) GetHostSSHCertificateAvailableSigningKeys(context.Context, *empty.Empty) (*KeyMetas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostSSHCertificateAvailableSigningKeys not implemented")
}
func (UnimplementedSigningServer) GetHostSSHCertificateSigningKey(context.Context, *KeyMeta) (*SSHKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHostSSHCertificateSigningKey not implemented")
}
func (UnimplementedSigningServer) PostHostSSHCertificate(context.Context, *SSHCertificateSigningRequest) (*SSHKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostHostSSHCertificate not implemented")
}
func (UnimplementedSigningServer) GetBlobAvailableSigningKeys(context.Context, *empty.Empty) (*KeyMetas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlobAvailableSigningKeys not implemented")
}
func (UnimplementedSigningServer) GetBlobSigningKey(context.Context, *KeyMeta) (*PublicKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlobSigningKey not implemented")
}
func (UnimplementedSigningServer) PostSignBlob(context.Context, *BlobSigningRequest) (*Signature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSignBlob not implemented")
}
func (UnimplementedSigningServer) mustEmbedUnimplementedSigningServer() {}

// UnsafeSigningServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SigningServer will
// result in compilation errors.
type UnsafeSigningServer interface {
	mustEmbedUnimplementedSigningServer()
}

func RegisterSigningServer(s grpc.ServiceRegistrar, srv SigningServer) {
	s.RegisterService(&_Signing_serviceDesc, srv)
}

func _Signing_GetX509CertificateAvailableSigningKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).GetX509CertificateAvailableSigningKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/GetX509CertificateAvailableSigningKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).GetX509CertificateAvailableSigningKeys(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_GetX509CACertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).GetX509CACertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/GetX509CACertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).GetX509CACertificate(ctx, req.(*KeyMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_PostX509Certificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(X509CertificateSigningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).PostX509Certificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/PostX509Certificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).PostX509Certificate(ctx, req.(*X509CertificateSigningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_GetUserSSHCertificateAvailableSigningKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).GetUserSSHCertificateAvailableSigningKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/GetUserSSHCertificateAvailableSigningKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).GetUserSSHCertificateAvailableSigningKeys(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_GetUserSSHCertificateSigningKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).GetUserSSHCertificateSigningKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/GetUserSSHCertificateSigningKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).GetUserSSHCertificateSigningKey(ctx, req.(*KeyMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_PostUserSSHCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSHCertificateSigningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).PostUserSSHCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/PostUserSSHCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).PostUserSSHCertificate(ctx, req.(*SSHCertificateSigningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_GetHostSSHCertificateAvailableSigningKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).GetHostSSHCertificateAvailableSigningKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/GetHostSSHCertificateAvailableSigningKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).GetHostSSHCertificateAvailableSigningKeys(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_GetHostSSHCertificateSigningKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).GetHostSSHCertificateSigningKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/GetHostSSHCertificateSigningKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).GetHostSSHCertificateSigningKey(ctx, req.(*KeyMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_PostHostSSHCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSHCertificateSigningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).PostHostSSHCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/PostHostSSHCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).PostHostSSHCertificate(ctx, req.(*SSHCertificateSigningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_GetBlobAvailableSigningKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).GetBlobAvailableSigningKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/GetBlobAvailableSigningKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).GetBlobAvailableSigningKeys(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_GetBlobSigningKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyMeta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).GetBlobSigningKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/GetBlobSigningKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).GetBlobSigningKey(ctx, req.(*KeyMeta))
	}
	return interceptor(ctx, in, info, handler)
}

func _Signing_PostSignBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlobSigningRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningServer).PostSignBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v3.Signing/PostSignBlob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningServer).PostSignBlob(ctx, req.(*BlobSigningRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Signing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v3.Signing",
	HandlerType: (*SigningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetX509CertificateAvailableSigningKeys",
			Handler:    _Signing_GetX509CertificateAvailableSigningKeys_Handler,
		},
		{
			MethodName: "GetX509CACertificate",
			Handler:    _Signing_GetX509CACertificate_Handler,
		},
		{
			MethodName: "PostX509Certificate",
			Handler:    _Signing_PostX509Certificate_Handler,
		},
		{
			MethodName: "GetUserSSHCertificateAvailableSigningKeys",
			Handler:    _Signing_GetUserSSHCertificateAvailableSigningKeys_Handler,
		},
		{
			MethodName: "GetUserSSHCertificateSigningKey",
			Handler:    _Signing_GetUserSSHCertificateSigningKey_Handler,
		},
		{
			MethodName: "PostUserSSHCertificate",
			Handler:    _Signing_PostUserSSHCertificate_Handler,
		},
		{
			MethodName: "GetHostSSHCertificateAvailableSigningKeys",
			Handler:    _Signing_GetHostSSHCertificateAvailableSigningKeys_Handler,
		},
		{
			MethodName: "GetHostSSHCertificateSigningKey",
			Handler:    _Signing_GetHostSSHCertificateSigningKey_Handler,
		},
		{
			MethodName: "PostHostSSHCertificate",
			Handler:    _Signing_PostHostSSHCertificate_Handler,
		},
		{
			MethodName: "GetBlobAvailableSigningKeys",
			Handler:    _Signing_GetBlobAvailableSigningKeys_Handler,
		},
		{
			MethodName: "GetBlobSigningKey",
			Handler:    _Signing_GetBlobSigningKey_Handler,
		},
		{
			MethodName: "PostSignBlob",
			Handler:    _Signing_PostSignBlob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sign.proto",
}
