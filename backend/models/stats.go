package models

type Dashboard struct {
	TotalUsers         int     `json:"total_users"`
	TotalUpcycled      int     `json:"total_upcycled"`
	TotalProjects      int     `json:"total_projects"`
	TotalContainers    int     `json:"total_containers"`
	WasteAvoided       float64 `json:"waste_avoided"`
	TotalAnnouncements int     `json:"total_announcements"`
}
