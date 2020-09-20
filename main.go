package main

import (
	"log"
	"net/http"
)

func main() {
	//router := gin.New()
	//router.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello")
	//})
	//if err := router.Run(":8080"); err != nil {
	//	log.Fatal(err)
	//}
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello1231"))
	})
	log.Println("listen on :8080")
	//http.ListenAndServe(":8080", nil)
}
