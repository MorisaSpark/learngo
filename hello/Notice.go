package main

import "time"

type Notice struct {
	Id            int32     `json:"id"`
	CreateAt      time.Time `json:"create_at"`
	UpdateAt      time.Time `json:"update_at"`
	DeleteAt      time.Time `json:"delete_at"`
	ChannelId     string    `json:"channel_id"`
	Name          string    `json:"name"`
	PrevAt        time.Time `json:"prev_at"`
	NextAt        time.Time `json:"next_at"`
	Frequency     string    `json:"frequency"`
	ExecuteCounts int32     `json:"execute_counts"`
	Status        string    `json:"status"`
	Remark        string    `json:"remark"`
	Mode          string    `json:"mode"`
	UserIds       string    `json:"user_ids"`
}
