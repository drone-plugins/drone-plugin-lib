// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/drone-plugins/drone-plugin-lib/drone"
	"github.com/drone-plugins/drone-plugin-lib/trace"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// networkFlags has the cli.Flags for the drone.Network.
func networkFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "transport.skip-verify",
			Usage:   "skip ssl verify",
			EnvVars: []string{"PLUGIN_SKIP_VERIFY"},
		},
	}
}

// NetworkFromContext creates a drone.Network from the cli.Context.
func NetworkFromContext(c *cli.Context) drone.Network {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	ctx := context.Background()
	skipVerify := c.Bool("transport.skip-verify")

	if skipVerify {
		logrus.Warning("ssl verification is turned off")

		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	if c.String("log-level") == logrus.TraceLevel.String() {
		ctx = trace.HTTP(ctx)
	}

	client := &http.Client{
		Transport: transport,
	}

	return drone.Network{
		Context:    ctx,
		SkipVerify: skipVerify,
		Client:     client,
	}
}
