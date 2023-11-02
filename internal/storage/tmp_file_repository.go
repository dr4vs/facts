package storage

import (
	"math/rand"
	"os"
	"strings"
)

type TmpFileFactsRepository struct {
	tmpFile *os.File
}

func InitTmpFileFactsRepository() *TmpFileFactsRepository {
	fName, err := os.CreateTemp("", "tmp-facts-")
	if err != nil {
		return nil
	}
	return &TmpFileFactsRepository{
		tmpFile: fName,
	}
}

func (r *TmpFileFactsRepository) CloseTmpFileFactsRepository() {
	r.tmpFile.Close()
	os.Remove(r.tmpFile.Name())
}

func (r *TmpFileFactsRepository) SaveFact(fact string) error {
	if _, err := r.tmpFile.Write([]byte(fact + "\n")); err != nil {
		return err
	}
	return nil
}

func (r *TmpFileFactsRepository) GetFact() (string, error) {
	data, err := os.ReadFile(r.tmpFile.Name())
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(data), "\n")
	return lines[rand.Intn(len(lines))], nil
}
