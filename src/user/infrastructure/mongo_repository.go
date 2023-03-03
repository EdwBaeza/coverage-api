package infrastructure

import (
	"context"
	"os"

	domain "github.com/edwbaeza/coverage-api/src/user/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	COLLECTION = "users"
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

func (repository *MongoRepository) Create(user domain.User) (domain.User, error) {

	result, err := repository.collection.InsertOne(repository.context, user)

	if err != nil {
		return domain.User{}, err
	}

	userCreated, _ := repository.Find(result.InsertedID.(primitive.ObjectID).Hex())

	return userCreated, err
}

func (repository *MongoRepository) findBy(filters map[string]interface{}) (domain.User, error) {
	user := domain.User{}
	errFind := repository.collection.FindOne(repository.context, filters).Decode(&user)

	return user, errFind
}

func (repository *MongoRepository) Find(id string) (domain.User, error) {
	objectID, _ := primitive.ObjectIDFromHex(id)
	user := domain.User{}
	errFind := repository.collection.FindOne(repository.context, bson.M{"_id": objectID}).Decode(&user)

	return user, errFind
}

func (repository *MongoRepository) FindByEmail(email string) (domain.User, error) {
	return repository.findBy(bson.M{"email": email})
}

func (repository *MongoRepository) Update(id string, user domain.User) (domain.User, error) {
	var userUpdated domain.User
	objectID, _ := primitive.ObjectIDFromHex(id)
	errFind := repository.collection.FindOneAndUpdate(repository.context, bson.M{"_id": objectID}, bson.M{"$set": user}).Decode(&userUpdated)

	return user, errFind
}
