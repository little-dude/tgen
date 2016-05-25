package capnp

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

type Controller_Server interface {
	GetPorts(Controller_getPorts) error
}

func Controller_ServerToClient(s Controller_Server) Controller {
	c, _ := s.(server.Closer)
	return Controller{Client: server.New(Controller_Methods(nil, s), c)}
}

func Controller_Methods(methods []server.Method, s Controller_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
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

	return methods
}

// Controller_getPorts holds the arguments for a server call to Controller.getPorts.
type Controller_getPorts struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Controller_getPorts_Params
	Results Controller_getPorts_Results
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
func (c Port) GetStreams(ctx context.Context, params func(Port_getStreams_Params) error, opts ...capnp.CallOption) Port_getStreams_Results_Promise {
	if c.Client == nil {
		return Port_getStreams_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      2,
			InterfaceName: "main.capnp:Port",
			MethodName:    "getStreams",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Port_getStreams_Params{Struct: s}) }
	}
	return Port_getStreams_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Port) NewStream(ctx context.Context, params func(Port_newStream_Params) error, opts ...capnp.CallOption) Port_newStream_Results_Promise {
	if c.Client == nil {
		return Port_newStream_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      3,
			InterfaceName: "main.capnp:Port",
			MethodName:    "newStream",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Port_newStream_Params{Struct: s}) }
	}
	return Port_newStream_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Port) DelStream(ctx context.Context, params func(Port_delStream_Params) error, opts ...capnp.CallOption) Port_delStream_Results_Promise {
	if c.Client == nil {
		return Port_delStream_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      4,
			InterfaceName: "main.capnp:Port",
			MethodName:    "delStream",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Port_delStream_Params{Struct: s}) }
	}
	return Port_delStream_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Port_Server interface {
	GetConfig(Port_getConfig) error

	SetConfig(Port_setConfig) error

	GetStreams(Port_getStreams) error

	NewStream(Port_newStream) error

	DelStream(Port_delStream) error
}

func Port_ServerToClient(s Port_Server) Port {
	c, _ := s.(server.Closer)
	return Port{Client: server.New(Port_Methods(nil, s), c)}
}

func Port_Methods(methods []server.Method, s Port_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 5)
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
			MethodName:    "getStreams",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Port_getStreams{c, opts, Port_getStreams_Params{Struct: p}, Port_getStreams_Results{Struct: r}}
			return s.GetStreams(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      3,
			InterfaceName: "main.capnp:Port",
			MethodName:    "newStream",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Port_newStream{c, opts, Port_newStream_Params{Struct: p}, Port_newStream_Results{Struct: r}}
			return s.NewStream(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      4,
			InterfaceName: "main.capnp:Port",
			MethodName:    "delStream",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Port_delStream{c, opts, Port_delStream_Params{Struct: p}, Port_delStream_Results{Struct: r}}
			return s.DelStream(call)
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

// Port_getStreams holds the arguments for a server call to Port.getStreams.
type Port_getStreams struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Port_getStreams_Params
	Results Port_getStreams_Results
}

// Port_newStream holds the arguments for a server call to Port.newStream.
type Port_newStream struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Port_newStream_Params
	Results Port_newStream_Results
}

// Port_delStream holds the arguments for a server call to Port.delStream.
type Port_delStream struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Port_delStream_Params
	Results Port_delStream_Results
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

type Port_getStreams_Params struct{ capnp.Struct }

func NewPort_getStreams_Params(s *capnp.Segment) (Port_getStreams_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_getStreams_Params{}, err
	}
	return Port_getStreams_Params{st}, nil
}

func NewRootPort_getStreams_Params(s *capnp.Segment) (Port_getStreams_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_getStreams_Params{}, err
	}
	return Port_getStreams_Params{st}, nil
}

func ReadRootPort_getStreams_Params(msg *capnp.Message) (Port_getStreams_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_getStreams_Params{}, err
	}
	return Port_getStreams_Params{root.Struct()}, nil
}

// Port_getStreams_Params_List is a list of Port_getStreams_Params.
type Port_getStreams_Params_List struct{ capnp.List }

// NewPort_getStreams_Params creates a new list of Port_getStreams_Params.
func NewPort_getStreams_Params_List(s *capnp.Segment, sz int32) (Port_getStreams_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Port_getStreams_Params_List{}, err
	}
	return Port_getStreams_Params_List{l}, nil
}

func (s Port_getStreams_Params_List) At(i int) Port_getStreams_Params {
	return Port_getStreams_Params{s.List.Struct(i)}
}
func (s Port_getStreams_Params_List) Set(i int, v Port_getStreams_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_getStreams_Params_Promise is a wrapper for a Port_getStreams_Params promised by a client call.
type Port_getStreams_Params_Promise struct{ *capnp.Pipeline }

func (p Port_getStreams_Params_Promise) Struct() (Port_getStreams_Params, error) {
	s, err := p.Pipeline.Struct()
	return Port_getStreams_Params{s}, err
}

type Port_getStreams_Results struct{ capnp.Struct }

func NewPort_getStreams_Results(s *capnp.Segment) (Port_getStreams_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_getStreams_Results{}, err
	}
	return Port_getStreams_Results{st}, nil
}

func NewRootPort_getStreams_Results(s *capnp.Segment) (Port_getStreams_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_getStreams_Results{}, err
	}
	return Port_getStreams_Results{st}, nil
}

func ReadRootPort_getStreams_Results(msg *capnp.Message) (Port_getStreams_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_getStreams_Results{}, err
	}
	return Port_getStreams_Results{root.Struct()}, nil
}
func (s Port_getStreams_Results) Streams() (capnp.PointerList, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return capnp.PointerList{}, err
	}
	return capnp.PointerList{List: p.List()}, nil
}

