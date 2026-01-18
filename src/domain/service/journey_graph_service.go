package service

import (
	"gen-concept-api/domain/model"
	"gen-concept-api/infra/persistence/database"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JourneyGraphService struct {
	db *gorm.DB
}

func NewJourneyGraphService() *JourneyGraphService {
	return &JourneyGraphService{
		db: database.GetDb(),
	}
}

func (s *JourneyGraphService) GetGraph(journeyID uuid.UUID, level string, parentID *uuid.UUID) ([]model.JourneyNode, []model.JourneyEdge, error) {
	var nodes []model.JourneyNode
	var edges []model.JourneyEdge

	// 1. Fetch Nodes
	nodeQuery := s.db.Where("journey_id = ?", journeyID)

	if level != "" {
		nodeQuery = nodeQuery.Where("level = ?", level)
	}

	if parentID != nil {
		nodeQuery = nodeQuery.Where("parent_node_id = ?", parentID)
	} else {
		// If explicit nil passed for root nodes (optional logic depending on requirements)
		// For now we assume if parentID is nil we check 'IS NULL' only if strictly requested,
		// but usually 'nil' in Go arg means "ignore filter" or "root".
		// The prompt says: "If parentID is provided, filter by parent_node_id." called with a pointer.
		// We'll interpret passing a pointer means filtering is requested.
	}

	// Refined logic: If parentID is passed (non-nil), strictly filter.
	// If the user wants ROOTS, they might pass a pointer to a nil UUID?
	// For now let's stick to: if parentID != nil in the arg, we use it.
	// Note: GORM handles pointers well.

	if err := nodeQuery.Find(&nodes).Error; err != nil {
		return nil, nil, err
	}

	// Optimization: Extract Node IDs to filter edges
	nodeIDs := make([]uuid.UUID, len(nodes))
	for i, n := range nodes {
		nodeIDs[i] = n.Uuid
	}

	// 2. Fetch Edges
	// Only return edges where BOTH Source and Target are in the fetched nodes (dangling edge prevention)
	if len(nodeIDs) > 0 {
		err := s.db.Where("journey_id = ? AND source_id IN ? AND target_id IN ?", journeyID, nodeIDs, nodeIDs).
			Find(&edges).Error
		if err != nil {
			return nil, nil, err
		}
	}

	return nodes, edges, nil
}
