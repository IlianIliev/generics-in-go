package api

import (
	"encoding/json"
	"net/http"
)

type user struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type userResponse struct {
	Data   []user `json:"data"`
	Paging struct {
		Total int `json:"total"`
		Limit int `json:"limit"`
	} `json:"paging"`
}

type company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type companyResponse struct {
	Data   []company `json:"data"`
	Paging struct {
		Total int `json:"total"`
		Limit int `json:"limit"`
	} `json:"paging"`
}

func getUsers() ([]user, error) {
	resp, err := http.Get("http://127.0.0.1:8000/users")
	if err != nil {
		return nil, err
	}

	var result userResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

func getCompanies() ([]company, error) {
	resp, err := http.Get("http://127.0.0.1:8000/companies")
	if err != nil {
		return nil, err
	}

	var result companyResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}
