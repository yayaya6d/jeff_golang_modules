package mongodb

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var DB = MongoDB{}

type MongoDB struct {
	db          *mongo.Database
	client      *mongo.Client
	collections map[string]*mongo.Collection
	rwlock      sync.RWMutex
}

func (m *MongoDB) Init(uri string) error {
	if m.db != nil {
		return nil
	}

	cs, err := connstring.Parse(uri)
	if err != nil {
		return err
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}

	m.client = client
	m.db = client.Database(cs.Database)
	m.collections = make(map[string]*mongo.Collection)

	return nil
}

func (m *MongoDB) Release() {
	if m.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		if err := m.client.Disconnect(ctx); err != nil {
			println("error occurred when disconnect DB, err = %s", err.Error())
		}
	}
	m.db = nil
	m.collections = make(map[string]*mongo.Collection)
}

func (m *MongoDB) Collection(name string) *mongo.Collection {
	if m.db == nil {
		panic("please call Init() first.")
	}

	m.rwlock.RLock()
	c, ok := m.collections[name]
	m.rwlock.RUnlock()

	if ok {
		return c
	}

	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.collections[name] = m.db.Collection(name)

	return m.collections[name]
}
