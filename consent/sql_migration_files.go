// Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// migrations/sql/shared/.gitkeep
// migrations/sql/shared/1.sql
// migrations/sql/shared/2.sql
// migrations/sql/shared/3.sql
// migrations/sql/mysql/.gitkeep
// migrations/sql/mysql/4.sql
// migrations/sql/postgres/.gitkeep
// migrations/sql/postgres/4.sql
// migrations/sql/tests/.gitkeep
// migrations/sql/tests/1_test.sql
// migrations/sql/tests/2_test.sql
// migrations/sql/tests/3_test.sql
// migrations/sql/tests/4_test.sql
package consent

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

var _migrationsSqlSharedGitkeep = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrationsSqlSharedGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlSharedGitkeep,
		"migrations/sql/shared/.gitkeep",
	)
}

func migrationsSqlSharedGitkeep() (*asset, error) {
	bytes, err := migrationsSqlSharedGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/.gitkeep", size: 0, mode: os.FileMode(438), modTime: time.Unix(1540902707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared1Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\x5d\x6f\xda\x30\x14\x86\xaf\x1d\x29\xff\xe1\x5c\x16\x6d\x95\xa6\x6a\xbd\xea\x15\x1b\x99\x34\x8d\x41\x85\x40\x53\xaf\x2c\xe3\x1c\x88\xd7\xc4\x66\xb6\x53\xb6\x7f\x3f\x25\x04\x9a\x4f\xc7\x65\xe3\x96\xf7\x7c\x3e\xaf\x8f\x72\x7b\x0b\xef\x32\xb1\xd7\xcc\x22\x6c\x0e\x61\xf0\x79\x15\x4d\xd7\x11\xac\xa7\x9f\xe6\x11\x24\x7f\x62\xcd\xa8\x62\xb9\x4d\xee\x28\x57\xd2\xa0\xb4\x54\xe3\xaf\x1c\x8d\x85\x9b\x30\x20\x3c\x61\x69\x8a\x72\x8f\x00\x84\x10\x80\x17\xa6\x79\xc2\xf4\xcd\xc7\x0f\x13\x58\x2c\xd7\xb0\xd8\xcc\xe7\xf0\xb8\xfa\xfa\x7d\xba\x7a\x82\x6f\xd1\xd3\xfb\x30\x20\x2f\xa8\xc5\x4e\xa0\x86\xcb\xaf\x2f\xaa\x50\xf2\x54\x14\x05\x45\x5c\xe6\x7e\xd5\xdd\xdd\xdf\x37\x85\x26\xdf\xfe\x44\x6e\xc9\x98\xae\x6a\x9d\xe6\x3a\x2d\xa5\x16\x7f\xdb\x66\xa2\x67\x71\x38\x67\x01\xd8\x2a\x95\xf6\xc5\x63\x4c\x0d\x57\x07\x24\xa4\x93\x80\x1b\xbd\x7b\x4d\x30\x34\x58\xb1\x50\x94\x56\x70\x56\xe4\x62\x96\x58\x91\xa1\xb1\x2c\x3b\xf4\x94\x62\xb6\x58\x6e\x4d\x71\xde\xeb\x2c\xfa\x32\xdd\xcc\xd7\x20\xd5\xf1\x66\x52\xc4\x28\x11\xf3\x02\x53\xd1\x55\x67\xbc\x30\x98\x3c\xb8\xf0\xd6\x7a\x12\x4a\xfe\x33\xe5\xf1\x55\x5d\x7c\xd0\xa2\xd6\x71\x81\xd7\x4a\xaf\xb2\x00\x40\xb7\xad\x11\x0b\x78\x7b\xf2\x1a\x80\x1e\xbe\xf8\x9f\x8c\x0d\x1a\x23\x94\x2c\x19\x8b\xf8\x34\xb0\x03\x46\x1b\x71\xbb\x5b\x00\xc7\x8c\x8b\xe5\x8f\xd3\x8c\x15\xa8\x76\xa1\xc6\xfe\xc6\xe6\x68\x9d\x22\x9a\x30\x19\xa7\x18\xf7\x98\xd5\x7b\x98\xbd\x66\xb2\xe6\xd6\x32\xae\xe3\x0d\x8d\x19\x66\x5b\xd4\xce\x13\x71\x92\xd0\x9d\xd2\x95\x4a\xc8\x66\x16\xd4\xba\xfc\xef\x9c\xa3\xb7\x50\xd3\x3c\xc4\xb9\xdb\x8b\x7f\x2a\xa0\x94\x71\x8e\xc6\x50\xab\x9e\x51\xf6\x59\xbc\x92\x89\xf8\x2c\xe9\xbb\x85\x1d\x33\x36\x9b\xa8\x54\x47\x66\x68\x6e\x30\x86\x81\x9d\x5c\x77\x76\x1c\x44\x7d\x81\xb6\x7c\xe6\x7c\xa9\x17\xaa\x5e\x50\x9d\x4c\x07\x91\x32\xfe\x46\xe6\x7e\xc8\x3b\x94\xc6\x20\x0d\x32\x0a\x83\xfa\xb7\xc0\x4c\x1d\x65\x18\xcc\x56\xcb\x47\x8f\x07\xf8\x30\xac\xec\xe7\xeb\x1f\x50\x79\xd5\x11\x30\x70\x0c\xde\xdc\x53\x2d\xf0\x6f\x00\x00\x00\xff\xff\xc9\x66\xd1\x81\x15\x09\x00\x00")

