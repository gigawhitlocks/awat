package model

import (
	cloudModel "github.com/mattermost/mattermost-cloud/model"
)

const TranslationStateRequested = "transaction-requested"
const TranslationStateInProgress = "transaction-in-progress"
const TranslationStateComplete = "transaction-complete"

type Translation struct {
	ID             string
	InstallationID string
	Team           string
	Type           string
	Resource       string
	Error          string
	StartAt        int64
	CompleteAt     int64
	LockedBy       string
}

func (t *Translation) State() string {
	if t.StartAt == 0 {
		return TranslationStateRequested
	}

	if t.CompleteAt == 0 {
		return TranslationStateInProgress
	}

	return TranslationStateComplete
}

func NewTranslationFromRequest(translationRequest *TranslationRequest) *Translation {
	return &Translation{
		ID:             cloudModel.NewID(),
		InstallationID: translationRequest.InstallationID,
		Type:           translationRequest.Type,
		Resource:       translationRequest.Archive,
		Team:           translationRequest.Team,
	}
}
