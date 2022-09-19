package function

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

//发送消息到钉钉
func SendToDingDing(push_url string, title string, info string) {
	nowTime := GetNowTime()
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  info + "\n > ###### " + nowTime + "发布 报表 ",
			"at": map[string]interface{}{
				"atMobiles": "",
				"isAtAll":   "TRUE",
			},
		},
	}

	HttpRequestPost(push_url, data)
}

// 发送消息企业微信
func SendToWeixin(push_url string, info string) {
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"content": info,
		},
	}

	HttpRequestPost(push_url, data)
}

// HttpRequestPost /*url 请求地址 param请求参数
func HttpRequestPost(url string, param interface{}) string {
	defer func() {
		//捕获panic
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	client := &http.Client{}
	bytesData, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	postBody := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", url, postBody)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	resp, _ := client.Do(req)
	//返回内容
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	return string(body)
}

// HttpRequestPost /*url 请求地址 param请求参数
func HttpRequestPostCookie(url string, param interface{}, cookie string) string {
	defer func() {
		//捕获panic
		if err := recover(); err != nil {
			panic(err)
		}
	}()
	client := &http.Client{}
	bytesData, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	postBody := bytes.NewReader(bytesData)
	req, err := http.NewRequest("POST", url, postBody)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	req.Header.Add("Cookie", cookie)
	resp, _ := client.Do(req)
	//返回内容
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// HttpRequestPost /*url 请求地址 param请求参数
func HttpRequestGetCookie(url string, param interface{}, cookie string) string {

	defer func() {
		//捕获panic
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	client := &http.Client{}
	bytesData, err := json.Marshal(param)
	if err != nil {
		panic(err)
	}

	postBody := bytes.NewReader(bytesData)
	req, err := http.NewRequest("GET", url, postBody)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-type", "application/json;charset=utf-8")
	req.Header.Add("Cookie", cookie)
	resp, _ := client.Do(req)
	//返回内容
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// RunFuncName 获取当前运行的方法名称
func RunFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// 判断字符串是否在数组中
func InArray(target string, str_array []string) bool {
	for _, element := range str_array {

		if target == element {
			return true
		}

	}

	return false
}

func GetNowTime() string {
	DateTime := "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(DateTime)
}

func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	errs := json.Unmarshal([]byte(jsonStr), &m)
	if errs != nil {
		return nil, errs
	}
	return m, nil
}

func MapToJson(m map[string]string) (string, error) {
	jsonByte, errs := json.Marshal(m)

	if errs != nil {
		return "", errs
	}

	return string(jsonByte), nil
}

//过滤HTML 标签
func TrimHtml(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	return strings.TrimSpace(src)
}

func ArrayColumn(array interface{}, key string) (result map[string]interface{}, err error) {
	result = make(map[string]interface{})
	t := reflect.TypeOf(array)
	v := reflect.ValueOf(array)
	if t.Kind() != reflect.Slice {
		return nil, errors.New("array type not slice")
	}
	if v.Len() == 0 {
		return nil, errors.New("array len is zero")
	}

	for i := 0; i < v.Len(); i++ {
		indexv := v.Index(i)
		if indexv.Type().Kind() != reflect.Struct {
			return nil, errors.New("element type not struct")
		}
		mapKeyInterface := indexv.FieldByName(key)
		if mapKeyInterface.Kind() == reflect.Invalid {
			return nil, errors.New("key not exist")
		}
		mapKeyString, err := interfaceToString(mapKeyInterface.Interface())
		if err != nil {
			return nil, err
		}
		result[mapKeyString] = indexv.Interface()
	}
	return result, err
}

func interfaceToString(v interface{}) (result string, err error) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		result = fmt.Sprintf("%v", v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result = fmt.Sprintf("%v", v)
	case reflect.String:
		result = v.(string)
	default:
		err = errors.New("can't transition to string")
	}
	return result, err
}

//显示图标
func ShowLogo() {
	fmt.Println("\n        _                _                                                _ \n       (_)              | |                                              (_)\n  __ _  _  _ __  ______ | |__    ___   _ __    __ _  ______  __ _  _ __   _ \n / _` || || '_ \\|______|| '_ \\  / _ \\ | '_ \\  / _` ||______|/ _` || '_ \\ | |\n| (_| || || | | |       | | | || (_) || | | || (_| |       | (_| || |_) || |\n \\__, ||_||_| |_|       |_| |_| \\___/ |_| |_| \\__, |        \\__,_|| .__/ |_|\n  __/ |                                        __/ |              | |       \n |___/                                        |___/               |_|       \n")
}

// ErrorToString 错误类型转字符串
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		//error错误类型 发送钉钉发
		msg := v.Error()
		return msg
	default:
		return r.(string)
	}
}

// StuctSliceToSliceByField 将结构体切片中指定字段存入切片
func StuctSliceToSliceByField(array interface{}, key string) (result []string, err error) {
	result = make([]string, 10)
	t := reflect.TypeOf(array)
	v := reflect.ValueOf(array)
	if t.Kind() != reflect.Slice {
		return nil, errors.New("array type not slice")
	}
	if v.Len() == 0 {
		return nil, errors.New("array len is zero")
	}

	for i := 0; i < v.Len(); i++ {
		indexv := v.Index(i)
		if indexv.Type().Kind() != reflect.Struct {
			return nil, errors.New("element type not struct")
		}
		mapKeyInterface := indexv.FieldByName(key)
		if mapKeyInterface.Kind() == reflect.Invalid {
			return nil, errors.New("key not exist")
		}
		mapKeyString, err := interfaceToString(mapKeyInterface.Interface())
		if err != nil {
			return nil, err
		}
		result = append(result, mapKeyString)
	}
	return result, err
}

//只取文本中的中文以及a-zA-Z0-9
func GetZhAndLetter(words string) string {
	reg := regexp.MustCompile("[\u4e00-\u9fa5_a-zA-Z0-9]+") // 查找连续的非小写字母
	wordsSlice := reg.FindAllString(words, -1)
	var str string
	if len(wordsSlice) > 0 {
		for _, item := range wordsSlice {
			str += item
		}
	}
	return str
}

//判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

//创建文件夹
func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		panic("获取文件夹异常 -> %v")
		return
	}
	if _exist {

	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

// 字符串转时间
func ParseDuration(d string) (time.Duration, error) {
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.HasSuffix(d, "d") {
		h := strings.TrimSuffix(d, "d")
		hour, _ := strconv.Atoi(h)
		dr = time.Hour * 24 * time.Duration(hour)
		return dr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

//统计代码块执行时间
func ExeStart() int64 {
	return time.Now().UnixMilli()
}

//统计代码块执行时间
func PrintUseTime(startTime int64) {
	nowTime := time.Now().UnixMilli()
	fmt.Printf("程序耗时:%d", nowTime-startTime)
}
