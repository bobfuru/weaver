// Code generated by "weaver generate". DO NOT EDIT.
//go:build !ignoreWeaverGen

package diverge

import (
	"context"
	"errors"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/runtime/codegen"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"reflect"
)

var _ codegen.LatestVersion = codegen.Version[[0][18]struct{}](`

ERROR: You generated this file with 'weaver generate' v0.18.0 (codegen
version v0.18.0). The generated code is incompatible with the version of the
github.com/ServiceWeaver/weaver module that you're using. The weaver module
version can be found in your go.mod file or by running the following command.

    go list -m github.com/ServiceWeaver/weaver

We recommend updating the weaver module and the 'weaver generate' command by
running the following.

    go get github.com/ServiceWeaver/weaver@latest
    go install github.com/ServiceWeaver/weaver/cmd/weaver@latest

Then, re-run 'weaver generate' and re-build your code. If the problem persists,
please file an issue at https://github.com/ServiceWeaver/weaver/issues.

`)

func init() {
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Errer",
		Iface: reflect.TypeOf((*Errer)(nil)).Elem(),
		Impl:  reflect.TypeOf(errer{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return errer_local_stub{impl: impl.(Errer), tracer: tracer, errMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Errer", Method: "Err", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return errer_client_stub{stub: stub, errMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Errer", Method: "Err", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return errer_server_stub{impl: impl.(Errer), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(reflect.Type, string, []reflect.Value) []reflect.Value) any {
			return errer_reflect_stub{caller: caller}
		},
		RefData: "",
	})
	codegen.Register(codegen.Registration{
		Name:  "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Pointer",
		Iface: reflect.TypeOf((*Pointer)(nil)).Elem(),
		Impl:  reflect.TypeOf(pointer{}),
		LocalStubFn: func(impl any, caller string, tracer trace.Tracer) any {
			return pointer_local_stub{impl: impl.(Pointer), tracer: tracer, getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Pointer", Method: "Get", Remote: false})}
		},
		ClientStubFn: func(stub codegen.Stub, caller string) any {
			return pointer_client_stub{stub: stub, getMetrics: codegen.MethodMetricsFor(codegen.MethodLabels{Caller: caller, Component: "github.com/ServiceWeaver/weaver/weavertest/internal/diverge/Pointer", Method: "Get", Remote: true})}
		},
		ServerStubFn: func(impl any, addLoad func(uint64, float64)) codegen.Server {
			return pointer_server_stub{impl: impl.(Pointer), addLoad: addLoad}
		},
		ReflectStubFn: func(caller func(reflect.Type, string, []reflect.Value) []reflect.Value) any {
			return pointer_reflect_stub{caller: caller}
		},
		RefData: "",
	})
}

// weaver.InstanceOf checks.
var _ weaver.InstanceOf[Errer] = (*errer)(nil)
var _ weaver.InstanceOf[Pointer] = (*pointer)(nil)

// weaver.Router checks.
var _ weaver.Unrouted = (*errer)(nil)
var _ weaver.Unrouted = (*pointer)(nil)

// Local stub implementations.

type errer_local_stub struct {
	impl       Errer
	tracer     trace.Tracer
	errMetrics *codegen.MethodMetrics
}

// Check that errer_local_stub implements the Errer interface.
var _ Errer = (*errer_local_stub)(nil)

func (s errer_local_stub) Err(ctx context.Context, a0 int) (err error) {
	// Update metrics.
	begin := s.errMetrics.Begin()
	defer func() { s.errMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "diverge.Errer.Err", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Err(ctx, a0)
}

type pointer_local_stub struct {
	impl       Pointer
	tracer     trace.Tracer
	getMetrics *codegen.MethodMetrics
}

// Check that pointer_local_stub implements the Pointer interface.
var _ Pointer = (*pointer_local_stub)(nil)

func (s pointer_local_stub) Get(ctx context.Context) (r0 Pair, err error) {
	// Update metrics.
	begin := s.getMetrics.Begin()
	defer func() { s.getMetrics.End(begin, err != nil, 0, 0) }()
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.tracer.Start(ctx, "diverge.Pointer.Get", trace.WithSpanKind(trace.SpanKindInternal))
		defer func() {
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			}
			span.End()
		}()
	}

	return s.impl.Get(ctx)
}

// Client stub implementations.

type errer_client_stub struct {
	stub       codegen.Stub
	errMetrics *codegen.MethodMetrics
}

// Check that errer_client_stub implements the Errer interface.
var _ Errer = (*errer_client_stub)(nil)

func (s errer_client_stub) Err(ctx context.Context, a0 int) (err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.errMetrics.Begin()
	defer func() { s.errMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "diverge.Errer.Err", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	// Preallocate a buffer of the right size.
	size := 0
	size += 8
	enc := codegen.NewEncoder()
	enc.Reset(size)

	// Encode arguments.
	enc.Int(a0)
	var shardKey uint64

	// Call the remote method.
	requestBytes = len(enc.Data())
	var results []byte
	results, err = s.stub.Run(ctx, 0, enc.Data(), shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	err = dec.Error()
	return
}

type pointer_client_stub struct {
	stub       codegen.Stub
	getMetrics *codegen.MethodMetrics
}

// Check that pointer_client_stub implements the Pointer interface.
var _ Pointer = (*pointer_client_stub)(nil)

func (s pointer_client_stub) Get(ctx context.Context) (r0 Pair, err error) {
	// Update metrics.
	var requestBytes, replyBytes int
	begin := s.getMetrics.Begin()
	defer func() { s.getMetrics.End(begin, err != nil, requestBytes, replyBytes) }()

	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		// Create a child span for this method.
		ctx, span = s.stub.Tracer().Start(ctx, "diverge.Pointer.Get", trace.WithSpanKind(trace.SpanKindClient))
	}

	defer func() {
		// Catch and return any panics detected during encoding/decoding/rpc.
		if err == nil {
			err = codegen.CatchPanics(recover())
			if err != nil {
				err = errors.Join(weaver.RemoteCallError, err)
			}
		}

		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()

	}()

	var shardKey uint64

	// Call the remote method.
	var results []byte
	results, err = s.stub.Run(ctx, 0, nil, shardKey)
	replyBytes = len(results)
	if err != nil {
		err = errors.Join(weaver.RemoteCallError, err)
		return
	}

	// Decode the results.
	dec := codegen.NewDecoder(results)
	(&r0).WeaverUnmarshal(dec)
	err = dec.Error()
	return
}

// Server stub implementations.

type errer_server_stub struct {
	impl    Errer
	addLoad func(key uint64, load float64)
}

// Check that errer_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*errer_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s errer_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Err":
		return s.err
	default:
		return nil
	}
}

