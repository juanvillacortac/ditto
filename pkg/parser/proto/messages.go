package proto

import (
	protoparser "github.com/yoheimuta/go-protoparser/v4/parser"
)

func GetMessages(nodes []protoparser.Visitee) []*protoparser.Message {
	messages := make([]*protoparser.Message, 0)
	for _, n := range nodes {
		if n, ok := n.(*protoparser.Message); ok {
			messages = append(messages, n)
		}
	}
	return messages
}

func GetOptions(nodes []protoparser.Visitee) []*protoparser.Option {
	options := make([]*protoparser.Option, 0)
	for _, n := range nodes {
		if n, ok := n.(*protoparser.Option); ok {
			options = append(options, n)
		}
	}
	return options
}
