package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// Configはsisakulintの設定を表す構造体でのインスタンスは".github"に位置する"sisakulint.yml"を読み込んでparse
type Config struct {
	//selfhostedrunner : setting for self-hosted runner
	SelfHostedRunner struct {
		//Labelsはself-hosted runnerのラベル
		Labels []string `yaml:"labels"`
	} `yaml:"self-hosted-runner"`
	// ConfigVariablesはチェックされるworkflowで使用される設定変数の名前を示す
	//この値がnilの時にvarsのコンテキストのプロパティ名はチェックされない
	ConfigVariables []string `yaml:"config-variables"`
}

// parseConfigは与えられたbyte sliceをConfigにparseする
func parseConfig(b []byte, path string) (*Config, error) {
	var c Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		msg := strings.ReplaceAll(err.Error(), "\n", " ")
		return nil, fmt.Errorf("failed to parse config file %q: %s", path, msg)
	}
	return &c, nil
}