func migrationsSqlShared1SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared1Sql,
		"migrations/sql/shared/1.sql",
	)
}

func migrationsSqlShared1Sql() (*asset, error) {
	bytes, err := migrationsSqlShared1SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/1.sql", size: 2325, mode: os.FileMode(438), modTime: time.Unix(1540900221, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared2Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\xcd\x4e\x33\x21\x14\x86\xd7\x90\x70\x0f\x67\xd7\x36\x5f\xbb\x69\xd2\x55\x57\x7c\x1d\x8c\x46\xfa\x13\x32\x63\xd2\x15\xa1\x40\x1d\x4c\x05\x05\x46\xe3\xdd\x9b\x69\xa6\xd6\x18\xdb\xa8\x65\xc1\xea\xe5\x3c\x87\xf7\x19\x8d\xe0\xdf\xa3\xbb\x8f\x2a\x5b\xa8\x9e\x08\xa6\xbc\x64\x02\x4a\xfa\x9f\x33\xa8\xdf\x4c\x54\x32\xa8\x26\xd7\x63\xa9\x83\x4f\xd6\x67\x19\xed\x73\x63\x53\x06\x5a\x14\xb0\x0d\x51\x5b\x23\x53\xb3\x79\xb0\x3a\x4b\x67\xac\xcf\x6e\xeb\x6c\x84\x3b\x2a\x66\xd7\x54\xf4\xc7\x93\xc9\x00\x16\x15\xe7\x50\xb0\x2b\x5a\xf1\x12\x7a\xbd\xe9\x19\x4a\x7b\xb7\x43\xb4\xca\x2e\xf8\x03\x4c\xd6\xca\x9b\x9d\x35\x97\x41\x67\x82\xd1\x92\x7d\x47\x0d\x9b\x6d\x93\xb4\xca\xd6\x7c\x5d\x20\xd9\x94\x5c\xf0\xd0\x27\x18\x75\x44\x00\x84\x10\x74\xe7\x45\x45\x5d\xab\xd8\x21\x97\xe5\x1e\x3b\x24\x18\xe9\x9d\x6b\xdb\x72\xe6\x98\x3e\x99\x3d\x7c\xe5\xb8\x06\x3a\x93\x5e\x89\x9b\x39\x15\x6b\xb8\x65\xeb\x7e\xf7\x72\x08\x1f\xb8\x01\xc1\x83\x29\xc1\x04\x7f\x16\x5b\x84\x57\xff\x0b\xb5\x85\x58\xae\x60\xb6\xe4\xd5\x7c\x71\xba\xed\x0b\x2c\xfe\x74\xfe\x3e\xf7\x17\x5d\xd3\xf7\x00\x00\x00\xff\xff\x8a\x43\xb4\xc1\xd6\x02\x00\x00")

func migrationsSqlShared2SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared2Sql,
		"migrations/sql/shared/2.sql",
	)
}

