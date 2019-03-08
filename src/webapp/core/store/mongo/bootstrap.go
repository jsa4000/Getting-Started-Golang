package mongo

import (
	"context"
	log "webapp/core/logging"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// Boostrap Provide an abstraction for an initial bootstraping
func Boostrap(collection *mongo.Collection, oo []interface{}) {
	ordered := false
	result, err := collection.InsertMany(context.Background(), oo, &options.InsertManyOptions{
		Ordered: &ordered,
	})
	if result == nil {
		log.Errorf("%s Bootstapping: %s", collection.Name(), err)
		return
	}
	insertedIDs := len(result.InsertedIDs)
	if err != nil {
		if errors, ok := err.(mongo.BulkWriteException); ok {
			errorIDs := len(errors.WriteErrors)
			if errorIDs != insertedIDs {
				log.Debugf("Inserted %d new default %s", insertedIDs-errorIDs, collection.Name())
			}
		}
		return
	}
	log.Debugf("Inserted %d new default %s", insertedIDs, collection.Name())
}
