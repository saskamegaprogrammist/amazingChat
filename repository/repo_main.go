package repository

import 	"github.com/jackc/pgx"

type Repository struct {
	pool *pgx.ConnPool
	UsersRepo *UsersRepo
	ChatsRepo *ChatsRepo
	MessagesRepo *MessagesRepo
}

var repo Repository
const dbConnections = 20


func Init(config pgx.ConnConfig) error {
	var err error
	repo.pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: config,
		MaxConnections: dbConnections,
	})
	if err != nil {
		return err
	}
	err = repo.createTables()
	if err != nil {
		return err
	}
	repo.UsersRepo = &UsersRepo{}
	repo.ChatsRepo = &ChatsRepo{}
	repo.MessagesRepo = &MessagesRepo{}
	return nil
}

// relation style tables creation

func (repo *Repository)  createTables() error {
	_, err := repo.pool.Exec(`
CREATE TABLE IF NOT EXISTS users  (
    id SERIAL NOT NULL PRIMARY KEY,
    username text NOT NULL UNIQUE,
    created TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS chats (
    id SERIAL NOT NULL PRIMARY KEY,
    name text NOT NULL UNIQUE,
	created TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS chat_users (
	chatid int NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    userid int NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL NOT NULL PRIMARY KEY,
    text text NOT NULL,
	chatid int NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    userid int NOT NULL REFERENCES users(id) ON DELETE SET NULL,
	created TIMESTAMPTZ NOT NULL
);

`)
	if err != nil {
		return err
	}
	return nil
}

func getPool() *pgx.ConnPool {
	return repo.pool
}

func GetUsersRepo() UsersRepoInterface {
	return repo.UsersRepo
}

func GetChatsRepo() ChatsRepoInterface {
	return repo.ChatsRepo
}

func GetMessagesRepo() MessagesRepoInterface {
	return repo.MessagesRepo
}