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

type RegistryType string

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
	Image struct {
		Image  string `json:"image"`
		Digest string `json:"digest"`
	}
	Data struct {
		RegistryType RegistryType `json:"registryType"`
		RegistryUrl  string       `json:"registryUrl"`
		Images       []Image      `json:"images"`
	}
	DockerArtifact struct {
		Kind string `json:"kind"`
		Data Data   `json:"data"`
	}
)

func WritePluginArtifactFile(registryType RegistryType, artifactFilePath, registryUrl, imageName, digest string, tags []string) error {
	var images []Image
	for _, tag := range tags {
		images = append(images, Image{
			Image:  fmt.Sprintf("%s:%s", imageName, tag),
			Digest: digest,
		})
	}
	data := Data{
		RegistryType: registryType,
		RegistryUrl:  registryUrl,
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
