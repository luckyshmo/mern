package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/luckyshmo/api-example/models"
	"github.com/luckyshmo/api-example/pkg/repository"
)

//TODO! STORE not in open repo. It's OK only for demo
const (
	//https://en.wikipedia.org/wiki/Salt_(cryptography)
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "asdna12jk%s2fnv4ks#jd12nvs" //totally random key for encryption and decryption
	tokenTTL   = 24 * time.Hour               //Token Time To Live //? any requirements or frontend sync?
)

//Custom JWT claims //? What also we need to store in JWT?
type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (uuid.UUID, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {

	//Get user from DB
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	//HMAC-SHA256 algorithm, has several security issues //? change? //https://habr.com/en/post/450054/
	//claims are pieces of information asserted about a subject
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	//Sign token
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (uuid.UUID, error) {
	//using same structure for parse claims from token
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		//func witch will receive the parsed token and should return the key for validating
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { //cast syntax
			return nil, errors.New("invalid signing method")
		}

		//returns ours singingKey to parent func to encrypt token
		return []byte(signingKey), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*tokenClaims) //cast syntax
	if !ok {
		return uuid.Nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
