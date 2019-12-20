// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// Step represents the currently running step within the stage.
type Step struct {
	// Name for the name of the current step.
	Name string

	// Number is the numeric value of the step.
	Number int
}
