// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

// Package urfave provides helpers for interacting with the `urfave/cli`
// package when creating plugins for use by the Drone CI/CD service.
//
// Drone communicates to plugins by passing in environment variables that have
// information on the currently executing build. The `urfave/cli` package can
// read these environment variables and extract them into structs.
//
//     import(
//         "github.com/drone-plugins/drone-plugin-lib/pkg/urfave"
//         "github.com/urfave/cli"
//	   )
//
//     func main() {
//         app := cli.New()
//         app.Name = "my awesome Drone plugin"
//         app.Run = run
//         app.Flags = []cli.Flags{
//             // All my plugin flags
//         }
//         app.Flags = append(
//             app.Flags,
//             urfave.CommitFlags()...,
//         )
//     }
package urfave

import (
	"time"

	"github.com/drone-plugins/drone-plugin-lib/pkg/plugin"
	"github.com/urfave/cli"
)

//---------------------------------------------------------------------
// Build Flags
//---------------------------------------------------------------------

const (
	// BuildActionFlag corresponds to plugin.Build.Action.
	BuildActionFlag = "build.action"
	// BuildCreatedFlag corresponds to plugin.Build.Created.
	BuildCreatedFlag = "build.created"
	// BuildDeployToFlag corresponds to plugin.Build.DeployTo.
	BuildDeployToFlag = "build.deploy-to"
	// BuildEventFlag corresponds to plugin.Build.Event.
	BuildEventFlag = "build.event"
	// BuildFailedStagesFlag corresponds to plugin.Build.FailedStages.
	BuildFailedStagesFlag = "build.failed-stages"
	// BuildFailedStepsFlag corresponds to plugin.Build.FailedSteps.
	BuildFailedStepsFlag = "build.failed-steps"
	// BuildFinishedFlag corresponds to plugin.Build.Finished.
	BuildFinishedFlag = "build.finished"
	// BuildNumberFlag corresponds to plugin.Build.Created.
	BuildNumberFlag = "build.number"
	// BuildParentFlag corresponds to plugin.Build.Parent.
	BuildParentFlag = "build.parent"
	// BuildPullRequestFlag corresponds to plugin.Build.PullRequest.
	BuildPullRequestFlag = "build.pull-request"
	// BuildSourceBranchFlag corresponds to plugin.Build.SourceBranch.
	BuildSourceBranchFlag = "build.source-branch"
	// BuildStartedFlag corresponds to plugin.Build.Started.
	BuildStartedFlag = "build.started"
	// BuildStatusFlag corresponds to plugin.Build.Status.
	BuildStatusFlag = "build.status"
	// BuildTagFlag corresponds to plugin.Build.Tag.
	BuildTagFlag = "build.tag"
	// BuildTargetBranchFlag corresponds to plugin.Build.TargetBranch.
	BuildTargetBranchFlag = "build.target-branch"
)

// BuildFlags has the cli.Flags for the plugin.Build.
func BuildFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   BuildActionFlag,
			Usage:  "build action",
			EnvVar: plugin.BuildActionEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildCreatedFlag,
			Usage:  "build created",
			EnvVar: plugin.BuildCreatedEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildDeployToFlag,
			Usage:  "build deploy to",
			EnvVar: plugin.BuildDeployToEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildEventFlag,
			Usage:  "build event",
			EnvVar: plugin.BuildEventEnvVar,
			Hidden: true,
		},
		cli.StringSliceFlag{
			Name:   BuildFailedStagesFlag,
			Usage:  "build failed stages",
			EnvVar: plugin.BuildFailedStagesEnvVar,
			Hidden: true,
		},
		cli.StringSliceFlag{
			Name:   BuildFailedStepsFlag,
			Usage:  "build failed steps",
			EnvVar: plugin.BuildFailedStepsEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildFinishedFlag,
			Usage:  "build finished",
			EnvVar: plugin.BuildFinishedEnvVar,
			Hidden: true,
		},
		cli.IntFlag{
			Name:   BuildNumberFlag,
			Usage:  "build number",
			EnvVar: plugin.BuildNumberEnvVar,
			Hidden: true,
		},
		cli.IntFlag{
			Name:   BuildParentFlag,
			Usage:  "build parent",
			EnvVar: plugin.BuildParentEnvVar,
			Hidden: true,
		},
		cli.IntFlag{
			Name:   BuildPullRequestFlag,
			Usage:  "build pull request",
			EnvVar: plugin.BuildPullRequestEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildSourceBranchFlag,
			Usage:  "build source branch",
			EnvVar: plugin.BuildSourceBranchEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildStartedFlag,
			Usage:  "build started",
			EnvVar: plugin.BuildStartedEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildStatusFlag,
			Usage:  "build status",
			EnvVar: plugin.BuildStatusEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildTagFlag,
			Usage:  "build tag",
			EnvVar: plugin.BuildTagEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   BuildTargetBranchFlag,
			Usage:  "build target branch",
			EnvVar: plugin.BuildTargetBranchEnvVar,
			Hidden: true,
		},
	}
}

