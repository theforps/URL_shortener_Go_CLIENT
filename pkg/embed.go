package pkg

import _ "embed"

var (
	//go:embed resources/message_content.json
	messageFS []byte
)

func GetMessageContent() []byte {
	return messageFS
}
