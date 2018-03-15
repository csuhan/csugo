package models

import (
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
)

func InitDB() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists([]byte("wxuser"))
		return err
	})
	return nil
}

func GetDB() (*bolt.DB, error) {
	DB := &bolt.DB{}
	//初始化数据库连接
	WxAppDBConfig := beego.AppConfig.String("DB::WxAppDB")
	DB, err := bolt.Open(WxAppDBConfig, 0600, nil)
	if err != nil {
		beego.Debug("database error")
		return nil, err
	}
	return DB, nil
}
