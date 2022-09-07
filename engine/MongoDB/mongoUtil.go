package MongoDB

import (
	"BurninProject/aop/logger"
	"go.mongodb.org/mongo-driver/bson"
)

func (conn *MongoConn) buildQueryFilter(userId string, time string) interface{} {
	filterData := make(map[string]interface{})
	filterData["userid"] = userId
	filterData["time"] = time

	filter := bson.M{}
	data, err := bson.Marshal(filterData)
	if err != nil {
		logger.Logger.ErrorF("marshal error: %v", err)
		return filter
	}

	err = bson.Unmarshal(data, filter)
	if err != nil {
		logger.Logger.ErrorF("unmarshal error: %v", err)
	}
	logger.Logger.InfoF("filter: %v", filter)
	return filter
}
