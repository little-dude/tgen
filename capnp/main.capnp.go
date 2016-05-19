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

func (c Port) GetName(ctx context.Context, params func(Port_getName_Params) error, opts ...capnp.CallOption) Port_getName_Results_Promise {
	if c.Client == nil {
		return Port_getName_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      0,
			InterfaceName: "main.capnp:Port",
			MethodName:    "getName",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(Port_getName_Params{Struct: s}) }
	}
	return Port_getName_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type Port_Server interface {
	GetName(Port_getName) error
}

func Port_ServerToClient(s Port_Server) Port {
	c, _ := s.(server.Closer)
	return Port{Client: server.New(Port_Methods(nil, s), c)}
}

func Port_Methods(methods []server.Method, s Port_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xaf41a42ec9ad3bcd,
			MethodID:      0,
			InterfaceName: "main.capnp:Port",
			MethodName:    "getName",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := Port_getName{c, opts, Port_getName_Params{Struct: p}, Port_getName_Results{Struct: r}}
			return s.GetName(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// Port_getName holds the arguments for a server call to Port.getName.
type Port_getName struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  Port_getName_Params
	Results Port_getName_Results
}

type Port_getName_Params struct{ capnp.Struct }

func NewPort_getName_Params(s *capnp.Segment) (Port_getName_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_getName_Params{}, err
	}
	return Port_getName_Params{st}, nil
}

func NewRootPort_getName_Params(s *capnp.Segment) (Port_getName_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	if err != nil {
		return Port_getName_Params{}, err
	}
	return Port_getName_Params{st}, nil
}

func ReadRootPort_getName_Params(msg *capnp.Message) (Port_getName_Params, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_getName_Params{}, err
	}
	return Port_getName_Params{root.Struct()}, nil
}

// Port_getName_Params_List is a list of Port_getName_Params.
type Port_getName_Params_List struct{ capnp.List }

// NewPort_getName_Params creates a new list of Port_getName_Params.
func NewPort_getName_Params_List(s *capnp.Segment, sz int32) (Port_getName_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	if err != nil {
		return Port_getName_Params_List{}, err
	}
	return Port_getName_Params_List{l}, nil
}

func (s Port_getName_Params_List) At(i int) Port_getName_Params {
	return Port_getName_Params{s.List.Struct(i)}
}
func (s Port_getName_Params_List) Set(i int, v Port_getName_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_getName_Params_Promise is a wrapper for a Port_getName_Params promised by a client call.
type Port_getName_Params_Promise struct{ *capnp.Pipeline }

func (p Port_getName_Params_Promise) Struct() (Port_getName_Params, error) {
	s, err := p.Pipeline.Struct()
	return Port_getName_Params{s}, err
}

type Port_getName_Results struct{ capnp.Struct }

func NewPort_getName_Results(s *capnp.Segment) (Port_getName_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_getName_Results{}, err
	}
	return Port_getName_Results{st}, nil
}

func NewRootPort_getName_Results(s *capnp.Segment) (Port_getName_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	if err != nil {
		return Port_getName_Results{}, err
	}
	return Port_getName_Results{st}, nil
}

func ReadRootPort_getName_Results(msg *capnp.Message) (Port_getName_Results, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Port_getName_Results{}, err
	}
	return Port_getName_Results{root.Struct()}, nil
}
func (s Port_getName_Results) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Port_getName_Results) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Port_getName_Results) NameBytes() ([]byte, error) {
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

func (s Port_getName_Results) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

// Port_getName_Results_List is a list of Port_getName_Results.
type Port_getName_Results_List struct{ capnp.List }

// NewPort_getName_Results creates a new list of Port_getName_Results.
func NewPort_getName_Results_List(s *capnp.Segment, sz int32) (Port_getName_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	if err != nil {
		return Port_getName_Results_List{}, err
	}
	return Port_getName_Results_List{l}, nil
}

func (s Port_getName_Results_List) At(i int) Port_getName_Results {
	return Port_getName_Results{s.List.Struct(i)}
}
func (s Port_getName_Results_List) Set(i int, v Port_getName_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// Port_getName_Results_Promise is a wrapper for a Port_getName_Results promised by a client call.
type Port_getName_Results_Promise struct{ *capnp.Pipeline }

func (p Port_getName_Results_Promise) Struct() (Port_getName_Results, error) {
	s, err := p.Pipeline.Struct()
	return Port_getName_Results{s}, err
}

type Stream struct{ capnp.Struct }

func NewStream(s *capnp.Segment) (Stream, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	if err != nil {
		return Stream{}, err
	}
	return Stream{st}, nil
}

func NewRootStream(s *capnp.Segment) (Stream, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
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
func (s Stream) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return "", err
	}
	return p.Text(), nil
}

func (s Stream) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Stream) NameBytes() ([]byte, error) {
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

func (s Stream) SetName(v string) error {
	t, err := capnp.NewText(s.Struct.Segment(), v)
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, t.List.ToPtr())
}

