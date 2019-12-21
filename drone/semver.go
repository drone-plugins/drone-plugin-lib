// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// SemVer represents the semantic version of the currently running build.
//
// This value is only applicable for tags. If the tag cannot be parsed into
// a semantic version then SemVer.Error will have the reason.
type SemVer struct {
	// Version is the full semantic version.
	Version string

	// Major version number.
	Major string

	// Minor version number.
	Minor string

	// Patch version number.
	Patch string

	// Prerelease version.
	Prerelease string

	// Build version number.
	//
	// This is signified by a + at the end of the tag.
	Build string

	// Short version of the semantic version string where labels and
	// metadata are truncated.
	Short string

	// Error is the semantic version parsing error if the tag was invalid.
	Error string
}
