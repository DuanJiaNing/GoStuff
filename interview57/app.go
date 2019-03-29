package interview57

import (
	"./model"
	"fmt"
)

func main() {

	// 原材料及其生产信息
	RAW_EUCALYPTUS_001_ID := "RAW_EUCALYPTUS_001"
	RAW_EUCALYPTUS_001 := model.RawMaterial{RAW_EUCALYPTUS_001_ID, []model.SupplyTimeIntervalAndQTY{
		{model.TimeInterval{getDate(2014, 02, 04, 0, 0, 0), getDate(2014, 11, 31, 0, 0, 0)}, 6000, RAW_EUCALYPTUS_001_ID},
		{model.TimeInterval{getDate(2015, 02, 01, 0, 0, 0), getDate(2038, 01, 19, 0, 0, 0)}, 6000, RAW_EUCALYPTUS_001_ID}}}

	RAW_ROSE_005_ID := "RAW_ROSE_005"
	RAW_ROSE_005 := model.RawMaterial{RAW_ROSE_005_ID, []model.SupplyTimeIntervalAndQTY{
		{model.TimeInterval{getDate(2014, 10, 01, 0, 0, 0), getDate(2014, 10, 31, 0, 0, 0)}, 18, RAW_ROSE_005_ID},
		{model.TimeInterval{getDate(2015, 01, 01, 0, 0, 0), getDate(2015, 01, 31, 0, 0, 0)}, 666, RAW_ROSE_005_ID}}}

	CAPACITY_ID := "CAPACITY"
	CAPACITY := model.RawMaterial{CAPACITY_ID, []model.SupplyTimeIntervalAndQTY{
		{model.TimeInterval{getDate(2014, 10, 01, 0, 0, 0), getDate(2014, 10, 31, 0, 0, 0)}, 18, CAPACITY_ID},
		{model.TimeInterval{getDate(2015, 01, 01, 0, 0, 0), getDate(2015, 01, 31, 0, 0, 0)}, 666, CAPACITY_ID}}}

	p98100201 := model.Product{"98100201", []model.ProductRawMaterials{{RAW_ROSE_005_ID, 14}, {CAPACITY_ID, 1}}}
	p98102601 := model.Product{"98102601", []model.ProductRawMaterials{{RAW_ROSE_005_ID, 12}, {CAPACITY_ID, 1}, {RAW_EUCALYPTUS_001_ID, 4}}}

	raws := []model.RawMaterial{RAW_ROSE_005, RAW_EUCALYPTUS_001, CAPACITY}
	var res1 = Resolve(p98100201, &raws)
	var res2 = Resolve(p98102601, &raws)

	p(*res1)
	p(*res2)

}

func p(arg []model.SupplyTimeIntervalAndQTY) {
	for _, v := range arg {
		fmt.Println(v)
	}
}
