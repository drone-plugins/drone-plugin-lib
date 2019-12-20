// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// Repo represents the repository for the build.
type Repo struct {
	// Slug for the full name of a repo.
	Slug string

	// SCM for the used SCM.
	SCM string

	// Owner for the repo owner.
	Owner string

	// Name for the repo name.
	Name string

	// Link for the link to the repo.
	Link string

	// Branch for the default branch of the repo.
	Branch string

	// HTTPURL for the clone URL via HTTP.
	HTTPURL string

	// SSHURL for the clone URL via SSH
	SSHURL string

	// Visbility for the visbility of the repo.
	Visibility string

	// Private to show if the repo is private.
	Private bool
}
