package register

import "context"

type Register interface {
	Registered(ctx context.Context, name, host string)
	Query(name string) (string, error)
}
