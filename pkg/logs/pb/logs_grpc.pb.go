// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: logs.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LogsServiceClient is the client API for LogsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogsServiceClient interface {
	Logs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (LogsService_LogsClient, error)
}

type logsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogsServiceClient(cc grpc.ClientConnInterface) LogsServiceClient {
	return &logsServiceClient{cc}
}

func (c *logsServiceClient) Logs(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (LogsService_LogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &LogsService_ServiceDesc.Streams[0], "/logs.LogsService/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &logsServiceLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogsService_LogsClient interface {
	Recv() (*LogResponse, error)
	grpc.ClientStream
}

type logsServiceLogsClient struct {
	grpc.ClientStream
}

func (x *logsServiceLogsClient) Recv() (*LogResponse, error) {
	m := new(LogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LogsServiceServer is the server API for LogsService service.
// All implementations must embed UnimplementedLogsServiceServer
// for forward compatibility
type LogsServiceServer interface {
	Logs(*LogRequest, LogsService_LogsServer) error
	mustEmbedUnimplementedLogsServiceServer()
}

// UnimplementedLogsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLogsServiceServer struct {
}

func (UnimplementedLogsServiceServer) Logs(*LogRequest, LogsService_LogsServer) error {
	return status.Errorf(codes.Unimplemented, "method Logs not implemented")
}
func (UnimplementedLogsServiceServer) mustEmbedUnimplementedLogsServiceServer() {}

// UnsafeLogsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogsServiceServer will
// result in compilation errors.
type UnsafeLogsServiceServer interface {
	mustEmbedUnimplementedLogsServiceServer()
}

func RegisterLogsServiceServer(s grpc.ServiceRegistrar, srv LogsServiceServer) {
	s.RegisterService(&LogsService_ServiceDesc, srv)
}

func _LogsService_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogsServiceServer).Logs(m, &logsServiceLogsServer{stream})
}

type LogsService_LogsServer interface {
	Send(*LogResponse) error
	grpc.ServerStream
}

type logsServiceLogsServer struct {
	grpc.ServerStream
}

func (x *logsServiceLogsServer) Send(m *LogResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LogsService_ServiceDesc is the grpc.ServiceDesc for LogsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logs.LogsService",
	HandlerType: (*LogsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _LogsService_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "logs.proto",
}
