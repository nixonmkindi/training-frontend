package server

import (
	"fmt"
	"html/template"
	"reflect"
	"strings"
	"time"
	"training-frontend/package/time_parser"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Renderer fetches the template render
// Renderer fetches the template renderer
func Renderer() *echoview.ViewEngine {
	gvc := goview.Config{
		Root:      "server/systems",
		Extension: ".html",
		Master:    "layouts/master",
		Funcs: template.FuncMap{
			"formatDateString": func(dateStr string) string {
				// Attempt to parse the string as a time
				parsedTime, err := time.Parse("2006-01-02 15:04:05 +0000 UTC", dateStr)
				if err != nil {
					return dateStr // Return the original string if parsing fails
				}
				return parsedTime.Format("2006-01-02")
			},
			"sub": func(a, b int) int {
				return a - b
			},
			"add": func(a, b int) int {
				return a + b
			},
			"inc": func(i int) int {
				return i + 1
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
			"timestamp": func(t time.Time) string {
				return t.Format("02/01/2006 15:04")
			},
			"menu": func(path string, data ...string) string {
				for _, v := range data {
					if strings.HasPrefix(path, v) {
						return "menu-open"
					}
				}
				return ""
			},
			"active": func(path string, data ...string) string {
				for _, v := range data {
					if strings.HasPrefix(path, v) {
						return "active"
					}
				}
				return ""
			},
			"hasRole": func(roles []string, role string) bool {
				return hasOne(roles, role)
			},
			"hasAnyRole": func(roles []string, role ...string) bool {
				return hasAny(role, roles)
			},
			"hasPermission": func(permissions []string, permission string) bool {
				return hasOne(permissions, permission)
			},
			"subfloat": func(a, b float32) string {
				return fmt.Sprintf("%0.2f", a-b)
			},
			"humanitize": func(t time.Time) string {
				return time_parser.TimeDuration(t)
			},
			"addfloat": func(a, b float32) string {
				return fmt.Sprintf("%0.2f", a+b)
			},
			"amountInFloat": func(a float32) string {
				return fmt.Sprintf("%0.2f", a)
			},
			"amountInFloatOneDecimal": func(a float32) string {
				return fmt.Sprintf("%0.1f", a)
			},
			"amountInFloats": func(a float64) string {
				return fmt.Sprintf("%0.2f", a)
			},
			"lengthOfArray": func(arr interface{}) int {
				arrValue := reflect.ValueOf(arr)
				if arrValue.Kind() != reflect.Array && arrValue.Kind() != reflect.Slice {
					return 0
				}
				return arrValue.Len()
			},
			"formaldate": func(dateStr string) string {
				t, _ := time.Parse("2006-01-02", dateStr)
				return t.Format("02/01/2006")
			},
			"split":     strings.Split,
			"camelcase": CamelCase,
		},
		DisableCache: true,
	}
	return echoview.New(gvc)
}

// hasAny returns true if one array element is contained in another array
func hasAny(s1 []string, s2 []string) bool {
	for _, a := range s1 {
		for _, b := range s2 {
			if a == b {
				return true
			}
		}
	}
	return false
}

// hasOne returns true if a string s2 is contained in the array s1
func hasOne(s1 []string, s2 string) bool {
	for _, s := range s1 {
		if s == s2 {
			return true
		}
	}
	return false
}

// CamelCase converts a string to camel case.
func CamelCase(input string) string {
	// Convert the input to lowercase
	lc := language.English
	mapper := cases.Title(lc)

	// Split the string into words
	words := strings.Fields(input)

	// Capitalize the first letter of each word
	for i, word := range words {
		words[i] = mapper.String(word)
	}

	// Join the words back together with spaces
	result := strings.Join(words, " ")

	return result
}
