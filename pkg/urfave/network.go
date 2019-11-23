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

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

//---------------------------------------------------------------------
// Transport Flags
//---------------------------------------------------------------------

// Network contains options for connecting to the network.
type Network struct {
	// Context for making network requests.
	//
	// If `trace` logging is requested the context will use `httptrace` to
	// capture all network requests.
	Context context.Context

	// Client for making network requests.
	Client *http.Client
}

const networkSkipVerifyFlag = "transport.skip-verify"

// networkFlags has the cli.Flags for the Transport.
func networkFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    networkSkipVerifyFlag,
			Usage:   "skip ssl verify",
			EnvVars: []string{"PLUGIN_SKIP_VERIFY"},
		},
	}
}

// NetworkFromContext creates a Transport from the cli.Context.
func NetworkFromContext(ctx *cli.Context) Network {
	// Create the client
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if ctx.Bool(networkSkipVerifyFlag) {
		logrus.Warning("ssl verification is turned off")
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// Create the context
	context := context.Background()

	if ctx.String(logLevelFlag) == logrus.TraceLevel.String() {
		context = traceHTTP(context)
	}

	return Network{
		Client: &http.Client{
			Transport: transport,
		},
		Context: context,
	}
}
