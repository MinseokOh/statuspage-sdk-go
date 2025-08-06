package statuspage

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// IncidentsService handles all incident-related API operations for managing service disruptions and maintenance
type IncidentsService struct {
	client *Client
}

// IncidentRequest wraps incident input data for API requests
type IncidentRequest struct {
	Incident *IncidentInput `json:"incident"`
}

// IncidentInput contains the editable fields for creating or updating an incident
type IncidentInput struct {
	Name                                      string                 `json:"name,omitempty"`
	Status                                    string                 `json:"status,omitempty"`
	ImpactOverride                            string                 `json:"impact_override,omitempty"`
	ScheduledFor                              *time.Time             `json:"scheduled_for,omitempty"`
	ScheduledUntil                            *time.Time             `json:"scheduled_until,omitempty"`
	ScheduledRemindPrior                      *bool                  `json:"scheduled_remind_prior,omitempty"`
	ScheduledAutoInProgress                   *bool                  `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledAutoCompleted                    *bool                  `json:"scheduled_auto_completed,omitempty"`
	Body                                      string                 `json:"body,omitempty"`
	ComponentIDs                              []string               `json:"component_ids,omitempty"`
	Components                                map[string]string      `json:"components,omitempty"`
	DeliverNotifications                      *bool                  `json:"deliver_notifications,omitempty"`
	AutoTransitionDeliverNotificationsAtEnd   *bool                  `json:"auto_transition_deliver_notifications_at_end,omitempty"`
	AutoTransitionDeliverNotificationsAtStart *bool                  `json:"auto_transition_deliver_notifications_at_start,omitempty"`
	AutoTransitionToMaintenanceState          *bool                  `json:"auto_transition_to_maintenance_state,omitempty"`
	AutoTransitionToOperationalState          *bool                  `json:"auto_transition_to_operational_state,omitempty"`
	AutoTweetAtBeginning                      *bool                  `json:"auto_tweet_at_beginning,omitempty"`
	AutoTweetOnCompletion                     *bool                  `json:"auto_tweet_on_completion,omitempty"`
	AutoTweetOnCreation                       *bool                  `json:"auto_tweet_on_creation,omitempty"`
	AutoTweetOneHourBefore                    *bool                  `json:"auto_tweet_one_hour_before,omitempty"`
	BackfillDate                              string                 `json:"backfill_date,omitempty"`
	Backfilled                                *bool                  `json:"backfilled,omitempty"`
	Metadata                                  map[string]interface{} `json:"metadata,omitempty"`
}

// IncidentListOptions provides filtering and pagination options for incident queries
type IncidentListOptions struct {
	Q       string `url:"q,omitempty"`
	Impact  string `url:"impact,omitempty"`
	Status  string `url:"status,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
}

// Incident status and impact level constants for managing incident lifecycle
const (
	IncidentStatusInvestigating = "investigating"
	IncidentStatusIdentified    = "identified"
	IncidentStatusMonitoring    = "monitoring"
	IncidentStatusResolved      = "resolved"
	IncidentStatusScheduled     = "scheduled"
	IncidentStatusInProgress    = "in_progress"
	IncidentStatusVerifying     = "verifying"
	IncidentStatusCompleted     = "completed"

	IncidentImpactNone        = "none"
	IncidentImpactMinor       = "minor"
	IncidentImpactMajor       = "major"
	IncidentImpactCritical    = "critical"
	IncidentImpactMaintenance = "maintenance"
)

// List retrieves incidents for a status page with optional filtering and pagination
func (s *IncidentsService) List(ctx context.Context, pageID string, opts *IncidentListOptions) ([]*Incident, *Response, error) {
	u := fmt.Sprintf("pages/%s/incidents", pageID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var incidents []*Incident
	resp, err := s.client.Do(ctx, req, &incidents)
	if err != nil {
		return nil, resp, err
	}

	return incidents, resp, nil
}

// ListUnresolved retrieves all active incidents that have not been resolved
func (s *IncidentsService) ListUnresolved(ctx context.Context, pageID string) ([]*Incident, *Response, error) {
	u := fmt.Sprintf("pages/%s/incidents/unresolved", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var incidents []*Incident
	resp, err := s.client.Do(ctx, req, &incidents)
	if err != nil {
		return nil, resp, err
	}

	return incidents, resp, nil
}

// ListScheduled retrieves all scheduled maintenance incidents for future events
func (s *IncidentsService) ListScheduled(ctx context.Context, pageID string) ([]*Incident, *Response, error) {
	u := fmt.Sprintf("pages/%s/incidents/scheduled", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var incidents []*Incident
	resp, err := s.client.Do(ctx, req, &incidents)
	if err != nil {
		return nil, resp, err
	}

	return incidents, resp, nil
}

func (s *IncidentsService) Get(ctx context.Context, pageID, incidentID string) (*Incident, *Response, error) {
	u := fmt.Sprintf("pages/%s/incidents/%s", pageID, incidentID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	incident := new(Incident)
	resp, err := s.client.Do(ctx, req, incident)
	if err != nil {
		return nil, resp, err
	}

	return incident, resp, nil
}

func (s *IncidentsService) Create(ctx context.Context, pageID string, incident *IncidentInput) (*Incident, *Response, error) {
	u := fmt.Sprintf("pages/%s/incidents", pageID)
	incidentReq := &IncidentRequest{Incident: incident}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, incidentReq)
	if err != nil {
		return nil, nil, err
	}

	newIncident := new(Incident)
	resp, err := s.client.Do(ctx, req, newIncident)
	if err != nil {
		return nil, resp, err
	}

	return newIncident, resp, nil
}

func (s *IncidentsService) Update(ctx context.Context, pageID, incidentID string, incident *IncidentInput) (*Incident, *Response, error) {
	u := fmt.Sprintf("pages/%s/incidents/%s", pageID, incidentID)
	incidentReq := &IncidentRequest{Incident: incident}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, incidentReq)
	if err != nil {
		return nil, nil, err
	}

	updatedIncident := new(Incident)
	resp, err := s.client.Do(ctx, req, updatedIncident)
	if err != nil {
		return nil, resp, err
	}

	return updatedIncident, resp, nil
}

func (s *IncidentsService) Delete(ctx context.Context, pageID, incidentID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/incidents/%s", pageID, incidentID)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
