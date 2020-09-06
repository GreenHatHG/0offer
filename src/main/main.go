package main

import (
	"0offer/initialize"
	"net/http"
)

func init() {
	initialize.Mysql()
}

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("Hello World"))
		})
	_ = http.ListenAndServe(":8000", nil)
}

//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"message": "pong",
//		})
//	})
//	_ = r.Run()
//}

//func main() {
//	r := gin.Default()
//	group := r.Group("/api")
//	{
//		group.GET("/test", api.Test)
//		group.GET("/user_infos", api.SelectAllUserInfo)
//		group.POST("/user_infos/actions/batch_insert", api.BatchInsertUserInfo)
//		group.POST("/order/actions/miaosha", api.Miaosha)
//	}
//
//	if err := r.Run(); err != nil {
//		log.Fatalf("run gin failed with %s\n", err)
//	}
//}
