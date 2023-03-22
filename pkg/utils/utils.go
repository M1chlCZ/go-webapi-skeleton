package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	firebaseSDK "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"google.golang.org/api/option"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"io"
	_ "io/ioutil"
	"log"
	"math"
	mathRand "math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
	"unicode"
)

const (
	VERSION           = "0.0.0.1"
	STATUS     string = "status"
	OK         string = "OK"
	FAIL       string = "FAIL"
	ERROR      string = "hasError"
	ServerUrl  string = "184.174.35.183"
	SERVERPORT        = 7810
	MNPORT            = 7805
	STAKEPORT         = 7803
)

var colorReset = "\033[0m"
var colorRed = "\033[31m"

func InlineIFT[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}

func GetENV(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		WrapErrorLog("Error loading .env file")
	}
	return os.Getenv(key)
}

func ReportError(c *fiber.Ctx, err string, statusCode int) error {
	json := fiber.Map{
		"errorMessage": err,
		STATUS:         FAIL,
		ERROR:          true,
	}
	if statusCode == 500 {
		if !strings.Contains(err, "tx_id_UNIQUE") || strings.Contains(err, "Invalid Token, id RPCUser") {
			go logToFile(fmt.Sprintf("[WARNING] %s %s %s %s", "HTTP call failed : ", err, "  Status code: ", fmt.Sprintf("%d", statusCode)))
		}
	} else {
		go logToFile(fmt.Sprintf("[WARNING] %s %s %s %s", "HTTP call failed : ", err, "  Status code: ", fmt.Sprintf("%d", statusCode)))
	}
	return c.Status(statusCode).JSON(json)
}

func ReportOK(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		STATUS: OK,
		ERROR:  false,
	})
}

func ReportErrorSilent(c *fiber.Ctx, err string, statusCode int) error {
	json := fiber.Map{
		"errorMessage": err,
		STATUS:         FAIL,
		ERROR:          true,
	}

	return c.Status(statusCode).JSON(json)
}

//func CreateToken(userId uint64) (string, error) {
//	var err error
//	jwtKey := GetENV("JWT_KEY")
//
//	atClaims := jwt.MapClaims{}
//	atClaims["authorized"] = true
//	atClaims["idUser"] = userId
//	atClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
//	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
//	token, err := at.SignedString([]byte(jwtKey))
//	if err != nil {
//		return "", err
//	}
//	return token, nil
//}

func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func ScheduleFunc(f func(), interval time.Duration) *time.Ticker {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			f()

		}
	}()
	return ticker
}

var m sync.Mutex

func logToFile(message string) {
	m.Lock()
	defer m.Unlock()
	f, err := os.OpenFile("api.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v\n", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Println(message)
	log.Println("")
	_ = f.Close()
}

func WrapErrorLog(message string) {
	if !strings.Contains(message, "tx_id_UNIQUE") {
		go logToFile(fmt.Sprintf("[ERROR] %s", message))
	}
}

func WrapErrorLogF(message string, args ...any) {
	go logToFile(fmt.Sprintf(message, args))
}

func ReportWarning(message string) {
	if !strings.Contains(message, "tx_id_UNIQUE") {
		go logToFile(fmt.Sprintf("[WARNING] %s", message))
	}
}

func ReportSuccess(message string) {
	go logToFile(fmt.Sprintf("[SUCCESS] %s", message))
}

func ReportMessage(message string) {
	go logToFile(fmt.Sprintf("[INFO] %s", message))
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func GetHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func ArrContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GenerateInviteCode(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$&")
	b := make([]rune, length)
	for i := range b {
		if i%8 == 0 && i != 0 {
			b[i] = '-'
		} else {
			mathRand.Seed(time.Now().UnixNano())
			b[i] = letterRunes[mathRand.Intn(len(letterRunes))]
		}
	}
	return string(b)
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func Authorized(handler func(*fiber.Ctx) error) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if len(c.Get("Authorization")) == 0 {
			err := "no token provided"
			return ReportError(c, err, http.StatusUnauthorized)
		}

		tokenSplit := strings.Fields(c.Get("Authorization"))
		if len(tokenSplit) != 2 {
			return ReportErrorSilent(c, "Invalid Token", http.StatusUnauthorized)
		}

		if tokenSplit[0] == "JWT" {
			id, secret, err := ValidateKeyToken(tokenSplit[1])
			if secret != nil {
				decodeString, err := hex.DecodeString(secret.(string))
				if err != nil {
					return ReportErrorSilent(c, "Invalid Token", http.StatusUnauthorized)
				} else {
					length := len(decodeString)
					if length != 32 {
						return ReportErrorSilent(c, "Invalid Token", http.StatusUnauthorized)
					}
				}
			} else {
				return ReportErrorSilent(c, "Invalid Token", http.StatusUnauthorized)
			}

			if err != nil {
				return ReportError(c, "Invalid token", http.StatusUnauthorized)
			} else {
				c.Request().Header.Set("user_id", fmt.Sprintf("%d", id))
				c.Request().Header.Set("user_secret", secret.(string))
				return handler(c)
			}
		} else {
			return ReportError(c, "Invalid Token", http.StatusUnauthorized)
		}

	}
}

func GetUrl() (string, error) {
	nodeF, err := exec.Command("bash", "-c", "ifconfig | sed -En 's/127.0.0.1//;s/.*inet (addr:)?(([0-9]*\\.){3}[0-9]*).*/\\2/p'").Output()
	if err != nil {
		return "", err
	}
	node := strings.TrimSpace(string(nodeF))
	return node, nil
}

func SendMessage(token string, title string, body string, data map[string]string) {
	opts := []option.ClientOption{option.WithCredentialsFile("xdn-project.json")}
	c := &firebaseSDK.Config{
		ProjectID: "xdn-project",
	}
	firebase, err := firebaseSDK.NewApp(context.Background(), c, opts...)
	if err != nil {
		WrapErrorLog(err.Error())
		return
	}
	mess, err := firebase.Messaging(context.Background())
	if err != nil {
		WrapErrorLog(err.Error())
		return
	}
	_, err = mess.Send(context.Background(), &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Data:     data,
			Notification: &messaging.AndroidNotification{
				ChannelID: "xdn1",
				Title:     title,
				Body:      body,
				Icon:      "@drawable/ic_notification",
			},
		},
		Data:  data,
		Token: token, // a token that you received from a client
	})

	if err != nil {
		//WrapErrorLog(err.Error())
		return
	}
}
