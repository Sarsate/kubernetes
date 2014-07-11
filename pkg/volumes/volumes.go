/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package volumes

import (
	"errors"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
)


// All volume types are expected to implement this interface
type ExternalVolume interface {
	// Mounts the volume to a directory path.
	Mount()
	// Returns the directory path the volume is mounted to.
	GetPath() string
	// Unmounes the volume and removes traces of the Mount procedure.
	UnMount()
}

// Host Directory Volumes represent a bare host directory mount.
type HostDirectoryVolume struct {
	Name string
	Path string
}

// Simple host directory mounts require no setup or cleanup, but still
// need to fulfill the interface definitions.
func (hostVol *HostDirectoryVolume) Mount() {}

func (hostVol *HostDirectoryVolume) UnMount() {}

func (hostVol *HostDirectoryVolume) GetPath() string {
	return hostVol.Path
}

// Interprets API volume as a HostDirectory
func createHostDirectoryVolume(volume *api.Volume) *HostDirectoryVolume {
	return &HostDirectoryVolume{volume.Name, volume.Path}
}

// Interprets parameters passed in the API as an internal structure
// with utility procedures for mounting.
func CreateVolume(volume *api.Volume) (ExternalVolume, error) {
	switch volume.Type {
	case "HOST":
		return createHostDirectoryVolume(volume), nil
	default:
		return nil, errors.New("Unsupported volume type.")
	}
}
