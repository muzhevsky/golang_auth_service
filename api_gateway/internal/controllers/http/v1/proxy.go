package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type proxy struct {
	redirectionURL *url.URL
	serviceName    string
}

func NewProxy(url *url.URL, serviceName string) *proxy {
	return &proxy{
		redirectionURL: url,
		serviceName:    serviceName,
	}
}

func (p *proxy) Handle(c *gin.Context) {
	urlPath := c.Request.URL.Path
	prefixIndex := strings.Index(urlPath[1:], "/") + 1
	urlPrefix := urlPath[1:prefixIndex]
	urlPostfix := urlPath[prefixIndex:]

	if urlPrefix != p.serviceName {
		c.Next()
		return
	}

	fullRequestURL, _ := url.Parse(p.redirectionURL.String() + urlPostfix)

	proxy := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.Out.URL = fullRequestURL
			fmt.Println(r.Out.URL)
		},
		ModifyResponse: func(response *http.Response) error {
			c.Set("proxyResponseCode", response.StatusCode)
			return nil
		},
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}
