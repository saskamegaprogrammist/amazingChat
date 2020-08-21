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
	//err = repo.createTables()
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
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS chat_users;
DROP TABLE IF EXISTS chats;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    username text NOT NULL UNIQUE,
    created TIMESTAMPTZ NOT NULL
);

CREATE TABLE chats (
    id SERIAL NOT NULL PRIMARY KEY,
    name text NOT NULL UNIQUE,
	created TIMESTAMPTZ NOT NULL
);

CREATE TABLE chat_users (
	chatid int NOT NULL REFERENCES chats(id) ON DELETE CASCADE,
    userid int NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX chat_users_chatid ON chat_users (chatid);

CREATE TABLE messages (
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

func GetUsersRepo() *UsersRepo {
	return repo.UsersRepo
}

func GetChatsRepo() *ChatsRepo {
	return repo.ChatsRepo
}

func GetMessagesRepo() *MessagesRepo {
	return repo.MessagesRepo
}