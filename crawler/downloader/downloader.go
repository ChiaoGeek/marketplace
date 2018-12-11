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

	if err != nil {
		return nil, err
	}
	resp, err = client.Do(newrequest)
	if err != nil{
		fmt.Println(err.Error())
		return resp, err
	}
	return resp, err
}