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

		//ctx is needed by Redis to handle timeouts or cancellations in case the request takes too long or is aborted —
		// it helps manage long-running or stuck operations.
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
		// This is how the key value store look in my redis in this case
		//  Key:   rate_limit:10.0.0.5
		//  Value: 1
		//  TTL:   59m 50s
		// Its not strict that my key needs to be an integer only in redis

		// Continue to handler
		c.Next()
	}
}

// Gin expects middleware like this:
// func RateLimiter(maxRequests int, duration time.Duration) gin.HandlerFunc {
//     return func(c *gin.Context) {
//         // ← c is available *here*, because this inner function is called per request
//     }
// }
// So we define an outer function to capture the params (maxRequests, duration), and an inner function that Gin calls later with *gin.Context when a request comes in.
// This is a standard pattern to use whe we need to pass params to any middleware
