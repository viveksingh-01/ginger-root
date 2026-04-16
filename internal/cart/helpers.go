package cart

import "github.com/viveksingh-01/ginger-root/internal/menu"

func FindMenuItem(items []menu.MenuItem, id string) *menu.MenuItem {
	for _, item := range items {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func MergeItems(existing, incoming []CartItem) []CartItem {
	itemMap := make(map[string]int)

	for _, item := range existing {
		itemMap[item.MenuItemID] += item.Quantity
	}

	for _, item := range incoming {
		itemMap[item.MenuItemID] += item.Quantity
	}

	var result []CartItem
	for id, qty := range itemMap {
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
