package statuspage

import (
	"context"
	"fmt"
	"net/http"
)

type IncidentUpdatesService struct {
	client *Client
}

type IncidentUpdateRequest struct {
	IncidentUpdate *IncidentUpdateInput `json:"incident_update"`
}

type IncidentUpdateInput struct {
	Body                 string            `json:"body,omitempty"`
	Status               string            `json:"status,omitempty"`
	DeliverNotifications *bool             `json:"deliver_notifications,omitempty"`
	CustomTweet          string            `json:"custom_tweet,omitempty"`
	TweetID              string            `json:"tweet_id,omitempty"`
	Components           map[string]string `json:"components,omitempty"`
	AffectedComponents   []string          `json:"affected_components,omitempty"`
}

func (s *IncidentUpdatesService) List(ctx context.Context, pageID, incidentID string) ([]*IncidentUpdate, error) {
	u := fmt.Sprintf("pages/%s/incidents/%s/incident_updates", pageID, incidentID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var updates []*IncidentUpdate
	_, err = s.client.Do(ctx, req, &updates)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

func (s *IncidentUpdatesService) Get(ctx context.Context, pageID, incidentID, updateID string) (*IncidentUpdate, error) {
	u := fmt.Sprintf("pages/%s/incidents/%s/incident_updates/%s", pageID, incidentID, updateID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	update := new(IncidentUpdate)
	_, err = s.client.Do(ctx, req, update)
	if err != nil {
		return nil, err
	}

	return update, nil
}

func (s *IncidentUpdatesService) Create(ctx context.Context, pageID, incidentID string, update *IncidentUpdateInput) (*IncidentUpdate, error) {
	u := fmt.Sprintf("pages/%s/incidents/%s/incident_updates", pageID, incidentID)
	updateReq := &IncidentUpdateRequest{IncidentUpdate: update}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, updateReq)
	if err != nil {
		return nil, err
	}

	newUpdate := new(IncidentUpdate)
	_, err = s.client.Do(ctx, req, newUpdate)
	if err != nil {
		return nil, err
	}

	return newUpdate, nil
}

func (s *IncidentUpdatesService) Update(ctx context.Context, pageID, incidentID, updateID string, update *IncidentUpdateInput) (*IncidentUpdate, error) {
	u := fmt.Sprintf("pages/%s/incidents/%s/incident_updates/%s", pageID, incidentID, updateID)
	updateReq := &IncidentUpdateRequest{IncidentUpdate: update}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, updateReq)
	if err != nil {
		return nil, err
	}

	updatedUpdate := new(IncidentUpdate)
	_, err = s.client.Do(ctx, req, updatedUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUpdate, nil
}
