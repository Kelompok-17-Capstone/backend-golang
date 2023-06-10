package payload

type GetHome struct {
	Saldo    int `json:"saldo" form:"saldo"`
	Products []ProductsHome
}

type ProductsHome struct {
	Name  string `json:"name" form:"name"`
	Image string `json:"image" form:"image"`
}
