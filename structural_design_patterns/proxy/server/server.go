package server

type Server interface {
	HandleRequest(url, method string) (int, string)
}