func (s errer_server_stub) err(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// Decode arguments.
	dec := codegen.NewDecoder(args)
	var a0 int
	a0 = dec.Int()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	appErr := s.impl.Err(ctx, a0)

	// Encode the results.
	enc := codegen.NewEncoder()
	enc.Error(appErr)
	return enc.Data(), nil
}

type pointer_server_stub struct {
	impl    Pointer
	addLoad func(key uint64, load float64)
}

// Check that pointer_server_stub implements the codegen.Server interface.
var _ codegen.Server = (*pointer_server_stub)(nil)

// GetStubFn implements the codegen.Server interface.
func (s pointer_server_stub) GetStubFn(method string) func(ctx context.Context, args []byte) ([]byte, error) {
	switch method {
	case "Get":
		return s.get
	default:
		return nil
	}
}

func (s pointer_server_stub) get(ctx context.Context, args []byte) (res []byte, err error) {
	// Catch and return any panics detected during encoding/decoding/rpc.
	defer func() {
		if err == nil {
			err = codegen.CatchPanics(recover())
		}
	}()

	// TODO(rgrandl): The deferred function above will recover from panics in the
	// user code: fix this.
	// Call the local method.
	r0, appErr := s.impl.Get(ctx)

	// Encode the results.
	enc := codegen.NewEncoder()
	(r0).WeaverMarshal(enc)
	enc.Error(appErr)
	return enc.Data(), nil
}

// Reflect stub implementations.

type errer_reflect_stub struct {
	caller func(reflect.Type, string, []reflect.Value) []reflect.Value
}

// Check that errer_reflect_stub implements the Errer interface.
var _ Errer = (*errer_reflect_stub)(nil)

func (s errer_reflect_stub) Err(ctx context.Context, a0 int) (err error) {
	component := reflect.TypeOf((*Errer)(nil)).Elem()
	args := make([]reflect.Value, 2)
	args[0] = reflect.ValueOf(ctx)
	args[1] = reflect.ValueOf(a0)
	results := s.caller(component, "Err", args)
	if x := results[0].Interface(); x != nil {
		err = x.(error)
	}
	return
}

type pointer_reflect_stub struct {
	caller func(reflect.Type, string, []reflect.Value) []reflect.Value
}

// Check that pointer_reflect_stub implements the Pointer interface.
var _ Pointer = (*pointer_reflect_stub)(nil)

func (s pointer_reflect_stub) Get(ctx context.Context) (r0 Pair, err error) {
	component := reflect.TypeOf((*Pointer)(nil)).Elem()
	args := make([]reflect.Value, 1)
	args[0] = reflect.ValueOf(ctx)
	results := s.caller(component, "Get", args)
	r0 = results[0].Interface().(Pair)
	if x := results[1].Interface(); x != nil {
		err = x.(error)
	}
	return
}

// AutoMarshal implementations.

var _ codegen.AutoMarshal = (*Pair)(nil)

type __is_Pair[T ~struct {
	weaver.AutoMarshal
	X *int
	Y *int
}] struct{}

var _ __is_Pair[Pair]

func (x *Pair) WeaverMarshal(enc *codegen.Encoder) {
	if x == nil {
		panic(fmt.Errorf("Pair.WeaverMarshal: nil receiver"))
	}
	serviceweaver_enc_ptr_int_98a2a745(enc, x.X)
	serviceweaver_enc_ptr_int_98a2a745(enc, x.Y)
}

func (x *Pair) WeaverUnmarshal(dec *codegen.Decoder) {
	if x == nil {
		panic(fmt.Errorf("Pair.WeaverUnmarshal: nil receiver"))
	}
	x.X = serviceweaver_dec_ptr_int_98a2a745(dec)
	x.Y = serviceweaver_dec_ptr_int_98a2a745(dec)
}

func serviceweaver_enc_ptr_int_98a2a745(enc *codegen.Encoder, arg *int) {
	if arg == nil {
		enc.Bool(false)
	} else {
		enc.Bool(true)
		enc.Int(*arg)
	}
}

func serviceweaver_dec_ptr_int_98a2a745(dec *codegen.Decoder) *int {
	if !dec.Bool() {
		return nil
	}
	var res int
	res = dec.Int()
	return &res
}
