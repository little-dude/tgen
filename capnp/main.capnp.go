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
func (c Stream) GetLayers(ctx context.Context, params func(Stream_getLayers_Params) error, opts ...capnp.CallOption) Stream_getLayers_Results_Promise {
	if c.Client == nil {
		return Stream_getLayers_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xded84ab707587279,
			MethodID:      2,
			InterfaceName: "main.capnp:Stream",
			MethodName:    "getLayers",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Stream_getLayers_Params{Struct: s}) }
	}
	return Stream_getLayers_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c Stream) SetLayers(ctx context.Context, params func(Stream_setLayers_Params) error, opts ...capnp.CallOption) Stream_setLayers_Results_Promise {
	if c.Client == nil {
		return Stream_setLayers_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xded84ab707587279,
			MethodID:      3,
			InterfaceName: "main.capnp:Stream",
			MethodName:    "setLayers",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 1}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Stream_setLayers_Params{Struct: s}) }
	}
	return Stream_setLayers_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Stream_Server interface {
	GetConfig(Stream_getConfig) error

	SetConfig(Stream_setConfig) error

	GetLayers(Stream_getLayers) error

	SetLayers(Stream_setLayers) error
}

func Stream_ServerToClient(s Stream_Server) Stream {
	c, _ := s.(server.Closer)
	return Stream{Client: server.New(Stream_Methods(nil, s), c)}
}

func Stream_Methods(methods []server.Method, s Stream_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 4)
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

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xded84ab707587279,
			MethodID:      2,
			InterfaceName: "main.capnp:Stream",
			MethodName:    "getLayers",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Stream_getLayers{c, opts, Stream_getLayers_Params{Struct: p}, Stream_getLayers_Results{Struct: r}}
			return s.GetLayers(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xded84ab707587279,
			MethodID:      3,
			InterfaceName: "main.capnp:Stream",
			MethodName:    "setLayers",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Stream_setLayers{c, opts, Stream_setLayers_Params{Struct: p}, Stream_setLayers_Results{Struct: r}}
			return s.SetLayers(call)
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

// Stream_getLayers holds the arguments for a server call to Stream.getLayers.
type Stream_getLayers struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Stream_getLayers_Params
	Results Stream_getLayers_Results
}

// Stream_setLayers holds the arguments for a server call to Stream.setLayers.
type Stream_setLayers struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Stream_setLayers_Params
	Results Stream_setLayers_Results
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

type Stream_getLayers_Params struct{ capnp.Struct }

func NewStream_getLayers_Params(s *capnp.Segment) (Stream_getLayers_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Stream_getLayers_Params{}, err
	}
	return Stream_getLayers_Params{st}, nil
}

func NewRootStream_getLayers_Params(s *capnp.Segment) (Stream_getLayers_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Stream_getLayers_Params{}, err
	}
	return Stream_getLayers_Params{st}, nil
}

func ReadRootStream_getLayers_Params(msg *capnp.Message) (Stream_getLayers_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_getLayers_Params{}, err
	}
	return Stream_getLayers_Params{root.Struct()}, nil
}

// Stream_getLayers_Params_List is a list of Stream_getLayers_Params.
type Stream_getLayers_Params_List struct{ capnp.List }

// NewStream_getLayers_Params creates a new list of Stream_getLayers_Params.
func NewStream_getLayers_Params_List(s *capnp.Segment, sz int32) (Stream_getLayers_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Stream_getLayers_Params_List{}, err
	}
	return Stream_getLayers_Params_List{l}, nil
}

func (s Stream_getLayers_Params_List) At(i int) Stream_getLayers_Params {
	return Stream_getLayers_Params{s.List.Struct(i)}
}
func (s Stream_getLayers_Params_List) Set(i int, v Stream_getLayers_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Stream_getLayers_Params_Promise is a wrapper for a Stream_getLayers_Params promised by a client call.
type Stream_getLayers_Params_Promise struct{ *capnp.Pipeline }

func (p Stream_getLayers_Params_Promise) Struct() (Stream_getLayers_Params, error) {
	s, err := p.Pipeline.Struct()
	return Stream_getLayers_Params{s}, err
}

type Stream_getLayers_Results struct{ capnp.Struct }

func NewStream_getLayers_Results(s *capnp.Segment) (Stream_getLayers_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Stream_getLayers_Results{}, err
	}
	return Stream_getLayers_Results{st}, nil
}

