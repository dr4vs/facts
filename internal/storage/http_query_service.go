package storage

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type HttpFactsQueryService struct{}

type FactDTO struct{
  Id string `json:"id"`
  Text string `json:"text"`
}

func (*HttpFactsQueryService) GetFact() (string, error ){
	resp, err := http.Get("https://uselessfacts.jsph.pl/api/v2/facts/random?language=en")
  if err != nil{
   log.Fatal(err)
   return "",err 
  }

  defer resp.Body.Close()
  
  body, err := io.ReadAll(resp.Body)
  if err != nil{
   log.Fatal(err)
   return "",err 
  }

  var fact FactDTO
  err = json.Unmarshal(body, &fact)
  if err != nil {
   log.Fatal(err)
   return "",err 
  }

	return fact.Text, nil
}
