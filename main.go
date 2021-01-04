package main

import (
	"github.com/joho/godotenv"
	"go-docker-test/cache"
	"go-docker-test/model"
	"go-docker-test/server"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("读取环境变量失败")
	}
	model.Database(os.Getenv("MYSQL_DSN"))
	cache.Redis()

	r := server.NewRouter()
	r.Run(":8011")
}