// BuildFromContext creates a plugin.Build from the cli.Context.
func BuildFromContext(ctx *cli.Context) plugin.Build {
	return plugin.Build{
		Action:       ctx.String(BuildActionFlag),
		Created:      time.Unix(ctx.Int64(BuildCreatedFlag), 0),
		DeployTo:     ctx.String(BuildDeployToFlag),
		Event:        ctx.String(BuildEventFlag),
		FailedStages: ctx.StringSlice(BuildFailedStagesFlag),
		FailedSteps:  ctx.StringSlice(BuildFailedStepsFlag),
		Finished:     time.Unix(ctx.Int64(BuildFinishedFlag), 0),
		Number:       ctx.Int(BuildNumberFlag),
		Parent:       ctx.Int(BuildParentFlag),
		PullRequest:  ctx.Int(BuildPullRequestFlag),
		SourceBranch: ctx.String(BuildSourceBranchFlag),
		Started:      time.Unix(ctx.Int64(BuildStartedFlag), 0),
		Status:       ctx.String(BuildStatusFlag),
		Tag:          ctx.String(BuildTagFlag),
		TargetBranch: ctx.String(BuildTargetBranchFlag),
	}
}

//---------------------------------------------------------------------
// Repo Flags
//---------------------------------------------------------------------

const (
	// RepoDefaultBranchFlag corresponds to plugin.Repo.DefaultBranch.
	RepoDefaultBranchFlag = "repo.branch"
	// RepoFullNameFlag corresponds to plugin.Repo.FullName.
	RepoFullNameFlag = "repo.full-name"
	// RepoLinkFlag corresponds to plugin.Repo.Link.
	RepoLinkFlag = "repo.link"
	// RepoNameFlag corresponds to plugin.Repo.Name
	RepoNameFlag = "repo.name"
	// RepoOwnerFlag corresponds to plugin.Repo.Owner.
	RepoOwnerFlag = "repo.owner"
	// RepoPrivateFlag corresponds to plugin.Repo.Private.
	RepoPrivateFlag = "repo.private"
	// RepoRemoteURLFlag corresponds to plugin.Repo.RemoteURL.
	RepoRemoteURLFlag = "repo.remote-url"
	// RepoSCMFlag corresponds to plugin.Repo.SCM.
	RepoSCMFlag = "repo.scm"
	// RepoVisibilityFlag corresponds to plugin.Repo.Visbility.
	RepoVisibilityFlag = "repo.visibility"
)

// RepoFlags has the cli.Flags for the plugin.Repo
func RepoFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   RepoDefaultBranchFlag,
			Usage:  "repo default branch",
			EnvVar: plugin.RepoDefaultBranchEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   RepoFullNameFlag,
			Usage:  "repo full name",
			EnvVar: plugin.RepoFullNameEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   RepoLinkFlag,
			Usage:  "repo link",
			EnvVar: plugin.RepoLinkEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   RepoNameFlag,
			Usage:  "repo name",
			EnvVar: plugin.RepoNameEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   RepoOwnerFlag,
			Usage:  "repo owner",
			EnvVar: plugin.RepoOwnerEnvVar,
			Hidden: true,
		},
		cli.BoolFlag{
			Name:   RepoPrivateFlag,
			Usage:  "repo private",
			EnvVar: plugin.RepoPrivateEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   RepoRemoteURLFlag,
			Usage:  "repo remote url",
			EnvVar: plugin.RepoRemoteURLEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   RepoSCMFlag,
			Usage:  "repo scm",
			EnvVar: plugin.RepoSCMEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   RepoVisibilityFlag,
			Usage:  "repo visibility",
			EnvVar: plugin.RepoVisibilityEnvVar,
			Hidden: true,
		},
	}
}

