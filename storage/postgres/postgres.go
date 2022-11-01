package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `

CREATE TABLE IF NOT EXISTS author (
    id	CHAR(36) PRIMARY KEY,
	first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS article (
    id CHAR(36) PRIMARY KEY,
	title VARCHAR(255) UNIQUE NOT NULL,
	body text NOT NULL,
	author_id CHAR(36),
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP,
	deleted_at TIMESTAMP
);

ALTER TABLE article DROP CONSTRAINT IF EXISTS fk_article_author; 
ALTER TABLE article ADD CONSTRAINT fk_article_author FOREIGN KEY (author_id) REFERENCES author (id);
`
type Postgres struct{
	db *sqlx.DB
}

func InitDb(psqlConfig string) (*Postgres, error){
	var err error

	tempDb, err := sqlx.Connect("postgres", psqlConfig)
	if err != nil{
		return nil, err
	}

	tempDb.MustExec(schema)

	tx := tempDb.MustBegin()
	tx.MustExec("INSERT INTO author (id, first_name, last_name) VALUES ($1, $2, $3)ON CONFLICT DO NOTHING", "f4edbadf-6153-4d31-8955-918af3f967a4", "John", "Doe")
	tx.MustExec("INSERT INTO author (id, first_name, last_name) VALUES ($1, $2, $3)ON CONFLICT DO NOTHING", "349b2748-3480-4c33-ac67-6e44d23555fe", "Peter", "Parker")
	tx.Commit()

	_, err = tempDb.NamedExec(`INSERT INTO article (id, title, body, author_id) VALUES (:id, :title, :body, :author_id)ON CONFLICT DO NOTHING`,
		map[string]interface{}{
			"id":        "1b569b84-48d9-414f-8882-b265c1dec5cc",
			"title":     "Lorem",
			"body":      "Ipsume lorem something...",
			"author_id": "f4edbadf-6153-4d31-8955-918af3f967a4",
		})
	if err != nil {
		panic(err)
	}

	return &Postgres{
		db: tempDb,
	}, nil
}
