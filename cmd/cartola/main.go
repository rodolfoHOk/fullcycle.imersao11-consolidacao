package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/infra/db"
	httpHandler "github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/infra/http"
	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/internal/infra/repository"
	"github.com/rodolfoHOk/fullcycle.imersao11-consolidacao/pkg/uow"

	"github.com/go-chi/chi"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dtb, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/cartola?parseTime=true")
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
	
	if err = http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
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
