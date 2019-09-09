// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package plugin

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// List of enviornment variables set by Drone when running a step within a
// build stage.
//
// Multiple values are specified with `,` within the string. If there are
// multiple values this is to support backward compatibility with versions of
// Drone prior to the 1.0 release.

const (
	// The following environment variables are being ignored currently
	//
	// * DRONE_GIT_HTTP_URL
	// * DRONE_GIT_SSH_URL
	// * DRONE_REPO_NAMESPACE - Redundant to DRONE_REPO_OWNER

	//---------------------------------------------------------------------
	// Repo Enviornment Variables
	//---------------------------------------------------------------------

	// RepoDefaultBranchEnvVar corresponds to Repo.DefaultBranch.
	RepoDefaultBranchEnvVar = "DRONE_REPO_BRANCH"
	// RepoFullNameEnvVar corresponds to Repo.FullName.
	RepoFullNameEnvVar = "DRONE_REPO"
	// RepoLinkEnvVar corresponds to Repo.Link.
	RepoLinkEnvVar = "DRONE_REPO_LINK"
	// RepoNameEnvVar corresponds to Repo.Name
	RepoNameEnvVar = "DRONE_REPO_NAME"
	// RepoOwnerEnvVar corresponds to Repo.Owner.
	RepoOwnerEnvVar = "DRONE_REPO_OWNER"
	// RepoPrivateEnvVar corresponds to Repo.Private.
	RepoPrivateEnvVar = "DRONE_REPO_PRIVATE"
	// RepoRemoteURLEnvVar corresponds to Repo.RemoteURL.
	RepoRemoteURLEnvVar = "DRONE_REMOTE_URL"
	// RepoSCMEnvVar corresponds to Repo.SCM.
	RepoSCMEnvVar = "DRONE_REPO_SCM"
	// RepoVisibilityEnvVar corresponds to Repo.Visbility.
	RepoVisibilityEnvVar = "DRONE_REPO_VISIBILITY"

	//---------------------------------------------------------------------
	// Stage Enviornment Variables
	//---------------------------------------------------------------------

	// StageArchEnvVar corresponds to Stage.Arch.
	StageArchEnvVar = "DRONE_STAGE_ARCH"
	// StageDependsOnEnvVar corresponds to Stage.DependsOn.
	StageDependsOnEnvVar = "DRONE_STAGE_DEPENDS_ON"
	// StageFinishedEnvVar corresponds to Stage.Finished.
	StageFinishedEnvVar = "DRONE_STAGE_FINISHED"
	// StageKindEnvVar corresponds Stage.Kind.
	StageKindEnvVar = "DRONE_STAGE_KIND"
	// StageMachineEnvVar corresponds to Stage.Machine.
	StageMachineEnvVar = "DRONE_STAGE_MACHINE"
	// StageNameEnvVar corresponds to Stage.Name.
	StageNameEnvVar = "DRONE_STAGE_NAME"
	// StageNumberEnvVar corresponds to Stage.Number.
	StageNumberEnvVar = "DRONE_STAGE_NUMBER"
	// StageOSEnvVar corresponds to Stage.OS.
	StageOSEnvVar = "DRONE_STAGE_OS"
	// StageStartedEnvVar corresponds to Stage.Started.
	StageStartedEnvVar = "DRONE_STAGE_STARTED"
	// StageStatusEnvVar corresponds to Stage.Status.
	StageStatusEnvVar = "DRONE_STAGE_STATUS"
	// StageTypeEnvVar corresponds to Stage.Type.
	StageTypeEnvVar = "DRONE_STAGE_TYPE"
	// StageVariantEnvVar corresponds to Stage.Variant.
	StageVariantEnvVar = "DRONE_STAGE_VARIANT"
	// StageVersionEnvVar corresponds to Stage.Version.
	StageVersionEnvVar = "DRONE_STAGE_VERSION"

	//---------------------------------------------------------------------
	// Step Environment Variables
	//---------------------------------------------------------------------

	// StepNameEnvVar corresponds to Step.Name.
	StepNameEnvVar = "DRONE_STEP_NAME"
	// StepNumberEnvVar corresponds to Step.Number.
	StepNumberEnvVar = "DRONE_STEP_NUMBER"

	//---------------------------------------------------------------------
	// SemVer Variables
	//---------------------------------------------------------------------

	// SemVerBuildEnvVar corresponds to SemVer.Build.
	SemVerBuildEnvVar = "DRONE_SEMVER_BUILD"
	// SemVerErrorEnvVar corresponds to SemVer.Error.
	SemVerErrorEnvVar = "DRONE_SEMVER_ERROR"
	// SemVerMajorEnvVar corresponds to SemVer.Major.
	SemVerMajorEnvVar = "DRONE_SEMVER_MAJOR"
	// SemVerMinorEnvVar corresponds to SemVer.Minor.
	SemVerMinorEnvVar = "DRONE_SEMVER_MINOR"
	// SemVerPatchEnvVar corresponds to SemVer.Patch.
	SemVerPatchEnvVar = "DRONE_SEMVER_PATCH"
	// SemVerPrereleaseEnvVar corresponds to SemVer.Prerelease
	SemVerPrereleaseEnvVar = "DRONE_SEMVER_PRERELEASE"
	// SemVerShortEnvVar corresponds to SemVer.Short.
	SemVerShortEnvVar = "DRONE_SEMVER_SHORT"
	// SemVerVersionEnvVar corresponds to SemVer.Version
	SemVerVersionEnvVar = "DRONE_SEMVER"
)

// StringEnvVar gets the environment variable's value.
//
// If the value is not present then this returns the empty string.
func StringEnvVar(envVar string) string {
	return envVarValue(envVar)
}

// StringSliceEnvVar gets the environment variable as a string slice.
//
// If the value is not present then this returns an empty slice.
func StringSliceEnvVar(envVar string) []string {
	return strings.Split(envVarValue(envVar), ",")
}

// IntEnvVar gets the environment variable as an int.
//
// If the value is not present then this returns 0.
func IntEnvVar(envVar string) int {
	return int(Int64EnvVar(envVar))
}

// Int64EnvVar gets the environment variable as an int64.
//
// If the value is not present then this returns 0.
func Int64EnvVar(envVar string) int64 {
	value, err := strconv.ParseInt(envVarValue(envVar), 0, 64)

	if err != nil {
		return 0
	}

	return value
}

// UintEnvVar gets the environment variable as an uint.
//
// If the value is not present then this returns 0.
func UintEnvVar(envVar string) uint {
	return uint(Uint64EnvVar(envVar))
}

// Uint64EnvVar gets the environment variable as an uint64.
//
// If the value is not present then this returns 0.
func Uint64EnvVar(envVar string) uint64 {
	value, err := strconv.ParseUint(envVarValue(envVar), 0, 64)

	if err != nil {
		return 0
	}

	return value
}

// BoolEnvVar gets the environment variable as a bool.
//
// If the value is not present then this returns false.
func BoolEnvVar(envVar string) bool {
	value, err := strconv.ParseBool(envVarValue(envVar))

	if err != nil {
		return false
	}

	return value
}

// TimeEnvVar gets the environment variable as a Time created from a unix
// timestamp.
//
// If the value is not present then this returns time.Unix(0).
func TimeEnvVar(envVar string) time.Time {
	return time.Unix(Int64EnvVar(envVar), 0)
}

// envVarValue returns the first matched environment variable or the empty
// string if none was found.
func envVarValue(envVar string) string {
	return os.Getenv(envVar)
}
