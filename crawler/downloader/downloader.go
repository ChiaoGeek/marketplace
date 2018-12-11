package downloader

import (
	"crypto/tls"
	"fmt"
	"github.com/ChiaoGeek/marketplace/crawler/errorhandling"
	"github.com/ChiaoGeek/marketplace/crawler/reqres"
	"net/http"
)

func Getwebpage(mrequest *reqres.Mrequest) (*http.Response, error){
	if mrequest == nil{
		return nil, errorhandling.Newerror("Mrequest cannot be empty")
	}
	var resp *http.Response
	err := errorhandling.Newerror("Unexpected error")
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsconfig,
	}
	client := http.Client{Transport:transport}
	newrequest,err := http.NewRequest(mrequest.Method, mrequest.Url, nil)
	//newrequest.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	if err != nil {
		return nil, err
	}
	resp, err = client.Do(newrequest)
	if err != nil{
		fmt.Println(err.Error())
		return resp, err
	}
	//fmt.Println(mrequest.Url)
	//fmt.Println(mrequest.Method)
	return resp, err
}