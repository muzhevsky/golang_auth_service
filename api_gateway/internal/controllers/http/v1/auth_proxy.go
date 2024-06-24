package v1

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type authProxy struct {
	authHost string
	client   *http.Client
}

func NewAuthProxy(authName string, authHost string) *authProxy {
	return &authProxy{authHost: authHost, client: &http.Client{}}
}

func (p *authProxy) Handle(c *gin.Context) {
	c.Request.Header.Del("account_id")
	token := c.Request.Header.Get("Authorization")

	if len(token) == 0 {
		c.Next()
		return
	}

	req, err := http.NewRequest("GET", p.authHost+"/authenticate", nil)
	if err != nil {
		c.Set("authError", "Failed to send HTTP request to authentication")
		c.Next()
		return
	}

	req.Header.Set("Authorization", token)

	resp, err := p.client.Do(req)
	if err != nil {
		c.Set("authError", "Failed to send HTTP request to authentication")
		c.Next()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Set("authError", "Failed to read response body of authentication request")
		c.Next()
		return
	}

	bodystr := string(body)

	if resp.StatusCode != http.StatusOK {
		c.Set("authError", bodystr)
		c.Next()
		return
	}

	_, err = strconv.Atoi(bodystr)
	if err != nil {
		c.Next()
		return
	}

	c.Request.Header.Add("account_id", bodystr)
}
