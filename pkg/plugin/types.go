// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import "time"

type (
	// BaseConfig is the common configuration for a plugin.
	//
	// The configuration organizes all the information available to a plugin
	// executing as a step within a stage.
	//
	// Plugins can choose to compose this within their own config.
	//
	//     import "github.com/drone-plugins/drone-plugin-lib/pkg/plugin"
	//
	//     type MyPluginConfig struct {
	//         plugin.BaseConfig
	//         Foo string
	//         Bar string
	//     }
	BaseConfig struct {
		Stage Stage
		Step  Step
	}

	// Stage represents a build stage.
	Stage struct {
		// Arch is the platform architecture of the current build stage.
		Arch string

		// DependsOn is a list of dependencies for the current build stage.
		DependsOn []string

		// Finished is the unix timestamp for when the pipeline is finished.
		//
		// A running pipleine cannot have a finish timestamp, therefore, the
		// system aways sets this value to the current timestamp.
		Finished time.Time

		// Kind is the kind of resource being executed.
		//
		// This value is sourced from the `kind` attribute in the yaml
		// configuration file
		Kind string

		// Machine provides the name of the host machine on which the build
		// stage is currently running.
		Machine string

		// Name is the name for the current running build stage.
		Name string

		// Number is the stage number for the current running build stage.
		Number int

		// OS is the target operating system for the current build stage.
		OS string

		// Started is the unix timestamp for when a build stage was started by
		// the runner.
		Started time.Time

		// Status is the status for the current running build stage.
		//
		// If all of the stage's steps are passing, the status defaults to
		// success.
		Status string

		// Type is the type of resource being executed.
		Type string

		// Variant is the target architecture variant for the current build
		// stage.
		Variant string

		// Version is OS version for the current build stage.
		Version string
	}

	// Step represents the currently running step within the stage.
	Step struct {
		Name   string
		Number int
	}
)
