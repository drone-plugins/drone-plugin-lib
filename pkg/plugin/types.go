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

	// Repo represents the repository for the build.
	Repo struct {
		DefaultBranch string
		FullName      string
		Link          string
		Name          string
		Owner         string
		Private       bool
		RemoteURL     string
		SCM           string
		Visibility    string
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

	// SemVer represents the semantic version of the currently running build.
	//
	// This value is only applicable for tags. If the tag cannot be parsed into
	// a semantic version then SemVer.Error will have the reason.
	SemVer struct {
		// Build version number.
		//
		// This is signified by a + at the end of the tag.
		Build string
		// Error is the semantic version parsing error if the tag was invalid.
		Error string
		// Major version number.
		Major string
		// Minor version number.
		Minor string
		// Patch version number.
		Patch string
		// Prerelease version.
		Prerelease string
		// Short version of the semantic version string where labels and
		// metadata are truncated.
		Short string
		// Version is the full semantic version.
		Version string
	}
)

// RepoFromEnv creates a Repo from the environment variables used by Drone.
func RepoFromEnv() Repo {
	return Repo{
		DefaultBranch: StringEnvVar(RepoDefaultBranchEnvVar),
		FullName:      StringEnvVar(RepoFullNameEnvVar),
		Link:          StringEnvVar(RepoLinkEnvVar),
		Name:          StringEnvVar(RepoNameEnvVar),
		Owner:         StringEnvVar(RepoOwnerEnvVar),
		Private:       BoolEnvVar(RepoPrivateEnvVar),
		RemoteURL:     StringEnvVar(RepoRemoteURLEnvVar),
		SCM:           StringEnvVar(RepoSCMEnvVar),
		Visibility:    StringEnvVar(RepoVisibilityEnvVar),
	}
}

// StageFromEnv creates a Stage from the environment variables used by Drone.
func StageFromEnv() Stage {
	return Stage{
		Arch:      StringEnvVar(StageArchEnvVar),
		DependsOn: StringSliceEnvVar(StageDependsOnEnvVar),
		Finished:  TimeEnvVar(StageFinishedEnvVar),
		Kind:      StringEnvVar(StageKindEnvVar),
		Machine:   StringEnvVar(StageMachineEnvVar),
		Name:      StringEnvVar(StageNameEnvVar),
		Number:    IntEnvVar(StageNumberEnvVar),
		OS:        StringEnvVar(StageOSEnvVar),
		Started:   TimeEnvVar(StageStartedEnvVar),
		Status:    StringEnvVar(StageStatusEnvVar),
		Type:      StringEnvVar(StageTypeEnvVar),
		Variant:   StringEnvVar(StageVariantEnvVar),
		Version:   StringEnvVar(StageVersionEnvVar),
	}
}

// StepFromEnv creates a Step from the environment variables used by Drone.
func StepFromEnv() Step {
	return Step{
		Name:   StringEnvVar(StepNameEnvVar),
		Number: IntEnvVar(StepNumberEnvVar),
	}
}

// SemVerFromEnv creates a SemVer from the environment variables used by Drone.
func SemVerFromEnv() SemVer {
	return SemVer{
		Build:      StringEnvVar(SemVerBuildEnvVar),
		Error:      StringEnvVar(SemVerErrorEnvVar),
		Major:      StringEnvVar(SemVerMajorEnvVar),
		Minor:      StringEnvVar(SemVerMinorEnvVar),
		Patch:      StringEnvVar(SemVerPatchEnvVar),
		Prerelease: StringEnvVar(SemVerPrereleaseEnvVar),
		Short:      StringEnvVar(SemVerShortEnvVar),
		Version:    StringEnvVar(SemVerVersionEnvVar),
	}
}
