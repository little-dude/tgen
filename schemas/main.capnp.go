package schemas

// AUTO GENERATED - DO NOT EDIT

import (
	context "golang.org/x/net/context"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	server "zombiezen.com/go/capnproto2/server"
)

type Controller struct{ Client capnp.Client }

func (c Controller) GetPorts(ctx context.Context, params func(Controller_getPorts_Params) error, opts ...capnp.CallOption) Controller_getPorts_Results_Promise {
	if c.Client == nil {
		return Controller_getPorts_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      0,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "getPorts",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Controller_getPorts_Params{Struct: s}) }
	}
	return Controller_getPorts_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Controller) ListStreams(ctx context.Context, params func(Controller_listStreams_Params) error, opts ...capnp.CallOption) Controller_listStreams_Results_Promise {
	if c.Client == nil {
		return Controller_listStreams_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      1,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "listStreams",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Controller_listStreams_Params{Struct: s}) }
	}
	return Controller_listStreams_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Controller) FetchStream(ctx context.Context, params func(Controller_fetchStream_Params) error, opts ...capnp.CallOption) Controller_fetchStream_Results_Promise {
	if c.Client == nil {
		return Controller_fetchStream_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      2,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "fetchStream",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Controller_fetchStream_Params{Struct: s}) }
	}
	return Controller_fetchStream_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Controller) SaveStream(ctx context.Context, params func(Controller_saveStream_Params) error, opts ...capnp.CallOption) Controller_saveStream_Results_Promise {
	if c.Client == nil {
		return Controller_saveStream_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      3,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "saveStream",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Controller_saveStream_Params{Struct: s}) }
	}
	return Controller_saveStream_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Controller) DeleteStream(ctx context.Context, params func(Controller_deleteStream_Params) error, opts ...capnp.CallOption) Controller_deleteStream_Results_Promise {
	if c.Client == nil {
		return Controller_deleteStream_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      4,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "deleteStream",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 8, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Controller_deleteStream_Params{Struct: s}) }
	}
	return Controller_deleteStream_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Controller_Server interface {
	GetPorts(Controller_getPorts) error

	ListStreams(Controller_listStreams) error

	FetchStream(Controller_fetchStream) error

	SaveStream(Controller_saveStream) error

	DeleteStream(Controller_deleteStream) error
}

func Controller_ServerToClient(s Controller_Server) Controller {
	c, _ := s.(server.Closer)
	return Controller{Client: server.New(Controller_Methods(nil, s), c)}
}

func Controller_Methods(methods []server.Method, s Controller_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 5)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      0,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "getPorts",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Controller_getPorts{c, opts, Controller_getPorts_Params{Struct: p}, Controller_getPorts_Results{Struct: r}}
			return s.GetPorts(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      1,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "listStreams",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Controller_listStreams{c, opts, Controller_listStreams_Params{Struct: p}, Controller_listStreams_Results{Struct: r}}
			return s.ListStreams(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      2,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "fetchStream",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Controller_fetchStream{c, opts, Controller_fetchStream_Params{Struct: p}, Controller_fetchStream_Results{Struct: r}}
			return s.FetchStream(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      3,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "saveStream",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Controller_saveStream{c, opts, Controller_saveStream_Params{Struct: p}, Controller_saveStream_Results{Struct: r}}
			return s.SaveStream(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 8, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xde386d159ea19675,
			MethodID:      4,
			InterfaceName: "main.capnp:Controller",
			MethodName:    "deleteStream",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Controller_deleteStream{c, opts, Controller_deleteStream_Params{Struct: p}, Controller_deleteStream_Results{Struct: r}}
			return s.DeleteStream(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// Controller_getPorts holds the arguments for a server call to Controller.getPorts.
type Controller_getPorts struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Controller_getPorts_Params
	Results Controller_getPorts_Results
}

// Controller_listStreams holds the arguments for a server call to Controller.listStreams.
type Controller_listStreams struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Controller_listStreams_Params
	Results Controller_listStreams_Results
}

// Controller_fetchStream holds the arguments for a server call to Controller.fetchStream.
type Controller_fetchStream struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Controller_fetchStream_Params
	Results Controller_fetchStream_Results
}

// Controller_saveStream holds the arguments for a server call to Controller.saveStream.
type Controller_saveStream struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Controller_saveStream_Params
	Results Controller_saveStream_Results
}

// Controller_deleteStream holds the arguments for a server call to Controller.deleteStream.
type Controller_deleteStream struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Controller_deleteStream_Params
	Results Controller_deleteStream_Results
}

type Controller_getPorts_Params struct{ capnp.Struct }

func NewController_getPorts_Params(s *capnp.Segment) (Controller_getPorts_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Controller_getPorts_Params{}, err
	}
	return Controller_getPorts_Params{st}, nil
}

func NewRootController_getPorts_Params(s *capnp.Segment) (Controller_getPorts_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Controller_getPorts_Params{}, err
	}
	return Controller_getPorts_Params{st}, nil
}

func ReadRootController_getPorts_Params(msg *capnp.Message) (Controller_getPorts_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_getPorts_Params{}, err
	}
	return Controller_getPorts_Params{root.Struct()}, nil
}

// Controller_getPorts_Params_List is a list of Controller_getPorts_Params.
type Controller_getPorts_Params_List struct{ capnp.List }

// NewController_getPorts_Params creates a new list of Controller_getPorts_Params.
func NewController_getPorts_Params_List(s *capnp.Segment, sz int32) (Controller_getPorts_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Controller_getPorts_Params_List{}, err
	}
	return Controller_getPorts_Params_List{l}, nil
}

func (s Controller_getPorts_Params_List) At(i int) Controller_getPorts_Params {
	return Controller_getPorts_Params{s.List.Struct(i)}
}
func (s Controller_getPorts_Params_List) Set(i int, v Controller_getPorts_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_getPorts_Params_Promise is a wrapper for a Controller_getPorts_Params promised by a client call.
type Controller_getPorts_Params_Promise struct{ *capnp.Pipeline }

func (p Controller_getPorts_Params_Promise) Struct() (Controller_getPorts_Params, error) {
	s, err := p.Pipeline.Struct()
	return Controller_getPorts_Params{s}, err
}

type Controller_getPorts_Results struct{ capnp.Struct }

func NewController_getPorts_Results(s *capnp.Segment) (Controller_getPorts_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Controller_getPorts_Results{}, err
	}
	return Controller_getPorts_Results{st}, nil
}

func NewRootController_getPorts_Results(s *capnp.Segment) (Controller_getPorts_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Controller_getPorts_Results{}, err
	}
	return Controller_getPorts_Results{st}, nil
}

