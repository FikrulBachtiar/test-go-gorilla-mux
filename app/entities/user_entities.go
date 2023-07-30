package entities

import "time"

type DataCustomer struct {
	CustomerID     	 int   `json:"cst_id"`
	NationalityID    int   `json:"nationality_id"`
	CustomerName     string   `json:"cst_name"`
	CustomerDOB      time.Time `json:"cst_dob"`
	CustomerPhoneNum string `json:"cst_phone_num"`
	CustomerEmail    string   `json:"cst_email"`
}

type DataUserFamily struct {
	CustomerID     	 int   `json:"cst_id"`
	NationalityID    int   `json:"nationality_id"`
	CustomerName     string   `json:"cst_name"`
	CustomerDOB      string `json:"cst_dob"`
	CustomerPhoneNum string `json:"cst_phone_num"`
	CustomerEmail    string   `json:"cst_email"`
	Families		 []Family
}

type Family struct {
	Relation	string `json:"fl_relation"`
	Name		string `json:"fl_name"`
	Dob			string `json:"fl_dob"`
}