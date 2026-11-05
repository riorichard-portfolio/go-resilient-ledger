package service

type Repo interface {
	BulkFindByIds(ids []string) (map[string]AccountMapVal, error)
	BulkAddBalance(bmap map[string]int) error
}

type Service struct {
	repo Repo
}

func NewService(
	repo Repo,
) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindValidAccs(accIds []string) (map[string]struct{}, error) {
	accountMap, err := s.repo.BulkFindByIds(accIds)
	if err != nil {
		return map[string]struct{}{}, err
	}
	result := make(map[string]struct{}, len(accountMap))
	for id := range accountMap {
		result[id] = struct{}{}
	}
	return result, nil
}

func (s *Service) AddAccsBalance(bmap map[string]int) error {
	return s.repo.BulkAddBalance(bmap)
}
