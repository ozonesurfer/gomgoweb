package controllers

import (
	"fmt"
	"gomgoweb"
	"gomgoweb/models"
	"labix.org/v2/mgo/bson"
	"strings"
)

func ToObjectId(in string) bson.ObjectId {
	var result bson.ObjectId
	id1 := strings.Replace(in, "ObjectIdHex(", "", -1)
	//	id1 := strings.Replace(in, "ObjectIdHex%28", "", -1)

	id1 = strings.Replace(id1, ")", "", -1)
	//	id1 = strings.Replace(id1, "%29", "", -1)
	id1 = strings.Replace(id1, "\"", "", -1)
	//	id1 = strings.TrimSpace(id1)
	result = bson.ObjectIdHex(id1)
	return result
}

type MyQuery struct {
	Name string `bson:"Name"`
}

func (q *MyQuery) ConvertToInterface() interface{} {
	return map[string]interface{}{"Name": q.Name}
}

func DocDump() {
	database, session := models.GetDB()
	fmt.Println("session started")
	col := database.C(gomgoweb.BAND_COL)
	fmt.Println("Got collection")
	q := col.Find(bson.M{"Values.Name": "test1"})
	count, err := q.Count()
	if err != nil {
		println("Find error:", err)
	} else {
		fmt.Println("Count =", count)
	}
	session.Close()	
}
