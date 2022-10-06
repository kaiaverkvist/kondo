package kondo

import "reflect"
import "github.com/kelindar/binary"

type netMessage struct {
	Ref  string
	Data []byte
}

func newNetMessage(data any) netMessage {
	encodedData, _ := binary.Marshal(data)
	return netMessage{
		Ref:  reflect.TypeOf(data).String(),
		Data: encodedData,
	}
}

func parseIncoming(data []byte) (netMessage, error) {
	var nm netMessage

	err := binary.Unmarshal(data, &nm)
	return nm, err
}

func encode(data netMessage) ([]byte, error) {
	return binary.Marshal(data)
}

func message(data any) []byte {
	nm := newNetMessage(data)
	b, _ := encode(nm)

	return b
}