func (s Port_getStreams_Results) HasStreams() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Port_getStreams_Results) SetStreams(v capnp.PointerList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewStreams sets the streams field to a newly
// allocated capnp.PointerList, preferring placement in s's segment.
func (s Port_getStreams_Results) NewStreams(n int32) (capnp.PointerList, error) {
	l, err := capnp.NewPointerList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.PointerList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Port_getStreams_Results_List is a list of Port_getStreams_Results.
type Port_getStreams_Results_List struct{ capnp.List }

// NewPort_getStreams_Results creates a new list of Port_getStreams_Results.
func NewPort_getStreams_Results_List(s *capnp.Segment, sz int32) (Port_getStreams_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Port_getStreams_Results_List{}, err
	}
	return Port_getStreams_Results_List{l}, nil
}

func (s Port_getStreams_Results_List) At(i int) Port_getStreams_Results {
	return Port_getStreams_Results{s.List.Struct(i)}
}
func (s Port_getStreams_Results_List) Set(i int, v Port_getStreams_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_getStreams_Results_Promise is a wrapper for a Port_getStreams_Results promised by a client call.
type Port_getStreams_Results_Promise struct{ *capnp.Pipeline }

func (p Port_getStreams_Results_Promise) Struct() (Port_getStreams_Results, error) {
	s, err := p.Pipeline.Struct()
	return Port_getStreams_Results{s}, err
}

type Port_newStream_Params struct{ capnp.Struct }

func NewPort_newStream_Params(s *capnp.Segment) (Port_newStream_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_newStream_Params{}, err
	}
	return Port_newStream_Params{st}, nil
}

func NewRootPort_newStream_Params(s *capnp.Segment) (Port_newStream_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_newStream_Params{}, err
	}
	return Port_newStream_Params{st}, nil
}

func ReadRootPort_newStream_Params(msg *capnp.Message) (Port_newStream_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_newStream_Params{}, err
	}
	return Port_newStream_Params{root.Struct()}, nil
}

// Port_newStream_Params_List is a list of Port_newStream_Params.
type Port_newStream_Params_List struct{ capnp.List }

// NewPort_newStream_Params creates a new list of Port_newStream_Params.
func NewPort_newStream_Params_List(s *capnp.Segment, sz int32) (Port_newStream_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Port_newStream_Params_List{}, err
	}
	return Port_newStream_Params_List{l}, nil
}

func (s Port_newStream_Params_List) At(i int) Port_newStream_Params {
	return Port_newStream_Params{s.List.Struct(i)}
}
func (s Port_newStream_Params_List) Set(i int, v Port_newStream_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_newStream_Params_Promise is a wrapper for a Port_newStream_Params promised by a client call.
type Port_newStream_Params_Promise struct{ *capnp.Pipeline }

func (p Port_newStream_Params_Promise) Struct() (Port_newStream_Params, error) {
	s, err := p.Pipeline.Struct()
	return Port_newStream_Params{s}, err
}

type Port_newStream_Results struct{ capnp.Struct }

func NewPort_newStream_Results(s *capnp.Segment) (Port_newStream_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_newStream_Results{}, err
	}
	return Port_newStream_Results{st}, nil
}

func NewRootPort_newStream_Results(s *capnp.Segment) (Port_newStream_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_newStream_Results{}, err
	}
	return Port_newStream_Results{st}, nil
}

func ReadRootPort_newStream_Results(msg *capnp.Message) (Port_newStream_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_newStream_Results{}, err
	}
	return Port_newStream_Results{root.Struct()}, nil
}
func (s Port_newStream_Results) Stream() Stream {
	p, err := s.Struct.Ptr(0)
	if err != nil {

		return Stream{}
	}
	return Stream{Client: p.Interface().Client()}
}

func (s Port_newStream_Results) HasStream() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Port_newStream_Results) SetStream(v Stream) error {
	seg := s.Segment()
	if seg == nil {

		return nil
	}
	var in capnp.Interface
	if v.Client != nil {
		in = capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	}
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Port_newStream_Results_List is a list of Port_newStream_Results.
type Port_newStream_Results_List struct{ capnp.List }

// NewPort_newStream_Results creates a new list of Port_newStream_Results.
func NewPort_newStream_Results_List(s *capnp.Segment, sz int32) (Port_newStream_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Port_newStream_Results_List{}, err
	}
	return Port_newStream_Results_List{l}, nil
}

func (s Port_newStream_Results_List) At(i int) Port_newStream_Results {
	return Port_newStream_Results{s.List.Struct(i)}
}
func (s Port_newStream_Results_List) Set(i int, v Port_newStream_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_newStream_Results_Promise is a wrapper for a Port_newStream_Results promised by a client call.
type Port_newStream_Results_Promise struct{ *capnp.Pipeline }

func (p Port_newStream_Results_Promise) Struct() (Port_newStream_Results, error) {
	s, err := p.Pipeline.Struct()
	return Port_newStream_Results{s}, err
}

func (p Port_newStream_Results_Promise) Stream() Stream {
	return Stream{Client: p.Pipeline.GetPipeline(0).Client()}
}

type Port_delStream_Params struct{ capnp.Struct }

func NewPort_delStream_Params(s *capnp.Segment) (Port_delStream_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_delStream_Params{}, err
	}
	return Port_delStream_Params{st}, nil
}

func NewRootPort_delStream_Params(s *capnp.Segment) (Port_delStream_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_delStream_Params{}, err
	}
	return Port_delStream_Params{st}, nil
}