func ReadRootController_getPorts_Results(msg *capnp.Message) (Controller_getPorts_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_getPorts_Results{}, err
	}
	return Controller_getPorts_Results{root.Struct()}, nil
}
func (s Controller_getPorts_Results) Ports() (capnp.PointerList, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return capnp.PointerList{}, err
	}
	return capnp.PointerList{List: p.List()}, nil
}

func (s Controller_getPorts_Results) HasPorts() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Controller_getPorts_Results) SetPorts(v capnp.PointerList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewPorts sets the ports field to a newly
// allocated capnp.PointerList, preferring placement in s's segment.
func (s Controller_getPorts_Results) NewPorts(n int32) (capnp.PointerList, error) {
	l, err := capnp.NewPointerList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.PointerList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Controller_getPorts_Results_List is a list of Controller_getPorts_Results.
type Controller_getPorts_Results_List struct{ capnp.List }

// NewController_getPorts_Results creates a new list of Controller_getPorts_Results.
func NewController_getPorts_Results_List(s *capnp.Segment, sz int32) (Controller_getPorts_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Controller_getPorts_Results_List{}, err
	}
	return Controller_getPorts_Results_List{l}, nil
}

func (s Controller_getPorts_Results_List) At(i int) Controller_getPorts_Results {
	return Controller_getPorts_Results{s.List.Struct(i)}
}
func (s Controller_getPorts_Results_List) Set(i int, v Controller_getPorts_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_getPorts_Results_Promise is a wrapper for a Controller_getPorts_Results promised by a client call.
type Controller_getPorts_Results_Promise struct{ *capnp.Pipeline }

func (p Controller_getPorts_Results_Promise) Struct() (Controller_getPorts_Results, error) {
	s, err := p.Pipeline.Struct()
	return Controller_getPorts_Results{s}, err
}

type Controller_listStreams_Params struct{ capnp.Struct }

func NewController_listStreams_Params(s *capnp.Segment) (Controller_listStreams_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Controller_listStreams_Params{}, err
	}
	return Controller_listStreams_Params{st}, nil
}

func NewRootController_listStreams_Params(s *capnp.Segment) (Controller_listStreams_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Controller_listStreams_Params{}, err
	}
	return Controller_listStreams_Params{st}, nil
}

func ReadRootController_listStreams_Params(msg *capnp.Message) (Controller_listStreams_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_listStreams_Params{}, err
	}
	return Controller_listStreams_Params{root.Struct()}, nil
}

// Controller_listStreams_Params_List is a list of Controller_listStreams_Params.
type Controller_listStreams_Params_List struct{ capnp.List }

// NewController_listStreams_Params creates a new list of Controller_listStreams_Params.
func NewController_listStreams_Params_List(s *capnp.Segment, sz int32) (Controller_listStreams_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Controller_listStreams_Params_List{}, err
	}
	return Controller_listStreams_Params_List{l}, nil
}

func (s Controller_listStreams_Params_List) At(i int) Controller_listStreams_Params {
	return Controller_listStreams_Params{s.List.Struct(i)}
}
func (s Controller_listStreams_Params_List) Set(i int, v Controller_listStreams_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_listStreams_Params_Promise is a wrapper for a Controller_listStreams_Params promised by a client call.
type Controller_listStreams_Params_Promise struct{ *capnp.Pipeline }

func (p Controller_listStreams_Params_Promise) Struct() (Controller_listStreams_Params, error) {
	s, err := p.Pipeline.Struct()
	return Controller_listStreams_Params{s}, err
}

type Controller_listStreams_Results struct{ capnp.Struct }

func NewController_listStreams_Results(s *capnp.Segment) (Controller_listStreams_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Controller_listStreams_Results{}, err
	}
	return Controller_listStreams_Results{st}, nil
}

func NewRootController_listStreams_Results(s *capnp.Segment) (Controller_listStreams_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Controller_listStreams_Results{}, err
	}
	return Controller_listStreams_Results{st}, nil
}

func ReadRootController_listStreams_Results(msg *capnp.Message) (Controller_listStreams_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_listStreams_Results{}, err
	}
	return Controller_listStreams_Results{root.Struct()}, nil
}
func (s Controller_listStreams_Results) Ids() (capnp.UInt16List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return capnp.UInt16List{}, err
	}
	return capnp.UInt16List{List: p.List()}, nil
}

func (s Controller_listStreams_Results) HasIds() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Controller_listStreams_Results) SetIds(v capnp.UInt16List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewIds sets the ids field to a newly
// allocated capnp.UInt16List, preferring placement in s's segment.
func (s Controller_listStreams_Results) NewIds(n int32) (capnp.UInt16List, error) {
	l, err := capnp.NewUInt16List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt16List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Controller_listStreams_Results_List is a list of Controller_listStreams_Results.
type Controller_listStreams_Results_List struct{ capnp.List }

// NewController_listStreams_Results creates a new list of Controller_listStreams_Results.
func NewController_listStreams_Results_List(s *capnp.Segment, sz int32) (Controller_listStreams_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Controller_listStreams_Results_List{}, err
	}
	return Controller_listStreams_Results_List{l}, nil
}

func (s Controller_listStreams_Results_List) At(i int) Controller_listStreams_Results {
	return Controller_listStreams_Results{s.List.Struct(i)}
}
func (s Controller_listStreams_Results_List) Set(i int, v Controller_listStreams_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_listStreams_Results_Promise is a wrapper for a Controller_listStreams_Results promised by a client call.
type Controller_listStreams_Results_Promise struct{ *capnp.Pipeline }

func (p Controller_listStreams_Results_Promise) Struct() (Controller_listStreams_Results, error) {
	s, err := p.Pipeline.Struct()
	return Controller_listStreams_Results{s}, err
}

type Controller_fetchStream_Params struct{ capnp.Struct }

func NewController_fetchStream_Params(s *capnp.Segment) (Controller_fetchStream_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Controller_fetchStream_Params{}, err
	}
	return Controller_fetchStream_Params{st}, nil
}

func NewRootController_fetchStream_Params(s *capnp.Segment) (Controller_fetchStream_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Controller_fetchStream_Params{}, err
	}
	return Controller_fetchStream_Params{st}, nil
}

func ReadRootController_fetchStream_Params(msg *capnp.Message) (Controller_fetchStream_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_fetchStream_Params{}, err
	}
	return Controller_fetchStream_Params{root.Struct()}, nil
}
func (s Controller_fetchStream_Params) Id() uint16 {
	return s.Struct.Uint16(0)
}

func (s Controller_fetchStream_Params) SetId(v uint16) {
	s.Struct.SetUint16(0, v)
}

// Controller_fetchStream_Params_List is a list of Controller_fetchStream_Params.
type Controller_fetchStream_Params_List struct{ capnp.List }

