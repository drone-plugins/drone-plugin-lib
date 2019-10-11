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
//             urfave.Flags()...,
//         )
//     }
//
//     func run(ctx *cli.Context) {
//         	pipeline := urfave.PipelineFromContext(ctx)
//          ....
//     }
package urfave

import (
	"time"

	"github.com/urfave/cli"

	"github.com/drone-plugins/drone-plugin-lib/internal/environ"
	"github.com/drone-plugins/drone-plugin-lib/pkg/plugin"
)

//---------------------------------------------------------------------
// Flags
//---------------------------------------------------------------------

// Flags for a urfave cli Drone plugin
func Flags() []cli.Flag {
	flags := []cli.Flag{}

	flags = append(flags, pipelineFlags()...)
	flags = append(flags, networkFlags()...)
	flags = append(flags, loggingFlags()...)

	return flags
}

//---------------------------------------------------------------------
// Pipeline flags
//---------------------------------------------------------------------

// pipelineFlags has the cli.Flags for the plugin.Pipeline.
func pipelineFlags() []cli.Flag {
	flags := []cli.Flag{}

	flags = append(flags, buildFlags()...)
	flags = append(flags, repoFlags()...)
	flags = append(flags, commitFlags()...)
	flags = append(flags, stageFlags()...)
	flags = append(flags, stepFlags()...)
	flags = append(flags, semVerFlags()...)

	return flags
}

// PipelineFromContext creates a plugin.Pipeline from the cli.Context.
func PipelineFromContext(ctx *cli.Context) plugin.Pipeline {
	return plugin.Pipeline{
		Build:  buildFromContext(ctx),
		Repo:   repoFromContext(ctx),
		Commit: commitFromContext(ctx),
		Stage:  stageFromContext(ctx),
		Step:   stepFromContext(ctx),
		SemVer: semVerFromContext(ctx),
	}
}

//---------------------------------------------------------------------
// Build Flags
//---------------------------------------------------------------------

const (
	buildActionFlag       = "build.action"
	buildCreatedFlag      = "build.created"
	buildDeployToFlag     = "build.deploy-to"
	buildEventFlag        = "build.event"
	buildFailedStagesFlag = "build.failed-stages"
	buildFailedStepsFlag  = "build.failed-steps"
	buildFinishedFlag     = "build.finished"
	buildNumberFlag       = "build.number"
	buildParentFlag       = "build.parent"
	buildPullRequestFlag  = "build.pull-request"
	buildSourceBranchFlag = "build.source-branch"
	buildStartedFlag      = "build.started"
	buildStatusFlag       = "build.status"
	buildTagFlag          = "build.tag"
	buildTargetBranchFlag = "build.target-branch"
)

// buildFlags has the cli.Flags for the plugin.Build.
func buildFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   buildActionFlag,
			Usage:  "build action",
			EnvVar: environ.BuildActionEnvVar,
		},
		cli.StringFlag{
			Name:   buildCreatedFlag,
			Usage:  "build created",
			EnvVar: environ.BuildCreatedEnvVar,
		},
		cli.StringFlag{
			Name:   buildDeployToFlag,
			Usage:  "build deploy to",
			EnvVar: environ.BuildDeployToEnvVar,
		},
		cli.StringFlag{
			Name:   buildEventFlag,
			Usage:  "build event",
			EnvVar: environ.BuildEventEnvVar,
		},
		cli.StringSliceFlag{
			Name:   buildFailedStagesFlag,
			Usage:  "build failed stages",
			EnvVar: environ.BuildFailedStagesEnvVar,
		},
		cli.StringSliceFlag{
			Name:   buildFailedStepsFlag,
			Usage:  "build failed steps",
			EnvVar: environ.BuildFailedStepsEnvVar,
		},
		cli.StringFlag{
			Name:   buildFinishedFlag,
			Usage:  "build finished",
			EnvVar: environ.BuildFinishedEnvVar,
		},
		cli.IntFlag{
			Name:   buildNumberFlag,
			Usage:  "build number",
			EnvVar: environ.BuildNumberEnvVar,
		},
		cli.IntFlag{
			Name:   buildParentFlag,
			Usage:  "build parent",
			EnvVar: environ.BuildParentEnvVar,
		},
		cli.IntFlag{
			Name:   buildPullRequestFlag,
			Usage:  "build pull request",
			EnvVar: environ.BuildPullRequestEnvVar,
		},
		cli.StringFlag{
			Name:   buildSourceBranchFlag,
			Usage:  "build source branch",
			EnvVar: environ.BuildSourceBranchEnvVar,
		},
		cli.StringFlag{
			Name:   buildStartedFlag,
			Usage:  "build started",
			EnvVar: environ.BuildStartedEnvVar,
		},
		cli.StringFlag{
			Name:   buildStatusFlag,
			Usage:  "build status",
			EnvVar: environ.BuildStatusEnvVar,
		},
		cli.StringFlag{
			Name:   buildTagFlag,
			Usage:  "build tag",
			EnvVar: environ.BuildTagEnvVar,
		},
		cli.StringFlag{
			Name:   buildTargetBranchFlag,
			Usage:  "build target branch",
			EnvVar: environ.BuildTargetBranchEnvVar,
		},
	}
}

