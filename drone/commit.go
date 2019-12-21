// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// Commit represents the current commit being built.
type Commit struct {
	// SHA for the current commit.
	SHA string

	// Before contains the commit sha before the patch is applied.
	Before string

	// After contains the commit sha after the patch is applied.
	After string

	// Ref for the current commit.
	Ref string

	// Branch target for the push or pull request. This may be empty for
	// tag events.
	Branch string

	// Link to the commit or object in the source control management system.
	Link string

	// Message for the current commit.
	Message string

	// Author of the commit.
	Author string

	// AuthorName of the commit.
	AuthorName string

	// AuthorEmail of the commit.
	AuthorEmail string

	// AuthorAvatar of the commit.
	AuthorAvatar string
}
