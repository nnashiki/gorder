package test

import (
	"github.com/magiconair/properties/assert"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 開始処理
	log.Print("setup")
	// パッケージ内のテストの実行
	code := m.Run()
	// 終了処理
	log.Print("tear-down")
	// テストの終了コードで exit
	os.Exit(code)
}

func TestHello(t *testing.T) {
	assert.Equal(t, true,true)
}

