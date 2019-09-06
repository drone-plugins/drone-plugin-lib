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
	"github.com/drone-plugins/drone-plugin-lib/pkg/plugin"
	"github.com/urfave/cli"
)

//---------------------------------------------------------------------
// Step Flags
//---------------------------------------------------------------------

const (
	// StepNameFlag is the flag name for setting plugin.Step.Name.
	StepNameFlag = "step.name"
	// StepNumberFlag is the flag name for setting plugin.Step.Number.
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
func StepFromContext(ctx cli.Context) plugin.Step {
	return plugin.Step{
		Name:   ctx.String(StepNameFlag),
		Number: ctx.Int(StepNumberFlag),
	}
}