// buildFromContext creates a plugin.Build from the cli.Context.
func buildFromContext(ctx *cli.Context) plugin.Build {
	return plugin.Build{
		Action:       ctx.String(buildActionFlag),
		Created:      time.Unix(ctx.Int64(buildCreatedFlag), 0),
		DeployTo:     ctx.String(buildDeployToFlag),
		Event:        ctx.String(buildEventFlag),
		FailedStages: ctx.StringSlice(buildFailedStagesFlag),
		FailedSteps:  ctx.StringSlice(buildFailedStepsFlag),
		Finished:     time.Unix(ctx.Int64(buildFinishedFlag), 0),
		Number:       ctx.Int(buildNumberFlag),
		Parent:       ctx.Int(buildParentFlag),
		PullRequest:  ctx.Int(buildPullRequestFlag),
		SourceBranch: ctx.String(buildSourceBranchFlag),
		Started:      time.Unix(ctx.Int64(buildStartedFlag), 0),
		Status:       ctx.String(buildStatusFlag),
		Tag:          ctx.String(buildTagFlag),
		TargetBranch: ctx.String(buildTargetBranchFlag),
	}
}

//---------------------------------------------------------------------
// Repo Flags
//---------------------------------------------------------------------

const (
	repoDefaultBranchFlag = "repo.branch"
	repoFullNameFlag      = "repo.full-name"
	repoLinkFlag          = "repo.link"
	repoNameFlag          = "repo.name"
	repoOwnerFlag         = "repo.owner"
	repoPrivateFlag       = "repo.private"
	repoRemoteURLFlag     = "repo.remote-url"
	repoSCMFlag           = "repo.scm"
	repoVisibilityFlag    = "repo.visibility"
)

// repoFlags has the cli.Flags for the plugin.Repo
func repoFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   repoDefaultBranchFlag,
			Usage:  "repo default branch",
			EnvVar: environ.RepoDefaultBranchEnvVar,
		},
		cli.StringFlag{
			Name:   repoFullNameFlag,
			Usage:  "repo full name",
			EnvVar: environ.RepoFullNameEnvVar,
		},
		cli.StringFlag{
			Name:   repoLinkFlag,
			Usage:  "repo link",
			EnvVar: environ.RepoLinkEnvVar,
		},
		cli.StringFlag{
			Name:   repoNameFlag,
			Usage:  "repo name",
			EnvVar: environ.RepoNameEnvVar,
		},
		cli.StringFlag{
			Name:   repoOwnerFlag,
			Usage:  "repo owner",
			EnvVar: environ.RepoOwnerEnvVar,
		},
		cli.BoolFlag{
			Name:   repoPrivateFlag,
			Usage:  "repo private",
			EnvVar: environ.RepoPrivateEnvVar,
		},
		cli.StringFlag{
			Name:   repoRemoteURLFlag,
			Usage:  "repo remote url",
			EnvVar: environ.RepoRemoteURLEnvVar,
		},
		cli.StringFlag{
			Name:   repoSCMFlag,
			Usage:  "repo scm",
			EnvVar: environ.RepoSCMEnvVar,
		},
		cli.StringFlag{
			Name:   repoVisibilityFlag,
			Usage:  "repo visibility",
			EnvVar: environ.RepoVisibilityEnvVar,
		},
	}
}

