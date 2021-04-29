package httpCpntroller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"httpDto"
	"net/http"
	"service"
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
	}
	str := getTypeFromTask("stream", task)
	sstr := getTypeFromTask("substream", task)
	scr := getTypeFromTask("screenshot", task)
	id, qat, tErr := service.CreateTask(
		task.Dsno,
		time.Unix(task.StartTime, 0).UTC(),
		time.Unix(task.EndTime, 0).UTC(),
		str|sstr|scr,
		str,
		sstr,
		scr,
	)
	if tErr != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	resp := &httpDto.TaskResponse{
		Id:       id,
		QueuedAt: qat,
	}
	c.JSON(http.StatusOK, resp)
}

//1,2,3 adas,dsm,ip   bits
//0 empty
func getTypeFromTask(s string, t httpDto.Task) int {
	var res = 0
	if t.Channels.Ip != nil {
		res = funcName(s, t.Channels.Ip.FileTypes, res, 3)
	}
	if t.Channels.Dsm != nil {
		res = funcName(s, t.Channels.Dsm.FileTypes, res, 2)
	}
	if t.Channels.Adas != nil {
		res = funcName(s, t.Channels.Adas.FileTypes, res, 1)
	}
	return res
}

func funcName(s string, ar []string, res int, ch int) int {
	typesAdas := ar
	for i := 0; i < len(typesAdas); i++ {
		if typesAdas[i] == s {
			res = service.PutFlagToInt(ch, res)
		}
	}
	return res
}
