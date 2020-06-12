package nytimes_com

import (
	"github.com/dreamerjackson/fast-crawl/engine"
	"github.com/dreamerjackson/fast-crawl/model/news"
	"regexp"
	"time"
)

var headerRe = regexp.MustCompile(`<h3 class="regularSummaryHeadline"><a.*?>(.*?)</a>[\s\S]*?<p class="summary">(.*?)</p>`)
func ParseHome(contents []byte) engine.ParseResult{
	matches:= headerRe.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	result.Flag = engine.FlagEnd
	result.Items = make([]interface{},0,5)
	now := time.Now()
	for _,m:= range matches{
		if len(m) < 3 {
			continue
		}
		item := news.News{
			Maintitle: string(m[1]),
			Subtitle:  string(m[2]),
			Type:"cn",
			Time:now,
			SourceName:"纽约时报",
			Push: engine.TELEGRAM,
			// Link:      "",
		}
		//不用用户名了
		result.Items = append(result.Items,item)
	}
	return result

}
