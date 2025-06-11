package internal

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kalifun/pmindt/entity"
	"github.com/kalifun/pmindt/utils"
)

const (
	InsertHead = "<meta charset='utf-8'><html><head></head><body><plait>"
	InsertEnd  = "</plait></body></html>"
)

func MdToPMind(markdownStr []string) (string, error) {
	rootDatum, _ := parseLines(markdownStr, 0)
	pcmind := entity.Pcmind{
		Type: "elements",
		Data: []entity.Datum{*rootDatum},
	}
	jsonData, err := json.Marshal(pcmind)
	if err != nil {
		return "", err
	}

	data := fmt.Sprintf("%s%s%s", InsertHead, string(jsonData), InsertEnd)
	return data, nil
}

func parseLines(lines []string, level int) (*entity.Datum, []string) {
	if len(lines) == 0 {
		return nil, lines
	}

	name := strings.TrimSpace(lines[0])
	datum := &entity.Datum{
		ID:       utils.GenShortId(),
		Data:     entity.Data{Topic: entity.Topic{Children: []entity.TopicChild{{Text: name}}}},
		Children: make([]entity.DatumChild, 0),
		Width:    int64(len(name) * 10),
		Height:   25,
		Layout:   "right",
		IsRoot:   true,
		Type:     "mindmap",
		Points:   [][]int64{{0, 12}},
	}
	lines = lines[1:]

	for len(lines) > 0 && getIndentationLevel(lines[0]) > level {
		child, remainingLines := parseLines(lines, getIndentationLevel(lines[0]))
		width := 80
		cl := child.Data.Topic.Children
		if len(cl) != 0 {
			width = len(cl[0].Text) * 8
		}
		datum.Children = append(datum.Children, entity.DatumChild{
			ID:       utils.GenShortId(),
			Data:     child.Data,
			Children: child.Children,
			Width:    int64(width),
			Height:   20,
		})
		lines = remainingLines
	}

	return datum, lines
}

func getIndentationLevel(line string) int {
	return len(line) - len(strings.TrimLeft(line, " "))
}
