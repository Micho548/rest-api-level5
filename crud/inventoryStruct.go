package crud

type productIntentory struct {
	Product  Products `json:"Products"`
	Quantity int      `json:"Quantity"`
}

type Products struct {
	ID    string  `json:"ID"`
	Code  string  `json:"Code"`
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}
