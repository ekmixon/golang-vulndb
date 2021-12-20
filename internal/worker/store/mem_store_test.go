// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.17
// +build go1.17

package store

import "testing"

func TestMemStore(t *testing.T) {
	testStore(t, NewMemStore())
}
