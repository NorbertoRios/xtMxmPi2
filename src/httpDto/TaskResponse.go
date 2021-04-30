package httpDto

type TaskResponse struct {
	Id       int64 `json:"id"`
	QueuedAt int64 `json:"queued_at"`
}
