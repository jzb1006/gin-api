package request

type EventRequest struct {
	Id     int    `form:"id" json:"id"`
	Search string `form:"search" json:"search"`
	HostId int    `form:"host_id" json:"host_id"`
	TypeId int    `form:"type_id" json:"type_id"`
	Time   string `json:"time" form:"time"`
}
