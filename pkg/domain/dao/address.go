package dao


type Address struct {
	Street   string `bson:"street"`
	PinCode  string `bson:"pin_code"`
	City     string `bson:"city"`
	District string `bson:"district"`
	State    string `bson:"state"`
	Country  string `bson:"country"`
}