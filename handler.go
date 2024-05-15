package loginqr

import (
	"strings"

	"github.com/aiteung/module/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlerQRLogin(msg model.IteungMessage, WAKeyword string, WAPhoneNumber string, db *mongo.Database, WAAPIQRLogin string) (resp Response, err error) {
	dt := &WhatsauthRequest{
		Uuid:        GetUUID(msg, WAKeyword),
		Phonenumber: msg.Phone_number,
		Delay:       msg.From_link_delay,
	}
	structtoken, err := WAAPIToken(WAPhoneNumber, db)
	if err != nil {
		return
	}
	resp, err = PostStructWithToken[Response]("Token", structtoken.Token, dt, WAAPIQRLogin)
	return
}

func WAAPIToken(phonenumber string, db *mongo.Database) (apitoken Profile, err error) {
	filter := bson.M{"phonenumber": phonenumber}
	apitoken, err = GetOneDoc[Profile](db, "profile", filter)

	return
}

func GetUUID(msg model.IteungMessage, keyword string) string {
	return strings.Replace(msg.Message, keyword, "", 1)
}
