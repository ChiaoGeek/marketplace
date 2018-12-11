package parse

import (
	"github.com/ChiaoGeek/marketplace/crawler/reqres"
	"golang.org/x/net/html"
	"net/url"
)

func Getcatogories(node *html.Node, originurl string, res *[]string){
	if node == nil {
		return
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _,a := range node.Attr  {
			if a.Key == "class" && a.Val == "filter-item py-1"{
				for _,b := range node.Attr {
					if b.Key == "href" {
						*res = append(*res, Fixurl(b.Val, originurl))
					}
				}
			}
		}
	}
	for n := node.FirstChild; n != nil;  n = n.NextSibling{
		Getcatogories(n, originurl, res)
	}
}

func Getdetail(node *html.Node, originurl string, res *[]*reqres.Detailres) {
	if node == nil {
		return
	}
	//c := make(chan string, 3)
	re := make([]*html.Node, 0)
	Getdetaillink(node, &re)

	for _, x := range re {
		var u, n, p string

		if x != nil {
			for _,b := range x.Attr {
				if b.Key == "href" {
					u = Fixurl(b.Val, originurl)
				}
			}
			n = x.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
			p = x.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
			//n = x.Data
			//p = x.Data
			singler := reqres.Detailres{
				n,
				u,
				p,
			}
			//fmt.Printf("%+v\n", singler)
			*res = append(*res, &singler)
		}
		//u = "test"
		//n = "t"
		//n = x.FirstChild.NextSibling.FirstChild.FirstChild.Data
		//p = "e"
		//n = x.FirstChild.NextSibling.FirstChild.FirstChild.Data



	}
	//singler := reqres.Detailres{
	//	n,
	//	u,
	//	p,
	//}

	//fmt.Printf("%+v\n", singler)

 	//*res = append(*res, &singler)
}

func Getdetaillink(node *html.Node,  res *[]*html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {

		for _,a := range node.Attr  {
			if a.Key == "class" && a.Val == "col-md-6 mb-4 d-flex no-underline"{
				//for _,b := range node.Attr {
				*res = append(*res, node)
			}
		}
	}
	for n := node.FirstChild; n != nil;  n = n.NextSibling{
		Getdetaillink(n,  res)
	}
}


func Fixurl(currenturi, originurl string) string{
	uri, err := url.Parse(currenturi)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(originurl)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}