func (s Stream) Loop() bool {
	return s.Struct.Bit(0)
}

func (s Stream) SetLoop(v bool) {
	s.Struct.SetBit(0, v)
}

func (s Stream) Repeat() int8 {
	return int8(s.Struct.Uint8(1))
}

func (s Stream) SetRepeat(v int8) {
	s.Struct.SetUint8(1, uint8(v))
}

func (s Stream) PacketsPerSec() int32 {
	return int32(s.Struct.Uint32(4))
}

func (s Stream) SetPacketsPerSec(v int32) {
	s.Struct.SetUint32(4, uint32(v))
}

func (s Stream) Payload() (Payload, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Payload{}, err
	}
	return Payload{Struct: p.Struct()}, nil
}

func (s Stream) HasPayload() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Stream) SetPayload(v Payload) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewPayload sets the payload field to a newly
// allocated Payload struct, preferring placement in s's segment.
func (s Stream) NewPayload() (Payload, error) {
	ss, err := NewPayload(s.Struct.Segment())
	if err != nil {
		return Payload{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Stream) Layers() (Protocol_List, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return Protocol_List{}, err
	}
	return Protocol_List{List: p.List()}, nil
}

func (s Stream) HasLayers() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Stream) SetLayers(v Protocol_List) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewLayers sets the layers field to a newly
// allocated Protocol_List, preferring placement in s's segment.
func (s Stream) NewLayers(n int32) (Protocol_List, error) {
	l, err := NewProtocol_List(s.Struct.Segment(), n)
	if err != nil {
		return Protocol_List{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

// Stream_List is a list of Stream.
type Stream_List struct{ capnp.List }

// NewStream creates a new list of Stream.
func NewStream_List(s *capnp.Segment, sz int32) (Stream_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
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

func (p Stream_Promise) Payload() Payload_Promise {
	return Payload_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

type Field struct{ capnp.Struct }

func NewField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
	if err != nil {
		return Field{}, err
	}
	return Field{st}, nil
}

func NewRootField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0})
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
func (s Field) Mask() uint64 {
	return s.Struct.Uint64(0)
}

func (s Field) SetMask(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Field) Value() uint64 {
	return s.Struct.Uint64(8)
}

func (s Field) SetValue(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s Field) Step() uint64 {
	return s.Struct.Uint64(16)
}

func (s Field) SetStep(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s Field) Mode() Field_Mode {
	return Field_Mode(s.Struct.Uint16(24))
}

func (s Field) SetMode(v Field_Mode) {
	s.Struct.SetUint16(24, uint16(v))
}

// Field_List is a list of Field.
type Field_List struct{ capnp.List }

// NewField creates a new list of Field.
func NewField_List(s *capnp.Segment, sz int32) (Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 0}, sz)
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

type Field_Mode uint16

// Values of Field_Mode.
const (
	Field_Mode_increment Field_Mode = 0
	Field_Mode_decrement Field_Mode = 1
	Field_Mode_random    Field_Mode = 2
)

// String returns the enum's constant name.
func (c Field_Mode) String() string {
	switch c {
	case Field_Mode_increment:
		return "increment"
	case Field_Mode_decrement:
		return "decrement"
	case Field_Mode_random:
		return "random"

	default:
		return ""
	}
}

// Field_ModeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func Field_ModeFromString(c string) Field_Mode {
	switch c {
	case "increment":
		return Field_Mode_increment
	case "decrement":
		return Field_Mode_decrement
	case "random":
		return Field_Mode_random

	default:
		return 0
	}
}

type Field_Mode_List struct{ capnp.List }

func NewField_Mode_List(s *capnp.Segment, sz int32) (Field_Mode_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	if err != nil {
		return Field_Mode_List{}, err
	}
	return Field_Mode_List{l.List}, nil
}

func (l Field_Mode_List) At(i int) Field_Mode {
	ul := capnp.UInt16List{List: l.List}
	return Field_Mode(ul.At(i))
}

func (l Field_Mode_List) Set(i int, v Field_Mode) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type Payload struct{ capnp.Struct }

func NewPayload(s *capnp.Segment) (Payload, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Payload{}, err
	}
	return Payload{st}, nil
}

func NewRootPayload(s *capnp.Segment) (Payload, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	if err != nil {
		return Payload{}, err
	}
	return Payload{st}, nil
}

func ReadRootPayload(msg *capnp.Message) (Payload, error) {
	root, err := msg.RootPtr()
	if err != nil {
		return Payload{}, err
	}
	return Payload{root.Struct()}, nil
}
func (s Payload) Data() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return nil, err
	}
	return []byte(p.Data()), nil
}

func (s Payload) HasData() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Payload) SetData(v []byte) error {
	d, err := capnp.NewData(s.Struct.Segment(), []byte(v))
	if err != nil {
		return err
	}
	return s.Struct.SetPtr(0, d.List.ToPtr())
}

