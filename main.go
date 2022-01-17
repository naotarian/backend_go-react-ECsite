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
		&model.Product{},
	)
	db.AutoMigrate(
		&model.User{},
	)

	//テーブル削除処理
	// db.DropTable(&model.User{})
	// db.DropTable(&model.Product{})

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
	api := r.Group("/api")

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "top root.",
		})
	})
	api.GET("/get_product", GetProducts)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test root.",
		})
	})
	r.GET("/add", AddDatas)
	api.POST("get_item", GetItem)
	// api.GET("/get_product", handler.myFunction)
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
	CONNECT := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, PROTOCOL, DBNAME)
	return gorm.Open(DBMS, CONNECT)
}

func AddDatas(c *gin.Context) {
	// db接続
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// user := &model.User{}
	// user.ID = 0
	// user.Name = "test_user"
	// db.Create(&user)
	error := db.Create(&model.Product{
		ID:                  0,
		UserId:              1,
		ProductName:         "【追加予約】BEAMS HEART / キルティング フードコート アウター",
		ProductIntroduction: "これは完全なるテスト商品です。これは完全なるテスト商品です。これは完全なるテスト商品です。",
		Price:               3980,
	}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加成功")
	}
}

func GetProducts(c *gin.Context) {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	products := []model.Product{}
	db.Find(&products)
	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// }
	c.JSON(200, gin.H{
		"message": products,
	})
}
func GetItem(c *gin.Context) {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	products := []model.Product{}
	db.First(&products)
	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// }
	c.JSON(200, gin.H{
		"message": products,
	})
}

func getDate() string {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}
