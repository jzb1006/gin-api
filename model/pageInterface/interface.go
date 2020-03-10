package pageInterface

type PageInfo struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

type Paging interface {
	GetInfoList(PageInfo) (list interface{}, total int, currentPage int, totalPage int, err error)
}

