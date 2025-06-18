package app

import (
	"github.com/go-lark/lark/v2"

	"github.com/crispgm/foosbot/internal/def"
)

func newBot() *lark.Bot {
	return lark.NewChatBot(def.AppID, def.AppSecret)
}
