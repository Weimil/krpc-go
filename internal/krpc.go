package internal

import (
	"github.com/golang/protobuf/proto"
	krpcgo "github.com/weimil/krpc-go"
	"github.com/weimil/krpc-go/types"
	"github.com/ztrue/tracerr"
)

// BasicKRPC is a partial implementation of the KRPC service. This should only
// be used to fetch the rest of the services.
type BasicKRPC struct {
	client *krpcgo.KRPCClient
}

func NewBasicKRPC(client *krpcgo.KRPCClient) *BasicKRPC {
	return &BasicKRPC{client: client}
}

func (s *BasicKRPC) GetStatus() (*types.Status, error) {
	request := &types.ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetStatus",
	}
	result, err := s.client.Call(request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	var status types.Status
	if err := proto.Unmarshal(result.Value, &status); err != nil {
		return nil, tracerr.Wrap(err)
	}
	return &status, nil
}

func (s *BasicKRPC) GetServices() (*types.Services, error) {
	request := &types.ProcedureCall{
		Service:   "KRPC",
		Procedure: "GetServices",
	}
	result, err := s.client.Call(request)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	var services types.Services
	if err := proto.Unmarshal(result.Value, &services); err != nil {
		return nil, tracerr.Wrap(err)
	}
	return &services, nil
}
