package models

import (
	"gorm.io/gorm"
)

type Menu struct {
	BaseModel
	Name      string     `gorm:"not null;varchar(255)" json:"name"`
	MenuItems []MenuItem `json:"menuItems"`
}

type MenuService struct {
	db *gorm.DB
}

func NewMenuService(db *gorm.DB) *MenuService {
	return &MenuService{
		db: db,
	}
}

func (ms *MenuService) GetMenuById(id uint64) (*Menu, error) {
	var menu Menu
	if err := ms.db.First(&menu, id).Error; err != nil {
		return nil, err
	} else {
		return &menu, nil
	}
}

func (ms *MenuService) CreateMenu(menu *Menu) error {
	return ms.db.Create(menu).Error
}

func (ms *MenuService) GetAll() (*[]Menu, error) {
	var menus []Menu

	if err := ms.db.Find(&menus).Error; err != nil {
		return nil, err
	} else {
		return &menus, nil
	}
}

func (ms *MenuService) DeleteMenu(menuId uint) error {
	return ms.db.Delete(&Menu{}, menuId).Error
}
