package setting

import "github.com/spf13/viper"

/**
* @Author $
* @Description //TODO $
* @Date $ $
* @Param $
* @return $
**/

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string) (*Setting, error) {
	vp := viper.New()
	// 设置配置文件名称
	vp.SetConfigName("config")

	// 设置配置文件地址
	for _, config := range configs {
		if config != "" {
			vp.AddConfigPath(config)
		}
	}
	// 设置配置文件类型
	vp.SetConfigType("yaml")
	// 读取文件内容
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
