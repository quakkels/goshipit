package images

func (i *ImageOptions) GetCategories() map[string]int {
	categories := make(map[string]int)

	for _, item := range i.imageOptions {
		hasCategory := false
		for cat, _ := range categories {
			if cat == item.Category {
				hasCategory = true
			}
		}

		if !hasCategory {
			categories[item.Category] = 0
		}

		categories[item.Category] = categories[item.Category] + 1
	}

	return categories
}
