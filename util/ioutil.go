package util

import "os"

func MkDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsExist(err) {
		return true, nil
	}
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return false, err
	}
	return true, nil
}
