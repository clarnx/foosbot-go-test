package app

import (
	"testing"

	"github.com/go-lark/lark/v2"
)

func TestMessageCard(t *testing.T) {
	bot := newBot()
	card := buildCard("zhangwanlong@dcarlife.com", "zhangwanlong@dcarlife.com")
	resp, err := bot.PostMessage(t.Context(), card)
	t.Log(err, resp)
}

func TestGetChat(t *testing.T) {
	bot := newBot()
	bot.WithUserIDType(lark.UIDOpenID)
	resp, _ := bot.ListChat(t.Context(), "ByCreateTimeAsc", "", 10)
	t.Log(resp.Data.Items)
}

func TestGetMessage(t *testing.T) {
	bot := newBot()
	resp, _ := bot.PostText(t.Context(), "hello", lark.WithEmail("zhangwanlong@dcarlife.com"))
	msgResp, _ := bot.WithUserIDType(lark.UIDOpenID).GetMessage(t.Context(), resp.Data.MessageID)
	t.Log(msgResp)
}
