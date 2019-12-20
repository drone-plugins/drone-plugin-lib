// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// systemFlags has the cli.Flags for the drone.System.
func systemFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "system.proto",
			Usage: "system proto",
			EnvVars: []string{
				"DRONE_SYSTEM_PROTO",
			},
		},
		&cli.StringFlag{
			Name:  "system.host",
			Usage: "system host",
			EnvVars: []string{
				"DRONE_SYSTEM_HOST",
				"DRONE_SYSTEM_HOSTNAME",
			},
		},
		&cli.StringFlag{
			Name:  "system.version",
			Usage: "system version",
			EnvVars: []string{
				"DRONE_SYSTEM_VERSION",
			},
		},
	}
}

// systemFromContext creates a drone.System from the cli.Context.
func systemFromContext(ctx *cli.Context) drone.System {
	return drone.System{
		Proto:   ctx.String("system.proto"),
		Host:    ctx.String("system.host"),
		Version: ctx.String("system.version"),
	}
}
