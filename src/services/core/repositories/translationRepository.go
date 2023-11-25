package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"golang-app/src/data/databases/mongo"
	"golang-app/src/services/core/entities"
	"golang-app/utils/errorsAndPanics"
)

type translationRepository struct {
	collectionName string
	core           *mongo.RepositoryCore
}

func NewTranslationRepository(core *mongo.RepositoryCore) translationRepository {
	return translationRepository{"translations", core}
}

func (t translationRepository) Insert(word *entities.Word) {
	t.core.InsertOne(t.collectionName, word)
}

func (t translationRepository) SelectByWord(word string) []entities.Word {
	var result []entities.Word
	cursor := t.core.Find(t.collectionName, bson.D{{Key: "notation", Value: word}})

	for cursor.Next(context.Background()) {
		var person entities.Word
		err := cursor.Decode(&person)
		errorsAndPanics.HandleError(err)
		result = append(result, person)
	}

	return result
}

// WORK IN PROGRESS
func (t translationRepository) SelectByTranslation(translation string) []entities.Word {
	var result = make([]entities.Word, 4)
	err := t.core.Find(t.collectionName, bson.D{
		{
			Key: "translations",
			Value: bson.D{{
				Key: "$elemMatch",
				Value: bson.D{{
					Key:   "Text",
					Value: translation,
				}},
			}},
		},
	}).Decode(&result)
	errorsAndPanics.HandleError(err)
	return result
}
