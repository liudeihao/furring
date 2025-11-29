package initialize

import (
    "github.com/liudeihao/furring/config"
    "github.com/spf13/viper"
)

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        panic("Viper读取配置失败" + err.Error())
    }
    if err := viper.Unmarshal(&config.Instance); err != nil {
        panic("Viper传入config.Instance失败" + err.Error())
    }
}