func ReadRootPort_delStream_Params(msg *capnp.Message) (Port_delStream_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_delStream_Params{}, err
	}
	return Port_delStream_Params{root.Struct()}, nil
}
func (s Port_delStream_Params) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Port_delStream_Params) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Port_delStream_Params) NameBytes() ([]byte, error) {
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

func (s Port_delStream_Params) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Port_delStream_Params_List is a list of Port_delStream_Params.
type Port_delStream_Params_List struct{ capnp.List }

// NewPort_delStream_Params creates a new list of Port_delStream_Params.
func NewPort_delStream_Params_List(s *capnp.Segment, sz int32) (Port_delStream_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Port_delStream_Params_List{}, err
	}
	return Port_delStream_Params_List{l}, nil
}

func (s Port_delStream_Params_List) At(i int) Port_delStream_Params {
	return Port_delStream_Params{s.List.Struct(i)}
}
func (s Port_delStream_Params_List) Set(i int, v Port_delStream_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_delStream_Params_Promise is a wrapper for a Port_delStream_Params promised by a client call.
type Port_delStream_Params_Promise struct{ *capnp.Pipeline }

func (p Port_delStream_Params_Promise) Struct() (Port_delStream_Params, error) {
	s, err := p.Pipeline.Struct()
	return Port_delStream_Params{s}, err
}

type Port_delStream_Results struct{ capnp.Struct }

func NewPort_delStream_Results(s *capnp.Segment) (Port_delStream_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_delStream_Results{}, err
	}
	return Port_delStream_Results{st}, nil
}

func NewRootPort_delStream_Results(s *capnp.Segment) (Port_delStream_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_delStream_Results{}, err
	}
	return Port_delStream_Results{st}, nil
}

func ReadRootPort_delStream_Results(msg *capnp.Message) (Port_delStream_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_delStream_Results{}, err
	}
	return Port_delStream_Results{root.Struct()}, nil
}

// Port_delStream_Results_List is a list of Port_delStream_Results.
type Port_delStream_Results_List struct{ capnp.List }

// NewPort_delStream_Results creates a new list of Port_delStream_Results.
func NewPort_delStream_Results_List(s *capnp.Segment, sz int32) (Port_delStream_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Port_delStream_Results_List{}, err
	}
	return Port_delStream_Results_List{l}, nil
}

func (s Port_delStream_Results_List) At(i int) Port_delStream_Results {
	return Port_delStream_Results{s.List.Struct(i)}
}
func (s Port_delStream_Results_List) Set(i int, v Port_delStream_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_delStream_Results_Promise is a wrapper for a Port_delStream_Results promised by a client call.
type Port_delStream_Results_Promise struct{ *capnp.Pipeline }

func (p Port_delStream_Results_Promise) Struct() (Port_delStream_Results, error) {
	s, err := p.Pipeline.Struct()
	return Port_delStream_Results{s}, err
}

type Stream struct{ Client capnp.Client }

func (c Stream) GetConfig(ctx context.Context, params func(Stream_getConfig_Params) error, opts ...capnp.CallOption) Stream_getConfig_Results_Promise {
	if c.Client == nil {
		return Stream_getConfig_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xded84ab707587279,
			MethodID:      0,
			InterfaceName: "main.capnp:Stream",
			MethodName:    "getConfig",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Stream_getConfig_Params{Struct: s}) }
	}
	return Stream_getConfig_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Stream) SetConfig(ctx context.Context, params func(Stream_setConfig_Params) error, opts ...capnp.CallOption) Stream_setConfig_Results_Promise {
	if c.Client == nil {
		return Stream_setConfig_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xded84ab707587279,
			MethodID:      1,
			InterfaceName: "main.capnp:Stream",
			MethodName:    "setConfig",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Stream_setConfig_Params{Struct: s}) }
	}
	return Stream_setConfig_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Stream_Server interface {
	GetConfig(Stream_getConfig) error

	SetConfig(Stream_setConfig) error
}

func Stream_ServerToClient(s Stream_Server) Stream {
	c, _ := s.(server.Closer)
	return Stream{Client: server.New(Stream_Methods(nil, s), c)}
}

func Stream_Methods(methods []server.Method, s Stream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xded84ab707587279,
			MethodID:      0,
			InterfaceName: "main.capnp:Stream",
			MethodName:    "getConfig",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Stream_getConfig{c, opts, Stream_getConfig_Params{Struct: p}, Stream_getConfig_Results{Struct: r}}
			return s.GetConfig(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xded84ab707587279,
			MethodID:      1,
			InterfaceName: "main.capnp:Stream",
			MethodName:    "setConfig",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Stream_setConfig{c, opts, Stream_setConfig_Params{Struct: p}, Stream_setConfig_Results{Struct: r}}
			return s.SetConfig(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 0},
	})

	return methods
}

// Stream_getConfig holds the arguments for a server call to Stream.getConfig.
type Stream_getConfig struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Stream_getConfig_Params
	Results Stream_getConfig_Results
}

// Stream_setConfig holds the arguments for a server call to Stream.setConfig.
type Stream_setConfig struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Stream_setConfig_Params
	Results Stream_setConfig_Results
}

type Stream_Config struct{ capnp.Struct }

func NewStream_Config(s *capnp.Segment) (Stream_Config, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Stream_Config{}, err
	}
	return Stream_Config{st}, nil
}

func NewRootStream_Config(s *capnp.Segment) (Stream_Config, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Stream_Config{}, err
	}
	return Stream_Config{st}, nil
}

func ReadRootStream_Config(msg *capnp.Message) (Stream_Config, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_Config{}, err
	}
	return Stream_Config{root.Struct()}, nil
}
func (s Stream_Config) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Stream_Config) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Stream_Config) NameBytes() ([]byte, error) {
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

func (s Stream_Config) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Stream_Config) Loop() bool {
	return s.Struct.Bit(0)
}

func (s Stream_Config) SetLoop(v bool) {
	s.Struct.SetBit(0, v)
}

func (s Stream_Config) Repeat() uint32 {
	return s.Struct.Uint32(4)
}

func (s Stream_Config) SetRepeat(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s Stream_Config) PacketsPerSec() uint32 {
	return s.Struct.Uint32(8)
}

