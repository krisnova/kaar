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
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"

	"github.com/kris-nova/logger"
	"github.com/mitchellh/go-homedir"
)

const (

	//	Reference: https://yaml.org/spec/1.2/spec.html
	//
	YAMLDelimiter string = "\n---\n"
)

var Version string

// Create will attempt to create a kaarball from a directory
func Create(dir string, path string) (*Archive, error) {
	logger.Info("CREATE file: %s from directory: %s", path, dir)
	// Glob directory
	archive := NewArchive(dir)
	err := archive.Load(dir)
	if err != nil {
		return nil, err
	}
	err = archive.WriteArchive(path)
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

	// Reference path for the archive on disk
	Path string

	// Manifests are valid Kubernetes manifests
	Manifests []*Manifest

	// Files is every file in the directory
	Files map[string]*fs.FileInfo
}

// NewArchive will create a new archive from a registered path
// By convention the path should be the directory that may or may not
// exist for a specific archive.
//
// path: directory to consider
func NewArchive(path string) *Archive {
	return &Archive{
		Path:  path,
		Files: make(map[string]*fs.FileInfo),
	}
}

// WriteArchive will write an archive to a given path
func (a *Archive) WriteArchive(path string) error {
	return nil
}

// Extract will restore an archive from memory
//
// dir: the name of the directory to write
func (a *Archive) Extract(dir string) error {
	return nil
}

// Load will recursively load the contents of a directory into memory
func (a *Archive) Load(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("unable to read dir in filesystem: %v", err)
	}
	// Recursively look for Kubernetes manifests
	for _, file := range files {
		filename := filepath.Join(dir, file.Name())
		// Look for Kubernetes YAML
		if file.IsDir() {
			a.Load(filename)
			continue
		}
		// Add files no matter what!
		a.Files[filename] = &file
		err := a.LoadManifests(filename)
		if err != nil {
			continue
		}
	}
	return nil
}

type Manifest struct {
	// Path is the relative path of the manifest from the root directory
	Path    string
	Images  map[string]*EmbeddedImage
	Decoded runtime.Object
}

// LoadManifests will attempt to parse a Kubernetes manifest
// and sync it's relevant embedded images locally
func (a *Archive) LoadManifests(path string) error {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("unable to read file path: %s: %v", path, err)
	}
	rawStr := string(raw)
	yamlStrs := strings.Split(rawStr, YAMLDelimiter)
	for _, yaml := range yamlStrs {
		manifest := &Manifest{}
		serializer := scheme.Codecs.UniversalDeserializer()
		var decoded runtime.Object
		decoded, _, err = serializer.Decode([]byte(yaml), nil, nil)
		if err != nil {
			logger.Warning("unable to serialize file: %v", err)
			continue
		}
		manifest.Path = path
		manifest.Decoded = decoded

		a.Manifests = append(a.Manifests, manifest)
		// TODO Container Images
		fmt.Printf("** Manifest Found: %s **\n", manifest.Path)
	}
	return nil
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
