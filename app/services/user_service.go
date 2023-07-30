package services

import (
	"test-golang/app/entities"
	"test-golang/app/repositories"
)

type UserService interface {
	AllUserFamily() ([]entities.DataUserFamily, error)
}

type userService struct {
	userRepo repositories.UserRepo
}

func NewUserService(userRepo repositories.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (service *userService) AllUserFamily() ([]entities.DataUserFamily, error) {
	var dataUserFamily []entities.DataUserFamily
	user, err := service.userRepo.GetCustomer()
	if err != nil {
		return nil, err;
	}

	for _, item := range user {
		var dobTemp string
		
		families, err := service.userRepo.GetFamilyByIDCustomer(item.CustomerID)
		if err != nil {
			return nil, err
		}

		dobTemp = item.CustomerDOB.Format("2006-01-02")
		dataUser := entities.DataUserFamily{
			CustomerID: item.CustomerID,
			NationalityID: item.NationalityID,
			CustomerName: item.CustomerName,
			CustomerDOB: dobTemp,
			CustomerPhoneNum: item.CustomerPhoneNum,
			CustomerEmail: item.CustomerEmail,
			Families: families,
		}

		dataUserFamily = append(dataUserFamily, dataUser)
	}

	return dataUserFamily, nil
}