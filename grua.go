package grua

import (
	"fmt"
	"time"
	"math/rand"
	"io/ioutil"
	"github.com/json-iterator/go"
)

/* Constant error number. */
const errc1 string = "61"

/* json iterator variable. */
var json = jsoniter.ConfigCompatibleWithStandardLibrary


/* Type value array. */
type jsonString []string

/* Generate random number. */
func Random(path string) (string) {
	data, _ := loadDataJson(path)
	number := rand.NewSource(time.Now().UnixNano())
	newNumber := rand.New(number)
	randomNumber := newNumber.Intn(len(data))
	return data[randomNumber]
}


/* Load user agents. */
func loadDataJson(path string) (jsonString, string) {
	var data jsonString
	err := ""
	file, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		fmt.Printf("File error: %v\n", err1)
		err = errc1
	}
	json.Unmarshal([]byte(string(file)), &data)
	return data, err
}