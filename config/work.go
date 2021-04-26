package config

/**
 * @Description: 企业微信接口地址
 */
const WorkApiHost = "https://qyapi.weixin.qq.com"

type WorkConfig struct {
	CorpId     string `yaml:"corpId"`
	AgentId    int    `yaml:"agentId"`
	CorpSecret string `yaml:"corpSecret"`
}
