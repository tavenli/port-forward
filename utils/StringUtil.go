package utils

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ToInt(str string) int {
	_num, _ := strconv.Atoi(str)
	return _num
}

func ToInt64(str string) int64 {
	_num, _ := strconv.ParseInt(str, 10, 64)
	return _num
}

func ToInteger(str string) (int, error) {
	_num, _err := strconv.Atoi(str)
	return _num, _err
}

func ToLong(str string) (int64, error) {
	_num, _err := strconv.ParseInt(str, 10, 64)
	return _num, _err
}

func ToFloat64(str string) (float64, error) {
	_num, _err := strconv.ParseFloat(str, 64)
	return _num, _err
}

func BinaryToInt(str string) (int64, error) {
	_num, _err := strconv.ParseInt(str, 2, 64)
	return _num, _err
}

func IntToBinary(num int64) string {
	bin := strconv.FormatInt(num, 2)
	return bin
}

func IsBinaryOverInt(binStr string, number int64) bool {
	_num, _ := strconv.ParseInt(binStr, 2, 64)
	return (_num & number) == number
}

func IsBinNumOverInt(binNum int64, number int64) bool {

	return (binNum & number) == number
}

func ToStr(_num int) string {
	return strconv.Itoa(_num)
}

func FormatInt(_num int) string {
	return strconv.FormatInt(int64(_num), 10)
}

func FormatInt64(_num int64) string {
	return strconv.FormatInt(_num, 10)
}

func FormatFloat64(_num float64) string {
	return strconv.FormatFloat(_num, 'f', 2, 64)
}

func IsEmpty(str string) bool {

	return Len(str) <= 0
}

func IsNotEmpty(str string) bool {

	return !IsEmpty(str)
}

func Replace(str string, find string, to string) string {

	return strings.Replace(str, find, to, 1)
}

func ReplaceAll(str string, find string, to string) string {

	return strings.Replace(str, find, to, -1)
}

func Split(str string, spChar string) []string {

	return strings.Split(str, spChar)
}

func Contains(str string, find string) bool {

	return strings.Contains(str, find)
}

//	strings.HasPrefix("ABC_xyz", "ABC")
func StartsWith(str string, find string) bool {

	return strings.HasPrefix(str, find)
}

//	strings.HasSuffix("ABC_xyz", "xyz")
func EndsWith(str string, find string) bool {

	return strings.HasSuffix(str, find)
}

//  strings.Count("cheese", "e") = 3
func Count(str string, find string) int {

	return strings.Count(str, find)
}

//  返回第一个匹配字符的位置，返回-1为未找到
//  strings.Index("ABC_xyz", "xyz") = 4
//  strings.Index("ABC_xyz", "B") = 1
func Index(str string, find string) int {

	return strings.Index(str, find)
}

//strings.Join(arrays, ",") = "foo, bar, bas"
func Join(strs []string, spChar string) string {

	return strings.Join(strs, spChar)
}

//  字母转为小写
//  strings.ToLower("Love GoLang") = "love golang"
func ToLower(str string) string {

	return strings.ToLower(str)
}

//  字母转为大写
//  strings.ToTitle("love 中国") = "LOVE 中国"
func ToUpper(str string) string {
	return strings.ToUpper(str)
	//return strings.ToTitle(str)
}

func Len(str string) int {

	return len(str)
}

func Print(str string) {
	//var show = fmt.Println
	//show(str)
	fmt.Println(str)
}

func FilterByRegex(expr, input, placeTo string) string {
	regx, _ := regexp.Compile(expr)
	return regx.ReplaceAllString(input, placeTo)
}

func FilterStyle(input string) string {
	//regx, _ := regexp.Compile("<style((?:.|\\n)*?)</style>")
	regx, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	return regx.ReplaceAllString(input, "")
}

func FilterScript(input string) string {
	//regx, _ := regexp.Compile("<script((?:.|\\n)*?)</script>")
	regx, _ := regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	return regx.ReplaceAllString(input, "")
}

func FilterHtml(input string) string {
	regx, _ := regexp.Compile("<.+?>")
	//regx, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	return regx.ReplaceAllString(input, "")
}

func FilterA(input string) string {

	regx, _ := regexp.Compile("<.?a(.|\n)*?>")
	return regx.ReplaceAllString(input, "")
}

func FilterImage(input string) string {

	regx, _ := regexp.Compile("<img(.|\\n)*?>")
	return regx.ReplaceAllString(input, "")
}

func FilterSpecialChar(input string) string {

	regx, _ := regexp.Compile("[+=|{}':;',]")
	return regx.ReplaceAllString(input, "")
}

func FilterUrlPrefix(input string) string {

	regx, _ := regexp.Compile("\\w+://")
	return regx.ReplaceAllString(input, "")
}

func IsNumber(input string) bool {

	match, _ := regexp.MatchString("^\\d+$", input)
	return match
}

func IsIP(input string) bool {

	match, _ := regexp.MatchString("^((2[0-4]\\d|25[0-5]|[01]?\\d\\d?)\\.){3}(2[0-4]\\d|25[0-5]|[01]?\\d\\d?)$", input)
	return match
}

func IsEMail(input string) bool {

	match, _ := regexp.MatchString("^([a-z0-9A-Z]+[-|\\.]?)+[a-z0-9A-Z]@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)?\\.)+[a-zA-Z]{2,}$", input)
	return match
}

//高效拼接字符串
func LinkStrs(inputs ...string) string {
	var buf bytes.Buffer
	for _, v := range inputs {
		buf.WriteString(v)
	}
	return buf.String()
}

func LinkInputs(inputs ...interface{}) string {
	var buf bytes.Buffer
	for _, v := range inputs {
		switch t := v.(type) {
		case string:
			buf.WriteString(t)
		default:
			buf.WriteString(fmt.Sprint(t))

		}
	}
	return buf.String()
}

// func LinkInputs(inputs ...interface{}) string {
// 	var buf bytes.Buffer
// 	for _, v := range inputs {
// 		switch t := v.(type) {
// 		case string:
// 			buf.WriteString(t)
// 		//case int, int64:
// 		case int64:
// 			buf.WriteString(FormatInt64(t))
// 		case int:
// 			buf.WriteString(FormatInt(t))
// 		case float64:
// 			buf.WriteString(FormatFloat64(t))
// 		default:
// 			buf.WriteString(fmt.Sprint(t))

// 		}

// 		fmt.Println("v:", v)

// 	}
// 	return buf.String()
// }
