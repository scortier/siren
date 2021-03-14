package rules

import (
	"context"
	cortexClient "github.com/grafana/cortex-tools/pkg/client"
	"github.com/grafana/cortex-tools/pkg/rules/rwrulefmt"
	"github.com/odpf/siren/domain"
	"gorm.io/gorm"
)

type cortexCaller interface {
	CreateRuleGroup(ctx context.Context, namespace string, rg rwrulefmt.RuleGroup) error
	DeleteRuleGroup(ctx context.Context, namespace, groupName string) error
	GetRuleGroup(ctx context.Context, namespace, groupName string) (*rwrulefmt.RuleGroup, error)
	ListRules(ctx context.Context, namespace string) (map[string][]rwrulefmt.RuleGroup, error)
}

// Service handles business logic
type Service struct {
	repository RuleRepository
	client     cortexCaller
}

// NewService returns repository struct
func NewService(db *gorm.DB, cortex domain.Cortex) domain.RuleService {
	cfg := cortexClient.Config{
		Address:         cortex.Host,
		UseLegacyRoutes: true,
	}
	client, err := cortexClient.New(cfg)
	if err != nil {
		return nil
	}
	return &Service{repository: NewRepository(db), client: client}
}

func (service Service) Migrate() error {
	return service.repository.Migrate()
}

func (service Service) Upsert(rule *domain.Rule) (*domain.Rule, error) {
	r := &Rule{}
	r, err := r.fromDomain(rule)
	upsertedRule, err := service.repository.Upsert(r, service.client)
	if err != nil {
		return nil, err
	}
	return upsertedRule.toDomain()
}

func (service Service) Get(namespace, entity, groupName, status, template string) ([]domain.Rule, error) {
	rules, err := service.repository.Get(namespace, entity, groupName, status, template)
	if err != nil {
		return nil, err
	}
	domainRules := make([]domain.Rule, 0, len(rules))
	for i := 0; i < len(rules); i++ {
		r, _ := rules[i].toDomain()
		domainRules = append(domainRules, *r)
	}
	return domainRules, nil
}