func (s Stream_Config) SetPacketsPerSec(v uint32) {
	s.Struct.SetUint32(8, v)
}

// Stream_Config_List is a list of Stream_Config.
type Stream_Config_List struct{ capnp.List }

// NewStream_Config creates a new list of Stream_Config.
func NewStream_Config_List(s *capnp.Segment, sz int32) (Stream_Config_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	if err != nil {
		return Stream_Config_List{}, err
	}
	return Stream_Config_List{l}, nil
}

func (s Stream_Config_List) At(i int) Stream_Config           { return Stream_Config{s.List.Struct(i)} }
func (s Stream_Config_List) Set(i int, v Stream_Config) error { return s.List.SetStruct(i, v.Struct) }

// Stream_Config_Promise is a wrapper for a Stream_Config promised by a client call.
type Stream_Config_Promise struct{ *capnp.Pipeline }

func (p Stream_Config_Promise) Struct() (Stream_Config, error) {
	s, err := p.Pipeline.Struct()
	return Stream_Config{s}, err
}

type Stream_getConfig_Params struct{ capnp.Struct }

func NewStream_getConfig_Params(s *capnp.Segment) (Stream_getConfig_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Stream_getConfig_Params{}, err
	}
	return Stream_getConfig_Params{st}, nil
}

func NewRootStream_getConfig_Params(s *capnp.Segment) (Stream_getConfig_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Stream_getConfig_Params{}, err
	}
	return Stream_getConfig_Params{st}, nil
}

func ReadRootStream_getConfig_Params(msg *capnp.Message) (Stream_getConfig_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_getConfig_Params{}, err
	}
	return Stream_getConfig_Params{root.Struct()}, nil
}

// Stream_getConfig_Params_List is a list of Stream_getConfig_Params.
type Stream_getConfig_Params_List struct{ capnp.List }

// NewStream_getConfig_Params creates a new list of Stream_getConfig_Params.
func NewStream_getConfig_Params_List(s *capnp.Segment, sz int32) (Stream_getConfig_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Stream_getConfig_Params_List{}, err
	}
	return Stream_getConfig_Params_List{l}, nil
}

func (s Stream_getConfig_Params_List) At(i int) Stream_getConfig_Params {
	return Stream_getConfig_Params{s.List.Struct(i)}
}
func (s Stream_getConfig_Params_List) Set(i int, v Stream_getConfig_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Stream_getConfig_Params_Promise is a wrapper for a Stream_getConfig_Params promised by a client call.
type Stream_getConfig_Params_Promise struct{ *capnp.Pipeline }

func (p Stream_getConfig_Params_Promise) Struct() (Stream_getConfig_Params, error) {
	s, err := p.Pipeline.Struct()
	return Stream_getConfig_Params{s}, err
}

type Stream_getConfig_Results struct{ capnp.Struct }

func NewStream_getConfig_Results(s *capnp.Segment) (Stream_getConfig_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Stream_getConfig_Results{}, err
	}
	return Stream_getConfig_Results{st}, nil
}

func NewRootStream_getConfig_Results(s *capnp.Segment) (Stream_getConfig_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Stream_getConfig_Results{}, err
	}
	return Stream_getConfig_Results{st}, nil
}

func ReadRootStream_getConfig_Results(msg *capnp.Message) (Stream_getConfig_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_getConfig_Results{}, err
	}
	return Stream_getConfig_Results{root.Struct()}, nil
}
func (s Stream_getConfig_Results) Config() (Stream_Config, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Stream_Config{}, err
	}
	return Stream_Config{Struct: p.Struct()}, nil
}

func (s Stream_getConfig_Results) HasConfig() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Stream_getConfig_Results) SetConfig(v Stream_Config) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewConfig sets the config field to a newly
// allocated Stream_Config struct, preferring placement in s's segment.
func (s Stream_getConfig_Results) NewConfig() (Stream_Config, error) {
	ss, err := NewStream_Config(s.Struct.Segment())
	if err != nil {
		return Stream_Config{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Stream_getConfig_Results_List is a list of Stream_getConfig_Results.
type Stream_getConfig_Results_List struct{ capnp.List }

// NewStream_getConfig_Results creates a new list of Stream_getConfig_Results.
func NewStream_getConfig_Results_List(s *capnp.Segment, sz int32) (Stream_getConfig_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Stream_getConfig_Results_List{}, err
	}
	return Stream_getConfig_Results_List{l}, nil
}

func (s Stream_getConfig_Results_List) At(i int) Stream_getConfig_Results {
	return Stream_getConfig_Results{s.List.Struct(i)}
}
func (s Stream_getConfig_Results_List) Set(i int, v Stream_getConfig_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Stream_getConfig_Results_Promise is a wrapper for a Stream_getConfig_Results promised by a client call.
type Stream_getConfig_Results_Promise struct{ *capnp.Pipeline }

func (p Stream_getConfig_Results_Promise) Struct() (Stream_getConfig_Results, error) {
	s, err := p.Pipeline.Struct()
	return Stream_getConfig_Results{s}, err
}

func (p Stream_getConfig_Results_Promise) Config() Stream_Config_Promise {
	return Stream_Config_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Stream_setConfig_Params struct{ capnp.Struct }

func NewStream_setConfig_Params(s *capnp.Segment) (Stream_setConfig_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Stream_setConfig_Params{}, err
	}
	return Stream_setConfig_Params{st}, nil
}

func NewRootStream_setConfig_Params(s *capnp.Segment) (Stream_setConfig_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Stream_setConfig_Params{}, err
	}
	return Stream_setConfig_Params{st}, nil
}

func ReadRootStream_setConfig_Params(msg *capnp.Message) (Stream_setConfig_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_setConfig_Params{}, err
	}
	return Stream_setConfig_Params{root.Struct()}, nil
}
func (s Stream_setConfig_Params) Config() (Stream_Config, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Stream_Config{}, err
	}
	return Stream_Config{Struct: p.Struct()}, nil
}

