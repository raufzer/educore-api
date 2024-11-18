package models

type Question struct {
	Text       string   `bson:"text" json:"text"`
	Options    []string `bson:"options" json:"options"`
	CorrectAns int      `bson:"correct_ans" json:"correctAns"`
}
