package pack

import "pc/internal/domain"

type Service struct{ repo domain.PackRepository }

func New(r domain.PackRepository) *Service { return &Service{r} }

func (s *Service) List() ([]domain.Pack, error)      { return s.repo.List() }
func (s *Service) Add(size int) (domain.Pack, error) { return s.repo.Create(size) }
func (s *Service) Remove(size int) error             { return s.repo.Delete(size) }
func (s *Service) Change(oldSize, newSize int) error { return s.repo.Update(oldSize, newSize) }
