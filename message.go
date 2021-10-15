package dtmgrpc

import (
	"context"

	"github.com/yedf/dtmcli"
)

// MsgGrpc reliable msg type
type MsgGrpc struct {
	dtmcli.TransBase
	Steps         []dtmcli.MsgStep `json:"steps"`
	QueryPrepared string           `json:"query_prepared"`
}

// NewMsgGrpc create new msg
func NewMsgGrpc(server string, gid string) *MsgGrpc {
	return &MsgGrpc{TransBase: *dtmcli.NewTransBase(gid, "msg", server, "")}
}

// Add add a new step
func (s *MsgGrpc) Add(action string, data []byte) *MsgGrpc {
	s.Steps = append(s.Steps, dtmcli.MsgStep{
		Action: action,
		Data:   string(data),
	})
	return s
}

// Submit submit the msg
func (s *MsgGrpc) Submit() error {
	_, err := MustGetDtmClient(s.Dtm).Submit(context.Background(), &DtmRequest{
		Gid:       s.Gid,
		TransType: s.TransType,
		Data:      dtmcli.MustMarshalString(&s.Steps),
	})
	return err
}

// Prepare prepare the msg
func (s *MsgGrpc) Prepare(queryPrepared string) error {
	s.QueryPrepared = dtmcli.OrString(queryPrepared, s.QueryPrepared)
	_, err := MustGetDtmClient(s.Dtm).Prepare(context.Background(), &DtmRequest{
		Gid:           s.Gid,
		TransType:     s.TransType,
		QueryPrepared: s.QueryPrepared,
		Data:          dtmcli.MustMarshalString(&s.Steps),
	})
	return err
}
