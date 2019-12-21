// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// repoFlags has the cli.Flags for the drone.Repo
func repoFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "repo.slug",
			Usage: "repo slug",
			EnvVars: []string{
				"DRONE_REPO",
			},
		},
		&cli.StringFlag{
			Name:  "repo.scm",
			Usage: "repo scm",
			EnvVars: []string{
				"DRONE_REPO_SCM",
			},
		},
		&cli.StringFlag{
			Name:  "repo.owner",
			Usage: "repo owner",
			EnvVars: []string{
				"DRONE_REPO_OWNER",
				"DRONE_REPO_NAMESPACE",
			},
		},
		&cli.StringFlag{
			Name:  "repo.name",
			Usage: "repo name",
			EnvVars: []string{
				"DRONE_REPO_NAME",
			},
		},
		&cli.StringFlag{
			Name:  "repo.link",
			Usage: "repo link",
			EnvVars: []string{
				"DRONE_REPO_LINK",
			},
		},
		&cli.StringFlag{
			Name:  "repo.branch",
			Usage: "repo branch",
			EnvVars: []string{
				"DRONE_REPO_BRANCH",
			},
		},
		&cli.StringFlag{
			Name:  "repo.http-url",
			Usage: "repo http url",
			EnvVars: []string{
				"DRONE_REMOTE_URL",
				"DRONE_GIT_HTTP_URL",
			},
		},
		&cli.StringFlag{
			Name:  "repo.ssh-url",
			Usage: "repo ssh url",
			EnvVars: []string{
				"DRONE_GIT_SSH_URL",
			},
		},
		&cli.StringFlag{
			Name:  "repo.visibility",
			Usage: "repo visibility",
			EnvVars: []string{
				"DRONE_REPO_VISIBILITY",
			},
		},
		&cli.BoolFlag{
			Name:  "repo.private",
			Usage: "repo private",
			EnvVars: []string{
				"DRONE_REPO_PRIVATE",
			},
		},
	}
}

// repoFromContext creates a drone.Repo from the cli.Context.
func repoFromContext(ctx *cli.Context) drone.Repo {
	return drone.Repo{
		Slug:       ctx.String("repo.slug"),
		SCM:        ctx.String("repo.scm"),
		Owner:      ctx.String("repo.owner"),
		Name:       ctx.String("repo.name"),
		Link:       ctx.String("repo.link"),
		Branch:     ctx.String("repo.branch"),
		HTTPURL:    ctx.String("repo.http-url"),
		SSHURL:     ctx.String("repo.ssh-url"),
		Visibility: ctx.String("repo.visibility"),
		Private:    ctx.Bool("repo.private"),
	}
}
