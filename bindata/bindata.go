// Code generated by go-bindata.
// sources:
// templates/404.html
// templates/footer.html
// templates/header.html
// templates/internal_error.html
// templates/main.html
// DO NOT EDIT!

package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templates404Html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x92\xc1\x6a\xc3\x30\x10\x44\xef\xfd\x8a\x41\x77\xc7\x49\xda\x4b\xc1\xf1\x27\xf4\x54\xe8\x31\x28\xd6\x2a\x2b\xa2\x48\x42\x5e\xdb\x04\xe3\x7f\x2f\x09\x69\x28\x8e\x5b\x5d\x24\x0d\x6f\x90\x76\x76\xc7\xd1\x90\x75\x81\xa0\x6c\xec\xf2\x3e\xf2\xfe\xba\xab\x69\x7a\x19\x47\xa1\x73\xf2\x5a\x08\x8a\x49\x1b\xca\x0a\xea\x23\x0a\x6c\xec\x82\xb9\x12\xb8\xaf\xca\xb8\x1e\x8d\xd7\x6d\xbb\x53\xe7\xe1\x1d\x0d\x05\xa1\x8c\xc4\xaf\x45\x68\x55\xfd\xe0\xe6\x6c\x63\x91\x78\xfb\xcc\xcc\x39\xeb\x31\x14\x9b\xf5\x1a\x49\x6f\x17\xd0\x39\x7e\x38\x16\x03\x3b\x21\x48\x83\xd4\xbf\xfd\xe1\xb8\xb9\x78\xf3\x78\x63\x0b\xcf\x85\x38\xf1\xa4\xea\x47\x95\x55\xc9\x9b\x7f\xec\xa9\xfe\x22\x0c\x94\x09\x21\x0a\xf4\xc1\x13\x24\xc2\xba\x60\x30\xb0\x16\x5c\x62\x87\x41\x07\x21\xb3\xc2\x67\xbe\xa0\x25\x9d\x1b\x76\xe1\x88\x18\x50\x69\x70\x26\xbb\x53\xa5\xfa\xf9\x84\x77\xe1\x04\xa3\xf3\xa9\x38\xf8\x8e\xc0\xb1\xa7\x7c\x3b\xaa\x5a\x98\x70\xd6\x2e\x20\xe9\x23\x55\xa5\xae\x57\xcb\x39\x94\xc6\xf5\x0b\x69\x3e\xcb\x33\xe9\x7e\xfd\xdd\x74\x1b\xa3\xd0\x7d\x14\x28\x98\x69\xfa\x0e\x00\x00\xff\xff\xc2\xe3\x3d\x77\x2c\x02\x00\x00")

func templates404HtmlBytes() ([]byte, error) {
	return bindataRead(
		_templates404Html,
		"templates/404.html",
	)
}

func templates404Html() (*asset, error) {
	bytes, err := templates404HtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/404.html", size: 556, mode: os.FileMode(436), modTime: time.Unix(1476638098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFooterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\x4a\xcb\xcf\x2f\x49\x2d\x52\xaa\xad\xe5\x52\x50\x50\x50\xb0\xd1\x4f\xca\x4f\xa9\xb4\xe3\xb2\xd1\xcf\x28\xc9\xcd\xb1\xe3\xaa\xae\x4e\xcd\x4b\xa9\xad\x05\x04\x00\x00\xff\xff\xfc\x8e\x72\x53\x2f\x00\x00\x00")

func templatesFooterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFooterHtml,
		"templates/footer.html",
	)
}

