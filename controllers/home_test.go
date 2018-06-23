package controllers

import (
	"testing"
)

func TestNewHomeController(t *testing.T) {
	_, c := NewHomeController()
	f := c.GetMethod("index")
	if f == nil {
		t.Error("No index action!")
	}
}