// RepoFromContext creates a plugin.Repo from the cli.Context.
func RepoFromContext(ctx *cli.Context) plugin.Repo {
	return plugin.Repo{
		DefaultBranch: ctx.String(RepoDefaultBranchFlag),
		FullName:      ctx.String(RepoFullNameFlag),
		Link:          ctx.String(RepoLinkFlag),
		Name:          ctx.String(RepoNameFlag),
		Owner:         ctx.String(RepoOwnerFlag),
		Private:       ctx.Bool(RepoPrivateFlag),
		RemoteURL:     ctx.String(RepoRemoteURLFlag),
		SCM:           ctx.String(RepoSCMFlag),
		Visibility:    ctx.String(RepoVisibilityFlag),
	}
}

//---------------------------------------------------------------------
// Commit Flags
//---------------------------------------------------------------------

const (
	// CommitAfterFlag corresponds to plugin.Commit.After.
	CommitAfterFlag = "commit.after"
	// CommitAuthorFlag corresponds to plugin.Commit.Author.
	CommitAuthorFlag = "commit.author"
	// CommitAuthorAvatarFlag corresponds to plugin.Commit.AuthorAvatar.
	CommitAuthorAvatarFlag = "commit.author-avatar"
	// CommitAuthorEmailFlag corresponds to plugin.Commit.AuthorEmail.
	CommitAuthorEmailFlag = "commit.author-email"
	// CommitAuthorNameFlag corresponds to plugin.Commit.AuthorName.
	CommitAuthorNameFlag = "commit.author-name"
	// CommitBeforeFlag corresponds to plugin.Commit.Before.
	CommitBeforeFlag = "commit.before"
	// CommitBranchFlag corresponds to plugin.Commit.Branch.
	CommitBranchFlag = "commit.branch"
	// CommitLinkFlag corresponds to plugin.Commit.Link.
	CommitLinkFlag = "commit.link"
	// CommitMessageFlag corresponds to plugin.Commit.Message.
	CommitMessageFlag = "commit.message"
	// CommitRefFlag corresponds to plugin.Commit.Ref.
	CommitRefFlag = "commit.ref"
	// CommitSHAFlag corresponds to plugin.Commit.SHA.
	CommitSHAFlag = "commit.sha"
)

// CommitFlags has the cli.Flags for the plugin.Commit.
func CommitFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   CommitAfterFlag,
			Usage:  "commit after",
			EnvVar: plugin.CommitAfterEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitAuthorFlag,
			Usage:  "commit author",
			EnvVar: plugin.CommitAuthorEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitAuthorAvatarFlag,
			Usage:  "commit author avatar",
			EnvVar: plugin.CommitAuthorAvatarEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitAuthorEmailFlag,
			Usage:  "commit author email",
			EnvVar: plugin.CommitAuthorEmailEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitAuthorNameFlag,
			Usage:  "commit author name",
			EnvVar: plugin.CommitAuthorNameEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitBeforeFlag,
			Usage:  "commit before",
			EnvVar: plugin.CommitBeforeEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitBranchFlag,
			Usage:  "commit branch",
			EnvVar: plugin.CommitBranchEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitLinkFlag,
			Usage:  "commit link",
			EnvVar: plugin.CommitLinkEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitMessageFlag,
			Usage:  "commit message",
			EnvVar: plugin.CommitMessageEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitRefFlag,
			Usage:  "commit ref",
			EnvVar: plugin.CommitRefEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   CommitSHAFlag,
			Usage:  "commit sha",
			EnvVar: plugin.CommitSHAEnvVar,
			Hidden: true,
		},
	}
}

// CommitFromContext creates a plugin.Commit from the cli.Context.
func CommitFromContext(ctx *cli.Context) plugin.Commit {
	return plugin.Commit{
		After:        ctx.String(CommitAfterFlag),
		Author:       ctx.String(CommitAuthorFlag),
		AuthorAvatar: ctx.String(CommitAuthorAvatarFlag),
		AuthorEmail:  ctx.String(CommitAuthorEmailFlag),
		AuthorName:   ctx.String(CommitAuthorNameFlag),
		Before:       ctx.String(CommitBeforeFlag),
		Branch:       ctx.String(CommitBranchFlag),
		Link:         ctx.String(CommitLinkFlag),
		Message:      ctx.String(CommitMessageFlag),
		Ref:          ctx.String(CommitRefFlag),
		SHA:          ctx.String(CommitSHAFlag),
	}
}