// repoFromContext creates a plugin.Repo from the cli.Context.
func repoFromContext(ctx *cli.Context) plugin.Repo {
	return plugin.Repo{
		DefaultBranch: ctx.String(repoDefaultBranchFlag),
		FullName:      ctx.String(repoFullNameFlag),
		Link:          ctx.String(repoLinkFlag),
		Name:          ctx.String(repoNameFlag),
		Owner:         ctx.String(repoOwnerFlag),
		Private:       ctx.Bool(repoPrivateFlag),
		RemoteURL:     ctx.String(repoRemoteURLFlag),
		SCM:           ctx.String(repoSCMFlag),
		Visibility:    ctx.String(repoVisibilityFlag),
	}
}

//---------------------------------------------------------------------
// Commit Flags
//---------------------------------------------------------------------

const (
	commitAfterFlag        = "commit.after"
	commitAuthorFlag       = "commit.author"
	commitAuthorAvatarFlag = "commit.author-avatar"
	commitAuthorEmailFlag  = "commit.author-email"
	commitAuthorNameFlag   = "commit.author-name"
	commitBeforeFlag       = "commit.before"
	commitBranchFlag       = "commit.branch"
	commitLinkFlag         = "commit.link"
	commitMessageFlag      = "commit.message"
	commitRefFlag          = "commit.ref"
	commitSHAFlag          = "commit.sha"
)

// commitFlags has the cli.Flags for the plugin.Commit.
func commitFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   commitAfterFlag,
			Usage:  "commit after",
			EnvVar: environ.CommitAfterEnvVar,
		},
		cli.StringFlag{
			Name:   commitAuthorFlag,
			Usage:  "commit author",
			EnvVar: environ.CommitAuthorEnvVar,
		},
		cli.StringFlag{
			Name:   commitAuthorAvatarFlag,
			Usage:  "commit author avatar",
			EnvVar: environ.CommitAuthorAvatarEnvVar,
		},
		cli.StringFlag{
			Name:   commitAuthorEmailFlag,
			Usage:  "commit author email",
			EnvVar: environ.CommitAuthorEmailEnvVar,
		},
		cli.StringFlag{
			Name:   commitAuthorNameFlag,
			Usage:  "commit author name",
			EnvVar: environ.CommitAuthorNameEnvVar,
		},
		cli.StringFlag{
			Name:   commitBeforeFlag,
			Usage:  "commit before",
			EnvVar: environ.CommitBeforeEnvVar,
		},
		cli.StringFlag{
			Name:   commitBranchFlag,
			Usage:  "commit branch",
			EnvVar: environ.CommitBranchEnvVar,
		},
		cli.StringFlag{
			Name:   commitLinkFlag,
			Usage:  "commit link",
			EnvVar: environ.CommitLinkEnvVar,
		},
		cli.StringFlag{
			Name:   commitMessageFlag,
			Usage:  "commit message",
			EnvVar: environ.CommitMessageEnvVar,
		},
		cli.StringFlag{
			Name:   commitRefFlag,
			Usage:  "commit ref",
			EnvVar: environ.CommitRefEnvVar,
		},
		cli.StringFlag{
			Name:   commitSHAFlag,
			Usage:  "commit sha",
			EnvVar: environ.CommitSHAEnvVar,
		},
	}
}

