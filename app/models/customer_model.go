package models

import "time"

func (Customer) TableName() string {
	return "trx.customer"
}

type Customer struct {
	CustomerID          uint      `gorm:"column:cst_id;not null;primaryKey;autoIncrement:true"`
	NationalityID       uint      `gorm:"column:nationality_id;not null;"`
	CustomerName        string      `gorm:"column:cst_name;not null;"`
	CustomerDOB         time.Time `gorm:"column:cst_dob;not null;"`
	CustomerPhoneNumber string      `gorm:"column:cst_phone_num;not null;"`
	CustomerEmail       string      `gorm:"column:cst_email;not null;"`
}