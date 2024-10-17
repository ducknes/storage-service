package database

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"storage-service/tools/storagecontext"

	"go.mongodb.org/mongo-driver/mongo"
)

type StorageRepository interface {
	GetProducts(ctx storagecontext.StorageContext, limit int64, cursor string) ([]Product, error)
	GetProduct(ctx storagecontext.StorageContext, productId string) (Product, error)
	AddProducts(ctx storagecontext.StorageContext, products []Product) error
	UpdateProducts(ctx storagecontext.StorageContext, products []Product) error
	DeleteProducts(ctx storagecontext.StorageContext, productIds []string) error

	TestData() error
}

type StorageRepositoryImpl struct {
	mongoClient *mongo.Client
	database    string
	collection  string
}

func NewStorageRepository(mongoClient *mongo.Client, database, collection string) StorageRepository {
	return &StorageRepositoryImpl{
		mongoClient: mongoClient,
		database:    database,
		collection:  collection,
	}
}

func (r *StorageRepositoryImpl) GetProducts(ctx storagecontext.StorageContext, limit int64, cursor string) (products []Product, err error) {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)

	filter := bson.M{}
	if cursor != "" {
		filter = bson.M{
			"_id": bson.M{
				"$gt": cursor,
			},
		}
	}

	findOptions := options.Find().SetLimit(limit).SetSort(bson.M{"_id": 1})

	collectionCursor, err := collection.Find(ctx.Ctx(), filter, findOptions)
	if err != nil {
		return nil, err
	}

	defer collectionCursor.Close(ctx.Ctx())

	return products, collectionCursor.All(ctx.Ctx(), &products)
}

func (r *StorageRepositoryImpl) GetProduct(ctx storagecontext.StorageContext, productId string) (product Product, err error) {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)

	objID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return product, err
	}

	filter := bson.M{
		"_id": objID,
	}

	return product, collection.FindOne(ctx.Ctx(), filter).Decode(&product)
}

func (r *StorageRepositoryImpl) AddProducts(ctx storagecontext.StorageContext, products []Product) error {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)
	_, err := collection.InsertMany(ctx.Ctx(), toAny(ToInsertItems(products)))
	return err
}

func (r *StorageRepositoryImpl) UpdateProducts(ctx storagecontext.StorageContext, products []Product) error {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)

	operations := make([]mongo.WriteModel, 0, len(products))
	for _, product := range products {
		objID, err := primitive.ObjectIDFromHex(product.Id)
		if err != nil {
			return err
		}

		filter := bson.M{
			"_id": objID,
		}

		update := bson.M{
			"$set": product,
		}

		operations = append(operations, mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update))
	}

	bulkOptions := options.BulkWrite().SetOrdered(false)
	_, err := collection.BulkWrite(ctx.Ctx(), operations, bulkOptions)
	return err
}

func (r *StorageRepositoryImpl) DeleteProducts(ctx storagecontext.StorageContext, productIds []string) error {
	collection := r.mongoClient.Database(r.database).Collection(r.collection)

	objIds, err := toPrimitiveObjectIds(productIds)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": bson.M{
			"$in": objIds,
		},
	}

	_, err = collection.DeleteMany(ctx.Ctx(), filter)
	return err
}

func (r *StorageRepositoryImpl) TestData() error {
	file, err := os.ReadFile("./database/mocks/mocks.json")
	if err != nil {
		return err
	}

	var products []Product
	if err = json.Unmarshal(file, &products); err != nil {
		return err
	}

	collection := r.mongoClient.Database(r.database).Collection(r.collection)
	_, err = collection.InsertMany(context.Background(), toAny(ToInsertItems(products)))
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
