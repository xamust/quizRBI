package model

type Data struct {
	ID         int    `json:"id" bson:"id"`
	SimpleData string `json:"simple_data" bson:"simple_data"`
	InsertedID string `json:"inserted_id" bson:"inserted_id"`
	Error      string `json:"error"`
}
