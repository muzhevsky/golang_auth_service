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
		c.Next()
		return
	}

	req.Header.Set("Authorization", token)

	resp, err := p.client.Do(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to send HTTP request"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.Next()
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	bodystr := string(body)
	_, err = strconv.Atoi(bodystr)
	if err != nil {
		c.Next()
		return
	}

	c.Request.Header.Add("account_id", bodystr)
}
