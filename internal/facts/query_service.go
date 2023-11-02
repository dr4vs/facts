package facts

type FactsQueryService interface{
  GetFact() (string, error)
}
