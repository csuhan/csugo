package models

import (
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/utils"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const JWC_BASE_URL = "http://csujwc.its.csu.edu.cn/jsxsd/"
const JWC_LOGIN_URL = JWC_BASE_URL + "xk/LoginToXk"
const JWC_GRADE_URL = JWC_BASE_URL + "kscj/yscjcx_list"
const JWC_RANK_URL = JWC_BASE_URL + "kscj/zybm_cx"
const JWC_CLASS_URL = JWC_BASE_URL + "xskb/xskb_list.do"

type JwcUser struct {
	Id, Pwd, Name, College, Margin, Class string
}

type JwcGrade struct {
	ClassNo int
	FirstTerm, GottenTerm, ClassName,
	MiddleGrade, FinalGrade, Grade,
	ClassScore, ClassType, ClassProp string
}

type Rank struct {
	TotalScore, ClassRank, AverScore string
}

type JwcRank struct {
	User  JwcUser
	Ranks []map[string]Rank
}

type Class struct {
	ClassName, Teacher, Weeks, Place string
}

type Jwc struct{}

//成绩查询
func (this *Jwc) Grade(user *JwcUser) ([]JwcGrade, error) {
	response, err := this.LogedRequest(user, "GET", JWC_GRADE_URL, nil)
	if err != nil {
		return []JwcGrade{}, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if !strings.Contains(string(data), "学生个人考试成绩") {
		return []JwcGrade{}, utils.ERROR_UNKOWN
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return []JwcGrade{}, utils.ERROR_SERVER
	}
	Grades := []JwcGrade{}
	doc.Find("table#dataList tr").Each(func(i int, selection *goquery.Selection) {
		if i != 0 {
			s := selection.Find("td")
			jwcgrade := JwcGrade{
				ClassNo:     i,
				FirstTerm:   s.Eq(1).Text(),
				GottenTerm:  s.Eq(2).Text(),
				ClassName:   s.Eq(3).Text(),
				MiddleGrade: s.Eq(4).Text(),
				FinalGrade:  s.Eq(5).Text(),
				Grade:       s.Eq(6).Text(),
				ClassScore:  s.Eq(7).Text(),
				ClassType:   s.Eq(8).Text(),
				ClassProp:   s.Eq(9).Text(),
			}
			Grades = append(Grades, jwcgrade)
		}
	})
	return Grades, nil
}

//专业排名查询
func (this *Jwc) Rank(user *JwcUser) ([]map[string]Rank, error) {
	response, err := this.LogedRequest(user, "POST", JWC_RANK_URL, strings.NewReader("xqfw="+url.QueryEscape("入学以来")))
	if err != nil {
		return []map[string]Rank{}, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return []map[string]Rank{}, utils.ERROR_SERVER
	}
	terms := make([]string, 0)
	doc.Find("#xqfw option").Each(func(i int, s *goquery.Selection) {
		terms = append(terms, s.Text())
	})

	ranks := make([]map[string]Rank, 0)
	for _, term := range terms {
		resp, _ := this.LogedRequest(user, "POST", JWC_RANK_URL, strings.NewReader("xqfw="+url.QueryEscape(term)))
		data, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
		td := doc.Find("#dataList tr").Eq(1).Find("td")
		rank := Rank{
			TotalScore: td.Eq(1).Text(),
			ClassRank:  td.Eq(2).Text(),
			AverScore:  td.Eq(3).Text(),
		}
		ranks = append(ranks, map[string]Rank{term: rank})
	}
	return ranks, nil
}

//课表查询
func (this *Jwc) Class(user *JwcUser, Week, Term string) ([][]Class, error) {
	if Week == "0" {
		Week = ""
	}
	body := strings.NewReader("zc=" + url.QueryEscape(Week) + "&xnxq01id=" + url.QueryEscape(Term) + "&sfFD=1")
	response, err := this.LogedRequest(user, "POST", JWC_CLASS_URL, body)
	if err != nil {
		return [][]Class{}, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	classes := make([][]Class, 0)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	doc.Find("table#kbtable").Eq(0).Find("td div.kbcontent").Each(func(i int, s *goquery.Selection) {
		font := s.Find("font")
		if font.Size() == 3 { //一节课
			class := Class{
				ClassName: s.Nodes[0].FirstChild.Data,
				Teacher:   font.Eq(0).Text(),
				Weeks:     font.Eq(1).Text(),
				Place:     font.Eq(2).Text(),
			}
			classes = append(classes, []Class{class})
		} else if font.Size() == 6 { //两节课
			class:=[]Class{
				Class{
					ClassName: s.Nodes[0].FirstChild.Data,
					Teacher:   font.Eq(0).Text(),
					Weeks:     font.Eq(1).Text(),
					Place:     font.Eq(2).Text(),
				},
				Class{
					ClassName: font.Eq(3).Nodes[0].PrevSibling.PrevSibling.Data,
					Teacher:   font.Eq(3).Text(),
					Weeks:     font.Eq(4).Text(),
					Place:     font.Eq(5).Text(),
				},
			}
			classes=append(classes,class)
		} else {
			classes = append(classes, make([]Class, 1))
		}
	})

	return classes, nil
}

//登录后请求
func (this *Jwc) LogedRequest(user *JwcUser, Method, Url string, Params io.Reader) (*http.Response, error) {
	//登录系统
	cookies, err := this.Login(user)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}
	//查询分数
	Req, err := http.NewRequest(Method, Url, Params)
	Req.Header.Add("content-type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	for _, cookie := range cookies {
		Req.AddCookie(cookie)
	}
	return http.DefaultClient.Do(Req)
}

//教务系统登录
func (this *Jwc) Login(user *JwcUser) ([]*http.Cookie, error) {
	//获取cookie
	response, err := http.Get(JWC_BASE_URL)
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	cookies := response.Cookies()
	//账号密码拼接字符
	encoded := base64.StdEncoding.EncodeToString([]byte(user.Id)) + "%%%" + base64.StdEncoding.EncodeToString([]byte(user.Pwd))
	Req, err := http.NewRequest("POST", JWC_LOGIN_URL, strings.NewReader("encoded="+url.QueryEscape(encoded)))
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	//添加cookie
	for _, cookie := range cookies {
		Req.AddCookie(cookie)
	}
	Req.Header.Add("content-type", "application/x-www-form-urlencoded")
	response, err = http.DefaultClient.Do(Req)
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	//登陆成功
	if strings.Contains(string(body), "我的桌面") {
		return cookies, nil
	}
	//账号或密码错误
	return nil, utils.ERROR_ID_PWD
}
