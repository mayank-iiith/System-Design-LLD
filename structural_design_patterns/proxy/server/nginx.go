package server

type Nginx struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNginxServer() *Nginx {
	return &Nginx{
		application:       &Application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func (n *Nginx) HandleRequest(url, method string) (int, string) {
	if n.isRateLimitExceeded(url) {
		return 403, "Not Allowed"
	}
	return n.application.HandleRequest(url, method)
}

func (n *Nginx) isRateLimitExceeded(url string) bool {
	if n.rateLimiter[url]+1 > n.maxAllowedRequest {
		return true
	}
	n.rateLimiter[url]++
	return false
}