// commitFromContext creates a plugin.Commit from the cli.Context.
func commitFromContext(ctx *cli.Context) plugin.Commit {
	return plugin.Commit{
		After:        ctx.String(commitAfterFlag),
		Author:       ctx.String(commitAuthorFlag),
		AuthorAvatar: ctx.String(commitAuthorAvatarFlag),
		AuthorEmail:  ctx.String(commitAuthorEmailFlag),
		AuthorName:   ctx.String(commitAuthorNameFlag),
		Before:       ctx.String(commitBeforeFlag),
		Branch:       ctx.String(commitBranchFlag),
		Link:         ctx.String(commitLinkFlag),
		Message:      ctx.String(commitMessageFlag),
		Ref:          ctx.String(commitRefFlag),
		SHA:          ctx.String(commitSHAFlag),
	}
}

//---------------------------------------------------------------------
// Stage Flags
//---------------------------------------------------------------------

const (
	stageArchFlag      = "stage.arch"
	stageDependsOnFlag = "stage.depends-on"
	stageFinishedFlag  = "stage.finished"
	stageKindFlag      = "stage.kind"
	stageMachineFlag   = "stage.machine"
	stageNameFlag      = "stage.name"
	stageNumberFlag    = "stage.number"
	stageOSFlag        = "stage.os"
	stageStartedFlag   = "stage.started"
	stageStatusFlag    = "stage.status"
	stageTypeFlag      = "stage.type"
	stageVariantFlag   = "stage.variant"
	stageVersionFlag   = "stage.version"
)

// stageFlags has the cli.Flags for the plugin.Stage
func stageFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   stageArchFlag,
			Usage:  "stage arch",
			EnvVar: environ.StageArchEnvVar,
		},
		cli.StringSliceFlag{
			Name:   stageDependsOnFlag,
			Usage:  "stage depends on",
			EnvVar: environ.StageDependsOnEnvVar,
		},
		cli.Int64Flag{
			Name:   stageFinishedFlag,
			Usage:  "stage finished",
			EnvVar: environ.StageFinishedEnvVar,
		}, cli.StringFlag{
			Name:   stageKindFlag,
			Usage:  "stage kind",
			EnvVar: environ.StageKindEnvVar,
		}, cli.StringFlag{
			Name:   stageMachineFlag,
			Usage:  "stage machine",
			EnvVar: environ.StageMachineEnvVar,
		}, cli.StringFlag{
			Name:   stageNameFlag,
			Usage:  "stage name",
			EnvVar: environ.StageNameEnvVar,
		}, cli.IntFlag{
			Name:   stageNumberFlag,
			Usage:  "stage number",
			EnvVar: environ.StageNumberEnvVar,
		}, cli.StringFlag{
			Name:   stageOSFlag,
			Usage:  "stage os",
			EnvVar: environ.StageOSEnvVar,
		}, cli.Int64Flag{
			Name:   stageStartedFlag,
			Usage:  "stage started",
			EnvVar: environ.StageStartedEnvVar,
		}, cli.StringFlag{
			Name:   stageStatusFlag,
			Usage:  "stage status",
			EnvVar: environ.StageStatusEnvVar,
		}, cli.StringFlag{
			Name:   stageTypeFlag,
			Usage:  "stage type",
			EnvVar: environ.StageTypeEnvVar,
		}, cli.StringFlag{
			Name:   stageVariantFlag,
			Usage:  "stage variant",
			EnvVar: environ.StageVariantEnvVar,
		}, cli.StringFlag{
			Name:   stageVersionFlag,
			Usage:  "stage version",
			EnvVar: environ.StageVersionEnvVar,
		},
	}
}

