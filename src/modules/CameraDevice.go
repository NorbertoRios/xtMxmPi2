package modules

type CameraDevice struct {
	Id string
}

func (c CameraDevice) GetId() string {
	return c.Id
}
