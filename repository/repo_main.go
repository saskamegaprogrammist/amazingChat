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

func (repo *Repository)  createTables() error {
	_, err := repo.pool.Exec(`
DROP TABLE IF EXISTS pets;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS coords;
DROP TABLE IF EXISTS ads;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    firstname text NOT NULL ,
    lastname text NOT NULL,
    email text NOT NULL UNIQUE,
    nickname text NOT NULL UNIQUE,
    phone text UNIQUE,
    password text NOT NULL
    CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$')
);

CREATE INDEX users_nickname ON users (nickname);
CREATE INDEX users_email ON users (email);
CREATE INDEX users_email_password ON users (email, password);


CREATE TABLE ads (
    id SERIAL NOT NULL PRIMARY KEY,
    userid int NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    type int NOT NULL  CHECK (type = 0 OR type = 1),
    title text DEFAULT '',
    text text DEFAULT '',
    time text DEFAULT '',
    contacts text DEFAULT '',
    comments int DEFAULT 0,
    date TIMESTAMPTZ NOT NULL
);

CREATE INDEX ads_userid ON ads (userid);
CREATE INDEX ads_type ON ads (type);
CREATE INDEX ads_userid_type ON ads (userid, type);


CREATE TABLE pets (
    id SERIAL NOT NULL PRIMARY KEY,
    adid int NOT NULL REFERENCES ads(id) ON DELETE CASCADE,
    name text DEFAULT '',
    animal text DEFAULT '',
    breed text DEFAULT '',
    color text DEFAULT ''
);

CREATE INDEX pets_adid ON pets (adid);


CREATE TABLE coords (
    id SERIAL NOT NULL PRIMARY KEY,
    adid int NOT NULL REFERENCES ads(id) ON DELETE CASCADE,
    x double precision NOT NULL ,
    y double precision NOT NULL 
);

CREATE INDEX coords_adid ON coords (adid);


CREATE TABLE comments (
    id SERIAL NOT NULL PRIMARY KEY,
    adid int NOT NULL REFERENCES ads(id) ON DELETE CASCADE,
    userid int NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    text text NOT NULL ,
	date TIMESTAMPTZ NOT NULL
);

CREATE INDEX comments_adid ON coords (adid);


CREATE OR REPLACE FUNCTION ad_comment() RETURNS TRIGGER
LANGUAGE  plpgsql
AS $ad_comment$
BEGIN
   UPDATE ads SET comments = comments + 1 WHERE id = NEW.adid;
    RETURN NEW;
END
$ad_comment$;


CREATE TRIGGER AdComment
    AFTER INSERT on comments
    FOR EACH ROW
    EXECUTE PROCEDURE ad_comment();


CREATE OR REPLACE FUNCTION delete_comment() RETURNS TRIGGER
LANGUAGE  plpgsql
AS $delete_comment$
BEGIN
   UPDATE ads SET comments = comments - 1 WHERE id = OLD.adid;
    RETURN NEW;
END
$delete_comment$;


CREATE TRIGGER DeleteComment
    AFTER DELETE on comments
    FOR EACH ROW
    EXECUTE PROCEDURE delete_comment();



`)
	if err != nil {
		return err
	}
	return nil
}