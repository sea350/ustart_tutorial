package storage

import "context"

// Storage is a database-agnostic interface for persisting auth data
type Storage interface {
	Insert(context.Context, string, string, string, int) error
}
