package main

import (
	"0offer/api"
	"0offer/initialize"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	initialize.Mysql()
}

func main() {
	r := gin.Default()
	group := r.Group("/api")
	{
		group.GET("/test", api.Test)
		group.GET("/user_infos", api.SelectAllUserInfo)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("run gin failed with %s\n", err)
	}
}
