package service

import (
	"testing"
)

var infoService = &InfoServiceImpl{}

func TestGetInfo(t *testing.T) {

	var info = infoService.GetInfo()

	if info == nil {
		t.Fatal("info is nil")
	}

	if info.Text != "Example Resource Server" {
		t.Fatalf("wrong info value: '%s'", info.Text)
	}
}
