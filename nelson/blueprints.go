package nelson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type BlueprintResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Revision    string `json:"revision"`
	Sha256      string `json:"sha256"`
	Template    string `json:"template"`
	CreatedAt   int64  `json:"created_at"`
}

type CreateBlueprintRequest struct {
	// Example: gpu-general-deployment
	Name string `json:"name"`

	// Example: a blueprint for intensive graphics consumption
	Description string `json:"description"`

	// Sha256 sum of the template for versioning purposes
	Sha256 string `json:"sha256"`

	// Base64 Encoding of the blueprint template
	Template string `json:"template"`
}

func (n *Nelson) CreateBlueprint(cbr CreateBlueprintRequest) error {
	requestBody, err := json.Marshal(cbr)
	if err != nil {
		return err
	}

	endpoint := []string{n.NelsonConfig.Endpoint, n.apiVersion, "blueprints"}
	req, _ := http.NewRequest(http.MethodPost, strings.Join(endpoint, "/"), bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", LibraryVersion())

	response, err := n.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	var blueprintResponse BlueprintResponse
	json.NewDecoder(response.Body).Decode(&blueprintResponse)

	if response.StatusCode != 200 {
		bodyBytes, _ := ioutil.ReadAll(response.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("nelson server error: %s", bodyString)
	}

	return nil
}

func (n *Nelson) GetBlueprints() ([]*BlueprintResponse, error) {
	var b []byte
	requestBody := bytes.NewBuffer(b)

	endpoint := []string{n.NelsonConfig.Endpoint, n.apiVersion, "blueprints"}
	req, _ := http.NewRequest(http.MethodGet, strings.Join(endpoint, "/"), requestBody)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", LibraryVersion())

	response, err := n.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("nelson server error")
	}

	var list []*BlueprintResponse
	json.NewDecoder(response.Body).Decode(&list)
	return list, nil
}

func (n *Nelson) GetBlueprint(blueprintName string) (*BlueprintResponse, error) {
	blueprints, err := n.GetBlueprints()
	if err != nil {
		return nil, err
	}

	for _, blueprint := range blueprints {
		if blueprint.Name == blueprintName {
			return blueprint, nil
		}
	}

	return nil, fmt.Errorf("blueprint not found")
}
