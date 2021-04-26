package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"work-wechat/config"
	"work-wechat/service/request"
	"work-wechat/util"
)

type WorkClient struct {
	config    *config.WorkConfig
	tokenInfo *token
}

var (
	WxApp = new(WorkClient)
	// 翻译器容器实例
	uni *ut.UniversalTranslator
	// 当前翻译器
	trans ut.Translator
)

/**
 * @Description: 初始化数据
 */
func init() {
	WxApp.config = config.App.Work
	WxApp.tokenInfo = new(token)

	zh := zh2.New()
	uni = ut.New(zh, zh)
	trans,_  = uni.GetTranslator("zh")
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhtrans.RegisterDefaultTranslations(v, trans)
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			fmt.Println(field.Tag.Get("label"))
			if label := field.Tag.Get("label"); label != ""{
				return label + " "
			}
			return field.Name
		})
	}
}

/**
 * @Description: 验证结果翻译
 * @receiver wx
 * @param err
 * @return []string
 */
func (wx *WorkClient) Translate(err error) []string  {
	var errorList []string
	errors := err.(validator.ValidationErrors)

	for _, e := range errors {
		errorList = append(errorList, e.Translate(trans))
	}

	return errorList
}

/**
 * @Description: 执行请求
 * @receiver wx
 * @param isToken 是否需要 access_token
 * @param method 请求方式
 * @param path 请求地址，可不包含 host
 * @param data 请求数据
 * @param headers 请求头
 * @return []byte
 */
func (wx *WorkClient) execQuery(isToken bool, method, path string, data interface{}, headers map[string]string) []byte {

	if isToken {
		if !strings.Contains(path, "http") && !strings.Contains(path, "https") {
			path = config.WorkApiHost + "/" + strings.TrimLeft(path, "/")
		}

		urlInfo, error := url.Parse(path)
		if error != nil {
			panic(error)
		}

		if len(urlInfo.RawQuery) == 0 || !strings.Contains(urlInfo.RawQuery, "access_token") {
			urlInfo.RawQuery += "&access_token=" + wx.tokenInfo.GetToken()
			urlInfo.RawQuery = strings.TrimLeft(urlInfo.RawQuery, "&")
		}

		path = urlInfo.String()
	}

	return util.DoRequest(method, path, data, headers)
}

/**
 * @Description: 获取打卡日报数据
 * @receiver wx
 * @param c
 */
func (wx *WorkClient) GetDailyPunch(c *gin.Context) {
	var data request.PunchBody
	if error := c.ShouldBindJSON(&data); error != nil {
		c.JSON(http.StatusUnprocessableEntity, wx.Translate(error))
		return
	}

	c.String(http.StatusOK, string(
		wx.execQuery(true,
		"POST",
		"cgi-bin/checkin/getcheckin_daydata",
		data,
		map[string]string{"content-type": "application/json; charset=utf-8"})))
}

/**
 * @Description: 获取打卡月报数据
 * @receiver wx
 * @param c
 */
func (wx *WorkClient) GetMonthPunch(c *gin.Context)  {
	var data request.PunchBody
	if error := c.ShouldBindJSON(&data); error != nil{
		c.JSON(http.StatusUnprocessableEntity, wx.Translate(error))
		return
	}

	c.String(http.StatusOK, string(
		wx.execQuery(true,
		"POST",
		"cgi-bin/checkin/getcheckin_monthdata",
		data,
		map[string]string{"content-type": "application/json; charset=utf-8"})))
}

/**
 * @Description: 获取打卡记录
 * @receiver wx
 * @param c
 */
func (wx *WorkClient) GetPunchRecord(c *gin.Context)  {
	var data request.PunchRecord
	if error := c.ShouldBindJSON(&data); error != nil {
		c.JSON(http.StatusUnprocessableEntity, wx.Translate(error))
		return
	}

	c.String(http.StatusOK, string(
		wx.execQuery(true,
			"POST",
			"cgi-bin/checkin/getcheckindata",
			data,
			map[string]string{"content-type": "application/json; charset=utf-8"})))
}
