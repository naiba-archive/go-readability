/*
 * Copyright (c) 2018, 奶爸<1@5.nu>
 * All rights reserved.
 */

package readability

import (
	"io/ioutil"
	"testing"
	"strconv"
)

func TestParse(t *testing.T) {
	pageUrls := []string{"https://www.lifelonglearning.cc/p82_interview.html", "https://www.qdaily.com/articles/52259.html", "https://www.cnblogs.com/163yun/p/8867738.html"}
	for page := 0; page < 3; page++ {
		htmlStr, _ := ioutil.ReadFile("./test/" + strconv.Itoa(page) + ".html")
		article, err := Parse(string(htmlStr), Option{Debug: false, PageUrl: pageUrls[page]})
		if err != nil {
			t.Log(err)
		} else {
			t.Log("标题", article.Title)
			t.Log("链接", article.URL)
			t.Log("作者", article.Byline)
			t.Log("长度", article.Length)
			t.Log("目录", article.Dir)
			t.Log("摘要", article.Excerpt)
			t.Log(article.Content)
		}
	}
}