func (s Stream_setConfig_Params) HasConfig() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Stream_setConfig_Params) SetConfig(v Stream_Config) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewConfig sets the config field to a newly
// allocated Stream_Config struct, preferring placement in s's segment.
func (s Stream_setConfig_Params) NewConfig() (Stream_Config, error) {
	ss, err := NewStream_Config(s.Struct.Segment())
	if err != nil {
		return Stream_Config{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// Stream_setConfig_Params_List is a list of Stream_setConfig_Params.
type Stream_setConfig_Params_List struct{ capnp.List }

// NewStream_setConfig_Params creates a new list of Stream_setConfig_Params.
func NewStream_setConfig_Params_List(s *capnp.Segment, sz int32) (Stream_setConfig_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Stream_setConfig_Params_List{}, err
	}
	return Stream_setConfig_Params_List{l}, nil
}

func (s Stream_setConfig_Params_List) At(i int) Stream_setConfig_Params {
	return Stream_setConfig_Params{s.List.Struct(i)}
}
func (s Stream_setConfig_Params_List) Set(i int, v Stream_setConfig_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Stream_setConfig_Params_Promise is a wrapper for a Stream_setConfig_Params promised by a client call.
type Stream_setConfig_Params_Promise struct{ *capnp.Pipeline }

func (p Stream_setConfig_Params_Promise) Struct() (Stream_setConfig_Params, error) {
	s, err := p.Pipeline.Struct()
	return Stream_setConfig_Params{s}, err
}

func (p Stream_setConfig_Params_Promise) Config() Stream_Config_Promise {
	return Stream_Config_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type Stream_setConfig_Results struct{ capnp.Struct }

func NewStream_setConfig_Results(s *capnp.Segment) (Stream_setConfig_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Stream_setConfig_Results{}, err
	}
	return Stream_setConfig_Results{st}, nil
}

func NewRootStream_setConfig_Results(s *capnp.Segment) (Stream_setConfig_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Stream_setConfig_Results{}, err
	}
	return Stream_setConfig_Results{st}, nil
}

func ReadRootStream_setConfig_Results(msg *capnp.Message) (Stream_setConfig_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_setConfig_Results{}, err
	}
	return Stream_setConfig_Results{root.Struct()}, nil
}

// Stream_setConfig_Results_List is a list of Stream_setConfig_Results.
type Stream_setConfig_Results_List struct{ capnp.List }

// NewStream_setConfig_Results creates a new list of Stream_setConfig_Results.
func NewStream_setConfig_Results_List(s *capnp.Segment, sz int32) (Stream_setConfig_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Stream_setConfig_Results_List{}, err
	}
	return Stream_setConfig_Results_List{l}, nil
}

func (s Stream_setConfig_Results_List) At(i int) Stream_setConfig_Results {
	return Stream_setConfig_Results{s.List.Struct(i)}
}
func (s Stream_setConfig_Results_List) Set(i int, v Stream_setConfig_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Stream_setConfig_Results_Promise is a wrapper for a Stream_setConfig_Results promised by a client call.
type Stream_setConfig_Results_Promise struct{ *capnp.Pipeline }

func (p Stream_setConfig_Results_Promise) Struct() (Stream_setConfig_Results, error) {
	s, err := p.Pipeline.Struct()
	return Stream_setConfig_Results{s}, err
}

type Field8 struct{ capnp.Struct }

func NewField8(s *capnp.Segment) (Field8, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Field8{}, err
	}
	return Field8{st}, nil
}

func NewRootField8(s *capnp.Segment) (Field8, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Field8{}, err
	}
	return Field8{st}, nil
}

func ReadRootField8(msg *capnp.Message) (Field8, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Field8{}, err
	}
	return Field8{root.Struct()}, nil
}
func (s Field8) Value() uint8 {
	return s.Struct.Uint8(0)
}

func (s Field8) SetValue(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Field8) Step() uint8 {
	return s.Struct.Uint8(1)
}

func (s Field8) SetStep(v uint8) {
	s.Struct.SetUint8(1, v)
}

func (s Field8) Mode() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Field8) HasMode() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Field8) ModeBytes() ([]byte, error) {
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

func (s Field8) SetMode(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Field8) Count() uint8 {
	return s.Struct.Uint8(2)
}

func (s Field8) SetCount(v uint8) {
	s.Struct.SetUint8(2, v)
}

// Field8_List is a list of Field8.
type Field8_List struct{ capnp.List }

// NewField8 creates a new list of Field8.
func NewField8_List(s *capnp.Segment, sz int32) (Field8_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return Field8_List{}, err
	}
	return Field8_List{l}, nil
}

func (s Field8_List) At(i int) Field8           { return Field8{s.List.Struct(i)} }
func (s Field8_List) Set(i int, v Field8) error { return s.List.SetStruct(i, v.Struct) }

// Field8_Promise is a wrapper for a Field8 promised by a client call.
type Field8_Promise struct{ *capnp.Pipeline }

func (p Field8_Promise) Struct() (Field8, error) {
	s, err := p.Pipeline.Struct()
	return Field8{s}, err
}

type Field16 struct{ capnp.Struct }

func NewField16(s *capnp.Segment) (Field16, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Field16{}, err
	}
	return Field16{st}, nil
}

func NewRootField16(s *capnp.Segment) (Field16, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Field16{}, err
	}
	return Field16{st}, nil
}

func ReadRootField16(msg *capnp.Message) (Field16, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Field16{}, err
	}
	return Field16{root.Struct()}, nil
}
func (s Field16) Value() uint16 {
	return s.Struct.Uint16(0)
}

