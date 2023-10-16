package services

import (
	"database/sql"
	user_domain "server/domain"
)

type UserPort interface {
	CreateUser(*user_domain.Users) (*user_domain.Users, error)
	GetUser(*user_domain.Users) (*user_domain.Users, error)
	UpdatePlan(*user_domain.Users, int) (error)
	BanUser(*user_domain.Users) (error)
	Login(*user_domain.Users) (bool, error)
}

type Users_services_implmentation struct {
	DB *sql.DB
}


func CreateUsersServicesImplementation(DB *sql.DB) UserPort {
	return &Users_services_implmentation{DB: DB}
}
