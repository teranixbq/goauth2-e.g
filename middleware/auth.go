package auth

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		token, err := ParseToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		data := token.Claims.(jwt.MapClaims)
		c.Locals("user", fiber.Map{
			"id":   data["id"],
			"role": data["role"],
		})

		return c.Next()
	}
}

func CreateToken(id string) (string, error) {
	godotenv.Load(".env")
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	if len(tokenString) >= 7 && tokenString[0:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	godotenv.Load(".env")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(c *fiber.Ctx) (string, error) {

	user := c.Locals("user")
	if user == nil {
		return "", errors.New("invalid token")
	}

	claims := user.(fiber.Map)

	id, ok := claims["id"].(string)
	if !ok {
		return "", errors.New("invalid token")
	}

	return id, nil
}
