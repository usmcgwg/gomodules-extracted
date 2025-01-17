// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modconv

import "github.com/dependabot/gomodules-extracted/cmd/go/_internal_/modfile"

var Converters = map[string]func(string, []byte) (*modfile.File, error){
	"GLOCKFILE":		ParseGLOCKFILE,
	"Godeps/Godeps.json":	ParseGodepsJSON,
	"Gopkg.lock":		ParseGopkgLock,
	"dependencies.tsv":	ParseDependenciesTSV,
	"glide.lock":		ParseGlideLock,
	"vendor.conf":		ParseVendorConf,
	"vendor.yml":		ParseVendorYML,
	"vendor/manifest":	ParseVendorManifest,
	"vendor/vendor.json":	ParseVendorJSON,
}
