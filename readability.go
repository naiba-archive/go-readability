/*
 * Copyright (c) 2018, 奶爸<1@5.nu>
 * All rights reserved.
 */

package readability

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"log"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	whitespacePattern  = regexp.MustCompile(`^\s*$`)
	defaultTagsToScore = map[string]int{
		"section": 0,
		"h2":      0,
		"h3":      0,
		"h4":      0,
		"h5":      0,
		"h6":      0,
		"p":       0,
		"td":      0,
		"pre":     0,
	}
	bylinePattern               = regexp.MustCompile(`byline|author|dateline|writtenby|p-author`)
	okMaybeItsACandidatePattern = regexp.MustCompile(`and|article|body|column|main|shadow|app|container`)
	unlikelyCandidatesPattern   = regexp.MustCompile(`banner|breadcrumbs|combx|comment|community|cover-wrap|disqus|extra|foot|header|legends|menu|related|remark|replies|rss|shoutbox|sidebar|skyscraper|social|sponsor|supplemental|ad-break|agegate|pagination|pager|popup|yom-remote`)
	option                      = new(Option)
	flags                       = flagCleanConditionally | flagStripUnlikely | flagWeightClasses
)

const (
	flagStripUnlikely      = 0x1
	flagWeightClasses      = 0x2
	flagCleanConditionally = 0x4
)

//Option 解析配置
type Option struct {
	MaxNodeNum    int
	Debug         bool
	ArticleByline string
}

//Metadata 文章摘要信息
type Metadata struct {
	Title   string
	Excerpt string
	Byline  string
}

//Readability 解析结果
type Readability struct {
	Title  string
	Byline string
}

//Parse 进行解析
func Parse(s string, opt Option) (*Readability, error) {
	d, err := goquery.NewDocumentFromReader(strings.NewReader(s))
	if err != nil {
		return nil, err
	}
	option = &opt
	// 超出最大解析限制
	if opt.MaxNodeNum > 0 && len(d.Nodes) > opt.MaxNodeNum {
		return nil, fmt.Errorf("Node 数量超出最大限制：%d 。 ", opt.MaxNodeNum)
	}
	article := new(Readability)
	// 预处理HTML文档以提高可读性。 这包括剥离JavaScript，CSS和处理没用的标记等内容。
	prepDocument(d)

	// 获取文章的摘要和作者信息
	md := getArticleMetadata(d)
	article.Title = md.Title

	//todo 提取文章正文
	getArticle(d)

	return article, nil
}

// 提取文章正文
func getArticle(d *goquery.Document) *goquery.Selection {
	page := d.Find("body").First().Children()
	if page.Length() == 0 {
		l("getArticle", "没有 body，哪里来的正文？")
		return nil
	}
	stripUnlikelyCandidates := flagIsActive(flagStripUnlikely)
	selectionsToScore := make([]*goquery.Selection, 0)
	sel := page.First()
	for sel != nil {
		node := sel.Get(0)
		// 首先，节点预处理。 清理看起来很糟糕的垃圾节点（比如类名为“comment”的垃圾节点），
		// 并将div转换为P标签，清理空节点。
		matchString, _ := sel.Attr("id")
		class, _ := sel.Attr("class")
		matchString += " " + class
		// 作者信息行
		if checkByline(sel, matchString) {
			l("getArticle", "作者信息", "清除")
			sel = removeAndGetNext(sel)
			continue
		}
		// 清理垃圾标签
		if stripUnlikelyCandidates {
			if unlikelyCandidatesPattern.MatchString(matchString) &&
				!okMaybeItsACandidatePattern.MatchString(matchString) &&
				node.Data != "body" &&
				node.Data != "a" {
				l("getArticle", "垃圾标签", "清除")
				sel = removeAndGetNext(sel)
				continue
			}
		}
		// 清理不含任何内容的 DIV, SECTION, 和 HEADER
		tags := map[string]int{"div": 0, "section": 0, "header": 0, "h1": 0, "h2": 0, "h3": 0, "h4": 0, "h5": 0, "h6": 0}
		if _, has := tags[node.Data]; has && len(ts(sel.Text())) == 0 {
			l("getArticle", "清理空块级元素", "清除")
			sel = removeAndGetNext(sel)
			continue
		}
		// 内容标签，加分项
		if _, has := defaultTagsToScore[node.Data]; has {
			selectionsToScore = append(selectionsToScore, sel)
		}
		// 将所有没有 children 的 div 转换为 p
		if node.Data == "div" {
			// 将只包含一个 p 标签的 div 标签去掉，将 p 提出来

		}

		sel = getNextSelection(sel, false)
	}
	return nil
}

