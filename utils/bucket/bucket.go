package bucket

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func HttpServer(addr string) {
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		token := r.FormValue("token")
		fmt.Printf("add key %v  and token %v \n", key, token)
		SetAuth(key, token)

		_, err := os.Stat("./auth/" + key)
		if err != nil {
			ioutil.WriteFile("./auth/"+key, []byte(token), os.ModeAppend)
		}

		retObj := map[string]interface{}{
			"ret_code": 0,
			"message":  "success",
		}
		retByt, _ := json.Marshal(retObj)
		w.Write(retByt)
	})

	log.Fatal(http.ListenAndServe(addr, nil))
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
