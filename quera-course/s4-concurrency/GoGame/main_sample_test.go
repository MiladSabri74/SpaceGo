package main

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestSample_Player_GetName_ReturnsFormattedName(t *testing.T) {
	p := &Player{name: "aLiCe"}
	expected := "Alice"

	if got := p.GetName(); got != expected {
		t.Errorf("متد GetName نام بازیکن را به‌درستی قالب‌بندی نکرد. مقدار مورد انتظار: %q، مقدار دریافت شده: %q", expected, got)
	}
}

func TestSample_Map_FanOutMessages_BroadcastsToOtherPlayers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := &Map{
		id:        1,
		players:   make(map[string]*Player),
		broadcast: make(chan Message, 1),
	}

	p1 := &Player{name: "Alice", messages: make(chan string, 1)}
	p2 := &Player{name: "Bob", messages: make(chan string, 1)}
	m.players["alice"] = p1
	m.players["bob"] = p2

	var wg sync.WaitGroup
	wg.Add(1)
	go m.FanOutMessages(ctx, &wg)

	m.broadcast <- Message{SenderName: "Alice", Text: "Hi Bob!"}

	select {
	case msg := <-p2.messages:
		if msg != "Alice says: Hi Bob!" {
			t.Errorf("پیام ارسال شده به باب نادرست است. مقدار دریافت شده: %s", msg)
		}
	case <-time.After(500 * time.Millisecond):
		t.Error("باب پیام ارسالی آلیس را دریافت نکرد")
	}

	select {
	case <-p1.messages:
		t.Error("آلیس نباید پیام ارسالی توسط خودش را دریافت کند")
	default:
	}
}

func TestSample_Game_ConnectPlayer_HandlesCasingAndRejectsDuplicates(t *testing.T) {
	g, err := NewGame(context.Background(), []int{101})
	if err != nil {
		t.Fatalf("خطا در ایجاد بازی با شناسه معتبر: %v", err)
	}
	defer g.Shutdown()

	err = g.ConnectPlayer("Charlie")
	if err != nil {
		t.Fatalf("خطا در اتصال بازیکن جدید: %v", err)
	}

	_, err = g.GetPlayer("cHaRlIe")
	if err != nil {
		t.Errorf("بازیکن اضافه شده با جست‌وجوی غیرحساس به حروف پیدا نشد")
	}

	err = g.ConnectPlayer("CHARLIE")
	if err == nil {
		t.Error("اتصال مجدد بازیکنی با نام یکسان (با حروف بزرگ/کوچک متفاوت) باید با خطا مواجه شود")
	}
}
