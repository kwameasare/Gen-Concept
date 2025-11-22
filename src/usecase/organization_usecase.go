package usecase

import (
	"context"
	"gen-concept-api/config"
	"gen-concept-api/domain/model"
	"gen-concept-api/domain/repository"
	"gen-concept-api/pkg/logging"
	"gen-concept-api/usecase/dto"

	"golang.org/x/crypto/bcrypt"
)

type OrganizationUsecase struct {
	logger                 logging.Logger
	cfg                    *config.Config
	organizationRepository repository.OrganizationRepository
	userRepository         repository.UserRepository
}

func NewOrganizationUsecase(cfg *config.Config, orgRepo repository.OrganizationRepository, userRepo repository.UserRepository) *OrganizationUsecase {
	logger := logging.NewLogger(cfg)
	return &OrganizationUsecase{
		cfg:                    cfg,
		organizationRepository: orgRepo,
		userRepository:         userRepo,
		logger:                 logger,
	}
}

func (u *OrganizationUsecase) OnboardOrganization(ctx context.Context, req dto.OnboardOrganizationRequest) (*model.Organization, *model.User, error) {
	// 1. Create Organization
	org := model.Organization{
		Name:             req.Organization.Name,
		Description:      req.Organization.Description,
		Domain:           req.Organization.Domain,
		SubscriptionPlan: req.Organization.SubscriptionPlan,
	}
	createdOrg, err := u.organizationRepository.Create(ctx, org)
	if err != nil {
		return nil, nil, err
	}

	// 2. Create Admin User
	user := dto.ToUserModel(req.User)
	user.OrganizationID = createdOrg.ID
	// Hash password
	bp := []byte(req.User.Password)
	hp, err := bcrypt.GenerateFromPassword(bp, bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return nil, nil, err
	}
	user.Password = string(hp)

	createdUser, err := u.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, nil, err
	}

	return &createdOrg, &createdUser, nil
}
