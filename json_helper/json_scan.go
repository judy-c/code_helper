package json_helper

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type JsonScanner struct {
	JsonString 		string
	ResultModelArr	[]StructTypeModel
}

func (scan *JsonScanner) ScanJson() error {
	JsonModelMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(scan.JsonString), &JsonModelMap)
	if err != nil {
		fmt.Println("illegal json string")
		return err
	}
	resultModelArr := make([]StructTypeModel, 0)
	rootModel, err := getStructModelFromInterface("RootModel", JsonModelMap, &resultModelArr)
	if err == nil {
		resultModelArr = append(resultModelArr, rootModel)
	}
	fmt.Printf("get result model:%v\n", resultModelArr)
	scan.ResultModelArr = resultModelArr
	return nil
}

func getStructModelFromInterface(ObjectName string, source map[string]interface{}, resultModelArr *[]StructTypeModel) (StructTypeModel, error) {
	result := StructTypeModel{
		TypeName:    ObjectName,
		StructItems: make([]StructItemModel, 0),
	}
	for modelName, val := range source {
		typeName := reflect.TypeOf(val)
		switch typeName.String() {
		case "map[string]interface {}":
			// TODO:
			objectName := modelName + "Model"
			model, err := getStructModelFromInterface(objectName, val.(map[string]interface{}), resultModelArr)
			if err != nil {
				return result, err
			}
			*resultModelArr = append(*resultModelArr, model)
			result.StructItems = append(result.StructItems, StructItemModel{
				ItemType: model.TypeName,
				VarName:  modelName,
			})
			fmt.Println("get an object")
		case "[]interface {}":
			// TODO:
			arrModel := val.([]interface{})
			arrItemType := reflect.TypeOf(arrModel[0])
			if arrItemType.String() == "map[string]interface {}" {
				objectName := modelName + "ItemModel"
				model, err := getStructModelFromInterface(objectName, arrModel[0].(map[string]interface{}), resultModelArr)
				if err != nil {
					return result, err
				}
				*resultModelArr = append(*resultModelArr, model)
				result.StructItems = append(result.StructItems, StructItemModel{
					ItemType: "[]" + model.TypeName,
					VarName:  modelName,
				})
			} else {
				result.StructItems = append(result.StructItems, StructItemModel{
					ItemType: "[]" + arrItemType.String(),
					VarName:  modelName,
				})
			}
			fmt.Println("get an array")
		default:
			result.StructItems = append(result.StructItems, StructItemModel{
				ItemType: typeName.String(),
				VarName:  modelName,
			})
		}
	}
	return result, nil
}