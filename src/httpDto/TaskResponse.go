package httpDto

type TaskResponse struct {
	Id       int64 `json:"id"`
	QueuedAt int   `json:"queued_at"`
}
