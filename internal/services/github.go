package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchGitHubData(username string) ([]map[string]interface{}, error) {

	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username);

	resp, err := http.Get(url);

	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %v", err)
	}

	defer resp.Body.Close();

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na resposta: %s", resp.Status)
	}

	var data []map[string]interface{}
	
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a resposta: %v", err)
	}

	return data, nil
}
