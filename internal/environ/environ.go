// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package environ

// List of enviornment variables set by Drone when running a step within a
// build stage.
//
// Multiple values are specified with `,` within the string. If there are
// multiple values this is to support backward compatibility with versions of
// Drone prior to the 1.0 release.

const (
	// The following environment variables are being ignored currently
	//
	// * DRONE_COMMIT - Redundant with DRONE_COMMIT_SHA
	// * DRONE_GIT_HTTP_URL
	// * DRONE_GIT_SSH_URL
	// * DRONE_REPO_NAMESPACE - Redundant to DRONE_REPO_OWNER

	//---------------------------------------------------------------------
	// Build Enviornment Variables
	//---------------------------------------------------------------------

	// BuildActionEnvVar corresponds to Build.Action.
	BuildActionEnvVar = "DRONE_BUILD_ACTION"
	// BuildCreatedEnvVar corresponds to Build.Created.
	BuildCreatedEnvVar = "DRONE_BUILD_CREATED"
	// BuildDeployToEnvVar corresponds to Build.DeployTo.
	BuildDeployToEnvVar = "DRONE_DEPLOY_TO"
	// BuildEventEnvVar corresponds to Build.Event.
	BuildEventEnvVar = "DRONE_BUILD_EVENT"
	// BuildFailedStagesEnvVar corresponds to Build.FailedStages.
	BuildFailedStagesEnvVar = "DRONE_FAILED_STAGES"
	// BuildFailedStepsEnvVar corresponds to Build.FailedSteps.
	BuildFailedStepsEnvVar = "DRONE_FAILED_STEPS"
	// BuildFinishedEnvVar corresponds to Build.Finished.
	BuildFinishedEnvVar = "DRONE_BUILD_FINISHED"
	// BuildNumberEnvVar corresponds to Build.Created.
	BuildNumberEnvVar = "DRONE_BUILD_NUMBER"
	// BuildParentEnvVar corresponds to Build.Parent.
	BuildParentEnvVar = "DRONE_BUILD_PARENT"
	// BuildPullRequestEnvVar corresponds to Build.PullRequest.
	BuildPullRequestEnvVar = "DRONE_PULL_REQUEST"
	// BuildStartedEnvVar corresponds to Build.Started.
	BuildStartedEnvVar = "DRONE_BUILD_STARTED"
	// BuildStatusEnvVar corresponds to Build.Status.
	BuildStatusEnvVar = "DRONE_BUILD_STATUS"
	// BuildSourceBranchEnvVar corresponds to Build.SourceBranch.
	BuildSourceBranchEnvVar = "DRONE_SOURCE_BRANCH"
	// BuildTagEnvVar corresponds to Build.Tag.
	BuildTagEnvVar = "DRONE_TAG"
	// BuildTargetBranchEnvVar corresponds to Build.TargetBranch.
	BuildTargetBranchEnvVar = "DRONE_TARGET_BRANCH"

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
	// Commit Enviornment Variables
	//---------------------------------------------------------------------

	// CommitAfterEnvVar corresponds to Commit.After.
	CommitAfterEnvVar = "DRONE_COMMIT_AFTER"
	// CommitAuthorEnvVar corresponds to Commit.Author.
	CommitAuthorEnvVar = "DRONE_COMMIT_AUTHOR"
	// CommitAuthorAvatarEnvVar corresponds to Commit.AuthorAvatar.
	CommitAuthorAvatarEnvVar = "DRONE_COMMIT_AUTHOR_AVATAR"
	// CommitAuthorEmailEnvVar corresponds to Commit.AuthorEmail.
	CommitAuthorEmailEnvVar = "DRONE_COMMIT_AUTHOR_EMAIL"
	// CommitAuthorNameEnvVar corresponds to Commit.AuthorName.
	CommitAuthorNameEnvVar = "DRONE_COMMIT_AUTHOR_NAME"
	// CommitBeforeEnvVar corresponds to Commit.Before.
	CommitBeforeEnvVar = "DRONE_COMMIT_BEFORE"
	// CommitBranchEnvVar corresponds to Commit.Branch.
	CommitBranchEnvVar = "DRONE_COMMIT_BRANCH"
	// CommitLinkEnvVar corresponds to Commit.Link.
	CommitLinkEnvVar = "DRONE_COMMIT_LINK"
	// CommitMessageEnvVar corresponds to Commit.Message.
	CommitMessageEnvVar = "DRONE_COMMIT_MESSAGE"
	// CommitRefEnvVar corresponds to Commit.Ref.
	CommitRefEnvVar = "DRONE_COMMIT_REF"
	// CommitSHAEnvVar corresponds to Commit.SHA.
	CommitSHAEnvVar = "DRONE_COMMIT_SHA"

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
