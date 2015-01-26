// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package user allows user account lookups by name or id.

// Package user allows user account lookups
// by name or id.
package user

// UnknownUserError is returned by Lookup when a user cannot be found.

// UnknownUserError is returned by Lookup
// when a user cannot be found.
type UnknownUserError string

func (e UnknownUserError) Error() string

// UnknownUserIdError is returned by LookupId when a user cannot be found.

// UnknownUserIdError is returned by
// LookupId when a user cannot be found.
type UnknownUserIdError int

func (e UnknownUserIdError) Error() string

// User represents a user account.
//
// On posix systems Uid and Gid contain a decimal number representing uid and gid.
// On windows Uid and Gid contain security identifier (SID) in a string format. On
// Plan 9, Uid, Gid, Username, and Name will be the contents of /dev/user.

// User represents a user account.
//
// On posix systems Uid and Gid contain a
// decimal number representing uid and gid.
// On windows Uid and Gid contain security
// identifier (SID) in a string format. On
// Plan 9, Uid, Gid, Username, and Name
// will be the contents of /dev/user.
type User struct {
	Uid      string // user id
	Gid      string // primary group id
	Username string
	Name     string
	HomeDir  string
}

// Current returns the current user.

// Current returns the current user.
func Current() (*User, error)

// Lookup looks up a user by username. If the user cannot be found, the returned
// error is of type UnknownUserError.

// Lookup looks up a user by username. If
// the user cannot be found, the returned
// error is of type UnknownUserError.
func Lookup(username string) (*User, error)

// LookupId looks up a user by userid. If the user cannot be found, the returned
// error is of type UnknownUserIdError.

// LookupId looks up a user by userid. If
// the user cannot be found, the returned
// error is of type UnknownUserIdError.
func LookupId(uid string) (*User, error)