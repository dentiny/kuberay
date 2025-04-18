// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_client

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComputeTemplateServiceClient is the client API for ComputeTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComputeTemplateServiceClient interface {
	// Creates a new compute template.
	CreateComputeTemplate(ctx context.Context, in *CreateComputeTemplateRequest, opts ...grpc.CallOption) (*ComputeTemplate, error)
	// Finds a specific compute template by its name and namespace.
	GetComputeTemplate(ctx context.Context, in *GetComputeTemplateRequest, opts ...grpc.CallOption) (*ComputeTemplate, error)
	// Finds all compute templates in a given namespace.
	ListComputeTemplates(ctx context.Context, in *ListComputeTemplatesRequest, opts ...grpc.CallOption) (*ListComputeTemplatesResponse, error)
	// Finds all compute templates in all namespaces.
	ListAllComputeTemplates(ctx context.Context, in *ListAllComputeTemplatesRequest, opts ...grpc.CallOption) (*ListAllComputeTemplatesResponse, error)
	// Deletes a compute template by its name and namespace
	DeleteComputeTemplate(ctx context.Context, in *DeleteComputeTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type computeTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewComputeTemplateServiceClient(cc grpc.ClientConnInterface) ComputeTemplateServiceClient {
	return &computeTemplateServiceClient{cc}
}

func (c *computeTemplateServiceClient) CreateComputeTemplate(ctx context.Context, in *CreateComputeTemplateRequest, opts ...grpc.CallOption) (*ComputeTemplate, error) {
	out := new(ComputeTemplate)
	err := c.cc.Invoke(ctx, "/proto.ComputeTemplateService/CreateComputeTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeTemplateServiceClient) GetComputeTemplate(ctx context.Context, in *GetComputeTemplateRequest, opts ...grpc.CallOption) (*ComputeTemplate, error) {
	out := new(ComputeTemplate)
	err := c.cc.Invoke(ctx, "/proto.ComputeTemplateService/GetComputeTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeTemplateServiceClient) ListComputeTemplates(ctx context.Context, in *ListComputeTemplatesRequest, opts ...grpc.CallOption) (*ListComputeTemplatesResponse, error) {
	out := new(ListComputeTemplatesResponse)
	err := c.cc.Invoke(ctx, "/proto.ComputeTemplateService/ListComputeTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeTemplateServiceClient) ListAllComputeTemplates(ctx context.Context, in *ListAllComputeTemplatesRequest, opts ...grpc.CallOption) (*ListAllComputeTemplatesResponse, error) {
	out := new(ListAllComputeTemplatesResponse)
	err := c.cc.Invoke(ctx, "/proto.ComputeTemplateService/ListAllComputeTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computeTemplateServiceClient) DeleteComputeTemplate(ctx context.Context, in *DeleteComputeTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ComputeTemplateService/DeleteComputeTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeTemplateServiceServer is the server API for ComputeTemplateService service.
// All implementations must embed UnimplementedComputeTemplateServiceServer
// for forward compatibility
type ComputeTemplateServiceServer interface {
	// Creates a new compute template.
	CreateComputeTemplate(context.Context, *CreateComputeTemplateRequest) (*ComputeTemplate, error)
	// Finds a specific compute template by its name and namespace.
	GetComputeTemplate(context.Context, *GetComputeTemplateRequest) (*ComputeTemplate, error)
	// Finds all compute templates in a given namespace.
	ListComputeTemplates(context.Context, *ListComputeTemplatesRequest) (*ListComputeTemplatesResponse, error)
	// Finds all compute templates in all namespaces.
	ListAllComputeTemplates(context.Context, *ListAllComputeTemplatesRequest) (*ListAllComputeTemplatesResponse, error)
	// Deletes a compute template by its name and namespace
	DeleteComputeTemplate(context.Context, *DeleteComputeTemplateRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedComputeTemplateServiceServer()
}

// UnimplementedComputeTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedComputeTemplateServiceServer struct {
}

func (UnimplementedComputeTemplateServiceServer) CreateComputeTemplate(context.Context, *CreateComputeTemplateRequest) (*ComputeTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComputeTemplate not implemented")
}
func (UnimplementedComputeTemplateServiceServer) GetComputeTemplate(context.Context, *GetComputeTemplateRequest) (*ComputeTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComputeTemplate not implemented")
}
func (UnimplementedComputeTemplateServiceServer) ListComputeTemplates(context.Context, *ListComputeTemplatesRequest) (*ListComputeTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListComputeTemplates not implemented")
}
func (UnimplementedComputeTemplateServiceServer) ListAllComputeTemplates(context.Context, *ListAllComputeTemplatesRequest) (*ListAllComputeTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllComputeTemplates not implemented")
}
func (UnimplementedComputeTemplateServiceServer) DeleteComputeTemplate(context.Context, *DeleteComputeTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComputeTemplate not implemented")
}
func (UnimplementedComputeTemplateServiceServer) mustEmbedUnimplementedComputeTemplateServiceServer() {
}

// UnsafeComputeTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComputeTemplateServiceServer will
// result in compilation errors.
type UnsafeComputeTemplateServiceServer interface {
	mustEmbedUnimplementedComputeTemplateServiceServer()
}

func RegisterComputeTemplateServiceServer(s grpc.ServiceRegistrar, srv ComputeTemplateServiceServer) {
	s.RegisterService(&ComputeTemplateService_ServiceDesc, srv)
}

func _ComputeTemplateService_CreateComputeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateComputeTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeTemplateServiceServer).CreateComputeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ComputeTemplateService/CreateComputeTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeTemplateServiceServer).CreateComputeTemplate(ctx, req.(*CreateComputeTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeTemplateService_GetComputeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComputeTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeTemplateServiceServer).GetComputeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ComputeTemplateService/GetComputeTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeTemplateServiceServer).GetComputeTemplate(ctx, req.(*GetComputeTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeTemplateService_ListComputeTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListComputeTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeTemplateServiceServer).ListComputeTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ComputeTemplateService/ListComputeTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeTemplateServiceServer).ListComputeTemplates(ctx, req.(*ListComputeTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeTemplateService_ListAllComputeTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllComputeTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeTemplateServiceServer).ListAllComputeTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ComputeTemplateService/ListAllComputeTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeTemplateServiceServer).ListAllComputeTemplates(ctx, req.(*ListAllComputeTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputeTemplateService_DeleteComputeTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteComputeTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeTemplateServiceServer).DeleteComputeTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ComputeTemplateService/DeleteComputeTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeTemplateServiceServer).DeleteComputeTemplate(ctx, req.(*DeleteComputeTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ComputeTemplateService_ServiceDesc is the grpc.ServiceDesc for ComputeTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComputeTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ComputeTemplateService",
	HandlerType: (*ComputeTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateComputeTemplate",
			Handler:    _ComputeTemplateService_CreateComputeTemplate_Handler,
		},
		{
			MethodName: "GetComputeTemplate",
			Handler:    _ComputeTemplateService_GetComputeTemplate_Handler,
		},
		{
			MethodName: "ListComputeTemplates",
			Handler:    _ComputeTemplateService_ListComputeTemplates_Handler,
		},
		{
			MethodName: "ListAllComputeTemplates",
			Handler:    _ComputeTemplateService_ListAllComputeTemplates_Handler,
		},
		{
			MethodName: "DeleteComputeTemplate",
			Handler:    _ComputeTemplateService_DeleteComputeTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config.proto",
}

// ImageTemplateServiceClient is the client API for ImageTemplateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageTemplateServiceClient interface {
	// Not implemented. Creates a new ImageTemplate.
	CreateImageTemplate(ctx context.Context, in *CreateImageTemplateRequest, opts ...grpc.CallOption) (*ImageTemplate, error)
	// Not implemented. Finds a specific ImageTemplate by ID.
	GetImageTemplate(ctx context.Context, in *GetImageTemplateRequest, opts ...grpc.CallOption) (*ImageTemplate, error)
	// Not Implemented. Finds all ImageTemplates.
	ListImageTemplates(ctx context.Context, in *ListImageTemplatesRequest, opts ...grpc.CallOption) (*ListImageTemplatesResponse, error)
	// Not implemented. Deletes an ImageTemplate.
	DeleteImageTemplate(ctx context.Context, in *DeleteImageTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type imageTemplateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageTemplateServiceClient(cc grpc.ClientConnInterface) ImageTemplateServiceClient {
	return &imageTemplateServiceClient{cc}
}

func (c *imageTemplateServiceClient) CreateImageTemplate(ctx context.Context, in *CreateImageTemplateRequest, opts ...grpc.CallOption) (*ImageTemplate, error) {
	out := new(ImageTemplate)
	err := c.cc.Invoke(ctx, "/proto.ImageTemplateService/CreateImageTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageTemplateServiceClient) GetImageTemplate(ctx context.Context, in *GetImageTemplateRequest, opts ...grpc.CallOption) (*ImageTemplate, error) {
	out := new(ImageTemplate)
	err := c.cc.Invoke(ctx, "/proto.ImageTemplateService/GetImageTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageTemplateServiceClient) ListImageTemplates(ctx context.Context, in *ListImageTemplatesRequest, opts ...grpc.CallOption) (*ListImageTemplatesResponse, error) {
	out := new(ListImageTemplatesResponse)
	err := c.cc.Invoke(ctx, "/proto.ImageTemplateService/ListImageTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageTemplateServiceClient) DeleteImageTemplate(ctx context.Context, in *DeleteImageTemplateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.ImageTemplateService/DeleteImageTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageTemplateServiceServer is the server API for ImageTemplateService service.
// All implementations must embed UnimplementedImageTemplateServiceServer
// for forward compatibility
type ImageTemplateServiceServer interface {
	// Not implemented. Creates a new ImageTemplate.
	CreateImageTemplate(context.Context, *CreateImageTemplateRequest) (*ImageTemplate, error)
	// Not implemented. Finds a specific ImageTemplate by ID.
	GetImageTemplate(context.Context, *GetImageTemplateRequest) (*ImageTemplate, error)
	// Not Implemented. Finds all ImageTemplates.
	ListImageTemplates(context.Context, *ListImageTemplatesRequest) (*ListImageTemplatesResponse, error)
	// Not implemented. Deletes an ImageTemplate.
	DeleteImageTemplate(context.Context, *DeleteImageTemplateRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedImageTemplateServiceServer()
}

// UnimplementedImageTemplateServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageTemplateServiceServer struct {
}

func (UnimplementedImageTemplateServiceServer) CreateImageTemplate(context.Context, *CreateImageTemplateRequest) (*ImageTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateImageTemplate not implemented")
}
func (UnimplementedImageTemplateServiceServer) GetImageTemplate(context.Context, *GetImageTemplateRequest) (*ImageTemplate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageTemplate not implemented")
}
func (UnimplementedImageTemplateServiceServer) ListImageTemplates(context.Context, *ListImageTemplatesRequest) (*ListImageTemplatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImageTemplates not implemented")
}
func (UnimplementedImageTemplateServiceServer) DeleteImageTemplate(context.Context, *DeleteImageTemplateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImageTemplate not implemented")
}
func (UnimplementedImageTemplateServiceServer) mustEmbedUnimplementedImageTemplateServiceServer() {}

// UnsafeImageTemplateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageTemplateServiceServer will
// result in compilation errors.
type UnsafeImageTemplateServiceServer interface {
	mustEmbedUnimplementedImageTemplateServiceServer()
}

func RegisterImageTemplateServiceServer(s grpc.ServiceRegistrar, srv ImageTemplateServiceServer) {
	s.RegisterService(&ImageTemplateService_ServiceDesc, srv)
}

func _ImageTemplateService_CreateImageTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageTemplateServiceServer).CreateImageTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImageTemplateService/CreateImageTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageTemplateServiceServer).CreateImageTemplate(ctx, req.(*CreateImageTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageTemplateService_GetImageTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageTemplateServiceServer).GetImageTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImageTemplateService/GetImageTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageTemplateServiceServer).GetImageTemplate(ctx, req.(*GetImageTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageTemplateService_ListImageTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImageTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageTemplateServiceServer).ListImageTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImageTemplateService/ListImageTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageTemplateServiceServer).ListImageTemplates(ctx, req.(*ListImageTemplatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageTemplateService_DeleteImageTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageTemplateServiceServer).DeleteImageTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ImageTemplateService/DeleteImageTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageTemplateServiceServer).DeleteImageTemplate(ctx, req.(*DeleteImageTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageTemplateService_ServiceDesc is the grpc.ServiceDesc for ImageTemplateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageTemplateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ImageTemplateService",
	HandlerType: (*ImageTemplateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateImageTemplate",
			Handler:    _ImageTemplateService_CreateImageTemplate_Handler,
		},
		{
			MethodName: "GetImageTemplate",
			Handler:    _ImageTemplateService_GetImageTemplate_Handler,
		},
		{
			MethodName: "ListImageTemplates",
			Handler:    _ImageTemplateService_ListImageTemplates_Handler,
		},
		{
			MethodName: "DeleteImageTemplate",
			Handler:    _ImageTemplateService_DeleteImageTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config.proto",
}
