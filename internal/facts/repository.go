package facts

type FactsRepository interface {
	SaveFact(fact string) error
}
