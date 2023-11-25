package entities

type Word struct {
	Notation     string        `json:"notation"`
	Translations []Translation `json:"translations"`
	UseCases     []UseCase     `json:"useCases"`
}

type UseCase struct {
	Text string `json:"text"`
}

type Translation struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
