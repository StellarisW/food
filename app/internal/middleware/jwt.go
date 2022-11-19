package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	g "main/app/global"
	"main/utils/cookie"
	myjwt "main/utils/jwt"
	"net/http"
	"time"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		var token string

		cookieConfig := g.Config.Auth.Cookie
		cookieWriter := cookie.NewCookieWriter(&cookie.Config{
			Secret: cookieConfig.Secret,
			Ctx:    c,
			Cookie: http.Cookie{
				Path:     "/",
				Domain:   cookieConfig.Domain,
				MaxAge:   cookieConfig.MaxAge,
				Secure:   cookieConfig.Secure,
				HttpOnly: cookieConfig.HttpOnly,
				SameSite: cookieConfig.SameSite,
			},
		})

		ok := cookieWriter.Get("x-token", &token)
		if token == "" || !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "not logged in",
				"ok":   false,
			})
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		// parseToken 解析token包含的信息
		jwtConfig := g.Config.Auth.Jwt
		j := myjwt.NewJWT(&myjwt.Config{
			SecretKey: jwtConfig.SecretKey},
		)

		mc, err := j.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
				"ok":   false,
			})
			c.Abort()
			return
		}

		if mc.ExpiresAt.Unix()-time.Now().Unix() < mc.BufferTime {
			mc.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(g.Config.Auth.Jwt.ExpiresTime) * time.Second))
			newToken, _ := j.GenerateToken(mc)
			newClaims, _ := j.ParseToken(newToken)
			cookieWriter.Set("x-token", newToken)
			err = g.Rdb.Set(c,
				fmt.Sprintf("jwt_%d", newClaims.BaseClaims.Id),
				newToken,
				time.Duration(jwtConfig.ExpiresTime)*time.Second).Err()
			if err != nil {
				g.Logger.Errorf("set [jwt] cache failed, %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "internal err",
					"ok":   false,
				})
				return
			}
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("id", mc.BaseClaims.Id)
		c.Set("username", mc.BaseClaims.Username)
		c.Next() // 后续的处理函数可以用过c.Search("username")来获取当前请求的用户信息
	}
}
