package servers

type msgType byte

const DATA msgType = 0
const CLOSE msgType = 1

type message struct {
	Addr string
	Type msgType
	Data []byte
}

func newCloseMessage(addr string) *message {
	return &message{Addr: addr, Type: CLOSE, Data: []byte{}}
}
func newDataMessage(addr string, data []byte) *message {
	return &message{Addr: addr, Type: DATA, Data: data}
}
