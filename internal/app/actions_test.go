package app

import (
	"testing"

	"github.com/go-lark/lark"
)

func TestMessageCard(t *testing.T) {
	bot := newBot()
	bot.GetTenantAccessTokenInternal(true)
	card := buildCard("zhangwanlong@dcarlife.com", "zhangwanlong@dcarlife.com")
	resp, err := bot.PostMessage(card)
	t.Log(err, resp)
}

func TestGetChat(t *testing.T) {
	bot := newBot()
	bot.GetTenantAccessTokenInternal(true)
	bot.WithUserIDType(lark.UIDOpenID)
	resp, _ := bot.ListChat("ByCreateTimeAsc", "", 10)
	t.Log(resp.Data.Items)
}

func TestGetMessage(t *testing.T) {
	bot := newBot()
	bot.GetTenantAccessTokenInternal(true)
	resp, _ := bot.PostText("hello", lark.WithEmail("zhangwanlong@dcarlife.com"))
	msgResp, _ := bot.WithUserIDType(lark.UIDOpenID).GetMessage(resp.Data.MessageID)
	t.Log(msgResp)
}