// NewController_fetchStream_Params creates a new list of Controller_fetchStream_Params.
func NewController_fetchStream_Params_List(s *capnp.Segment, sz int32) (Controller_fetchStream_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return Controller_fetchStream_Params_List{}, err
	}
	return Controller_fetchStream_Params_List{l}, nil
}

func (s Controller_fetchStream_Params_List) At(i int) Controller_fetchStream_Params {
	return Controller_fetchStream_Params{s.List.Struct(i)}
}
func (s Controller_fetchStream_Params_List) Set(i int, v Controller_fetchStream_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_fetchStream_Params_Promise is a wrapper for a Controller_fetchStream_Params promised by a client call.
type Controller_fetchStream_Params_Promise struct{ *capnp.Pipeline }

func (p Controller_fetchStream_Params_Promise) Struct() (Controller_fetchStream_Params, error) {
	s, err := p.Pipeline.Struct()
	return Controller_fetchStream_Params{s}, err
}

type Controller_fetchStream_Results struct{ capnp.Struct }

func NewController_fetchStream_Results(s *capnp.Segment) (Controller_fetchStream_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Controller_fetchStream_Results{}, err
	}
	return Controller_fetchStream_Results{st}, nil
}

func NewRootController_fetchStream_Results(s *capnp.Segment) (Controller_fetchStream_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Controller_fetchStream_Results{}, err
	}
	return Controller_fetchStream_Results{st}, nil
}

func ReadRootController_fetchStream_Results(msg *capnp.Message) (Controller_fetchStream_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_fetchStream_Results{}, err
	}
	return Controller_fetchStream_Results{root.Struct()}, nil
}
func (s Controller_fetchStream_Results) Stream() (Stream, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Stream{}, err
	}
	return Stream{Struct: p.Struct()}, nil
}

func (s Controller_fetchStream_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Controller_fetchStream_Results) SetStream(v Stream) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewStream sets the stream field to a newly
// allocated Stream struct, preferring placement in s's segment.
func (s Controller_fetchStream_Results) NewStream() (Stream, error) {
	ss, err := NewStream(s.Struct.Segment())
	if err != nil {
		return Stream{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Controller_fetchStream_Results_List is a list of Controller_fetchStream_Results.
type Controller_fetchStream_Results_List struct{ capnp.List }

// NewController_fetchStream_Results creates a new list of Controller_fetchStream_Results.
func NewController_fetchStream_Results_List(s *capnp.Segment, sz int32) (Controller_fetchStream_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Controller_fetchStream_Results_List{}, err
	}
	return Controller_fetchStream_Results_List{l}, nil
}

func (s Controller_fetchStream_Results_List) At(i int) Controller_fetchStream_Results {
	return Controller_fetchStream_Results{s.List.Struct(i)}
}
func (s Controller_fetchStream_Results_List) Set(i int, v Controller_fetchStream_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_fetchStream_Results_Promise is a wrapper for a Controller_fetchStream_Results promised by a client call.
type Controller_fetchStream_Results_Promise struct{ *capnp.Pipeline }

func (p Controller_fetchStream_Results_Promise) Struct() (Controller_fetchStream_Results, error) {
	s, err := p.Pipeline.Struct()
	return Controller_fetchStream_Results{s}, err
}

func (p Controller_fetchStream_Results_Promise) Stream() Stream_Promise {
	return Stream_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Controller_saveStream_Params struct{ capnp.Struct }

func NewController_saveStream_Params(s *capnp.Segment) (Controller_saveStream_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Controller_saveStream_Params{}, err
	}
	return Controller_saveStream_Params{st}, nil
}

func NewRootController_saveStream_Params(s *capnp.Segment) (Controller_saveStream_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Controller_saveStream_Params{}, err
	}
	return Controller_saveStream_Params{st}, nil
}

func ReadRootController_saveStream_Params(msg *capnp.Message) (Controller_saveStream_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_saveStream_Params{}, err
	}
	return Controller_saveStream_Params{root.Struct()}, nil
}
func (s Controller_saveStream_Params) Stream() (Stream, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Stream{}, err
	}
	return Stream{Struct: p.Struct()}, nil
}

func (s Controller_saveStream_Params) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Controller_saveStream_Params) SetStream(v Stream) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewStream sets the stream field to a newly
// allocated Stream struct, preferring placement in s's segment.
func (s Controller_saveStream_Params) NewStream() (Stream, error) {
	ss, err := NewStream(s.Struct.Segment())
	if err != nil {
		return Stream{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Controller_saveStream_Params_List is a list of Controller_saveStream_Params.
type Controller_saveStream_Params_List struct{ capnp.List }

// NewController_saveStream_Params creates a new list of Controller_saveStream_Params.
func NewController_saveStream_Params_List(s *capnp.Segment, sz int32) (Controller_saveStream_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Controller_saveStream_Params_List{}, err
	}
	return Controller_saveStream_Params_List{l}, nil
}

func (s Controller_saveStream_Params_List) At(i int) Controller_saveStream_Params {
	return Controller_saveStream_Params{s.List.Struct(i)}
}
func (s Controller_saveStream_Params_List) Set(i int, v Controller_saveStream_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_saveStream_Params_Promise is a wrapper for a Controller_saveStream_Params promised by a client call.
type Controller_saveStream_Params_Promise struct{ *capnp.Pipeline }

func (p Controller_saveStream_Params_Promise) Struct() (Controller_saveStream_Params, error) {
	s, err := p.Pipeline.Struct()
	return Controller_saveStream_Params{s}, err
}

func (p Controller_saveStream_Params_Promise) Stream() Stream_Promise {
	return Stream_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Controller_saveStream_Results struct{ capnp.Struct }

func NewController_saveStream_Results(s *capnp.Segment) (Controller_saveStream_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Controller_saveStream_Results{}, err
	}
	return Controller_saveStream_Results{st}, nil
}

func NewRootController_saveStream_Results(s *capnp.Segment) (Controller_saveStream_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Controller_saveStream_Results{}, err
	}
	return Controller_saveStream_Results{st}, nil
}

func ReadRootController_saveStream_Results(msg *capnp.Message) (Controller_saveStream_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_saveStream_Results{}, err
	}
	return Controller_saveStream_Results{root.Struct()}, nil
}
func (s Controller_saveStream_Results) Id() uint16 {
	return s.Struct.Uint16(0)
}

func (s Controller_saveStream_Results) SetId(v uint16) {
	s.Struct.SetUint16(0, v)
}

// Controller_saveStream_Results_List is a list of Controller_saveStream_Results.
type Controller_saveStream_Results_List struct{ capnp.List }

// NewController_saveStream_Results creates a new list of Controller_saveStream_Results.
func NewController_saveStream_Results_List(s *capnp.Segment, sz int32) (Controller_saveStream_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return Controller_saveStream_Results_List{}, err
	}
	return Controller_saveStream_Results_List{l}, nil
}

