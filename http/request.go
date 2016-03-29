package http


type Request struct {
	url string
	method string
	params map[string]string
	headers map[string]string
}

func NewRequest(url, method string, params map[string]string, args ...interface{}) *Request {
	return &Request{url: url, method: method, params: params}
}