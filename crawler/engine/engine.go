package engine

import (
	"fmt"
	"learngo/crawler/fetcher"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests)>0{
		r := requests[0]
		requests = requests[1:]
		body, err := fetcher.GetContent(r.Url)
		if err != nil{
			continue
		}

		parseResult := r.ParseFunc(body)
		requests = append(requests,parseResult.Request ...)//将解析的得出的地址一个个再塞进请求地址里面

		for k,v := range parseResult.Items{
			fmt.Printf("名称：%s",v);
			fmt.Printf("地址：%s",parseResult.Request[k].Url)
			fmt.Println();
		}
	}
}