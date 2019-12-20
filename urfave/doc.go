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
// 	import (
// 		"github.com/drone-plugins/drone-plugin-lib/urfave"
// 		"github.com/urfave/cli/v2"
// 	)
//
// 	func main() {
// 		app := cli.NewApp()
// 		app.Name = "plugin name"
// 		app.Action = run
// 		app.Flags = []cli.Flag{
// 			// All my plugin flags
// 		}
//
// 		app.Flags = append(
// 			app.Flags,
// 			urfave.Flags()...,
// 		)
// 	}
//
// 	func run(ctx *cli.Context) error {
// 		pipeline := urfave.FromContext(ctx)
// 		...
// 		return nil
// 	}
package urfave
