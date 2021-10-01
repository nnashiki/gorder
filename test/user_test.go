package test

import (
	"bytes"
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"goder/models"
	"net/http"
	"testing"
)

func TestUserPostの200(t *testing.T) {

	testServer := ServerInit()
	defer testServer.Close()

	// JsonHandler のテスト
	params := map[string]interface{}{"name": "nashiki", "email": "nashiki@example.com"}
	b, err := json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	response, err := http.Post(testServer.URL+"/api/user",
		"application/json", bytes.NewBuffer(b))
	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("invalid response: %v", response)
	}
	resp := models.User{}
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
	response.Body.Close()
	assert.Equal(t, response.StatusCode, http.StatusOK)
	assert.Equal(t, resp.Name, "nashiki")
}