func migrationsSqlShared2Sql() (*asset, error) {
	bytes, err := migrationsSqlShared2SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/2.sql", size: 726, mode: os.FileMode(438), modTime: time.Unix(1540900276, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlShared3Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x90\xb1\x8b\x83\x30\x14\x87\xe7\x08\xfe\x0f\x6f\xf3\x8e\x43\x38\x8e\xdb\x3a\xa5\xc6\xd2\x21\xd5\x12\x4c\xd7\x10\xf4\xa1\x01\x9b\xb4\x26\x52\xfa\xdf\x17\x97\x52\x70\x28\xda\xe5\x8d\xdf\xfb\x7e\x5f\x9a\xc2\xcf\xd9\xb4\x83\x0e\x08\xf2\x12\x47\x94\x57\xb9\x80\x8a\x6e\x79\x0e\xdd\xbd\x19\xb4\x72\x7a\x0c\xdd\x9f\xaa\x9d\xf5\x68\x83\x1a\xf0\x3a\xa2\x0f\x40\x19\x83\xde\xb5\xc6\x2a\x8f\xde\x1b\x67\x95\x69\xe0\x44\x45\xb6\xa7\xe2\xeb\xff\xf7\x1b\x0a\xc9\x39\xb0\x7c\x47\x25\xaf\x20\x49\x36\xab\xd8\x75\xa7\xfb\x1e\x6d\x8b\x6b\xd1\xd3\x45\x1b\x4c\xad\xc3\xe4\xf8\x81\x3d\x21\x84\xc4\xd1\x6b\x2f\xe6\x6e\x76\xc1\x2a\x26\xca\x23\x64\x25\x97\x87\x62\xf6\x7b\x49\x9d\x39\xe7\x59\x69\x45\x89\x37\x56\x8f\x00\x00\x00\xff\xff\xff\xe9\x89\x60\x20\x02\x00\x00")

func migrationsSqlShared3SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlShared3Sql,
		"migrations/sql/shared/3.sql",
	)
}

func migrationsSqlShared3Sql() (*asset, error) {
	bytes, err := migrationsSqlShared3SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/shared/3.sql", size: 544, mode: os.FileMode(438), modTime: time.Unix(1540900335, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlMysqlGitkeep = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrationsSqlMysqlGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlMysqlGitkeep,
		"migrations/sql/mysql/.gitkeep",
	)
}

func migrationsSqlMysqlGitkeep() (*asset, error) {
	bytes, err := migrationsSqlMysqlGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/mysql/.gitkeep", size: 0, mode: os.FileMode(438), modTime: time.Unix(1540902707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlMysql4Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\xd2\xc1\x4b\xc3\x30\x14\x06\xf0\x7b\xa1\xff\xc3\xbb\xed\x20\xbb\x78\x2d\x1e\xaa\xa9\x20\x64\xcb\x98\x29\xe8\x29\x3c\x9a\xc7\x1a\xd0\x17\xcd\x52\xc4\xff\x5e\xa6\x3b\x14\x49\x2a\xab\xbd\x94\x9e\x7e\xbc\xef\xfb\xb2\x5e\xc3\xd5\xab\x3b\x04\x8c\x04\xed\x5b\x59\xd4\x52\x37\x7b\xd0\xf5\xad\x6c\xa0\xff\xb4\x01\x8d\xc7\x21\xf6\xd7\xa6\xf3\x7c\x24\x8e\x26\xd0\xfb\x40\xc7\x08\xb5\x10\x70\xfe\x27\x6b\x30\x1a\x1c\xac\x23\xee\x08\x74\xf3\xa4\x61\xdb\x4a\x59\x4d\x70\xa7\x2f\x71\x74\x1d\x46\xe7\x79\x29\xf5\xd7\x91\xa6\x47\xb6\x2f\x64\xbf\xd9\x43\x40\x9e\x42\xcb\xa2\xdd\x89\x5a\xff\x11\xfb\xb1\xd1\xe9\x03\x6f\x56\xab\x2a\x4d\x64\xa2\xce\x90\x72\xf1\x4e\x54\x22\xde\x0f\x74\xc1\xa6\x1b\x25\x1e\xee\x9f\x27\x07\x50\xb3\xa7\x5d\x06\xcf\x55\x70\xd6\xf3\x23\xab\xd1\xd0\xe3\x47\x2f\xfc\x07\x5f\x50\x91\xd8\xab\x1d\xdc\x29\xd9\x6e\xb6\xe9\x28\x33\x9a\xf9\x97\x99\x2b\x64\x8c\x26\x5a\xa9\xca\xe2\x2b\x00\x00\xff\xff\xc3\xd9\x79\x09\xfb\x03\x00\x00")

func migrationsSqlMysql4SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlMysql4Sql,
		"migrations/sql/mysql/4.sql",
	)
}

func migrationsSqlMysql4Sql() (*asset, error) {
	bytes, err := migrationsSqlMysql4SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/mysql/4.sql", size: 1019, mode: os.FileMode(438), modTime: time.Unix(1540933030, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlPostgresGitkeep = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrationsSqlPostgresGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlPostgresGitkeep,
		"migrations/sql/postgres/.gitkeep",
	)
}

