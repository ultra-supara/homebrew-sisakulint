package root

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

// Position はファイル内の位置を表します。
type Position struct {
	// Line は位置の行番号です。この値は1から始まります。
	Line int
	// Col は位置の列番号です。この値は1から始まります。
	Col int
}

// コンソール出力時における色付けのための定数
var (
	BoldStyle   = color.New(color.Bold)
	GreenStyle  = color.New(color.FgGreen)
	YellowStyle = color.New(color.FgYellow)
	GrayStyle   = color.New(color.FgHiBlack)
	OrangeStyle = color.New(color.FgHiYellow)
	RedStyle    = color.New(color.FgRed)
)

// PolicyErrorはsisakupcにおけるpolicy errorの詳細を表す構造体
type PolicyError struct {
	//PolicyErrorの種類
	Description string
	//PolicyErrorが発生したファイルのパス
	FilePath string
	//PolicyErrorが発生した行番号
	LineNumber int
	//PolicyErrorが発生した列番号
	ColNumber int
	//PolicyErrorが発生した行の内容
	Type string
}

func (e *PolicyError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s [%s]", e.FilePath, e.LineNumber, e.ColNumber, e.Description, e.Type)
}

func (e *PolicyError) String() string {
	return e.Error()
}

func NewError(position *Position, errorType string, message string) *PolicyError {
	return &PolicyError{
		Description: message,
		LineNumber:  position.Line,
		ColNumber:   position.Col,
		Type:        errorType,
	}
}

func FormattedError(position *Position, errorType string, format string, args ...interface{}) *PolicyError {
	return &PolicyError{
		Description: fmt.Sprintf(format, args...),
		LineNumber:  position.Line,
		ColNumber:   position.Col,
		Type:        errorType,
	}
}

// ExtractTemplateFieldsはPolicyErrorからテンプレートの生成に必要なフィールドを抽出する
func (e *PolicyError) ExtractTemplateFields(sourceContent []byte) *TemplateFields {
	codeSnippet := ""

	if len(sourceContent) > 0 && e.LineNumber > 0 {
		if lineContent, found := e.extractLineContent(sourceContent); found {
			codeSnippet = lineContent
		}
	}
	return &TemplateFields{
		Message:  e.Description,
		Filepath: e.FilePath,
		Line:     e.LineNumber,
		Column:   e.ColNumber,
		Type:     e.Type,
		Snippet:  codeSnippet,
	}
}

type RuleTemplateField struct {
	Name        string
	Description string
}

type ByRuleTemplateField []*RuleTemplateField

func (a ByRuleTemplateField) Len() int {
	return len(a)
}

func (a ByRuleTemplateField) Less(i, j int) bool {
	return strings.Compare(a[i].Name, a[j].Name) < 0
}

func (a ByRuleTemplateField) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// DisplayErrorはエラーを見やすい形で出力する
// sourceがnilな場合はスニペットは表示しない
func (e *PolicyError) DisplayError(output io.Writer, sourceContent []byte) {
	printColored(output, GreenStyle, e.FilePath)
	printColored(output, GrayStyle, ":")
	fmt.Fprint(output, e.LineNumber)
	printColored(output, GrayStyle, ":")
	fmt.Fprint(output, e.ColNumber)
	printColored(output, GrayStyle, ": ")
	printColored(output, OrangeStyle, e.Description)
	printColored(output, RedStyle, fmt.Sprintf(" [%s]\n", e.Type))

	if len(sourceContent) == 0 || e.LineNumber == 0 {
		return
	}

	lineContent, found := e.extractLineContent(sourceContent)
	if !found || len(lineContent) < e.ColNumber-1 {
		return
	}

	lineHeader := fmt.Sprintf("%d 👈|", e.LineNumber)
	padding := strings.Repeat(" ", len(lineHeader)-2)
	printColored(output, GrayStyle, fmt.Sprintf("%s %s", padding, lineHeader))
	fmt.Fprintln(output, lineContent)
	printColored(output, GrayStyle, fmt.Sprintf("%s %s\n", padding, strings.Repeat(" ", e.ColNumber-1)))
}

// helper function to print with color
func printColored(output io.Writer, colorizer *color.Color, content string) {
	colorizer.Fprint(output, content)
}

// extractLineContentはソースコードの中からエラーが発生した行の内容を抽出する
func (e *PolicyError) extractLineContent(sourceContent []byte) (string, bool) {
	s := bufio.NewScanner(bytes.NewReader(sourceContent))
	lineNumber := 0
	for s.Scan() {
		lineNumber++
		if lineNumber == e.LineNumber {
			return s.Text(), true
		}
	}
	return "", false
}

type ByRuleErrorPosition []*PolicyError

func (a ByRuleErrorPosition) Len() int {
	return len(a)
}

func (a ByRuleErrorPosition) Less(i, j int) bool {
	if comparisonResult := strings.Compare(a[i].FilePath, a[j].FilePath); comparisonResult != 0 {
		return comparisonResult < 0
	}
	if a[i].LineNumber == a[j].LineNumber {
		return a[i].ColNumber < a[j].ColNumber
	}
	return a[i].LineNumber < a[j].LineNumber
}

func (a ByRuleErrorPosition) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// TemplateFieldsはエラーメッセージをフォーマットするためのフィールド保持
type TemplateFields struct {
	// Message はエラーメッセージの本文
	Message string `json:"message"`
	// Filepath は正規の相対ファイルパスです。入力が標準入力から読み取られた場合、このフィールドは空に
	// JSONにエンコードする際、ファイルパスが空の場合（このフィールドは省略される可能性あり）
	Filepath string `json:"filepath,omitempty"`
	// Line はエラー位置の行番号
	Line int `json:"line"`
	// Column はエラー位置の列番号
	Column int `json:"column"`
	// Type はエラーが属しているルールの名前
	Type string `json:"type"`
	// Snippet はエラーが発生した位置を示すコードスニペットおよびインジケーター
	// JSONにエンコードする際、スニペットが空の場合、(このフィールドは省略される可能性あり)
	Snippet string `json:"snippet,omitempty"`
}
