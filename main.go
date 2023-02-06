package main

//go:generate mockgen -package main -source=main.go -destination=mocks_test.go *

import (
	"errors"
	"fmt"
	"time"
)

var (
	errTooFast = errors.New("too fast")
)

type Item struct {
	CreatedAt time.Time
}

type Store interface {
	GetLastItem() (*Item, error)
	CreateItem(item Item) error
}

type Service struct {
	store Store
	now   func() time.Time
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
		now:   time.Now,
	}
}

func (s *Service) CreateItem(item Item) error {
	lastItem, err := s.store.GetLastItem()
	if err != nil {
		return err
	}

	nextAllowedTime := lastItem.CreatedAt.Add(10 * time.Second)
	if s.now().Before(nextAllowedTime) {
		return errTooFast
	}

	return s.store.CreateItem(item)
}

func main() {
	fmt.Println("just a demo")
}
