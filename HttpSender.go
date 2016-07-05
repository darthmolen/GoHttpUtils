package HttpSender

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// postFromReader takes an io.Reader and issues a POST to a destination server
func postFromReader(URL string, buffer io.Reader, mime string) (string, error) {
	r, err := http.Post(URL, mime, buffer)

	if err != nil {
		return "", err
	}

	defer r.Body.Close()
	resp, _ := ioutil.ReadAll(r.Body)

	ret := string(resp)
	return ret, nil
}

// PostByteArray sends a Byte Array to another server via POST
func PostByteArray(URL string, data []byte, mime string) (string, error) {

	reader := bytes.NewReader(data)
	response, err := postFromReader(URL, reader, mime)

	check(err)

	return response, err
}

// PostJSONFile takes a file path and reads it in then sends it to server via POST
func PostJSONFile(URL string, jsonFilePath string) (string, error) {
	dat, err := ioutil.ReadFile(jsonFilePath)

	check(err)

	r, serr := PostByteArray(URL, dat, "application/json")

	check(serr)

	return r, serr
}

// PostJSONString sends a JSON string to another server via POST
func PostJSONString(URL string, json string) (string, error) {
	b := []byte(json)

	r, err := PostByteArray(URL, b, "application/json")

	check(err)

	return r, err
}
