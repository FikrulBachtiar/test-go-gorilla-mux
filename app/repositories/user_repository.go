package repositories

import (
	"test-golang/app/entities"
	"test-golang/app/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetCustomer() ([]entities.DataCustomer, error)
	GetFamilyByIDCustomer(id int) ([]entities.Family, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) GetCustomer() ([]entities.DataCustomer, error) {
	rows, err := repo.db.Model(&models.Customer{}).Select("cst_id, nationality_id, cst_name, cst_dob, cst_phone_num, cst_email").Rows()
	if err != nil {
		return nil, err;
	}
	defer rows.Close()

	var result []entities.DataCustomer
	for rows.Next() {
		var data entities.DataCustomer
		err := rows.Scan(&data.CustomerID, &data.NationalityID, &data.CustomerName, &data.CustomerDOB, &data.CustomerPhoneNum, &data.CustomerEmail)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}
	
	return result, nil
}

func (repo *userRepo) GetFamilyByIDCustomer(id int) ([]entities.Family, error) {
	rows, err := repo.db.Model(&models.Family{}).Select("fl_relation, fl_name, fl_dob").Where("cst_id = ?", id).Rows()
	if err != nil {
		return nil, err;
	}
	defer rows.Close()

	var result []entities.Family
	for rows.Next() {
		var data entities.Family
		err := rows.Scan(&data.Relation, &data.Name, &data.Dob)
		if err != nil {
			return nil, err
		}

		result = append(result, data)
	}
	
	return result, nil
}