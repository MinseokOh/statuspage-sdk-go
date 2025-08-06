package statuspage

import "time"

// Page represents a status page configuration with customization settings and metadata
type Page struct {
	ID                       string    `json:"id,omitempty"`
	CreatedAt                time.Time `json:"created_at,omitempty"`
	UpdatedAt                time.Time `json:"updated_at,omitempty"`
	Name                     string    `json:"name,omitempty"`
	PageDescription          string    `json:"page_description,omitempty"`
	Headline                 string    `json:"headline,omitempty"`
	Branding                 string    `json:"branding,omitempty"`
	Subdomain                string    `json:"subdomain,omitempty"`
	Domain                   string    `json:"domain,omitempty"`
	URL                      string    `json:"url,omitempty"`
	SupportURL               string    `json:"support_url,omitempty"`
	HiddenFromSearch         bool      `json:"hidden_from_search,omitempty"`
	AllowPageSubscribers     bool      `json:"allow_page_subscribers,omitempty"`
	AllowIncidentSubscribers bool      `json:"allow_incident_subscribers,omitempty"`
	AllowEmailSubscribers    bool      `json:"allow_email_subscribers,omitempty"`
	AllowSmsSubscribers      bool      `json:"allow_sms_subscribers,omitempty"`
	AllowRssAtomFeeds        bool      `json:"allow_rss_atom_feeds,omitempty"`
	AllowWebhookSubscribers  bool      `json:"allow_webhook_subscribers,omitempty"`
	NotificationsFromEmail   string    `json:"notifications_from_email,omitempty"`
	NotificationsEmailFooter string    `json:"notifications_email_footer,omitempty"`
	ActivityScore            float64   `json:"activity_score,omitempty"`
	TwitterUsername          string    `json:"twitter_username,omitempty"`
	ViewersMustBeTeamMembers bool      `json:"viewers_must_be_team_members,omitempty"`
	IPRestrictions           string    `json:"ip_restrictions,omitempty"`
	City                     string    `json:"city,omitempty"`
	State                    string    `json:"state,omitempty"`
	Country                  string    `json:"country,omitempty"`
	TimeZone                 string    `json:"time_zone,omitempty"`
	CSSBodyBackgroundColor   string    `json:"css_body_background_color,omitempty"`
	CSSFontColor             string    `json:"css_font_color,omitempty"`
	CSSLightFontColor        string    `json:"css_light_font_color,omitempty"`
	CSSGreens                string    `json:"css_greens,omitempty"`
	CSSYellows               string    `json:"css_yellows,omitempty"`
	CSSOranges               string    `json:"css_oranges,omitempty"`
	CSSBlues                 string    `json:"css_blues,omitempty"`
	CSSReds                  string    `json:"css_reds,omitempty"`
	CSSBorderColor           string    `json:"css_border_color,omitempty"`
	CSSGraphColor            string    `json:"css_graph_color,omitempty"`
	CSSLinkColor             string    `json:"css_link_color,omitempty"`
	CSSNoData                string    `json:"css_no_data,omitempty"`
	FaviconLogo              string    `json:"favicon_logo,omitempty"`
	TransactionalLogo        string    `json:"transactional_logo,omitempty"`
	HeroCover                string    `json:"hero_cover,omitempty"`
	EmailLogo                string    `json:"email_logo,omitempty"`
	TwitterLogo              string    `json:"twitter_logo,omitempty"`
}

// Component represents a service or system component tracked on the status page
type Component struct {
	ID                 string     `json:"id,omitempty"`
	PageID             string     `json:"page_id,omitempty"`
	GroupID            string     `json:"group_id,omitempty"`
	CreatedAt          time.Time  `json:"created_at,omitempty"`
	UpdatedAt          time.Time  `json:"updated_at,omitempty"`
	Group              bool       `json:"group,omitempty"`
	Name               string     `json:"name,omitempty"`
	Description        string     `json:"description,omitempty"`
	Position           int        `json:"position,omitempty"`
	Status             string     `json:"status,omitempty"`
	Showcase           bool       `json:"showcase,omitempty"`
	OnlyShowIfDegraded bool       `json:"only_show_if_degraded,omitempty"`
	AutomationEmail    string     `json:"automation_email,omitempty"`
	StartDate          *time.Time `json:"start_date,omitempty"`
}