func migrationsSqlPostgresGitkeep() (*asset, error) {
	bytes, err := migrationsSqlPostgresGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/postgres/.gitkeep", size: 0, mode: os.FileMode(438), modTime: time.Unix(1540902707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlPostgres4Sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x90\x31\x4b\xc6\x30\x10\x40\xe7\x14\xfa\x1f\x6e\xeb\x20\x5d\x5c\x3b\x45\x13\xa7\xd8\x4a\x49\xc0\x2d\x1c\xcd\xd1\x06\xf4\xa2\x69\x8a\xf8\xef\x45\x71\x28\x52\x45\xfd\xf8\x96\xe3\xa6\xf7\xee\x5e\xdb\xc2\xc5\x63\x9c\x33\x16\x02\xf7\x54\x57\xd2\x58\x3d\x82\x95\x57\x46\xc3\xf2\x1a\x32\xfa\x84\x5b\x59\x2e\xfd\x94\x78\x25\x2e\x3e\xd3\xf3\x46\x6b\x01\xa9\x14\x7c\xee\x14\x3c\x16\x8f\x5b\x88\xc4\x13\x81\xd5\xf7\x16\x7a\x67\x0c\x28\x7d\x23\x9d\xb1\xd0\x34\xdd\x0f\xe4\xf7\x49\x5c\xe2\x84\x25\x26\x3e\x83\xe0\xcb\xe9\x7e\x41\x0e\x0f\x14\x3e\x0c\x73\x46\xfe\x25\x5f\x08\x21\xea\x6a\x1f\x4c\xa5\x17\xfe\x43\x32\x35\x0e\x77\x70\x3d\x18\x77\xdb\x1f\x7f\xf6\x8f\x4a\x27\x31\xbf\x0b\xb3\x87\x1e\x04\xea\xea\xea\x2d\x00\x00\xff\xff\x4d\x53\x1e\xa3\x37\x02\x00\x00")

func migrationsSqlPostgres4SqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlPostgres4Sql,
		"migrations/sql/postgres/4.sql",
	)
}

func migrationsSqlPostgres4Sql() (*asset, error) {
	bytes, err := migrationsSqlPostgres4SqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/postgres/4.sql", size: 567, mode: os.FileMode(438), modTime: time.Unix(1540900421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTestsGitkeep = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func migrationsSqlTestsGitkeepBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTestsGitkeep,
		"migrations/sql/tests/.gitkeep",
	)
}

func migrationsSqlTestsGitkeep() (*asset, error) {
	bytes, err := migrationsSqlTestsGitkeepBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/.gitkeep", size: 0, mode: os.FileMode(438), modTime: time.Unix(1540902707, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests1_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\x41\x8f\xda\x30\x10\x85\xcf\x1b\x29\xff\x61\x6e\x09\xaa\x23\x2d\xad\xd4\x4b\x4f\x95\xba\x07\xa4\x0a\xa4\x02\xed\xd1\x32\xf6\x84\xb8\x04\x9b\x8e\x9d\xd2\xaa\xea\x7f\xaf\x5c\x27\x10\x12\x04\x15\xe7\x9e\xf0\x0c\x63\xf9\xbd\xef\x8d\x52\x14\xf0\x6a\xaf\xb7\x24\x3c\xc2\xfa\x90\x26\xb3\xf9\xf2\xe5\xd3\x0a\x66\xf3\xd5\x22\x4d\x9e\xaa\x9f\x8a\x04\xb7\xa2\xf1\xd5\x6b\x2e\xad\x71\x68\x3c\x27\xfc\xd6\xa0\xf3\x90\xcb\x4a\xd4\x35\x9a\x2d\x32\xf8\x8e\xa4\x4b\x8d\xc4\x40\xd6\x3a\x0c\x69\xc5\xc0\x35\x9b\xaf\x28\x3d\x83\xf6\x06\x6f\xa8\x66\xe0\x76\xfa\x70\x6a\xa1\xe2\x4e\xda\x03\x32\x90\x8e\x4a\x06\xe1\x25\x34\x5e\x4b\x11\xfe\x12\xbe\x3f\x18\x2a\xab\x95\x0c\x42\x3c\xfe\xf0\x93\x34\xf9\xfc\xfe\xe3\xfa\x65\x99\x26\x4f\x79\x36\x2d\x4e\x72\x32\x06\xd9\xb4\xe8\x24\xc5\x2a\xca\x8a\xe7\x56\x57\x2c\x08\x95\xa6\x58\x95\xa2\x76\x18\x27\x82\xa6\xf6\xa2\xa3\x32\x63\x30\x5f\x7c\xc9\x27\xa7\x9f\xec\xd7\xef\x6c\xf2\x2e\x4d\x6e\xf2\xea\x79\xd1\xd6\xfc\xc7\xf6\x18\x36\x87\xce\x69\x6b\x20\x0f\x68\xc6\x3e\x5b\x51\x43\x53\x61\xf0\xfc\x7c\x14\x7f\xff\xe9\xc1\x86\xf3\x4a\x18\x55\xa3\xba\x88\x6c\x4b\xc2\xf4\xf8\x13\xee\x71\xbf\x09\x09\x76\x27\x5e\x5a\x62\x80\x44\x96\x86\x31\xb4\x5e\xb8\x90\x12\x9d\xe3\xde\xee\xd0\x9c\xbb\x5a\x75\x9d\xb1\xcd\xa3\x70\xbc\x71\xa8\xee\x84\xd7\x05\xe0\xa9\x41\x06\x6f\xde\x3e\x3f\x47\xe8\x17\x09\x0c\x5a\x7f\xe3\xbb\x82\x06\xfe\x61\x99\xaf\x12\xea\xed\xef\x4d\x36\x42\x8e\x00\x3d\x6e\xbc\xd9\x5c\xb5\x9d\x4d\x47\x4b\xd8\xf3\xdb\xff\xf8\x7d\xb0\x47\x93\x26\x7f\x02\x00\x00\xff\xff\x44\x70\x31\x6d\x0f\x05\x00\x00")

