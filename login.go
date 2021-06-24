package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type rsp struct {
	Message string `json:"msg"`
}

func (l LYExamination) Login(_ context.Context, in *LoginReq) (*LoginRsp, error) {

	req, err := buildRequest(in)
	if err != nil {
		// TODO 日志打印
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// TODO 日志打印
		return nil, err
	}

	rsp, err := parseResponse(resp)
	if err != nil {
		// TODO 日志打印
		return nil, err
	}

	switch rsp.Message {
	case "ok":
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "PHPSESSID" {
				return &LoginRsp{Session: cookie.Value}, nil
			}
		}

		return nil, errors.New("不可企及的会话 ID")

	case "mimacuowu":
		return nil, errors.New("密码错误")

	case "shangweizhuce":
		return nil, errors.New("帐号不存在")

	default:
		return nil, errors.New("未知错误。远程状态：" + rsp.Message)
	}
}

func buildRequest(in *LoginReq) (*http.Request, error) {

	formData := make(url.Values)
	formData.Set("Login_phone", in.GetPhone())
	formData.Set("parpwd", in.GetPwd())

	req, err := http.NewRequest(
		"POST",
		"https://mic.fjjxhl.com/pcnews/index.php/Home/User/parlogin",
		strings.NewReader(formData.Encode()),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "*/*")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("origin", "https://mic.fjjxhl.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://mic.fjjxhl.com/pcnews/index.php/Home/User/pclogin.html")
	req.Header.Set("sec-ch-ua", "\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"91\", \"Chromium\";v=\"91\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36 Edg/91.0.864.54")
	req.Header.Set("x-requested-with", "XMLHttpRequest")

	return req, nil
}

func parseResponse(resp *http.Response) (*rsp, error) {

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	goRsp := new(rsp)
	err = json.Unmarshal(bytes, goRsp)
	if err != nil {
		return nil, err
	}

	return goRsp, nil
}
