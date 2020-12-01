package NetUtils

import (
	"errors"
	"forward-core/Utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"bytes"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func GetIP(c *beego.Controller) string {
	//utils.GetIP(&c.Controller)
	//也可以直接用 c.Ctx.Input.IP() 取真实IP
	ip := c.Ctx.Request.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip = c.Ctx.Request.Header.Get("Remote_addr")
	if ip == "" {
		ip = c.Ctx.Request.RemoteAddr
	}
	return ip
}

func HttpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		logs.Error("HttpGet error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		//logs.Debug("HttpGet result: ", result)
	} else {
		logs.Error("HttpGet error: ", err)
	}

	return result, nil
}

func HttpPostJsonReturnByte(url string, json string) ([]byte, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(json))
	if err != nil {
		logs.Error("HttpPostJson error: ", err)
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("返回对象为空")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return body, err
		//logs.Debug("HttpPostJson result: ", result)
	} else {
		logs.Error("HttpPostJson error: ", err)
		return nil, err
	}

}

func HttpPost(url string, param map[string]string) (string, error) {

	var paramBuf bytes.Buffer
	paramBuf.WriteString("curTime=" + Utils.GetCurrentTime())
	for k, v := range param {
		paramBuf.WriteString("&" + k + "=" + v)
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(paramBuf.String()))
	if err != nil {
		logs.Error("HttpPost error: ", err)
		return "", err
	}

	if resp == nil {
		return "", errors.New("返回对象为空")
	}

	defer resp.Body.Close()
	result := ""
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		result = string(body)
		logs.Debug("HttpPost result: ", result)
	} else {
		logs.Error("HttpPost error: ", err)
	}

	return result, nil
}

func UrlEncode(input string) string {
	if Utils.IsEmpty(input) {
		return ""
	}
	return url.QueryEscape(input)
}

func UrlDecode(input string) string {
	if Utils.IsEmpty(input) {
		return ""
	}
	result, err := url.QueryUnescape(input)
	if err != nil {
		return input
	} else {
		return result
	}
}