func migrationsSqlTests1_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests1_testSql,
		"migrations/sql/tests/1_test.sql",
	)
}

func migrationsSqlTests1_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests1_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/1_test.sql", size: 1295, mode: os.FileMode(438), modTime: time.Unix(1540936819, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests2_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x4d\x6f\xd4\x30\x10\x3d\x37\x52\xfe\xc3\xdc\xb2\x2b\x1c\xa9\x2c\x12\x17\x4e\x48\xf4\x50\x09\x6d\x25\xda\xc2\xd1\xf2\xda\x93\x8d\x69\xd6\x5e\x66\x1c\x0a\x42\xfc\x77\x64\x9c\xaf\x26\xfd\x50\x7b\x81\x53\x32\xe3\x19\xcf\x9b\x37\x6f\x5c\x96\xf0\xea\x60\xf7\xa4\x02\xc2\xf5\x31\xcf\xce\xb7\x97\x67\x9f\xae\xe0\x7c\x7b\x75\x91\x67\x27\xf5\x4f\x43\x4a\x7a\xd5\x86\x7a\x23\xb5\x77\x8c\x2e\x48\xc2\x6f\x2d\x72\x80\x95\xae\x55\xd3\xa0\xdb\xa3\x80\xef\x48\xb6\xb2\x48\x02\x74\x63\x63\x90\x35\x02\xb8\xdd\x7d\x45\x1d\x04\x74\x19\xb2\xa5\x46\x00\xdf\xd8\xe3\xe0\x42\x23\x59\xfb\x23\x0a\xd0\x4c\x95\x80\x58\x09\x5d\xb0\x5a\xc5\x23\x15\xa6\x81\xd1\xf2\xd6\xe8\x08\x24\xe0\x8f\x20\xa0\xf2\xa4\xe3\x0d\xa9\x8e\xb4\x26\xa6\x46\x18\xeb\x3c\xfb\xfc\xfe\xe3\xf5\xd9\x65\x9e\x9d\xac\x8a\x4d\x39\x20\x2d\x04\x14\x9b\xb2\x47\x9b\xac\x84\x38\xfd\x77\x57\x25\x83\xd0\x58\x4a\x56\xa5\x1a\xc6\x14\x11\xe1\x76\x89\x4c\x55\x21\x60\x7b\xf1\x65\xb5\x1e\x3e\xc5\xaf\xdf\xe9\x34\x81\x8b\x37\x16\xeb\x77\x79\xf6\x28\xb5\x93\xb6\xad\x77\xff\x0d\xc3\xff\x98\xc6\xe7\xd2\xc6\xc8\x6c\xbd\x83\x55\xa4\x66\xd9\x67\x07\x6a\xde\x54\x0c\x1c\xcb\x27\xf0\x4f\x97\x9e\x2d\x83\xac\x95\x33\x0d\x9a\x3b\x23\xdb\x93\x72\x13\xfe\x09\x0f\x78\xd8\xc5\x09\xf6\x7f\xb2\xf2\x24\x00\x89\x3c\xcd\xc7\xd0\xf5\x22\x95\xd6\xc8\x2c\x83\xbf\x41\x37\x7a\xad\xe9\x3d\xcb\x36\x6f\x15\xcb\x96\xd1\x3c\x31\xbc\x7e\x00\x81\x5a\x14\xf0\xe6\xed\xe9\x69\xaf\xdd\xbb\x42\x9e\xba\xfe\x8e\xef\x85\x62\xbe\x97\xa1\x89\x7e\x1f\xe5\x46\xe9\x05\x41\x0f\x37\xfe\xf2\x77\x21\x4e\xfe\x3e\x46\x8a\xd7\x0b\x7d\x8e\x4a\x7e\xce\x9e\xfb\x5d\xd5\x72\x87\xf8\x01\xed\x0e\x94\x2c\xd7\x7c\x92\x3e\x6f\x24\x21\x9f\x6d\xe1\x18\xde\x21\x9b\x3e\xf6\x1f\xfc\xad\xcb\xb3\x3f\x01\x00\x00\xff\xff\xfc\x65\xc0\x9d\xff\x05\x00\x00")

