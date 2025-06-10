package internal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kalifun/pmindt/entity"
)

func jsonToMarkdownHelper(node entity.DatumChild, indentLevel int, builder *strings.Builder) {
	if len(node.Data.Topic.Children) == 0 {
		return
	}

	// 添加当前节点
	indent := strings.Repeat("\t", indentLevel)
	builder.WriteString(indent + fmt.Sprintf("- %s", node.Data.Topic.Children[0].Text) + "\n")

	// 递归处理子节点
	for _, child := range node.Children {
		jsonToMarkdownHelper(child, indentLevel+1, builder)
	}
}

func PmindToMarkdown(jsonData []byte) (string, error) {
	var data entity.Pcmind
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", err
	}

	if len(data.Data) == 0 {
		return "", fmt.Errorf("no root node found")
	}

	var builder strings.Builder

	// 处理根节点
	root := data.Data[0]
	if len(root.Data.Topic.Children) > 0 {
		builder.WriteString(fmt.Sprintf("- %s", root.Data.Topic.Children[0].Text) + "\n")
	}

	// 递归处理子节点
	for _, child := range root.Children {
		jsonToMarkdownHelper(child, 1, &builder)
	}

	return builder.String(), nil
}
