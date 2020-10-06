package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Fetch(url string) ([]byte, error) {
	//打开页面
	//resp, err := http.Get(url)


	//client := &http.Client{}
	//req, err := http.NewRequest(http.MethodGet, url, nil)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	//resp, err := client.Do(req)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//defer resp.Body.Close()

	client := &http.Client{}
	newUrl := strings.Replace(url, "http://", "https://", 1)
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		panic(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	cookie1 := "sid=b24c91bc-8b2d-4998-a761-2f8e5b450a3e; ec=Gjl2gxaR-1599928119029-6994399c306ce1287851462; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1599928134,1601910421; FSSBBIl1UgzbN7NO=5fx7_ycM4DFYDKVjTjEIXuFEi.Jem5ZC4.3OwsiVJdu5VXHIR.PeyN7cgJ92QH0.xsJ7owquAVAY8ITlTFGWtVa; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1601987790; FSSBBIl1UgzbN7NP=5U6S9uTeZRWgqqqm0d8vJhGjIjWTzEwihCAGgvKgQqHpwHBsBT0Rhlq3Ta4vEsy2n7Wx8s4o1RZGScXgrCxBK_JgwnkrsSs6ECVq0gbCi_zipbfuupFiBPMpV6r6JASdLCaKdhBlqqqdLDokxYzVb2GVgqTWReGGRamn6Ae4r7QnN15gVWdfa91yxdk6FDesYVA1sIvgx1GSq3torWy1pnAfWs2rvieZM811uOYxEjpXT4MoEntB17yhQdYUds9z2mfoMoC5uo9STSB4CwPuRxE; _efmdata=glR%2Ff61R6Z0PdhKKzZ8s87BAy%2FpTMm1kyl0TQciGduix3%2FKybX%2Bx78ipTgtevGpDSMrOIMyQ2StJ1%2FVOFfvlcZ9aQXiw1RZhPpFA9YmwFrQ%3D; _exid=uYoPnnWzQlemyACM6Dkrc0NKOEytGpYmQS5jjypT3LTIrfDPIbJUnLd3Z%2BN%2BRh1Reh94VUfJ0wWTZ6NzUZlFRQ%3D%3D"
	req.Header.Add("cookie", cookie1)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	//
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//获取页面内容，就是查看源代码看到的内容
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

//自动获取网站编码，因为网站不一定时gbk
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	//如果使用resp.Body直接读，读完之后无法继续再读所以使用bufio.NewReader()
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		//如果出错还是希望继续读，所以返回一个默认的encoding
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}