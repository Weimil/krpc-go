// Package remotetech provides methods to invoke procedures in the RemoteTech
// service.
//
// From service docs: this service provides functionality to interact with <a
// href="https://forum.kerbalspaceprogram.com/index.php?/topic/139167-13-remotetech-v188-2017-09-03/">RemoteTech</a>.
package remotetech

import (
	krpcgo "github.com/weimil/krpc-go"
	krpc "github.com/weimil/krpc-go/krpc"
	encode "github.com/weimil/krpc-go/lib/encode"
	service "github.com/weimil/krpc-go/lib/service"
	spacecenter "github.com/weimil/krpc-go/spacecenter"
	types "github.com/weimil/krpc-go/types"
	tracerr "github.com/ztrue/tracerr"
)

// Code generated by gen_services.go. DO NOT EDIT.

/*
Target - the type of object an antenna is targetting. See <see
cref="M:RemoteTech.Antenna.Target" />.
*/
type Target int32

const (
	// The active vessel.
	Target_ActiveVessel Target = 0
	// A celestial body.
	Target_CelestialBody Target = 1
	// A ground station.
	Target_GroundStation Target = 2
	// A specific vessel.
	Target_Vessel Target = 3
	// No target.
	Target_None Target = 4
)

func (v Target) Value() int32 {
	return int32(v)
}
func (v *Target) SetValue(val int32) {
	*v = Target(val)
}

// Antenna - a RemoteTech antenna. Obtained by calling <see
// cref="M:RemoteTech.Comms.Antennas" /> or <see cref="M:RemoteTech.Antenna" />.
type Antenna struct {
	service.BaseClass
}

// NewAntenna creates a new Antenna.
func NewAntenna(id uint64, client *krpcgo.KRPCClient) *Antenna {
	c := &Antenna{BaseClass: service.BaseClass{Client: client}}
	c.SetID(id)
	return c
}

// Comms - communications for a vessel.
type Comms struct {
	service.BaseClass
}

// NewComms creates a new Comms.
func NewComms(id uint64, client *krpcgo.KRPCClient) *Comms {
	c := &Comms{BaseClass: service.BaseClass{Client: client}}
	c.SetID(id)
	return c
}

// RemoteTech - this service provides functionality to interact with <a
// href="https://forum.kerbalspaceprogram.com/index.php?/topic/139167-13-remotetech-v188-2017-09-03/">RemoteTech</a>.
type RemoteTech struct {
	Client *krpcgo.KRPCClient
}

// New creates a new RemoteTech.
func New(client *krpcgo.KRPCClient) *RemoteTech {
	return &RemoteTech{Client: client}
}

