package statuspage

import (
	"context"
	"fmt"
	"net/http"
)

// PageAccessUsersService handles communication with page access users related to audience-specific status pages
type PageAccessUsersService struct {
	client *Client
}

type PageAccessUserRequest struct {
	PageAccessUser *PageAccessUserInput `json:"page_access_user"`
}

type PageAccessUserInput struct {
	Email              string   `json:"email,omitempty"`
	ExternalLogin      string   `json:"external_login,omitempty"`
	PageAccessGroupIDs []string `json:"page_access_group_ids,omitempty"`
	ComponentIDs       []string `json:"component_ids,omitempty"`
	MetricIDs          []string `json:"metric_ids,omitempty"`
}

// List gets a list of page access users for audience-specific status pages
func (s *PageAccessUsersService) List(ctx context.Context, pageID string) ([]*PageAccessUser, *Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_users", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*PageAccessUser
	resp, err := s.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}

// Get retrieves a specific page access user by ID
func (s *PageAccessUsersService) Get(ctx context.Context, pageID, userID string) (*PageAccessUser, *Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_users/%s", pageID, userID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := new(PageAccessUser)
	resp, err := s.client.Do(ctx, req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// Create creates a new page access user for audience-specific access control
func (s *PageAccessUsersService) Create(ctx context.Context, pageID string, user *PageAccessUserInput) (*PageAccessUser, *Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_users", pageID)
	userReq := &PageAccessUserRequest{PageAccessUser: user}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, userReq)
	if err != nil {
		return nil, nil, err
	}

	newUser := new(PageAccessUser)
	resp, err := s.client.Do(ctx, req, newUser)
	if err != nil {
		return nil, resp, err
	}

	return newUser, resp, nil
}

// Update modifies an existing page access user
func (s *PageAccessUsersService) Update(ctx context.Context, pageID, userID string, user *PageAccessUserInput) (*PageAccessUser, *Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_users/%s", pageID, userID)
	userReq := &PageAccessUserRequest{PageAccessUser: user}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, userReq)
	if err != nil {
		return nil, nil, err
	}

	updatedUser := new(PageAccessUser)
	resp, err := s.client.Do(ctx, req, updatedUser)
	if err != nil {
		return nil, resp, err
	}

	return updatedUser, resp, nil
}

