package main

import "os"

type FileSystem struct {

}

func (obj *FileSystem) IsDir(path string) bool {
	existed := true
	info, err := os.Stat(path)
	if err != nil {
		existed = false
	} else if os.IsNotExist(err) {
		existed = false
	} else if info.IsDir() == false {
		existed = false
	}
	return existed
}

func (obj *FileSystem) IsExist(path string) bool {
	existed := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

func (obj *FileSystem) IsFile(path string) bool {
	existed := true
	info, err := os.Stat(path)
	if err != nil {
		existed = false
	} else if os.IsNotExist(err) {
		existed = false
	} else if info.IsDir() {
		existed = false
	}
	return existed
}
