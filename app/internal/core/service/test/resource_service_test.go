package test

import (
	"github.com/stefnef/Flowingo/m/internal/core/domain"
	"github.com/stefnef/Flowingo/m/internal/core/service"
	"testing"
)

var resourceService = service.NewResourceService()

func TestGetResources(t *testing.T) {
	var expectedResources = []domain.Resource{
		{
			Id:          "some-id",
			Name:        "Some Name",
			MagicNumber: 41,
		},
	}
	var info = resourceService.GetResources()

	if info == nil {
		t.Fatal("info is nil")
	}

	if len(info) != len(expectedResources) {
		t.Fatalf("wrong size: '%d'", len(info))
	}

	if info[0] != expectedResources[0] {
		t.Fatalf("wrong resources:\nact: '%+v'\nexp: '%+v'", info, expectedResources)
	}
}
