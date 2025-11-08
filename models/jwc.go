package models

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego"
	"github.com/csuhan/csugo/utils"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
)

const JWC_BASE_URL = "http://csujwc.its.csu.edu.cn/jsxsd/"
const JWC_LOGIN_URL = JWC_BASE_URL + "xk/LoginToXk"
const JWC_GRADE_URL = JWC_BASE_URL + "kscj/yscjcx_list"
const JWC_RANK_URL = JWC_BASE_URL + "kscj/zybm_cx"
const JWC_CLASS_URL = JWC_BASE_URL + "xskb/xskb_list.do"
const CAS_LOGIN_URL = "https://ca.csu.edu.cn/authserver/login?service=http%3A%2F%2Fcsujwc.its.csu.edu.cn%2Fsso.jsp"

const aesCharSet = "ABCDEFGHJKMNPQRSTWXYZabcdefhijkmnprstwxyz2345678"

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
	Term, TotalScore, ClassRank, AverScore string
}

type JwcRank struct {
	User  JwcUser
	Ranks []Rank
}

type Class struct {
	ClassName, Teacher, Weeks, Place string
}

type Jwc struct{}

// 成绩查询
func (this *Jwc) Grade(user *JwcUser) ([]JwcGrade, error) {
	response, err := this.LogedRequest(user, "GET", JWC_GRADE_URL, nil)
	if err != nil {
		return []JwcGrade{}, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if !strings.Contains(string(data), "学生个人考试成绩") {
		return []JwcGrade{}, utils.ERROR_JWC
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

// 专业排名查询
func (this *Jwc) Rank(user *JwcUser) ([]Rank, error) {
	response, err := this.LogedRequest(user, "POST", JWC_RANK_URL, strings.NewReader("xqfw="+url.QueryEscape("入学以来")))
	if err != nil {
		return []Rank{}, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		return []Rank{}, utils.ERROR_SERVER
	}
	terms := make([]string, 0)
	doc.Find("#xqfw option").Each(func(i int, s *goquery.Selection) {
		terms = append(terms, s.Text())
	})
	err = nil
	ranks := make([]Rank, len(terms))
	ch := make(chan map[int]Rank)
	chanRanks := []map[int]Rank{}
	for key, term := range terms {
		go func(key int, term string, ch chan map[int]Rank) {
			resp, _ := this.LogedRequest(user, "POST", JWC_RANK_URL, strings.NewReader("xqfw="+url.QueryEscape(term)))
			data, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
			td := doc.Find("#dataList tr").Eq(1).Find("td")
			rank := Rank{
				Term:       term,
				TotalScore: td.Eq(1).Text(),
				ClassRank:  td.Eq(2).Text(),
				AverScore:  td.Eq(3).Text(),
			}
			ch <- map[int]Rank{key: rank}
		}(key, term, ch)
	}
	for range terms {
		chanRanks = append(chanRanks, <-ch)
	}
	for i := 0; i < len(terms); i++ {
		for j := 0; j < len(chanRanks); j++ {
			if v, ok := chanRanks[j][i]; ok {
				ranks[i] = v
			}
		}
	}
	return ranks, err
}

// 课表查询
func (this *Jwc) Class(user *JwcUser, Week, Term string) ([][]Class, string, error) {
	if Week == "0" {
		Week = ""
	}
	body := strings.NewReader("zc=" + url.QueryEscape(Week) + "&xnxq01id=" + url.QueryEscape(Term) + "&sfFD=1")
	response, err := this.LogedRequest(user, "POST", JWC_CLASS_URL, body)
	if err != nil {
		return [][]Class{}, "", err
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
			class := []Class{
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
			classes = append(classes, class)
		} else {
			classes = append(classes, make([]Class, 1))
		}
	})

	//每学期开学时间
	temp := doc.Find("table#kbtable").Eq(1).Find("td").Eq(0).Text()
	startWeekDay := regexp.MustCompile("第1周\u00a0(.*)日至").FindStringSubmatch(temp)

	return classes, startWeekDay[1], nil
}

// 登录后请求
func (this *Jwc) LogedRequest(user *JwcUser, Method, Url string, Params io.Reader) (*http.Response, error) {
	client, err := this.Login(user)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}
	Req, err := http.NewRequest(Method, Url, Params)
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	Req.Header.Add("content-type", "application/x-www-form-urlencoded")
	return client.Do(Req)
}

// 教务系统登录
func (this *Jwc) Login(user *JwcUser) (*http.Client, error) {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	resp, err := client.Get(CAS_LOGIN_URL)
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	formValues, err := this.buildCASForm(user, doc)
	if err != nil {
		return nil, err
	}
	loginReq, err := http.NewRequest("POST", CAS_LOGIN_URL, strings.NewReader(formValues.Encode()))
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	loginReq.Header.Set("content-type", "application/x-www-form-urlencoded")
	resp, err = client.Do(loginReq)
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	defer resp.Body.Close()
	_, _ = ioutil.ReadAll(resp.Body)
	if resp.Request == nil || !strings.Contains(resp.Request.URL.Host, "csujwc.its.csu.edu.cn") {
		return nil, utils.ERROR_ID_PWD
	}
	return client, nil
}

func (this *Jwc) buildCASForm(user *JwcUser, doc *goquery.Document) (url.Values, error) {
	lt := strings.TrimSpace(doc.Find("input[name=lt]").AttrOr("value", ""))
	execution := strings.TrimSpace(doc.Find("input[name=execution]").AttrOr("value", ""))
	eventId := strings.TrimSpace(doc.Find("input[name=_eventId]").AttrOr("value", "submit"))

	// 明确选择用户名密码登录方式的 cllt 和 dllt
	// 页面上有多个 cllt 字段对应不同登录方式（fidoLogin, dynamicLogin, userNameLogin, qrLogin）
	// 我们需要使用 userNameLogin 进行用户名密码登录
	cllt := strings.TrimSpace(doc.Find("input[name=cllt][value=userNameLogin]").AttrOr("value", "userNameLogin"))
	dllt := strings.TrimSpace(doc.Find("input[name=dllt]").AttrOr("value", "generalLogin"))

	salt := strings.TrimSpace(doc.Find("input#pwdEncryptSalt").AttrOr("value", ""))
	if salt == "" || execution == "" {
		return nil, utils.ERROR_SERVER
	}
	encryptedPwd, err := encryptPassword(user.Pwd, salt)
	if err != nil {
		return nil, utils.ERROR_SERVER
	}
	form := url.Values{}
	form.Set("username", user.Id)
	form.Set("password", encryptedPwd)
	form.Set("passwordText", "")
	form.Set("lt", lt)
	form.Set("execution", execution)
	form.Set("_eventId", eventId)
	form.Set("cllt", cllt)
	form.Set("dllt", dllt)
	return form, nil
}

func encryptPassword(password, salt string) (string, error) {
	if salt == "" {
		return "", errors.New("empty salt")
	}
	prefix, err := randomString(64)
	if err != nil {
		return "", err
	}
	iv, err := randomString(16)
	if err != nil {
		return "", err
	}
	plain := []byte(prefix + password)
	plain = pkcs7Pad(plain, aes.BlockSize)
	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return "", err
	}
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	ciphertext := make([]byte, len(plain))
	mode.CryptBlocks(ciphertext, plain)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func randomString(length int) (string, error) {
	if length <= 0 {
		return "", nil
	}
	var builder strings.Builder
	for i := 0; i < length; i++ {
		idx, err := randInt(len(aesCharSet))
		if err != nil {
			return "", err
		}
		builder.WriteByte(aesCharSet[idx])
	}
	return builder.String(), nil
}

func randInt(max int) (int, error) {
	if max <= 0 {
		return 0, errors.New("invalid max")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}
