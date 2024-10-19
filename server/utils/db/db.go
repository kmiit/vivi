package db

import (
	"context"
	"encoding/json"

	"github.com/kmiit/vivi/types"
	"github.com/kmiit/vivi/utils/log"
)

func GetNewId(ctx context.Context, namespace string) (int64, error) {
	id, err := rdb.Incr(ctx, namespace).Result()
	if err != nil {
		return -1, err
	}
	return id, err
}

func GetAllFiles(ctx context.Context, namespace string) ([]types.FDescriptor, error) {
	var (
		files  []types.FDescriptor
		cursor uint64
	)
	for {
		keys, nextCursor, err := rdb.Scan(ctx, cursor, namespace+"*", 0).Result()
		if err != nil {
			return nil, err
		}
		cursor = nextCursor
		for _, key := range keys {
			val, err := rdb.Get(ctx, key).Result()
			if err != nil {
				return nil, err
			}
			var file types.FDescriptor
			err = json.Unmarshal([]byte(val), &file)
			if err != nil {
				return nil, err
			}
			files = append(files, file)
		}

		if cursor == 0 {
			break
		}
	}
	log.W(TAG, "Get all files: ", files)
	return files, nil
}

func GetAllOuter(ctx context.Context, namespace string) ([]types.DescriptorO, error) {
	allFiles, err := GetAllFiles(ctx, namespace)
	if err != nil {
		return nil, err
	}
	files := []types.DescriptorO{}
	for _, file := range allFiles {
		files = append(files, file.Outer)
	}
	return files, nil
}

func GetKeys(ctx context.Context, namespace string) ([]string, error) {
	var (
		allKeys []string
		cursor  uint64
	)
	for {
		keys, cursor, err := rdb.Scan(ctx, cursor, namespace+"*", 0).Result()
		if err != nil {
			return nil, err
		}
		allKeys = append(allKeys, keys...)
		if cursor == 0 {
			break
		}
	}
	return allKeys, nil
}
