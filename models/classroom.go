package models

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/csuhan/csugo/utils"
	"strings"
)

const BURKETCLASSTEMP = "classtemp"
const BURKETCLASSROOM = "classroom"

type XQ struct {
	ID   string
	Name string
}

type JXL struct {
	ID   string `json:"jzwid"`
	Name string `json:"jzwmc"`
	XQ   XQ     `json:"XQ"`
}

type ClassRoom struct {
	JSID         string `json:"jsid"`
	ClassRoomID  string `json:"jsmc"`
	JXL          JXL    `json:"JXL"`
	FreeWeekTime []bool
}

var XQS = []XQ{
	{ID: "1", Name: "校本部"}, {ID: "2", Name: "南校区"}, {ID: "3", Name: "铁道校区"},
	{ID: "4", Name: "湘雅新校区"}, {ID: "5", Name: "湘雅老校区"}, {ID: "6", Name: "湘雅医院"},
	{ID: "7", Name: "湘雅二医院"}, {ID: "8", Name: "湘雅三医院"}, {ID: "9", Name: "新校区"},
}

func getDB() (*bolt.DB, error) {
	dbName := beego.AppConfig.String("DB::ClassesDB")
	db, err := bolt.Open(dbName, 777, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetFreeWeekTime(Term, Week, XQ, JXL string) ([]ClassRoom, error) {
	// -------获取教室空闲时间--------------
	db, err := getDB()
	if err != nil {
		return []ClassRoom{}, err
	}
	freeWeekTimes := []struct {
		Key   string
		Value []bool
	}{}
	db.View(func(tx *bolt.Tx) error {
		ccls := tx.Bucket([]byte(BURKETCLASSTEMP)).Cursor()
		prefix := Term + ":" + Week + ":" + XQ + ":" + JXL + ":"
		for k, v := ccls.Seek([]byte(prefix)); bytes.HasPrefix(k, []byte(prefix)); k, v = ccls.Next() {
			freeWeekTime := []bool{}
			json.Unmarshal(v, &freeWeekTime)
			freeWeekTimes = append(freeWeekTimes, struct {
				Key   string
				Value []bool
			}{Key: string(k), Value: freeWeekTime})
		}
		return nil
	})
	// -------获取教室名称--------------
	tcls := []ClassRoom{}
	db.View(func(tx *bolt.Tx) error {
		data := tx.Bucket([]byte(BURKETCLASSROOM)).Get([]byte(XQ + ":" + JXL))
		json.Unmarshal(data, &tcls)
		return nil
	})
	// -------教室与空闲时间关联--------------
	classrooms := []ClassRoom{}
	classroom := ClassRoom{}
	for _, freeWeekTime := range freeWeekTimes {
		keys := strings.Split(freeWeekTime.Key, ":")
		js := keys[4]
		for _, tcl := range tcls {
			if js == tcl.JSID {
				classroom = tcl
				classroom.FreeWeekTime = freeWeekTime.Value
				classrooms = append(classrooms, classroom)
			}
		}
	}
	db.Close()
	return classrooms, nil
}

//根据校区获取教学楼
func GetBuildingsByXQ(XQ string) ([]JXL, error) {
	db, err := getDB()
	if err != nil {
		return []JXL{}, utils.ERROR_SERVER
	}
	jxls := []JXL{}
	db.View(func(tx *bolt.Tx) error {
		cjxls := tx.Bucket([]byte(BURKETCLASSROOM)).Cursor()
		prefix := XQ + ":"
		for k, v := cjxls.Seek([]byte(prefix)); bytes.HasPrefix(k, []byte(prefix)); k, v = cjxls.Next() {
			cls := []ClassRoom{}
			json.Unmarshal(v, &cls)
			if len(cls) != 0 {
				jxls = append(jxls, cls[0].JXL)
			}
		}
		return nil
	})
	db.Close()
	return jxls, nil
}

//获取所有教学楼
func GetBuildings() map[string][]JXL {
	jxls := map[string][]JXL{}
	for _, XQ := range XQS {
		jxls[XQ.ID], _ = GetBuildingsByXQ(XQ.ID)
	}
	return jxls
}
