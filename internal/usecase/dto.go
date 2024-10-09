package usecase

type MonsterRating struct {
	No           int
	Name         string
	Game8Monster Game8Monster
}

type Game8Monster struct {
	URL             string
	OriginMonsterNo int
	Scores          []Game8monsterScore
}

type Game8monsterScore struct {
	Name           string
	LeaderPoint    string
	SubLeaderPoint string
	AssistPoint    string
}
