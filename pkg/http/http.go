package http

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func HttpGet(url string){

}

func HttpPost(url string,data interface{}){

}

func HttpHandle(url string,hreader map[string]string) []byte{
	client := &http.Client{}

	req,_ := http.NewRequest(http.MethodGet,url,nil)
	for key,val := range hreader {
		req.Header.Add(key,val)
	}
	//req.Header.Add("Accept","application/json")
	//req.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	//req.Header.Add("X-Requested-With","XMLHttpRequest")
	//req.Header.Add("Referer","https://m.douban.com/movie/subject/26588308/comments?sort=new_score&start=50")

	resp,err := client.Do(req)
	if err != nil {
		//logging.Error(url,"Request error:",err)
		fmt.Errorf("Url [%s] Request error:%v\n",url,err)
		return nil
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//logging.Error(url,"ReadAll error:",err)
		fmt.Errorf("Url [%s] ReadAll error:%v\n",url,err)
		return nil
	}
	//fmt.Printf("%s",body)
	return body
}
