package routes

import (
	"awesomeProject/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	testBody, err := json.Marshal(map[string]string{
		"email":    "test@gmail.com",
		"password": "test",
	})
	if err != nil {
		t.Skip("Failed to marshal test body")
	}
	res, err := http.Post(
		"http://localhost:3000/accounts",
		"application/json",
		bytes.NewReader(testBody))
	if err != nil {
		t.Skip("Failed to perform http request")
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	t.Log(string(bodyBytes))
}

func TestGetAccounts(t *testing.T) {
	res, err := http.Get("http://localhost:3000/accounts")
	if err != nil {
		return
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %s", err.Error())
	}

	var accounts []models.Account
	err = json.Unmarshal(bodyBytes, &accounts)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %s", err.Error())
	}

	t.Logf("%d accounts found", len(accounts))
}

func TestGetAccount(t *testing.T) {
	// get all accounts
	res, err := http.Get("http://localhost:3000/accounts")
	if err != nil {
		t.Skip("Failed to perform http request")
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %s", err.Error())
	}

	var accounts []models.Account
	err = json.Unmarshal(bodyBytes, &accounts)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %s", err.Error())
	}

	if len(accounts) == 0 {
		t.Skip("No accounts found")
	}

	// get the first account and get it specifically
	res, err = http.Get("http://localhost:3000/accounts/" + strconv.Itoa(int(accounts[0].ID)))
	if err != nil {
		t.Skip("Failed to perform http request")
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	bodyBytes, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %s", err.Error())
	}

	var account models.Account
	err = json.Unmarshal(bodyBytes, &account)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %s", err.Error())
	}

	t.Logf("Account found: %s", account.Email)
}

func TestUpdateAccount(t *testing.T) {
	// Get the first account
	res, err := http.Get("http://localhost:3000/accounts")
	if err != nil {
		t.Skip("Failed to perform http request")
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d: %s", res.StatusCode, string(bodyBytes))
	}

	var accounts []models.Account
	err = json.Unmarshal(bodyBytes, &accounts)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %s", err.Error())
	}

	if len(accounts) == 0 {
		t.Skip("No accounts found")
	}

	// Update the first account
	testBody, err := json.Marshal(map[string]string{
		"Email":    "test1@gmail.com",
		"Password": "test",
	})
	if err != nil {
		t.Skip("Failed to marshal test body")
	}

	req, err := http.NewRequest("PUT", "http://localhost:3000/accounts/"+strconv.Itoa(int(accounts[0].ID)), bytes.NewReader(testBody))
	if err != nil {
		t.Skip("Failed to create request")
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Skip("Failed to perform http request")
	}

	bodyBytes, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d: %s", res.StatusCode, string(bodyBytes))
	}

	var account models.Account
	err = json.Unmarshal(bodyBytes, &account)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %s", err.Error())
	}

	t.Logf("Account updated: %s", account.Email)
}

func TestDeleteAccount(t *testing.T) {
	// Get the first account
	res, err := http.Get("http://localhost:3000/accounts")
	if err != nil {
		t.Skip("Failed to perform http request")
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %s", err.Error())
	}

	var accounts []models.Account
	err = json.Unmarshal(bodyBytes, &accounts)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %s", err.Error())
	}

	if len(accounts) == 0 {
		t.Skip("No accounts found")
	}

	// Delete the first account
	req, err := http.NewRequest("DELETE", "http://localhost:3000/accounts/"+strconv.Itoa(int(accounts[0].ID)), nil)
	if err != nil {
		t.Skip("Failed to create http request")
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Skip("Failed to perform http request")
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}

	bodyBytes, err = io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %s", err.Error())
	}

	t.Log(string(bodyBytes))
}
