package parser

import (
	proto "github.com/yoheimuta/go-protoparser/v4/parser"
)

func GetMessages(nodes []proto.Visitee) []*proto.Message {
	messages := make([]*proto.Message, 0)
	for _, n := range nodes {
		if n, ok := n.(*proto.Message); ok {
			messages = append(messages, n)
		}
	}
	return messages
}
