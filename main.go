package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

const (
	TOKEN_YAML = "/home/ec2-user/golang/yaml/token.yml"
	DATA_YAML  = "/home/ec2-user/golang/yaml/data.yml"
)

type Token struct {
	Consumer, Access_token Options
}

type Options struct {
	Token, Secret string
}

type Statuses struct {
	Data []string
}

func getToken() (Token, error) {
	var token Token
	filename, _ := filepath.Abs(TOKEN_YAML)
	yml, err := ioutil.ReadFile(filename)
	if err != nil {
		return token, err
	}

	err = yaml.Unmarshal(yml, &token)
	if err != nil {
		return token, err
	}

	return token, nil
}

func getStatuses() (Statuses, error) {
	var statuses Statuses
	filename, _ := filepath.Abs(DATA_YAML)
	yml, err := ioutil.ReadFile(filename)
	if err != nil {
		return statuses, err
	}

	err = yaml.Unmarshal(yml, &statuses)
	if err != nil {
		return statuses, err
	}

	return statuses, nil
}

func createStatus(statuses Statuses) string {
	current := time.Now()
	currentUnixtime := current.Unix()
	rand.Seed(currentUnixtime)
	num := rand.Intn(len(statuses.Data))
	timeString := fmt.Sprintf(
		"%4d年%02d月%02d日 %02d:%02d の運勢: %s",
		current.Year(),
		current.Month(),
		current.Day(),
		current.Hour(),
		current.Minute(),
		statuses.Data[num],
	)

	return timeString
}

func main() {
	token, err := getToken()
	if err != nil {
		panic(err)
	}

	statuses, err := getStatuses()
	if err != nil {
		panic(err)
	}

	anaconda.SetConsumerKey(token.Consumer.Token)
	anaconda.SetConsumerSecret(token.Consumer.Secret)
	api := anaconda.NewTwitterApi(token.Access_token.Token, token.Access_token.Secret)
	api.PostTweet(createStatus(statuses), nil)
}
