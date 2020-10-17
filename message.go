package bus

import (
	"encoding/json"

	nsq "github.com/nsqio/go-nsq"
)

// Message carries nsq.Message fields and methods and
// adds extra fields for handling messages internally.
type Message struct {
	*nsq.Message
	ReplyTo string `json:"replyTo"`
	Payload []byte `json:"payload"`
	Topic   string `json:"-"`
}

// NewMessage returns a new bus.Message.
func NewMessage(p []byte, r string) *Message {
	return &Message{Payload: p, ReplyTo: r}
}

// DecodePayload deserializes data (as []byte) and creates a new struct passed by parameter.
func (m *Message) DecodePayload(v interface{}) (err error) {
	return json.Unmarshal(m.Payload, v)
}

func decodeMessage(message *nsq.Message) (m *Message, err error) {
	m = &Message{Message: message}
	err = json.Unmarshal(message.Body, m)
	return
}

func encodeMessage(payload interface{}, replyTo string) ([]byte, error) {
	p, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	message := NewMessage(p, replyTo)
	return json.Marshal(message)
}
