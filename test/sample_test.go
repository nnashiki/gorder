package test

import (
	"github.com/magiconair/properties/assert"
	_ "github.com/volatiletech/sqlboiler/v4/boil"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"goder/models"
	"goder/mysql"
	"log"
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

	// パッケージ内のテストの実行
	code := m.Run()

	// 終了処理
	log.Print("tear-down")
	ClearData(ctx)

	// テストの終了コードで exit
	os.Exit(code)
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

