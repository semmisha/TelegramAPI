package http

import (
	"github.com/sirupsen/logrus"
)

type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type Api struct {
	logging *logrus.Logger
	Message
}

func NewApi(logging *logrus.Logger) *Api {
	return &Api{logging: logging}
}
