// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

import (
	"time"
)

// Build represents a build of a repository.
type Build struct {
	// Branch defines the branch name of the build.
	Branch string

	// PullRequest number of the build.
	PullRequest int

	// Tag of the build.
	Tag string

	// SourceBranch for the pull request.
	SourceBranch string

	// TargetBranch for the pull request.
	TargetBranch string

	// Number for the build.
	Number int

	// Parent build number for the build.
	Parent int

	// Event that triggered the build.
	Event string

	// Action that triggered the build. This value is used to differentiate
	// bettween a pull request being opened vs synchronized.
	Action string

	// Status of the build.
	Status string

	// Created time of the build.
	Created time.Time

	// Started time of the build.
	Started time.Time

	// Finished time of the build.
	Finished time.Time

	// DeployTo the environment.
	DeployTo string

	// FailedStages of the build.
	FailedStages []string

	// FailedSteps of the build.
	FailedSteps []string
}
