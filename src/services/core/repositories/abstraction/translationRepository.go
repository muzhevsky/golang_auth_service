package abstraction

import (
	"golang-app/src/services/core/entities"
)

type TranslationRepository interface {
	Insert(word *entities.Word)
	SelectByWord(word string) []entities.Word
	SelectByTranslation(translation string) []entities.Word
}
