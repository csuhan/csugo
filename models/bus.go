package models

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/httplib"
	"github.com/csuhan/csugo/utils"
	"regexp"
	"strings"
)

const BUS_SEARCH_URL = "http://app.its.csu.edu.cn/csu-app/cgi-bin/depa/depa?method=search"

type Bus struct {
	StartTime, Start, End, RunTime, Num, Seat string
	Stations                                  []string
}

func (this *Bus) Search(Start, End, Time string) ([]Bus, error) {
	//校车
	buses := make([]Bus, 0)
	//获取页面
	req := httplib.Post(BUS_SEARCH_URL)
	req.Param("startValue", Start)
	req.Param("endValue", End)
	req.Param("timeValue", Time)
	req.Param("selTimeValue", "0")
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	response, err := req.String()
	if err != nil {
		return []Bus{}, utils.ERROR_SERVER
	}
	//解析页面
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		return []Bus{}, utils.ERROR_SERVER
	}
	doc.Find(".busClassDiv").Each(func(i int, s *goquery.Selection) {
		//每辆车
		busClassDiv, _ := s.Html()

		re := regexp.MustCompile("起站发车时间：(.*)  ")
		if temp := re.FindStringSubmatch(busClassDiv); len(temp) == 2 {
			this.StartTime = temp[1]
		}

		re = regexp.MustCompile("台数：(.*)台")
		if temp := re.FindStringSubmatch(busClassDiv); len(temp) == 2 {
			this.Num = temp[1]
		}

		re = regexp.MustCompile("座位数：(.*)座")
		if temp := re.FindStringSubmatch(busClassDiv); len(temp) == 2 {
			this.Seat = temp[1]
		}

		ul := s.Find("ul")
		this.RunTime = ul.Eq(0).Find("font").Text()

		temp := strings.Split(ul.Eq(1).Find("li").Text(), "→")
		if len(temp) == 2 {
			this.Start = strings.Trim(temp[0], " ")
			this.End = strings.Trim(temp[1], " ")
		}
		this.Stations = make([]string, 0)

		ul.Eq(2).Find("li").Not(".f_blue").Each(func(j int, station *goquery.Selection) {
			this.Stations = append(this.Stations, station.Text())
		})

		this.Start = Start
		this.End = End

		buses = append(buses, *this)
	})
	return buses, nil
}
