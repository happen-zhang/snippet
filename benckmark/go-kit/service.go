package kit

import (
	"context"
	"strings"
)

type Service interface {
	Say(context.Context, string, int) string
}

type service struct{}

func (svc *service) Say(_ context.Context, word string, repeat int) string {
	if repeat > 1 && word != "" {
		word = strings.Repeat(word, int(repeat))
	}
	return "hello" + word + "!"
}

func NewService() Service {
	return &service{}
}