func (s Field16) SetValue(v uint16) {
	s.Struct.SetUint16(0, v)
}

func (s Field16) Step() uint16 {
	return s.Struct.Uint16(2)
}

func (s Field16) SetStep(v uint16) {
	s.Struct.SetUint16(2, v)
}

func (s Field16) Mode() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Field16) HasMode() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Field16) ModeBytes() ([]byte, error) {
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

func (s Field16) SetMode(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Field16) Count() uint16 {
	return s.Struct.Uint16(4)
}

func (s Field16) SetCount(v uint16) {
	s.Struct.SetUint16(4, v)
}

// Field16_List is a list of Field16.
type Field16_List struct{ capnp.List }

// NewField16 creates a new list of Field16.
func NewField16_List(s *capnp.Segment, sz int32) (Field16_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return Field16_List{}, err
	}
	return Field16_List{l}, nil
}

func (s Field16_List) At(i int) Field16           { return Field16{s.List.Struct(i)} }
func (s Field16_List) Set(i int, v Field16) error { return s.List.SetStruct(i, v.Struct) }

// Field16_Promise is a wrapper for a Field16 promised by a client call.
type Field16_Promise struct{ *capnp.Pipeline }

func (p Field16_Promise) Struct() (Field16, error) {
	s, err := p.Pipeline.Struct()
	return Field16{s}, err
}

type Field32 struct{ capnp.Struct }

func NewField32(s *capnp.Segment) (Field32, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Field32{}, err
	}
	return Field32{st}, nil
}

func NewRootField32(s *capnp.Segment) (Field32, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	if err != nil {
		return Field32{}, err
	}
	return Field32{st}, nil
}

func ReadRootField32(msg *capnp.Message) (Field32, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Field32{}, err
	}
	return Field32{root.Struct()}, nil
}
func (s Field32) Value() uint32 {
	return s.Struct.Uint32(0)
}

func (s Field32) SetValue(v uint32) {
	s.Struct.SetUint32(0, v)
}

func (s Field32) Step() uint32 {
	return s.Struct.Uint32(4)
}

func (s Field32) SetStep(v uint32) {
	s.Struct.SetUint32(4, v)
}

func (s Field32) Mode() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Field32) HasMode() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Field32) ModeBytes() ([]byte, error) {
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

func (s Field32) SetMode(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Field32) Count() uint32 {
	return s.Struct.Uint32(8)
}

func (s Field32) SetCount(v uint32) {
	s.Struct.SetUint32(8, v)
}

// Field32_List is a list of Field32.
type Field32_List struct{ capnp.List }

// NewField32 creates a new list of Field32.
func NewField32_List(s *capnp.Segment, sz int32) (Field32_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	if err != nil {
		return Field32_List{}, err
	}
	return Field32_List{l}, nil
}

func (s Field32_List) At(i int) Field32           { return Field32{s.List.Struct(i)} }
func (s Field32_List) Set(i int, v Field32) error { return s.List.SetStruct(i, v.Struct) }

// Field32_Promise is a wrapper for a Field32 promised by a client call.
type Field32_Promise struct{ *capnp.Pipeline }

func (p Field32_Promise) Struct() (Field32, error) {
	s, err := p.Pipeline.Struct()
	return Field32{s}, err
}

type Field64 struct{ capnp.Struct }

func NewField64(s *capnp.Segment) (Field64, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	if err != nil {
		return Field64{}, err
	}
	return Field64{st}, nil
}

func NewRootField64(s *capnp.Segment) (Field64, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	if err != nil {
		return Field64{}, err
	}
	return Field64{st}, nil
}

func ReadRootField64(msg *capnp.Message) (Field64, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Field64{}, err
	}
	return Field64{root.Struct()}, nil
}
func (s Field64) Value() uint64 {
	return s.Struct.Uint64(0)
}

func (s Field64) SetValue(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Field64) Step() uint64 {
	return s.Struct.Uint64(8)
}

func (s Field64) SetStep(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Field64) Mode() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Field64) HasMode() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Field64) ModeBytes() ([]byte, error) {
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

func (s Field64) SetMode(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Field64) Count() uint64 {
	return s.Struct.Uint64(16)
}

func (s Field64) SetCount(v uint64) {
	s.Struct.SetUint64(16, v)
}

// Field64_List is a list of Field64.
type Field64_List struct{ capnp.List }

// NewField64 creates a new list of Field64.
func NewField64_List(s *capnp.Segment, sz int32) (Field64_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1}, sz)
	if err != nil {
		return Field64_List{}, err
	}
	return Field64_List{l}, nil
}

func (s Field64_List) At(i int) Field64           { return Field64{s.List.Struct(i)} }
func (s Field64_List) Set(i int, v Field64) error { return s.List.SetStruct(i, v.Struct) }

// Field64_Promise is a wrapper for a Field64 promised by a client call.
type Field64_Promise struct{ *capnp.Pipeline }

func (p Field64_Promise) Struct() (Field64, error) {
	s, err := p.Pipeline.Struct()
	return Field64{s}, err
}

type LongField struct{ capnp.Struct }

func NewLongField(s *capnp.Segment) (LongField, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5})
	if err != nil {
		return LongField{}, err
	}
	return LongField{st}, nil
}

func NewRootLongField(s *capnp.Segment) (LongField, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5})
	if err != nil {
		return LongField{}, err
	}
	return LongField{st}, nil
}

func ReadRootLongField(msg *capnp.Message) (LongField, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return LongField{}, err
	}
	return LongField{root.Struct()}, nil
}
func (s LongField) Value() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s LongField) HasValue() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s LongField) SetValue(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s LongField) Step() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s LongField) HasStep() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s LongField) SetStep(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, d.List.ToPtr())
}

func (s LongField) Mode() (string, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s LongField) HasMode() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s LongField) ModeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s LongField) SetMode(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, t.List.ToPtr())
}

