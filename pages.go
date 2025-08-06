package statuspage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// PagesService handles all page-related API operations for status page management
type PagesService struct {
	client *Client
}

// PageRequest wraps page input data for API requests
type PageRequest struct {
	Page *PageInput `json:"page"`
}

// PageInput contains the editable fields for creating or updating a status page
type PageInput struct {
	Name                     string `json:"name,omitempty"`
	Domain                   string `json:"domain,omitempty"`
	URL                      string `json:"url,omitempty"`
	Subdomain                string `json:"subdomain,omitempty"`
	Branding                 string `json:"branding,omitempty"`
	CSSBodyBackgroundColor   string `json:"css_body_background_color,omitempty"`
	CSSFontColor             string `json:"css_font_color,omitempty"`
	CSSLightFontColor        string `json:"css_light_font_color,omitempty"`
	CSSGreens                string `json:"css_greens,omitempty"`
	CSSYellows               string `json:"css_yellows,omitempty"`
	CSSOranges               string `json:"css_oranges,omitempty"`
	CSSBlues                 string `json:"css_blues,omitempty"`
	CSSReds                  string `json:"css_reds,omitempty"`
	CSSBorderColor           string `json:"css_border_color,omitempty"`
	CSSGraphColor            string `json:"css_graph_color,omitempty"`
	CSSLinkColor             string `json:"css_link_color,omitempty"`
	CSSNoData                string `json:"css_no_data,omitempty"`
	HiddenFromSearch         *bool  `json:"hidden_from_search,omitempty"`
	ViewersMustBeTeamMembers *bool  `json:"viewers_must_be_team_members,omitempty"`
	AllowPageSubscribers     *bool  `json:"allow_page_subscribers,omitempty"`
	AllowIncidentSubscribers *bool  `json:"allow_incident_subscribers,omitempty"`
	AllowEmailSubscribers    *bool  `json:"allow_email_subscribers,omitempty"`
	AllowSmsSubscribers      *bool  `json:"allow_sms_subscribers,omitempty"`
	AllowRssAtomFeeds        *bool  `json:"allow_rss_atom_feeds,omitempty"`
	AllowWebhookSubscribers  *bool  `json:"allow_webhook_subscribers,omitempty"`
	NotificationsFromEmail   string `json:"notifications_from_email,omitempty"`
	NotificationsEmailFooter string `json:"notifications_email_footer,omitempty"`
	TimeZone                 string `json:"time_zone,omitempty"`
	City                     string `json:"city,omitempty"`
	State                    string `json:"state,omitempty"`
	Country                  string `json:"country,omitempty"`
	TwitterUsername          string `json:"twitter_username,omitempty"`
	PageDescription          string `json:"page_description,omitempty"`
	Headline                 string `json:"headline,omitempty"`
	SupportURL               string `json:"support_url,omitempty"`
	IPRestrictions           string `json:"ip_restrictions,omitempty"`
	FaviconLogo              json.RawMessage `json:"favicon_logo,omitempty"`
	TransactionalLogo        json.RawMessage `json:"transactional_logo,omitempty"`
	HeroCover                json.RawMessage `json:"hero_cover,omitempty"`
	EmailLogo                json.RawMessage `json:"email_logo,omitempty"`
	TwitterLogo              json.RawMessage `json:"twitter_logo,omitempty"`
}

// List retrieves all status pages accessible with the current API key
func (s *PagesService) List(ctx context.Context) ([]*Page, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "pages", nil)
	if err != nil {
		return nil, nil, err
	}

	var pages []*Page
	resp, err := s.client.Do(ctx, req, &pages)
	if err != nil {
		return nil, resp, err
	}

	return pages, resp, nil
}

// Get retrieves a specific status page by its unique identifier
func (s *PagesService) Get(ctx context.Context, pageID string) (*Page, *Response, error) {
	u := fmt.Sprintf("pages/%s", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(Page)
	resp, err := s.client.Do(ctx, req, page)
	if err != nil {
		return nil, resp, err
	}

	return page, resp, nil
}

// Update modifies an existing status page with new configuration settings
func (s *PagesService) Update(ctx context.Context, pageID string, page *PageInput) (*Page, *Response, error) {
	u := fmt.Sprintf("pages/%s", pageID)
	pageReq := &PageRequest{Page: page}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, pageReq)
	if err != nil {
		return nil, nil, err
	}

	updatedPage := new(Page)
	resp, err := s.client.Do(ctx, req, updatedPage)
	if err != nil {
		return nil, resp, err
	}

	return updatedPage, resp, nil
}
