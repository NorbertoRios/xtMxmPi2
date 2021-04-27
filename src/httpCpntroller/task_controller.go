package httpCpntroller

import (
	"encoding/json"
	"httpDto"
	"net/http"
	"service"
	"time"
)

type TaskHandler struct{}

func (h TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//request.Body.Read()
		var task httpDto.Task
		if r.Body == nil {
			http.Error(w, "Empty body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		if !validateTask(task) {
			w.WriteHeader(400)
		}
		str := getTypeFromTask("stream", task)
		sstr := getTypeFromTask("substream", task)
		scr := getTypeFromTask("screenshot", task)
		err = service.CreateTask(
			task.Dsno,
			time.Unix(0, task.StartTime*1000000).UTC(),
			time.Unix(0, task.EndTime*1000000).UTC(),
			str|sstr|scr,
			str,
			sstr,
			scr,
		)
		if err != nil {
			w.WriteHeader(500)
		}
	}
}

func validateTask(t httpDto.Task) bool {
	if t.StartTime < 1556342281 || t.EndTime < 1556342281 || t.StartTime > 4112486281 || t.EndTime > 4112486281 {
		return false
	}
	if len(t.Dsno) < 3 {
		return false
	}
	if t.Channels == nil || (t.Channels.Adas == nil && t.Channels.Dsm == nil && t.Channels.Ip == nil) {
		return false
	}
	return true
}

//1,2,3 adas,dsm,ip   bits
//0 empty
func getTypeFromTask(s string, t httpDto.Task) int {
	var res int = 0
	if t.Channels.Ip != nil {
		types := t.Channels.Ip.FileTypes
		for i := 0; i < len(types); i++ {
			if types[i] == s {
				res = service.PutFlagToInt(3, res)
			}
		}
	}
	if t.Channels.Dsm != nil {
		typesDsm := t.Channels.Dsm.FileTypes
		for i := 0; i < len(typesDsm); i++ {
			if typesDsm[i] == s {
				res = service.PutFlagToInt(2, res)
			}
		}
	}
	if t.Channels.Adas != nil {
		typesAdas := t.Channels.Adas.FileTypes
		for i := 0; i < len(typesAdas); i++ {
			if typesAdas[i] == s {
				res = service.PutFlagToInt(1, res)
			}
		}
	}
	return res
}
