package builder


type MethodType string

const (
	Post MethodType = "POST"
	Get MethodType = "GET"
	Put MethodType = "PUT"
	Delete MethodType = "DELETE"
)

type Request struct{
	Host string
	Method MethodType
	Url string
	Headers map[string]string
	Body any
}

type RequestBuilder struct{
	request Request
}

func NewRequestBuilder() *RequestBuilder{
	return &RequestBuilder{
		request: Request{Headers: map[string]string{}},
	}
}

func (r *RequestBuilder) Host(host string) *RequestBuilder{
	r.request.Host = host
	return r
}

func (r *RequestBuilder) Method(method MethodType) *RequestBuilder{
	r.request.Method = method
	return r
}

func (r *RequestBuilder) Url(url string) *RequestBuilder{
	r.request.Url = url
	return r
}

func (r *RequestBuilder) AddHeader(header string,value string) *RequestBuilder{
	if r.request.Headers == nil {
		r.request.Headers = map[string]string{}
	}
	r.request.Headers[header] = value
	return r
}

func (r *RequestBuilder) Body(body any) *RequestBuilder{
	r.request.Body = body
	return r
}

func (r *RequestBuilder) Build() *Request{
	return &r.request
}


// func main(){

// 	ApiRequest := NewRequestBuilder().Host("localhost:8090").Url("/users").Method(Get).AddHeader("Authorisation","Bearer Token").Body(nil).Build()

// 	fmt.Println(ApiRequest)

// }