func (s Payload) Randomize() bool {
	return s.Struct.Bit(0)
}

func (s Payload) SetRandomize(v bool) {
	s.Struct.SetBit(0, v)
}

func (s Payload) Length() uint32 {
	return s.Struct.Uint32(4)
}

func (s Payload) SetLength(v uint32) {
	s.Struct.SetUint32(4, v)
}

// Payload_List is a list of Payload.
type Payload_List struct{ capnp.List }

// NewPayload creates a new list of Payload.
func NewPayload_List(s *capnp.Segment, sz int32) (Payload_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	if err != nil {
		return Payload_List{}, err
	}
	return Payload_List{l}, nil
}

func (s Payload_List) At(i int) Payload           { return Payload{s.List.Struct(i)} }
func (s Payload_List) Set(i int, v Payload) error { return s.List.SetStruct(i, v.Struct) }

// Payload_Promise is a wrapper for a Payload promised by a client call.
type Payload_Promise struct{ *capnp.Pipeline }

func (p Payload_Promise) Struct() (Payload, error) {
	s, err := p.Pipeline.Struct()
	return Payload{s}, err
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
func (s Protocol_ethernet) Source() (Field, error) {
	p, err := s.Struct.Ptr(0)
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: p.Struct()}, nil
}

func (s Protocol_ethernet) HasSource() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet) SetSource(v Field) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewSource sets the source field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ethernet) NewSource() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ethernet) Destination() (Field, error) {
	p, err := s.Struct.Ptr(1)
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: p.Struct()}, nil
}

func (s Protocol_ethernet) HasDestination() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet) SetDestination(v Field) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDestination sets the destination field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ethernet) NewDestination() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ethernet) EthernetType() (Field, error) {
	p, err := s.Struct.Ptr(2)
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: p.Struct()}, nil
}

func (s Protocol_ethernet) HasEthernetType() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet) SetEthernetType(v Field) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewEthernetType sets the ethernetType field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ethernet) NewEthernetType() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s Protocol_ethernet) Length() (Field, error) {
	p, err := s.Struct.Ptr(3)
	if err != nil {
		return Field{}, err
	}
	return Field{Struct: p.Struct()}, nil
}

func (s Protocol_ethernet) HasLength() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Protocol_ethernet) SetLength(v Field) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewLength sets the length field to a newly
// allocated Field struct, preferring placement in s's segment.
func (s Protocol_ethernet) NewLength() (Field, error) {
	ss, err := NewField(s.Struct.Segment())
	if err != nil {
		return Field{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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
	return Field{Struct: p.Struct()}, nil
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

func (p Protocol_Promise) Ethernet() Protocol_ethernet_Promise {
	return Protocol_ethernet_Promise{p.Pipeline}
}

// Protocol_ethernet_Promise is a wrapper for a Protocol_ethernet promised by a client call.
type Protocol_ethernet_Promise struct{ *capnp.Pipeline }

func (p Protocol_ethernet_Promise) Struct() (Protocol_ethernet, error) {
	s, err := p.Pipeline.Struct()
	return Protocol_ethernet{s}, err
}

func (p Protocol_ethernet_Promise) Source() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Protocol_ethernet_Promise) Destination() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Protocol_ethernet_Promise) EthernetType() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Protocol_ethernet_Promise) Length() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Protocol_Promise) Ipv4() Protocol_ipv4_Promise { return Protocol_ipv4_Promise{p.Pipeline} }

// Protocol_ipv4_Promise is a wrapper for a Protocol_ipv4 promised by a client call.
type Protocol_ipv4_Promise struct{ *capnp.Pipeline }

func (p Protocol_ipv4_Promise) Struct() (Protocol_ipv4, error) {
	s, err := p.Pipeline.Struct()
	return Protocol_ipv4{s}, err
}

func (p Protocol_ipv4_Promise) Version() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p Protocol_ipv4_Promise) Ihl() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Protocol_ipv4_Promise) Tos() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p Protocol_ipv4_Promise) Length() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Protocol_ipv4_Promise) Id() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(4)}
}

func (p Protocol_ipv4_Promise) Flags() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(5)}
}

func (p Protocol_ipv4_Promise) FragOffset() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(6)}
}

func (p Protocol_ipv4_Promise) Ttl() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(7)}
}

func (p Protocol_ipv4_Promise) Protocol() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(8)}
}

func (p Protocol_ipv4_Promise) Checksum() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(9)}
}

func (p Protocol_ipv4_Promise) Srcip() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(10)}
}

func (p Protocol_ipv4_Promise) Dstip() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(11)}
}

func (p Protocol_ipv4_Promise) Options() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(12)}
}

func (p Protocol_ipv4_Promise) Padding() Field_Promise {
	return Field_Promise{Pipeline: p.Pipeline.GetPipeline(13)}
}
