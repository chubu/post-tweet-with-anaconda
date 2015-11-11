package main

import (
	"filepath"
	"github.com/ChimeraCoder/anaconda"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"math/rand"
	"time"
)

type Token struct {
	Consumer, AccessToken map[string]Options
}

type Options struct {
	Token, Secret string
}

type Statuses struct {
	Data []string
}

func main() {
	// read yaml
	filename, _ := filepath.Abs("./yaml/token.yml")
	yaml, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var token Token
	err = yaml.Unmarshal(yaml, &token)
	if err != nil {
		panic(err)
	}

	// read statuses
	filename, _ := filepath.Abs("./yaml/data.yml")
	yaml, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var statuses Statuses
	err = yaml.Unmarshal(yaml, &statuses)
	if err != nil {
		panic(err)
	}

	// create api instance
	anaconda.SetConsumerKey(token.Consumer.Token)
	anaconda.SetConsumerSecret(token.Consumer.Secret)
	api := anaconda.NewTwitterApi(token.AccessToken.Token, token.AccessToken.Secret)

	// post tweet
	current := time.Now().Unix()
	rand.Seed(current)
	num := rand.Intn(len(statuses))
	api.PostTweet("今の運勢: "+statuses[num], nil)
}
