package category

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Categories []*Category

// initialize data
var CategoriesData = Categories{
	{
		ID:          1,
		Name:        "Electronics",
		Description: "Electronics category",
	},
}

func (c *Categories) GetByID(id int) *Category {
	for _, category := range CategoriesData {
		if category.ID == id {
			return category
		}
	}
	return nil
}
