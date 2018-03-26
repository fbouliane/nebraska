// Code generated by go-bindata.
// sources:
// db/drop_all_tables.sql
// db/migrations/0001_initial.sql
// DO NOT EDIT!

package api

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

var _dbDrop_all_tablesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x41\x8e\x85\x30\x0c\x43\xf7\x9c\xa2\xf7\xe0\x30\x91\x09\x19\x88\x28\x6d\xd5\x04\x34\xdc\x7e\x24\xc4\x6c\xbe\xbe\x14\xf6\x2f\xb6\x63\xcf\xbd\xb6\xe4\x98\xb2\x24\xfd\x49\xf2\xab\xe6\x96\x5c\xb0\x27\x86\x31\x66\x19\x87\xaf\xc8\x61\xd2\x2d\x60\xd0\x5a\x56\x86\x6b\x2d\x01\xd9\xc0\x1b\x16\x09\x28\xae\x5d\xaa\x11\xf8\x85\x22\xaf\x28\x45\x72\x40\x2d\xbd\x1e\x2d\x7a\x43\x8b\x39\x0a\x47\xe9\xfe\x31\x32\x87\x1f\x6f\x45\xe9\x7d\x49\x1f\x06\xb4\xaa\x79\xed\x57\x70\x25\xa7\x14\x27\xbf\x5a\x94\xff\x06\xa3\x45\xd9\xf5\x54\x8f\x3c\x9f\x39\xe9\x19\x81\xa6\x0c\xde\xb2\x5a\x24\x3f\xc3\x31\xc1\x84\x76\x5d\xfa\x5d\x89\x8d\xc3\x5f\x00\x00\x00\xff\xff\xda\x81\x6e\xb2\xa2\x02\x00\x00")

func dbDrop_all_tablesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbDrop_all_tablesSql,
		"db/drop_all_tables.sql",
	)
}

