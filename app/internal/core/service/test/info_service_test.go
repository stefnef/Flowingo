package test

import (
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"testing"
)

var infoService = service.NewInfoService()

func TestGetInfo(t *testing.T) {

	var info = infoService.GetInfo()

	if info == nil {
		t.Fatal("info is nil")
	}

	if info.Text != "Example Resource Server" {
		t.Fatalf("wrong info value: '%s'", info.Text)
	}
}
