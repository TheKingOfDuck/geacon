package config

import (
	"time"
)

var (
	RsaPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
-----END PUBLIC KEY-----`)
	RsaPrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
-----END PRIVATE KEY-----`)

	C2        = "IP:PORT"
	plainHTTP = "https://"
	sslHTTP   = "https://"
	GetUrl    = plainHTTP + C2 + "/api/home"
	PostUrl   = plainHTTP + C2 + "/api/index"
	WaitTime  = 15000 * time.Millisecond

	IV        = []byte("abcdefghijklmnop")
	GlobalKey []byte
	AesKey    []byte
	HmacKey   []byte
	Counter   = 0
)

const (
	DebugMode = true
)
