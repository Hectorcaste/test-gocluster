package models

type BranchModel struct {
	BranchName      string  `gorm:"column:nombre"`
	BranchLatitude  float64 `gorm:"column:latitud"`
	BranchLongitude float64 `gorm:"column:longitud"`
}
