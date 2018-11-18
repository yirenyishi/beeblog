package models

import (
	"github.com/Unknwon/com"
	"os"
	"path"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const(
	_DB_NAME = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

func RegistDB()  {
	if !com.IsExist(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}
	//orm.RegisterModel(new(Attachment),new(Topic))
	orm.RegisterModel(new(Attachment),new(User),new(Blog),new(NLabel))
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)
}