package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"os"
	"storage-service/settings"
	"testing"
)

const (
	_redisAddress = "localhost:6379"
)

func TestInsert100ToCache(t *testing.T) {
	ctx := context.Background()

	testDataBytes, err := os.ReadFile(_testDataPath)
	if err != nil {
		t.Fatalf("не удалось прочитать файл с тестовыми даннымим: %v", err)
	}

	var products []Product
	if err = json.Unmarshal(testDataBytes, &products); err != nil {
		t.Fatalf("не удалось анмаршалить файл с тестовыми даннымим: %v", err)
	}

	redisClient, err := NewRedisClient(ctx, settings.RedisSettings{Address: _redisAddress})
	if err != nil {
		t.Fatalf("не удалось подлючиться к redis: %v", err)
	}

	testData := make([]Product, 0, _oneHundred)
	for i := 0; i < _oneHundred/len(products); i++ {
		testData = append(testData, products...)
	}

	inserted := 0

	for _, product := range testData {
		product.Id = uuid.NewString()

		if err = redisClient.Set(ctx, getKey(product.Id), ToProtoProduct(product), _defaultTTL).Err(); err != nil {
			t.Error(err)
			continue
		}

		inserted++
	}

	if inserted != _oneHundred {
		t.Fatalf("ожидалась вставка 100 элементов, а вставилось %d", inserted)
	}
}

func TestInsert100000ToCache(t *testing.T) {
	ctx := context.Background()

	testDataBytes, err := os.ReadFile(_testDataPath)
	if err != nil {
		t.Fatalf("не удалось прочитать файл с тестовыми даннымим: %v", err)
	}

	var products []Product
	if err = json.Unmarshal(testDataBytes, &products); err != nil {
		t.Fatalf("не удалось анмаршалить файл с тестовыми даннымим: %v", err)
	}

	redisClient, err := NewRedisClient(ctx, settings.RedisSettings{Address: _redisAddress})
	if err != nil {
		t.Fatalf("не удалось подлючиться к redis: %v", err)
	}

	testData := make([]Product, 0, _oneHundredThousand)
	for i := 0; i < _oneHundredThousand/len(products); i++ {
		testData = append(testData, products...)
	}

	inserted := 0

	for _, product := range testData {
		product.Id = uuid.NewString()

		if err = redisClient.Set(ctx, getKey(product.Id), ToProtoProduct(product), _defaultTTL).Err(); err != nil {
			t.Error(err)
			continue
		}

		inserted++
	}

	if inserted != _oneHundred {
		t.Fatalf("ожидалась вставка 100000 элементов, а вставилось %d", inserted)
	}
}

//func DeleteAll(t *testing.T) {
//	ctx := context.Background()
//
//	testDataBytes, err := os.ReadFile(_testDataPath)
//	if err != nil {
//		t.Fatalf("не удалось прочитать файл с тестовыми даннымим: %v", err)
//	}
//
//	var products []Product
//	if err = json.Unmarshal(testDataBytes, &products); err != nil {
//		t.Fatalf("не удалось анмаршалить файл с тестовыми даннымим: %v", err)
//	}
//
//	redisClient, err := NewRedisClient(ctx, settings.RedisSettings{Address: _redisAddress})
//	if err != nil {
//		t.Fatalf("не удалось подлючиться к redis: %v", err)
//	}
//}

func getKey(guid string) string {
	return fmt.Sprintf("%s:%s", guid, _productsKey)
}
