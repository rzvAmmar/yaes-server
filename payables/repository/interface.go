package repository

import "algogrit.com/yaes-server/entities"

type PayableRepository interface {
	RetrieveBy(entities.User) ([]*entities.Payable, error)
	FindBy(uint) (*entities.Payable, error)
	Update(*entities.Payable) error
}
