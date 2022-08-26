// Calling GitHub API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type User struct {
	Login    string
	Name     string
	NumRepos string // `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	resp, err := http.Get("https://api.github.com/users/" + url.PathEscape(login))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(resp.Status)
	}
	// io.Copy(os.Stdout, resp.Body)
	dec := json.NewDecoder(resp.Body)
	var user User
	if err := dec.Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func main() {
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)
}
