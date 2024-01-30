package main

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/open-policy-agent/opa/rego"
	"gopkg.in/yaml.v3"
)

// LoadYAMLFiles は指定されたディレクトリ内のYAMLファイルを読み込みます。
func LoadYAMLFiles(dir string) ([]map[string]interface{}, error) {
	var files []map[string]interface{}

	// ディレクトリ内のファイルを走査
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// YAMLファイルを読み込み
		if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var data map[string]interface{}
			if err := yaml.Unmarshal(file, &data); err != nil {
				return err
			}

			files = append(files, data)
		}
		return nil
	})

	return files, err
}

// ProcessResults はrego.ResultSetを処理し、重複しないメッセージを出力します。
func ProcessResults(results rego.ResultSet, seenMessages map[string]bool) error {
	if len(results) == 0 {
		return nil
	}

	// 結果を解析し、メッセージを出力
	for _, result := range results {
		for _, expression := range result.Expressions {
			msgs, ok := expression.Value.([]interface{})
			if !ok {
				continue
			}

			for _, msg := range msgs {
				msgStr, ok := msg.(string)
				if !ok {
					continue
				}

				if _, seen := seenMessages[msgStr]; !seen {
					// Regoクエリからのメッセージを整形して出力
					return &PolicyError{
						File:    ".github/workflows",
						Line:    0,
						Column:  0,
						Message: msgStr,
					}
				}
			}
		}
	}
	return nil
}

func Policy() error {
	// ワークフローファイルを読み込み
	files, err := LoadYAMLFiles(".github/workflows")
	if err != nil {
		return err
	}

	// Regoポリシーファイルを読み込み
	policy, err := os.ReadFile("script/issueinjection.rego")
	if err != nil {
		return err
	}

	// Regoクエリを準備
	r := rego.New(
		rego.Query("data.core.deny"),
		rego.Module("issueinjection.rego", string(policy)),
	)

	query, err := r.PrepareForEval(context.Background())
	if err != nil {
		return err
	}

	seenMessages := make(map[string]bool)

	// 各ファイルに対してRegoクエリを実行
	for _, file := range files {
		results, err := query.Eval(context.Background(), rego.EvalInput(file))
		if err != nil {
			return err
		}

		if err := ProcessResults(results, seenMessages); err != nil {
			return err
		}
	}
	return nil
}
