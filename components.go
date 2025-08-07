package statuspage

import (
	"context"
	"fmt"
	"net/http"
)

// ComponentsService handles all component-related API operations for tracking service status
type ComponentsService struct {
	client *Client
}

// ComponentRequest wraps component input data for API requests
type ComponentRequest struct {
	Component *ComponentInput `json:"component"`
}

// ComponentInput contains the editable fields for creating or updating a component
type ComponentInput struct {
	Name               string     `json:"name,omitempty"`
	Description        string     `json:"description,omitempty"`
	Status             string     `json:"status,omitempty"`
	OnlyShowIfDegraded *bool      `json:"only_show_if_degraded,omitempty"`
	GroupID            string     `json:"group_id,omitempty"`
	Showcase           *bool      `json:"showcase,omitempty"`
	StartDate          *Time `json:"start_date,omitempty"`
}

// ComponentStatusInput is used specifically for updating only the status of a component
type ComponentStatusInput struct {
	Component struct {
		Status string `json:"status"`
	} `json:"component"`
}

// Component status constants for updating component states
const (
	ComponentStatusOperational         = "operational"
	ComponentStatusDegradedPerformance = "degraded_performance"
	ComponentStatusPartialOutage       = "partial_outage"
	ComponentStatusMajorOutage         = "major_outage"
	ComponentStatusUnderMaintenance    = "under_maintenance"
)

// List retrieves all components for a specific status page
func (s *ComponentsService) List(ctx context.Context, pageID string) ([]*Component, error) {
	u := fmt.Sprintf("pages/%s/components", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var components []*Component
	_, err = s.client.Do(ctx, req, &components)
	if err != nil {
		return nil, err
	}

	return components, nil
}

// Get retrieves a specific component by its unique identifier
func (s *ComponentsService) Get(ctx context.Context, pageID, componentID string) (*Component, error) {
	u := fmt.Sprintf("pages/%s/components/%s", pageID, componentID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	component := new(Component)
	_, err = s.client.Do(ctx, req, component)
	if err != nil {
		return nil, err
	}

	return component, nil
}

// Create adds a new component to track on the status page
func (s *ComponentsService) Create(ctx context.Context, pageID string, component *ComponentInput) (*Component, error) {
	u := fmt.Sprintf("pages/%s/components", pageID)
	componentReq := &ComponentRequest{Component: component}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, componentReq)
	if err != nil {
		return nil, err
	}

	newComponent := new(Component)
	_, err = s.client.Do(ctx, req, newComponent)
	if err != nil {
		return nil, err
	}

	return newComponent, nil
}

// Update modifies an existing component's configuration
func (s *ComponentsService) Update(ctx context.Context, pageID, componentID string, component *ComponentInput) (*Component, error) {
	u := fmt.Sprintf("pages/%s/components/%s", pageID, componentID)
	componentReq := &ComponentRequest{Component: component}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, componentReq)
	if err != nil {
		return nil, err
	}

	updatedComponent := new(Component)
	_, err = s.client.Do(ctx, req, updatedComponent)
	if err != nil {
		return nil, err
	}

	return updatedComponent, nil
}

// Delete removes a component from the status page
func (s *ComponentsService) Delete(ctx context.Context, pageID, componentID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/components/%s", pageID, componentID)
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

// UpdateStatus changes only the operational status of a component
func (s *ComponentsService) UpdateStatus(ctx context.Context, pageID, componentID, status string) (*Component, error) {
	u := fmt.Sprintf("pages/%s/components/%s", pageID, componentID)
	statusInput := &ComponentStatusInput{}
	statusInput.Component.Status = status

	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, statusInput)
	if err != nil {
		return nil, err
	}

	component := new(Component)
	_, err = s.client.Do(ctx, req, component)
	if err != nil {
		return nil, err
	}

	return component, nil
}
