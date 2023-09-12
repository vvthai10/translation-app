package main

import (
	"log"
	"translation-app/controller/httpapi"
	"translation-app/infrastructure/googlesv"
	"translation-app/infrastructure/postgres"
	"translation-app/service"

	"github.com/gin-gonic/gin"
)

func main() {

	// Setup Dependencies
	// Tạo/Connect với DB
	// mongoRepo := mongodb.NewMongoDB(nil)
	// mysqlRepo := mysql.NewMySQLRepo(nil)
	postgresDB := postgres.NewPostgresDB()
	// Tạo/Connect với 1 dịch vụ
	googleTranslate := googlesv.New()
	// Khởi tạo use case
	service := service.NewService(postgresDB, googleTranslate)

	// Tạo các controller
	controller := httpapi.NewAPIController(service)

	engine := gin.Default()

	v1 := engine.Group("/v1")
	controller.SetUpRoute(v1)

	if err := engine.Run(); err != nil {
		log.Fatalln(err)
	}

}

// db, err := connectDBWithRetry(5)

// 	if err != nil {
// 		log.Fatalln(err)
// }

// func connectDBWithRetry(times int) (*gorm.DB, error) {
// 	var e error

// 	for i := 1; i <= times; i++ {
// 		dsn := os.Getenv("MYSQL_DSN")
// 		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 		if err == nil {
// 			return db, nil
// 		}

// 		e = err

// 		time.Sleep(time.Second * 2)
// 	}

// 	return nil, e
// }