package main

import (
	color "github.com/gookit/color"
	selection "github.com/subeen-playground/tools/src/utils/selection"
)

func main() {
	// fmt.Println("hello")
	selection := selection.NewSelection([]string{"item1", "item2", "item3"})

	for _, choice := range selection.Items {
		color.Greenln(choice)
	}
}
