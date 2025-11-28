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

							// Map BackendJourney steps if necessary (omitted for brevity, can be added if steps have IDs)
							break
						}
					}
				}
				break
			}
		}
	}
}
