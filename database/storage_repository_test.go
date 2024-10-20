package database

import (
	"context"
	"encoding/json"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"sync"
	"testing"
)

const (
	_oneHundred         = 100
	_oneHundredThousand = 100_000

	_batchSize = 100

	_mongoConnectionString = "mongodb://localhost:27017/?direct=true"
	_mongoDatabaseName     = "storage-service"
	_mongoCollectionName   = "products"

	_testDataPath = "/Users/ilyaantonov/Downloads/4 курс/microservices/storage-service/database/mocks/test_data.json"
)

// Тест добавления 100 элементов в MongoDB
func TestInsert100Documents(t *testing.T) {
	ctx := context.Background()

	testDataBytes, err := os.ReadFile(_testDataPath)
	if err != nil {
		t.Fatalf("не удалось прочитать файл с тестовыми даннымим: %v", err)
	}

	var products []Product
	if err = json.Unmarshal(testDataBytes, &products); err != nil {
		t.Fatalf("не удалось анмаршалить файл с тестовыми даннымим: %v", err)
	}

	testData := ToInsertItems(products)

	client, err := MongoConnect(ctx, _mongoConnectionString)
	if err != nil {
		t.Fatalf("не удалось подключиться к монге: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(_mongoDatabaseName).Collection(_mongoCollectionName)

	data := make([]any, 0, _oneHundred)
	for i := 0; i < _oneHundred/len(testData); i++ {
		data = append(data, toAny(testData)...)
	}

	// Вставляем документы
	res, err := collection.InsertMany(ctx, data)
	if err != nil {
		t.Fatalf("Ошибка при вставке 100 документов: %v", err)
	}

	if len(res.InsertedIDs) != _oneHundred {
		t.Fatalf("Ожидалось 100 документов, но вставлено %d", len(res.InsertedIDs))
	}
}

// Тест добавления 100 000 элементов в MongoDB
func TestInsert100000Documents(t *testing.T) {
	ctx := context.Background()

	testDataBytes, err := os.ReadFile(_testDataPath)
	if err != nil {
		t.Fatalf("не удалось прочитать файл с тестовыми даннымим: %v", err)
	}

	var products []Product
	if err = json.Unmarshal(testDataBytes, &products); err != nil {
		t.Fatalf("не удалось анмаршалить файл с тестовыми даннымим: %v", err)
	}

	testData := ToInsertItems(products)

	client, err := MongoConnect(ctx, _mongoConnectionString)
	if err != nil {
		t.Fatalf("не удалось подключиться к монге: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(_mongoDatabaseName).Collection(_mongoCollectionName)

	data := make([]any, 0, _oneHundredThousand)
	for i := 0; i < _oneHundredThousand/len(testData); i++ {
		data = append(data, toAny(testData)...)
	}

	dataBatches := lo.Chunk(data, _batchSize)

	totalInserted := 0
	wg := sync.WaitGroup{}
	for _, batch := range dataBatches {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, insertErr := collection.InsertMany(ctx, batch)
			if insertErr != nil {
				t.Errorf("Ошибка при вставке документов: %v", err)
				return
			}

			totalInserted += len(res.InsertedIDs)
		}()
	}

	wg.Wait()

	// Вставляем документы
	if totalInserted != _oneHundredThousand {
		t.Fatalf("Ожидалось 100000 документов, но вставлено %d", totalInserted)
	}
}

// Тест удаления всех элементов из MongoDB
func TestDeleteAllDocuments(t *testing.T) {
	ctx := context.Background()

	client, err := MongoConnect(ctx, _mongoConnectionString)
	if err != nil {
		t.Fatalf("не удалось подключиться к монге: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database(_mongoDatabaseName).Collection(_mongoCollectionName)

	// Удаляем все документы
	res, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Ошибка удаления всех документов: %v", err)
	}

	if res.DeletedCount == 0 {
		t.Fatalf("Ожидалось удаление всех документов, но удалено %d", res.DeletedCount)
	}

	count, err := collection.CountDocuments(ctx, bson.D{})
	if err != nil {
		t.Fatalf("Ошибка подсчета документов после удаления: %v", err)
	}

	if count != 0 {
		t.Fatalf("Ожидалось 0 документов после удаления, но найдено %d", count)
	}
}