func (s Controller_saveStream_Results_List) At(i int) Controller_saveStream_Results {
	return Controller_saveStream_Results{s.List.Struct(i)}
}
func (s Controller_saveStream_Results_List) Set(i int, v Controller_saveStream_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_saveStream_Results_Promise is a wrapper for a Controller_saveStream_Results promised by a client call.
type Controller_saveStream_Results_Promise struct{ *capnp.Pipeline }

func (p Controller_saveStream_Results_Promise) Struct() (Controller_saveStream_Results, error) {
	s, err := p.Pipeline.Struct()
	return Controller_saveStream_Results{s}, err
}

type Controller_deleteStream_Params struct{ capnp.Struct }

func NewController_deleteStream_Params(s *capnp.Segment) (Controller_deleteStream_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Controller_deleteStream_Params{}, err
	}
	return Controller_deleteStream_Params{st}, nil
}

func NewRootController_deleteStream_Params(s *capnp.Segment) (Controller_deleteStream_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	if err != nil {
		return Controller_deleteStream_Params{}, err
	}
	return Controller_deleteStream_Params{st}, nil
}

func ReadRootController_deleteStream_Params(msg *capnp.Message) (Controller_deleteStream_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_deleteStream_Params{}, err
	}
	return Controller_deleteStream_Params{root.Struct()}, nil
}
func (s Controller_deleteStream_Params) Id() uint16 {
	return s.Struct.Uint16(0)
}

func (s Controller_deleteStream_Params) SetId(v uint16) {
	s.Struct.SetUint16(0, v)
}

// Controller_deleteStream_Params_List is a list of Controller_deleteStream_Params.
type Controller_deleteStream_Params_List struct{ capnp.List }

// NewController_deleteStream_Params creates a new list of Controller_deleteStream_Params.
func NewController_deleteStream_Params_List(s *capnp.Segment, sz int32) (Controller_deleteStream_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	if err != nil {
		return Controller_deleteStream_Params_List{}, err
	}
	return Controller_deleteStream_Params_List{l}, nil
}

func (s Controller_deleteStream_Params_List) At(i int) Controller_deleteStream_Params {
	return Controller_deleteStream_Params{s.List.Struct(i)}
}
func (s Controller_deleteStream_Params_List) Set(i int, v Controller_deleteStream_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_deleteStream_Params_Promise is a wrapper for a Controller_deleteStream_Params promised by a client call.
type Controller_deleteStream_Params_Promise struct{ *capnp.Pipeline }

func (p Controller_deleteStream_Params_Promise) Struct() (Controller_deleteStream_Params, error) {
	s, err := p.Pipeline.Struct()
	return Controller_deleteStream_Params{s}, err
}

type Controller_deleteStream_Results struct{ capnp.Struct }

func NewController_deleteStream_Results(s *capnp.Segment) (Controller_deleteStream_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Controller_deleteStream_Results{}, err
	}
	return Controller_deleteStream_Results{st}, nil
}

func NewRootController_deleteStream_Results(s *capnp.Segment) (Controller_deleteStream_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Controller_deleteStream_Results{}, err
	}
	return Controller_deleteStream_Results{st}, nil
}

func ReadRootController_deleteStream_Results(msg *capnp.Message) (Controller_deleteStream_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Controller_deleteStream_Results{}, err
	}
	return Controller_deleteStream_Results{root.Struct()}, nil
}

// Controller_deleteStream_Results_List is a list of Controller_deleteStream_Results.
type Controller_deleteStream_Results_List struct{ capnp.List }

// NewController_deleteStream_Results creates a new list of Controller_deleteStream_Results.
func NewController_deleteStream_Results_List(s *capnp.Segment, sz int32) (Controller_deleteStream_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Controller_deleteStream_Results_List{}, err
	}
	return Controller_deleteStream_Results_List{l}, nil
}

func (s Controller_deleteStream_Results_List) At(i int) Controller_deleteStream_Results {
	return Controller_deleteStream_Results{s.List.Struct(i)}
}
func (s Controller_deleteStream_Results_List) Set(i int, v Controller_deleteStream_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Controller_deleteStream_Results_Promise is a wrapper for a Controller_deleteStream_Results promised by a client call.
type Controller_deleteStream_Results_Promise struct{ *capnp.Pipeline }

func (p Controller_deleteStream_Results_Promise) Struct() (Controller_deleteStream_Results, error) {
	s, err := p.Pipeline.Struct()
	return Controller_deleteStream_Results{s}, err
}

type Port struct{ Client capnp.Client }

func (c Port) GetConfig(ctx context.Context, params func(Port_getConfig_Params) error, opts ...capnp.CallOption) Port_getConfig_Results_Promise {
	if c.Client == nil {
		return Port_getConfig_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      0,
			InterfaceName: "main.capnp:Port",
			MethodName:    "getConfig",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Port_getConfig_Params{Struct: s}) }
	}
	return Port_getConfig_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Port) SetConfig(ctx context.Context, params func(Port_setConfig_Params) error, opts ...capnp.CallOption) Port_setConfig_Results_Promise {
	if c.Client == nil {
		return Port_setConfig_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      1,
			InterfaceName: "main.capnp:Port",
			MethodName:    "setConfig",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Port_setConfig_Params{Struct: s}) }
	}
	return Port_setConfig_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Port) StartSend(ctx context.Context, params func(Port_startSend_Params) error, opts ...capnp.CallOption) Port_startSend_Results_Promise {
	if c.Client == nil {
		return Port_startSend_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      2,
			InterfaceName: "main.capnp:Port",
			MethodName:    "startSend",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Port_startSend_Params{Struct: s}) }
	}
	return Port_startSend_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Port_Server interface {
	GetConfig(Port_getConfig) error

	SetConfig(Port_setConfig) error

	StartSend(Port_startSend) error
}

func Port_ServerToClient(s Port_Server) Port {
	c, _ := s.(server.Closer)
	return Port{Client: server.New(Port_Methods(nil, s), c)}
}

