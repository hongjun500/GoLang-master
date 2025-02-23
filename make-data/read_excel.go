package main

import (
	"github.com/xuri/excelize/v2"
	"log"
)

func readStandFieldsFromExcel() (*excelize.File, error) {
	openFile, err := excelize.OpenFile(`./resources/院内单病种采集数据接口字段整理-V1.1.xlsx`)
	if err != nil {
		log.Fatal("read by open file err", err)
		return nil, err
	}
	return openFile, nil
}
