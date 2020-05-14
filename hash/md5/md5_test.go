package md5_test

import (
	"testing"

	"github.com/danielfmelo/myhttphash/hash/md5"
)

func TestCreateMD5(t *testing.T) {
	m := md5.New()
	data := []byte("to.be.hashed")
	expectedMd5 := "ff0c1b127f5e1170afb5d053b295d561"
	result := m.CreateHash(data)
	if result != expectedMd5 {
		t.Errorf("expected %s but got %s", expectedMd5, result)
	}
}
