package account

type Account interface {
	GetName() string
	GetId() int64
	GetPhone() string
}

type WordsHistory interface {
	GetWords(id int64) []string
	GetPopularWords(id int64) []map[string]int32
}
