package model

import "pretest-privyid/modules/v1/category/model"

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type Product struct {
	ID          int              `json:"id,omitempty"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Categories  []model.Category `json:"categories"`
	Image       []Image          `json:"image,omitempty"`
}

type ProductCategory struct {
	ProductId  int
	CategoryId int
}

type ProductImage struct {
	ProductId int
	ImageId   int
}

type Image struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	File string `json:"file"`
}