package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	filepath = "./arquivo-go.txt"
)

func handler(w http.ResponseWriter, req *http.Request) {
	// Lê os dados de um arquivo específico,
	// concatena aos dados uma string randômica
	// e salva o arquivo com os novos dados
	data, err := ioutil.ReadFile(filepath)
	checkError(err, w)
	rs, err := generateRandomString(256)
	checkError(err, w)
	content := append(data, fmt.Sprintf("%s\n", rs)...)
	err = ioutil.WriteFile(filepath, content, 0644)
	checkError(err, w)

	// Usa os dados randômicos gerados anteriormente
	// para criar um novo arquivo numa pasta específica
	// O nome do arquivo é gerado randômicamente
	// err = generateRandomFile(content)
	// checkError(err, w)

	// Chama uma api http com dados mock,
	// escreve no stdout as informações retornadas
	// logMockApiresponse()

	// Retorna os dados no response body
	fmt.Fprintf(w, string(content))
}

func checkError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func generateRandomString(size int) (string, error) {
	rb := make([]byte, size)
	_, err := rand.Read(rb)
	if err != nil {
		return "", err
	}

	rs := base64.URLEncoding.EncodeToString(rb)

	return rs, nil
}

func generateRandomFile(content []byte) error {
	rs, err := generateRandomString(10)
	if err != nil {
		return err
	}

	ioutil.WriteFile(fmt.Sprintf("./contents/%s.txt", rs), content, 0644)

	return nil
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type users []user

func logMockApiresponse() error {
	url := "http://5d879522cd71160014aaeac7.mockapi.io/api/v1/users"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	users := new(users)
	json.NewDecoder(resp.Body).Decode(users)

	for _, u := range *users {
		fmt.Fprintf(os.Stdout, "\nid: %s \nnome: %s\n", u.ID, u.Name)
	}

	return nil
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Fprintf(os.Stdout, "listening on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
