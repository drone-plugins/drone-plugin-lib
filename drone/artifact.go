// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

import (
	"encoding/json"
	"fmt"

	"os"
	"path/filepath"
)

// Type of registry
type RegistryType string

// Registries const
const (
	Docker RegistryType = "Docker"
	ECR    RegistryType = "ECR"
	GCR    RegistryType = "GCR"
	ACR    RegistryType = "ACR"
)

const (
	dockerArtifactV1 string = "docker/v1"
)

type (
	//  Image stores the image data
	Image struct {
		Image  string `json:"image"`
		Digest string `json:"digest"`
	}
	// Date stores the registry data
	Data struct {
		RegistryType RegistryType `json:"registryType"`
		RegistryUrl  string       `json:"registryUrl"`
		Images       []Image      `json:"images"`
	}
	// DockerArtifact is the current artifact
	DockerArtifact struct {
		Kind string `json:"kind"`
		Data Data   `json:"data"`
	}
)

// WritePluginArtifactFile writes the docker artifact data to the provided artifact file
func WritePluginArtifactFile(registryType RegistryType, artifactFilePath, registryURL, imageName, digest string, tags []string) error {
	var images []Image
	for _, tag := range tags {
		images = append(images, Image{
			Image:  fmt.Sprintf("%s:%s", imageName, tag),
			Digest: digest,
		})
	}
	data := Data{
		RegistryType: registryType,
		RegistryUrl:  registryURL,
		Images:       images,
	}

	dockerArtifact := DockerArtifact{
		Kind: dockerArtifactV1,
		Data: data,
	}

	b, err := json.MarshalIndent(dockerArtifact, "", "\t")
	if err != nil {
		return fmt.Errorf("failed with err %s to marshal output %+v", err, dockerArtifact)
	}

	dir := filepath.Dir(artifactFilePath)
	err = os.MkdirAll(dir, 0644)
	if err != nil {
		return fmt.Errorf("failed with err %s to create %s directory for artifact file", err, dir)
	}

	err = os.WriteFile(artifactFilePath, b, 0644)
	if err != nil {
		return fmt.Errorf("failed to write artifact to artifact file %s", artifactFilePath)
	}
	return nil
}
