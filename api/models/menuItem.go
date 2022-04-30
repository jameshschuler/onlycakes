package models

import "gorm.io/gorm"

type MenuItem struct {
	BaseModel
	Name          string          `gorm:"not null;varchar(255)" json:"name"`
	Price         float32         `gorm:"not null" json:"price"`
	Active        bool            `gorm:"not null;default:true" json:"active"`
	MenuId        uint            `gorm:"not null" json:"-"`
	Configurable  bool            `gorm:"not null;default:false" json:"configurable"`
	IsParentItem  bool            `gorm:"not null;default:true" json:"isParentItem"`
	MenuItemSteps []*MenuItemStep `gorm:"many2many:menuitem_menuitemsteps;" json:"menuItemSteps,omitempty"`
}

type MenuItemStep struct {
	BaseModel
	Name                string      `gorm:"not null;varchar(255)"`
	MenuItems           []*MenuItem `gorm:"many2many:menuitem_menuitemsteps;"`
	MenuItemStepOptions []MenuItemStepOption
}

type MenuItemStepOption struct {
	BaseModel
	MenuItemId     uint `gorm:"not null"`
	MenuItemStepId uint `gorm:"not null"`
	MenuItem       MenuItem
	MenuItemStep   MenuItemStep
}

type MenuItemService struct {
	db *gorm.DB
}

func NewMenuItemService(db *gorm.DB) *MenuItemService {
	return &MenuItemService{
		db: db,
	}
}

func (menuItemService *MenuItemService) GetAll(menuId uint, includeInactive bool) (*[]MenuItem, error) {
	var menuItems []MenuItem

	query := "menu_id = ?"

	if !includeInactive {
		query += "AND active = true"
	}

	if err := menuItemService.db.Where(query, menuId).Find(&menuItems).Error; err != nil {
		return nil, err
	} else {
		return &menuItems, nil
	}
}
