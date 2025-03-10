package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres_user"
	password = "postgres_password"
	dbname   = "postgres"
)

func main() {
	var migrationsPath, migrationsTable, direction string

	// Получаем необходимые значения из флагов запуска

	// Путь до файла БД.
	// Его достаточно, т.к. мы используем SQLite, другие креды не нужны.
	//flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	// Путь до папки с миграциями.
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	// Таблица, в которой будет храниться информация о миграциях. Она нужна
	// для того, чтобы понимать, какие миграции уже применены, а какие нет.
	// Дефолтное значение - 'migrations'.
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.StringVar(&direction, "direction", "up", "path to migrations")
	flag.Parse() // Выполняем парсинг флагов

	// Валидация параметров
	//if storagePath == "" {
	//	// Простейший способ обработки ошибки :)
	//	// При необходимости, можете выбрать более подходящий вариант.
	//	// Меня паника пока устраивает, поскольку это вспомогательная утилита.
	//	panic("storage-path is required")
	//}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres", driver)
	// Создаем объект мигратора, передав креды нашей БД
	//m, err := migrate.New(
	//	"file://"+migrationsPath,
	//	fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	//)
	if err != nil {
		panic(err)
	}

	if direction == "up" {
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")

				return
			}

			panic(err)
		}
	} else if direction == "down" {
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")

				return
			}

			panic(err)
		}
	}
	// Выполняем миграции до последней версии

}
