package model

func GetModels() []interface{} {
	return []interface{}{
		&Category{},
		&Word{},
		&List{},
	}
}