func (s LongField) Count() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s LongField) HasCount() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s LongField) SetCount(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, d.List.ToPtr())
}

func (s LongField) Mask() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s LongField) HasMask() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s LongField) SetMask(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(4, d.List.ToPtr())
}

// LongField_List is a list of LongField.
type LongField_List struct{ capnp.List }

// NewLongField creates a new list of LongField.
func NewLongField_List(s *capnp.Segment, sz int32) (LongField_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 5}, sz)
	if err != nil {
		return LongField_List{}, err
	}
	return LongField_List{l}, nil
}

func (s LongField_List) At(i int) LongField           { return LongField{s.List.Struct(i)} }
func (s LongField_List) Set(i int, v LongField) error { return s.List.SetStruct(i, v.Struct) }

// LongField_Promise is a wrapper for a LongField promised by a client call.
type LongField_Promise struct{ *capnp.Pipeline }

func (p LongField_Promise) Struct() (LongField, error) {
	s, err := p.Pipeline.Struct()
	return LongField{s}, err
}

type Protocol struct{ capnp.Struct }
type Protocol_ethernet Protocol
type Protocol_ipv4 Protocol
type Protocol_Which uint16

const (
	Protocol_Which_ethernet Protocol_Which = 0
	Protocol_Which_ipv4     Protocol_Which = 1
)

