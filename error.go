package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

// コンソール出力時における色付けのための定数
var (
	BoldStyle   = color.New(color.Bold)
	GreenStyle  = color.New(color.FgGreen)
	YellowStyle = color.New(color.FgYellow)
	GrayStyle   = color.New(color.FgHiBlack)
	OrangeStyle = color.New(color.FgHiYellow)
	RedStyle    = color.New(color.FgRed)
)

type Position struct {
	// Line は位置の行番号です。この値は1から始まります。
	Line int
	// Col は位置の列番号です。この値は1から始まります。
	Col int
}

type PolicyError struct {
	Description string
	FilePath    string
	Line    int
	Column  int
	Type string
}

func (e *PolicyError) Error() string {
	return fmt.Sprintf("%s:%d:%d: %s [%s]", e.FilePath, e.Line, e.Column, e.Description, e.Type)
}

func NewError(position *Position, errorType string, message string) *PolicyError {
	return &PolicyError{
		Description: message,
		Line:  position.Line,
		Column:   position.Col,
		Type:        errorType,
	}
}

func FormattedError(position *Position, errorType string, format string, args ...interface{}) *PolicyError {
	return &PolicyError{
		Description: fmt.Sprintf(format, args...),
		Line:  position.Line,
		Column:   position.Col,
		Type:        errorType,
	}
}

// extractLineContentはソースコードの中からエラーが発生した行の内容を抽出する
func (e *PolicyError) extractLineContent(sourceContent []byte) (string, bool) {
	s := bufio.NewScanner(bytes.NewReader(sourceContent))
	lineNumber := 0
	for s.Scan() {
		lineNumber++
		if lineNumber == e.Line {
			return s.Text(), true
		}
	}
	return "", false
}

// DisplayErrorはエラーを見やすい形で出力する
// sourceがnilな場合はスニペットは表示しない
func (e *PolicyError) DisplayError(output io.Writer, sourceContent []byte) {
	printColored(output, GreenStyle, e.FilePath)
	printColored(output, GrayStyle, ":")
	fmt.Fprint(output, e.Line)
	printColored(output, GrayStyle, ":")
	fmt.Fprint(output, e.Column)
	printColored(output, GrayStyle, ": ")
	printColored(output, OrangeStyle, e.Description)
	printColored(output, RedStyle, fmt.Sprintf(" [%s]\n", e.Type))

	if len(sourceContent) == 0 || e.Line == 0 {
		return
	}

	lineContent, found := e.extractLineContent(sourceContent)
	if !found || len(lineContent) < e.Column-1 {
		return
	}

	lineHeader := fmt.Sprintf("%d 👈|", e.Line)
	padding := strings.Repeat(" ", len(lineHeader)-2)
	printColored(output, GrayStyle, fmt.Sprintf("%s %s", padding, lineHeader))
	fmt.Fprintln(output, lineContent)
	printColored(output, GrayStyle, fmt.Sprintf("%s %s\n", padding, strings.Repeat(" ", e.Column-1)))
}

// helper function to print with color
func printColored(output io.Writer, colorizer *color.Color, content string) {
	colorizer.Fprint(output, content)
}
