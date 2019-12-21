// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// commitFlags has the cli.Flags for the drone.Commit.
func commitFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "commit.sha",
			Usage: "commit sha",
			EnvVars: []string{
				"DRONE_COMMIT",
				"DRONE_COMMIT_SHA",
			},
		},
		&cli.StringFlag{
			Name:  "commit.before",
			Usage: "commit before",
			EnvVars: []string{
				"DRONE_COMMIT_BEFORE",
			},
		},
		&cli.StringFlag{
			Name:  "commit.after",
			Usage: "commit after",
			EnvVars: []string{
				"DRONE_COMMIT_AFTER",
			},
		},
		&cli.StringFlag{
			Name:  "commit.ref",
			Usage: "commit ref",
			EnvVars: []string{
				"DRONE_COMMIT_REF",
			},
		},
		&cli.StringFlag{
			Name:  "commit.branch",
			Usage: "commit branch",
			EnvVars: []string{
				"DRONE_COMMIT_BRANCH",
			},
		}, &cli.StringFlag{
			Name:  "commit.link",
			Usage: "commit link",
			EnvVars: []string{
				"DRONE_COMMIT_LINK",
			},
		},
		&cli.StringFlag{
			Name:  "commit.message",
			Usage: "commit message",
			EnvVars: []string{
				"DRONE_COMMIT_MESSAGE",
			},
		},
		&cli.StringFlag{
			Name:  "commit.author",
			Usage: "commit author",
			EnvVars: []string{
				"DRONE_COMMIT_AUTHOR",
			},
		},
		&cli.StringFlag{
			Name:  "commit.author-name",
			Usage: "commit author name",
			EnvVars: []string{
				"DRONE_COMMIT_AUTHOR_NAME",
			},
		},
		&cli.StringFlag{
			Name:  "commit.author-email",
			Usage: "commit author email",
			EnvVars: []string{
				"DRONE_COMMIT_AUTHOR_EMAIL",
			},
		},
		&cli.StringFlag{
			Name:  "commit.author-avatar",
			Usage: "commit author avatar",
			EnvVars: []string{
				"DRONE_COMMIT_AUTHOR_AVATAR",
			},
		},
	}
}

// commitFromContext creates a drone.Commit from the cli.Context.
func commitFromContext(ctx *cli.Context) drone.Commit {
	return drone.Commit{
		SHA:          ctx.String("commit.sha"),
		Before:       ctx.String("commit.before"),
		After:        ctx.String("commit.after"),
		Ref:          ctx.String("commit.ref"),
		Branch:       ctx.String("commit.branch"),
		Link:         ctx.String("commit.link"),
		Message:      ctx.String("commit.message"),
		Author:       ctx.String("commit.author"),
		AuthorName:   ctx.String("commit.author-name"),
		AuthorEmail:  ctx.String("commit.author-email"),
		AuthorAvatar: ctx.String("commit.author-avatar"),
	}
}
