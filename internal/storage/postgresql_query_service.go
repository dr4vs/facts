package storage

import (
	"database/sql"
	"math/rand"
)

type PostgresqlFactsQueryService struct{
  db *sql.DB
}

func InitPostgresqlFactsQueryService(db *sql.DB) *PostgresqlFactsQueryService{
  return &PostgresqlFactsQueryService{
    db: db,
  }
}

func (qs *PostgresqlFactsQueryService) GetFact() (string, error){
  rows, err := qs.db.Query("SELECT fact FROM facts")
  if err != nil {
    return "", err
  }
  defer rows.Close()

  var facts []string
  for rows.Next(){
    var fact string
    if err := rows.Scan(&fact); err != nil{
      return "", err
    }
    facts = append(facts, fact)
  }
  if rows.Err(); err != nil{
    return "", err
  }
  return facts[rand.Intn(len(facts))], nil
}
