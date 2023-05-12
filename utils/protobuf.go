package utils

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"stac/database"
)

func GetProtoFromDB(key string, m proto.Message) error {
	rawPB, err := database.DB.Get([]byte(key), nil)
	if CheckError(err) {
		return errors.New(`db error`)
	}
	err = proto.Unmarshal(rawPB, m)
	if CheckError(err) {
		return errors.New(`proto unmarshal error`)
	}
	return nil
}

func PutProtoToDB(key string, m proto.Message) error {
	out, err := proto.Marshal(m)
	if err != nil {
		return errors.New(`proto marshal error`)
	}
	if err := database.DB.Put([]byte(key), out, nil); err != nil {
		return errors.New(`db error`)
	}
	return nil
}
