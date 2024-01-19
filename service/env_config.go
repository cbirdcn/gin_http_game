package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EnvConfig struct {
	gorm.Model
	Code  string
	Price uint
}

func GetEnv(c *gin.Context) int{
	return 0
	//if dbName == "" {
	//	dbName = "global"
	//}
	//dbStr := "mysql_" + dbName
	//
	//dbParams := conf.Connections[dbStr].(conf.MysqlConnectionParams)
	//// loc = 服务器本地时间，如果确定Asia%2FShanghai这种时区也可以写，不利于全球化
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local", dbParams.Username, dbParams.Password, dbParams.Host, dbParams.Port, dbParams.Database)
	//DbDefault, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if helpers.CheckNil(DbDefault, true) || helpers.CheckError(err, true) {
	//	// DB()实际上是对*sql.DB(即go-sql-driver)的封装，连接池需要更改*sql.DB配置
	//	sqlDb, err := DbDefault.DB()
	//	if !helpers.CheckError(err, true) {
	//		sqlDb.SetMaxOpenConns(defaultMaxOpenConns)
	//		sqlDb.SetMaxIdleConns(defaultMaxIdleConns)
	//		sqlDb.SetConnMaxLifetime(defaultMaxLifetime)
	//	}
	//}
	//
	//db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// Migrate the schema
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}