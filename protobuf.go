package log

import (
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func protoBuf(key string, msg proto.Message) Field {
	return zap.Reflect(key, &protoMessage{msg: msg})
}

type protoMessage struct {
	msg proto.Message
}

func (p *protoMessage) MarshalJSON() ([]byte, error) {
	if p.msg == nil {
		return make([]byte, 0), nil
	}

	m := &protojson.MarshalOptions{}
	buf, err := m.Marshal(p.msg)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