func Port_Methods(methods []server.Method, s Port_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 3)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      0,
			InterfaceName: "main.capnp:Port",
			MethodName:    "getConfig",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Port_getConfig{c, opts, Port_getConfig_Params{Struct: p}, Port_getConfig_Results{Struct: r}}
			return s.GetConfig(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      1,
			InterfaceName: "main.capnp:Port",
			MethodName:    "setConfig",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Port_setConfig{c, opts, Port_setConfig_Params{Struct: p}, Port_setConfig_Results{Struct: r}}
			return s.SetConfig(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      2,
			InterfaceName: "main.capnp:Port",
			MethodName:    "startSend",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Port_startSend{c, opts, Port_startSend_Params{Struct: p}, Port_startSend_Results{Struct: r}}
			return s.StartSend(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// Port_getConfig holds the arguments for a server call to Port.getConfig.
type Port_getConfig struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Port_getConfig_Params
	Results Port_getConfig_Results
}

// Port_setConfig holds the arguments for a server call to Port.setConfig.
type Port_setConfig struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Port_setConfig_Params
	Results Port_setConfig_Results
}

// Port_startSend holds the arguments for a server call to Port.startSend.
type Port_startSend struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Port_startSend_Params
	Results Port_startSend_Results
}

type Port_Config struct{ capnp.Struct }

func NewPort_Config(s *capnp.Segment) (Port_Config, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_Config{}, err
	}
	return Port_Config{st}, nil
}

func NewRootPort_Config(s *capnp.Segment) (Port_Config, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_Config{}, err
	}
	return Port_Config{st}, nil
}

func ReadRootPort_Config(msg *capnp.Message) (Port_Config, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_Config{}, err
	}
	return Port_Config{root.Struct()}, nil
}
func (s Port_Config) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Port_Config) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Port_Config) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Port_Config) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Port_Config_List is a list of Port_Config.
type Port_Config_List struct{ capnp.List }

// NewPort_Config creates a new list of Port_Config.
func NewPort_Config_List(s *capnp.Segment, sz int32) (Port_Config_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Port_Config_List{}, err
	}
	return Port_Config_List{l}, nil
}

func (s Port_Config_List) At(i int) Port_Config           { return Port_Config{s.List.Struct(i)} }
func (s Port_Config_List) Set(i int, v Port_Config) error { return s.List.SetStruct(i, v.Struct) }

// Port_Config_Promise is a wrapper for a Port_Config promised by a client call.
type Port_Config_Promise struct{ *capnp.Pipeline }

func (p Port_Config_Promise) Struct() (Port_Config, error) {
	s, err := p.Pipeline.Struct()
	return Port_Config{s}, err
}

type Port_getConfig_Params struct{ capnp.Struct }

func NewPort_getConfig_Params(s *capnp.Segment) (Port_getConfig_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_getConfig_Params{}, err
	}
	return Port_getConfig_Params{st}, nil
}

func NewRootPort_getConfig_Params(s *capnp.Segment) (Port_getConfig_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_getConfig_Params{}, err
	}
	return Port_getConfig_Params{st}, nil
}

func ReadRootPort_getConfig_Params(msg *capnp.Message) (Port_getConfig_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_getConfig_Params{}, err
	}
	return Port_getConfig_Params{root.Struct()}, nil
}

// Port_getConfig_Params_List is a list of Port_getConfig_Params.
type Port_getConfig_Params_List struct{ capnp.List }

// NewPort_getConfig_Params creates a new list of Port_getConfig_Params.
func NewPort_getConfig_Params_List(s *capnp.Segment, sz int32) (Port_getConfig_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Port_getConfig_Params_List{}, err
	}
	return Port_getConfig_Params_List{l}, nil
}

func (s Port_getConfig_Params_List) At(i int) Port_getConfig_Params {
	return Port_getConfig_Params{s.List.Struct(i)}
}
func (s Port_getConfig_Params_List) Set(i int, v Port_getConfig_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_getConfig_Params_Promise is a wrapper for a Port_getConfig_Params promised by a client call.
type Port_getConfig_Params_Promise struct{ *capnp.Pipeline }

func (p Port_getConfig_Params_Promise) Struct() (Port_getConfig_Params, error) {
	s, err := p.Pipeline.Struct()
	return Port_getConfig_Params{s}, err
}

type Port_getConfig_Results struct{ capnp.Struct }

func NewPort_getConfig_Results(s *capnp.Segment) (Port_getConfig_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_getConfig_Results{}, err
	}
	return Port_getConfig_Results{st}, nil
}

func NewRootPort_getConfig_Results(s *capnp.Segment) (Port_getConfig_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_getConfig_Results{}, err
	}
	return Port_getConfig_Results{st}, nil
}

func ReadRootPort_getConfig_Results(msg *capnp.Message) (Port_getConfig_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_getConfig_Results{}, err
	}
	return Port_getConfig_Results{root.Struct()}, nil
}
func (s Port_getConfig_Results) Config() (Port_Config, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Port_Config{}, err
	}
	return Port_Config{Struct: p.Struct()}, nil
}

func (s Port_getConfig_Results) HasConfig() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Port_getConfig_Results) SetConfig(v Port_Config) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewConfig sets the config field to a newly
// allocated Port_Config struct, preferring placement in s's segment.
func (s Port_getConfig_Results) NewConfig() (Port_Config, error) {
	ss, err := NewPort_Config(s.Struct.Segment())
	if err != nil {
		return Port_Config{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Port_getConfig_Results_List is a list of Port_getConfig_Results.
type Port_getConfig_Results_List struct{ capnp.List }

// NewPort_getConfig_Results creates a new list of Port_getConfig_Results.
func NewPort_getConfig_Results_List(s *capnp.Segment, sz int32) (Port_getConfig_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Port_getConfig_Results_List{}, err
	}
	return Port_getConfig_Results_List{l}, nil
}

func (s Port_getConfig_Results_List) At(i int) Port_getConfig_Results {
	return Port_getConfig_Results{s.List.Struct(i)}
}
func (s Port_getConfig_Results_List) Set(i int, v Port_getConfig_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_getConfig_Results_Promise is a wrapper for a Port_getConfig_Results promised by a client call.
type Port_getConfig_Results_Promise struct{ *capnp.Pipeline }

func (p Port_getConfig_Results_Promise) Struct() (Port_getConfig_Results, error) {
	s, err := p.Pipeline.Struct()
	return Port_getConfig_Results{s}, err
}

func (p Port_getConfig_Results_Promise) Config() Port_Config_Promise {
	return Port_Config_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Port_setConfig_Params struct{ capnp.Struct }

func NewPort_setConfig_Params(s *capnp.Segment) (Port_setConfig_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_setConfig_Params{}, err
	}
	return Port_setConfig_Params{st}, nil
}

func NewRootPort_setConfig_Params(s *capnp.Segment) (Port_setConfig_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_setConfig_Params{}, err
	}
	return Port_setConfig_Params{st}, nil
}

func ReadRootPort_setConfig_Params(msg *capnp.Message) (Port_setConfig_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_setConfig_Params{}, err
	}
	return Port_setConfig_Params{root.Struct()}, nil
}
func (s Port_setConfig_Params) Config() (Port_Config, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Port_Config{}, err
	}
	return Port_Config{Struct: p.Struct()}, nil
}

