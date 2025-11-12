package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var CONFIG Server

func InitConfig() error {
	// 1. 初始化 Viper
	v := viper.New()

	// 2. 设置配置文件信息
	v.SetConfigName("config") // 配置文件名称（不带后缀）
	v.SetConfigType("yaml")   // 配置文件类型（可选，Viper 会自动根据文件名推断）
	v.AddConfigPath(".")      // 配置文件路径（当前目录）
	// 可添加多个路径：v.AddConfigPath("/etc/myapp")、v.AddConfigPath("$HOME/.myapp")

	// 3. 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}

	// 4. 将配置映射到结构体
	if err := v.Unmarshal(&CONFIG); err != nil {
		return fmt.Errorf("配置映射到结构体失败: %w", err)
	}

	return nil
}
