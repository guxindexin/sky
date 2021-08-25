package migrate

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gorm.io/gorm"
)

func InitDb(db *gorm.DB, filePath string) (err error) {
	err = ExecSql(db, filePath)
	return err
}

func ExecSql(db *gorm.DB, filePath string) error {
	sql, err := Ioutil(filePath)
	if err != nil {
		fmt.Println("数据库基础数据初始化脚本读取失败！原因:", err.Error())
		return err
	}
	if err = db.Exec(sql).Error; err != nil {
		log.Printf("error sql: %s", sql)
		if !strings.Contains(err.Error(), "Query was empty") {
			return err
		}
	}
	return nil
}

func Ioutil(filePath string) (string, error) {
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		return result, nil
	} else {
		return "", err
	}
}
