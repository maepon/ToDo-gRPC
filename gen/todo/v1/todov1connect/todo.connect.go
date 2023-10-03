// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: todo/v1/todo.proto

package todov1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "example.com/todo/gen/todo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// ToDoServiceName is the fully-qualified name of the ToDoService service.
	ToDoServiceName = "todo.v1.ToDoService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ToDoServiceReadProcedure is the fully-qualified name of the ToDoService's Read RPC.
	ToDoServiceReadProcedure = "/todo.v1.ToDoService/Read"
	// ToDoServiceCreateProcedure is the fully-qualified name of the ToDoService's Create RPC.
	ToDoServiceCreateProcedure = "/todo.v1.ToDoService/Create"
	// ToDoServiceUpdateProcedure is the fully-qualified name of the ToDoService's Update RPC.
	ToDoServiceUpdateProcedure = "/todo.v1.ToDoService/Update"
	// ToDoServiceDeleteProcedure is the fully-qualified name of the ToDoService's Delete RPC.
	ToDoServiceDeleteProcedure = "/todo.v1.ToDoService/Delete"
)

// ToDoServiceClient is a client for the todo.v1.ToDoService service.
type ToDoServiceClient interface {
	Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error)
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewToDoServiceClient constructs a client for the todo.v1.ToDoService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewToDoServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ToDoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &toDoServiceClient{
		read: connect.NewClient[v1.ReadRequest, v1.ReadResponse](
			httpClient,
			baseURL+ToDoServiceReadProcedure,
			opts...,
		),
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+ToDoServiceCreateProcedure,
			opts...,
		),
		update: connect.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+ToDoServiceUpdateProcedure,
			opts...,
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+ToDoServiceDeleteProcedure,
			opts...,
		),
	}
}

// toDoServiceClient implements ToDoServiceClient.
type toDoServiceClient struct {
	read   *connect.Client[v1.ReadRequest, v1.ReadResponse]
	create *connect.Client[v1.CreateRequest, v1.CreateResponse]
	update *connect.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Read calls todo.v1.ToDoService.Read.
func (c *toDoServiceClient) Read(ctx context.Context, req *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	return c.read.CallUnary(ctx, req)
}

// Create calls todo.v1.ToDoService.Create.
func (c *toDoServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Update calls todo.v1.ToDoService.Update.
func (c *toDoServiceClient) Update(ctx context.Context, req *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls todo.v1.ToDoService.Delete.
func (c *toDoServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// ToDoServiceHandler is an implementation of the todo.v1.ToDoService service.
type ToDoServiceHandler interface {
	Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error)
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewToDoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewToDoServiceHandler(svc ToDoServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	toDoServiceReadHandler := connect.NewUnaryHandler(
		ToDoServiceReadProcedure,
		svc.Read,
		opts...,
	)
	toDoServiceCreateHandler := connect.NewUnaryHandler(
		ToDoServiceCreateProcedure,
		svc.Create,
		opts...,
	)
	toDoServiceUpdateHandler := connect.NewUnaryHandler(
		ToDoServiceUpdateProcedure,
		svc.Update,
		opts...,
	)
	toDoServiceDeleteHandler := connect.NewUnaryHandler(
		ToDoServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	return "/todo.v1.ToDoService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ToDoServiceReadProcedure:
			toDoServiceReadHandler.ServeHTTP(w, r)
		case ToDoServiceCreateProcedure:
			toDoServiceCreateHandler.ServeHTTP(w, r)
		case ToDoServiceUpdateProcedure:
			toDoServiceUpdateHandler.ServeHTTP(w, r)
		case ToDoServiceDeleteProcedure:
			toDoServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedToDoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedToDoServiceHandler struct{}

func (UnimplementedToDoServiceHandler) Read(context.Context, *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.ToDoService.Read is not implemented"))
}

func (UnimplementedToDoServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.ToDoService.Create is not implemented"))
}

func (UnimplementedToDoServiceHandler) Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.ToDoService.Update is not implemented"))
}

func (UnimplementedToDoServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("todo.v1.ToDoService.Delete is not implemented"))
}
