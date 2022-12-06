package factory

import "github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/infra/kafka/event"

func CreateProcessMessageStrategy(topic string) event.ProcessEventStrategy {
	switch topic {
	case "chooseTeam":
		return event.ProcessChooseTeam{}
	case "newPlayer":
		return event.ProcessNewPlayer{}
	case "newMatch":
		return event.ProcessNewMatch{}
	case "newAction":
		return event.ProcessNewAction{}
	case "matchUpdateResult":
		return event.ProcessMatchUpdateResult{}
	}
	return nil
}

// Test data kafka
// chooseTeam: {"my_team_id":"1", "players":["1","2","3","4","5"]}
// newPlayer: {"id": "10","name": "Wesley","initial_price": 10.5}
// newMatch: {"id":"3","match_date":"2021-05-01T00:00:00Z","team_a_id":"1","team_b_id":"2"}
// newAction: {"match_id":"3","team_id":"1","player_id":"1","action":"goal","minutes":10}
// matchUpdateResult: {"match_id":"3","result":"2-0"}
