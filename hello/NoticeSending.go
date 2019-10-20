package main

import "time"

type NoticeSending struct {
	Id       int32     `json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	NoticeAt time.Time `json:"notice_at"`
	UserId   int32     `json:"user_id"`
	NoticeId int32     `json:"notice_id"`
	Status   string    `json:"status"`
}
