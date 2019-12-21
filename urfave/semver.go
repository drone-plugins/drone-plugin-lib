// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// semVerFlags has the cli.Flags for the drone.SemVer.
func semVerFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "semver.version",
			Usage: "semver version",
			EnvVars: []string{
				"DRONE_SEMVER",
			},
		},
		&cli.StringFlag{
			Name:  "semver.major",
			Usage: "semver major",
			EnvVars: []string{
				"DRONE_SEMVER_MAJOR",
			},
		},
		&cli.StringFlag{
			Name:  "semver.minor",
			Usage: "semver minor",
			EnvVars: []string{
				"DRONE_SEMVER_MINOR",
			},
		},
		&cli.StringFlag{
			Name:  "semver.patch",
			Usage: "semver patch",
			EnvVars: []string{
				"DRONE_SEMVER_PATCH",
			},
		},
		&cli.StringFlag{
			Name:  "semver.prerelease",
			Usage: "semver prerelease",
			EnvVars: []string{
				"DRONE_SEMVER_PRERELEASE",
			},
		},
		&cli.StringFlag{
			Name:  "semver.build",
			Usage: "semver build",
			EnvVars: []string{
				"DRONE_SEMVER_BUILD",
			},
		},
		&cli.StringFlag{
			Name:  "semver.short",
			Usage: "semver short",
			EnvVars: []string{
				"DRONE_SEMVER_SHORT",
			},
		},
		&cli.StringFlag{
			Name:  "semver.error",
			Usage: "semver error",
			EnvVars: []string{
				"DRONE_SEMVER_ERROR",
			},
		},
	}
}

// semVerFromContext creates a drone.Step from the cli.Context.
func semVerFromContext(ctx *cli.Context) drone.SemVer {
	return drone.SemVer{
		Version:    ctx.String("semver.version"),
		Major:      ctx.String("semver.major"),
		Minor:      ctx.String("semver.minor"),
		Patch:      ctx.String("semver.patch"),
		Prerelease: ctx.String("semver.prerelease"),
		Build:      ctx.String("semver.build"),
		Short:      ctx.String("semver.short"),
		Error:      ctx.String("semver.error"),
	}
}
