package model

type LineDocument struct {
	User     string    `bson:"user"`
	Messages []Message `bson:"messages"`
}

type Message struct {
	Time    int64  `bson:"time"`
	Content string `bson:"content"`
}
