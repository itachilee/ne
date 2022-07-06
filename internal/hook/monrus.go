package hook

import (
	"context"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type hooker struct {
	c *mongo.Collection
}

type M bson.M

func NewHookerFromCollection(collection *mongo.Collection) *hooker {
	return &hooker{c: collection}
}
func NewHookerWithAuth(mgoUrl, db, collection, user, pass string) (*hooker, error) {
	var err error
	var mgoCli *mongo.Client
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s/%s",
		user, pass, mgoUrl, db,
	))

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &hooker{c: mgoCli.Database(db).Collection(collection)}, nil
}

func (h *hooker) Fire(entry *logrus.Entry) error {
	data := make(logrus.Fields)
	data["Level"] = entry.Level.String()
	data["Time"] = entry.Time
	data["Message"] = entry.Message

	for k, v := range entry.Data {
		if errData, isError := v.(error); logrus.ErrorKey == k && v != nil && isError {
			data[k] = errData.Error()
		} else {
			data[k] = v
		}
	}

	_, mgoErr := h.c.InsertOne(context.TODO(), M(data))

	if mgoErr != nil {
		return fmt.Errorf("failed to send log entry to mongodb: %v", mgoErr)
	}

	return nil
}

func (h *hooker) Levels() []logrus.Level {
	return logrus.AllLevels
}
