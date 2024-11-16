package database

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"user-service/domain"
	"user-service/tools/usercontext"
)

type UserRepository interface {
	GetUser(ctx usercontext.UserContext, id string) (User, error)
	GetUsers(ctx usercontext.UserContext, limit int64, cursor string) ([]User, error)
	SaveUsers(ctx usercontext.UserContext, users []User) error
	UpdateUsers(ctx usercontext.UserContext, users []User) error
	DeleteUsers(ctx usercontext.UserContext, ids []string) error
}

type RepositoryImpl struct {
	mongoClient *mongo.Client
	database    string
	collection  string
}

func NewUserRepository(mongoClient *mongo.Client, database, collection string) UserRepository {
	return &RepositoryImpl{
		mongoClient: mongoClient,
		database:    database,
		collection:  collection,
	}
}

func (r *RepositoryImpl) GetUser(ctx usercontext.UserContext, id string) (user User, err error) {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return User{}, err
	}

	filter := bson.M{"_id": objID}
	return user, collection.FindOne(ctx.Ctx(), filter).Decode(&user)
}

func (r *RepositoryImpl) GetUsers(ctx usercontext.UserContext, limit int64, cursor string) (users []User, err error) {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)

	filter := bson.M{}
	if cursor != "" {
		objId, err := primitive.ObjectIDFromHex(cursor)
		if err != nil {
			return nil, err
		}

		filter = bson.M{
			"_id": bson.M{
				"$gt": objId,
			},
		}
	}

	findOptions := options.Find().SetLimit(limit).SetSort(bson.M{"_id": 1})

	collectionCursor, err := collection.Find(ctx.Ctx(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	defer collectionCursor.Close(ctx.Ctx())

	return users, collectionCursor.All(ctx.Ctx(), &users)
}

func (r *RepositoryImpl) SaveUsers(ctx usercontext.UserContext, users []User) error {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)
	_, err := collection.InsertMany(ctx.Ctx(), toAny(users))
	return err
}

func (r *RepositoryImpl) UpdateUsers(ctx usercontext.UserContext, users []User) error {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)

	operations := make([]mongo.WriteModel, 0, len(users))
	for _, user := range users {
		objID, err := primitive.ObjectIDFromHex(user.Id)
		if err != nil {
			return err
		}

		filter := bson.M{
			"_id": objID,
		}

		updatedProduct := User{
			Surname:  user.Surname,
			Name:     user.Name,
			Lastname: user.Lastname,
			Role:     user.Role,
		}

		update := bson.M{
			"$set": updatedProduct,
		}

		operations = append(operations, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update))
	}

	bulkOptions := options.BulkWrite().SetOrdered(false)
	_, err := collection.BulkWrite(ctx.Ctx(), operations, bulkOptions)
	return err
}

func (r *RepositoryImpl) DeleteUsers(ctx usercontext.UserContext, ids []string) error {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)

	objIds, err := toPrimitiveObjectIds(ids)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": bson.M{
			"$in": objIds,
		},
	}

	res, err := collection.DeleteMany(ctx.Ctx(), filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return domain.NoDocumentAffected
	}

	return err
}

func toAny[T any](any []T) (result []any) {
	for _, v := range any {
		result = append(result, v)
	}
	return result
}

func toPrimitiveObjectIds(ids []string) ([]primitive.ObjectID, error) {
	objectIds := make([]primitive.ObjectID, 0, len(ids))
	for _, id := range ids {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, err
		}
		objectIds = append(objectIds, objectId)
	}
	return objectIds, nil
}
