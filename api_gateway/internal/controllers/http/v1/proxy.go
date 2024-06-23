package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type proxy struct {
	serviceMap map[string]string
}

func NewProxy(serviceMap map[string]string) *proxy {
	return &proxy{serviceMap: serviceMap}
}

func (p *proxy) Handle(c *gin.Context) {
	urlPath := c.Request.URL.Path
	prefixIndex := strings.Index(urlPath[1:], "/") + 1
	urlPrefix := urlPath[1:prefixIndex]
	urlPostfix := urlPath[prefixIndex:]

	var redirectionURL *string

	for k, v := range p.serviceMap {
		if urlPrefix == k {
			redirectionURL = &v
			break
		}
	}

	if redirectionURL == nil {
		c.Next()
		return
	}

	fullRequestRawURL := *redirectionURL + urlPostfix
	fullRequestURL, err := url.Parse(fullRequestRawURL)
	if err != nil {
		c.Next()
		return
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.Out.URL = fullRequestURL
		},
		ModifyResponse: func(response *http.Response) error {
			c.Set("proxyResponseCode", response.StatusCode)
			return nil
		},
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
