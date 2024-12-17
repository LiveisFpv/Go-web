package domain

import "time"

type Group struct {
	Id_group                uint64
	Name_group              string
	Studies_direction_group string
	Studies_profile_group   string
	Start_date_group        time.Time
	Studies_period_group    uint8
}
