package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"goder/models"
	"goder/mysql"
	"goder/router"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"context"
)

func TestMain(m *testing.M) {
	// 開始処理
	log.Print("setup")

	mysql.New()
	ctx := context.Background()
	ClearData(ctx)

	LoadData(ctx)
	fmt.Println(models.Users().One(ctx,mysql.DB))



	// パッケージ内のテストの実行
	code := m.Run()

	// 終了処理
	log.Print("tear-down")
	//ClearData(ctx)

	// テストの終了コードで exit
	os.Exit(code)
}

type UsersFixture struct {
	TestCases []models.User `json:"test_cases"`
}

func LoadUsers(ctx context.Context) error {
	b, err := ioutil.ReadFile("testdata/fixture.json")
	if err != nil {
		log.Fatal(err)
	}
	f := new(UsersFixture)
	if err := json.Unmarshal(b, f); err != nil {
		log.Fatal(err)
	}
	for _, user  := range f.TestCases{
		user.Insert(ctx,mysql.DB,boil.Infer())
	}
	return nil
}

func LoadData(ctx context.Context) error {
	LoadUsers(ctx)
	return nil
}

func ClearData(ctx context.Context) error {
	ClearUsers(ctx)
	return nil
}

func ClearUsers(ctx context.Context) error {
	users, err := models.Users().All(ctx,mysql.DB)
	if err == nil {
		users.DeleteAll(ctx,mysql.DB)
		return nil
	}
	return err
}

func TestHello(t *testing.T) {
	assert.Equal(t, true,true)
}

func ServerInit() *httptest.Server {
	// test server を起動
	routers := router.Init()
	return httptest.NewServer(routers)
}

func TestUserPost(t *testing.T) {

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
