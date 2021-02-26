package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsc04fc1a61692547f914f6072065498f096af1f01 = "console.log('hello assets')"
var _Assetsbfa8d115ce0617d89507412d5393a462f8e9b003 = "<!doctype html>\n<body>\n<p>Can you see this? → {{.Bar}}</p>\n</body>"
var _Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2 = "<!doctype html>\n<body>\n<p>Hello, {{.Foo}}</p>\n<script type=\"application/javascript\" src=\"app.js\"></script>\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"bar.tmpl", "index.tmpl", "app.js"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1614265490, 1614265490812318878),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1614265488, 1614265488032612105),
		Data:     nil,
	}, "/html/bar.tmpl": &assets.File{
		Path:     "/html/bar.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1614264056, 1614264056931186864),
		Data:     []byte(_Assetsbfa8d115ce0617d89507412d5393a462f8e9b003),
	}, "/html/index.tmpl": &assets.File{
		Path:     "/html/index.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1614265488, 1614265488032519790),
		Data:     []byte(_Assets3737a75b5254ed1f6d588b40a3449721f9ea86c2),
	}, "/html/app.js": &assets.File{
		Path:     "/html/app.js",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1614264863, 1614264863653762577),
		Data:     []byte(_Assetsc04fc1a61692547f914f6072065498f096af1f01),
	}}, "")
