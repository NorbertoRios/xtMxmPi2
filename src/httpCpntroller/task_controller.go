package httpCpntroller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"streamax-go/httpDto"
	"streamax-go/httpService"
	"time"
)

type HttpTaskController struct{}

func (h HttpTaskController) CreateTaskPOST(c *gin.Context) {
	var task httpDto.Task
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	err := json.NewDecoder(c.Request.Body).Decode(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	if !task.ValidateTask() {
		c.JSON(http.StatusBadRequest, "")
		return
	}
	streamBits := task.GetTypeFromTask("stream")
	substreamBits := task.GetTypeFromTask("substream")
	screenshotBits := task.GetTypeFromTask("screenshot")
	taskResponse, tErr := httpService.CreateTask(
		&task,
		task.Dsno,
		time.Unix(task.StartTime, 0).UTC(),
		time.Unix(task.EndTime, 0).UTC(),
		streamBits|substreamBits|screenshotBits,
		streamBits,
		substreamBits,
		screenshotBits,
	)
	if tErr != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	resp := taskResponse
	c.JSON(http.StatusOK, resp)
}
