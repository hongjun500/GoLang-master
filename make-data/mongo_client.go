package main

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoProperties struct {
	Uri      string
	DataBase string
	Password string
}

func (p MongoProperties) NewMongoClient() (*mongo.Client, error) {
	if p.Uri == "" {
		return nil, errors.New("no Uri Provided")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(p.Uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// 在连接后进行 Ping 操作以验证 MongoDB 服务是否可用
	err = client.Database("mongodb").Client().Ping(context.Background(), nil)
	if err != nil {
		// 更详细的错误返回
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	// 在这里不使用 defer 关闭连接，而是交给调用者来控制连接关闭的时机
	return client, nil
}
