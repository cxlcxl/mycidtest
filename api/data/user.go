package data

type LoginData struct {
	Email          string `json:"email" binding:"required"`
	Password       string `json:"password" binding:"required"`
	ProductVersion int    `json:"product_version" binding:"required"`
}
