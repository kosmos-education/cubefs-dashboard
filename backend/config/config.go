// Copyright 2023 The CubeFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package config

import (
	"github.com/spf13/viper"
)

var Conf *Config

func Init(configPath string) (err error) {
	serverConfig := viper.New()
	mysqlConfig := viper.New()
	serverConfig.SetConfigFile(configPath + "config.yaml")
	mysqlConfig.SetConfigFile(configPath + "mysql.yaml")

	err = serverConfig.ReadInConfig()
	if err != nil {
		return err
	}

	err = mysqlConfig.ReadInConfig()
	if err != nil {
		return err
	}

	err = serverConfig.Unmarshal(&Conf)
	if err != nil {
		return err
	}

	err = mysqlConfig.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	return
}

type Config struct {
	Server *ServerConfig `mapstructure:"server"`
	Prefix *PrefixConfig `mapstructure:"prefix"`
	Mysql  *MysqlConfig  `mapstructure:"mysql"`
}

type ServerConfig struct {
	Port           int            `mapstructure:"port"`
	Mode           string         `mapstructure:"mode"`
	ClientIDKey    string         `mapstructure:"clientIDKey"`
	StaticResource StaticResource `mapstructure:"static_resource"`
}

type StaticResource struct {
	Enable       bool   `mapstructure:"enable"`
	RelativePath string `mapstructure:"relative_path"`
	RootPath     string `mapstructure:"root_path"`
}

type PrefixConfig struct {
	Api string `mapstructure:"api"`
}

type MysqlConfig struct {
	Host        string `mapstructure:"host"`
	Port        string `mapstructure:"port"`
	SlaveHost   string `mapstructure:"slaveHost"`
	SlavePort   string `mapstructure:"slavePort"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Database    string `mapstructure:"database"`
	MaxIdleConn int    `mapstructure:"maxIdleConn"`
	MaxOpenConn int    `mapstructure:"maxOpenConn"`
}