func migrationsSqlTests2_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests2_testSql,
		"migrations/sql/tests/2_test.sql",
	)
}

func migrationsSqlTests2_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests2_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/2_test.sql", size: 1535, mode: os.FileMode(438), modTime: time.Unix(1540936819, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests3_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x4f\x6f\xd3\x30\x14\x3f\x2f\x52\xbe\xc3\xbb\xa5\x15\x8e\x34\xa8\xc4\x85\x13\x12\x3b\x4c\x42\x9d\xc4\x36\x38\x5a\xae\xfd\xd2\x9a\xa5\x76\x79\x76\x18\x08\xf1\xdd\x91\x6b\x27\xf5\x12\xd6\x08\x4e\x3b\x35\xef\xf9\x3d\xfb\xf7\xe7\xd9\xad\x6b\x78\xb5\xd7\x5b\x12\x1e\xe1\xfe\x50\x16\xd7\xeb\xdb\xab\x4f\x77\x70\xbd\xbe\xbb\x29\x8b\x8b\xdd\x4f\x45\x82\x5b\xd1\xf9\xdd\x1b\x2e\xad\x71\x68\x3c\x27\xfc\xd6\xa1\xf3\xb0\x90\x3b\xd1\xb6\x68\xb6\xc8\xe0\x3b\x92\x6e\x34\x12\x03\xd9\xea\x50\xa4\x15\x03\xd7\x6d\xbe\xa2\xf4\x0c\x52\x07\xef\xa8\x65\xe0\x1e\xf4\x61\x48\xa1\xe2\x4e\xda\x03\x32\x90\x8e\x1a\x06\xe1\x24\x34\x5e\x4b\x11\x96\x84\xcf\x0b\x43\x64\xb5\x92\x01\x88\xc7\x1f\x9e\x41\x63\x49\x86\x1d\xe2\x39\x5c\xab\xd0\x1a\x61\xb4\x76\xab\x0d\x77\xe8\x9c\xb6\xe6\x88\x26\x66\x06\xcc\xcb\xb2\xf8\xfc\xfe\xe3\xfd\xd5\x6d\x59\x5c\x2c\xaa\x55\x3d\x2c\x54\x0c\xaa\x55\xdd\x13\x8a\x51\x24\x15\xbf\xd3\x69\x31\x20\x54\x9a\x62\xd4\x88\xd6\x61\xac\x08\x8c\x52\xa3\xa3\xa6\x62\xb0\xbe\xf9\xb2\x58\x0e\x3f\xd5\xaf\xdf\x71\x35\xe2\x0f\x3b\xc6\x78\x0c\x3a\xcb\x66\xf8\x96\xef\xca\xe2\xac\x51\x99\x88\x61\x9f\x97\xe3\xd7\x98\xdf\x4b\xf0\x20\xaa\x9b\x30\xd5\x5a\xfd\xb3\xbc\xa9\x15\x16\x41\xc2\xa9\x1e\x09\xea\x98\x6a\x28\x3c\x81\x8a\x94\xe6\x8f\x1e\x5d\x41\xbe\x13\x46\xb5\xa8\x9e\x58\xbb\x25\x61\x32\x9f\x08\xf7\xb8\xdf\x04\xa7\xfb\x2f\xde\x58\x62\x80\x44\x96\xc6\x76\xf5\xd6\x08\x29\xd1\x39\xee\xed\x03\x9a\x53\x56\xab\x3e\x33\xa5\xf9\x28\x1c\xef\x1c\xce\x59\xda\xdb\xe2\xa9\x43\x06\xab\xb7\x97\x97\xbd\x15\x4f\x7d\xc9\x53\x47\x53\xff\x73\xe8\xff\xaa\x50\x36\xe7\x67\xb5\x11\x72\x22\xd0\xf3\xc4\xcf\xbc\x46\x73\x9a\x1c\xaf\xff\x54\x91\xea\xf5\x64\x6a\x4f\xf3\x9d\x3d\x1d\xb3\xd2\xd8\x4d\xd3\xb9\x84\xf8\x99\xd9\x1d\x24\x99\x3e\x07\x59\xfb\x98\xc8\xf0\x70\xe5\x77\xf3\x54\x9e\x90\xe5\x7f\x31\x1f\xec\xa3\x29\x8b\x3f\x01\x00\x00\xff\xff\x44\xb2\xfb\xa1\x75\x06\x00\x00")

