package calc

import (
	"errors"
	"github.com/h6x0r/pack-calculator/internal/application/calc/dto"
	"github.com/h6x0r/pack-calculator/internal/domain"
	"reflect"
	"testing"
)

type mockPackRepo struct {
	packs   []domain.Pack
	listErr error
}

func (m *mockPackRepo) List() ([]domain.Pack, error) {
	return m.packs, m.listErr
}
func (m *mockPackRepo) Create(size int) (domain.Pack, error) { return domain.Pack{}, nil }
func (m *mockPackRepo) Delete(size int) error                { return nil }
func (m *mockPackRepo) Update(oldSize, newSize int) error    { return nil }

type mockOrderRepo struct {
	saveErr error
}

func (m *mockOrderRepo) Save(o *domain.Order) error { return m.saveErr }

func TestServiceImpl_Calculate_Success(t *testing.T) {
	packs := []domain.Pack{{ID: 1, Size: 5}, {ID: 2, Size: 10}}
	service := New(&mockPackRepo{packs: packs}, &mockOrderRepo{})
	resp, err := service.Calculate(dto.CalculateRequest{Items: 12})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Total < 12 {
		t.Errorf("total should be >= items")
	}
}

func TestServiceImpl_Calculate_NoPacks(t *testing.T) {
	service := New(&mockPackRepo{packs: nil}, &mockOrderRepo{})
	_, err := service.Calculate(dto.CalculateRequest{Items: 5})
	if err == nil || err.Error() != "no pack sizes configured" {
		t.Errorf("expected error for no packs, got: %v", err)
	}
}

func TestServiceImpl_Calculate_SaveError(t *testing.T) {
	packs := []domain.Pack{{ID: 1, Size: 5}}
	service := New(&mockPackRepo{packs: packs}, &mockOrderRepo{saveErr: errors.New("db error")})
	_, err := service.Calculate(dto.CalculateRequest{Items: 5})
	if err == nil || err.Error() != "db error" {
		t.Errorf("expected db error, got: %v", err)
	}
}

func TestServiceImpl_Calculate_InvalidItems(t *testing.T) {
	service := New(&mockPackRepo{packs: []domain.Pack{{ID: 1, Size: 5}}}, &mockOrderRepo{})
	_, err := service.Calculate(dto.CalculateRequest{Items: -1})
	if err == nil || err.Error() != "items must be ≥0" {
		t.Errorf("expected error for negative items, got: %v", err)
	}
}

func TestCalculate_CoreCases(t *testing.T) {
	cases := []struct {
		name      string
		sizes     []int
		order     int
		wantPacks map[int]int
		wantOver  int
	}{
		{"exact 1×250", []int{250, 500}, 250, map[int]int{250: 1}, 0},
		{"251→1×500", []int{250, 500}, 251, map[int]int{500: 1}, 249},
		{"501→500+250", []int{250, 500, 1000}, 501, map[int]int{500: 1, 250: 1}, 249},
		{"big edge 500k", []int{23, 31, 53}, 500000, map[int]int{23: 2, 31: 7, 53: 9429}, 0},
		{"prime order 97", []int{10, 20, 25}, 97, map[int]int{25: 4}, 3},

		{"zero order", []int{250, 500}, 0, map[int]int{}, 0},

		{"below min size", []int{6, 9}, 1, map[int]int{6: 1}, 5},

		{"tie overshoot choose fewer packs",
			[]int{4, 6, 8}, 14,
			map[int]int{8: 1, 6: 1}, 0},

		{"large non-exact", []int{500, 700}, 999, map[int]int{500: 2}, 1},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calculate(tc.order, tc.sizes)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got.Packs, tc.wantPacks) {
				t.Fatalf("packs mismatch\nwant: %v\ngot : %v", tc.wantPacks, got.Packs)
			}
			if got.Overshoot != tc.wantOver {
				t.Fatalf("overshoot: want %d, got %d", tc.wantOver, got.Overshoot)
			}
		})
	}
}

func TestCalculate_Invalid(t *testing.T) {
	tests := []struct {
		order int
		sizes []int
	}{
		{-1, []int{10}},
		{10, nil},
		{10, []int{}},
	}

	for _, tt := range tests {
		if _, err := Calculate(tt.order, tt.sizes); err == nil {
			t.Fatalf("expected error for order=%d sizes=%v", tt.order, tt.sizes)
		}
	}
}
