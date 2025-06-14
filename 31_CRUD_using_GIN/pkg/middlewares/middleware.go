package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Abhishek-2400/crud_gin/pkg/config"
	model "github.com/Abhishek-2400/crud_gin/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var ctx context.Context

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		idFloat := claims["user_id"].(float64)
		id := int64(idFloat)

		user, error := model.FindUser(id)
		if error != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user", user)
		//c.Set("user", user) stores the user data in the Gin context so it can be accessed later in the same request,
		// like in route handlers, without querying the database again.
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}

func RateLimiter(maxRequests int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		redisClient = config.GetRedisClient()
		ctx = config.GetContext()
		ip := c.ClientIP()
		key := fmt.Sprintf("rate_limit:%s", ip)

		// Increment request count
		count, err := redisClient.Incr(ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Redis error"})
			return
		}

		if count == 1 {
			// Set expiration only the first time
			redisClient.Expire(ctx, key, duration)
		}

		if count > int64(maxRequests) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Try again later.",
			})
			return
		}

		// Continue to handler
		c.Next()
	}
}
