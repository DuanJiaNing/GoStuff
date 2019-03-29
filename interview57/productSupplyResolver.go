package interview57

import "./model"

var combineList = make([]model.TimeInterval, 10)
var COMBINE_PRESENT model.TimeInterval

func Resolve(pro model.Product, raws *[]model.RawMaterial) *[]model.SupplyTimeIntervalAndQTY {

}

func findIntersection(intervals []model.TimeInterval) model.TimeInterval {

	// sort
	descSortList := []model.TimeInterval{intervals[0]}
	for j := 1; j < len(intervals); j++ {
		interval := intervals[j]

		d1 := calcTimeDurationInDays(interval)
		curIndex := len(descSortList)
		descSortList[curIndex] = interval
		for i := len(descSortList) - 1; i >= 0; i-- {
			curCheck := descSortList[i]
			if d1 > calcTimeDurationInDays(curCheck) {
				descSortList[i], descSortList[curIndex] = descSortList[curIndex], descSortList[i]
				curIndex = i
			}
		}
	}

	firstInterval := descSortList[0]
	from, to := firstInterval.StartTime, firstInterval.EndTime
	for j := 1; j < len(descSortList); j++ {
		interval := descSortList[j]
		startTime, endTime := interval.StartTime, interval.EndTime
		if startTime.After(to) || endTime.Before(from) {
			return model.TimeInterval{}
		}

		if startTime.After(from) {
			from = startTime
		}

		if endTime.Before(to) {
			to = endTime
		}
	}

	return model.TimeInterval{StartTime: from, EndTime: to}
}

func combineTimeInterval(tempList []model.TimeInterval, intervals []model.TimeInterval, n int) {
	if n == 0 {
		for _, v := range tempList {
			combineList = append(combineList, v)
			return
		}
	}

	for _, interval := range intervals {
		notContain := true
		for _, v := range tempList {
			if v == interval {
				notContain = false
			}
		}
		if notContain {
			tempList[len(tempList)-n] = interval
			combineTimeInterval(tempList, intervals, n-1)
			tempList[len(tempList)-n] = COMBINE_PRESENT
		}
	}

}
