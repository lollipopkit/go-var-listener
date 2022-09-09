package govarlistener

import (
	"testing"
	"time"
)

var (
	v Var[int]
	vv int = 1
)

func TestNew(t *testing.T) {
	v = *NewVar(1)
	if v.Get() != 1 {
		t.Error("NewVar failed")
	}
}

func TestAddCallback(t *testing.T) {
	v.Listen(Callback[int]{
		fn:   func(i int) {
			vv += 1
		},
		name: "add1-onboth",
		typ:  OnBoth,
	})
	if len(v.callbacks.callback) != 1{
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
	if len(v.callbacks.callback) != 0 {
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
	v.Listen(Callback[int]{
		fn:   func(i int) {
			vv += 1
		},
		name: "add1-onboth",
		typ:  OnBoth,
	})
	err := v.Listen(Callback[int]{
		fn:   func(i int) {
			vv += 1
		},
		name: "add1-onboth",
		typ:  OnBoth,
	})
	if err != ErrSameCallbackName {
		t.Error("listen err failed")
	}
}

func TestHaveListen(t *testing.T) {
	if !v.HaveListen("add1-onboth") {
		t.Error("have listen failed")
	}
}

func TestHaveListenErr(t *testing.T) {
	if v.HaveListen("add1-onboth-err") {
		t.Error("have listen err failed")
	}
}

func TestSet(t *testing.T) {
	v.Set(3)
	if v.Get() != 3 {
		t.Error("set failed")
	}
}

func TestGet(t *testing.T) {
	if v.Get() != 3 {
		t.Error("get failed")
	}
}