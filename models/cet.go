package models

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/csuhan/csugo/utils"
	"github.com/djimenez/iconv-go"
	"io/ioutil"
	"net/http"
	"strings"
)

const CET_ZKZH_URL = "http://202.197.61.241/cetmodifyb.asp"
const CET_HISTORY_URL = "http://exam.csu.edu.cn/engfen.asp"

type Cet struct {
}
type ZKZH struct {
	ZKZH, Type, Classroom, Seat, Name, ClassID, School, ExamTime, ExamPlace string
}
type HGrade struct {
	Type, ZKZH, ZSH, Grade, Time string
}

//获取准考证号
func (this *Cet) GetZKZ(ID, CETType string) (ZKZH, error) {
	//请求登录
	CETTypes := []string{"%CB%C4%BC%B6", "%C1%F9%BC%B6"} //cet类别
	var bmlb string
	if CETType == "4" {
		bmlb = CETTypes[0]
	} else if CETType == "6" {
		bmlb = CETTypes[1]
	} else {
		return ZKZH{}, utils.ERROR_INPUT
	}
	reqData := "username=" + ID + "&bmlb=" + bmlb
	req, _ := http.NewRequest("POST", CET_ZKZH_URL, strings.NewReader(reqData))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ZKZH{}, utils.ERROR_SERVER
	}
	//将数据转为
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ZKZH{}, utils.ERROR_SERVER
	}
	defer resp.Body.Close()
	utf8data, err := iconv.ConvertString(string(data), "gbk", "utf8")
	if err != nil {
		return ZKZH{}, utils.ERROR_SERVER
	}
	if !strings.Contains(utf8data, "中南大学CET考生信息") { //登录失败
		return ZKZH{}, errors.New("学号或者类别错误")
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(utf8data))
	if err != nil {
		return ZKZH{}, utils.ERROR_SERVER
	}
	zkzh := doc.Find("#zkz").AttrOr("value", "")
	zkz := &ZKZH{
		ZKZH:      zkzh,
		Type:      string([]byte(zkzh)[9]),
		Classroom: string([]byte(zkzh)[10:13]),
		Seat:      string([]byte(zkzh)[13:15]),
		Name:      strings.Trim(doc.Find("#zkz0").AttrOr("value", ""), " "),
		ClassID:   strings.Trim(doc.Find("#yx0").AttrOr("value", ""), " "),
		School:    strings.Trim(doc.Find("option[selected]").AttrOr("value", ""), " "),
		ExamTime:  doc.Find("#yx1").AttrOr("value", ""),
		ExamPlace: strings.Trim(doc.Find("#yx").AttrOr("value", ""), " "),
	}
	return *zkz, nil
}

//获取历史成绩
func (this *Cet) GetHGrade(ID, Name string) ([]HGrade, error) {
	gbkName, _ := iconv.ConvertString(Name, "utf8", "gbk")
	resp, err := http.Get(CET_HISTORY_URL + "?xm=" + gbkName + "&sfzh=&zkzh=&xh=" + ID)
	if err != nil {
		return []HGrade{}, utils.ERROR_SERVER
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []HGrade{}, utils.ERROR_SERVER
	}
	utf8data, err := iconv.ConvertString(string(data), "gbk", "utf8")
	if err != nil {
		return []HGrade{}, utils.ERROR_SERVER
	}
	if !strings.Contains(utf8data, "考试成绩查询结果") { //登录失败
		return []HGrade{}, errors.New("学号或者姓名错误")
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(utf8data))
	if err != nil {
		return []HGrade{}, utils.ERROR_SERVER
	}
	hgrades := make([]HGrade, 0)
	hgrade := &HGrade{}

	doc.Find("tr[height='20']").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td")
		hgrade.Type = td.Eq(0).Text()
		hgrade.ZKZH = td.Eq(1).Text()
		hgrade.ZSH = td.Eq(2).Text()
		hgrade.Grade = td.Eq(3).Text()
		year := "20" + string([]byte(hgrade.ZKZH)[6:8])
		month := ""
		if string([]byte(hgrade.ZKZH)[8]) == "1" {
			month = "06"
		} else {
			month = "12"
		}
		hgrade.Time = year + "." + month
		hgrades = append(hgrades, *hgrade)
	})
	return hgrades, nil
}
