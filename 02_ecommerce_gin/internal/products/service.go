package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) (error)
}