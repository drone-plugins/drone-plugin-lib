// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

import (
	"time"
)

// Stage represents a build stage.
type Stage struct {
	// Kind is the kind of resource being executed.
	//
	// This value is sourced from the `kind` attribute in the yaml
	// configuration file
	Kind string

	// Type is the type of resource being executed.
	Type string

	// Name is the name for the current running build stage.
	Name string

	// Number is the stage number for the current running build stage.
	Number int

	// Machine provides the name of the host machine on which the build
	// stage is currently running.
	Machine string

	// OS is the target operating system for the current build stage.
	OS string

	// Arch is the platform architecture of the current build stage.
	Arch string

	// Variant is the target architecture variant for the current build
	// stage.
	Variant string

	// Version is OS version for the current build stage.
	Version string

	// Status is the status for the current running build stage.
	//
	// If all of the stage's steps are passing, the status defaults to
	// success.
	Status string

	// Started is the unix timestamp for when a build stage was started by
	// the runner.
	Started time.Time

	// Finished is the unix timestamp for when the pipeline is finished.
	//
	// A running pipleine cannot have a finish timestamp, therefore, the
	// system aways sets this value to the current timestamp.
	Finished time.Time

	// DependsOn is a list of dependencies for the current build stage.
	DependsOn []string
}
