package test_data

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
	Name  string `json:"name"`
	Score int    `json:"score"`
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
		Name:  "name",
		Score: 10,
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
		Name:  Challenge1.Name,
		Score: 20,
	},
}
