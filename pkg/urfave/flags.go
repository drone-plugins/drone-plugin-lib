// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"strings"

	"github.com/urfave/cli"
)

// ShowAllFlags sets Hidden to false on all of the flags.
//
// By default all Drone environment flags are hidden by default. For
// notification plugins its unlikely this is the desired behavior. This function
// modifies all the flags to be shown in the help menu.
func ShowAllFlags(flags []cli.Flag) []cli.Flag {
	for i := 0; i < len(flags); i++ {
		flags[i] = ShowFlag(flags[i])
	}

	return flags
}

// ShowFlagsByName sets Hidden to false on all flags with the given names.
//
//     flags := urfave.ShowFlagsByName(flags, urfave.RepoNameFlag, urfave.RepoOwnerFlag)
func ShowFlagsByName(flags []cli.Flag, names ...string) []cli.Flag {
	for i := 0; i < len(flags); i++ {
		flag := flags[i]
		flagName := flag.GetName()

		for _, name := range names {
			if flagName == name {
				flags[i] = ShowFlag(flag)
				break
			}
		}
	}

	return flags
}

// ShowFlagsByType sets Hidden to false on all flags with the given type names.
//
// All of the Drone environment flags are prefixed with their corresponding type
// e.g. "repo". Any flags that contain the type name are shown.
//
//     flags := urfave.ShowFlagsByType(flags, urfave.RepoFlagType)
func ShowFlagsByType(flags []cli.Flag, typenames ...string) []cli.Flag {
	for i := 0; i < len(flags); i++ {
		flag := flags[i]
		flagName := flag.GetName()

		for _, typename := range typenames {
			if strings.HasPrefix(flagName, typename) {
				flags[i] = ShowFlag(flag)
				break
			}
		}
	}

	return flags
}

// ShowFlag sets Hidden to false on the flag.
func ShowFlag(flag cli.Flag) cli.Flag {
	// Currently we're assuming that the only flag types applicable are the ones
	// used within this library when creating the Drone environment.
	switch f := flag.(type) {
	case cli.StringFlag:
		f.Hidden = false
		return f
	case cli.StringSliceFlag:
		f.Hidden = false
		return f
	case cli.IntFlag:
		f.Hidden = false
		return f
	case cli.Int64Flag:
		f.Hidden = false
		return f
	case cli.BoolFlag:
		f.Hidden = false
		return f
	}

	// This shouldn't be hit
	return flag
}
