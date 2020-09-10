// Copyright (c) 2019, the Drone Plugins project authors.
// Please see the AUTHORS file for details. All rights reserved.
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

type (
	// Commit represents the current commit being built.
	Commit struct {
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
		Message Message

		// Author of the commit.
		Author Author

		// AuthorName of the commit.
		AuthorName string

		// AuthorEmail of the commit.
		AuthorEmail string

		// AuthorAvatar of the commit.
		AuthorAvatar string
	}

	// Author of a Commit.
	Author struct {
		// Username of the Commit author.
		Username string
		// Name of the Commit author.
		Name string
		// Email for the Commit author.
		Email string
		// Avatar for the Commit author.
		Avatar string
	}

	// Message for a Commit.
	Message struct {
		// Title for the Commit.
		Title string
		// Body of the Commit message.
		Body string
	}
)

func (a Author) String() string {
	return a.Username
}

func (m Message) String() string {
	return m.Title + m.Body
}
