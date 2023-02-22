package usecase

type TwinComponent struct {
	Id      string      `json:"id"`
	Name    string      `json:"name"`
	Classes []TwinClass `json:"classes"`
	Service TwinService `json:"service"`
}

type TwinClass struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TwinProperty struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type TwinEnum struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type TwinService struct{}
