package model

import "github.com/samber/oops"

func GetModels() []interface{} {
	return []interface{}{
		&Category{},
		&Word{},
		&List{},
	}
}

var errBuilder = oops.In("model").
	Tags("conversion")
