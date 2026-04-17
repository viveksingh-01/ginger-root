package cart

import (
	"github.com/google/uuid"
	"github.com/viveksingh-01/ginger-root/internal/menu"
)

func generateGuestID() string {
	return uuid.New().String()
}

func FindMenuItem(items []menu.MenuItem, id string) *menu.MenuItem {
	for _, item := range items {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func MergeItems(existing, incoming []CartItem) []CartItem {
	m := make(map[string]int)

	for _, i := range existing {
		m[i.MenuItemID] += i.Quantity
	}
	for _, i := range incoming {
		m[i.MenuItemID] += i.Quantity
	}

	var result []CartItem
	for id, qty := range m {
		result = append(result, CartItem{
			MenuItemID: id,
			Quantity:   qty,
		})
	}
	return result
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
