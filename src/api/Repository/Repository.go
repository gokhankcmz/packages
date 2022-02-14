package CustomerRepository

import (
	"Packages/src/api/Type/EntityTypes"
	"Packages/src/api/Type/ErrorTypes"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type IRepository interface {
	GetSingle(ID string, model interface{})
	GetMany(options *options.FindOptions, filter *bson.M) []EntityTypes.User
	Create(c *EntityTypes.User) string
	Update(c *EntityTypes.User) *mongo.UpdateResult
	Delete(ID string) *mongo.DeleteResult
}

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *Repository {
	return &Repository{mc: mc}
}

func (r Repository) GetSingle(ID string, model interface{}) {
	objID, err := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	if err = r.mc.FindOne(context.Background(), filter).Decode(model); err != nil {
		if err == mongo.ErrNoDocuments {
			panic(ErrorTypes.EntityNotFound.SetArgs(ID))
		}
	}
}

func (r Repository) GetMany(options *options.FindOptions, filter *bson.M) []EntityTypes.User {
	var results []EntityTypes.User
	cur, _ := r.mc.Find(context.Background(), filter, options)
	defer cur.Close(context.Background())
	cur.All(context.Background(), &results)
	return results
}

func (r Repository) Create(c *EntityTypes.User) string {
	now := time.Now()
	c.CreatedAt = now
	c.UpdatedAt = now
	res, _ := r.mc.InsertOne(context.Background(), c)
	return res.InsertedID.(primitive.ObjectID).Hex()
}

func (r Repository) Update(c *EntityTypes.User) *mongo.UpdateResult {
	c.UpdatedAt = time.Now()
	filter := bson.M{"_id": c.ID}
	res, _ := r.mc.ReplaceOne(context.Background(), filter, c)
	return res
}

func (r Repository) Delete(ID string) *mongo.DeleteResult {
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": objID}
	res, _ := r.mc.DeleteOne(context.Background(), filter)
	return res
}
