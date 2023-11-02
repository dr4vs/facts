package storage

import (
  "database/sql"
)

type PostgresqlFactsRepository struct{
  db *sql.DB
}

func InitPostgresqlFactsRepository(db *sql.DB) *PostgresqlFactsRepository{
  return &PostgresqlFactsRepository{
    db: db,
  }
}

func (r *PostgresqlFactsRepository) SaveFact(fact string) error {
  if _, err := r.db.Exec("INSERT INTO facts(fact) VALUES($1)", fact); err != nil {
     return err
  }
  return nil
}
