package root

import (
	"context"
	"fmt"
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
func ProcessResults(results rego.ResultSet, seenMessages map[string]bool) {
	if len(results) == 0 {
		return
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
					fmt.Printf("Warning: %s\n", msgStr)
					seenMessages[msgStr] = true
				}
			}
		}
	}
}

func policy() {
	// ワークフローファイルを読み込み
	files, err := LoadYAMLFiles("script")
	if err != nil {
		fmt.Println("Error reading policy files:", err)
		return
	}

	// Regoポリシーファイルを読み込み
	policy, err := os.ReadFile("issueinjection.rego")
	if err != nil {
		fmt.Println("Error reading policy file:", err)
		return
	}

	// Regoクエリを準備
	r := rego.New(
		rego.Query("data.core.deny"),
		rego.Module("issueinjection.rego", string(policy)),
	)

	query, err := r.PrepareForEval(context.Background())
	if err != nil {
		fmt.Println("Error preparing query for evaluation:", err)
		return
	}

	seenMessages := make(map[string]bool)

	// 各ファイルに対してRegoクエリを実行
	for _, file := range files {
		results, err := query.Eval(context.Background(), rego.EvalInput(file))
		if err != nil {
			fmt.Println("Error evaluating query:", err)
			return
		}
		// Print evaluation results for debugging
		// fmt.Println("Evaluation results:", results)

		if len(results) > 0 {
			expressions := results[0].Expressions
			if len(expressions) > 0 {
				if exprValue, ok := expressions[0].Value.(bool); ok && exprValue {
					fmt.Println("Direct usage of ${{ ... }} in run script detected at jobs steps. Use env to set variables.[security-hardening]")
				}
			}
		}
		ProcessResults(results, seenMessages)
	}
}