func NewRootStream_getLayers_Results(s *capnp.Segment) (Stream_getLayers_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Stream_getLayers_Results{}, err
	}
	return Stream_getLayers_Results{st}, nil
}

func ReadRootStream_getLayers_Results(msg *capnp.Message) (Stream_getLayers_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_getLayers_Results{}, err
	}
	return Stream_getLayers_Results{root.Struct()}, nil
}
func (s Stream_getLayers_Results) Layers() (Protocol_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Protocol_List{}, err
	}
	return Protocol_List{List: p.List()}, nil
}

func (s Stream_getLayers_Results) HasLayers() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Stream_getLayers_Results) SetLayers(v Protocol_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewLayers sets the layers field to a newly
// allocated Protocol_List, preferring placement in s's segment.
func (s Stream_getLayers_Results) NewLayers(n int32) (Protocol_List, error) {
	l, err := NewProtocol_List(s.Struct.Segment(), n)
	if err != nil {
		return Protocol_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Stream_getLayers_Results_List is a list of Stream_getLayers_Results.
type Stream_getLayers_Results_List struct{ capnp.List }

// NewStream_getLayers_Results creates a new list of Stream_getLayers_Results.
func NewStream_getLayers_Results_List(s *capnp.Segment, sz int32) (Stream_getLayers_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Stream_getLayers_Results_List{}, err
	}
	return Stream_getLayers_Results_List{l}, nil
}

func (s Stream_getLayers_Results_List) At(i int) Stream_getLayers_Results {
	return Stream_getLayers_Results{s.List.Struct(i)}
}
func (s Stream_getLayers_Results_List) Set(i int, v Stream_getLayers_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Stream_getLayers_Results_Promise is a wrapper for a Stream_getLayers_Results promised by a client call.
type Stream_getLayers_Results_Promise struct{ *capnp.Pipeline }

func (p Stream_getLayers_Results_Promise) Struct() (Stream_getLayers_Results, error) {
	s, err := p.Pipeline.Struct()
	return Stream_getLayers_Results{s}, err
}

type Stream_setLayers_Params struct{ capnp.Struct }

func NewStream_setLayers_Params(s *capnp.Segment) (Stream_setLayers_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Stream_setLayers_Params{}, err
	}
	return Stream_setLayers_Params{st}, nil
}

func NewRootStream_setLayers_Params(s *capnp.Segment) (Stream_setLayers_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Stream_setLayers_Params{}, err
	}
	return Stream_setLayers_Params{st}, nil
}

func ReadRootStream_setLayers_Params(msg *capnp.Message) (Stream_setLayers_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_setLayers_Params{}, err
	}
	return Stream_setLayers_Params{root.Struct()}, nil
}
func (s Stream_setLayers_Params) Layers() (Protocol_List, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Protocol_List{}, err
	}
	return Protocol_List{List: p.List()}, nil
}

