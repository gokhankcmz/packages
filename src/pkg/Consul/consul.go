package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

type Client struct {
	Client *api.Client
	Url    string
	Token  string
}

func Init(url string) Client {
	conf := api.DefaultConfig()
	conf.Address = url
	consulClient, err := api.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}
	return Client{Client: consulClient, Url: url}
}

func (c Client) Set(key, value string) error {
	kv := c.Client.KV()
	p := &api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) Get(key string) (string, error) {
	kv := c.Client.KV()

	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return "", err
	}

	return string(pair.Value), nil
}

func (c Client) GetConfigs(relativeUrl string) map[string]string {
	kv, _, err := c.Client.KV().List(relativeUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	result := map[string]string{}
	for _, v := range kv {
		result[v.Key] = string(v.Value)
	}

	return result
}
