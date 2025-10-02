package entities

type Localization struct {
	Scenarios []*Scenario `json:"localizations"`
}

type Scenario struct {
	Language string `json:"language"`
	HelloMes string `json:"hello"`
	GoodMes  string `json:"good"`
	BadMes   string `json:"bad"`
	WrongMes string `json:"wrong"`
}
