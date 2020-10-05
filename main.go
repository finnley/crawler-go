package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	//打开页面
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//获取页面内容，就是查看源代码看到的内容
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	////不需要头部，只需要body
	//all, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}

	//处理乱码 这里还是乱码
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)
	//if err != nil {
	//	panic(err)
	//}

	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", all)
	//提取城市和url
	printCityList(all)
}

//自动获取网站编码，因为网站不一定时gbk
func determineEncoding(r io.Reader) encoding.Encoding {
	//如果使用resp.Body直接读，读完之后无法继续再读所以使用bufio.NewReader()
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte)  {
	//re := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`)
	////-1表示所有匹配
	//matches := re.FindAll(contents, -1)
	//fmt.Printf("%s\n", matches)
	//for _, m := range matches {
	//	fmt.Printf("%s\n", m)
	//}
	//fmt.Printf("Matches found: %d\n", len(matches))

	//上面提取内容: <a href="http://www.zhenai.com/zhenghun/tongliang" data-v-2cb5b6a2>铜梁</a>
	//现在只需要提取城市和url,改造如下

	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//for _, subMatch := range m {
		//	fmt.Printf("%s ", subMatch)
		//}
		//fmt.Println()

		//上面输出格式: <a href="http://www.zhenai.com/zhenghun/yuxi" data-v-2cb5b6a2>玉溪</a> http://www.zhenai.com/zhenghun/yuxi 玉溪
		//不需要全部打印，只需要城市和url

		fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
}