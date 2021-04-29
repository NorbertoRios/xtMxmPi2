package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"httpCpntroller"
	"net/http"
	"strconv"
	"time"
)

func CreateGinServer() {
	router := gin.Default()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(GetConfig().WebServerPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	taskController := &httpCpntroller.HttpTaskController{}
	router.POST("/api/v1/tasks", taskController.CreateTaskPOST)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
