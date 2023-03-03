package infrastructure

import (
	"context"
	"os"

	domain "github.com/edwbaeza/coverage-api/src/purchase/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION = "purchases"
)

// TODO: Improve mongo repository to reuse base logic (CRUD by generic types maybe)
type MongoRepository struct {
	client     mongo.Client
	context    context.Context
	collection *mongo.Collection
}

func NewMongoRepository() *MongoRepository {
	ctx, _ := context.WithCancel(context.Background())
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	return &MongoRepository{
		client:     *client,
		context:    ctx,
		collection: client.Database(os.Getenv("MONGO_DB_NAME")).Collection(COLLECTION),
	}
}
func (repository *MongoRepository) Create(purchase domain.Purchase) (domain.Purchase, error) {
	result, err := repository.collection.InsertOne(repository.context, purchase)

	if err != nil {
		return domain.Purchase{}, err
	}

	purchaseCreated, err := repository.Find(result.InsertedID.(primitive.ObjectID).Hex())

	return purchaseCreated, err
}

func (repository *MongoRepository) Find(id string) (domain.Purchase, error) {
	purchase := domain.Purchase{}
	objectID, _ := primitive.ObjectIDFromHex(id)
	errFind := repository.collection.FindOne(repository.context, bson.M{"_id": objectID}).Decode(&purchase)

	return purchase, errFind
}
func (repository *MongoRepository) List(filters map[string]interface{}) ([]domain.Purchase, error) {
	purchases := []domain.Purchase{}
	cursor, err := repository.collection.Find(repository.context, filters)
	if err != nil {
		return purchases, err
	}

	err = cursor.All(repository.context, &purchases)

	return purchases, err
}

func (repository *MongoRepository) UpdateStatus(id string, status domain.PurchaseStatus) (domain.Purchase, error) {
	purchase := domain.Purchase{}
	objectID, _ := primitive.ObjectIDFromHex(id)
	errFind := repository.collection.FindOneAndUpdate(repository.context, bson.M{"_id": objectID}, bson.M{"$set": bson.M{"status": status}}).Decode(&purchase)

	return purchase, errFind
}
