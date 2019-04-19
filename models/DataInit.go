package models

import (
	"github.com/astaxie/beego"
	//"github.com/Unknwon/com"
	//"os"
	//"path"
	"github.com/astaxie/beego/orm"
	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

func RegistDB() {
	//if !com.IsExist(_DB_NAME){
	//	os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
	//	os.Create(_DB_NAME)
	//}
	//orm.RegisterModel(new(Attachment),new(Topic))
	//orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)
	//orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)

	orm.RegisterModel(new(User), new(Blog), new(NLabel), new(Note), new(NoteColl), new(Category), new(Like), new(Comment))
	orm.RegisterDataBase("default", "mysql", "root:"+beego.AppConfig.String("dburl")+"/beeblog?charset=utf8&loc=Local", 30)
}