func (s Stream_setLayers_Params) HasLayers() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Stream_setLayers_Params) SetLayers(v Protocol_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewLayers sets the layers field to a newly
// allocated Protocol_List, preferring placement in s's segment.
func (s Stream_setLayers_Params) NewLayers(n int32) (Protocol_List, error) {
	l, err := NewProtocol_List(s.Struct.Segment(), n)
	if err != nil {
		return Protocol_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Stream_setLayers_Params_List is a list of Stream_setLayers_Params.
type Stream_setLayers_Params_List struct{ capnp.List }

// NewStream_setLayers_Params creates a new list of Stream_setLayers_Params.
func NewStream_setLayers_Params_List(s *capnp.Segment, sz int32) (Stream_setLayers_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Stream_setLayers_Params_List{}, err
	}
	return Stream_setLayers_Params_List{l}, nil
}

func (s Stream_setLayers_Params_List) At(i int) Stream_setLayers_Params {
	return Stream_setLayers_Params{s.List.Struct(i)}
}
func (s Stream_setLayers_Params_List) Set(i int, v Stream_setLayers_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Stream_setLayers_Params_Promise is a wrapper for a Stream_setLayers_Params promised by a client call.
type Stream_setLayers_Params_Promise struct{ *capnp.Pipeline }

func (p Stream_setLayers_Params_Promise) Struct() (Stream_setLayers_Params, error) {
	s, err := p.Pipeline.Struct()
	return Stream_setLayers_Params{s}, err
}

type Stream_setLayers_Results struct{ capnp.Struct }

func NewStream_setLayers_Results(s *capnp.Segment) (Stream_setLayers_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Stream_setLayers_Results{}, err
	}
	return Stream_setLayers_Results{st}, nil
}

func NewRootStream_setLayers_Results(s *capnp.Segment) (Stream_setLayers_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Stream_setLayers_Results{}, err
	}
	return Stream_setLayers_Results{st}, nil
}

func ReadRootStream_setLayers_Results(msg *capnp.Message) (Stream_setLayers_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Stream_setLayers_Results{}, err
	}
	return Stream_setLayers_Results{root.Struct()}, nil
}

// Stream_setLayers_Results_List is a list of Stream_setLayers_Results.
type Stream_setLayers_Results_List struct{ capnp.List }

// NewStream_setLayers_Results creates a new list of Stream_setLayers_Results.
func NewStream_setLayers_Results_List(s *capnp.Segment, sz int32) (Stream_setLayers_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Stream_setLayers_Results_List{}, err
	}
	return Stream_setLayers_Results_List{l}, nil
}

func (s Stream_setLayers_Results_List) At(i int) Stream_setLayers_Results {
	return Stream_setLayers_Results{s.List.Struct(i)}
}
func (s Stream_setLayers_Results_List) Set(i int, v Stream_setLayers_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Stream_setLayers_Results_Promise is a wrapper for a Stream_setLayers_Results promised by a client call.
type Stream_setLayers_Results_Promise struct{ *capnp.Pipeline }

func (p Stream_setLayers_Results_Promise) Struct() (Stream_setLayers_Results, error) {
	s, err := p.Pipeline.Struct()
	return Stream_setLayers_Results{s}, err
}

type Field struct{ capnp.Struct }

func NewField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	if err != nil {
		return Field{}, err
	}
	return Field{st}, nil
}

func NewRootField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
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

func (s Field) Mode() (string, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Field) HasMode() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Field) ModeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return nil, err
	}
	d := p.Data()
	if len(d) == 0 {
		return d, nil
	}
	return d[:len(d)-1], nil
}

func (s Field) SetMode(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(1, t.List.ToPtr())
}

func (s Field) Step() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Field) HasStep() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Field) SetStep(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(2, d.List.ToPtr())
}

func (s Field) Mask() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Field) HasMask() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Field) SetMask(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(3, d.List.ToPtr())
}

func (s Field) Count() uint64 {
	return s.Struct.Uint64(0)
}

func (s Field) SetCount(v uint64) {
	s.Struct.SetUint64(0, v)
}

// Field_List is a list of Field.
type Field_List struct{ capnp.List }

// NewField creates a new list of Field.
func NewField_List(s *capnp.Segment, sz int32) (Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
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
	ss, err := p.StructDefault(x_ef97cf4069588836[0:80])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[80:160])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[160:240])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[240:320])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[320:400])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[400:480])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[480:560])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[560:640])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[640:720])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[720:800])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[800:880])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[880:960])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[960:1040])
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

func (s Protocol_ipv4) Srcip() (Field, error) {
	p, err := s.Struct.Ptr(10)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[1040:1120])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasSrcip() bool {
	p, err := s.Struct.Ptr(10)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetSrcip(v Field) error {
	return s.Struct.SetPtr(10, v.Struct.ToPtr())
}

// NewSrcip sets the srcip field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewSrcip() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(10, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ipv4) Dstip() (Field, error) {
	p, err := s.Struct.Ptr(11)
	if err != nil {
		return Field{}, err
	}
	ss, err := p.StructDefault(x_ef97cf4069588836[1120:1200])
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: ss}, nil
}