//---------------------------------------------------------------------
// Stage Flags
//---------------------------------------------------------------------

const (
	// StageArchFlag corresponds to plugin.Stage.Arch.
	StageArchFlag = "stage.arch"
	// StageDependsOnFlag corresponds to plugin.Stage.DependsOn.
	StageDependsOnFlag = "stage.depends-on"
	// StageFinishedFlag corresponds to plugin.Stage.Finished.
	StageFinishedFlag = "stage.finished"
	// StageKindFlag corresponds Stage.Kind.
	StageKindFlag = "stage.kind"
	// StageMachineFlag corresponds to plugin.Stage.Machine.
	StageMachineFlag = "stage.machine"
	// StageNameFlag corresponds to plugin.Stage.Name.
	StageNameFlag = "stage.name"
	// StageNumberFlag corresponds to plugin.Stage.Number.
	StageNumberFlag = "stage.number"
	// StageOSFlag corresponds to plugin.Stage.OS.
	StageOSFlag = "stage.os"
	// StageStartedFlag corresponds to plugin.Stage.Started.
	StageStartedFlag = "stage.started"
	// StageStatusFlag corresponds to plugin.Stage.Status.
	StageStatusFlag = "stage.status"
	// StageTypeFlag corresponds to plugin.Stage.Type.
	StageTypeFlag = "stage.type"
	// StageVariantFlag corresponds to plugin.Stage.Variant.
	StageVariantFlag = "stage.variant"
	// StageVersionFlag corresponds to plugin.Stage.Version.
	StageVersionFlag = "stage.version"
)

// StageFlags has the cli.Flags for the plugin.Stage
func StageFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   StageArchFlag,
			Usage:  "stage arch",
			EnvVar: plugin.StageArchEnvVar,
			Hidden: true,
		},
		cli.StringSliceFlag{
			Name:   StageDependsOnFlag,
			Usage:  "stage depends on",
			EnvVar: plugin.StageDependsOnEnvVar,
			Hidden: true,
		},
		cli.Int64Flag{
			Name:   StageFinishedFlag,
			Usage:  "stage finished",
			EnvVar: plugin.StageFinishedEnvVar,
			Hidden: true,
		}, cli.StringFlag{
			Name:   StageKindFlag,
			Usage:  "stage kind",
			EnvVar: plugin.StageKindEnvVar,
			Hidden: true,
		}, cli.StringFlag{
			Name:   StageMachineFlag,
			Usage:  "stage machine",
			EnvVar: plugin.StageMachineEnvVar,
			Hidden: true,
		}, cli.StringFlag{
			Name:   StageNameFlag,
			Usage:  "stage name",
			EnvVar: plugin.StageNameEnvVar,
			Hidden: true,
		}, cli.IntFlag{
			Name:   StageNumberFlag,
			Usage:  "stage number",
			EnvVar: plugin.StageNumberEnvVar,
			Hidden: true,
		}, cli.StringFlag{
			Name:   StageOSFlag,
			Usage:  "stage os",
			EnvVar: plugin.StageOSEnvVar,
			Hidden: true,
		}, cli.Int64Flag{
			Name:   StageStartedFlag,
			Usage:  "stage started",
			EnvVar: plugin.StageStartedEnvVar,
			Hidden: true,
		}, cli.StringFlag{
			Name:   StageStatusFlag,
			Usage:  "stage status",
			EnvVar: plugin.StageStatusEnvVar,
			Hidden: true,
		}, cli.StringFlag{
			Name:   StageTypeFlag,
			Usage:  "stage type",
			EnvVar: plugin.StageTypeEnvVar,
			Hidden: true,
		}, cli.StringFlag{
			Name:   StageVariantFlag,
			Usage:  "stage variant",
			EnvVar: plugin.StageVariantEnvVar,
			Hidden: true,
		}, cli.StringFlag{
			Name:   StageVersionFlag,
			Usage:  "stage version",
			EnvVar: plugin.StageVersionEnvVar,
			Hidden: true,
		},
	}
}

