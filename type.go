package loginqr

type WhatsauthRequest struct {
	Uuid        string `json:"uuid,omitempty" bson:"uuid,omitempty"`
	Phonenumber string `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	Delay       uint32 `json:"delay,omitempty" bson:"delay,omitempty"`
}

type Profile struct {
	Token       string `bson:"token"`
	Phonenumber string `bson:"phonenumber"`
}

type Response struct {
	Response string `json:"response"`
	Info     string `json:"info,omitempty"`
}
