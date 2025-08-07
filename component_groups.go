package statuspage

import (
	"context"
	"fmt"
	"net/http"
)

type ComponentGroupsService struct {
	client *Client
}

type ComponentGroupRequest struct {
	ComponentGroup *ComponentGroupInput `json:"component_group"`
}

type ComponentGroupInput struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Components  []string `json:"components,omitempty"`
	Position    int      `json:"position,omitempty"`
}

func (s *ComponentGroupsService) List(ctx context.Context, pageID string) ([]*ComponentGroup, error) {
	u := fmt.Sprintf("pages/%s/component-groups", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var groups []*ComponentGroup
	_, err = s.client.Do(ctx, req, &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (s *ComponentGroupsService) Get(ctx context.Context, pageID, groupID string) (*ComponentGroup, error) {
	u := fmt.Sprintf("pages/%s/component-groups/%s", pageID, groupID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	group := new(ComponentGroup)
	_, err = s.client.Do(ctx, req, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (s *ComponentGroupsService) Create(ctx context.Context, pageID string, group *ComponentGroupInput) (*ComponentGroup, error) {
	u := fmt.Sprintf("pages/%s/component-groups", pageID)
	groupReq := &ComponentGroupRequest{ComponentGroup: group}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, groupReq)
	if err != nil {
		return nil, err
	}

	newGroup := new(ComponentGroup)
	_, err = s.client.Do(ctx, req, newGroup)
	if err != nil {
		return nil, err
	}

	return newGroup, nil
}

func (s *ComponentGroupsService) Update(ctx context.Context, pageID, groupID string, group *ComponentGroupInput) (*ComponentGroup, error) {
	u := fmt.Sprintf("pages/%s/component-groups/%s", pageID, groupID)
	groupReq := &ComponentGroupRequest{ComponentGroup: group}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, groupReq)
	if err != nil {
		return nil, err
	}

	updatedGroup := new(ComponentGroup)
	_, err = s.client.Do(ctx, req, updatedGroup)
	if err != nil {
		return nil, err
	}

	return updatedGroup, nil
}

func (s *ComponentGroupsService) Delete(ctx context.Context, pageID, groupID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/component-groups/%s", pageID, groupID)
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
