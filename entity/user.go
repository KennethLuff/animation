package entity

type User struct {
	Id           int64  `json:"id" form:"id"`
	Tel          string `json:"tel" form:"tel"`
	Password     string `json:"password" form:"password"`
	NickName     string `json:"nick_name" form:"nickname"`
	Account      string `json:"account" form:"account"`
	HeadIcon     string `json:"head_icon" form:"head_icon"`
	UserCity     string `json:"user_city" form:"user_city"`
	UserCityCode string `json:"user_city_code" form:"user_city_code"`
}
