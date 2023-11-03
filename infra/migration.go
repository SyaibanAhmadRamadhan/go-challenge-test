package infra

import (
	"bufio"
	"database/sql"
	"embed"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/rs/zerolog/log"

	"challenge-test-synapsis/conf"
	"challenge-test-synapsis/helper"
)

//go:embed migration
var embedMigration embed.FS

func ConnForMigrate(url string) *sql.DB {
	db, err := goose.OpenDBWithDriver("pgx", url)
	if err != nil {
		helper.PanicIf(err)
	}

	return db
}

func MigrateMaster(cmd string, arg string, url string) {
	goose.SetSequential(true)
	if url == "" {
		postgresConf := conf.EnvPostgresConf()
		url = postgresConf.DBUrl()
	}
	if cmd == "" && arg == "" {
		cmd, arg = cmdMigrateMaster()
	}

	goose.SetBaseFS(embedMigration)

	db := ConnForMigrate(url)
	defer func() {
		if err := db.Close(); err != nil {
			log.Err(err).Msgf("failed close goose db")
		}
		log.Info().Msgf("closed db for migration")
	}()

	var err error

	switch cmd {
	case "up":
		err = goose.Up(db, "migration")
	case "up-by-one":
		err = goose.UpByOne(db, "migration")
	case "status":
		err = goose.Status(db, "migration")
	case "down":
		err = goose.Down(db, "migration")
	case "redo":
		err = goose.Redo(db, "migration")
	case "reset":
		err = goose.Reset(db, "migration")
	case "version":
		err = goose.Version(db, "migration")
	case "fix":
		err = goose.Fix("migration")
	case "create":
		err = goose.Create(db, "infra/"+"migration", arg, "sql")
	case "create-go":
		err = goose.Create(db, "infra/"+"migration", arg, "go")
	case "up-to":
		argInt, err := strconv.Atoi(arg)
		helper.PanicIf(err)
		err = goose.UpTo(db, "migration", int64(argInt))
	case "down-to":
		argInt, err := strconv.Atoi(arg)
		helper.PanicIf(err)
		err = goose.DownTo(db, "migration", int64(argInt))
	default:
		err = goose.Status(db, "migration")
	}
	helper.PanicIf(err)
}

func cmdMigrateMaster() (cmd string, arg string) {
	fmt.Println("\nSelect command migrate:")
	fmt.Println("up\t\t\tMigrate the DB to the most recent arg available")
	fmt.Println("up-by-one\t\tMigrate the DB up by 1")
	fmt.Println("up-to VERSION\t\tMigrate the DB to a specific VERSION")
	fmt.Println("down\t\t\tRoll back the migration by 1. desc migration priority")
	fmt.Println("down-to VERSION\t\tRoll back the DB to a specific VERSION")
	fmt.Println("redo\t\t\tRe-run the latest migration")
	fmt.Println("reset\t\t\tRoll back all migrations")
	fmt.Println("status\t\t\tDump the migration status for the current DB")
	fmt.Println("version\t\t\tPrint the current version of the database")
	fmt.Println("create NAME\t\tCreates new migration file with the increment current migration")
	fmt.Println("create-go NAME\t\tCreates new migration file go with the increment current migration")
	fmt.Print("Masukkan perintah yang ingin dijalankan any key for stop: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Split(input, " ")
	cmd = parts[0]

	switch cmd {
	case "status", "up", "down", "reset", "up-by-one", "redo", "version":
		fmt.Println("\n=== your command valid. start migration ===")
		return cmd, arg
	case "create", "create-go", "up-to", "down-to":
		if len(parts) < 2 {
			panic("argument 2 is empty")
		}
		arg = parts[1]
		log.Info().Msgf("%v", arg)
		return cmd, arg
	default:
		panic("invalid command")
	}
}
