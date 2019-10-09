// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import "time"

type (
	// Environment for the plugin.
	//
	// Represents the full Drone environment that the plugin is executing in.
	Environment struct {
		Build  Build
		Repo   Repo
		Commit Commit
		Stage  Stage
		Step   Step
		SemVer SemVer
	}

	// Build represents a build of a repository.
	Build struct {
		// Action that triggered the build. This value is used to differentiate
		// bettween a pull request being opened vs synchronized.
		Action string
		// Created time of the build.
		Created time.Time
		// DeployTo the environment.
		DeployTo string
		// Event that triggered the build.
		Event string
		// FailedStages of the build.
		FailedStages []string
		// FailedSteps of the build.
		FailedSteps []string
		// Finished time of the build.
		Finished time.Time
		// Number for the build.
		Number int
		// Parent build number for the build.
		Parent int
		// PullRequest number of the build.
		PullRequest int
		// Started time of the build.
		Started time.Time
		// Status of the build.
		Status string
		// SourceBranch for the pull request.
		SourceBranch string
		// Tag of the build.
		Tag string
		// TargetBranch for the pull request.
		TargetBranch string
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

	// Commit represents the current commit being built.
	Commit struct {
		// After contains the commit sha after the patch is applied.
		After string
		// Author of the commit.
		Author string
		// AuthorAvatar of the commit.
		AuthorAvatar string
		// AuthorEmail of the commit.
		AuthorEmail string
		// AuthorName of the commit.
		AuthorName string
		// Before contains the commit sha before the patch is applied.
		Before string
		// Branch target for the push or pull request. This may be empty for
		// tag events.
		Branch string
		// Link to the commit or object in the source control management system.
		Link string
		// Message for the current commit.
		Message string
		// Ref for the current commit.
		Ref string
		// SHA for the current commit.
		SHA string
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

// Environ creates an Environment from all environment variables used by Drone.
func Environ() Environment {
	return Environment{
		Build:  BuildFromEnv(),
		Repo:   RepoFromEnv(),
		Commit: CommitFromEnv(),
		Stage:  StageFromEnv(),
		Step:   StepFromEnv(),
		SemVer: SemVerFromEnv(),
	}
}

// BuildFromEnv creates a Build from the environment variables used by Drone.
func BuildFromEnv() Build {
	return Build{
		Action:       StringEnvVar(BuildActionEnvVar),
		Created:      TimeEnvVar(BuildCreatedEnvVar),
		DeployTo:     StringEnvVar(BuildDeployToEnvVar),
		Event:        StringEnvVar(BuildEventEnvVar),
		FailedStages: StringSliceEnvVar(BuildFailedStagesEnvVar),
		FailedSteps:  StringSliceEnvVar(BuildFailedStepsEnvVar),
		Finished:     TimeEnvVar(BuildFinishedEnvVar),
		Number:       IntEnvVar(BuildNumberEnvVar),
		Parent:       IntEnvVar(BuildParentEnvVar),
		PullRequest:  IntEnvVar(BuildPullRequestEnvVar),
		SourceBranch: StringEnvVar(BuildSourceBranchEnvVar),
		Started:      TimeEnvVar(BuildStartedEnvVar),
		Status:       StringEnvVar(BuildStatusEnvVar),
		Tag:          StringEnvVar(BuildTagEnvVar),
		TargetBranch: StringEnvVar(BuildTargetBranchEnvVar),
	}
}

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

// CommitFromEnv creates a Commit from the environment variables used by Drone.
func CommitFromEnv() Commit {
	return Commit{
		After:        StringEnvVar(CommitAfterEnvVar),
		Author:       StringEnvVar(CommitAuthorEnvVar),
		AuthorAvatar: StringEnvVar(CommitAuthorAvatarEnvVar),
		AuthorEmail:  StringEnvVar(CommitAuthorEmailEnvVar),
		AuthorName:   StringEnvVar(CommitAuthorNameEnvVar),
		Before:       StringEnvVar(CommitBeforeEnvVar),
		Branch:       StringEnvVar(CommitBranchEnvVar),
		Link:         StringEnvVar(CommitLinkEnvVar),
		Message:      StringEnvVar(CommitMessageEnvVar),
		Ref:          StringEnvVar(CommitRefEnvVar),
		SHA:          StringEnvVar(CommitSHAEnvVar),
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
