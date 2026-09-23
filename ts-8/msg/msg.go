package msg

import (
	"fmt"
	"math/rand"
)

const maxMessage = 10

type Message struct {
	Id  int
	Msg string
}

var couter = 0

func GetMessages() []Message {

	msgCoupasity := rand.Intn(maxMessage)
	messages := make([]Message, 0, msgCoupasity)

	for i := 0; i < msgCoupasity; i++ {
		couter++
		messages = append(messages, Message{
			Id:  couter,
			Msg: fmt.Sprintf("message %d", couter),
		})
	}
	return messages
}
