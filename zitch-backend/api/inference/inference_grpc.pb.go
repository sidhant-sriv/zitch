// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package inference

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

// InferenceClient is the client API for Inference service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InferenceClient interface {
	Predict(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Prediction, error)
}

type inferenceClient struct {
	cc grpc.ClientConnInterface
}

func NewInferenceClient(cc grpc.ClientConnInterface) InferenceClient {
	return &inferenceClient{cc}
}

func (c *inferenceClient) Predict(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Prediction, error) {
	out := new(Prediction)
	err := c.cc.Invoke(ctx, "/inference.Inference/Predict", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InferenceServer is the server API for Inference service.
// All implementations must embed UnimplementedInferenceServer
// for forward compatibility
type InferenceServer interface {
	Predict(context.Context, *Image) (*Prediction, error)
	mustEmbedUnimplementedInferenceServer()
}

// UnimplementedInferenceServer must be embedded to have forward compatible implementations.
type UnimplementedInferenceServer struct {
}

func (UnimplementedInferenceServer) Predict(context.Context, *Image) (*Prediction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Predict not implemented")
}
func (UnimplementedInferenceServer) mustEmbedUnimplementedInferenceServer() {}

// UnsafeInferenceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InferenceServer will
// result in compilation errors.
type UnsafeInferenceServer interface {
	mustEmbedUnimplementedInferenceServer()
}

func RegisterInferenceServer(s grpc.ServiceRegistrar, srv InferenceServer) {
	s.RegisterService(&Inference_ServiceDesc, srv)
}

func _Inference_Predict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InferenceServer).Predict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inference.Inference/Predict",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InferenceServer).Predict(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

// Inference_ServiceDesc is the grpc.ServiceDesc for Inference service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inference_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inference.Inference",
	HandlerType: (*InferenceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Predict",
			Handler:    _Inference_Predict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inference.proto",
}
