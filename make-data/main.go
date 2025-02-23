package main

import (
	"context"
	"fmt"
	_ "fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func main() {
	print("hello make-data")

	properties := MongoProperties{
		Uri:      "mongodb://mongodb:mongodb@localhost:27017/admin", // 本地
		DataBase: "inner-drugs-auto",
	}
	client, err := properties.NewMongoClient()
	defer client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal("connect fail")
	}
	timestamp := time.Now().Format("20060102_150405")
	// collection := client.Database(properties.DateBase).Collection("dc_standard_field_from_excel")
	collection := client.Database(properties.DataBase).Collection("dc_standard_field" + "_" + timestamp)
	// 查全部
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return
	}
	maps := make([]any, 1)
	err = cur.All(context.Background(), &maps)
	if err != nil {
		log.Fatal("err")
	}

	excel, err := readStandFieldsFromExcel()
	if err != nil {
		log.Fatal("open files fail...")
		return
	}
	sheetNameMap := make(map[string]string, 16)
	sourceMap := make(map[string]string, 16)
	standFields := make([]any, 0)
	// 获取所有sheet
	sheets := excel.GetSheetList()
	for sheetIndex, sheet := range sheets {
		rows, err := excel.GetRows(sheet)
		if err != nil {
			log.Fatal("get rows fail...")
		}
		for rowIndex, row := range rows {
			if rowIndex == 0 {
				continue
			}
			if sheetIndex == 0 {
				if sheetIndex == 0 {
					v := row[2]
					if v != "" {
						sheetNameMap[v] = row[1]
						sourceMap[v] = row[3]
					}
				}
			} else {
				// 拆解需要入库的 sheet 数据, 从索引 = 1 的行开始

				// 获取 索引列 = 1 的数据开始
				name := row[1]
				// 针对特殊列内容进行转换
				desc := row[2]
				// 1. 如果是换行符的数据需要转成 对象切片
				valueRange, str := Build(desc)
				if valueRange != nil {
					// 额外填充数组值域
					// todo
					fmt.Println("is valueRange", valueRange)
				} else {
					fmt.Println("not valueRange", str)
				}

				code := row[4]

				standField := &StandField{
					Name:        name,
					Code:        code,
					Type:        "",
					ValueRanges: valueRange,
					Description: str,
					Enable:      true,
				}

				// 定义字段的类型
				buildFieldType(standField)
				standFields = append(standFields, standField)
				// fmt.Println("name", name, "regexs", valueRange, "code", code)
				fmt.Println(standField)
			}
		}
	}

	result, err := collection.InsertMany(context.TODO(), standFields)
	if err != nil {
		log.Fatal("insertmany err ", err)
		return
	}
	for _, insertedID := range result.InsertedIDs {
		fmt.Println("insert successful id = ", insertedID)
	}
}
