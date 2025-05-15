package app

import (
	"github.com/crispgm/foosbot/internal/def"
	"github.com/go-lark/lark"
)

func newBot() *lark.Bot {
	return lark.NewChatBot(def.AppID, def.AppSecret)
}
