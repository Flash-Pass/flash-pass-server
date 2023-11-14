package model

type Plan struct {
	Base
	Title          string `gorm:"type:varchar(20)" json:"title"`
	Description    string `gorm:"type:varchar(255)" json:"description"`
	CycleSize      int    `json:"cycle_size"`
	LearnPerCycle  int    `json:"learn_per_cycle"`
	ReviewPerCycle int    `json:"review_per_cycle"`
	ReviewPerLearn int    `json:"review_per_learn"`
	GroupSize      int    `json:"group_size"`
	ReviewCycles   int    `json:"review_cycles"`
	LearnStrategy  string `gorm:"type:varchar(20)" json:"learn_strategy"`
	ReviewStrategy string `gorm:"type:varchar(20)" json:"review_strategy"`
	CreatedBy      int64  `gorm:"index" json:"created_by,string"`
}

type Queries interface{}
