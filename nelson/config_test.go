package nelson

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	fixture := `
---
endpoint: http://nelson.local
session:
  token: sZoqgvx5h019nxgWwDPDdQMAAENKsPf2m/9261SAtS1FK
  expires_at: 1566976604632
`
	ioutil.WriteFile("/tmp/nelson.yml", []byte(fixture), os.FileMode(0666))

	var client Nelson
	client.configPath = "/tmp/nelson.yml"
	client.NelsonConfig = &NelsonConfig{
		Endpoint: "",
		Session: &NelsonSession{
			Token:     "",
			ExpiresAt: 0,
		},
	}
	if err := client.readConfig(); err != nil {
		t.Error(err)
	}
}

func TestWriteConfig(t *testing.T) {
	fixture := `
endpoint: https://nelson.local
session:
  token: ""
  expires_at: 0
`
	client, _ := CreateNelson("https://nelson.local", "v1", "/tmp/nelson.yml")
	if err := client.writeConfig(); err != nil {
		t.Error(err)
	}

	b := []byte(fixture)
	read, err := ioutil.ReadFile(client.configPath)
	if err != nil {
		t.Error(err)
	}

	if bytes.Compare(b, read) > 0 {
		t.Fail()
	}
}
