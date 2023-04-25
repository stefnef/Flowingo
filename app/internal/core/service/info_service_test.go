package service

import (
	"testing"
)

var infoService = &InfoService{}

func TestGetInfo(t *testing.T) {

	var info = infoService.getInfo()

	if info == nil {
		t.Fatal("info is nil")
	}

	if info.Text != "Example Resource Server" {
		t.Fatalf("wrong info value: '%s'", info.Text)
	}
}
