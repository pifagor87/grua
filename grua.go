package grua

import (
	"fmt"
	"github.com/json-iterator/go"
	"io/ioutil"
	"math/rand"
	"time"
)

/* Constant error number. */
const errc1 string = "61"
const errc2 string = "62"

/* json iterator variable. */
var json = jsoniter.ConfigCompatibleWithStandardLibrary

/* Type value json. */
type jsonString []string

/* Generate random user agent. */
func Random(path string) (string, string) {
	data, error := loadDataJson(path)
	if error != "" {
		return "", error
	}
	number := rand.NewSource(time.Now().UnixNano())
	newNumber := rand.New(number)
	randomNumber := newNumber.Intn(len(data))
	result := data[randomNumber]
	if result == "" {
		return "", errc2
	}
	return result, ""
}

/* Load user agents. */
func loadDataJson(path string) (jsonString, string) {
	var data jsonString
	file, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		fmt.Printf("File error: %v\n", err1)
	}
	json.Unmarshal([]byte(string(file)), &data)
	return data, errc1
}
