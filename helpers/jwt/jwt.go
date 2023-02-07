package jwt

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/workjaedsada3/modules/helpers/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/golang-jwt/jwt/v4"
)

// simple key
// var sampleSecretKey = []byte("SecretYouShouldHide")
var AccessExpire = 3600 * time.Minute
var RefreshExpire = 3600 * time.Minute

type JWTConfig struct {
	AccessExpire   time.Duration
	RefreshExpire  time.Duration
	PrivateKeyFile string
	PublicKeyFile  string
}

func (j JWTConfig) Sign(data TokenDetail) Token {
	sampleSecretKey := loadRsaPrivateKey(j.PrivateKeyFile)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(sampleSecretKey)
	if err != nil {
		exception.InternalServerErrorException(err)
		return Token{}
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(AccessExpire).Unix()
	claims["authorized"] = true
	claims["access_uuid"] = utils.UUIDv4()
	claims["data"] = data
	refresh_token, _ := j.createRefreshToken()
	fmt.Println(refresh_token)
	// Generate encoded token and send it as response.
	if t, err := token.SignedString(privateKey); err != nil {
		exception.InternalServerErrorException(err)
		return Token{}
	} else {
		return Token{
			AccessToken:      t,
			ExpiresIn:        claims["exp"],
			TokenType:        "Bearer",
			RefreshToken:     refresh_token.RefreshToken,
			RefreshExpiresIn: refresh_token.RefreshExpiresIn,
		}
	}
}

func (j JWTConfig) Decode(token string) (jwt.MapClaims, error) {
	sampleSecretKey := loadRsaPublicKey(j.PublicKeyFile)
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(sampleSecretKey)
	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if err != nil {
			return publicKey, err
		}
		return publicKey, nil
	})
	if err != nil {
		return jwt.MapClaims{}, err
	}
	return claims["data"].(map[string]interface{}), nil
	// token_type := fmt.Sprintf("%s", claims["type"])
}

func (j JWTConfig) createRefreshToken() (RefreshToken, error) {
	sampleSecretKey := loadRsaPrivateKey(j.PrivateKeyFile)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(sampleSecretKey)
	if err != nil {
		log.Fatal(err)
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(RefreshExpire).Unix()
	claims["authorized"] = true
	claims["typ"] = "refresh_token"
	claims["refresh_uuid"] = utils.UUIDv4()
	// Generate encoded token and send it as response.
	if t, err := token.SignedString(privateKey); err != nil {
		log.Println(err.Error())
		return RefreshToken{}, err
	} else {
		return RefreshToken{
			RefreshExpiresIn: claims["exp"],
			RefreshToken:     t,
		}, nil
	}
}

func Accessible(c *fiber.Ctx) error {
	return c.SendString("Accessible")
}

func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}

func loadRsaPrivateKey(file string) []byte {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err.Error())
	}
	// start := strings.Split(string(bytes), "-----BEGIN PRIVATE KEY-----")
	// content := strings.Split(start[1], "-----END PRIVATE KEY-----")
	// key, err := ssh.ParseRawPrivateKey(bytes)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	return bytes
}

func loadRsaPublicKey(file string) []byte {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println(err.Error())
	}
	// start := strings.Split(string(bytes), "-----BEGIN PUBLIC KEY-----")
	// content := strings.Split(start[1], "-----END PUBLIC KEY-----")
	return bytes
	// key, err := ssh.ParsePublicKey(bytes)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// return key.(*rsa.PublicKey)
}
