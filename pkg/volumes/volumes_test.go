package volumes

import (
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)

func TestCreateVolumes(t *testing.T) {
	volumes := []api.Volume{
		{
			Name: "host-dir",
			Type: "HOST",
			Path: "/dir/path",
		},
	}
	expectedPaths := []string{"/dir/path"}
	for i, volume := range volumes {
		extVolume, _ := CreateVolume(&volume)
		expectedPath := expectedPaths[i]
		path := extVolume.GetPath()
		if expectedPath != path {
			t.Errorf("Unexpected bind path. Expected %v, got %v", expectedPath, path)
		}
	}
}
