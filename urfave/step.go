// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// stepFlags has the cli.Flags for the drone.Step.
func stepFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "step.name",
			Usage: "step name",
			EnvVars: []string{
				"DRONE_STEP_NAME",
			},
		},
		&cli.IntFlag{
			Name:  "step.number",
			Usage: "step number",
			EnvVars: []string{
				"DRONE_STEP_NUMBER",
			},
		},
	}
}

// stepFromContext creates a drone.Step from the cli.Context.
func stepFromContext(ctx *cli.Context) drone.Step {
	return drone.Step{
		Name:   ctx.String("step.name"),
		Number: ctx.Int("step.number"),
	}
}
