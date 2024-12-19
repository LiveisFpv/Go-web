package domain

import "backend/internal/mytype"

type Group struct {
	Name_group              string
	Studies_direction_group string
	Studies_profile_group   string
	Start_date_group        mytype.JsonDate
	Studies_period_group    uint8
}
