// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command blog is a web server for the Go blog that can run on App Engine or
// as a stand-alone HTTP server.
package main

const hostname = "on-golang.appspot.com" // default hostname for blog server

var config = Config{
	Hostname:     hostname,
	BaseURL:      "//" + hostname,
	GodocURL:     "//on-golang.appspot.com",
	HomeArticles: 5,  // articles to display on the home page
	FeedArticles: 10, // articles to include in Atom and JSON feeds
	PlayEnabled:  true,
	FeedTitle:    "OnGolang - Regular news about the Go programming language",
}

