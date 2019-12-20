// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"time"

	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// buildFlags has the cli.Flags for the drone.Build.
func buildFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "build.branch",
			Usage:   "build branch",
			EnvVars: []string{"DRONE_BRANCH"},
		},
		&cli.IntFlag{
			Name:    "build.pull-request",
			Usage:   "build pull request",
			EnvVars: []string{"DRONE_PULL_REQUEST"},
		},
		&cli.StringFlag{
			Name:    "build.tag",
			Usage:   "build tag",
			EnvVars: []string{"DRONE_TAG"},
		},
		&cli.StringFlag{
			Name:    "build.source-branch",
			Usage:   "build source branch",
			EnvVars: []string{"DRONE_SOURCE_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "build.target-branch",
			Usage:   "build target branch",
			EnvVars: []string{"DRONE_TARGET_BRANCH"},
		},
		&cli.IntFlag{
			Name:    "build.number",
			Usage:   "build number",
			EnvVars: []string{"DRONE_BUILD_NUMBER"},
		},
		&cli.IntFlag{
			Name:    "build.parent",
			Usage:   "build parent",
			EnvVars: []string{"DRONE_BUILD_PARENT"},
		},
		&cli.StringFlag{
			Name:    "build.event",
			Usage:   "build event",
			EnvVars: []string{"DRONE_BUILD_EVENT"},
		},
		&cli.StringFlag{
			Name:    "build.action",
			Usage:   "build action",
			EnvVars: []string{"DRONE_BUILD_ACTION"},
		},
		&cli.StringFlag{
			Name:    "build.status",
			Usage:   "build status",
			EnvVars: []string{"DRONE_BUILD_STATUS"},
		},
		&cli.StringFlag{
			Name:    "build.created",
			Usage:   "build created",
			EnvVars: []string{"DRONE_BUILD_CREATED"},
		},
		&cli.StringFlag{
			Name:    "build.started",
			Usage:   "build started",
			EnvVars: []string{"DRONE_BUILD_STARTED"},
		},
		&cli.StringFlag{
			Name:    "build.finished",
			Usage:   "build finished",
			EnvVars: []string{"DRONE_BUILD_FINISHED"},
		},
		&cli.StringFlag{
			Name:    "build.deploy-to",
			Usage:   "build deploy to",
			EnvVars: []string{"DRONE_DEPLOY_TO"},
		},
		&cli.StringSliceFlag{
			Name:    "build.failed-stages",
			Usage:   "build failed stages",
			EnvVars: []string{"DRONE_FAILED_STAGES"},
		},
		&cli.StringSliceFlag{
			Name:    "build.failed-steps",
			Usage:   "build failed steps",
			EnvVars: []string{"DRONE_FAILED_STEPS"},
		},
	}
}

// buildFromContext creates a drone.Build from the cli.Context.
func buildFromContext(ctx *cli.Context) drone.Build {
	return drone.Build{
		Branch:       ctx.String("build.branch"),
		PullRequest:  ctx.Int("build.pull-request"),
		Tag:          ctx.String("build.tag"),
		SourceBranch: ctx.String("build.source-branch"),
		TargetBranch: ctx.String("build.target-branch"),
		Number:       ctx.Int("build.number"),
		Parent:       ctx.Int("build.parent"),
		Event:        ctx.String("build.event"),
		Action:       ctx.String("build.action"),
		Status:       ctx.String("build.status"),
		Created:      time.Unix(ctx.Int64("build.created"), 0),
		Started:      time.Unix(ctx.Int64("build.started"), 0),
		Finished:     time.Unix(ctx.Int64("build.finished"), 0),
		DeployTo:     ctx.String("build.deploy-to"),
		FailedStages: ctx.StringSlice("build.failed-stages"),
		FailedSteps:  ctx.StringSlice("build.failed-steps"),
	}
}
