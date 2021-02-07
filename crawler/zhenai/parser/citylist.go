package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const cityListRegexp = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult{
	//首先创建一个正则表达式,如果是``表示不会对字符串里面的内容进行什么编译转码
	matchs :=regexp.MustCompile(cityListRegexp);
	lists :=matchs.FindAllSubmatch(contents,-1);
	result := engine.ParseResult{};

	//将匹配到的城市列表放入result里面
	for _,v := range lists{

		result.Items = append(result.Items,v[2]) // 将城市进行追加

		result.Request = append(result.Request,engine.Request{
			Url:string(v[1]),
			ParseFunc: ParseCity,
		})
	}

	return result
}
