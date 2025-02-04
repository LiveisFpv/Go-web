package queries

import "context"

//TODO work with database
func (q *Queries) Login(ctx context.Context, login, password string) (string, error) {
	return "", nil
}
func (q *Queries) Register(ctx context.Context, email, login, password string) error {
	return nil
}
