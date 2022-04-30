package domain

import "onlycakes/models"

type MenuItemsResponse struct {
	MenuItems []models.MenuItem `json:"menuItems"`
}