func templatesFooterHtml() (*asset, error) {
	bytes, err := templatesFooterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/footer.html", size: 47, mode: os.FileMode(436), modTime: time.Unix(1476625875, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x50\xc1\x6e\xa3\x30\x14\xbc\xe7\x2b\xde\xfa\xbc\x60\xad\xd8\xec\x61\x85\x51\xbf\xa1\xfd\x02\x63\x3f\xe2\xa7\x18\x1b\xd9\x2f\x44\x88\xf2\xef\x55\xa0\x49\xe8\xa1\x52\x4f\x30\xa3\xf1\xcc\xbc\x99\x67\x8b\x1d\x05\x04\xe1\x50\x5b\x4c\x62\x59\x0e\xf5\x2f\x1b\x0d\x4f\x03\x82\xe3\xde\x37\x87\x7a\xfb\x00\x00\xd4\x37\xd5\xf6\xbb\x42\x26\xf6\xd8\xcc\x73\xb9\x2c\xf0\x0e\x6f\xfa\x84\xb5\xdc\xb8\xa7\xa6\x47\xd6\x10\x74\x8f\x4a\x8c\x84\xd7\x21\x26\x16\x60\x62\x60\x0c\xac\xc4\x95\x2c\x3b\x65\x71\x24\x83\xc5\x0a\x7e\x03\x05\x62\xd2\xbe\xc8\x46\x7b\x54\x7f\xc4\xce\xcc\x53\x38\x43\x42\xaf\x44\xe6\xc9\x63\x76\x88\x2c\xc0\x25\xec\x94\x70\xcc\x43\xfe\x2f\xe5\x25\x0c\xe7\x53\x69\x62\x2f\x59\x1b\x37\xc5\x90\x5f\xfe\x96\xc7\xb2\x92\x26\xe7\x07\x55\xf6\x14\x4a\x93\xb3\x90\x9f\x97\xc9\xe7\x69\x75\x1b\xed\xb4\x0b\x0d\x7a\x04\xe3\x75\xce\x4a\x0c\xba\x82\x41\x57\x45\xc8\xd0\x9e\x8a\xd6\x5f\x70\xd7\x6e\x15\xeb\xbb\x74\xad\x6a\xa9\x87\xab\x23\x46\x80\x16\xba\x7f\xd0\x1d\x6f\x4f\x2d\xb5\xd0\xa7\xea\x5e\x5c\x0a\x58\x47\x53\xe2\xb6\xa0\x68\xb6\x1d\xf5\x4f\x8d\xe1\x7b\xe3\x84\x43\xcc\xc4\x31\x11\xe6\x47\xc8\xeb\x9e\x6c\xf6\xe8\x4b\x68\x2d\x83\x1e\x9b\xc3\x3c\x63\xb0\xcb\xf2\x11\x00\x00\xff\xff\xda\x3d\xd1\x88\x29\x02\x00\x00")

func templatesHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeaderHtml,
		"templates/header.html",
	)
}

func templatesHeaderHtml() (*asset, error) {
	bytes, err := templatesHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/header.html", size: 553, mode: os.FileMode(436), modTime: time.Unix(1476625876, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInternal_errorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x53\xc1\x6e\xdc\x20\x10\xbd\xf7\x2b\xa6\xdc\x6d\xb2\xdb\x5e\x1a\x11\x4b\x3d\xf4\x5c\xa9\xed\xbd\xc2\xf0\x6c\xd0\x62\x40\x80\xbd\xaa\x50\xfe\xbd\x8a\xb5\x1b\x45\xde\x8d\x0f\xe1\x04\x33\xef\xcd\x3c\x3d\x66\x6a\xd5\x18\xac\x07\x31\xeb\x0b\x92\x97\xee\x2f\x52\x0a\x89\x3d\x3f\x7f\xaa\xb5\x60\x8a\x4e\x16\x10\x33\x90\x1a\x89\x11\xfb\x71\xcd\xd2\xe5\x08\x6d\x17\x52\x4e\xe6\xfc\xc4\xa6\xf3\x37\x52\x78\x29\x44\xd1\x7c\x69\x7c\x66\xdd\x2b\x6e\x8b\x55\x03\x45\x73\xbc\xc5\x6c\x71\x83\xa3\x73\x73\x78\x78\xa0\x28\x8f\x77\xa0\x5b\x78\x3f\x36\x67\x63\x0b\xa8\x28\x8a\xcb\xd7\x77\x18\x2b\xcb\x1c\x5e\x7b\x1c\xc9\x99\xa6\xd8\xe2\xc0\xba\x9f\x21\xe6\xcf\x82\x9b\xc3\xfb\xd4\x5a\xed\x40\xed\x6a\xc5\x1b\x27\x6e\x1a\xc4\xee\xbb\xa7\xd5\x4e\x32\x32\x53\x50\x6a\x4e\x09\xfa\x51\xf0\xb8\x23\x2b\x26\x74\x42\x05\x8d\xae\xd6\x6b\x0f\xc1\xd7\x80\xe0\x2f\xc9\x1d\x59\x70\x19\x1f\x50\xd4\xee\x2a\xaa\x15\x5e\xef\x57\xfd\x63\x6c\x26\x9b\x69\x0a\xb9\x90\xb3\x27\xb8\x7f\x24\xa9\x9f\x47\xb2\x9e\x7e\xcb\x11\x2d\xfd\x42\x0c\xa9\x50\x31\x20\x9b\xf3\x0c\x12\x92\x4c\xc2\xf0\xc4\x4c\x29\x31\x3f\x72\x3e\xda\x62\xe6\xbe\x55\x61\xe2\x93\xf5\x50\x49\x0e\x05\x89\x67\x39\x82\xaf\x94\xcc\xae\x1f\xe6\xac\x3f\x91\x96\xe9\xd4\xf4\x6e\x06\x99\xb0\x20\xad\x57\xd6\x19\x24\x08\x2e\xbb\xf6\xfe\xa8\x70\x6d\x97\x3b\x03\x77\x1b\xde\x84\x2e\xcf\xb7\x3b\x31\x84\x50\x70\xd9\x94\xd5\xa1\xff\x01\x00\x00\xff\xff\xb7\x65\x6c\xfc\x4d\x03\x00\x00")

