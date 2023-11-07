package config

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
)

var (
	DBConfig DatabaseConfig
)

func InitConfig() {
	InitDB("database", "FLASH_PASS")
}

func InitDB(dataId string, group string) {
	// TODO：部署时，test & pro 环境需要设置这个环境变量
	namespace := os.Getenv("FLASH_PASS_NAMESPACE")
	if namespace == "" {
		namespace = "386a677f-cc4f-40f3-b596-ee991acf2a68"
	}
	user := os.Getenv("FLASH_PASS_USER")

	sc := []constant.ServerConfig{
		{
			IpAddr: "0.0.0.0",
			Port:   18848,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId: namespace,
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		log.Println("创建 Nacos 配置客户端失败:", err)
		return
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		log.Println("获取 Nacos 配置失败:", err)
		return
	}

	err = json.Unmarshal([]byte(content), &DBConfig)
	if err != nil {
		log.Println("解析配置失败:", err)
		return
	}
	// 本地 dev 环境开发人员的 db 相互独立，通过 username 区分，该 db 需要事先联系管理员创建
	if DBConfig.Database == "" {
		DBConfig.Database = user
	}
	//log.Println(DBConfig)
}
