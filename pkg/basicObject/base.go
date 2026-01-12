package basicObject

type ResponseError struct {
	Message    string   `json:"message"`
	StatusCode int      `json:"code"`
	Reasons    []string `json:"details"`
}

func NewResponseError(msg string, hc int) (int, ResponseError) {
	return hc, ResponseError{
		Message:    msg,
		StatusCode: hc,
		Reasons:    []string{},
	}
}

func NewResponseWithReasonError(msg string, rsn []string, hc int) (int, ResponseError) {
	return hc, ResponseError{
		Message:    msg,
		StatusCode: hc,
		Reasons:    rsn,
	}
}

type Meta struct {
	Count         string `json:"count"`
	CurrentPage   string `json:"currentPage"`
	LastPage      string `json:"lastPage"`
	RecordPerPage string `json:"recordPerPage"`
}
