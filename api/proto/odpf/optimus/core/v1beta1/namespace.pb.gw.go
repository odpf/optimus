// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: odpf/optimus/core/v1beta1/namespace.proto

/*
Package optimus is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package optimus

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_NamespaceService_RegisterProjectNamespace_0(ctx context.Context, marshaler runtime.Marshaler, client NamespaceServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RegisterProjectNamespaceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["project_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "project_name")
	}

	protoReq.ProjectName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "project_name", err)
	}

	msg, err := client.RegisterProjectNamespace(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NamespaceService_RegisterProjectNamespace_0(ctx context.Context, marshaler runtime.Marshaler, server NamespaceServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RegisterProjectNamespaceRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["project_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "project_name")
	}

	protoReq.ProjectName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "project_name", err)
	}

	msg, err := server.RegisterProjectNamespace(ctx, &protoReq)
	return msg, metadata, err

}

func request_NamespaceService_ListProjectNamespaces_0(ctx context.Context, marshaler runtime.Marshaler, client NamespaceServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListProjectNamespacesRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["project_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "project_name")
	}

	protoReq.ProjectName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "project_name", err)
	}

	msg, err := client.ListProjectNamespaces(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_NamespaceService_ListProjectNamespaces_0(ctx context.Context, marshaler runtime.Marshaler, server NamespaceServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListProjectNamespacesRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["project_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "project_name")
	}

	protoReq.ProjectName, err = runtime.String(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "project_name", err)
	}

	msg, err := server.ListProjectNamespaces(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterNamespaceServiceHandlerServer registers the http handlers for service NamespaceService to "mux".
// UnaryRPC     :call NamespaceServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterNamespaceServiceHandlerFromEndpoint instead.
func RegisterNamespaceServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server NamespaceServiceServer) error {

	mux.Handle("POST", pattern_NamespaceService_RegisterProjectNamespace_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/odpf.optimus.core.v1beta1.NamespaceService/RegisterProjectNamespace", runtime.WithHTTPPathPattern("/v1beta1/project/{project_name}/namespace"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NamespaceService_RegisterProjectNamespace_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NamespaceService_RegisterProjectNamespace_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_NamespaceService_ListProjectNamespaces_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/odpf.optimus.core.v1beta1.NamespaceService/ListProjectNamespaces", runtime.WithHTTPPathPattern("/v1beta1/project/{project_name}/namespace"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_NamespaceService_ListProjectNamespaces_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NamespaceService_ListProjectNamespaces_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterNamespaceServiceHandlerFromEndpoint is same as RegisterNamespaceServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterNamespaceServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterNamespaceServiceHandler(ctx, mux, conn)
}

// RegisterNamespaceServiceHandler registers the http handlers for service NamespaceService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterNamespaceServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterNamespaceServiceHandlerClient(ctx, mux, NewNamespaceServiceClient(conn))
}

// RegisterNamespaceServiceHandlerClient registers the http handlers for service NamespaceService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "NamespaceServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "NamespaceServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "NamespaceServiceClient" to call the correct interceptors.
func RegisterNamespaceServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client NamespaceServiceClient) error {

	mux.Handle("POST", pattern_NamespaceService_RegisterProjectNamespace_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/odpf.optimus.core.v1beta1.NamespaceService/RegisterProjectNamespace", runtime.WithHTTPPathPattern("/v1beta1/project/{project_name}/namespace"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NamespaceService_RegisterProjectNamespace_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NamespaceService_RegisterProjectNamespace_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_NamespaceService_ListProjectNamespaces_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/odpf.optimus.core.v1beta1.NamespaceService/ListProjectNamespaces", runtime.WithHTTPPathPattern("/v1beta1/project/{project_name}/namespace"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_NamespaceService_ListProjectNamespaces_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_NamespaceService_ListProjectNamespaces_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_NamespaceService_RegisterProjectNamespace_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3}, []string{"v1beta1", "project", "project_name", "namespace"}, ""))

	pattern_NamespaceService_ListProjectNamespaces_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3}, []string{"v1beta1", "project", "project_name", "namespace"}, ""))
)

var (
	forward_NamespaceService_RegisterProjectNamespace_0 = runtime.ForwardResponseMessage

	forward_NamespaceService_ListProjectNamespaces_0 = runtime.ForwardResponseMessage
)
