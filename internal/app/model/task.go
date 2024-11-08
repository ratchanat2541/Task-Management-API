package model

import "time"

type Task struct {
	ID          int              `json:"id" gorm:"primaryKey;type:autoIncrement;not null"`
	Title       string           `json:"title" gorm:"type:varchar(255);not null"`
	Description string           `json:"description" gorm:"type:varchar(255);not null"`
	Status      StatusOfTaskEnum `json:"status" gorm:"type:int;not null"`
	CreatedAt   time.Time        `json:"created_at" gorm:"type:timestamptz;not null"`
	UpdatedAt   time.Time        `json:"updated_at" gorm:"type:timestamptz;not null"`
}

type TaskRequest struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      StatusOfTaskEnum `json:"status"`
}

type TaskUpdateRequest struct {
	Title       *string           `json:"title"`
	Description *string           `json:"description"`
	Status      *StatusOfTaskEnum `json:"status"`
}

type TaskUpdateStatusRequest struct {
	Status StatusOfTaskEnum `json:"status"`
}

type StatusOfTaskEnum uint64

var STATUS_OF_TASK_TO_DO StatusOfTaskEnum = 1
var STATUS_OF_TASK_IN_PROGRESS StatusOfTaskEnum = 2
var STATUS_OF_TASK_DONE StatusOfTaskEnum = 3

var STATUS_OF_TASK_TO_DO_NAME = "To Do"
var STATUS_OF_TASK_IN_PROGRESS_NAME = "In Progress"
var STATUS_OF_TASK_DONE_NAME = "Done"

var StatusOfEventCupNameEnums = map[StatusOfTaskEnum]string{
	STATUS_OF_TASK_TO_DO:       STATUS_OF_TASK_TO_DO_NAME,
	STATUS_OF_TASK_IN_PROGRESS: STATUS_OF_TASK_IN_PROGRESS_NAME,
	STATUS_OF_TASK_DONE:        STATUS_OF_TASK_DONE_NAME,
}

var StatusOfEventCupIDEnums = map[int]*StatusOfTaskEnum{
	1: &STATUS_OF_TASK_TO_DO,
	2: &STATUS_OF_TASK_IN_PROGRESS,
	3: &STATUS_OF_TASK_DONE,
}
