// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import "github.com/urfave/cli"

//---------------------------------------------------------------------
// Transport Flags
//---------------------------------------------------------------------

// Transport contains options for the underlying http client.
type Transport struct {
	// SSLVerify certificate information.
	SSLVerify bool
}

const transportSSLVerifyFlag = "transport.ssl-verify"

// TransportFlags has the cli.Flags for the Transport.
func TransportFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:   transportSSLVerifyFlag,
			Usage:  "transport ssl verify",
			EnvVar: "PLUGIN_SSL_VERIFY",
			Hidden: true,
		},
	}
}

// TransportFromContext creates a Transport from the cli.Context.
func TransportFromContext(ctx cli.Context) Transport {
	return Transport{
		SSLVerify: ctx.Bool(transportSSLVerifyFlag),
	}
}
