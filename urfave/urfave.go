// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// Flags has the cli.Flags for the Drone plugin.
func Flags() []cli.Flag {
	flags := []cli.Flag{}

	flags = append(flags, buildFlags()...)
	flags = append(flags, repoFlags()...)
	flags = append(flags, commitFlags()...)
	flags = append(flags, stageFlags()...)
	flags = append(flags, stepFlags()...)
	flags = append(flags, semVerFlags()...)
	flags = append(flags, systemFlags()...)
	flags = append(flags, networkFlags()...)
	flags = append(flags, loggingFlags()...)

	return flags
}

// FromContext creates a drone.Pipeline from the cli.Context.
func FromContext(ctx *cli.Context) drone.Pipeline {
	loggingFromContext(ctx)

	return drone.Pipeline{
		Network: networkFromContext(ctx),
		Build:   buildFromContext(ctx),
		Repo:    repoFromContext(ctx),
		Commit:  commitFromContext(ctx),
		Stage:   stageFromContext(ctx),
		Step:    stepFromContext(ctx),
		SemVer:  semVerFromContext(ctx),
		System:  systemFromContext(ctx),
	}
}
