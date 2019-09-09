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
// Repo Flags
//---------------------------------------------------------------------

const (
	// RepoDefaultBranchFlag corresponds to Repo.DefaultBranch.
	RepoDefaultBranchFlag = "repo.branch"
	// RepoFullNameFlag corresponds to Repo.FullName.
	RepoFullNameFlag = "repo.full-name"
	// RepoLinkFlag corresponds to Repo.Link.
	RepoLinkFlag = "repo.link"
	// RepoNameFlag corresponds to Repo.Name
	RepoNameFlag = "repo.name"
	// RepoOwnerFlag corresponds to Repo.Owner.
	RepoOwnerFlag = "repo.owner"
	// RepoPrivateFlag corresponds to Repo.Private.
	RepoPrivateFlag = "repo.private"
	// RepoRemoteURLFlag corresponds to Repo.RemoteURL.
	RepoRemoteURLFlag = "repo.remote-url"
	// RepoSCMFlag corresponds to Repo.SCM.
	RepoSCMFlag = "repo.scm"
	// RepoVisibilityFlag corresponds to Repo.Visbility.
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
	// SemVerBuildFlag corresponds to SemVer.Build.
	SemVerBuildFlag = "semver.build"
	// SemVerErrorFlag corresponds to SemVer.Error.
	SemVerErrorFlag = "semver.error"
	// SemVerMajorFlag corresponds to SemVer.Major.
	SemVerMajorFlag = "semver.major"
	// SemVerMinorFlag corresponds to SemVer.Minor.
	SemVerMinorFlag = "semver.minor"
	// SemVerPatchFlag corresponds to SemVer.Patch.
	SemVerPatchFlag = "semver.patch"
	// SemVerPrereleaseFlag corresponds to SemVer.Prerelease
	SemVerPrereleaseFlag = "semver.prerelease"
	// SemVerShortFlag corresponds to SemVer.Short.
	SemVerShortFlag = "semver.short"
	// SemVerVersionFlag corresponds to SemVer.Version
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
