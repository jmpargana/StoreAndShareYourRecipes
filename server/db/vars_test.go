package db

import "server/models"

var (
	selectRecipeByIDQueryTest = "SELECT id, private, title, ingridients, time, method FROM recipes WHERE id = \\?"
	recipeTableTop            = []string{"id", "private", "title", "ingridients", "time", "method"}
	r                         = models.Recipe{
		ID:          1,
		Private:     1,
		Title:       "curry",
		Ingridients: "rice",
		Time:        30,
		Method:      "cook",
	}
)
