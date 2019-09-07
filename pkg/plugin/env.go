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
)
