package model

type Error struct {
	ErrorMessage string `json:"errorMessage"`
	TimeStamp    string `json:"timeStamp"`
}

func (e *Error) SetErrorMessage(err string) {
	e.ErrorMessage = err
}

func (e *Error) SetTimeStamp(timeStamp string) {
	e.TimeStamp = timeStamp
}

func (e *Error) GetErrorMessage() string {
	return e.ErrorMessage
}

func (e *Error) GetTimeStamp() string {
	return e.TimeStamp
}
