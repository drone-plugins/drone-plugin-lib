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
// Stage Flags
//---------------------------------------------------------------------

const (
	// StageArchFlag corresponds to plugin.Stage.Arch.
	StageArchFlag = "DRONE_STAGE_ARCH"
	// StageDependsOnFlag corresponds to plugin.Stage.DependsOn.
	StageDependsOnFlag = "DRONE_STAGE_DEPENDS_ON"
	// StageFinishedFlag corresponds to plugin.Stage.Finished.
	StageFinishedFlag = "DRONE_STAGE_FINISHED"
	// StageKindFlag corresponds Stage.Kind.
	StageKindFlag = "DRONE_STAGE_KIND"
	// StageMachineFlag corresponds to plugin.Stage.Machine.
	StageMachineFlag = "DRONE_STAGE_MACHINE"
	// StageNameFlag corresponds to plugin.Stage.Name.
	StageNameFlag = "DRONE_STAGE_NAME"
	// StageNumberFlag corresponds to plugin.Stage.Number.
	StageNumberFlag = "DRONE_STAGE_NUMBER"
	// StageOSFlag corresponds to plugin.Stage.OS.
	StageOSFlag = "DRONE_STAGE_OS"
	// StageStartedFlag corresponds to plugin.Stage.Started.
	StageStartedFlag = "DRONE_STAGE_STARTED"
	// StageStatusFlag corresponds to plugin.Stage.Status.
	StageStatusFlag = "DRONE_STAGE_STATUS"
	// StageTypeFlag corresponds to plugin.Stage.Type.
	StageTypeFlag = "DRONE_STAGE_TYPE"
	// StageVariantFlag corresponds to plugin.Stage.Variant.
	StageVariantFlag = "DRONE_STAGE_VARIANT"
	// StageVersionFlag corresponds to plugin.Stage.Version.
	StageVersionFlag = "DRONE_STAGE_VERSION"
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
