package parser
import (
	"learngo/crawler/engine"
	"regexp"
)

const cityRegexp = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
func ParseCity(contents []byte) engine.ParseResult{
	//首先创建一个正则表达式,如果是``表示不会对字符串里面的内容进行什么编译转码
	matchs :=regexp.MustCompile(cityRegexp);
	lists :=matchs.FindAllSubmatch(contents,-1);
	result := engine.ParseResult{};

	for _,v := range lists{

		result.Items = append(result.Items,v[2]) // 将城市进行追加

		result.Request = append(result.Request,engine.Request{
			Url:string(v[1]),
			ParseFunc: ParseProfile,//解析单人信息
		})
	}

	return result
}