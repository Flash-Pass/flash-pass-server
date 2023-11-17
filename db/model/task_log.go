package model

type TaskLog struct {
	Base
	TaskCardRecordId int64  `json:"task_card_record_id"`
	Type             string `json:"type"`
	LearnStatus      string `json:"learn_status"`
	CreatedBy        int64  `json:"created_by"`
}
