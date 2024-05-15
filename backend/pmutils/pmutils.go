package pmutils

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"protos/parser"
	"reflect"
	"runtime"
	"strings"
	"unicode"

	"golang.org/x/exp/constraints"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func StrToBool(s string) bool {
	return s != ""
}

func SetupLogging(module string) {
	debug := StrToBool(GetEnv("DEBUG", ""))
	logCode := slog.LevelInfo
	if debug {
		logCode = slog.LevelDebug
	}
	slogOpts := slog.HandlerOptions{
		Level: logCode,
	}
	logFile, err := os.OpenFile(module+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file %s for output: %s", module+".log", err)
		panic(err)
	}

	handler := slog.NewJSONHandler(io.MultiWriter(os.Stdout, logFile), &slogOpts)

	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func ContainsEmptyString(ss ...string) bool {
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}

func InterSection[T constraints.Ordered](pS ...[]T) []T {
	hash := make(map[T]*int) // value, counter
	result := make([]T, 0)
	for _, slice := range pS {
		duplicationHash := make(map[T]bool) // duplication checking for individual slice
		for _, value := range slice {
			if _, isDup := duplicationHash[value]; !isDup { // is not duplicated in slice
				if counter := hash[value]; counter != nil { // is found in hash counter map
					if *counter++; *counter >= len(pS) { // is found in every slice
						result = append(result, value)
					}
				} else { // not found in hash counter map
					i := 1
					hash[value] = &i
				}
				duplicationHash[value] = true
			}
		}
	}
	return result
}

func BoolToStringList(b bool) *parser.StringList {
	if b {
		return &parser.StringList{Values: []string{"Да"}}
	}
	return &parser.StringList{Values: []string{"Нет"}}
}

func GetFunctionName(f interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func CapitalFirstLet(s string) string {
	r := []rune(s)
	return string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
}