// ComponentGroup represents a logical grouping of related components for better organization
type ComponentGroup struct {
	ID          string    `json:"id,omitempty"`
	PageID      string    `json:"page_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Components  []string  `json:"components,omitempty"`
	Position    int       `json:"position,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// Incident represents a service disruption or maintenance event affecting components
type Incident struct {
	ID                            string              `json:"id,omitempty"`
	Components                    []Component         `json:"components,omitempty"`
	CreatedAt                     time.Time           `json:"created_at,omitempty"`
	Impact                        string              `json:"impact,omitempty"`
	ImpactOverride                string              `json:"impact_override,omitempty"`
	IncidentUpdates               []IncidentUpdate    `json:"incident_updates,omitempty"`
	MonitoringAt                  *time.Time          `json:"monitoring_at,omitempty"`
	Name                          string              `json:"name,omitempty"`
	PageID                        string              `json:"page_id,omitempty"`
	PostmortemBody                string              `json:"postmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt   *time.Time          `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored             bool                `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers bool                `json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter     bool                `json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt         *time.Time          `json:"postmortem_published_at,omitempty"`
	ResolvedAt                    *time.Time          `json:"resolved_at,omitempty"`
	ScheduledAutoCompleted        bool                `json:"scheduled_auto_completed,omitempty"`
	ScheduledAutoInProgress       bool                `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledFor                  *time.Time          `json:"scheduled_for,omitempty"`
	ScheduledRemindPrior          bool                `json:"scheduled_remind_prior,omitempty"`
	ScheduledRemindedAt           *time.Time          `json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil                *time.Time          `json:"scheduled_until,omitempty"`
	Shortlink                     string              `json:"shortlink,omitempty"`
	Status                        string              `json:"status,omitempty"`
	UpdatedAt                     time.Time           `json:"updated_at,omitempty"`
	ComponentIDs                  []string            `json:"component_ids,omitempty"`
	AffectedComponents            []AffectedComponent `json:"affected_components,omitempty"`
}

// IncidentUpdate represents a status update for an ongoing incident with communication details
type IncidentUpdate struct {
	ID                   string              `json:"id,omitempty"`
	IncidentID           string              `json:"incident_id,omitempty"`
	AffectedComponents   []AffectedComponent `json:"affected_components,omitempty"`
	Body                 string              `json:"body,omitempty"`
	CreatedAt            time.Time           `json:"created_at,omitempty"`
	CustomTweet          string              `json:"custom_tweet,omitempty"`
	DeliverNotifications bool                `json:"deliver_notifications,omitempty"`
	DisplayAt            time.Time           `json:"display_at,omitempty"`
	Status               string              `json:"status,omitempty"`
	TweetID              string              `json:"tweet_id,omitempty"`
	TwitterUpdatedAt     *time.Time          `json:"twitter_updated_at,omitempty"`
	UpdatedAt            time.Time           `json:"updated_at,omitempty"`
	WantsTwitterUpdate   bool                `json:"wants_twitter_update,omitempty"`
}

// AffectedComponent represents a component impacted by an incident with status change information
type AffectedComponent struct {
	Code      string `json:"code,omitempty"`
	Name      string `json:"name,omitempty"`
	OldStatus string `json:"old_status,omitempty"`
	NewStatus string `json:"new_status,omitempty"`
}

// Subscriber represents a user subscribed to receive notifications about status updates
type Subscriber struct {
	ID                           string      `json:"id,omitempty"`
	SkipConfirmationNotification bool        `json:"skip_confirmation_notification,omitempty"`
	Mode                         string      `json:"mode,omitempty"`
	Email                        string      `json:"email,omitempty"`
	Endpoint                     string      `json:"endpoint,omitempty"`
	PhoneCountry                 string      `json:"phone_country,omitempty"`
	PhoneNumber                  string      `json:"phone_number,omitempty"`
	DisplayPhoneNumber           string      `json:"display_phone_number,omitempty"`
	ObfuscatedChannelName        string      `json:"obfuscated_channel_name,omitempty"`
	WorkspaceURL                 string      `json:"workspace_url,omitempty"`
	CreatedAt                    time.Time   `json:"created_at,omitempty"`
	ConfirmedAt                  *time.Time  `json:"confirmed_at,omitempty"`
	PurgeAt                      *time.Time  `json:"purge_at,omitempty"`
	UnsubscribedAt               *time.Time  `json:"unsubscribed_at,omitempty"`
	PageID                       string      `json:"page_id,omitempty"`
	PageAccessUserID             string      `json:"page_access_user_id,omitempty"`
	Components                   []Component `json:"components,omitempty"`
	Quarantined                  bool        `json:"quarantined,omitempty"`
	QuarantinedAt                *time.Time  `json:"quarantined_at,omitempty"`
	ComponentIDs                 []string    `json:"component_ids,omitempty"`
}

// Metric represents a performance metric displayed on the status page with configuration settings
type Metric struct {
	ID                  string     `json:"id,omitempty"`
	PageID              string     `json:"page_id,omitempty"`
	MetricsProviderID   string     `json:"metrics_provider_id,omitempty"`
	Name                string     `json:"name,omitempty"`
	DisplayName         string     `json:"display_name,omitempty"`
	Tooltip             string     `json:"tooltip,omitempty"`
	BackfillPercentage  float64    `json:"backfill_percentage,omitempty"`
	BackfilledAt        *time.Time `json:"backfilled_at,omitempty"`
	YAxisMin            float64    `json:"y_axis_min,omitempty"`
	YAxisMax            float64    `json:"y_axis_max,omitempty"`
	YAxisHidden         bool       `json:"y_axis_hidden,omitempty"`
	Suffix              string     `json:"suffix,omitempty"`
	DecimalPlaces       int        `json:"decimal_places,omitempty"`
	MostRecentDataAt    *time.Time `json:"most_recent_data_at,omitempty"`
	CreatedAt           time.Time  `json:"created_at,omitempty"`
	UpdatedAt           time.Time  `json:"updated_at,omitempty"`
	LastFetchedAt       *time.Time `json:"last_fetched_at,omitempty"`
	BackfillStatus      string     `json:"backfill_status,omitempty"`
	BackfillErrorsCount int        `json:"backfill_errors_count,omitempty"`
}

// MetricData represents a single data point for a metric with timestamp and value
type MetricData struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

// MetricProvider represents a third-party service providing metric data with authentication details
type MetricProvider struct {
	ID             string    `json:"id,omitempty"`
	Type           string    `json:"type,omitempty"`
	Email          string    `json:"email,omitempty"`
	Password       string    `json:"password,omitempty"`
	APIKey         string    `json:"api_key,omitempty"`
	APIToken       string    `json:"api_token,omitempty"`
	ApplicationKey string    `json:"application_key,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

// PageAccessUser represents a user with restricted access to specific components and metrics on audience-specific pages
type PageAccessUser struct {
	ID                 string    `json:"id,omitempty"`
	PageID             string    `json:"page_id,omitempty"`
	Email              string    `json:"email,omitempty"`
	ExternalLogin      string    `json:"external_login,omitempty"`
	PageAccessGroupIDs []string  `json:"page_access_group_ids,omitempty"`
	ComponentIDs       []string  `json:"component_ids,omitempty"`
	MetricIDs          []string  `json:"metric_ids,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	StatusPageURL      string    `json:"status_page_url,omitempty"`
}

// PageAccessGroup represents a group of users with similar access permissions for audience-specific status pages
type PageAccessGroup struct {
	ID                 string    `json:"id,omitempty"`
	PageID             string    `json:"page_id,omitempty"`
	Name               string    `json:"name,omitempty"`
	Description        string    `json:"description,omitempty"`
	Color              string    `json:"color,omitempty"`
	ComponentIDs       []string  `json:"component_ids,omitempty"`
	MetricIDs          []string  `json:"metric_ids,omitempty"`
	PageAccessUserIDs  []string  `json:"page_access_user_ids,omitempty"`
	ExternalIdentifier string    `json:"external_identifier,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

// Template represents a pre-configured incident template for faster incident creation with predefined content
type Template struct {
	ID                      string    `json:"id,omitempty"`
	PageID                  string    `json:"page_id,omitempty"`
	Name                    string    `json:"name,omitempty"`
	Body                    string    `json:"body,omitempty"`
	GroupID                 string    `json:"group_id,omitempty"`
	UpdateStatus            string    `json:"update_status,omitempty"`
	ShouldTweet             bool      `json:"should_tweet,omitempty"`
	ShouldSendNotifications bool      `json:"should_send_notifications,omitempty"`
	CreatedAt               time.Time `json:"created_at,omitempty"`
	UpdatedAt               time.Time `json:"updated_at,omitempty"`
}

// StatusEmbedConfig represents configuration for embedded status widgets with customizable appearance settings
type StatusEmbedConfig struct {
	PageID                     string    `json:"page_id,omitempty"`
	Position                   string    `json:"position,omitempty"`
	IncidentBackgroundColor    string    `json:"incident_background_color,omitempty"`
	IncidentTextColor          string    `json:"incident_text_color,omitempty"`
	MaintenanceBackgroundColor string    `json:"maintenance_background_color,omitempty"`
	MaintenanceTextColor       string    `json:"maintenance_text_color,omitempty"`
	CreatedAt                  time.Time `json:"created_at,omitempty"`
	UpdatedAt                  time.Time `json:"updated_at,omitempty"`
}

// ListOptions represents pagination parameters for API list requests
type ListOptions struct {
	Page    int `url:"page,omitempty"`
	PerPage int `url:"per_page,omitempty"`
}

// ErrorEntity represents an error response from the API with descriptive message
type ErrorEntity struct {
	Error string `json:"error,omitempty"`
}
