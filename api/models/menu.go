package models

import (
	"errors"

	"gorm.io/gorm"
)

var (
	EntityNotFound = errors.New("Entity not found")
)

type Menu struct {
	BaseModel
	Name      string `gorm:"not null;varchar(255)"`
	MenuItems []MenuItem
}

type MenuService struct {
	db *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{
		db: db,
	}
}

func (ms *MenuService) GetMenuById(id uint) (*Menu, error) {
	var menu Menu
	err := ms.db.First(&menu, id).Error

	switch err {
	case nil:
		return &menu, nil
	case gorm.ErrRecordNotFound:
		return nil, EntityNotFound
	default:
		return nil, err
	}
}

func (ms *MenuService) CreateMenu(menu *Menu) error {
	return ms.db.Create(menu).Error
}