func (s Port_setConfig_Params) HasConfig() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Port_setConfig_Params) SetConfig(v Port_Config) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewConfig sets the config field to a newly
// allocated Port_Config struct, preferring placement in s's segment.
func (s Port_setConfig_Params) NewConfig() (Port_Config, error) {
	ss, err := NewPort_Config(s.Struct.Segment())
	if err != nil {
		return Port_Config{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Port_setConfig_Params_List is a list of Port_setConfig_Params.
type Port_setConfig_Params_List struct{ capnp.List }

// NewPort_setConfig_Params creates a new list of Port_setConfig_Params.
func NewPort_setConfig_Params_List(s *capnp.Segment, sz int32) (Port_setConfig_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Port_setConfig_Params_List{}, err
	}
	return Port_setConfig_Params_List{l}, nil
}

func (s Port_setConfig_Params_List) At(i int) Port_setConfig_Params {
	return Port_setConfig_Params{s.List.Struct(i)}
}
func (s Port_setConfig_Params_List) Set(i int, v Port_setConfig_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_setConfig_Params_Promise is a wrapper for a Port_setConfig_Params promised by a client call.
type Port_setConfig_Params_Promise struct{ *capnp.Pipeline }

func (p Port_setConfig_Params_Promise) Struct() (Port_setConfig_Params, error) {
	s, err := p.Pipeline.Struct()
	return Port_setConfig_Params{s}, err
}

func (p Port_setConfig_Params_Promise) Config() Port_Config_Promise {
	return Port_Config_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Port_setConfig_Results struct{ capnp.Struct }

func NewPort_setConfig_Results(s *capnp.Segment) (Port_setConfig_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_setConfig_Results{}, err
	}
	return Port_setConfig_Results{st}, nil
}

func NewRootPort_setConfig_Results(s *capnp.Segment) (Port_setConfig_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_setConfig_Results{}, err
	}
	return Port_setConfig_Results{st}, nil
}

func ReadRootPort_setConfig_Results(msg *capnp.Message) (Port_setConfig_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_setConfig_Results{}, err
	}
	return Port_setConfig_Results{root.Struct()}, nil
}

// Port_setConfig_Results_List is a list of Port_setConfig_Results.
type Port_setConfig_Results_List struct{ capnp.List }

// NewPort_setConfig_Results creates a new list of Port_setConfig_Results.
func NewPort_setConfig_Results_List(s *capnp.Segment, sz int32) (Port_setConfig_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Port_setConfig_Results_List{}, err
	}
	return Port_setConfig_Results_List{l}, nil
}

func (s Port_setConfig_Results_List) At(i int) Port_setConfig_Results {
	return Port_setConfig_Results{s.List.Struct(i)}
}
func (s Port_setConfig_Results_List) Set(i int, v Port_setConfig_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_setConfig_Results_Promise is a wrapper for a Port_setConfig_Results promised by a client call.
type Port_setConfig_Results_Promise struct{ *capnp.Pipeline }

func (p Port_setConfig_Results_Promise) Struct() (Port_setConfig_Results, error) {
	s, err := p.Pipeline.Struct()
	return Port_setConfig_Results{s}, err
}

type Port_startSend_Params struct{ capnp.Struct }

func NewPort_startSend_Params(s *capnp.Segment) (Port_startSend_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_startSend_Params{}, err
	}
	return Port_startSend_Params{st}, nil
}

func NewRootPort_startSend_Params(s *capnp.Segment) (Port_startSend_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_startSend_Params{}, err
	}
	return Port_startSend_Params{st}, nil
}

func ReadRootPort_startSend_Params(msg *capnp.Message) (Port_startSend_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_startSend_Params{}, err
	}
	return Port_startSend_Params{root.Struct()}, nil
}
func (s Port_startSend_Params) Ids() (capnp.UInt16List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return capnp.UInt16List{}, err
	}
	return capnp.UInt16List{List: p.List()}, nil
}

func (s Port_startSend_Params) HasIds() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Port_startSend_Params) SetIds(v capnp.UInt16List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewIds sets the ids field to a newly
// allocated capnp.UInt16List, preferring placement in s's segment.
func (s Port_startSend_Params) NewIds(n int32) (capnp.UInt16List, error) {
	l, err := capnp.NewUInt16List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.UInt16List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Port_startSend_Params_List is a list of Port_startSend_Params.
type Port_startSend_Params_List struct{ capnp.List }

// NewPort_startSend_Params creates a new list of Port_startSend_Params.
func NewPort_startSend_Params_List(s *capnp.Segment, sz int32) (Port_startSend_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Port_startSend_Params_List{}, err
	}
	return Port_startSend_Params_List{l}, nil
}

func (s Port_startSend_Params_List) At(i int) Port_startSend_Params {
	return Port_startSend_Params{s.List.Struct(i)}
}
func (s Port_startSend_Params_List) Set(i int, v Port_startSend_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_startSend_Params_Promise is a wrapper for a Port_startSend_Params promised by a client call.
type Port_startSend_Params_Promise struct{ *capnp.Pipeline }

func (p Port_startSend_Params_Promise) Struct() (Port_startSend_Params, error) {
	s, err := p.Pipeline.Struct()
	return Port_startSend_Params{s}, err
}

type Port_startSend_Results struct{ capnp.Struct }

func NewPort_startSend_Results(s *capnp.Segment) (Port_startSend_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_startSend_Results{}, err
	}
	return Port_startSend_Results{st}, nil
}

func NewRootPort_startSend_Results(s *capnp.Segment) (Port_startSend_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_startSend_Results{}, err
	}
	return Port_startSend_Results{st}, nil
}

func ReadRootPort_startSend_Results(msg *capnp.Message) (Port_startSend_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_startSend_Results{}, err
	}
	return Port_startSend_Results{root.Struct()}, nil
}

// Port_startSend_Results_List is a list of Port_startSend_Results.
type Port_startSend_Results_List struct{ capnp.List }

// NewPort_startSend_Results creates a new list of Port_startSend_Results.
func NewPort_startSend_Results_List(s *capnp.Segment, sz int32) (Port_startSend_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Port_startSend_Results_List{}, err
	}
	return Port_startSend_Results_List{l}, nil
}

func (s Port_startSend_Results_List) At(i int) Port_startSend_Results {
	return Port_startSend_Results{s.List.Struct(i)}
}
func (s Port_startSend_Results_List) Set(i int, v Port_startSend_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_startSend_Results_Promise is a wrapper for a Port_startSend_Results promised by a client call.
type Port_startSend_Results_Promise struct{ *capnp.Pipeline }

func (p Port_startSend_Results_Promise) Struct() (Port_startSend_Results, error) {
	s, err := p.Pipeline.Struct()
	return Port_startSend_Results{s}, err
}

type Stream struct{ capnp.Struct }

func NewStream(s *capnp.Segment) (Stream, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Stream{}, err
	}
	return Stream{st}, nil
}

