package mongo

import (
	"context"
	"fmt"
	"time"
	log "webapp/core/logging"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
)

// CreateIndexModel fetches all the values form the database
func CreateIndexModel(name string, asc, unique bool) mongo.IndexModel {
	// Set index options (unique)
	options := options.Index()
	options.SetUnique(unique)
	// Set asc desc index
	var value int32 = 1
	if !asc {
		value = -1
	}
	index := mongo.IndexModel{}
	index.Keys = bsonx.Doc{{Key: name, Value: bsonx.Int32(value)}}
	index.Options = options
	return index
}

// CreateIndex fetches all the values form the database
func CreateIndex(ctx context.Context, c *mongo.Collection, indexes ...mongo.IndexModel) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var err error
	if len(indexes) == 1 {
		log.Debugf("Creating index %s for collection %s", indexes[0].Keys, c.Name())
		_, err = c.Indexes().CreateOne(ctx, indexes[0])
	} else {
		log.Debugf("Creating indexes for collection %s", c.Name())
		_, err = c.Indexes().CreateMany(ctx, indexes)
	}
	if err != nil {
		log.Error(fmt.Sprintf("Error creating indexes. '%s'", err))
	}
	return err
}