func (s Protocol_ipv4) HasDstip() bool {
	p, err := s.Struct.Ptr(11)
	return p.IsValid() || err != nil
}

func (s Protocol_ipv4) SetDstip(v Field) error {
	return s.Struct.SetPtr(11, v.Struct.ToPtr())
}

// NewDstip sets the dstip field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ipv4) NewDstip() (Field, error) {
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
	ss, err := p.StructDefault(x_ef97cf4069588836[1200:1280])
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
	ss, err := p.StructDefault(x_ef97cf4069588836[1280:1360])
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
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(0, x_ef97cf4069588836[1360:1440])}
}

func (p Protocol_ethernet2_Promise) Destination() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(1, x_ef97cf4069588836[1440:1520])}
}

func (p Protocol_ethernet2_Promise) EthernetType() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(2, x_ef97cf4069588836[1520:1600])}
}

func (p Protocol_Promise) Ipv4() Protocol_ipv4_Promise { return Protocol_ipv4_Promise{p.Pipeline} }

// Protocol_ipv4_Promise is a wrapper for a Protocol_ipv4 promised by a client call.
type Protocol_ipv4_Promise struct{ *capnp.Pipeline }

func (p Protocol_ipv4_Promise) Struct() (Protocol_ipv4, error) {
	s, err := p.Pipeline.Struct()
	return Protocol_ipv4{s}, err
}

func (p Protocol_ipv4_Promise) Version() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(0, x_ef97cf4069588836[1600:1680])}
}

func (p Protocol_ipv4_Promise) Ihl() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(1, x_ef97cf4069588836[1680:1760])}
}

func (p Protocol_ipv4_Promise) Tos() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(2, x_ef97cf4069588836[1760:1840])}
}

func (p Protocol_ipv4_Promise) Length() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(3, x_ef97cf4069588836[1840:1920])}
}

func (p Protocol_ipv4_Promise) Id() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(4, x_ef97cf4069588836[1920:2000])}
}

func (p Protocol_ipv4_Promise) Flags() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(5, x_ef97cf4069588836[2000:2080])}
}

func (p Protocol_ipv4_Promise) FragOffset() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(6, x_ef97cf4069588836[2080:2160])}
}

func (p Protocol_ipv4_Promise) Ttl() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(7, x_ef97cf4069588836[2160:2240])}
}

func (p Protocol_ipv4_Promise) Protocol() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(8, x_ef97cf4069588836[2240:2320])}
}

func (p Protocol_ipv4_Promise) Checksum() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(9, x_ef97cf4069588836[2320:2400])}
}

func (p Protocol_ipv4_Promise) Srcip() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(10, x_ef97cf4069588836[2400:2480])}
}

func (p Protocol_ipv4_Promise) Dstip() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(11, x_ef97cf4069588836[2480:2560])}
}

func (p Protocol_ipv4_Promise) Options() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(12, x_ef97cf4069588836[2560:2640])}
}

func (p Protocol_ipv4_Promise) Padding() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipelineDefault(13, x_ef97cf4069588836[2640:2720])}
}

var x_ef97cf4069588836 = []byte{
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 50, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 18, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	8, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	5, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	97, 117, 116, 111, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 18, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	97, 117, 116, 111, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	97, 117, 116, 111, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 34, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 34, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 34, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 50, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 255, 255, 255, 255, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 18, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	8, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	4, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	5, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	97, 117, 116, 111, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 18, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 18, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	97, 117, 116, 111, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 42, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	97, 117, 116, 111, 0, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 34, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 34, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 34, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 34, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 255, 255, 255, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 9, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 10, 0, 0, 0,
	13, 0, 0, 0, 50, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	9, 0, 0, 0, 10, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	102, 105, 120, 101, 100, 0, 0, 0,
	255, 0, 0, 0, 0, 0, 0, 0,
}