// stageFromContext creates a plugin.Stage from the cli.Context.
func stageFromContext(ctx *cli.Context) plugin.Stage {
	return plugin.Stage{
		Arch:      ctx.String(stageArchFlag),
		DependsOn: ctx.StringSlice(stageDependsOnFlag),
		Finished:  time.Unix(ctx.Int64(stageFinishedFlag), 0),
		Kind:      ctx.String(stageKindFlag),
		Machine:   ctx.String(stageMachineFlag),
		Name:      ctx.String(stageNameFlag),
		Number:    ctx.Int(stageNumberFlag),
		OS:        ctx.String(stageOSFlag),
		Started:   time.Unix(ctx.Int64(stageStartedFlag), 0),
		Status:    ctx.String(stageStatusFlag),
		Type:      ctx.String(stageTypeFlag),
		Variant:   ctx.String(stageVariantFlag),
		Version:   ctx.String(stageVersionFlag),
	}
}

//---------------------------------------------------------------------
// Step Flags
//---------------------------------------------------------------------

const (
	// stepNameFlag corresponds to plugin.Step.Name.
	stepNameFlag = "step.name"
	// stepNumberFlag corresponds to plugin.Step.Number.
	stepNumberFlag = "step.number"
)

// stepFlags has the cli.Flags for the plugin.Step.
func stepFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   stepNameFlag,
			Usage:  "step name",
			EnvVar: environ.StepNameEnvVar,
		},
		cli.StringFlag{
			Name:   stepNumberFlag,
			Usage:  "step number",
			EnvVar: environ.StepNumberEnvVar,
		},
	}
}

// stepFromContext creates a plugin.Step from the cli.Context.
func stepFromContext(ctx *cli.Context) plugin.Step {
	return plugin.Step{
		Name:   ctx.String(stepNameFlag),
		Number: ctx.Int(stepNumberFlag),
	}
}

//---------------------------------------------------------------------
// SemVer Flags
//---------------------------------------------------------------------

const (
	semVerBuildFlag      = "semver.build"
	semVerErrorFlag      = "semver.error"
	semVerMajorFlag      = "semver.major"
	semVerMinorFlag      = "semver.minor"
	semVerPatchFlag      = "semver.patch"
	semVerPrereleaseFlag = "semver.prerelease"
	semVerShortFlag      = "semver.short"
	semVerVersionFlag    = "semver.version"
)

// semVerFlags has the cli.Flags for the plugin.SemVer.
func semVerFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   semVerBuildFlag,
			Usage:  "semver build",
			EnvVar: environ.SemVerBuildEnvVar,
		},
		cli.StringFlag{
			Name:   semVerErrorFlag,
			Usage:  "semver error",
			EnvVar: environ.SemVerErrorEnvVar,
		},
		cli.StringFlag{
			Name:   semVerMajorFlag,
			Usage:  "semver major",
			EnvVar: environ.SemVerMajorEnvVar,
		},
		cli.StringFlag{
			Name:   semVerMinorFlag,
			Usage:  "semver minor",
			EnvVar: environ.SemVerMinorEnvVar,
		},
		cli.StringFlag{
			Name:   semVerPatchFlag,
			Usage:  "semver patch",
			EnvVar: environ.SemVerPatchEnvVar,
		},
		cli.StringFlag{
			Name:   semVerPrereleaseFlag,
			Usage:  "semver prerelease",
			EnvVar: environ.SemVerPrereleaseEnvVar,
		},
		cli.StringFlag{
			Name:   semVerShortFlag,
			Usage:  "semver short",
			EnvVar: environ.SemVerShortEnvVar,
		},
		cli.StringFlag{
			Name:   semVerVersionFlag,
			Usage:  "semver version",
			EnvVar: environ.SemVerVersionEnvVar,
		},
	}
}

// semVerFromContext creates a plugin.Step from the cli.Context.
func semVerFromContext(ctx *cli.Context) plugin.SemVer {
	return plugin.SemVer{
		Build:      ctx.String(semVerBuildFlag),
		Error:      ctx.String(semVerErrorFlag),
		Major:      ctx.String(semVerMajorFlag),
		Minor:      ctx.String(semVerMinorFlag),
		Patch:      ctx.String(semVerPatchFlag),
		Prerelease: ctx.String(semVerPrereleaseFlag),
		Short:      ctx.String(semVerShortFlag),
		Version:    ctx.String(semVerVersionFlag),
	}
}
