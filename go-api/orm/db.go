package orm

import(
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB(){
	dsn :=os.Getenv("MYSQL_DNS")
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		panic("fail to connect database")
	}
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Car{})
	Db.AutoMigrate(&Booking{})
}