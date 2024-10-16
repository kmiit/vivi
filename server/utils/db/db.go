package db

import "context"

func GetNewId(ctx context.Context, namespace string) (int64, error) {
	id, err := RDB.Incr(ctx, namespace).Result()
	if err != nil {
		return -1, err
	}
	return id, err
}
