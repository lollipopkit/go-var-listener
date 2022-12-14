package govarlistener

import (
	"testing"
	"time"
)

var (
	v  Var[int] = *NewVar(1)
	vv int = 1
)

func TestNew(t *testing.T) {
	if v.Get() != 1 {
		t.Error("NewVar failed")
	}
}

func TestCallbackType(t *testing.T) {
	if OnChange|OnGet|OnListen|OnUnlisten|OnError != OnAll {
		t.Error("CallbackType failed")
	}
}

func TestAddCallback(t *testing.T) {
	v.Listen(&Callback{
		Fn: func() {
			vv += 1
		},
		Name: "add1-onboth",
		Type: OnAll,
	})
	if len(v.callbacks) != 1 {
		t.Error("listen failed")
	}
}

func TestCallback(t *testing.T) {
	if vv != 1 {
		t.Error("vv initial value != 1")
	}

	v.Get()
	time.Sleep(time.Millisecond * 100)
	if vv != 2 {
		t.Error("vv get value != 2")
	}

	v.Set(2)
	time.Sleep(time.Millisecond * 100)
	if vv != 3 {
		t.Error("vv set value != 3")
	}
	if v.Get() != 2 {
		t.Error("v get value != 2")
	}
}

func TestUnlisten(t *testing.T) {
	v.Unlisten("add1-onboth")
	if len(v.callbacks) != 0 {
		t.Error("unlisten failed")
	}
}

func TestUnlistenErr(t *testing.T) {
	err := v.Unlisten("add1-onboth")
	if err != ErrThisNoListenName {
		t.Error("unlisten err failed")
	}
}

func TestListenErr(t *testing.T) {
	v.Listen(&Callback{
		Fn: func() {
			vv += 1
		},
		Name: "add1-onboth",
		Type: OnAll,
	})
	err := v.Listen(&Callback{
		Fn: func() {
			vv += 1
		},
		Name: "add1-onboth",
		Type: OnAll,
	})
	if err != ErrSameCallbackName {
		t.Error("listen err failed")
	}
}

func TestIsListened(t *testing.T) {
	if !v.IsListening("add1-onboth") {
		t.Error("TestIsListened failed")
	}
}
