package nelson

import "testing"

func TestCreateNelson(t *testing.T) {
	client, err := CreateNelson("https://nelson.local", "v1", "")

	if client.NelsonConfig.Endpoint != "https://nelson.local" {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
