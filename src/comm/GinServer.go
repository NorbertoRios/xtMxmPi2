package comm

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"streamax-go/config"
	"streamax-go/httpCpntroller"
	"time"
)

type GinServer struct {
}

func (gs GinServer) RunGinServer() {
	router := gin.Default()
	s := &http.Server{
		Addr:           ":" + strconv.Itoa(config.GetConfig().WebServerPort),
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
