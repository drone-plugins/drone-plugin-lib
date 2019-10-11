// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const logLevelFlag = "log-level"

func loggingFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  logLevelFlag,
			Usage: "logging-level",
		},
	}
}

// LoggingFromContext sets the logrus logging level.
func LoggingFromContext(ctx *cli.Context) {
	lvl, err := logrus.ParseLevel(ctx.String(logLevelFlag))

	if err != nil {
		lvl = logrus.InfoLevel
	}

	logrus.SetLevel(lvl)
	logrus.WithField("level", lvl.String()).Info("setup logging")
}
