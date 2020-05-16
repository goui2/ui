package message

type MessageType string

const (
	MessageType_Success     MessageType = "success"
	MessageType_Error                   = "error"
	MessageType_None                    = "none"
	MessageType_Information             = "information"
	MessageType_Warning                 = "warning"
)

/*
type EventData interface {
	UnWrap() interface{}
}
*/

type Message struct {
	Id          string
	Message     string
	Description string
	Type        MessageType
}

func (m Message) UnWrap() interface{} { return m }