func NewRootStream(s *capnp.Segment) (Stream, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Stream{}, err
	}
	return Stream{st}, nil
}

func ReadRootStream(msg *capnp.Message) (Stream, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream{}, err
	}
	return Stream{root.Struct()}, nil
}
func (s Stream) Id() uint16 {
	return s.Struct.Uint16(0)
}

func (s Stream) SetId(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s Stream) Count() uint32 {
	return s.Struct.Uint32(4) ^ 1
}

func (s Stream) SetCount(v uint32) {
	s.Struct.SetUint32(4, v^1)
}

func (s Stream) PacketsPerSec() uint32 {
	return s.Struct.Uint32(8) ^ 1
}

func (s Stream) SetPacketsPerSec(v uint32) {
	s.Struct.SetUint32(8, v^1)
}

func (s Stream) Layers() (Protocol_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Protocol_List{}, err
	}
	return Protocol_List{List: p.List()}, nil
}

func (s Stream) HasLayers() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Stream) SetLayers(v Protocol_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewLayers sets the layers field to a newly
// allocated Protocol_List, preferring placement in s's segment.
func (s Stream) NewLayers(n int32) (Protocol_List, error) {
	l, err := NewProtocol_List(s.Struct.Segment(), n)
	if err != nil {
		return Protocol_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Stream_List is a list of Stream.
type Stream_List struct{ capnp.List }

// NewStream creates a new list of Stream.
func NewStream_List(s *capnp.Segment, sz int32) (Stream_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	if err != nil {
		return Stream_List{}, err
	}
	return Stream_List{l}, nil
}

func (s Stream_List) At(i int) Stream           { return Stream{s.List.Struct(i)} }
func (s Stream_List) Set(i int, v Stream) error { return s.List.SetStruct(i, v.Struct) }

// Stream_Promise is a wrapper for a Stream promised by a client call.
type Stream_Promise struct{ *capnp.Pipeline }

func (p Stream_Promise) Struct() (Stream, error) {
	s, err := p.Pipeline.Struct()
	return Stream{s}, err
}

type Field struct{ capnp.Struct }

func NewField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return Field{}, err
	}
	return Field{st}, nil
}

func NewRootField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return Field{}, err
	}
	return Field{st}, nil
}

func ReadRootField(msg *capnp.Message) (Field, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Field{}, err
	}
	return Field{root.Struct()}, nil
}
func (s Field) Value() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Field) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Field) SetValue(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s Field) Mode() uint8 {
	return s.Struct.Uint8(0)
}

func (s Field) SetMode(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Field) Step() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Field) HasStep() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Field) SetStep(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s Field) Mask() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Field) HasMask() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Field) SetMask(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, d.List.ToPtr())
}

func (s Field) Count() uint16 {
	return s.Struct.Uint16(2) ^ 1
}

func (s Field) SetCount(v uint16) {
	s.Struct.SetUint16(2, v^1)
}

// Field_List is a list of Field.
type Field_List struct{ capnp.List }

// NewField creates a new list of Field.
func NewField_List(s *capnp.Segment, sz int32) (Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	if err != nil {
		return Field_List{}, err
	}
	return Field_List{l}, nil
}

func (s Field_List) At(i int) Field           { return Field{s.List.Struct(i)} }
func (s Field_List) Set(i int, v Field) error { return s.List.SetStruct(i, v.Struct) }

// Field_Promise is a wrapper for a Field promised by a client call.
type Field_Promise struct{ *capnp.Pipeline }

func (p Field_Promise) Struct() (Field, error) {
	s, err := p.Pipeline.Struct()
	return Field{s}, err
}

type Protocol struct{ capnp.Struct }
type Protocol_ethernet2 Protocol
type Protocol_ipv4 Protocol
type Protocol_Which uint16

const (
	Protocol_Which_ethernet2 Protocol_Which = 0
	Protocol_Which_ipv4      Protocol_Which = 1
)

func (w Protocol_Which) String() string {
	const s = "ethernet2ipv4"
	switch w {
	case Protocol_Which_ethernet2:
		return s[0:9]
	case Protocol_Which_ipv4:
		return s[9:13]

	}
	return "Protocol_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

func NewProtocol(s *capnp.Segment) (Protocol, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 14})
	if err != nil {
		return Protocol{}, err
	}
	return Protocol{st}, nil
}

func NewRootProtocol(s *capnp.Segment) (Protocol, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 14})
	if err != nil {
		return Protocol{}, err
	}
	return Protocol{st}, nil
}

func ReadRootProtocol(msg *capnp.Message) (Protocol, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Protocol{}, err
	}
	return Protocol{root.Struct()}, nil
}

func (s Protocol) Which() Protocol_Which {
	return Protocol_Which(s.Struct.Uint16(0))
}
func (s Protocol) Ethernet2() Protocol_ethernet2 { return Protocol_ethernet2(s) }
func (s Protocol) SetEthernet2() {
	s.Struct.SetUint16(0, 0)
}
func (s Protocol_ethernet2) Source() (Field, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[0:64])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ethernet2) HasSource() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet2) SetSource(v Field) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSource sets the source field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ethernet2) NewSource() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ethernet2) Destination() (Field, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[64:128])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ethernet2) HasDestination() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet2) SetDestination(v Field) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDestination sets the destination field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ethernet2) NewDestination() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ethernet2) EthernetType() (Field, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[128:192])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ethernet2) HasEthernetType() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet2) SetEthernetType(v Field) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewEthernetType sets the ethernetType field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ethernet2) NewEthernetType() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol) Ipv4() Protocol_ipv4 { return Protocol_ipv4(s) }
func (s Protocol) SetIpv4() {
	s.Struct.SetUint16(0, 1)
}
func (s Protocol_ipv4) Version() (Field, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[192:256])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasVersion() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetVersion(v Field) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewVersion sets the version field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewVersion() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Ihl() (Field, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[256:320])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasIhl() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetIhl(v Field) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewIhl sets the ihl field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewIhl() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Tos() (Field, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[320:384])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasTos() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetTos(v Field) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewTos sets the tos field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewTos() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Length() (Field, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[384:448])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasLength() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetLength(v Field) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewLength sets the length field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewLength() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Id() (Field, error) {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[448:512])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasId() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetId(v Field) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewId sets the id field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewId() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Flags() (Field, error) {
	p, err := s.Struct.Ptr(5)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[512:576])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasFlags() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetFlags(v Field) error {
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewFlags sets the flags field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewFlags() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) FragOffset() (Field, error) {
	p, err := s.Struct.Ptr(6)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[576:640])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasFragOffset() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetFragOffset(v Field) error {
	return s.Struct.SetPtr(6, v.Struct.ToPtr())
}

