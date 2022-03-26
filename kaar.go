// Copyright © 2022 Kris Nóva <kris@nivenly.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//  Kubernetes Application Archive
//
//  ██╗  ██╗ █████╗  █████╗ ██████╗
//  ██║ ██╔╝██╔══██╗██╔══██╗██╔══██╗
//  █████╔╝ ███████║███████║██████╔╝
//  ██╔═██╗ ██╔══██║██╔══██║██╔══██╗
//  ██║  ██╗██║  ██║██║  ██║██║  ██║
//  ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝
//

package kaar

import (
	"fmt"
	"io/ioutil"

	"github.com/kris-nova/logger"
	"github.com/mitchellh/go-homedir"
)

// Create will attempt to create a kaarball from a directory
func Create(dir string, path string) (*Archive, error) {
	logger.Info("CREATE file: %s from directory: %s", path, dir)
	// Glob directory
	archive := NewArchive(dir)
	err := archive.Read(dir)
	if err != nil {
		return nil, err
	}
	err = archive.Write(path)
	if err != nil {
		return nil, fmt.Errorf("unable to write kaarball: %v", err)
	}
	return archive, nil
}

// Extract will attempt to extract a kaarball
func Extract(dir string, path string) (*Archive, error) {
	logger.Info("EXTRACT file: %s", path)
	return nil, nil
}

type Archive struct {
	Path      string
	Manifests map[string]*Manifest
}

func NewArchive(path string) *Archive {
	return &Archive{
		Path:      path,
		Manifests: make(map[string]*Manifest),
	}
}

// Write will write an archive to a given path
func (a *Archive) Write(path string) error {
	return nil
}

// Restore will restore an archive from memory
//     dir the name of the directory to write
func (a *Archive) Restore(dir string) error {
	return nil
}

// Read will read the contents of a directory into memory
func (a *Archive) Read(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("unable to read dir in filesystem: %v", err)
	}
	// Recursively look for Kubernetes manifests
	for _, file := range files {
		// Look for Kubernetes YAML
		if file.IsDir() {
			a.Read(file.Name())
			continue
		}
		manifest, err := NewManifest(file.Name())
		if err != nil {
			continue
		}
		if manifest != nil {
			a.Manifests[manifest.Path] = manifest
		}
	}
	return nil
}

type Manifest struct {
	// Path is the relative path of the manifest from the root directory
	Path   string
	Images map[string]*EmbeddedImage
}

// NewManifest will attempt to parse a Kubernetes manifest
// and sync it's relevant embedded images locally
func NewManifest(path string) (*Manifest, error) {
	fmt.Println(path)
	return nil, nil
}

// resolveDir will handle POSIX parlance
//
// ~
// .
func resolveDir(dir string) string {
	var ret string
	ret, err := homedir.Expand(dir)
	if err != nil {
		return dir
	}
	return ret
}

type EmbeddedImage struct {
}
