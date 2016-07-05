package HttpSender_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/darthmolen/GoHttpUtils"
)

var srv *httptest.Server
var testFile string
var mime string

func TestMain(m *testing.M) {
	startTestServer("TestFile.json")
	os.Exit(m.Run())
}

func startTestServer(tFile string) {
	testFile = tFile
	mime = "application/json"
	srv = httptest.NewServer(http.HandlerFunc(parseTest))
}

func readFileToString(location string) (string, error) {
	dat, err := ioutil.ReadFile(testFile)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func parseTest(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	resp, _ := ioutil.ReadAll(req.Body)

	ret := string(resp)
	fmt.Fprint(rw, ret)
}

func TestPostJSONFile(t *testing.T) {

	resp, err := HttpSender.PostJSONFile(srv.URL, testFile)

	if err != nil {
		t.Error("Error connecting to server from file post", err)
	}

	dat, err := readFileToString("TestFile.json")

	if err != nil {
		t.Error("Unable to open Test File", err)
	}

	if dat != resp {
		t.Log(string(dat))
		t.Log(resp)
		t.Error("File Post Failure")

	}

}

func TestPostJSONString(t *testing.T) {

	s, ferr := readFileToString(testFile)

	if ferr != nil {
		t.Error("Unable to open Test File", ferr)
	}

	resp, err := HttpSender.PostJSONString(srv.URL, s)

	if err != nil {
		t.Error("Error connecting to server from string post", err)
	}

	if s != resp {
		t.Log(s)
		t.Log(resp)
		t.Error("String Post Failure")
	}
}

func TestPostByteArray(t *testing.T) {
	dat, ferr := ioutil.ReadFile(testFile)

	if ferr != nil {
		t.Error("Unable to open Test File", ferr)
	}

	resp, err := HttpSender.PostByteArray(srv.URL, dat, mime)

	if err != nil {
		t.Error("Error connecting to server from byte post", err)
	}

	if string(dat) != resp {
		t.Log(string(dat))
		t.Log(resp)
		t.Error("Byte Post Failure")
	}
}
