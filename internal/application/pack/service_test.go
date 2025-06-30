package pack

import (
	"errors"
	"github.com/h6x0r/pack-calculator/internal/application/pack/dto"
	"github.com/h6x0r/pack-calculator/internal/domain"
	"testing"
)

type mockRepo struct {
	packs     []domain.Pack
	listErr   error
	createErr error
	deleteErr error
	updateErr error
}

func (m *mockRepo) List() ([]domain.Pack, error) { return m.packs, m.listErr }
func (m *mockRepo) Create(size int) (domain.Pack, error) {
	if m.createErr != nil {
		return domain.Pack{}, m.createErr
	}
	return domain.Pack{ID: 1, Size: size}, nil
}
func (m *mockRepo) Delete(size int) error             { return m.deleteErr }
func (m *mockRepo) Update(oldSize, newSize int) error { return m.updateErr }

func TestServiceImpl_List_Success(t *testing.T) {
	repo := &mockRepo{packs: []domain.Pack{{ID: 1, Size: 5}}}
	svc := New(repo)
	resp, err := svc.List()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Packs) != 1 || resp.Packs[0].Size != 5 {
		t.Errorf("unexpected packs: %+v", resp.Packs)
	}
}

func TestServiceImpl_Add_Success(t *testing.T) {
	repo := &mockRepo{}
	svc := New(repo)
	resp, err := svc.Add(dto.PackAddRequest{Size: 10})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Size != 10 {
		t.Errorf("unexpected size: %d", resp.Size)
	}
}

func TestServiceImpl_Add_Error(t *testing.T) {
	repo := &mockRepo{createErr: errors.New("fail")}
	svc := New(repo)
	_, err := svc.Add(dto.PackAddRequest{Size: 10})
	if err == nil || err.Error() != "fail" {
		t.Errorf("expected error, got: %v", err)
	}
}

func TestServiceImpl_Remove_Error(t *testing.T) {
	repo := &mockRepo{deleteErr: errors.New("fail")}
	svc := New(repo)
	err := svc.Remove(dto.PackDeleteRequest{Size: 1})
	if err == nil || err.Error() != "fail" {
		t.Errorf("expected error, got: %v", err)
	}
}

func TestServiceImpl_Change_Error(t *testing.T) {
	repo := &mockRepo{updateErr: errors.New("fail")}
	svc := New(repo)
	err := svc.Change(dto.PackUpdateRequest{OldSize: 1, NewSize: 2})
	if err == nil || err.Error() != "fail" {
		t.Errorf("expected error, got: %v", err)
	}
}
