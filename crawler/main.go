package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
)

func main()  {
	//url := "http://www.zhenai.com/zhenghun";
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
/*	data,err :=fetcher.GetContent("https://album.zhenai.com/u/1972565883");
	if err != nil{
		fmt.Printf("错误")
	}
	fmt.Printf("%s",data);*/
}





