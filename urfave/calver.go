// Copyright (c) 2020, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// calVerFlags has the cli.Flags for the drone.CalVer.
func calVerFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "calver.version",
			Usage: "calver version",
			EnvVars: []string{
				"DRONE_CALVER",
			},
		},
		&cli.StringFlag{
			Name:  "calver.major",
			Usage: "calver major",
			EnvVars: []string{
				"DRONE_CALVER_MAJOR",
			},
		},
		&cli.StringFlag{
			Name:  "calver.minor",
			Usage: "calver minor",
			EnvVars: []string{
				"DRONE_CALVER_MINOR",
			},
		},
		&cli.StringFlag{
			Name:  "calver.micro",
			Usage: "calver micro",
			EnvVars: []string{
				"DRONE_CALVER_MICRO",
			},
		},
		&cli.StringFlag{
			Name:  "calver.modifier",
			Usage: "calver modifier",
			EnvVars: []string{
				"DRONE_CALVER_MODIFIER",
			},
		},
		&cli.StringFlag{
			Name:  "calver.short",
			Usage: "calver short",
			EnvVars: []string{
				"DRONE_CALVER_SHORT",
			},
		},
	}
}

// calVerFromContext creates a drone.CalVer from the cli.Context.
func calVerFromContext(ctx *cli.Context) drone.CalVer {
	return drone.CalVer{
		Version:  ctx.String("calver.version"),
		Major:    ctx.String("calver.major"),
		Minor:    ctx.String("calver.minor"),
		Micro:    ctx.String("calver.micro"),
		Modifier: ctx.String("calver.modifier"),
		Short:    ctx.String("calver.short"),
	}
}
