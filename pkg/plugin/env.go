// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

// List of enviornment variables set by Drone when running a step within a
// build stage.
//
// Multiple values are specified with `,` within the string. If there are
// multiple values this is to support backward compatibility with versions of
// Drone prior to the 1.0 release.

const (
	// StepNameEnvVar is the environment variable for setting Step.Name.
	StepNameEnvVar = "DRONE_STEP_NAME"
	// StepNumberEnvVar is the environment variable for setting Step.Number.
	StepNumberEnvVar = "DRONE_STEP_NUMBER"
)
