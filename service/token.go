package service

import (
	"encoding/json"
	"net/url"
	"sync"
	"time"
	"work-wechat/config"
	"work-wechat/util"
)

type token struct {
	mutex        sync.RWMutex
	accessToken  string `json:"access_token"`
	refreshToken string
	lastRefresh  time.Time
	expiresIn    time.Duration `json:"expires_in"`
}

type response struct {
	ErrCode     int64         `json:"errcode"`
	ErrMsg      string        `json:"errmsg"`
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
}

/**
 * @Description: 获取 token
 * @receiver t
 * @return string
 */
func (t *token) GetToken() string {

	t.mutex.RLock()
	if t.accessToken == "" || time.Now().Sub(t.lastRefresh) > t.expiresIn {
		t.mutex.RUnlock()
		t.syncToken()
	} else {
		t.mutex.RUnlock()
	}

	return t.accessToken
}

/**
 * @Description: 调用接口
 * @receiver t
 */
func (t *token) syncToken() {

	t.mutex.Lock()
	defer t.mutex.Unlock()

	resp := util.HttpGet(config.WorkApiHost + "/cgi-bin/gettoken" + "?" + t.intoUrlValues())
	var respData response
	if error := json.Unmarshal(resp, &respData); error != nil {
		panic(error)
	}

	t.lastRefresh = time.Now()
	t.accessToken = respData.AccessToken
	t.expiresIn = respData.ExpiresIn * time.Second
}

/**
 * @Description: 拼凑获取 token
 * @receiver t
 * @return string
 */
func (t *token) intoUrlValues() string {

	var query = url.Values{}
	query.Add("corpid", config.App.Work.CorpId)
	query.Add("corpsecret", config.App.Work.CorpSecret)

	return query.Encode()
}
