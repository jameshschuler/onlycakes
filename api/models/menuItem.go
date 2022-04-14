package models

type MenuItem struct {
	BaseModel
	Name          string          `gorm:"not null;varchar(255)"`
	Price         float32         `gorm:"not null"`
	Active        bool            `gorm:"not null;default:true"`
	MenuId        uint            `gorm:"not null"`
	Configurable  bool            `gorm:"not null;default:false"`
	IsParentItem  bool            `gorm:"not null;default:true"`
	MenuItemSteps []*MenuItemStep `gorm:"many2many:menuitem_menuitemsteps;"`
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
