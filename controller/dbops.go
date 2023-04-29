// Defined some database operations, used in development only

package controller

import (
	"net/http"
	"stac/database"
	"stac/models"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb/util"
	"google.golang.org/protobuf/proto"
)

type DBresult struct {
	Key   string             `json:"key"`
	Value *models.GithubHook `json:"protobuf"`
}

func GetAllFromDB(c *gin.Context) {
	iter := database.DB.NewIterator(nil, nil)
	defer iter.Release()

	var res []DBresult

	for iter.Next() {
		key := iter.Key()
		val := iter.Value()

		pb := &models.GithubHook{}
		proto.Unmarshal(val, pb)

		res = append(res, DBresult{Key: string(key), Value: pb})
	}
	c.JSON(http.StatusOK, res)
}

func DeleteAllFromDB(c *gin.Context) {
	iter := database.DB.NewIterator(nil, nil)
	defer iter.Release()

	for iter.Next() {
		database.DB.Delete(iter.Key(), nil)
	}
	database.DB.CompactRange(util.Range{Start: nil, Limit: nil})
	c.JSON(http.StatusOK, OPSuccess)
}

func DeleteSingleKey(c *gin.Context) {
	database.DB.Delete([]byte(c.GetHeader("key")), nil)
	c.JSON(http.StatusOK, OPSuccess)
}
