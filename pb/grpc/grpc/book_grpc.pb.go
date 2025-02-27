// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: grpc/book.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Book_Create_FullMethodName      = "/book.Book/Create"
	Book_Delete_FullMethodName      = "/book.Book/Delete"
	Book_GetBookById_FullMethodName = "/book.Book/GetBookById"
	Book_GetBooks_FullMethodName    = "/book.Book/GetBooks"
	Book_UpdateBook_FullMethodName  = "/book.Book/UpdateBook"
)

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Book service create and get book in library
type BookClient interface {
	Create(ctx context.Context, in *BookCreateRequest, opts ...grpc.CallOption) (*BookResponse, error)
	Delete(ctx context.Context, in *BookDeleteRequest, opts ...grpc.CallOption) (*BookDeleteResponse, error)
	GetBookById(ctx context.Context, in *BookGetBookByIdRequest, opts ...grpc.CallOption) (*BookResponse, error)
	GetBooks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[BookResponse], error)
	UpdateBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
}

type bookClient struct {
	cc grpc.ClientConnInterface
}

func NewBookClient(cc grpc.ClientConnInterface) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) Create(ctx context.Context, in *BookCreateRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, Book_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) Delete(ctx context.Context, in *BookDeleteRequest, opts ...grpc.CallOption) (*BookDeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookDeleteResponse)
	err := c.cc.Invoke(ctx, Book_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBookById(ctx context.Context, in *BookGetBookByIdRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, Book_GetBookById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBooks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[BookResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Book_ServiceDesc.Streams[0], Book_GetBooks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, BookResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Book_GetBooksClient = grpc.ServerStreamingClient[BookResponse]

func (c *bookClient) UpdateBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, Book_UpdateBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility.
//
// Book service create and get book in library
type BookServer interface {
	Create(context.Context, *BookCreateRequest) (*BookResponse, error)
	Delete(context.Context, *BookDeleteRequest) (*BookDeleteResponse, error)
	GetBookById(context.Context, *BookGetBookByIdRequest) (*BookResponse, error)
	GetBooks(*emptypb.Empty, grpc.ServerStreamingServer[BookResponse]) error
	UpdateBook(context.Context, *BookRequest) (*BookResponse, error)
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServer struct{}

func (UnimplementedBookServer) Create(context.Context, *BookCreateRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBookServer) Delete(context.Context, *BookDeleteRequest) (*BookDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBookServer) GetBookById(context.Context, *BookGetBookByIdRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBookById not implemented")
}
func (UnimplementedBookServer) GetBooks(*emptypb.Empty, grpc.ServerStreamingServer[BookResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedBookServer) UpdateBook(context.Context, *BookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedBookServer) mustEmbedUnimplementedBookServer() {}
func (UnimplementedBookServer) testEmbeddedByValue()              {}

// UnsafeBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServer will
// result in compilation errors.
type UnsafeBookServer interface {
	mustEmbedUnimplementedBookServer()
}

func RegisterBookServer(s grpc.ServiceRegistrar, srv BookServer) {
	// If the following call pancis, it indicates UnimplementedBookServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Book_ServiceDesc, srv)
}

func _Book_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).Create(ctx, req.(*BookCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).Delete(ctx, req.(*BookDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBookById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookGetBookByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBookById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBookById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBookById(ctx, req.(*BookGetBookByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServer).GetBooks(m, &grpc.GenericServerStream[emptypb.Empty, BookResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Book_GetBooksServer = grpc.ServerStreamingServer[BookResponse]

func _Book_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_UpdateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).UpdateBook(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Book_ServiceDesc is the grpc.ServiceDesc for Book service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Book_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Book_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Book_Delete_Handler,
		},
		{
			MethodName: "GetBookById",
			Handler:    _Book_GetBookById_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _Book_UpdateBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBooks",
			Handler:       _Book_GetBooks_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/book.proto",
}
