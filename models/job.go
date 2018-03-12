package models

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/httplib"
	"github.com/csuhan/csugo/utils"
	"regexp"
	"strings"
)

const JOB_BASE_URL = "http://jobsky.csu.edu.cn"
const JOB_ARTICLE_URL = JOB_BASE_URL + "/Home/PartialArticleList"

type Job struct {
	Link  string
	Title string
	Time  string
	Place string
}

func (this *Job) List(Typeid, Pageindex, Pagesize, HasTime string) ([]Job, error) {
	req := httplib.Post(JOB_ARTICLE_URL)
	req.Header("content-Type", "application/x-www-form-urlencoded")
	req.Param("pageindex", Pageindex)
	req.Param("pagesize", Pagesize)
	req.Param("typeid", Typeid)
	req.Param("followingdates", "-1")

	response, err := req.String()
	jobs := make([]Job, 0)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader("<html><body><table>" + response + "</table></body></html>"))
	if err != nil {
		return []Job{}, utils.ERROR_SERVER
	}
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a")
		this.Link = JOB_BASE_URL + a.AttrOr("href", "")
		this.Title = a.Text()
		this.Time = s.Find(".spanDate").Text()
		jobs = append(jobs, *this)
	})
	if HasTime == "1" {
		for k, j := range jobs {
			resp, _ := httplib.Get(j.Link).String()
			re := regexp.MustCompile("<p class=\"text-center place\">招聘地点：(.*)</p>")
			temp := re.FindStringSubmatch(resp)
			if len(temp) == 2 {
				jobs[k].Place = temp[1]
			}
		}
	}
	return jobs, nil
}
