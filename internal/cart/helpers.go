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
	itemMap := make(map[string]CartItem)

	// existing items
	for _, item := range existing {
		itemMap[item.MenuItemID] = item
	}

	// incoming items
	for _, item := range incoming {
		if ex, ok := itemMap[item.MenuItemID]; ok {
			ex.Quantity = item.Quantity
			itemMap[item.MenuItemID] = ex
		} else {
			itemMap[item.MenuItemID] = item
		}
	}

	// convert back to slice
	result := make([]CartItem, 0, len(itemMap))
	for _, v := range itemMap {
		result = append(result, v)
	}

	return result
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
