package LearnGo_crawl

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"bufio"
	"golang.org/x/text/encoding"
	"log"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

func main(){

	resp,err:= http.Get("https://book.douban.com/")

	if err!=nil{
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		fmt.Printf("Error status code:%d",resp.StatusCode)
	}

	bodyReader:= bufio.NewReader(resp.Body)
	e:= determinEncoding(bodyReader)

	utf8Reader:= transform.NewReader(bodyReader,e.NewDecoder())



	result,err:= ioutil.ReadAll(utf8Reader)
	if err!=nil{
		panic(err)
	}

	fmt.Printf("%s",result)

}
 

func determinEncoding(r * bufio.Reader) encoding.Encoding{

		bytes,err:= r.Peek(1024)

		if err!=nil{
			log.Printf("fetch error:%v",err)
			return unicode.UTF8
		}

		e,_,_:=charset.DetermineEncoding(bytes,"")
		return e
}