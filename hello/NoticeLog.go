package main

import "time"

type NoticeLog struct {
	Id          int32     `json:"id"`
	CreateAt    time.Time `json:"create_at"`
	UserId      string    `json:"user_id"`
	ActionId    int32     `json:"action_id"`
	ActionTable string    `json:"action_table"`
	Type        string    `json:"type"`
}