func migrationsSqlTests3_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests3_testSql,
		"migrations/sql/tests/3_test.sql",
	)
}

func migrationsSqlTests3_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests3_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/3_test.sql", size: 1653, mode: os.FileMode(438), modTime: time.Unix(1540936819, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrationsSqlTests4_testSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\xcd\x6e\x13\x31\x10\x3e\x77\xa5\x7d\x87\xb9\x6d\x22\xbc\x52\x81\x8a\x0b\x27\x24\x7a\xa8\x84\x52\x89\xb6\x70\xb4\x1c\x7b\x36\x31\xdd\xd8\x61\x6c\x53\x10\xe2\xdd\x91\xeb\xfd\x71\x77\x69\x82\xe0\x92\x53\x76\xc6\xe3\xf1\xf7\x63\x4f\xea\x1a\x5e\xec\xf4\x86\x84\x47\xb8\xdb\x97\xc5\xd5\xea\xe6\xf2\xe3\x2d\x5c\xad\x6e\xaf\xcb\xe2\x6c\xfb\x43\x91\xe0\x56\x04\xbf\x7d\xc5\xa5\x35\x0e\x8d\xe7\x84\x5f\x03\x3a\x0f\x0b\xb9\x15\x6d\x8b\x66\x83\x0c\xbe\x21\xe9\x46\x23\x31\x90\xad\x8e\x45\x5a\x31\x70\x61\xfd\x05\xa5\x67\xd0\xed\xe0\x81\x5a\x06\xee\x5e\xef\x87\x14\x2a\xee\xa4\xdd\x23\x03\xe9\xa8\x61\x10\x4f\x42\xe3\xb5\x14\x71\x49\xf8\xbc\x30\x46\x56\x2b\x19\x81\x78\xfc\xee\x19\x34\x96\x64\xec\x90\xce\xe1\x5a\xc5\xad\x09\x46\x6b\x37\xda\x70\x87\xce\x69\x6b\x1e\xd1\xa4\x4c\x86\x39\xef\xcc\x45\x50\x1a\x8d\xc4\x65\x59\x7c\x7a\xf7\xe1\xee\xf2\xa6\x2c\xce\x16\xd5\x45\x3d\xd4\x57\x0c\xaa\x8b\xba\xe7\x99\xa2\xc4\x35\x7d\x77\x20\x52\x40\xa8\x34\xa5\xa8\x11\xad\xc3\x54\x11\x89\x76\x1b\x1d\x35\x15\x83\xd5\xf5\xe7\xc5\x72\xf8\xa9\x7e\xfe\x4a\xab\x89\x56\xec\x98\xe2\x29\x97\x2c\x3b\xc5\x27\x82\xaa\x96\x6f\xcb\xe2\xa0\x91\x99\xc8\xb1\xe1\xe9\xf8\x39\x37\xed\x64\x3d\x4a\xea\x77\x50\xeb\xde\x93\x7f\x91\xbf\xeb\x01\x8b\xc8\x77\xae\x57\x87\x79\xca\x39\x16\x8e\xe8\x12\xb7\xe3\x47\x4f\x9e\x30\xdf\x0a\xa3\x5a\x54\x4f\xac\xdf\x90\x30\x99\x8f\x84\x3b\xdc\xad\xe3\x4d\xe8\xbf\x78\x63\x89\x01\x12\x59\x9a\xda\xd9\x5b\x27\xa4\x44\xe7\xb8\xb7\xf7\x68\xc6\xac\x56\x7d\x66\x4e\xf3\x41\x38\x1e\x1c\xaa\xf1\xfc\xbf\x37\xbc\x37\xcd\x53\x40\x06\xaf\xdf\x9c\x9f\xf7\x46\x3d\x75\x2d\x4f\x8d\x96\xff\xc7\x93\xf9\xa3\x7e\xd9\x2b\x39\xa8\x9c\x90\x33\xf9\x0e\xc9\xf2\xec\xac\x3b\x26\xce\xe3\x14\x99\x4b\x53\xbd\x9c\x5d\xee\x51\x93\x6c\x02\x1d\x95\xc6\xae\x9b\xe0\x3a\xc4\xcf\xdc\xec\x41\x92\xf9\x30\xc9\xb6\x4f\x89\x0c\xf3\x2f\x7f\xc2\x63\x79\x87\x2c\xff\x03\x7b\x6f\x1f\x4c\x59\xfc\x0e\x00\x00\xff\xff\xf7\x44\x67\x29\xd3\x06\x00\x00")