// Comms - get a communications object, representing the communication
// capability of a particular vessel.
//
// Allowed game scenes: any.
func (s *RemoteTech) Comms(vessel *spacecenter.Vessel) (*Comms, error) {
	var err error
	var argBytes []byte
	var vv Comms
	request := &types.ProcedureCall{
		Procedure: "Comms",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(vessel)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	vv.Client = s.Client
	return &vv, nil
}

// Antenna - get the antenna object for a particular part.
//
// Allowed game scenes: any.
func (s *RemoteTech) Antenna(part *spacecenter.Part) (*Antenna, error) {
	var err error
	var argBytes []byte
	var vv Antenna
	request := &types.ProcedureCall{
		Procedure: "Antenna",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(part)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	vv.Client = s.Client
	return &vv, nil
}

// Available - whether RemoteTech is installed.
//
// Allowed game scenes: any.
func (s *RemoteTech) Available() (bool, error) {
	var err error
	var vv bool
	request := &types.ProcedureCall{
		Procedure: "get_Available",
		Service:   "RemoteTech",
	}
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// AvailableStream - whether RemoteTech is installed.
//
// Allowed game scenes: any.
func (s *RemoteTech) AvailableStream() (*krpcgo.Stream[bool], error) {
	var err error
	request := &types.ProcedureCall{
		Procedure: "get_Available",
		Service:   "RemoteTech",
	}
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) bool {
		var value bool
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// GroundStations - the names of the ground stations.
//
// Allowed game scenes: any.
func (s *RemoteTech) GroundStations() ([]string, error) {
	var err error
	var vv []string
	request := &types.ProcedureCall{
		Procedure: "get_GroundStations",
		Service:   "RemoteTech",
	}
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// GroundStationsStream - the names of the ground stations.
//
// Allowed game scenes: any.
func (s *RemoteTech) GroundStationsStream() (*krpcgo.Stream[[]string], error) {
	var err error
	request := &types.ProcedureCall{
		Procedure: "get_GroundStations",
		Service:   "RemoteTech",
	}
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) []string {
		var value []string
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// Part - get the part containing this antenna.
//
// Allowed game scenes: any.
func (s *Antenna) Part() (*spacecenter.Part, error) {
	var err error
	var argBytes []byte
	var vv spacecenter.Part
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_Part",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	vv.Client = s.Client
	return &vv, nil
}

// HasConnection - whether the antenna has a connection.
//
// Allowed game scenes: any.
func (s *Antenna) HasConnection() (bool, error) {
	var err error
	var argBytes []byte
	var vv bool
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_HasConnection",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// HasConnectionStream - whether the antenna has a connection.
//
// Allowed game scenes: any.
func (s *Antenna) HasConnectionStream() (*krpcgo.Stream[bool], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_HasConnection",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) bool {
		var value bool
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// Target - the object that the antenna is targetting. This property can be used
// to set the target to <see cref="M:RemoteTech.Target.None" /> or <see
// cref="M:RemoteTech.Target.ActiveVessel" />. To set the target to a celestial
// body, ground station or vessel see <see
// cref="M:RemoteTech.Antenna.TargetBody" />, <see
// cref="M:RemoteTech.Antenna.TargetGroundStation" /> and <see
// cref="M:RemoteTech.Antenna.TargetVessel" />.
//
// Allowed game scenes: any.
func (s *Antenna) Target() (Target, error) {
	var err error
	var argBytes []byte
	var vv Target
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_Target",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// TargetStream - the object that the antenna is targetting. This property can
// be used to set the target to <see cref="M:RemoteTech.TargetStream.None" /> or
// <see cref="M:RemoteTech.TargetStream.ActiveVessel" />. To set the target to a
// celestial body, ground station or vessel see <see
// cref="M:RemoteTech.Antenna.TargetStreamBody" />, <see
// cref="M:RemoteTech.Antenna.TargetStreamGroundStation" /> and <see
// cref="M:RemoteTech.Antenna.TargetStreamVessel" />.
//
// Allowed game scenes: any.
func (s *Antenna) TargetStream() (*krpcgo.Stream[Target], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_Target",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) Target {
		var value Target
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// SetTarget - the object that the antenna is targetting. This property can be
// used to set the target to <see cref="M:RemoteTech.Target.None" /> or <see
// cref="M:RemoteTech.Target.ActiveVessel" />. To set the target to a celestial
// body, ground station or vessel see <see
// cref="M:RemoteTech.Antenna.TargetBody" />, <see
// cref="M:RemoteTech.Antenna.TargetGroundStation" /> and <see
// cref="M:RemoteTech.Antenna.TargetVessel" />.
//
// Allowed game scenes: any.
func (s *Antenna) SetTarget(value Target) error {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Antenna_set_Target",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	argBytes, err = encode.Marshal(value)
	if err != nil {
		return tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x1),
		Value:    argBytes,
	})
	_, err = s.Client.Call(request)
	if err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}

// TargetBody - the celestial body the antenna is targetting.
//
// Allowed game scenes: any.
func (s *Antenna) TargetBody() (*spacecenter.CelestialBody, error) {
	var err error
	var argBytes []byte
	var vv spacecenter.CelestialBody
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_TargetBody",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	vv.Client = s.Client
	return &vv, nil
}

// SetTargetBody - the celestial body the antenna is targetting.
//
// Allowed game scenes: any.
func (s *Antenna) SetTargetBody(value *spacecenter.CelestialBody) error {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Antenna_set_TargetBody",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	argBytes, err = encode.Marshal(value)
	if err != nil {
		return tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x1),
		Value:    argBytes,
	})
	_, err = s.Client.Call(request)
	if err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}

// TargetGroundStation - the ground station the antenna is targetting.
//
// Allowed game scenes: any.
func (s *Antenna) TargetGroundStation() (string, error) {
	var err error
	var argBytes []byte
	var vv string
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_TargetGroundStation",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// TargetGroundStationStream - the ground station the antenna is targetting.
//
// Allowed game scenes: any.
func (s *Antenna) TargetGroundStationStream() (*krpcgo.Stream[string], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_TargetGroundStation",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) string {
		var value string
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// SetTargetGroundStation - the ground station the antenna is targetting.
//
// Allowed game scenes: any.
func (s *Antenna) SetTargetGroundStation(value string) error {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Antenna_set_TargetGroundStation",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	argBytes, err = encode.Marshal(value)
	if err != nil {
		return tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x1),
		Value:    argBytes,
	})
	_, err = s.Client.Call(request)
	if err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}

// TargetVessel - the vessel the antenna is targetting.
//
// Allowed game scenes: any.
func (s *Antenna) TargetVessel() (*spacecenter.Vessel, error) {
	var err error
	var argBytes []byte
	var vv spacecenter.Vessel
	request := &types.ProcedureCall{
		Procedure: "Antenna_get_TargetVessel",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	vv.Client = s.Client
	return &vv, nil
}

// SetTargetVessel - the vessel the antenna is targetting.
//
// Allowed game scenes: any.
func (s *Antenna) SetTargetVessel(value *spacecenter.Vessel) error {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Antenna_set_TargetVessel",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	argBytes, err = encode.Marshal(value)
	if err != nil {
		return tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x1),
		Value:    argBytes,
	})
	_, err = s.Client.Call(request)
	if err != nil {
		return tracerr.Wrap(err)
	}
	return nil
}

// SignalDelayToVessel - the signal delay between the this vessel and another
// vessel, in seconds.
//
// Allowed game scenes: any.
func (s *Comms) SignalDelayToVessel(other *spacecenter.Vessel) (float64, error) {
	var err error
	var argBytes []byte
	var vv float64
	request := &types.ProcedureCall{
		Procedure: "Comms_SignalDelayToVessel",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	argBytes, err = encode.Marshal(other)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x1),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// SignalDelayToVesselStream - the signal delay between the this vessel and
// another vessel, in seconds.
//
// Allowed game scenes: any.
func (s *Comms) SignalDelayToVesselStream(other *spacecenter.Vessel) (*krpcgo.Stream[float64], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Comms_SignalDelayToVessel",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	argBytes, err = encode.Marshal(other)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x1),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) float64 {
		var value float64
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// Vessel - get the vessel.
//
// Allowed game scenes: any.
func (s *Comms) Vessel() (*spacecenter.Vessel, error) {
	var err error
	var argBytes []byte
	var vv spacecenter.Vessel
	request := &types.ProcedureCall{
		Procedure: "Comms_get_Vessel",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return &vv, tracerr.Wrap(err)
	}
	vv.Client = s.Client
	return &vv, nil
}

// HasLocalControl - whether the vessel can be controlled locally.
//
// Allowed game scenes: any.
func (s *Comms) HasLocalControl() (bool, error) {
	var err error
	var argBytes []byte
	var vv bool
	request := &types.ProcedureCall{
		Procedure: "Comms_get_HasLocalControl",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// HasLocalControlStream - whether the vessel can be controlled locally.
//
// Allowed game scenes: any.
func (s *Comms) HasLocalControlStream() (*krpcgo.Stream[bool], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Comms_get_HasLocalControl",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) bool {
		var value bool
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// HasFlightComputer - whether the vessel has a flight computer on board.
//
// Allowed game scenes: any.
func (s *Comms) HasFlightComputer() (bool, error) {
	var err error
	var argBytes []byte
	var vv bool
	request := &types.ProcedureCall{
		Procedure: "Comms_get_HasFlightComputer",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// HasFlightComputerStream - whether the vessel has a flight computer on board.
//
// Allowed game scenes: any.
func (s *Comms) HasFlightComputerStream() (*krpcgo.Stream[bool], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Comms_get_HasFlightComputer",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) bool {
		var value bool
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// HasConnection - whether the vessel has any connection.
//
// Allowed game scenes: any.
func (s *Comms) HasConnection() (bool, error) {
	var err error
	var argBytes []byte
	var vv bool
	request := &types.ProcedureCall{
		Procedure: "Comms_get_HasConnection",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// HasConnectionStream - whether the vessel has any connection.
//
// Allowed game scenes: any.
func (s *Comms) HasConnectionStream() (*krpcgo.Stream[bool], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Comms_get_HasConnection",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) bool {
		var value bool
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// HasConnectionToGroundStation - whether the vessel has a connection to a
// ground station.
//
// Allowed game scenes: any.
func (s *Comms) HasConnectionToGroundStation() (bool, error) {
	var err error
	var argBytes []byte
	var vv bool
	request := &types.ProcedureCall{
		Procedure: "Comms_get_HasConnectionToGroundStation",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// HasConnectionToGroundStationStream - whether the vessel has a connection to a
// ground station.
//
// Allowed game scenes: any.
func (s *Comms) HasConnectionToGroundStationStream() (*krpcgo.Stream[bool], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Comms_get_HasConnectionToGroundStation",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) bool {
		var value bool
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// SignalDelay - the shortest signal delay to the vessel, in seconds.
//
// Allowed game scenes: any.
func (s *Comms) SignalDelay() (float64, error) {
	var err error
	var argBytes []byte
	var vv float64
	request := &types.ProcedureCall{
		Procedure: "Comms_get_SignalDelay",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// SignalDelayStream - the shortest signal delay to the vessel, in seconds.
//
// Allowed game scenes: any.
func (s *Comms) SignalDelayStream() (*krpcgo.Stream[float64], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Comms_get_SignalDelay",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) float64 {
		var value float64
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// SignalDelayToGroundStation - the signal delay between the vessel and the
// closest ground station, in seconds.
//
// Allowed game scenes: any.
func (s *Comms) SignalDelayToGroundStation() (float64, error) {
	var err error
	var argBytes []byte
	var vv float64
	request := &types.ProcedureCall{
		Procedure: "Comms_get_SignalDelayToGroundStation",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// SignalDelayToGroundStationStream - the signal delay between the vessel and
// the closest ground station, in seconds.
//
// Allowed game scenes: any.
func (s *Comms) SignalDelayToGroundStationStream() (*krpcgo.Stream[float64], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Comms_get_SignalDelayToGroundStation",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) float64 {
		var value float64
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}

// Antennas - the antennas for this vessel.
//
// Allowed game scenes: any.
func (s *Comms) Antennas() ([]*Antenna, error) {
	var err error
	var argBytes []byte
	var vv []*Antenna
	request := &types.ProcedureCall{
		Procedure: "Comms_get_Antennas",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	result, err := s.Client.Call(request)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	err = encode.Unmarshal(result.Value, &vv)
	if err != nil {
		return vv, tracerr.Wrap(err)
	}
	return vv, nil
}

// AntennasStream - the antennas for this vessel.
//
// Allowed game scenes: any.
func (s *Comms) AntennasStream() (*krpcgo.Stream[[]*Antenna], error) {
	var err error
	var argBytes []byte
	request := &types.ProcedureCall{
		Procedure: "Comms_get_Antennas",
		Service:   "RemoteTech",
	}
	argBytes, err = encode.Marshal(s)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	request.Arguments = append(request.Arguments, &types.Argument{
		Position: uint32(0x0),
		Value:    argBytes,
	})
	krpc := krpc.New(s.Client)
	st, err := krpc.AddStream(request, true)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	rawStream := s.Client.GetStream(st.Id)
	stream := krpcgo.MapStream(rawStream, func(b []byte) []*Antenna {
		var value []*Antenna
		encode.Unmarshal(b, &value)
		return value
	})
	stream.AddCloser(func() error {
		return tracerr.Wrap(krpc.RemoveStream(st.Id))
	})
	return stream, nil
}
