package usecase

import (
	"gen-concept-api/domain/model"
)

func mapJourneyIDs(existing *model.Journey, incoming *model.Journey) {
	// Map EntityJourneys
	for i := range incoming.EntityJourneys {
		incomingEJ := &incoming.EntityJourneys[i]
		for _, existingEJ := range existing.EntityJourneys {
			if incomingEJ.Uuid == existingEJ.Uuid {
				incomingEJ.ID = existingEJ.ID
				incomingEJ.JourneyID = existing.ID
				incomingEJ.CreatedAt = existingEJ.CreatedAt
				incomingEJ.CreatedBy = existingEJ.CreatedBy

				// Map Operations within EntityJourney
				for j := range incomingEJ.Operations {
					incomingOp := &incomingEJ.Operations[j]
					for _, existingOp := range existingEJ.Operations {
						if incomingOp.Uuid == existingOp.Uuid {
							incomingOp.ID = existingOp.ID
							incomingOp.EntityJourneyID = existingEJ.ID
							incomingOp.CreatedAt = existingOp.CreatedAt
							incomingOp.CreatedBy = existingOp.CreatedBy

							// Map BackendJourney steps
							mapJourneySteps(existingOp.BackendJourney, incomingOp.BackendJourney)
							break
						}
					}
				}
				break
			}
		}
	}
}

func mapJourneySteps(existingSteps []model.JourneyStep, incomingSteps []model.JourneyStep) {
	for i := range incomingSteps {
		incomingStep := &incomingSteps[i]
		for _, existingStep := range existingSteps {
			if incomingStep.Uuid == existingStep.Uuid {
				incomingStep.ID = existingStep.ID
				incomingStep.OperationID = existingStep.OperationID
				incomingStep.ParentStepID = existingStep.ParentStepID
				incomingStep.CreatedAt = existingStep.CreatedAt
				incomingStep.CreatedBy = existingStep.CreatedBy

				// Recursively map SubSteps
				if len(incomingStep.SubSteps) > 0 && len(existingStep.SubSteps) > 0 {
					mapJourneySteps(existingStep.SubSteps, incomingStep.SubSteps)
				}
				break
			}
		}
	}
}
