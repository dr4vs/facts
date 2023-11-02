package facts

type FactsService struct {
	queryServcie FactsQueryService
	repository   FactsRepository
}

func InitFactService(queryServcie FactsQueryService, repository FactsRepository) *FactsService {
	return &FactsService{
		queryServcie: queryServcie,
		repository:   repository,
	}
}

func (s *FactsService) GetRandomFact() (string, error) {
	fact, err := s.queryServcie.GetFact()
	if err != nil {
		return "", err
	}
	return fact, nil
}
