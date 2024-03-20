package models

// Items struct
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ItemCode    string `json:"item_code"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	CategoryID  int    `json:"category_id"`
	Status      string `json:"status"`
	BrandID     int    `json:"brand_id"`
	LocationID  int    `json:"location_id"`
}

// Categories struct
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Brands struct
type Brand struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Locations struct
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
