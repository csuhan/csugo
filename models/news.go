package models

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/httplib"
	"github.com/csuhan/csugo/utils"
	"regexp"
	"strings"
)

const NEWSARTICLE_URL = "http://tz.its.csu.edu.cn/Home/Release_TZTG_zd/"
const NEWSLIST_URL = "http://tz.its.csu.edu.cn/Home/Release_TZTG/"

type NewsItem struct {
	ID, Title, Dept, ViewCount, Time string
	Link, Content                    string
}

type NewsList struct {
	NowPage, TotalPage, TotalNews string
	News                          []NewsItem
}

func GetNewsList(PageID string) (NewsList, error) {
	req := httplib.Post(NEWSLIST_URL + PageID)
	req.Header("x-forwarded-for", "202.197.71.84") //模仿校内登录
	resp, err := req.String()
	if err != nil {
		return NewsList{}, utils.ERROR_SERVER
	}
	news := NewsList{}
	//查找总页数,总信息数
	re := regexp.MustCompile("共有数据：(.*)条&nbsp;共(.*)页&nbsp;当前")
	res := re.FindStringSubmatch(resp)
	news.NowPage = PageID
	if len(res) == 3 {
		news.TotalNews = res[1]
		news.TotalPage = res[2]
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp))
	if err != nil {
		return NewsList{}, utils.ERROR_SERVER
	}
	newsItems := []NewsItem{}
	//查找每个tr
	doc.Find(".trs tr").Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		temp := tds.Find("a").AttrOr("onclick", "")
		link := regexp.MustCompile(`/Home/Release_TZTG_zd/(.*)', '', 'left=0`).FindStringSubmatch(temp)[1]
		newsItems = append(newsItems, NewsItem{
			ID:        strings.Trim(tds.Eq(3).Text(), "\n "),
			Title:     strings.Trim(tds.Eq(4).Text(), "\n "),
			Dept:      strings.Trim(tds.Eq(5).Text(), "\n "),
			ViewCount: strings.Trim(tds.Eq(6).Text(), "\n "),
			Time:      strings.Trim(tds.Eq(7).Text(), "\n "),
			Link:      NEWSARTICLE_URL + link,
		})
	})
	news.News = newsItems
	return NewsList{}, nil
}
