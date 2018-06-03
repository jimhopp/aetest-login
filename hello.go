// Copyright 2018 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
)

func main() {
	http.HandleFunc("/", handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	if user.Current(c) == nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "sorry, not logged in")
		return
	} else {
		fmt.Fprintln(w, "Hello, world!")
		fmt.Fprintln(w, "user: %v", user.Current(c))

	}
}
