package model

type Word struct {
	Text       string `gorm: "type: varchar(50)"`
	Frequency  int
	Complexity int
	LanguageID int
	Language   Language
}

type WordService interface {
	GetTranslations() []Word
	GetSynonyms() []Word
}

type WordRepository interface {
	GetTranslations() []Word
	GetSynonyms() []Word
}

type wordService struct {
	Word
	repo WordRepository
}

func NewWordService(repository WordRepository) WordService {
	return &wordService{
		repo: repository,
	}
}

func (w wordService) GetTranslations() []Word {
	return w.repo.GetTranslations()
}

func (w wordService) GetSynonyms() []Word {
	return w.repo.GetSynonyms()
}
