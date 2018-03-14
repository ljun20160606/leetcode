package main

import (
	"encoding/json"
	"io"
	"os"
)

// FileSystem ...
type FileSystem interface {
	Exists(path string) bool
	MkdirP(path string) error
	WriteFile(name string, writeFunc func(writer io.Writer) error) error
	WriteFileNotExist(name string, writeFunc func(writer io.Writer) error) error
	WriteJSON(path string, data interface{}) error
	ReadJSON(path string, data interface{}) error
}

// File ...
type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	Stat() (os.FileInfo, error)
}

var fs = &osFileSystem{}

type osFileSystem struct{}

func (*osFileSystem) WriteJSON(path string, data interface{}) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return err
	}
	return nil
}

func (*osFileSystem) ReadJSON(path string, data interface{}) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	err = json.NewDecoder(file).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (*osFileSystem) Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func (*osFileSystem) MkdirP(path string) error {
	if fileInfo, err := os.Stat(path); os.IsNotExist(err) || (fileInfo != nil && !fileInfo.IsDir()) {
		return os.MkdirAll(path, 0751)
	}
	return nil
}

func (*osFileSystem) WriteFile(name string, writeFunc func(writer io.Writer) error) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	return writeFunc(file)
}

func (o *osFileSystem) WriteFileNotExist(name string, writeFunc func(writer io.Writer) error) error {
	_, err := os.Stat(name)
	if err != nil {
		o.WriteFile(name, writeFunc)
	}
	return nil
}