// StageFromContext creates a plugin.Stage from the cli.Context.
func StageFromContext(ctx *cli.Context) plugin.Stage {
	return plugin.Stage{
		Arch:      ctx.String(StageArchFlag),
		DependsOn: ctx.StringSlice(StageDependsOnFlag),
		Finished:  time.Unix(ctx.Int64(StageFinishedFlag), 0),
		Kind:      ctx.String(StageKindFlag),
		Machine:   ctx.String(StageMachineFlag),
		Name:      ctx.String(StageNameFlag),
		Number:    ctx.Int(StageNumberFlag),
		OS:        ctx.String(StageOSFlag),
		Started:   time.Unix(ctx.Int64(StageStartedFlag), 0),
		Status:    ctx.String(StageStatusFlag),
		Type:      ctx.String(StageTypeFlag),
		Variant:   ctx.String(StageVariantFlag),
		Version:   ctx.String(StageVersionFlag),
	}
}

//---------------------------------------------------------------------
// Step Flags
//---------------------------------------------------------------------

const (
	// StepNameFlag corresponds to plugin.Step.Name.
	StepNameFlag = "step.name"
	// StepNumberFlag corresponds to plugin.Step.Number.
	StepNumberFlag = "step.number"
)

// StepFlags has the cli.Flags for the plugin.Step.
func StepFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   StepNameFlag,
			Usage:  "step name",
			EnvVar: plugin.StepNameEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   StepNumberFlag,
			Usage:  "step number",
			EnvVar: plugin.StepNumberEnvVar,
			Hidden: true,
		},
	}
}

// StepFromContext creates a plugin.Step from the cli.Context.
func StepFromContext(ctx *cli.Context) plugin.Step {
	return plugin.Step{
		Name:   ctx.String(StepNameFlag),
		Number: ctx.Int(StepNumberFlag),
	}
}

//---------------------------------------------------------------------
// SemVer Flags
//---------------------------------------------------------------------

const (
	// SemVerBuildFlag corresponds to plugin.SemVer.Build.
	SemVerBuildFlag = "semver.build"
	// SemVerErrorFlag corresponds to plugin.SemVer.Error.
	SemVerErrorFlag = "semver.error"
	// SemVerMajorFlag corresponds to plugin.SemVer.Major.
	SemVerMajorFlag = "semver.major"
	// SemVerMinorFlag corresponds to plugin.SemVer.Minor.
	SemVerMinorFlag = "semver.minor"
	// SemVerPatchFlag corresponds to plugin.SemVer.Patch.
	SemVerPatchFlag = "semver.patch"
	// SemVerPrereleaseFlag corresponds to plugin.SemVer.Prerelease
	SemVerPrereleaseFlag = "semver.prerelease"
	// SemVerShortFlag corresponds to plugin.SemVer.Short.
	SemVerShortFlag = "semver.short"
	// SemVerVersionFlag corresponds to plugin.SemVer.Version
	SemVerVersionFlag = "semver.version"
)

// SemVerFlags has the cli.Flags for the plugin.SemVer.
func SemVerFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   SemVerBuildFlag,
			Usage:  "semver build",
			EnvVar: plugin.SemVerBuildEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   SemVerErrorFlag,
			Usage:  "semver error",
			EnvVar: plugin.SemVerErrorEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   SemVerMajorFlag,
			Usage:  "semver major",
			EnvVar: plugin.SemVerMajorEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   SemVerMinorFlag,
			Usage:  "semver minor",
			EnvVar: plugin.SemVerMinorEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   SemVerPatchFlag,
			Usage:  "semver patch",
			EnvVar: plugin.SemVerPatchEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   SemVerPrereleaseFlag,
			Usage:  "semver prerelease",
			EnvVar: plugin.SemVerPrereleaseEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   SemVerShortFlag,
			Usage:  "semver short",
			EnvVar: plugin.SemVerShortEnvVar,
			Hidden: true,
		},
		cli.StringFlag{
			Name:   SemVerVersionFlag,
			Usage:  "semver version",
			EnvVar: plugin.SemVerVersionEnvVar,
			Hidden: true,
		},
	}
}

// SemVerFromContext creates a plugin.Step from the cli.SemVer.
func SemVerFromContext(ctx *cli.Context) plugin.SemVer {
	return plugin.SemVer{
		Build:      ctx.String(SemVerBuildFlag),
		Error:      ctx.String(SemVerErrorFlag),
		Major:      ctx.String(SemVerMajorFlag),
		Minor:      ctx.String(SemVerMinorFlag),
		Patch:      ctx.String(SemVerPatchFlag),
		Prerelease: ctx.String(SemVerPrereleaseFlag),
		Short:      ctx.String(SemVerShortFlag),
		Version:    ctx.String(SemVerVersionFlag),
	}
}
