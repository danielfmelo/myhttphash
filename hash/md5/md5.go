package md5

import (
	"crypto/md5"
	"encoding/hex"
)

type Md5 struct{}

func (m *Md5) CreateHash(data []byte) string {
	result := md5.Sum(data)
	return hex.EncodeToString(result[:])
}

func New() *Md5 {
	return &Md5{}
}
