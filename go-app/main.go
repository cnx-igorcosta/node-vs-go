package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	filepath = "./arquivo-go.txt"
)

func main() {
	http.HandleFunc("/", handle)
	fmt.Fprintf(os.Stdout, "listening on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handle(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadFile(filepath)
	checkError(err, w)

	rs, err := generateRandomString()
	checkError(err, w)

	content := append(data, rs+"\n"...)

	err = ioutil.WriteFile(filepath, content, 0644)
	checkError(err, w)

	// err = generateRandomFile(content)
	// checkError(err, w)

	fmt.Fprintf(w, string(content))
}

func checkError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func generateRandomString() (string, error) {
	size := 32
	rb := make([]byte, size)

	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}

	rs := base64.URLEncoding.EncodeToString(rb)

	return rs, nil
}

func generateRandomFile(content []byte) error {
	rs, err := generateRandomString()
	if err != nil {
		return err
	}

	ioutil.WriteFile("./contents/"+rs+".txt", content, 0644)

	return nil
}
