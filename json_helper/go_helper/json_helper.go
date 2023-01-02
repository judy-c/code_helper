package go_helper

import (
	"fmt"
	"github.com/judy-c/code_helper/json_helper"
)

func dealValueName(valueJsonName string) string {
	return valueJsonName
}

func printStructDefine(ResultModelArr []json_helper.StructTypeModel) string {
	result := ""
	structStrArr := make([]string, 0)

	for _, resultModel := range ResultModelArr {
		oneStructStrLines := make([]string, 0)
		firstLine := fmt.Sprintf("type %s struct {\n", resultModel.TypeName)
		oneStructStrLines = append(oneStructStrLines, firstLine)
		for _, itemInfo := range resultModel.StructItems {
			jsonStr := fmt.Sprintf("\"json:`%s`\"", itemInfo.VarName)
			valueName := dealValueName(itemInfo.VarName)
			itemStr := fmt.Sprintf("\t%s\t\t%s\t\t%s\n", valueName, itemInfo.ItemType, jsonStr)
			oneStructStrLines = append(oneStructStrLines, itemStr)
		}
		oneStructStrLines = append(oneStructStrLines, "}")
		structStr := ""
		for _, lineStr := range oneStructStrLines {
			structStr = structStr + lineStr
		}
		structStrArr = append(structStrArr, structStr)
	}

	for _, structStr := range structStrArr {
		result = result + structStr + "\n\n"
	}
	return result
}

type Example struct {
	Name string
	Value string
}

func GetGoStructDefine(jsonStr string) (string, error) {
	scanner := json_helper.JsonScanner{
		JsonString:     jsonStr,
		ResultModelArr: nil,
	}
	err := scanner.ScanJson()
	if err != nil {
		return "", err
	}
	res := printStructDefine(scanner.ResultModelArr)
	return res, nil
}