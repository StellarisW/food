package cookie

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type (
	Cookie struct {
		Config *Config
	}

	Config struct {
		Secret string
		Ctx    *gin.Context
		http.Cookie
	}
)

func NewCookieWriter(config *Config) *Cookie {
	return &Cookie{
		Config: config,
	}
}

func (c *Cookie) Set(key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	setSecureCookie(c, key, string(bytes))
}

func (c *Cookie) Get(key string, obj interface{}) bool {
	tempData, ok := getSecureCookie(c, key)
	if !ok {
		return false
	}
	_ = json.Unmarshal([]byte(tempData), obj)
	return true
}

func (c *Cookie) Remove(key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	setSecureCookie(c, key, string(bytes))
}

func setSecureCookie(c *Cookie, name, value string) {
	vs := base64.URLEncoding.EncodeToString([]byte(value))
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	h := hmac.New(sha256.New, []byte(c.Config.Secret))
	_, _ = fmt.Fprintf(h, "%s%s", vs, timestamp)

	sig := fmt.Sprintf("%02x", h.Sum(nil))
	cookie := strings.Join([]string{vs, timestamp, sig}, "|")

	http.SetCookie(c.Config.Ctx.Writer, &http.Cookie{
		Name:     name,
		Value:    cookie,
		MaxAge:   c.Config.MaxAge,
		Path:     "/",
		Domain:   c.Config.Domain,
		SameSite: http.SameSite(1),
		Secure:   c.Config.Secure,
		HttpOnly: c.Config.HttpOnly,
	})
}

func getSecureCookie(c *Cookie, key string) (string, bool) {
	cookie, err := c.Config.Ctx.Request.Cookie(key)
	if err != nil {
		return "", false
	}
	val, err := url.QueryUnescape(cookie.Value)
	if val == "" || err != nil {
		return "", false
	}

	parts := strings.SplitN(val, "|", 3)
	if len(parts) != 3 {
		return "", false
	}

	vs := parts[0]
	timestamp := parts[1]
	sig := parts[2]

	h := hmac.New(sha256.New, []byte(c.Config.Secret))
	_, _ = fmt.Fprintf(h, "%s%s", vs, timestamp)

	if fmt.Sprintf("%02x", h.Sum(nil)) != sig {
		return "", false
	}
	res, _ := base64.URLEncoding.DecodeString(vs)
	return string(res), true
}
