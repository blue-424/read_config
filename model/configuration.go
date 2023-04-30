package Cnf

type GiftItem struct {
	Gift string `json:"gift"`
}

type AdressItem struct {
	Street      string
	State       string
	City        string
	Postal_Code string
}

type Configuration struct {
	Username string     `json:"username"`
	Password string     `json:"password"`
	Flag     bool       `json:"flag"`
	Products []int      `jason:"products"`
	Gifts    []GiftItem `json:"gifts"`
	Adr      AdressItem `json:"adr"`
	Birth    string     `json:"birth"`
}

func (a AdressItem) ToString() string {
	str := a.City + " " + a.State + " " + a.Street + " " + a.Postal_Code
	return str
}
