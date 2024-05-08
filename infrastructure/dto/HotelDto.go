package dto

type HotelDto struct {
	Name string `json:"name" validate:"required,min=5,max=20"`
	//Email     string `json:"email" validate:"required,email"`
	//Age       int    `json:"age" validate:"gte=18,lte=120"`
}
