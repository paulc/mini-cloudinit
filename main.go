package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Hostname      string `yaml:"hostname"`
	InstanceId    int    `yaml:"instance-id"`
	LocalIPV4     string `yaml:"local-ipv4"`
	PublicIPV4    string `yaml:"public-ipv4"`
	NetworkConfig struct {
		Config  []map[string]interface{}
		Version int
	} `yaml:"network-config"`
	NetworkSysconfig string   `yaml:"network-sysconfig"`
	PublicKeys       []string `yaml:"public-keys"`
	VendorData       string   `yaml:"vendor_data"`
}

func main() {

	resp, err := http.Get("http://127.0.0.1:8000/meta.yml")
	if err != nil {
		log.Fatalf("HTTP GET failed: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var meta Metadata
	yaml.Unmarshal(body, &meta)
	spew.Dump(meta)
}
