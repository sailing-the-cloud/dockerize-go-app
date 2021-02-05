package model

/**
 * Created by Manggala Pramuditya Wiryawan on 08/11/19 Nov, 2019
 * email : manggala.wiryawan@gmail.com
 */

type Category struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name"`
	Enable string `json:"-"`
}

type Categories []Category