// 删除并获取下一个
func removeAndGetNext(s *goquery.Selection) *goquery.Selection {
	l("removeAndGetNext", s.Get(0))
	t := getNextSelection(s, true)
	s.Remove()
	return t
}

/*
 * 从 node 开始遍历DOM，
 * 如果 ignoreSelfAndKids 为 true 则不遍历子 element
 * 改为遍历 兄弟 和 父级兄弟 element
 */
func getNextSelection(s *goquery.Selection, ignoreSelfAndChildren bool) *goquery.Selection {
	if s.Length() == 0 {
		l("getNextSelection", "空空如也😂")
		return nil
	}
	// 如果 ignoreSelfAndKids 不为 true 且 node 有子 element 返回第一个子 element
	if !ignoreSelfAndChildren && s.Children().Length() > 0 {
		t := s.Children().First()
		if t.Length() > 0 {
			l("getNextSelection", "儿子", t.Get(0))
			return t
		}
	}
	// 然后是兄弟 element
	if s.Next().Length() > 0 {
		l("getNextSelection", "兄弟", s.Next().Get(0))
		return s.Next()
	}
	// 最后，父节点的兄弟 element
	//（因为这是深度优先遍历，我们已经遍历了父节点本身）。
	for {
		s = s.Parent()
		t := s.Next()
		if t.Length() == 0 {
			if s.Parent().Length() > 0 {
				continue
			}
			break
		} else {
			l("getNextSelection", "父兄", t.Get(0))
			return t
		}
	}
	l("getNextSelection", "遍历完毕😂")
	return nil
}

// 是否是作者信息
func checkByline(s *goquery.Selection, matchString string) bool {
	if len(option.ArticleByline) > 0 {
		return false
	}
	innerText := s.Text()
	if s.AttrOr("rel", "") == "author" || bylinePattern.MatchString(matchString) && isValidByline(innerText) {
		option.ArticleByline = ts(innerText)
		return true
	}
	return false
}

// 合理的作者信息行
func isValidByline(line string) bool {
	length := utf8.RuneCountInString(ts(line))
	return length > 10 && length < 100
}

// 是否启用
func flagIsActive(flag int) bool {
	return flags&flag > 0
}

// 从 metadata 获取文章的摘要和作者信息
func getArticleMetadata(d *goquery.Document) Metadata {
	var md Metadata
	values := make(map[string]string)

	namePattern := regexp.MustCompile(`/^\s*((twitter)\s*:\s*)?(description|title)\s*$/gi`)
	propertyPattern := regexp.MustCompile(`/^\s*((twitter)\s*:\s*)?(description|title)\s*$/gi`)

	// 提取元数据
	d.Find("meta").Each(func(i int, s *goquery.Selection) {
		elementName, _ := s.Attr("name")
		elementProperty, _ := s.Attr("property")

		if _, has := map[string]string{elementName: "", elementProperty: ""}["author"]; has {
			md.Byline, _ = s.Attr("content")
		}

		var name string
		if namePattern.MatchString(elementName) {
			name = elementName
		} else if propertyPattern.MatchString(elementProperty) {
			name = elementProperty
		}

		if len(name) > 0 {
			elementContent, _ := s.Attr("content")
			if len(elementContent) > 0 {
				name = whitespacePattern.ReplaceAllString(strings.ToLower(name), " ")
				values["name"] = ts(elementContent)
			}
		}

	})

	// 取文章摘要
	if val, has := values["description"]; has {
		md.Excerpt = val
	} else if val, has := values["og:description"]; has {
		md.Excerpt = val
	} else if val, has := values["twitter:description"]; has {
		md.Excerpt = val
	}

	// 取网页标题
	md.Title = getArticleTitle(d)
	if len(md.Title) < 1 {
		if val, has := values["og:title"]; has {
			md.Title = val
		} else if val, has := values["twitter:title"]; has {
			md.Title = val
		}
	}

	return md
}

