// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// Plugin is an interface for a Drone plugin written in Go.
//
// This is a higly opinionated interface for what a Plugin should do. Its
// not required that a plugin author follow it.
type Plugin interface {
	// Validate checks the inputs to the Plugin and verifies that the
	// configuration is correct before executing.
	//
	// An error is returned if there are any issues with the current
	// configuration, such as missing information or files not being
	// present. A Plugin may choose to populate additional information to
	// ensure a successful execution, for example if a URL is parsed
	// successfully it can be stored off for later use.
	//
	// Validate needs to be called before Execute.
	Validate() error

	// Execute runs the plugin in the current configuration.
	//
	// An error is returned if the Plugin did not run successfully that
	// describes the runtime error.
	//
	// Execute needs to be called after Validate.
	Execute() error
}
