package sysconfig

import (
	messagev1alpha1 "github.com/myoperator/messageoperator/pkg/apis/message/v1alpha1"
	"github.com/myoperator/messageoperator/pkg/common"
	"io/ioutil"
	"os"
	"sigs.k8s.io/yaml"
)



var SysConfig1 = new(SysConfig)

func InitConfig() error {
	// 读取yaml配置
	config, err := ioutil.ReadFile("./app.yaml")
	if err != nil {
		return err
	}


	err = yaml.Unmarshal(config, SysConfig1)
	if err != nil {
		return err
	}


	return nil

}

type SysConfig struct {
	Sender Sender  `yaml:"sender"`
	Server Server  `yaml:"server"`
}

type Sender struct {
	Remote   string `json:"remote"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Targets  string `json:"targets"`
}

type Server struct {
	Ip     string `yaml:"ip"`
	Port   int 	  `yaml:"port"`
}

func AppConfig(message *messagev1alpha1.Message) error {

	SysConfig1.Sender.Remote = message.Spec.Sender.Remote
	SysConfig1.Sender.Email = message.Spec.Sender.Email
	SysConfig1.Sender.Targets = message.Spec.Sender.Targets
	SysConfig1.Sender.Port = message.Spec.Sender.Port
	SysConfig1.Sender.Password = message.Spec.Sender.Password

	// 保存配置文件
	if err := saveConfigToFile(); err != nil {
		return err
	}

	return ReloadConfig()
}

// ReloadConfig 重载配置
func ReloadConfig() error {
	return InitConfig()
}

//saveConfigToFile 把config配置放入文件中
func saveConfigToFile() error {

	b, err := yaml.Marshal(SysConfig1)
	if err != nil {
		return err
	}
	// 读取文件
	path := common.GetWd()
	filePath := path + "/app.yaml"
	appYamlFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 644)
	if err != nil {
		return err
	}

	defer appYamlFile.Close()
	_, err = appYamlFile.Write(b)
	if err != nil {
		return err
	}

	return nil
}