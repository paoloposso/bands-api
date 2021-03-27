package repositorymongodb

import (
	"bands-api/user"
	"context"
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const collection string = "users"

type mongoUserRepository struct {
	client *mongo.Client
	database string
	timeout time.Duration
}

func newMongoDbClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, err
}

// NewMongoUserRepository returns a reference to an implementation of UserRepository interface that implements the comunication with MongoDb for the User domain
func NewMongoUserRepository(mongoURL, database string, mongoTimeout int) (user.Repository, error)  {
	repo := &mongoUserRepository{
		timeout: time.Duration(mongoTimeout)*time.Second,
		database: database,
	}
	client, err := newMongoDbClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.client = client
	createIndex(*repo.client, repo.database, "email", true)
	return repo, nil
}

func (r *mongoUserRepository) Create(user *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(collection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil{
		return errors.WithStack(err) 
	}
	return nil
}

func (r *mongoUserRepository) GetByEmail(email string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	var user user.User
	collection := r.client.Database(r.database).Collection(collection)
	if err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		if strings.Contains(err.Error(), "no documents") {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *mongoUserRepository) GetByID(id string) (*user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	var user user.User
	collection := r.client.Database(r.database).Collection(collection)
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func createIndex(client mongo.Client, databaseName string, field string, unique bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
    mod := mongo.IndexModel{
        Keys: bson.M{field: 1},
		Options: options.Index().SetUnique(unique),
	}
    collection :=  client.Database(databaseName).Collection(collection)
	_, err := collection.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Fatal("Error creating index: %w", errors.WithStack(err))
		return err
    }
	return nil
}