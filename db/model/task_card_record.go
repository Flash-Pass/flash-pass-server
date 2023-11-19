package model

type TaskCardRecord struct {
	Base
	TaskId    int64  `json:"task_id"`
	UserId    int64  `json:"user_id"`
	CardId    int64  `json:"card_id"`
	Card      Card   `gorm:"ForeignKey:CardId;AssociationForeignKey:Id" json:"card"`
	Date      string `json:"date"`
	Timestamp int64  `json:"timestamp"`
	IsViewed  bool   `json:"is_viewed"`
}
