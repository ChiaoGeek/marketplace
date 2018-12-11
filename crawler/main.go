package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ChiaoGeek/marketplace/crawler/downloader"
	"github.com/ChiaoGeek/marketplace/crawler/parse"
	"github.com/ChiaoGeek/marketplace/crawler/reqres"
	"golang.org/x/net/html"
	"os"
	"strings"
	"sync"
)
var global int
var lock = new(sync.Mutex)
func getdetail(mrequest *reqres.Mrequest, c chan map[string]*[]*reqres.Detailres) {
	resp, err := downloader.Getwebpage(mrequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	node, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	res := make([]*reqres.Detailres, 0)
	m := make(map[string]*[]*reqres.Detailres)
	parse.Getdetail(node, mrequest.Url, &res)

	m[mrequest.Url] = &res

	c <- m
	lock.Lock()
	global--
	if(global == 0) {
		close(c)
	}

	lock.Unlock()

}

func main()  {
	mrequest := &reqres.Mrequest{
		Method: "GET",
		Url: "https://github.com/marketplace",
	}
	c := make(chan map[string]*[]*reqres.Detailres)
	resp, err := downloader.Getwebpage(mrequest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	node, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	linkres := make([]string, 0)
	parse.Getcatogories(node, mrequest.Url, &linkres)
	global = len(linkres)
	go func() {
		for _,x := range linkres {
			tem := reqres.Mrequest{
				Url: x,
				Method:"GET",
			}
			go getdetail(&tem, c)
		}
	}()
	resbytes := bytes.Buffer{}
	resbytes.WriteString("{")
	finalresslice := make([]reqres.Jsonres, 0)
	for x := range c  {
		finalresslice = append(finalresslice, maptojsonstring(x))
	}

	str, err := json.Marshal(reqres.Bigjson{
		finalresslice,
	})

	fmt.Println(string(str))
	file, err := os.Create("marketplace.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	file.WriteString(string(str))
	file.Close()
}


func maptojsonstring(m map[string]*[]*reqres.Detailres) reqres.Jsonres{

	jsonslice := make([]reqres.Jsonresslice, 0)
	var jsonkey string
	for key, val := range m {
		jsonkey = getclassbyurl(key)
		for _,v := range *val{
			jsonslice = append(jsonslice, reqres.Jsonresslice{
				v.Name,
				v.Url,
				v.Description,
				getclassbyurl(key),
			})

		}
	}
	return reqres.Jsonres{
		jsonkey,
		jsonslice,
	}

}

func getclassbyurl(url string) string{
	strslice := strings.Split(url, "/")
	return strslice[len(strslice) - 1]
}
