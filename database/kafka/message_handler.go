package kafka

import (
	"encoding/json"
	"fmt"
	"storage-service/database"
	"storage-service/domain"
	"storage-service/tools/storagecontext"
	"time"
)

type MessageHandler interface {
	HandleMessage(ctx storagecontext.StorageContext, message []byte) error
}

type Impl struct {
	storageRepository database.StorageRepository
	kafkaProducer     *Producer
}

func NewMessageHandler(storageRepository database.StorageRepository, producer *Producer) MessageHandler {
	return &Impl{
		storageRepository: storageRepository,
		kafkaProducer:     producer,
	}
}

func (h *Impl) HandleMessage(ctx storagecontext.StorageContext, message []byte) error {
	var approvedItems []database.ApprovedItem
	if err := json.Unmarshal(message, &approvedItems); err != nil {
		return fmt.Errorf("не удалось десериализовать сообщение: %w", err)
	}

	for _, approvedItem := range approvedItems {
		ctx.Log().Info(fmt.Sprintf("Продукт %s был подтвержден в %s", approvedItem.ProductId, approvedItem.ApproveTime.Format(time.RFC3339)))

		product, err := h.storageRepository.GetProduct(ctx, approvedItem.ProductId)
		if err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось получить продукт из базы: %v", err))
			continue
		}

		product.Status = string(domain.Approved)
		product.ApproveTime = approvedItem.ApproveTime

		if err = h.storageRepository.UpdateProducts(ctx, []database.Product{product}); err != nil {
			ctx.Log().Error(fmt.Sprintf("не удалось обновить продукт в базы: %v", err))
			continue
		}

		ctx.Log().Info(fmt.Sprintf("Статус продукта %s был обновлен", approvedItem.ProductId))
	}

	return nil
}
