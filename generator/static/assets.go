// Code generated by fileb0x at "2019-08-23 21:32:47.704668 -0700 PDT m=+0.002800974" from config file "assets.toml" DO NOT EDIT.
// modification hash(2da2dd1857ab925bc22746b0da2e831a.e17f13136a3154d6a6bbde0013668fb9)

package static

import (
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

// FileAcronymsJSON is "acronyms.json"
var FileAcronymsJSON = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8a\xe6\x52\x50\x50\x72\xf4\xf5\x54\xd2\x01\x31\x5c\xfc\x82\x21\x0c\x67\x7f\xdf\x00\xd7\x10\xcf\x10\x4f\x7f\x3f\x25\xae\x58\x2e\x40\x00\x00\x00\xff\xff\xf5\x6c\x11\x85\x26\x00\x00\x00")

// FileEnumTmpl is "enum.tmpl"
var FileEnumTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xcc\x5a\x5b\x73\xdb\xb8\x15\x7e\x36\x7f\xc5\x59\x4d\xdc\x92\x5b\x85\xf6\xb6\xd3\x97\xec\xf8\x61\x27\x49\xdd\x6c\x73\x71\xd7\xde\xce\xec\x78\x32\x0a\x44\x82\x12\x2a\x12\xe0\x02\x90\x22\x55\xe5\x7f\xef\x1c\x00\xa4\x40\x8a\xd4\x65\x53\xa7\x7d\xc8\x44\x02\x0e\xce\xe5\x3b\x57\x40\xde\x6e\x9f\x43\x4a\x33\xc6\x29\x8c\xe6\x94\xa4\x54\x8e\xaa\x2a\xb8\xba\x82\x97\x22\xa5\x30\xa3\x9c\x4a\xa2\x69\x0a\xd3\x0d\x48\xca\x97\x05\x84\x33\xa6\xe7\xcb\x69\x9c\x88\xe2\x6a\x46\xf9\x75\xc2\x52\x7a\x65\xb6\x22\x3c\xf6\xea\x03\xbc\xff\xf0\x00\xaf\x5f\xbd\x79\xf8\x26\x08\x4a\x92\x2c\xc8\x8c\xc2\x76\x1b\xbb\x8f\x55\x15\x04\xac\x28\x85\xd4\x10\x06\x00\x00\xa3\xac\xd0\xa3\x20\x0a\xb6\x5b\xca\x53\x78\x8e\xfb\xbe\x4e\xe6\x7f\xa6\x99\xe0\x4e\xaf\xed\x36\x4e\x04\xcf\xd8\x2c\x7e\xcd\x97\xc5\x9b\x57\x55\x05\x4c\x01\xf1\x54\xd5\x9b\x92\x02\xc9\x19\x51\x90\x09\x09\x7a\x4e\xfb\x0e\xa1\xc6\x71\x60\x68\x7b\x76\x77\x4b\xb7\x22\x7e\xd8\x94\xa8\x37\xaa\x85\x3a\x76\x55\x44\x4e\xa8\x5c\x22\xb8\x42\xab\x70\xef\x19\x2e\xbe\x27\x05\x85\x17\x37\xd0\xe6\x6e\x6c\x34\x34\x2b\x22\x15\xee\xa7\x2c\xd1\x30\xca\x89\xd2\x22\xcb\x14\xd5\x23\xb8\x76\x44\x20\x09\x9f\x51\x78\x26\xdf\xf0\x94\xae\xc7\x78\x24\x5f\xb6\x78\xfe\x03\x17\x14\xea\x74\x61\x78\x22\x97\x0f\x86\x0b\x52\x95\xf9\x32\x59\xb4\x59\x5b\xa9\xff\x86\x8c\x49\xa5\xa1\xaa\xb6\x5b\x78\x26\x9a\x03\x6a\x39\x75\x42\x2c\xe7\x5a\xb4\x13\x00\x2c\x03\xfa\x6b\x4d\x71\x47\x54\x42\x72\x18\x4d\x46\x55\x75\x75\x05\xf7\x0b\x56\x96\x34\x05\xb3\xb9\xdd\xd2\x5c\x51\xb3\xbe\xdd\xd6\xf4\x92\x66\x6c\x4d\x53\x7b\xce\x39\x8e\x1b\x4f\xd8\x43\xd6\x5f\xd6\x25\x0d\x82\x55\x15\x9b\xd8\x40\x0d\x7c\x6e\x2f\x45\x51\x50\xae\x6b\xcb\x59\x56\xab\xf5\x8a\xaa\x44\xb2\x12\x63\xa6\x7b\xa6\xb3\xe5\x3c\x5a\x55\x01\x1c\xd0\x72\x67\xb6\xc3\xe2\xda\xc0\xe6\x29\x08\x37\xc0\x84\x26\x96\x90\x53\xb8\x6e\x20\xad\x2a\xf8\x03\x78\x10\xe3\x41\x23\xd1\xe2\xe3\xe8\x7d\xaf\xf9\x94\xfb\x22\x06\xb9\x3d\x9b\x18\xf7\x21\x03\xe3\xe0\xb6\xcf\xed\x07\x1b\xb8\x9e\xc5\x28\xbe\x0e\xa4\xbb\x7c\x39\x63\x5c\xc5\x28\x4e\x95\x24\xa1\xf1\x6b\x4e\xa6\x39\x35\xf1\x0e\x30\xd9\x4f\x92\x86\x12\x6e\xe0\xd3\x6e\x7b\x9f\x51\xf3\xa9\xaa\x3e\x19\xb9\x4d\x1e\x5d\xf4\xb0\xbd\x5b\xcc\x4c\xe2\xb4\x98\xde\x8a\xf8\xce\x16\x10\x0b\xc7\xa7\xa1\xa3\x77\x44\xcf\x87\x8e\xe2\x1e\x1e\x8d\x0c\x0c\x9a\x16\x65\x4e\x34\x85\x91\xd2\x92\xf1\x19\x95\x23\x88\x4d\x72\x5f\x5d\xc1\x1d\x91\x8a\xf6\xd4\x05\xa2\xf1\x98\x56\xa0\x05\x24\x82\xaf\xa8\xd4\x40\xc0\x32\xc0\x35\xd2\x53\x4c\x82\x6c\xc9\x93\x21\x8e\x21\x47\x5b\x2d\x83\x08\xc2\x7d\x82\x31\x50\x29\x85\x8c\x60\x1b\x5c\xb0\x0c\xd6\x63\x10\x0b\x74\x75\x8f\xf9\x26\x63\x1f\x91\xe1\xc7\xef\x91\x6a\x1b\x5c\x5c\x48\xaa\x97\x92\xe3\x31\xce\xf2\xe0\xc2\x04\x3d\x7a\x1d\xa9\x94\xa9\x32\x35\x49\x8f\x6e\xd7\xd1\x18\xb2\x42\xc7\xaf\x51\x83\x2c\x1c\x5d\x2a\x4c\x59\x2e\xd0\xe6\x15\xc9\x59\x0a\x7d\xfa\x6a\xb9\x81\xc7\x4b\xf5\x71\x34\x06\x94\x32\x76\xd6\xa9\xf8\x47\xc1\x78\x38\x14\x48\x63\x18\x8d\x61\x14\x45\x2e\x2b\x31\x39\x9e\x40\x3b\xa7\x53\xe4\xe7\xbe\xcd\x8a\x9e\x54\x78\xa0\x6b\xcc\x31\x8c\x87\x77\x44\xaa\x39\xc9\xcd\x0a\x2b\xca\x9c\x62\xe9\x51\xa6\xb7\x68\x5c\x2b\xec\x7e\x4e\x25\x14\x54\xcf\x45\x6a\x9d\x1e\xae\xe1\xdb\x7d\x25\x22\x9f\x5d\x18\x41\xf8\xf8\x71\xba\xd1\xd4\xf7\xb4\xb3\xda\x6e\x84\xeb\xf8\xde\x40\x18\x46\x91\xf5\xa3\x0d\xd2\x9f\x79\x71\x44\xad\x25\x3f\x53\xb1\x16\xcb\xd0\xf0\xb0\x3a\x44\x56\x39\xd4\x8d\xbb\xb6\x66\xdd\x6a\x88\xa2\xe0\x42\x17\xa5\x31\x00\x77\x0e\xc5\x7a\x64\xc2\x18\x09\xbf\xb9\x41\x5b\xfc\x28\xa5\x52\x9a\x10\xfd\x76\x0d\x37\xa0\x8b\xb2\xc1\xc1\xda\x5c\xb7\x81\x01\x77\xdd\xff\xfd\xad\xf3\xd6\x7d\x42\x78\x17\x0f\x5c\xe3\x54\x02\xe3\x9a\xca\x0c\xcb\xd2\x61\x24\x90\x3e\xb4\x5d\xa9\x39\xb2\xad\x3c\x18\x56\x44\x82\x97\xbc\x41\x70\xa1\x3e\x33\x9d\xcc\x61\x85\x10\xd8\x5e\x12\x62\x33\x33\x0e\x4d\x88\xaa\x29\x5f\x04\x17\x16\xc3\x1b\x58\xb9\x0d\x8b\xb1\xb7\xe1\xb0\x5d\x45\x8e\x80\xb3\x1c\x77\x0d\x30\xbd\xa9\xb0\x03\xd1\xe5\xf9\x57\xf1\x08\x82\x6d\xe7\x84\x0e\xda\xa9\x64\x2b\x2a\xed\x5e\x2f\xe6\x7d\x90\x1b\x6a\xcc\x07\x7b\xda\x4e\x20\x3d\x59\xb1\x4b\x87\xf1\x69\x91\xf1\x97\x9c\xcc\x54\x1d\x1b\x74\x2f\x55\x6e\x45\x4e\xf8\x0c\xb2\x9c\xb8\x79\x6a\xa7\x30\xa0\xbe\xc7\x22\x85\x6a\x0c\x94\xa6\x88\xef\x22\xe4\x28\xfc\x2b\x92\x47\x0e\xdc\x55\xe0\x63\x6e\xa1\xbd\x3d\xac\xeb\x2d\xd5\xda\x47\xf7\x14\x65\x6f\x29\x56\x1c\x2f\xa2\x3d\x5c\xbf\x5d\x3b\xb9\x38\xf0\x76\x05\x7b\x63\xbf\x2a\xb3\xef\xfe\x74\x55\x22\xaa\xd0\xc1\xeb\x88\x74\x64\x1c\x46\x75\xbf\xdc\x49\x1e\xf5\x54\xea\xe3\x6e\xc5\x2b\x8a\x8a\xef\x8d\xa2\xce\xbb\x6f\x85\x58\x2c\xcb\x93\x1b\x77\xcf\x98\x8f\x04\x4c\xff\x5e\x01\xfd\x75\xc9\x56\x34\xa7\x5c\xf7\x5d\x10\x6c\x7e\x5b\x6b\x87\x84\x86\xad\x16\xd4\x48\x38\xa5\xcd\x4f\x0e\xb4\xf9\x77\xa4\x7c\xec\x15\x16\xed\xb5\xfd\x7e\xb2\xff\xce\x2c\xf0\x10\x5e\xae\xa2\x33\xe7\x01\x96\xda\x7f\x5f\x6b\x22\x38\x51\xc7\x46\xb3\xee\x5c\x50\xdf\x83\x3b\xb9\x80\x6b\x27\x57\x35\x24\xb6\x29\x07\xdb\x00\xc0\x69\xcf\xb8\x0e\xd7\xd1\xf1\x18\xff\xf1\xfe\xc3\xfb\xf6\x0c\x62\x56\x3a\x0a\xfd\x53\x09\x1e\xbb\xfd\x33\x7a\x9c\xc7\xf1\xe0\x18\xe2\xb3\xf7\x87\x91\xee\x1c\x32\xa8\x5a\x43\x71\x86\x72\x2d\xae\xe1\x74\x7f\x0e\xc1\x06\xac\xea\xee\x0b\xe0\x7a\xd8\x8b\x9b\x8e\xc8\x70\x3a\x86\xdf\xa9\xe8\xfb\x76\x87\x03\x68\x5c\xe1\x47\x8c\x65\xde\x8c\x4d\x58\xa5\x8c\x51\x26\xdd\x5f\xc0\xe5\x6a\x64\xe0\x89\x02\x80\xea\xb4\x0e\xab\x9e\x6e\xe0\xf9\xe5\x87\x77\x6f\xdb\xb1\x61\x56\x3a\x0e\xd8\x90\x22\xff\xed\xb1\x81\x1c\x31\x36\xbc\x8e\x71\x52\x47\x6e\xc5\xc5\xa0\x5a\x5f\x14\x17\x46\xb5\xc6\x55\xa6\xf9\x85\xfb\xa3\xda\x29\x01\xd3\x30\x09\xbf\x24\x50\x8c\x95\xbd\x81\xf2\xf5\x22\xa5\x89\x98\xd6\x0b\x55\x73\xc1\x6d\x5e\xa9\x06\x6a\x2e\xdc\x60\x2b\x86\x25\x4f\xa9\x54\x89\x90\x74\x62\x8f\xb2\x6c\x53\x47\x1f\x60\x6b\x0e\xdc\x8b\x4b\x37\x22\xdd\x8b\xd0\x21\x31\x0f\x4c\xe7\x54\x39\x41\xa5\xa1\x67\xff\x1a\x92\xd3\x7e\x72\x83\x5d\xb3\xaa\x2a\x74\xe6\x50\xe7\x30\x43\x32\x27\x85\xc7\xce\xdd\xeb\x87\x0e\x58\x30\x15\x10\xc8\x99\xd2\x20\x32\x28\x85\x52\x6c\x9a\xd7\x43\xbb\x75\xac\xc2\x9d\x7d\x1e\x2e\x62\x07\x98\x87\x11\x3c\x7e\xdc\x8d\x3c\xba\x28\x31\x0a\x0a\xb2\xa0\x61\xbd\x3e\x86\x9c\x0e\x37\x42\x6c\x81\x89\x28\x37\xa1\x89\xa2\x41\xaa\x26\x24\x30\x3c\x2a\xf7\xd2\x63\x5f\x54\x07\xc0\x7a\x47\x4a\x03\x15\x14\xa4\x6c\x43\x6f\xc0\xb2\x39\xbd\x77\x91\x72\xc1\x74\x6a\xff\xab\x2b\x83\x37\xf5\xb1\x0c\xbf\x1c\x19\x71\xd6\x7b\xd3\x8c\xd2\x36\x07\xbc\x74\xbc\x2f\x25\xe3\x3a\x0b\x7b\x06\xc8\xf0\x32\x8d\x46\x63\x30\x4d\x76\x08\x80\xbf\x31\x9e\x2a\x1f\x82\xc9\xc2\xac\x74\x80\x40\x32\x2f\x44\x9c\x21\x22\x73\x13\xb9\x7d\xb2\xac\xdf\x9a\x67\x6c\x45\x39\x14\x54\x29\x32\x3b\x02\x0d\xf2\xf5\x81\x31\x35\xe9\x30\x32\x46\xe3\x1d\x36\x5e\x75\x42\x74\x4c\xb1\x01\xb0\x63\xdd\x92\x2f\xb8\xf8\xcc\xed\xfb\xdd\x89\x8c\xf7\xd7\x7f\xb6\x5c\x0e\x08\x44\x79\xae\xea\xb4\x4a\xe5\x17\xfa\xe6\x5e\x2c\x65\x42\xdb\xde\xc9\x96\x79\xde\xef\x22\x4b\xbd\x73\x92\xb9\xff\x0b\xa9\xa1\x24\x7a\x0e\x29\x93\x34\xd1\xf9\x06\x87\x7b\xf3\x38\xb2\x29\x8f\x45\xad\xe1\x77\xa6\x73\x9c\xca\x4f\xe1\x9e\x9a\xf5\xff\x91\x83\xee\x88\x9e\xf7\xb8\xa7\x34\xcb\x4f\xec\x1e\x94\x7d\xa6\x73\x8c\xba\x4f\xe1\x1a\xcb\xf8\x89\x1d\x73\xf0\x55\x3d\x6e\xbf\x76\x9b\x1f\x30\xe2\x7b\x4e\x16\x78\x2d\x9e\x88\xa5\x9e\x88\x6c\x32\x15\x4b\x9e\xaa\x09\xe5\xcb\x62\x72\x99\x36\xce\x35\x8f\xdd\xcd\xcb\x7a\xe3\x21\x74\x83\x79\x8b\x72\x55\xae\x24\x12\x6f\xc2\xf5\x0f\x7a\xb6\xd8\x31\x75\x92\xab\x1a\xee\x7d\x57\xff\xe1\x5f\x00\x82\xaa\x99\x33\x38\x85\xfe\xe7\x7c\x18\x8d\xdc\x04\xfc\xc6\xc4\x93\x59\xf3\x6d\xc0\x90\x6c\xc5\x5a\xaf\x3d\x07\x0d\xd8\x71\x3e\x5d\x7f\xa4\x0e\x9a\x9f\x0c\x9b\x81\xec\xd0\xef\x2d\xce\x8e\xdd\xf7\x4e\xdf\xa5\x05\xd3\x9a\xca\xdd\x81\xf4\xd4\x0e\xdc\x9c\x68\x67\xcc\xb0\x01\xcd\x81\xb6\x09\x03\x75\xc0\x3e\x01\x99\x3a\x80\xf3\x70\x6b\x92\x88\x73\xf1\x99\x4a\xf3\x84\x59\x55\x83\x18\x78\x3f\xcb\x21\xd9\x80\x1c\x8f\xaa\x5d\x76\x52\x7f\xa3\x53\x78\x7c\xd6\x03\x80\xbe\xa2\x9a\xb0\xfc\x74\x38\x3d\x96\x67\x96\x20\xdf\x82\xa7\xa8\x44\x2d\xfe\xff\x83\x4e\xe1\xdd\x57\x9f\x0f\x3f\xd8\xfd\xf2\xc3\x4f\x77\x2f\x0f\xf8\xd9\xec\x1b\xd2\x96\x97\x37\x44\x96\xc9\x24\x31\xeb\x1d\x27\x3f\x08\xcb\x73\xc0\xc1\x66\xf3\x27\xaa\x4a\xc1\xcf\x78\xb7\x71\x4c\xc3\x08\x8c\x64\x73\xe1\xb3\x26\xd4\xce\x4e\x0e\xb8\x62\x67\x44\xaf\xa3\x93\xda\xcd\xee\x7b\x57\x84\x73\x55\x50\xf5\x5d\x81\x7a\x51\xfd\xeb\xc3\xc3\xdd\x01\x50\x71\x7b\x1f\xd3\xb9\xd6\xe5\x10\xa4\x86\xe1\x00\xa2\xb8\xf7\x1b\x00\xc5\x63\xde\x53\xd8\x11\x08\x1b\x95\x4f\x42\xf0\xcf\x7f\xbc\x3e\x03\xad\x0f\xf7\xaf\xd7\x4c\x1f\xc0\xcb\x12\xec\x23\x26\xd4\x84\xae\x99\xfe\x8a\xa0\x59\x4d\x4e\x87\xcd\xd3\xfc\x24\xe0\xbe\x0b\xf6\xff\xb4\xa5\x07\x37\xf3\xf0\x71\x00\xb0\x77\xf6\xc2\xd3\x46\xab\xa8\x17\x3b\x40\x39\xe2\x56\xa7\x36\x7f\x06\xa2\xe6\xd8\xa8\xdd\x31\xf0\x0a\xfb\x61\x94\x1c\xbf\x33\x2f\x99\x4e\xb9\xe1\x9b\x66\xff\x55\x93\x8b\x46\x41\xb2\x22\x2c\x27\xd3\xdc\x0e\x44\xc6\x82\x4b\x85\xd5\xb0\xfd\x46\x3a\x00\x99\x81\xb4\x0d\x98\x29\x01\x83\xb0\x59\x1f\x74\xc3\xcb\x2e\x9e\x16\x4e\x86\xc3\x79\x30\x59\x2d\xcf\x05\xe9\x52\x0d\x0f\xa5\x55\x35\xb9\x54\x80\x8d\xe3\x45\x0d\x97\x37\xa2\x8c\x5b\xaf\x89\xeb\xd8\xbe\x9f\xe3\xa7\xc6\xcb\x51\x6b\x2e\xa9\xbb\xce\x7f\x02\x00\x00\xff\xff\x5f\x5f\xee\x76\xcb\x26\x00\x00")

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	var f webdav.File

	var rb *bytes.Reader
	var r *gzip.Reader

	rb = bytes.NewReader(FileAcronymsJSON)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "acronyms.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	rb = bytes.NewReader(FileEnumTmpl)
	r, err = gzip.NewReader(rb)
	if err != nil {
		panic(err)
	}

	err = r.Close()
	if err != nil {
		panic(err)
	}

	f, err = FS.OpenFile(CTX, "enum.tmpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, r)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}