// NewFragOffset sets the fragOffset field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewFragOffset() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(6, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Ttl() (Field, error) {
	p, err := s.Struct.Ptr(7)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[640:704])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasTtl() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetTtl(v Field) error {
	return s.Struct.SetPtr(7, v.Struct.ToPtr())
}

// NewTtl sets the ttl field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewTtl() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(7, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Protocol() (Field, error) {
	p, err := s.Struct.Ptr(8)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[704:768])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasProtocol() bool {
	p, err := s.Struct.Ptr(8)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetProtocol(v Field) error {
	return s.Struct.SetPtr(8, v.Struct.ToPtr())
}

// NewProtocol sets the protocol field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewProtocol() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(8, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Checksum() (Field, error) {
	p, err := s.Struct.Ptr(9)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[768:832])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasChecksum() bool {
	p, err := s.Struct.Ptr(9)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetChecksum(v Field) error {
	return s.Struct.SetPtr(9, v.Struct.ToPtr())
}

// NewChecksum sets the checksum field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewChecksum() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(9, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Source() (Field, error) {
	p, err := s.Struct.Ptr(10)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[832:896])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasSource() bool {
	p, err := s.Struct.Ptr(10)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetSource(v Field) error {
	return s.Struct.SetPtr(10, v.Struct.ToPtr())
}

// NewSource sets the source field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewSource() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(10, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Destination() (Field, error) {
	p, err := s.Struct.Ptr(11)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[896:960])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasDestination() bool {
	p, err := s.Struct.Ptr(11)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetDestination(v Field) error {
	return s.Struct.SetPtr(11, v.Struct.ToPtr())
}

// NewDestination sets the destination field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewDestination() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(11, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Options() (Field, error) {
	p, err := s.Struct.Ptr(12)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[960:1024])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasOptions() bool {
	p, err := s.Struct.Ptr(12)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetOptions(v Field) error {
	return s.Struct.SetPtr(12, v.Struct.ToPtr())
}

// NewOptions sets the options field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewOptions() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(12, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Padding() (Field, error) {
	p, err := s.Struct.Ptr(13)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[1024:1088])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasPadding() bool {
	p, err := s.Struct.Ptr(13)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetPadding(v Field) error {
	return s.Struct.SetPtr(13, v.Struct.ToPtr())
}

// NewPadding sets the padding field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewPadding() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(13, ss.Struct.ToPtr())
	return ss, err
}

// Protocol_List is a list of Protocol.
type Protocol_List struct{ capnp.List }

// NewProtocol creates a new list of Protocol.
func NewProtocol_List(s *capnp.Segment, sz int32) (Protocol_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 14}, sz)
	if err != nil {
		return Protocol_List{}, err
	}
	return Protocol_List{l}, nil
}

func (s Protocol_List) At(i int) Protocol           { return Protocol{s.List.Struct(i)} }
func (s Protocol_List) Set(i int, v Protocol) error { return s.List.SetStruct(i, v.Struct) }

// Protocol_Promise is a wrapper for a Protocol promised by a client call.
type Protocol_Promise struct{ *capnp.Pipeline }

func (p Protocol_Promise) Struct() (Protocol, error) {
	s, err := p.Pipeline.Struct()
	return Protocol{s}, err
}

func (p Protocol_Promise) Ethernet2() Protocol_ethernet2_Promise {
	return Protocol_ethernet2_Promise{p.Pipeline}
}

// Protocol_ethernet2_Promise is a wrapper for a Protocol_ethernet2 promised by a client call.
type Protocol_ethernet2_Promise struct{ *capnp.Pipeline }

func (p Protocol_ethernet2_Promise) Struct() (Protocol_ethernet2, error) {
	s, err := p.Pipeline.Struct()
	return Protocol_ethernet2{s}, err
}

func (p Protocol_ethernet2_Promise) Source() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(0, x_ef97cf4069588836[1088:1152])}
}

func (p Protocol_ethernet2_Promise) Destination() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(1, x_ef97cf4069588836[1152:1216])}
}

func (p Protocol_ethernet2_Promise) EthernetType() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(2, x_ef97cf4069588836[1216:1280])}
}

func (p Protocol_Promise) Ipv4() Protocol_ipv4_Promise { return Protocol_ipv4_Promise{p.Pipeline} }

// Protocol_ipv4_Promise is a wrapper for a Protocol_ipv4 promised by a client call.
type Protocol_ipv4_Promise struct{ *capnp.Pipeline }

func (p Protocol_ipv4_Promise) Struct() (Protocol_ipv4, error) {
	s, err := p.Pipeline.Struct()
	return Protocol_ipv4{s}, err
}

func (p Protocol_ipv4_Promise) Version() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(0, x_ef97cf4069588836[1280:1344])}
}

func (p Protocol_ipv4_Promise) Ihl() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(1, x_ef97cf4069588836[1344:1408])}
}

func (p Protocol_ipv4_Promise) Tos() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(2, x_ef97cf4069588836[1408:1472])}
}

func (p Protocol_ipv4_Promise) Length() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(3, x_ef97cf4069588836[1472:1536])}
}

func (p Protocol_ipv4_Promise) Id() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(4, x_ef97cf4069588836[1536:1600])}
}

func (p Protocol_ipv4_Promise) Flags() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(5, x_ef97cf4069588836[1600:1664])}
}

func (p Protocol_ipv4_Promise) FragOffset() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(6, x_ef97cf4069588836[1664:1728])}
}

func (p Protocol_ipv4_Promise) Ttl() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(7, x_ef97cf4069588836[1728:1792])}
}

func (p Protocol_ipv4_Promise) Protocol() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(8, x_ef97cf4069588836[1792:1856])}
}

func (p Protocol_ipv4_Promise) Checksum() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(9, x_ef97cf4069588836[1856:1920])}
}

func (p Protocol_ipv4_Promise) Source() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(10, x_ef97cf4069588836[1920:1984])}
}

func (p Protocol_ipv4_Promise) Destination() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(11, x_ef97cf4069588836[1984:2048])}
}

func (p Protocol_ipv4_Promise) Options() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(12, x_ef97cf4069588836[2048:2112])}
}

func (p Protocol_ipv4_Promise) Padding() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(13, x_ef97cf4069588836[2112:2176])}
}

var x_ef97cf4069588836 = []byte{
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 50, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	8, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	5, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 34, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 50, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	8, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	5, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 34, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 7, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 3, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
}
