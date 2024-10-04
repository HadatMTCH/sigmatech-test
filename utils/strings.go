package utils

import (
	"base-api/constants"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	idCode = "+62"
	male   = "M"
	female = "F"
)

var (
	phoneNumberRegex = regexp.MustCompile("^[+]?[0-9]{8,}$")
	emailRegex       = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// ValidateEmail return true if email address has a valid format
func ValidateEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// NormalizePhoneNumber standarized mobile phone number format
func NormalizePhoneNumber(number string) string {
	number = strings.Replace(strings.TrimSpace(number), " ", "", -1)

	if len(number) < 8 {
		return number
	}

	if strings.HasPrefix(number, "0") {
		return fmt.Sprintf("%s%s", idCode, number[1:])
	}

	if strings.HasPrefix(number, "+62") {
		return fmt.Sprintf("%s%s", idCode, number[3:])
	}

	origNumber := number
	for {
		if len(number) <= 2 || !strings.HasPrefix(number, "62") {
			break
		}

		number = number[2:]
	}

	if len(number) < 8 {
		number = origNumber
	}

	return fmt.Sprintf("%s%s", idCode, number)
}

// ValidatePhoneNumber return true if number has a valid phone number format
func ValidatePhoneNumber(number string) bool {
	number = NormalizePhoneNumber(number)
	return phoneNumberRegex.MatchString(number)
}

func trimQuote(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

// DateString return string representation of a date pointer
func DateString(dateTime *time.Time) string {
	if dateTime == nil {
		return ""
	}

	return dateTime.Format(constants.DateFormatStd)
}

func NullStringScan(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}

func NullBoolScanToString(x *bool) string {
	if x != nil {
		return strconv.FormatBool(*x)
	} else {
		return "-"
	}
}

func ConvertBytesToString(data []byte) string {
	return string(data[:])
}

func FormatMediaPath(rootPath string, value *string) string {
	if NullStringScan(value) != "" {
		return fmt.Sprintf("%s%s", rootPath, NullStringScan(value))
	} else {
		return ""
	}
}

func StructToByte(data interface{}) []byte {
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(data)
	if err != nil {
		logrus.Error(err)
	}

	return reqBodyBytes.Bytes()
}

func HourMinuteToTimeFormat(hour, minute int) string {
	hourString := fmt.Sprintf("00%d", hour)
	minuteString := fmt.Sprintf("00%d", minute)

	hourStringFormatted := hourString[len(hourString)-2:]
	minuteStringFormatted := minuteString[len(minuteString)-2:]

	return fmt.Sprintf("%s:%s", hourStringFormatted, minuteStringFormatted)
}

func TimeFormatToHourMinute(s string) (int, int, error) {
	timeArray := strings.Split(s, ":")
	hour, err := strconv.Atoi(timeArray[0])
	if err != nil {
		return 0, 0, err
	}
	minute, err := strconv.Atoi(timeArray[1])
	if err != nil {
		return 0, 0, err
	}

	return hour, minute, nil
}

func PrintStruct(data interface{}) {
	dataByte := StructToByte(data)
	fmt.Println(string(dataByte))
}

func Uid(length int) string {
	rand.Seed(time.Now().UnixNano())
	buf := []string{}
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-="
	charRune := []rune(chars)
	charlen := len(chars)

	for i := 0; i < length; i++ {
		index := rand.Intn(charlen)
		buf = append(buf, string(charRune[index]))
	}

	return strings.Join(buf, "")
}

func ConvertMapToString(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\" ", key, value)
	}
	return b.String()
}
