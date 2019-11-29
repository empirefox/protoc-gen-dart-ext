package zero

import (
	"time"

	"github.com/araddon/dateparse"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func (m *Duration) Eval() (*duration.Duration, error) {
	r, err := m.eval()
	if err != nil {
		return nil, err
	}
	if r != nil && r.Seconds == 0 && r.Nanos == 0 {
		return nil, nil
	}
	return r, nil
}

func (m *Duration) eval() (*duration.Duration, error) {
	switch is := m.Is.(type) {
	case *Duration_Wkt:
		return is.Wkt, nil
	case *Duration_Parse:
		r, err := time.ParseDuration(is.Parse)
		if err != nil {
			return nil, err
		}
		return ptypes.DurationProto(r), nil
	}
	return nil, nil
}

type EvalTime struct {
	Wkt *timestamp.Timestamp

	Now    bool
	NowAdd time.Duration
}

func (m *Time) Eval() (*EvalTime, error) {
	r, err := m.eval()
	if err != nil {
		return nil, err
	}

	if r != nil {
		if r.Wkt != nil && r.Wkt.Seconds == 0 && r.Wkt.Nanos == 0 {
			r.Wkt = nil
		}
		if r.Wkt == nil && !r.Now && r.NowAdd == 0 {
			return nil, nil
		}
	}

	return r, nil
}

func (m *Time) eval() (*EvalTime, error) {
	switch is := m.Is.(type) {
	case *Time_Wkt:
		return &EvalTime{Wkt: is.Wkt}, nil
	case *Time_Parse:
		t, err := dateparse.ParseAny(is.Parse)
		if err != nil {
			return nil, err
		}
		ts, err := ptypes.TimestampProto(t)
		if err != nil {
			return nil, err
		}
		return &EvalTime{Wkt: ts}, nil
	case *Time_Now:
		return &EvalTime{Now: is.Now}, nil
	case *Time_NowAdd:
		add, err := ptypes.Duration(is.NowAdd)
		if err != nil {
			return nil, err
		}
		return &EvalTime{Now: true, NowAdd: add}, nil
	case *Time_NowAddParse:
		add, err := time.ParseDuration(is.NowAddParse)
		if err != nil {
			return nil, err
		}
		return &EvalTime{Now: true, NowAdd: add}, nil
	}
	return nil, nil
}
