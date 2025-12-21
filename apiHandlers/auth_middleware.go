package apiHandlers

import (
	"Suppliers/dto"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/patrickmn/go-cache"
)

type AuthMiddleware struct {
	config *dto.AuthConfig
	cache  *cache.Cache
}

type UserInfo struct {
	Sub          string `json:"sub"`
	Nickname     string `json:"nickname"`
	Name         string `json:"name"`
	Picture      string `json:"picture"`
	UpdatedAt    string `json:"updated_at"`
	Email        string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	UserMetadata struct {
		OrganizationId string `json:"organizationid"`
		Role            string `json:"role"`
	} `json:"user_metadata"`
}

func NewAuthMiddleware(config dto.AuthConfig) *AuthMiddleware {
	return &AuthMiddleware{
		config: &config,
		cache:  cache.New(2*time.Hour, 10*time.Minute),
	}
}

func (a *AuthMiddleware) ValidateToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	issuerURL, err := url.Parse("https://" + a.config.AUTH0_DOMAIN + "/")
	if err != nil {
		log.Fatalf("failed to parse the issuer url: %v", err)
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{a.config.AUTH0_AUDIENCE},
	)
	if err != nil {
		log.Fatalln("failed to set up the jwt validator")
	}

	_, err = jwtValidator.ValidateToken(c.Context(), authHeaderParts[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	if cachedData, found := a.cache.Get(authHeaderParts[1]); found {
		userData := cachedData.(dto.UserCacheData)
		c.Locals("user", userData)
		return c.Next()
	}

	userInfo, err := a.getUserInfo(authHeaderParts[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized access",
		})
	}

	userData := dto.UserCacheData{
		Sub:      userInfo.Sub,
		Role:     userInfo.UserMetadata.Role,
		OrganizationId: userInfo.UserMetadata.OrganizationId,
	}
	a.cache.Set(authHeaderParts[1], userData, 2*time.Hour)
	
	c.Locals("user", userData)
	return c.Next()
}

func (a *AuthMiddleware) getUserInfo(token string) (*UserInfo, error) {
	client := &http.Client{}
	UserInfoUrl := fmt.Sprintf("https://%s/userinfo", a.config.AUTH0_DOMAIN)
	req, err := http.NewRequest("GET", UserInfoUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo UserInfo
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}