// 获取文章标题
func getArticleTitle(d *goquery.Document) string {
	titleSplitPattern := regexp.MustCompile(`(.*)[\|\-\\\/>»](.*)`)
	var title, originTitle string

	// 从 title 标签获取标题
	elementTitle := d.Find("title").First()
	originTitle = ts(elementTitle.Text())
	title = originTitle

	hasSplit := titleSplitPattern.MatchString(title)
	if hasSplit {
		// 是否有分隔符，判断主题在前还是在后
		title = titleSplitPattern.ReplaceAllString(originTitle, "$1")
		if utf8.RuneCountInString(title) < 3 {
			title = titleSplitPattern.ReplaceAllString(originTitle, "$2")
		}
	} else if strings.Index("：", originTitle) != -1 || strings.Index(":", originTitle) != -1 {
		// 判断是否有 "：" 符号
		flag := false
		d.Find("h1,h2").EachWithBreak(func(i int, s *goquery.Selection) bool {
			// 提取的标题是否在正文中存在
			if ts(s.Text()) == title {
				flag = true
			}
			return !flag
		})
		if !flag {
			// 如果不存在取 ":" 前后的文字
			i := strings.LastIndex(originTitle, "：")
			if i == -1 {
				i = strings.LastIndex(originTitle, ":")
			} else {
				title = originTitle[i:]
				if utf8.RuneCountInString(title) < 3 {
					i = strings.Index(originTitle, "：")
					if i == -1 {
						i = strings.Index(originTitle, ":")
					} else {
						title = originTitle[i:]
					}
				} else if utf8.RuneCountInString(originTitle[0:i]) > 5 {
					title = originTitle
				}
			}
		}
	} else if utf8.RuneCountInString(title) > 150 || utf8.RuneCountInString(title) < 15 {
		// 如果标题字数很离谱切只有一个h1标签，取其文字
		h1s := d.Find("h1")
		if h1s.Length() == 1 {
			title = ts(h1s.First().Text())
		}
	}

	titleCount := utf8.RuneCountInString(title)

	if titleCount < 4 && (!hasSplit || utf8.RuneCountInString(titleSplitPattern.ReplaceAllString(originTitle, "$1$2"))-1 != titleCount) {
		// 如果提取的标题很短 取网页标题
		title = originTitle
	}

	return title
}

// 预处理HTML文档以提高可读性。 这包括剥离JavaScript，CSS和处理没用的标记等内容。
func prepDocument(d *goquery.Document) {
	// 移除所有script标签
	removeTags("script,noscript", d)

	// 移除所有style标签
	removeTags("style", d)

	// 将多个连续的<br>替换成<p>
	replaceBrs(d)

	// 将所有的font替换成span
	replaceSelectionTags(d.Find("font"), "span")
}

// 将多个连续的<br>替换成<p>
func replaceBrs(d *goquery.Document) {
	d.Find("br").Each(func(i int, br *goquery.Selection) {
		// 当有 2 个或多个 <br> 时替换成 <p>
		replaced := false

		// 如果找到了一串相连的 <br>，忽略中间的空格，移除所有相连的 <br>
		next := nextElement(br.Get(0).NextSibling)
		for next != nil && next.Data == "br" {
			replaced = true
			t := nextElement(next.NextSibling)
			next.Parent.RemoveChild(next)
			next = t
		}

		// 如果移除了 <br> 链，将其余的 <br> 替换为 <p>，将其他相邻节点添加到 <p> 下。直到遇到第二个 <br>
		if replaced {
			pNode := br.Get(0)
			pNode.Data = "p"
			pNode.Namespace = "p"
			br.Text()

			next = pNode.NextSibling
			for next != nil {
				// 如果我们遇到了其他的 <br><br> 结束添加
				if pNode.Data == "br" {
					innerNext := nextElement(next)
					if innerNext.Data == "br" {
						break
					}
				}
				// 否则将节点添加为 <p> 的子节点
				temp := next.NextSibling
				next.Parent.RemoveChild(next)
				next.Parent = nil
				next.PrevSibling = nil
				next.NextSibling = nil
				pNode.AppendChild(next)
				next = temp
			}
		}
	})
}

// 获取下一个Element
func nextElement(n *html.Node) *html.Node {
	for n != nil &&
		n.Type != html.ElementNode && (whitespacePattern.MatchString(n.Data) ||
		n.Type == html.CommentNode) {
		l("nextElement", n)
		n = n.NextSibling
	}
	return n
}

// 移除所有 tags 标签
// 例如 "script,noscript" 清理所有script
func removeTags(tags string, d *goquery.Document) {
	d.Find(tags).Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})
}

// 将所有的s的标签替换成tag
func replaceSelectionTags(s *goquery.Selection, tag string) {
	s.Each(func(i int, is *goquery.Selection) {
		is.Get(0).Data = tag
		is.Get(0).Namespace = tag
	})
}

// 调试日志
func l(ms ...interface{}) {
	if option.Debug {
		log.Println(ms...)
	}
}

// TrimSpace
func ts(s string) string {
	return strings.TrimSpace(s)
}
