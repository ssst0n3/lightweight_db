package test_data

import (
	"github.com/ssst0n3/awesome_libs"
)

/*
CREATE TABLE IF NOT EXISTS challenge
(
    id    integer primary key autoincrement,
    name  text,
    score integer
);
*/
type ResourceWrapper struct {
	Id       uint
	Resource interface{}
}

type Challenge struct {
	Name   string                            `json:"name"`
	Score  int                               `json:"score"`
	Solved awesome_libs.NumberCompatibleBool `json:"solved"`
}

type ChallengeWithId struct {
	Id uint `json:"id"`
	Challenge
}

const (
	TableNameChallenge = "challenge"
)

const (
	ColumnNameChallengeName = "name"
)

var Challenge1 = ChallengeWithId{
	Id: 1,
	Challenge: Challenge{
		Name:   "name",
		Score:  10,
		Solved: true,
	},
}

var ChallengeWithIdList = []ChallengeWithId{
	Challenge1,
}

var Challenges = func() (resourceWrapper []ResourceWrapper) {
	for _, resourceWithId := range ChallengeWithIdList {
		resourceWrapper = append(resourceWrapper, ResourceWrapper{
			Id:       resourceWithId.Id,
			Resource: resourceWithId.Challenge,
		})
	}
	return resourceWrapper
}()

var Challenge1Update = ChallengeWithId{
	Id: Challenge1.Id,
	Challenge: Challenge{
		Name:   Challenge1.Name,
		Score:  20,
		Solved: true,
	},
}

func Bool2Int64(b awesome_libs.NumberCompatibleBool) int64 {
	if b {
		return int64(1)
	} else {
		return int64(0)
	}
}

var Challenge1FromDbSimulate = awesome_libs.Dict{
	"id":     int64(Challenge1.Id),
	"name":   Challenge1.Name,
	"score":  int64(Challenge1.Score),
	"solved": Bool2Int64(Challenge1.Solved),
}
