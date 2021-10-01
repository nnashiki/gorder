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
	b, err := json.Marshal(map[string]interface{}{"name": "nashiki", "email":"nashiki@example.com"})
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.Post(testServer.URL+"/api/user", "application/json", bytes.NewBuffer(b))
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("invalid response: %v", res)
	}
	resp := models.User{}
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
	res.Body.Close()

	assert.Equal(t, resp.Name,"nashiki")
}
