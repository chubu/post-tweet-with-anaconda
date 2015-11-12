package main

import (
	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

const (
	TOKEN_YAML = "./yaml/token.yml"
	DATA_YAML  = "./yaml/data.yml"
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

func main() {
	// read yaml
	filename, _ := filepath.Abs(TOKEN_YAML)
	yml, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var token Token
	err = yaml.Unmarshal(yml, &token)
	if err != nil {
		panic(err)
	}

	// read statuses
	filename, _ = filepath.Abs(DATA_YAML)
	yml, err = ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var statuses Statuses
	err = yaml.Unmarshal(yml, &statuses)
	if err != nil {
		panic(err)
	}

	// create api instance
	anaconda.SetConsumerKey(token.Consumer.Token)
	anaconda.SetConsumerSecret(token.Consumer.Secret)
	api := anaconda.NewTwitterApi(token.Access_token.Token, token.Access_token.Secret)

	// post tweet
	current := time.Now().Unix()
	rand.Seed(current)
	num := rand.Intn(len(statuses.Data))
	api.PostTweet("今の運勢: "+statuses.Data[num], nil)
}
