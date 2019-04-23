package controllers


func PageCal(totalNum int64,pageNo int32,pageSize int) (bool,int32,int32) {
	isPage := true
	if int(totalNum) < pageSize {
		isPage = false
	}
	pagePrev := pageNo
	pageNext := pageNo + 2
	temp := int(totalNum) / pageSize
	if int(pageNo) == temp {
		pageNext = pageNo + 1
	}
	if pageNo == 0 {
		pagePrev = pageNo + 1
	}
	return isPage,pagePrev,pageNext
}


func PageRangeCal(totalNum int64,pageNo int32,pageSize int) (bool,[]int,int32,int32) {
	isPage := true
	if int(totalNum) < pageSize {
		isPage = false
	}
	temp := int(totalNum) / pageSize
	var t  []int
	for i := 0; i <= temp;i ++ {
		t = append(t, i+1)
	}
	pageRange := t
	pagePrev := pageNo
	pageNext := pageNo + 2
	if int(pageNo) == temp {
		pageNext = pageNo + 1
	}
	if pageNo == 0 {
		pagePrev = pageNo + 1
	}
	return isPage,pageRange,pagePrev,pageNext
}