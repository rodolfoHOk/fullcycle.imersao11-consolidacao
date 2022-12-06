package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/infra/db"
	httpHandler "github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/infra/http"
	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/infra/kafka/consumer"
	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/infra/repository"
	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/pkg/uow"

	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/cartola?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer dtb.Close()
	uow, err := uow.NewUow(ctx, dtb);
	if err != nil {
		panic(err)
	}
	registerRepositories(uow)

	r := chi.NewRouter()
	r.Get("/players", httpHandler.ListPlayersHandler(ctx, *db.New(dtb)))
	r.Get("/my-teams/{teamID}/players", httpHandler.ListMyTeamPlayers(ctx, *db.New(dtb)))
	r.Get("/my-teams/{teamID}/balance", httpHandler.GetMyTeamBalanceHandler(ctx, *db.New(dtb)))
	r.Get("/matches", httpHandler.ListMatchesHandler(ctx, repository.NewMatchRepository(dtb)))
	r.Get("/matches/{matchID}", httpHandler.ListMatchByIDHandler(ctx, repository.NewMatchRepository(dtb)))
	
	go http.ListenAndServe(":8080", r)

	var topics = []string {"newMatch", "chooseTeam", "newPlayer", "matchUpdateResult", "newAction"}
	msgChan := make(chan *kafka.Message)
	go consumer.Consume(topics, "broker:9094", msgChan)
	consumer.ProcessEvents(ctx, msgChan, uow)
}

func registerRepositories(uow *uow.Uow) {
	uow.Register("PlayerRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewPlayerRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MatchRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMatchRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("TeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})

	uow.Register("MyTeamRepository", func(tx *sql.Tx) interface{} {
		repo := repository.NewMyTeamRepository(uow.Db)
		repo.Queries = db.New(tx)
		return repo
	})
}
