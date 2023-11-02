package storage

import (
	"math/rand"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type TmpDirFactsRepository struct {
	tmpDirName string
}

func InitTmpDirFactsRepository() *TmpDirFactsRepository {
	dirName, err := os.MkdirTemp("", "facts_tmp")
	if err != nil {
		return nil
	}
	return &TmpDirFactsRepository{
		tmpDirName: dirName,
	}
}

func (r *TmpDirFactsRepository) CloseTmpDirFactsRepository() {
	os.RemoveAll(r.tmpDirName)
}

func (r *TmpDirFactsRepository) SaveFact(fact string) error {
	fileName := filepath.Join(r.tmpDirName, uuid.New().String())
	if err := os.WriteFile(fileName, []byte(fact), 0777); err != nil {
		return err
	}
	return nil
}

func (r *TmpDirFactsRepository) GetFact() (string, error) {
	dirEntry, err := os.ReadDir(r.tmpDirName)
	if err != nil {
		return "", err
	}
	randomFile := r.tmpDirName + "/" + dirEntry[rand.Intn(len(dirEntry))].Name()
	b, err := os.ReadFile(randomFile)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
