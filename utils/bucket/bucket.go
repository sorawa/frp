package bucket

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	_authInfo map[string]string
)

func SetAuth(key, token string) {
	_authInfo[key] = token
}

func GetAuth(key string) string {
	data, flag := _authInfo[key]
	if flag {
		return data
	}
	return ""
}

func ListAuth() {
	for k, v := range _authInfo {
		fmt.Printf("Key : %s Tokey %s \n", k, v)
	}
}

func init() {
	_authInfo = make(map[string]string)
	filepath.Walk("./auth", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		byt, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		SetAuth(filepath.Base(path), string(byt))
		return nil
	})
}
