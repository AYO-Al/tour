package word

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
	"unicode"
)

// 转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下划线转大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	titleCaser := cases.Title(language.Und)
	s = titleCaser.String(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线转小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var r []rune
	for i, v := range s {
		if i == 0 {
			r = append(r, unicode.ToLower(v))
			continue
		}
		if unicode.IsUpper(v) {
			r = append(r, '_') // 单引号
		}
		r = append(r, unicode.ToLower(v))
	}
	return string(r)
}
