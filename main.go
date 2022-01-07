package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// "net/http"
	"ec_site/model"
	"github.com/joho/godotenv"
	"os"
	"time"
)

func main() {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB接続成功")
	}
	//マイグレーション処理
	db.AutoMigrate(
		&model.User{},
	)

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	// config.AllowMethods = []string{}
	config.AllowHeaders = []string{
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	}
	config.AllowCredentials = false
	config.MaxAge = 24 * time.Hour
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "top root.",
		})
	})
	r.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "api root.",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test root.",
		})
	})
	// r.GET("/add", AddDatas)
	r.Run(":8080")
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	userEnv := godotenv.Load("env/dev.env")
	if userEnv != nil {
		panic("Error loading .env file")
	}
	DBMS := os.Getenv("DBMS")
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	//dockerを使う場合は、DBのコンテナ名を指定する
	PROTOCOL := os.Getenv("DB_PROTOCOL")
	DBNAME := os.Getenv("DBNAME")

	// CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	CONNECT := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, PROTOCOL, DBNAME)
	return gorm.Open(DBMS, CONNECT)
}

// func AddDatas(c *gin.Context) {
// 	// db接続
// 	db, err := sqlConnect()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	error := db.Create(&Users{
// 		Name:     "テスト太郎",
// 		Age:      18,
// 		Address:  "東京都千代田区",
// 		UpdateAt: getDate(),
// 	}).Error
// 	if error != nil {
// 		fmt.Println(error)
// 	} else {
// 		fmt.Println("データ追加成功")
// 	}
// }

// func getDate() string {
// 	const layout = "2006-01-02 15:04:05"
// 	now := time.Now()
// 	return now.Format(layout)
// }

// type Users struct {
// 	ID       int
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	Address  string `json:"address"`
// 	UpdateAt string `json:"updateAt" sql:"not null;type:date"`
// }