func templatesInternal_errorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesInternal_errorHtml,
		"templates/internal_error.html",
	)
}

func templatesInternal_errorHtml() (*asset, error) {
	bytes, err := templatesInternal_errorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/internal_error.html", size: 845, mode: os.FileMode(436), modTime: time.Unix(1476627982, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMainHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\xd1\xc1\x4e\xc3\x30\x0c\x06\xe0\x3b\x4f\x61\xe5\x1e\xba\x16\x38\x20\x95\xbe\x01\x27\x78\x01\xaf\x75\x1b\x8b\x2c\xb5\x62\xab\x3d\x54\x7d\x77\x34\xd8\x26\xd4\x6d\xe4\x96\x5f\xdf\x6f\x29\xf1\xb2\x74\xd4\x73\x22\x70\x07\xe4\xe4\xd6\xf5\x61\x59\x8c\x0e\x12\xd1\x08\x5c\x20\xec\x28\x3b\x70\xef\xc8\x09\x04\x07\x3a\x0a\x38\x9d\xba\xe3\x09\xda\x88\xaa\x6f\xee\x30\x97\x15\xb4\x94\x8c\x32\x48\x78\xf2\x49\x61\x3f\xf8\x44\x98\xfd\x1c\xd8\xc8\x35\x97\xda\xb6\xda\xf6\x20\xa1\xf2\x49\x37\x66\xeb\xfa\x08\xb3\x2f\x77\x3b\x10\xac\x6e\xd0\x2d\xb7\x16\x64\x7a\xbe\x03\x7f\x70\x28\x2f\xa3\x2b\x88\xc1\x1b\x5b\x24\xd7\x7c\xe0\x40\x75\x11\xca\x7f\x9a\xd2\x7c\x06\x56\xe0\xa4\x86\x31\x02\x2b\x28\xe5\x89\xd3\x00\xb5\x0a\xa6\xf3\xd8\xbd\x6b\xca\x5d\x5d\x1c\xa3\x06\x32\xc9\xa8\x6c\x63\x66\x52\x98\xd9\xc2\xb5\x7d\x79\x3d\x63\xc1\xf6\x0b\x07\xd2\xc7\xba\x90\x3b\x4f\x2d\x3a\x9e\x6e\x7c\xd8\x75\xbc\x89\x4e\xd7\xbf\x6b\xee\xc7\xd1\x28\xff\x2e\x9f\x52\xb7\xae\xdf\x01\x00\x00\xff\xff\x5c\xcb\xfc\x8e\x16\x02\x00\x00")

func templatesMainHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesMainHtml,
		"templates/main.html",
	)
}

func templatesMainHtml() (*asset, error) {
	bytes, err := templatesMainHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/main.html", size: 534, mode: os.FileMode(436), modTime: time.Unix(1476638562, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/404.html": templates404Html,
	"templates/footer.html": templatesFooterHtml,
	"templates/header.html": templatesHeaderHtml,
	"templates/internal_error.html": templatesInternal_errorHtml,
	"templates/main.html": templatesMainHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"404.html": &bintree{templates404Html, map[string]*bintree{}},
		"footer.html": &bintree{templatesFooterHtml, map[string]*bintree{}},
		"header.html": &bintree{templatesHeaderHtml, map[string]*bintree{}},
		"internal_error.html": &bintree{templatesInternal_errorHtml, map[string]*bintree{}},
		"main.html": &bintree{templatesMainHtml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
