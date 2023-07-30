package models

import "time"

func (Family) TableName() string {
	return "trx.family_list"
}

type Family struct {
	FamilyListID        uint		`gorm:"column:fl_id;not null;primaryKey;autoIncrement:true"`
	FamilyListRelation  string      `gorm:"column:fl_relation;not null;"`
	FamilyListName		string      `gorm:"column:fl_name;not null;"`
	FamilyListDOB       time.Time   `gorm:"column:fl_dob;not null;"`
}