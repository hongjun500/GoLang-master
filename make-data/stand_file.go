package main

import (
	"strings"
)

type StandField struct {
	Name        string        `bson:"name"`
	Code        string        `bson:"code"`
	Type        string        `bson:"type"`
	ValueRanges []*ValueRange `bson:"valueRange"`
	Description string        `bson:"description"`
	Enable      bool          `bson:"enable"`
	View        string        `bson:"view"`
}

type ValueRange struct {
	Name        string
	Code        string
	DefaultFlag bool
}

var (
	ENUM       = "ENUM"
	MULTI_ENUM = "MULTI_ENUM"
	TEXT       = "TEXT"
	NUMBER     = "NUMBER"
	DATE       = "DATE"
)

func Build(str string) ([]*ValueRange, string) {
	lineCount := strings.Count(str, "\n") + 1
	if lineCount == 1 {
		return nil, str
	}
	lines := make([]string, 0, lineCount)

	for remaining := str; ; {
		line, rest, found := strings.Cut(remaining, "\n")
		lines = append(lines, line)
		if !found {
			break
		}
		remaining = rest
	}
	valueRanges := make([]*ValueRange, len(lines))

	for i, line := range lines {
		split1 := strings.Split(line, "：")
		if len(split1) > 1 {
			valueRanges[i] = &ValueRange{
				Code: strings.TrimSpace(split1[0]),
				Name: strings.TrimSpace(split1[1]),
			}
		} else if split2 := strings.Split(line, "-"); len(split2) > 1 {
			valueRanges[i] = &ValueRange{
				Code: strings.TrimSpace(split2[0]),
				Name: strings.TrimSpace(split2[1]),
			}
		}
	}
	return valueRanges, ""
}

func buildFieldType(field *StandField) {
	// 首先是值域的枚举类型
	if field.ValueRanges != nil {
		field.Type = ENUM
		return
	}
	containDate := strings.Contains(field.Name, "日期") || strings.Contains(field.Name, "时间")
	if containDate {
		field.Type = DATE
		return
	}
	containNumber := strings.Contains(field.Name, "分值") ||
		strings.Contains(field.Name, "数值") ||
		strings.Contains(field.Name, "费")
	if containNumber {
		field.Type = NUMBER
		return
	}
	field.Type = TEXT

}
