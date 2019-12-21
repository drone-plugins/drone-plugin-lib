// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

import (
	"context"
	"net/http"
)

// Network contains options for connecting to the network.
type Network struct {
	// Context for making network requests.
	//
	// If `trace` logging is requested the context will use `httptrace` to
	// capture all network requests.
	Context context.Context

	/// Whether SSL verification is skipped
	SkipVerify bool

	// Client for making network requests.
	Client *http.Client
}
