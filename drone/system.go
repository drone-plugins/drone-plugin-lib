// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// System represents the available system variables.
type System struct {
	// Proto for the system protocol.
	Proto string

	// Host for the system host name.
	Host string

	// Version for the system version.
	Version string
}
