package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetContent(url string )([]byte,error) {
	newUrl := strings.Replace(url, "http://", "https://", 1)
	//url := "http://www.zhenai.com/zhenghun";
	client := &http.Client{}
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	/*req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
	*/
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.104 Safari/537.36")
	cookie := "sid=2ece965e-a8ce-4a9e-90bb-e0263bec350d; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1612599250; ec=EKvY4mZ0-1612599730549-d8a32692b3061-1945136429; spliterabparams=1612599746950:4552482931751395180; FSSBBIl1UgzbN7NO=5Uf61suKMNkNnlFicvkwbReVWoYGVzS4ZhR0ApcoipKAI3sIkKF4EXyMgA4y3dyMhbWoeiMdMRdiE0m09wwYwmA; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1612713581; FSSBBIl1UgzbN7NP=53cqiwbqGkbaqqqm6IsPcvGp2tkh38KsWDCrO4XqX3QR.4aP19LxhOMjs0SNWwcUK49Iqe8gsuBnyMp1E1HteXg0Vf7D4.rLF6CwyfvbFshziqwy0PVLPfQ1okWI48euNM3zxiatxkRSGx9lvrom7uVq7p2ILQMRO5HY0Sq8JKhDeRgfLXYpqb0ebCSNzHUEWf_KKFuxNOBtgOEpYUszubol2CBP_CUVvo_s.NM2Vlq.qFN0adZ2MDBRQKvhc_0sr0; _efmdata=KE7gVLRzDmlsfQMDjqpd/4cJK/D7eRTr95/qogr3xdlKftukmnTeJtxFnWxPk3axCjUOFv5Lx9kBOFi66kkTYu9G47bKDmzS+XhHUgOS4Ac=; _exid=tNrDRwceHS0SBUkFBTDbYXy8r3m4LTNGLsyFyqB0vZzefcYPtk4Iov1qLoNbS9Dkomw50ePn98sPe4OoOR6Lug=="
	req.Header.Add("cookie", cookie)
	resp, err := client.Do(req)

	if resp.StatusCode != http.StatusOK{
		//log.Fatal("状态码出错%s",resp.StatusCode);
		fmt.Println(resp.StatusCode);
		return nil,nil
	}
	defer resp.Body.Close()

	//进行网页编码猜测
	e := determineEncoding(resp.Body)

	//转化成UT8编码
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)

	if err != nil{
		return  nil,fmt.Errorf(`wrong status %s`,"转码错误")
	}

	return all,err
}

/*
  去判断网页是什么样的格式
*/
func determineEncoding(r io.Reader)  encoding.Encoding{
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil{
		return unicode.UTF8
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e
}

