package internal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kalifun/pmindt/entity"
	"github.com/kalifun/pmindt/utils"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

const (
	InsertHead = "<meta charset='utf-8'><html><head></head><body><plait>"
	InsertEnd  = "</plait></body></html>"
)

func MdToPMind(markdownStr string) (string, error) {
	md := goldmark.New()
	reader := text.NewReader([]byte(markdownStr))
	doc := md.Parser().Parse(reader)

	var root *entity.DatumChild
	stack := []*entity.DatumChild{}
	currentLevel := 0

	ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering || n.Kind() != ast.KindParagraph {
			return ast.WalkContinue, nil
		}

		lines := n.Lines()
		for i := 0; i < lines.Len(); i++ {
			line := lines.At(i)
			lineContent := string(line.Value(reader.Source()))

			// 计算缩进级别（同时处理\t和空格）
			indent := 0
			for _, r := range lineContent {
				if r == '\t' {
					indent++
				} else if r == ' ' {
					// 如果有空格缩进，4个空格=1个\t
					// 这里为了兼容性，但建议统一使用\t
				} else {
					break
				}
			}

			content := strings.TrimSpace(lineContent)
			if content == "" {
				continue
			}

			node := entity.DatumChild{
				ID: utils.GenShortId(),
				Data: entity.Data{
					Topic: entity.Topic{
						Children: []entity.TopicChild{
							{Text: content},
						},
					},
				},
				Children: []entity.DatumChild{},
				Width:    int64(len(content)*8 + 20),
				Height:   20,
			}
			if root == nil {
				root = &node
				stack = append(stack, root)
				currentLevel = indent
				continue
			}

			// 调整栈以匹配当前级别
			for indent <= currentLevel && len(stack) > 1 {
				stack = stack[:len(stack)-1]
				currentLevel--
			}

			// 添加到父节点的子节点
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, node)
				parent.Data = entity.Data{Topic: entity.Topic{Children: []entity.TopicChild{}}}
				stack = append(stack, &parent.Children[len(parent.Children)-1])
				currentLevel = indent
			}
		}
		return ast.WalkContinue, nil
	})

	if root == nil {
		return "", fmt.Errorf("no valid content found in markdown")
	}

	result := entity.Pcmind{
		Type: "elements",
		Data: []entity.Datum{
			{
				ID:       "root",
				Data:     root.Data,
				Children: root.Children,
				Width:    100,
				Height:   25,
				Layout:   "right",
				IsRoot:   true,
				Type:     "mindmap",
				Points:   [][]int64{{0, 12}},
			},
		},
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s%s", InsertHead, string(jsonData), InsertEnd), nil
}