func dbDrop_all_tablesSql() (*asset, error) {
	bytes, err := dbDrop_all_tablesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/drop_all_tables.sql", size: 674, mode: os.FileMode(420), modTime: time.Unix(1522052498, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations0001_initialSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5a\xdb\x73\xdb\xb6\xb3\x7e\x8e\xfe\x0a\xcc\xef\x3c\x28\x19\x0b\x36\x2f\xe0\x05\xee\x69\x67\x62\xa7\x89\xdd\x3a\x69\x1c\xc7\xcd\xb8\x2f\x1a\x10\x58\x4a\x8c\x29\x82\x01\x48\xd9\xf2\x5f\x7f\x06\x14\x49\x51\xb2\x24\xcb\x97\x9c\xd3\x39\x0f\xcd\xb8\xc2\x02\x8b\xfd\xf6\xfb\x16\x0b\x92\x18\xa3\xbd\x49\x32\x52\xac\x00\x74\x99\xf7\x7a\x18\xa3\x63\xa9\xe0\x8b\x4c\x53\x50\x48\xf3\x31\x4c\x58\xaf\xc7\xd2\x02\x14\x12\xac\x60\x11\xd3\x80\xb8\x54\xa0\x6a\x0b\x28\x50\x91\x4c\xe0\x4e\x66\x80\x7e\x45\xfd\xb2\xe0\xfd\x5f\x7a\x3d\xae\xc0\xac\x08\xb7\x05\x64\x3a\x91\x19\x4a\x62\x94\xc9\x02\xc1\x6d\xa2\x0b\x8d\xfe\x53\x96\x89\xc0\x52\xeb\xfc\x3f\x0b\xe3\x82\x45\x29\xa0\x02\xd8\x04\xbd\xee\xbd\x4a\x04\x32\x46\x28\x57\xc9\x84\xa9\x19\xba\x86\x19\x12\x10\xb3\x32\x2d\xaa\x81\xe1\x08\x32\x30\xdb\x1e\x4e\xc9\xeb\x37\x83\xde\xab\x8c\x4d\x00\x4d\x99\xe2\x63\xa6\x5e\x3b\xde\x9b\xca\x5f\x56\xa6\x29\xe2\x63\xe0\xd7\xe8\x75\x65\xf0\xdf\xbf\xa1\x7e\xff\x0d\x2a\xb3\xe4\x47\x09\x83\xde\xab\xb9\x6f\x31\x2c\x74\x15\x86\x2e\xd8\x24\x2f\xee\x5a\x4f\xbc\x54\x0a\xb2\x62\xd8\x8e\xb5\xab\xf6\xde\xac\xee\xbc\xd4\xa0\xf4\x93\xb6\x6e\x66\x3e\xb8\xfd\xd6\x68\x35\x04\x0d\x5c\x41\xd1\xce\xf5\xac\xfb\x73\x6b\x93\xf9\xcc\xe7\x44\x3d\xe8\xbd\x32\x09\x1a\x36\x21\xb6\x8e\x14\xc4\xa0\x20\xe3\xa0\xeb\x0c\x26\xe2\x0d\x92\x19\x12\x90\x42\x01\x88\x33\xcd\x99\x80\x2e\x68\x49\x26\xe0\xd6\x98\xd4\xb8\xd5\xeb\xde\x83\x95\xe5\x79\x9a\x70\x56\x18\x16\x3d\x9b\x17\xeb\xc0\xe9\x80\x3a\xe8\xbd\x12\xa0\xb9\x4a\xf2\xca\x5d\x01\xb7\xc5\xff\x21\x5a\x86\x19\x55\x92\x5b\x70\x06\xc8\x6c\xf6\xcd\x7d\xee\xe5\x8c\x5f\xb3\x11\x3c\x09\xa0\x62\x96\x9b\x6c\x14\xf7\x80\xa9\x06\x7e\x43\x96\x31\x9a\x82\xaa\x74\xbc\x60\xe8\x1a\x8a\x36\x46\x2d\x98\xa5\x4a\x3b\x33\xfc\x35\xa4\x56\xe9\xc2\x3a\x4e\x52\x58\xca\x96\x6d\x59\x1b\x52\xa2\x93\xbb\x8e\x5a\x2a\xab\x31\xd3\xe3\xf6\x27\x9f\x3c\x97\xe6\x1d\xda\x6d\xcd\xdf\x12\x3d\xb7\xa7\xf1\xf5\xf2\x9a\x03\x54\xe3\xb5\x26\x9f\xa6\xc2\x4a\x3d\x64\xfc\xc9\xb4\x87\x29\x64\x45\x17\xa2\xd6\xbc\x9f\x4b\x5d\x24\x99\x2e\x58\x9a\xf6\x07\xa8\xf7\x8a\x8f\x95\x9c\x18\x77\x6b\x93\xdc\x4e\xab\x6c\xf5\x98\x39\x9e\xbf\x82\x73\x06\x20\xf4\x90\x89\x49\x92\xa1\x48\xca\x14\x58\xd6\x4e\x8b\x59\xaa\x0d\x04\x89\x1e\x0a\x48\x0b\xb6\xd1\x40\x24\xda\x84\x3e\xcc\xd9\x2c\x95\x4c\x0c\x23\xc6\xaf\x65\x1c\xdf\xb3\x2f\x54\x55\xf5\x26\x50\x30\x73\x20\x0d\x75\x32\xca\x58\x51\x2a\x18\x2a\xcd\x96\xd9\xd6\xd9\xfa\xd2\x84\xbb\x65\x8a\x2d\xdb\x09\x60\x22\x4d\xb2\x6d\x26\xcf\xe1\x55\xad\xd4\xad\x9c\x6a\xd5\xbc\x7b\x11\x5d\x21\xcc\xc2\xcb\x7d\x6a\x8d\x59\x96\x41\xfa\xd3\xcf\x58\x83\x93\x4c\xa5\x5a\x6b\xff\x6f\x12\xe7\x6a\x46\x1e\x4e\x84\x69\x7a\xea\x9d\x34\x05\x7a\x55\xda\xf7\xea\xf4\x22\x53\x0d\xfe\x5b\x72\x34\x52\xb2\xcc\x9f\xd6\x4b\x3c\xf3\xb8\xeb\xa2\x6c\x5a\x3c\x59\x16\xc3\x24\x1b\xe6\x4a\x8e\x14\x68\xbd\x5e\xbd\x4b\xfc\x96\x69\xc2\x67\xc3\x32\x17\xac\x00\x3d\x84\xcc\x04\x24\xd6\xaa\x78\xcd\x34\xcd\x62\x18\x4e\xa4\x80\x5d\x27\xc8\x38\x4e\x38\x0c\xc7\xb2\x54\xbb\x6f\xae\x6d\x58\x1b\x9c\x48\x75\x82\xd4\xa3\x39\xa8\x44\x8a\x61\x92\x15\xa0\xa6\x2c\x5d\xaa\xa1\xab\x60\x6e\x98\xd2\xc2\x5b\x8f\x4f\xd8\x6d\x8b\x48\x0e\xaa\xb6\x37\x87\x2e\x8c\x40\x6d\x5a\x75\xc3\xac\xfa\x44\x5e\x02\xba\x0a\x49\x96\xc5\x2e\x9b\x5d\x99\xf1\x22\x4d\xe1\x8b\x0a\xb2\x56\xc8\x3a\x41\xb6\xe2\x79\x41\x41\x36\x62\x5b\xb8\xbd\xa7\xc7\xea\xb4\xcc\x78\xd3\x5f\x75\xe5\xd5\x15\x66\x8d\x73\x22\x16\xa0\x26\x39\x4a\x32\x28\x5e\xa4\xf0\xdd\x6f\x13\x9a\x7d\x0d\x75\xc1\x8a\xb2\x29\x18\x1a\x54\xc2\xd2\xee\xce\xee\x95\x6e\x6b\xb9\x14\x6f\xab\xd3\x09\xef\x76\x04\x9d\x89\x5b\xb6\xb3\xd2\xb3\x3f\xad\x79\x7c\x0e\x1f\x6b\x3c\x6a\x81\x0d\x7a\xaf\x52\xa6\x8b\x61\xe5\x69\x18\x4b\xd5\xc8\xea\xf1\x0b\x57\xeb\xd4\x0a\x1a\x29\x96\xdd\xdf\xe0\x8a\xd1\xda\xd8\xeb\x90\xd7\xd9\x2d\x7a\xe7\xf9\xef\x8f\x2a\xbe\x6d\x02\x56\x48\xba\x4e\x8a\x0b\x4e\x6f\xd2\xe1\x8b\x8a\xba\x52\xd9\x3a\x49\x37\xf2\xdb\xa8\xe8\xae\xc2\x5e\x77\x22\x1c\xa0\xe5\x0d\xae\x57\xf7\x7a\x4e\x76\x96\x79\xf3\xcb\x8e\x73\x56\xbc\xed\x3a\xad\x09\xfc\x21\xf1\x0e\xc7\x89\x2e\xa4\x9a\x6d\x13\xf1\x3d\x5a\x6f\x63\xd7\x8b\x0a\xea\xff\x0f\xb3\xb6\x75\xd1\x1b\x53\xb2\x23\x5f\x56\xa7\xed\x4e\x99\xd5\x99\x1b\x59\x53\xdd\xeb\x86\xd5\xd5\x7c\x0b\x51\x9a\x3b\xfd\x52\x7b\x61\x5a\x3a\xd0\x26\xd5\x6b\x46\xba\x8d\xe0\xd2\xcd\x67\x73\xb9\x9f\x5f\x31\xb7\xec\xe2\x59\xf7\x24\x05\xd3\x44\x96\xeb\x2f\xa5\xe6\x7e\xab\x94\x54\x43\x6e\x5a\xc5\x95\xc7\x05\xff\x62\xaa\x2e\x92\x67\xd6\xba\xd7\xfe\x75\x96\xeb\xa6\xd9\xac\xb6\x96\xae\x75\x02\xb6\x93\xb3\x36\x7a\x90\x8a\xb5\xdd\xd2\x16\xef\x3f\x8c\xe3\x45\x32\x4d\x8a\xad\x45\xea\x39\x59\xe7\x29\xd3\x7a\x1d\x3d\x35\x4c\x41\x19\xc7\x6b\xc6\x9e\xd6\x5e\xfc\x0b\x6a\xd0\x93\x7a\xdd\xc5\xdc\x4d\x44\xdf\x9d\xdf\x6b\x59\xb5\xc8\xf1\x83\x9c\x59\x98\x76\xea\xd5\x66\xa3\xa5\x0e\x7b\xb3\xd9\x32\x9f\xd7\x3e\xe8\x1c\x36\x4b\x45\x29\xe3\xd7\x69\xa2\xab\x3a\xf4\xfc\x67\x2b\x6b\x32\xb2\x6e\x91\x07\x53\xb3\xd4\xb3\x2c\xb6\x35\x40\x1d\x08\x2a\xf0\x31\x46\xa7\x59\x52\x18\x15\x09\x56\xb0\xea\x87\x77\xcd\x95\x17\xd8\xa4\x97\x64\x1a\x54\x55\xb1\x65\xfb\xac\xb8\xbe\xcd\xa0\x29\x4b\x4b\xd0\xe8\x75\x5f\x84\xd4\x25\x8e\xe0\x98\x3a\x36\xc1\x84\xd8\x02\x33\xc2\x62\x1c\x09\x11\xba\x01\x73\x23\xc7\xa5\xfd\x01\xea\xd7\x1a\xec\xd7\x9e\x7f\xaf\x24\x6f\xb4\xae\x97\xfc\x74\x8b\x8f\xf9\x77\x80\xe6\xe7\xc6\x00\x75\x4e\x89\x85\x7b\x77\x80\xac\x01\xea\x9f\x36\x5c\x53\x90\x4b\x55\x80\x40\x2c\x43\x55\x99\x46\xa2\x54\x49\x36\x32\xff\x3f\xef\x6a\x91\x2e\x20\xdf\x37\xdb\x78\x9e\x5b\x7b\x80\xfa\x97\xd5\x8a\x0a\x8d\x99\x46\xb9\x92\x1c\xb4\xae\x5c\x8b\xb9\x72\x41\x34\x29\x7f\x01\x7f\x4e\x27\x4c\x8d\xca\x7c\xa4\x98\x00\x81\x0a\xd9\x54\xb5\x96\x19\x75\xb9\x79\xae\x4f\xbb\x0e\xf2\x9d\xbc\xc9\x52\xc9\x84\x41\x31\x35\x37\x97\xe2\xc5\x3c\x90\x2e\x8c\xad\x3c\x98\x52\xc9\x14\x04\xd2\x25\x37\x88\xc6\x65\x9a\xce\x9e\xeb\x2a\xb4\xac\xb9\xaf\xd3\xf9\xc3\xe7\x66\xf1\x7d\x54\x3b\xe7\x72\x92\xa7\x50\x55\x5a\x73\xfc\x83\xb9\x5d\xa1\x68\xd6\x56\xb1\xfd\x86\xb8\xc7\x52\xc1\x5f\x17\xdd\xd2\xbc\xb4\xaf\x95\x92\x3d\x97\xcb\xd2\xa6\x06\xa8\x79\xdd\xb4\x10\x11\x50\xdf\x09\x6d\xe6\x63\x61\xb3\x18\x93\x48\x00\xa6\xcc\x62\x98\x06\x51\xe0\x83\xe7\x0b\xee\x05\x46\x44\x73\xe7\xe6\xaf\xb3\x24\x2b\x6f\x51\x2c\x15\x9a\x30\xad\x93\xa9\xb9\xb6\xa8\x29\x28\x24\x20\x4f\xe5\x6c\x02\x59\xa1\x2b\xd9\xed\x22\xcf\x15\x6c\x9b\x44\xb4\xdb\x73\x22\x46\x38\x0d\x09\xf6\x80\x46\x98\xd8\x36\xe0\x28\xe0\x2e\x8e\x5c\x88\x6c\x12\x07\xcc\xf1\xed\xfe\x1c\xdf\xc0\xf7\xf7\xdd\x7d\xcb\xb8\x1e\x17\x45\xae\x0f\x0f\x0e\xb8\x9c\x4c\x64\x66\x4a\x8c\x69\x30\x8d\x1a\x46\x52\x8e\x52\x60\x79\xa2\xf7\xb9\x9c\x1c\xcc\x85\x89\x9b\x51\x2e\x15\x60\xa9\xf7\x33\x28\x0e\xd8\x44\xf8\x04\x97\x5a\x1d\xd4\x0b\x1f\x98\x95\xe7\x13\xf6\x47\x77\xfd\x01\xfa\x74\x79\x76\x36\x40\x7d\xdb\x23\xd4\x0f\x88\x17\x9a\xf1\x94\xfc\x79\x13\xbc\x85\xa3\x33\xf5\xf7\xe9\x3b\xfa\x47\x14\xcf\x3e\xca\x3f\xe0\xc8\x9b\xfd\x39\xfa\xd5\x8c\x3b\x96\xed\x61\x8b\x62\xc7\x42\x96\x75\x68\x3b\x87\x6e\xb0\xef\x39\x2e\x75\xab\xd9\x3b\x25\xe3\x21\xc8\x5c\x37\x88\xdc\x38\x00\x1c\xc7\x0e\xc5\x24\x80\x10\x33\xcb\x73\x70\x6c\x85\x2e\x11\x8e\x17\x89\xc8\xeb\x40\x46\x7e\x16\x64\x64\x1b\x64\x9e\x65\x87\xd4\x76\xcc\x78\xac\xae\xaf\xd9\xde\xd1\xc1\xdd\xd7\x69\xf0\xd7\xe7\x6f\xa3\x44\x5c\xed\x5d\x93\x8b\xec\xdd\xc5\x7d\xc8\xfc\x43\xdb\x3b\x74\xe8\xbe\x6d\x85\x8e\xef\xbf\x18\x64\xdc\x61\xae\xef\xda\x0e\x8e\x68\x48\x31\xb1\x5c\xc0\x2c\xf2\x02\x6c\xf9\xdc\xf2\x3c\x16\x00\xe3\x4e\x0d\x59\x68\x85\xfb\xd6\xcf\x80\xac\x5e\x78\x23\x64\x41\x10\xd8\x01\xb1\x89\x19\x8f\x7e\xb8\xf1\xf9\x97\x93\xcf\xe1\xed\x91\xfb\xe5\xfd\xe5\xf7\x63\xf1\x36\x76\x6f\xce\xaf\x8e\x9d\xdf\xd7\xb0\xcc\xa2\x87\x96\xbf\x1f\xba\x94\x86\xf4\xc5\x20\x23\xae\x17\x5a\x21\x75\x30\x67\x22\xc4\xc4\x0f\x19\x66\x56\x14\x61\x88\xa8\xb0\xc0\xa2\xc0\x19\x69\x20\xb3\xbd\x9f\x04\xd9\x7c\xe1\xcd\x90\x85\x3e\x71\xbd\xa0\x0a\xfa\xfa\x13\x61\x13\xf9\xe7\xd5\xdf\xff\x5c\x7e\x70\xbe\xc9\x0b\x71\x7e\x62\x7f\x3e\xf9\x7c\xa7\xbc\xb7\xcb\x90\x79\xc8\x76\x0f\x3d\xef\xd0\xb1\xf6\x0d\xdc\xf6\xcb\x41\xe6\x84\x44\x38\xd4\x8b\xb0\x67\x87\x31\x26\xc2\x0f\x30\xa5\x14\x30\x25\xd4\x0f\x85\x05\x20\xa8\xd5\x40\xe6\xd0\x9f\x04\xd9\x7c\xe1\x8d\x90\x85\xbe\x43\x3c\x6f\xce\x32\x27\x1d\xcb\xcb\xe9\x34\x93\x57\xae\x47\xf3\xc4\x79\x9f\xb1\x8b\x83\x5b\x3d\x2a\x92\x8e\x30\x6d\x0b\xdb\x16\x72\xdc\x43\xdb\x3e\xb4\xad\xfd\xd0\xf1\x68\xe8\x3d\x15\xb2\xb6\x8f\x68\x4f\x27\xcb\xb7\x7c\xc2\x84\x39\x3e\x08\x26\xd4\x22\x98\x32\x1f\x70\x2c\x88\xef\x79\xd4\x15\x76\x54\xd5\x10\x5d\xf5\xc9\xe6\xaf\xff\xb2\x49\x44\x85\xdf\x4d\xa9\x4d\x91\xe5\x19\x15\xb8\x64\xdf\xf1\x6d\x87\xd8\x3b\xef\x6f\x80\x76\x2b\xa7\x0f\xc5\x61\x3b\x61\x14\x72\x87\x62\xcf\xf2\x8c\x5a\x88\x8b\x43\xf0\x03\xcc\x6c\x66\x81\xcb\x7d\x62\xf3\x8a\x67\x11\x14\xac\x8a\x22\xe6\x41\xec\xba\x9b\xa3\x20\xae\x4b\xfe\xd7\xa3\x60\x61\xc0\x2c\xd7\x64\xc3\x9c\xc8\x24\x60\x36\x0e\x05\x27\xd8\xf5\xac\x20\x62\x40\x6d\x80\x0a\x5a\x96\xe6\xe3\x79\x18\x76\x1c\x45\xe1\x96\x64\x78\x81\xf7\xa8\x30\x76\x92\xd0\x4a\x18\xf5\x65\xb4\x8d\x82\x32\x47\x40\x14\x58\xd8\x0d\x22\xc0\xc4\x72\x7c\x1c\x7a\x6e\x8c\xa3\x58\x08\x3f\x72\x49\x10\x45\x15\x93\x2e\x5a\x4e\xbd\x97\xca\x34\xd9\xa2\x9c\xbf\x67\xe6\x69\xa9\x0b\x50\xa6\xc9\x99\xbf\xc6\x9f\xbf\x9d\xaf\xff\xad\x7f\xea\xbf\x2d\x75\xa1\x58\x9a\xb0\x83\x8b\x99\xc8\x60\xd6\xaf\xce\x3d\x34\x49\xb2\xb2\x00\x33\xd7\xb4\xd5\xbe\xd5\xf9\x61\x03\x46\xd4\xf2\x9d\xc7\x60\xb4\x93\x66\x1e\xc0\xc8\x8d\xc1\xb6\x08\xb5\xb0\x10\x81\x8b\x49\x44\x28\x8e\x02\x87\x61\x27\x64\xdc\xa6\x8c\xc7\x5c\x70\xe3\xeb\xa8\xe6\xeb\x67\x25\x27\xb2\xba\xfc\x98\xd4\x23\x05\x29\x30\x0d\x7a\x50\x5d\x12\x58\xc1\xc7\x28\x2a\x47\x1a\xe9\x1c\x78\x12\x27\xdc\xfc\x3c\x93\xa5\x42\x5c\x66\x71\x32\x2a\x55\xd5\xb5\xf6\x97\x81\xfc\x39\x70\x06\xae\x43\x1e\x45\xb9\x9d\xa4\xfb\x00\x9c\x5e\x14\xda\x96\x1f\x5a\x18\x5c\x9f\x61\x12\x06\xd4\x74\x1a\x0c\x93\x98\x86\x14\x42\x2b\x0a\x69\x25\xff\xb7\x8d\x70\xbe\x2a\xc6\xaf\x75\x7b\xbf\x12\x30\x85\x54\xe6\xa6\xb5\x46\x37\x52\x5d\x57\x17\xbd\x44\x37\x38\x0b\x14\x2b\xf8\x51\x42\x56\xa4\xb3\x67\x92\xd2\x1c\x3e\xee\x0e\x28\x12\x6a\x3f\xaa\x8a\xee\x54\x3a\x56\xeb\xcf\xd2\xd7\x1d\x2d\x98\x91\x13\xd9\x3e\x38\x80\xbd\x20\x0e\x31\x09\x02\x0f\x87\x4e\x10\xe3\x30\xb6\x22\xdb\x66\x40\x23\x51\xe9\x65\xf9\x83\x9f\x7e\xd5\x04\x84\xc7\x47\x61\xf1\x11\xac\x8f\xe1\xbb\xd9\x3f\xde\x97\x7f\x6e\xee\xde\x9d\xcd\xbe\x8a\xeb\x93\xef\x7f\x1d\x5c\x8d\xe2\x3f\xff\xce\x9c\x2f\xa3\xcb\x8f\xf2\x9a\xff\xba\x00\x72\x09\xcf\x7e\xbf\xf3\xdf\xda\x6e\xde\x75\x9c\xb0\xc2\x66\xa7\xbb\xcb\x6e\x31\x0b\x8f\x39\x3c\x8a\x5d\x6c\x88\x84\x09\x84\x1c\x87\x21\x84\xd8\x17\x31\xb5\x63\xee\xf8\x84\xfb\x1b\x62\x3e\xbf\xfc\x90\x4d\x3e\x7b\xf6\x38\x0f\xee\x66\x7b\x7b\x7b\xd2\x8b\x8f\x4e\x6f\x7e\x4f\x4f\xb3\xaf\x6f\x27\x3a\x38\xc8\xbe\x67\xd7\xb7\x65\x91\x1d\x9c\x9f\x3e\x3a\xe6\xb6\x1d\xb7\xfd\xf9\xa1\xff\xa4\x63\x66\x7d\xc8\x0e\xa5\xdc\x23\xc2\xc6\xae\x4b\x08\x26\x0c\x28\x0e\x99\x70\xb0\xc7\x7d\x57\x78\xbe\xf0\xf9\xbc\x4b\x59\x13\xf2\xc5\xf1\x34\xa4\x1f\xae\xee\x6e\x83\xd3\xdb\xbd\xaf\xe9\xf7\x1f\xd1\x27\x2d\x02\x99\xf9\x9e\x94\xdf\x7e\x1c\xdd\xf1\x63\x75\xf6\xfe\x8c\xdc\x1c\x8f\xcf\x9f\x90\xe6\x79\x3b\x4d\x9d\x80\x54\x57\xcf\xdd\x2e\x0f\xbb\xc5\x1c\x90\x50\xc4\x5e\xcc\xb1\xed\x30\x0f\x13\xc1\x04\x66\x81\x0d\x98\xf8\x1e\xe7\xb6\xef\x87\x16\x09\x37\xc4\x4c\x4f\x2e\x35\xb9\x19\x27\x77\xf1\x6c\x1a\x91\xc9\x28\xdd\xfb\xc6\x3e\xb1\x6f\xee\xdf\x67\xe7\x57\x37\xfa\x9b\xfd\xe1\xe4\xd3\xc9\x1f\x9f\x72\xfe\x7e\x44\x1e\x17\x73\xa7\x1f\x0e\x1d\xcf\x21\x95\xb6\x76\xea\xfe\x77\x8b\x99\x72\x41\x02\xc2\x3d\x0c\x31\x73\x4d\x79\xa0\xe6\x48\x77\xb0\x17\x0a\x11\x85\x9e\x03\x22\xdc\x94\x67\xfb\x82\xde\x9d\x1f\x9f\x7d\xf8\x3e\x81\xf8\x2a\xfb\xfd\x80\xbd\xe7\xf9\xf1\xf7\x33\xfb\x93\x1e\x95\x27\xe3\xf3\x0f\xdf\xad\xcb\xe3\x89\x17\x5b\x1f\x69\xb8\x7b\xcc\x2b\x0d\x2d\xb5\xdd\x20\x08\x1f\xd3\x7b\xf4\xba\x5f\x97\xbf\x93\x37\x59\xaf\x27\x94\xcc\x9b\xf7\x9f\x71\xf3\x39\x78\xf5\x5c\xb3\x7e\x78\xfa\xcb\x7a\x93\xf9\x47\xc2\xdb\x6d\xba\x8f\x7d\xb6\x5b\x36\x17\x92\xed\x56\xcb\x69\x7a\xc0\xb6\xee\x10\xb7\x5b\xd5\xa7\xe1\x76\xa3\xf6\x99\xfd\x6e\x66\xcd\xb7\x1f\x3b\x5a\xef\x0e\xd2\xa6\x77\x93\xdb\x67\x75\x9e\x09\xee\x60\xf8\x50\x46\x9b\x77\x02\x3b\xa5\x73\xcd\x4b\x81\x76\xde\xff\x04\x00\x00\xff\xff\xb8\x5a\x57\xef\xe8\x30\x00\x00")

func dbMigrations0001_initialSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations0001_initialSql,
		"db/migrations/0001_initial.sql",
	)
}

func dbMigrations0001_initialSql() (*asset, error) {
	bytes, err := dbMigrations0001_initialSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/0001_initial.sql", size: 12520, mode: os.FileMode(420), modTime: time.Unix(1522053093, 0)}
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
	"db/drop_all_tables.sql": dbDrop_all_tablesSql,
	"db/migrations/0001_initial.sql": dbMigrations0001_initialSql,
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
	"db": &bintree{nil, map[string]*bintree{
		"drop_all_tables.sql": &bintree{dbDrop_all_tablesSql, map[string]*bintree{}},
		"migrations": &bintree{nil, map[string]*bintree{
			"0001_initial.sql": &bintree{dbMigrations0001_initialSql, map[string]*bintree{}},
		}},
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