func (w Protocol_Which) String() string {
	const s = "ethernetipv4"
	switch w {
	case Protocol_Which_ethernet:
		return s[0:8]
	case Protocol_Which_ipv4:
		return s[8:12]

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
func (s Protocol) Ethernet() Protocol_ethernet { return Protocol_ethernet(s) }
func (s Protocol) SetEthernet() {
	s.Struct.SetUint16(0, 0)
}
func (s Protocol_ethernet) Source() (LongField, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return LongField{}, err
	}
	return LongField{Struct: p.Struct()}, nil
}

func (s Protocol_ethernet) HasSource() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet) SetSource(v LongField) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSource sets the source field to a newly
// allocated LongField struct, preferring placement in s's segment.
func (s Protocol_ethernet) NewSource() (LongField, error) {
	ss, err := NewLongField(s.Struct.Segment())
	if err != nil {
		return LongField{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ethernet) Destination() (LongField, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return LongField{}, err
	}
	return LongField{Struct: p.Struct()}, nil
}

func (s Protocol_ethernet) HasDestination() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet) SetDestination(v LongField) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDestination sets the destination field to a newly
// allocated LongField struct, preferring placement in s's segment.
func (s Protocol_ethernet) NewDestination() (LongField, error) {
	ss, err := NewLongField(s.Struct.Segment())
	if err != nil {
		return LongField{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ethernet) EthernetType() (Field16, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return Field16{}, err
	}
	return Field16{Struct: p.Struct()}, nil
}

func (s Protocol_ethernet) HasEthernetType() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet) SetEthernetType(v Field16) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewEthernetType sets the ethernetType field to a newly
// allocated Field16 struct, preferring placement in s's segment.
func (s Protocol_ethernet) NewEthernetType() (Field16, error) {
	ss, err := NewField16(s.Struct.Segment())
	if err != nil {
		return Field16{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ethernet) Length() (Field16, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return Field16{}, err
	}
	return Field16{Struct: p.Struct()}, nil
}

func (s Protocol_ethernet) HasLength() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet) SetLength(v Field16) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewLength sets the length field to a newly
// allocated Field16 struct, preferring placement in s's segment.
func (s Protocol_ethernet) NewLength() (Field16, error) {
	ss, err := NewField16(s.Struct.Segment())
	if err != nil {
		return Field16{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol) Ipv4() Protocol_ipv4 { return Protocol_ipv4(s) }
func (s Protocol) SetIpv4() {
	s.Struct.SetUint16(0, 1)
}
func (s Protocol_ipv4) Version() (Field8, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasVersion() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetVersion(v Field8) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewVersion sets the version field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewVersion() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Ihl() (Field8, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasIhl() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetIhl(v Field8) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewIhl sets the ihl field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewIhl() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Tos() (Field8, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasTos() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetTos(v Field8) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewTos sets the tos field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewTos() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Length() (Field8, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasLength() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetLength(v Field8) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewLength sets the length field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewLength() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Id() (Field8, error) {
	p, err := s.Struct.Ptr(4)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasId() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetId(v Field8) error {
	return s.Struct.SetPtr(4, v.Struct.ToPtr())
}

// NewId sets the id field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewId() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(4, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Flags() (Field8, error) {
	p, err := s.Struct.Ptr(5)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasFlags() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetFlags(v Field8) error {
	return s.Struct.SetPtr(5, v.Struct.ToPtr())
}

// NewFlags sets the flags field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewFlags() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(5, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) FragOffset() (Field16, error) {
	p, err := s.Struct.Ptr(6)
	if err != nil {
		return Field16{}, err
	}
	return Field16{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasFragOffset() bool {
	p, err := s.Struct.Ptr(6)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetFragOffset(v Field16) error {
	return s.Struct.SetPtr(6, v.Struct.ToPtr())
}

// NewFragOffset sets the fragOffset field to a newly
// allocated Field16 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewFragOffset() (Field16, error) {
	ss, err := NewField16(s.Struct.Segment())
	if err != nil {
		return Field16{}, err
	}
	err = s.Struct.SetPtr(6, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Ttl() (Field8, error) {
	p, err := s.Struct.Ptr(7)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasTtl() bool {
	p, err := s.Struct.Ptr(7)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetTtl(v Field8) error {
	return s.Struct.SetPtr(7, v.Struct.ToPtr())
}

// NewTtl sets the ttl field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewTtl() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(7, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Protocol() (Field8, error) {
	p, err := s.Struct.Ptr(8)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasProtocol() bool {
	p, err := s.Struct.Ptr(8)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetProtocol(v Field8) error {
	return s.Struct.SetPtr(8, v.Struct.ToPtr())
}

// NewProtocol sets the protocol field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewProtocol() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(8, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Checksum() (Field8, error) {
	p, err := s.Struct.Ptr(9)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasChecksum() bool {
	p, err := s.Struct.Ptr(9)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetChecksum(v Field8) error {
	return s.Struct.SetPtr(9, v.Struct.ToPtr())
}

// NewChecksum sets the checksum field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewChecksum() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
	}
	err = s.Struct.SetPtr(9, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Srcip() (Field32, error) {
	p, err := s.Struct.Ptr(10)
	if err != nil {
		return Field32{}, err
	}
	return Field32{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasSrcip() bool {
	p, err := s.Struct.Ptr(10)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetSrcip(v Field32) error {
	return s.Struct.SetPtr(10, v.Struct.ToPtr())
}

// NewSrcip sets the srcip field to a newly
// allocated Field32 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewSrcip() (Field32, error) {
	ss, err := NewField32(s.Struct.Segment())
	if err != nil {
		return Field32{}, err
	}
	err = s.Struct.SetPtr(10, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Dstip() (Field32, error) {
	p, err := s.Struct.Ptr(11)
	if err != nil {
		return Field32{}, err
	}
	return Field32{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasDstip() bool {
	p, err := s.Struct.Ptr(11)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetDstip(v Field32) error {
	return s.Struct.SetPtr(11, v.Struct.ToPtr())
}

// NewDstip sets the dstip field to a newly
// allocated Field32 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewDstip() (Field32, error) {
	ss, err := NewField32(s.Struct.Segment())
	if err != nil {
		return Field32{}, err
	}
	err = s.Struct.SetPtr(11, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Options() (LongField, error) {
	p, err := s.Struct.Ptr(12)
	if err != nil {
		return LongField{}, err
	}
	return LongField{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasOptions() bool {
	p, err := s.Struct.Ptr(12)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetOptions(v LongField) error {
	return s.Struct.SetPtr(12, v.Struct.ToPtr())
}

// NewOptions sets the options field to a newly
// allocated LongField struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewOptions() (LongField, error) {
	ss, err := NewLongField(s.Struct.Segment())
	if err != nil {
		return LongField{}, err
	}
	err = s.Struct.SetPtr(12, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Padding() (Field8, error) {
	p, err := s.Struct.Ptr(13)
	if err != nil {
		return Field8{}, err
	}
	return Field8{Struct: p.Struct()}, nil
}

func (s Protocol_ipv4) HasPadding() bool {
	p, err := s.Struct.Ptr(13)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetPadding(v Field8) error {
	return s.Struct.SetPtr(13, v.Struct.ToPtr())
}

// NewPadding sets the padding field to a newly
// allocated Field8 struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewPadding() (Field8, error) {
	ss, err := NewField8(s.Struct.Segment())
	if err != nil {
		return Field8{}, err
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

func (p Protocol_Promise) Ethernet() Protocol_ethernet_Promise {
	return Protocol_ethernet_Promise{p.Pipeline}
}

// Protocol_ethernet_Promise is a wrapper for a Protocol_ethernet promised by a client call.
type Protocol_ethernet_Promise struct{ *capnp.Pipeline }

func (p Protocol_ethernet_Promise) Struct() (Protocol_ethernet, error) {
	s, err := p.Pipeline.Struct()
	return Protocol_ethernet{s}, err
}

func (p Protocol_ethernet_Promise) Source() LongField_Promise {
	return LongField_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Protocol_ethernet_Promise) Destination() LongField_Promise {
	return LongField_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Protocol_ethernet_Promise) EthernetType() Field16_Promise {
	return Field16_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Protocol_ethernet_Promise) Length() Field16_Promise {
	return Field16_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Protocol_Promise) Ipv4() Protocol_ipv4_Promise { return Protocol_ipv4_Promise{p.Pipeline} }

// Protocol_ipv4_Promise is a wrapper for a Protocol_ipv4 promised by a client call.
type Protocol_ipv4_Promise struct{ *capnp.Pipeline }

func (p Protocol_ipv4_Promise) Struct() (Protocol_ipv4, error) {
	s, err := p.Pipeline.Struct()
	return Protocol_ipv4{s}, err
}

func (p Protocol_ipv4_Promise) Version() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Protocol_ipv4_Promise) Ihl() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Protocol_ipv4_Promise) Tos() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Protocol_ipv4_Promise) Length() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Protocol_ipv4_Promise) Id() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

func (p Protocol_ipv4_Promise) Flags() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Protocol_ipv4_Promise) FragOffset() Field16_Promise {
	return Field16_Promise{Pipeline: p.Pipeline.GetPipeline(6)}
}

func (p Protocol_ipv4_Promise) Ttl() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(7)}
}

func (p Protocol_ipv4_Promise) Protocol() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(8)}
}

func (p Protocol_ipv4_Promise) Checksum() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(9)}
}

func (p Protocol_ipv4_Promise) Srcip() Field32_Promise {
	return Field32_Promise{Pipeline: p.Pipeline.GetPipeline(10)}
}

func (p Protocol_ipv4_Promise) Dstip() Field32_Promise {
	return Field32_Promise{Pipeline: p.Pipeline.GetPipeline(11)}
}

func (p Protocol_ipv4_Promise) Options() LongField_Promise {
	return LongField_Promise{Pipeline: p.Pipeline.GetPipeline(12)}
}

func (p Protocol_ipv4_Promise) Padding() Field8_Promise {
	return Field8_Promise{Pipeline: p.Pipeline.GetPipeline(13)}
}
