package unit

import (
	"context"
	"testing"

	"gen-concept-api/config"
	"gen-concept-api/domain/filter"
	"gen-concept-api/domain/model"
	"gen-concept-api/usecase"
	"gen-concept-api/usecase/dto"

	"github.com/google/uuid"
)

// Mock Organization Repository
type MockOrganizationRepository struct{}

func (m *MockOrganizationRepository) Create(ctx context.Context, entity model.Organization) (model.Organization, error) {
	entity.ID = 1
	return entity, nil
}
func (m *MockOrganizationRepository) Update(ctx context.Context, uuid uuid.UUID, entity map[string]interface{}) (model.Organization, error) {
	return model.Organization{}, nil
}
func (m *MockOrganizationRepository) Delete(ctx context.Context, uuid uuid.UUID) error { return nil }
func (m *MockOrganizationRepository) GetById(ctx context.Context, uuid uuid.UUID) (model.Organization, error) {
	return model.Organization{}, nil
}
func (m *MockOrganizationRepository) GetByFilter(ctx context.Context, req filter.PaginationInputWithFilter) (int64, *[]model.Organization, error) {
	return 0, nil, nil
}

// Mock User Repository
type MockUserRepository struct{}

func (m *MockUserRepository) ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error) {
	return false, nil
}
func (m *MockUserRepository) ExistsUsername(ctx context.Context, username string) (bool, error) {
	return false, nil
}
func (m *MockUserRepository) ExistsEmail(ctx context.Context, email string) (bool, error) {
	return false, nil
}
func (m *MockUserRepository) FetchUserInfo(ctx context.Context, username string, password string) (model.User, error) {
	return model.User{}, nil
}
func (m *MockUserRepository) GetDefaultRole(ctx context.Context) (roleId uint, err error) {
	return 1, nil
}
func (m *MockUserRepository) CreateUser(ctx context.Context, u model.User) (model.User, error) {
	u.ID = 1
	return u, nil
}

func TestOnboardOrganization(t *testing.T) {
	cfg := &config.Config{
		Logger: config.LoggerConfig{
			Level:    "debug",
			Encoding: "console",
			Logger:   "zap",
		},
	}
	orgRepo := &MockOrganizationRepository{}
	userRepo := &MockUserRepository{}
	uc := usecase.NewOrganizationUsecase(cfg, orgRepo, userRepo)

	req := dto.OnboardOrganizationRequest{
		Organization: dto.OrganizationDto{
			Name:   "Test Org",
			Domain: "test.com",
		},
		User: dto.RegisterUserByUsername{
			Username: "admin",
			Email:    "admin@test.com",
			Password: "password",
		},
	}

	org, user, err := uc.OnboardOrganization(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if org.Name != "Test Org" {
		t.Errorf("Expected org name 'Test Org', got %s", org.Name)
	}

	if user.Username != "admin" {
		t.Errorf("Expected username 'admin', got %s", user.Username)
	}

	if user.OrganizationID != org.ID {
		t.Errorf("Expected user to be linked to org %d, got %d", org.ID, user.OrganizationID)
	}
}