func migrationsSqlTests4_testSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrationsSqlTests4_testSql,
		"migrations/sql/tests/4_test.sql",
	)
}

func migrationsSqlTests4_testSql() (*asset, error) {
	bytes, err := migrationsSqlTests4_testSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/sql/tests/4_test.sql", size: 1747, mode: os.FileMode(438), modTime: time.Unix(1540936819, 0)}
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
	"migrations/sql/shared/.gitkeep":   migrationsSqlSharedGitkeep,
	"migrations/sql/shared/1.sql":      migrationsSqlShared1Sql,
	"migrations/sql/shared/2.sql":      migrationsSqlShared2Sql,
	"migrations/sql/shared/3.sql":      migrationsSqlShared3Sql,
	"migrations/sql/mysql/.gitkeep":    migrationsSqlMysqlGitkeep,
	"migrations/sql/mysql/4.sql":       migrationsSqlMysql4Sql,
	"migrations/sql/postgres/.gitkeep": migrationsSqlPostgresGitkeep,
	"migrations/sql/postgres/4.sql":    migrationsSqlPostgres4Sql,
	"migrations/sql/tests/.gitkeep":    migrationsSqlTestsGitkeep,
	"migrations/sql/tests/1_test.sql":  migrationsSqlTests1_testSql,
	"migrations/sql/tests/2_test.sql":  migrationsSqlTests2_testSql,
	"migrations/sql/tests/3_test.sql":  migrationsSqlTests3_testSql,
	"migrations/sql/tests/4_test.sql":  migrationsSqlTests4_testSql,
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
	"migrations": &bintree{nil, map[string]*bintree{
		"sql": &bintree{nil, map[string]*bintree{
			"mysql": &bintree{nil, map[string]*bintree{
				".gitkeep": &bintree{migrationsSqlMysqlGitkeep, map[string]*bintree{}},
				"4.sql":    &bintree{migrationsSqlMysql4Sql, map[string]*bintree{}},
			}},
			"postgres": &bintree{nil, map[string]*bintree{
				".gitkeep": &bintree{migrationsSqlPostgresGitkeep, map[string]*bintree{}},
				"4.sql":    &bintree{migrationsSqlPostgres4Sql, map[string]*bintree{}},
			}},
			"shared": &bintree{nil, map[string]*bintree{
				".gitkeep": &bintree{migrationsSqlSharedGitkeep, map[string]*bintree{}},
				"1.sql":    &bintree{migrationsSqlShared1Sql, map[string]*bintree{}},
				"2.sql":    &bintree{migrationsSqlShared2Sql, map[string]*bintree{}},
				"3.sql":    &bintree{migrationsSqlShared3Sql, map[string]*bintree{}},
			}},
			"tests": &bintree{nil, map[string]*bintree{
				".gitkeep":   &bintree{migrationsSqlTestsGitkeep, map[string]*bintree{}},
				"1_test.sql": &bintree{migrationsSqlTests1_testSql, map[string]*bintree{}},
				"2_test.sql": &bintree{migrationsSqlTests2_testSql, map[string]*bintree{}},
				"3_test.sql": &bintree{migrationsSqlTests3_testSql, map[string]*bintree{}},
				"4_test.sql": &bintree{migrationsSqlTests4_testSql, map[string]*bintree{}},
			}},
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