// Delete removes a page access user
func (s *PageAccessUsersService) Delete(ctx context.Context, pageID, userID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_users/%s", pageID, userID)
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

// PageAccessGroupsService handles communication with page access groups for audience-specific status pages
type PageAccessGroupsService struct {
	client *Client
}

type PageAccessGroupRequest struct {
	PageAccessGroup *PageAccessGroupInput `json:"page_access_group"`
}

type PageAccessGroupInput struct {
	Name               string   `json:"name,omitempty"`
	Description        string   `json:"description,omitempty"`
	Color              string   `json:"color,omitempty"`
	ComponentIDs       []string `json:"component_ids,omitempty"`
	MetricIDs          []string `json:"metric_ids,omitempty"`
	PageAccessUserIDs  []string `json:"page_access_user_ids,omitempty"`
	ExternalIdentifier string   `json:"external_identifier,omitempty"`
}

// List gets a list of page access groups for audience-specific status pages
func (s *PageAccessGroupsService) List(ctx context.Context, pageID string) ([]*PageAccessGroup, *Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_groups", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var groups []*PageAccessGroup
	resp, err := s.client.Do(ctx, req, &groups)
	if err != nil {
		return nil, resp, err
	}

	return groups, resp, nil
}

// Get retrieves a specific page access group by ID
func (s *PageAccessGroupsService) Get(ctx context.Context, pageID, groupID string) (*PageAccessGroup, *Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_groups/%s", pageID, groupID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(PageAccessGroup)
	resp, err := s.client.Do(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

// Create creates a new page access group for organizing users with similar access needs
func (s *PageAccessGroupsService) Create(ctx context.Context, pageID string, group *PageAccessGroupInput) (*PageAccessGroup, *Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_groups", pageID)
	groupReq := &PageAccessGroupRequest{PageAccessGroup: group}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, groupReq)
	if err != nil {
		return nil, nil, err
	}

	newGroup := new(PageAccessGroup)
	resp, err := s.client.Do(ctx, req, newGroup)
	if err != nil {
		return nil, resp, err
	}

	return newGroup, resp, nil
}

// Update modifies an existing page access group
func (s *PageAccessGroupsService) Update(ctx context.Context, pageID, groupID string, group *PageAccessGroupInput) (*PageAccessGroup, *Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_groups/%s", pageID, groupID)
	groupReq := &PageAccessGroupRequest{PageAccessGroup: group}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, groupReq)
	if err != nil {
		return nil, nil, err
	}

	updatedGroup := new(PageAccessGroup)
	resp, err := s.client.Do(ctx, req, updatedGroup)
	if err != nil {
		return nil, resp, err
	}

	return updatedGroup, resp, nil
}

// Delete removes a page access group
func (s *PageAccessGroupsService) Delete(ctx context.Context, pageID, groupID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/page_access_groups/%s", pageID, groupID)
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

// TemplatesService handles communication with incident templates for faster incident creation
type TemplatesService struct {
	client *Client
}

type TemplateRequest struct {
	Template *TemplateInput `json:"template"`
}

type TemplateInput struct {
	Name                    string `json:"name,omitempty"`
	Body                    string `json:"body,omitempty"`
	GroupID                 string `json:"group_id,omitempty"`
	UpdateStatus            string `json:"update_status,omitempty"`
	ShouldTweet             *bool  `json:"should_tweet,omitempty"`
	ShouldSendNotifications *bool  `json:"should_send_notifications,omitempty"`
}

// List gets a list of incident templates for creating incidents with pre-filled information
func (s *TemplatesService) List(ctx context.Context, pageID string) ([]*Template, *Response, error) {
	u := fmt.Sprintf("pages/%s/incident_templates", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var templates []*Template
	resp, err := s.client.Do(ctx, req, &templates)
	if err != nil {
		return nil, resp, err
	}

	return templates, resp, nil
}

// Get retrieves a specific incident template by ID
func (s *TemplatesService) Get(ctx context.Context, pageID, templateID string) (*Template, *Response, error) {
	u := fmt.Sprintf("pages/%s/incident_templates/%s", pageID, templateID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	template := new(Template)
	resp, err := s.client.Do(ctx, req, template)
	if err != nil {
		return nil, resp, err
	}

	return template, resp, nil
}

// Create creates a new incident template with pre-filled name and message for faster incident creation
func (s *TemplatesService) Create(ctx context.Context, pageID string, template *TemplateInput) (*Template, *Response, error) {
	u := fmt.Sprintf("pages/%s/incident_templates", pageID)
	templateReq := &TemplateRequest{Template: template}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, templateReq)
	if err != nil {
		return nil, nil, err
	}

	newTemplate := new(Template)
	resp, err := s.client.Do(ctx, req, newTemplate)
	if err != nil {
		return nil, resp, err
	}

	return newTemplate, resp, nil
}

// Update modifies an existing incident template
func (s *TemplatesService) Update(ctx context.Context, pageID, templateID string, template *TemplateInput) (*Template, *Response, error) {
	u := fmt.Sprintf("pages/%s/incident_templates/%s", pageID, templateID)
	templateReq := &TemplateRequest{Template: template}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, templateReq)
	if err != nil {
		return nil, nil, err
	}

	updatedTemplate := new(Template)
	resp, err := s.client.Do(ctx, req, updatedTemplate)
	if err != nil {
		return nil, resp, err
	}

	return updatedTemplate, resp, nil
}

// Delete removes an incident template
func (s *TemplatesService) Delete(ctx context.Context, pageID, templateID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/incident_templates/%s", pageID, templateID)
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

// StatusEmbedConfigService handles communication with status embed configs for customizing embedded status widgets
type StatusEmbedConfigService struct {
	client *Client
}

type StatusEmbedConfigRequest struct {
	StatusEmbedConfig *StatusEmbedConfigInput `json:"status_embed_config"`
}

type StatusEmbedConfigInput struct {
	Position                   string `json:"position,omitempty"`
	IncidentBackgroundColor    string `json:"incident_background_color,omitempty"`
	IncidentTextColor          string `json:"incident_text_color,omitempty"`
	MaintenanceBackgroundColor string `json:"maintenance_background_color,omitempty"`
	MaintenanceTextColor       string `json:"maintenance_text_color,omitempty"`
}

// Get retrieves status embed config settings for customizing the appearance of embedded status widgets
func (s *StatusEmbedConfigService) Get(ctx context.Context, pageID string) (*StatusEmbedConfig, *Response, error) {
	u := fmt.Sprintf("pages/%s/status_embed_config", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	config := new(StatusEmbedConfig)
	resp, err := s.client.Do(ctx, req, config)
	if err != nil {
		return nil, resp, err
	}

	return config, resp, nil
}

// Update modifies the status embed config settings for customizing widget appearance
func (s *StatusEmbedConfigService) Update(ctx context.Context, pageID string, config *StatusEmbedConfigInput) (*StatusEmbedConfig, *Response, error) {
	u := fmt.Sprintf("pages/%s/status_embed_config", pageID)
	configReq := &StatusEmbedConfigRequest{StatusEmbedConfig: config}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, configReq)
	if err != nil {
		return nil, nil, err
	}

	updatedConfig := new(StatusEmbedConfig)
	resp, err := s.client.Do(ctx, req, updatedConfig)
	if err != nil {
		return nil, resp, err
	}

	return updatedConfig, resp, nil
}
