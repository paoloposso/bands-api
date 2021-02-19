package repositorymongodb

import (
	customerrors "bands-api/custom_errors"
	"bands-api/user"
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const collection string = "users"

type mongoRepository struct {
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

// NewMongoRepository returns a reference to an implementation of UserRepository interface that implements the comunication with MongoDb for the User domain
func NewMongoRepository(mongoURL, database string, mongoTimeout int) (user.Repository, error)  {
	repo := &mongoRepository{
		timeout: time.Duration(mongoTimeout)*time.Second,
		database: database,
	} 
	client, err := newMongoDbClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, &customerrors.DBConnectionError{ Err: errors.WithStack(err) }
	}
	repo.client = client
	return repo, nil
}

func (r *mongoRepository) Create(user *user.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(collection)
	_, err := collection.InsertOne(ctx, user)
	if err != nil{
		return &customerrors.DBConnectionError{ Err: errors.WithStack(err) }
	}
	return nil
}

func (r *mongoRepository) GetByEmail(email string) (*user.User, error) {
	return nil, nil
}
