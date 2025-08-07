package statuspage

import (
	"context"
	"fmt"
	"net/http"
)

type SubscribersService struct {
	client *Client
}

type SubscriberRequest struct {
	Subscriber *SubscriberInput `json:"subscriber"`
}

type SubscriberInput struct {
	Email                        string   `json:"email,omitempty"`
	Endpoint                     string   `json:"endpoint,omitempty"`
	PhoneCountry                 string   `json:"phone_country,omitempty"`
	PhoneNumber                  string   `json:"phone_number,omitempty"`
	SkipConfirmationNotification *bool    `json:"skip_confirmation_notification,omitempty"`
	ComponentIDs                 []string `json:"component_ids,omitempty"`
	PageAccessUserID             string   `json:"page_access_user_id,omitempty"`
}

type SubscriberListOptions struct {
	Q       string `url:"q,omitempty"`
	Sort    string `url:"sort,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
}

func (s *SubscribersService) List(ctx context.Context, pageID string, opts *SubscriberListOptions) ([]*Subscriber, error) {
	u := fmt.Sprintf("pages/%s/subscribers", pageID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var subscribers []*Subscriber
	_, err = s.client.Do(ctx, req, &subscribers)
	if err != nil {
		return nil, err
	}

	return subscribers, nil
}

func (s *SubscribersService) Get(ctx context.Context, pageID, subscriberID string) (*Subscriber, error) {
	u := fmt.Sprintf("pages/%s/subscribers/%s", pageID, subscriberID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	subscriber := new(Subscriber)
	_, err = s.client.Do(ctx, req, subscriber)
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}

func (s *SubscribersService) Create(ctx context.Context, pageID string, subscriber *SubscriberInput) (*Subscriber, error) {
	u := fmt.Sprintf("pages/%s/subscribers", pageID)
	subscriberReq := &SubscriberRequest{Subscriber: subscriber}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, subscriberReq)
	if err != nil {
		return nil, err
	}

	newSubscriber := new(Subscriber)
	_, err = s.client.Do(ctx, req, newSubscriber)
	if err != nil {
		return nil, err
	}

	return newSubscriber, nil
}

func (s *SubscribersService) Update(ctx context.Context, pageID, subscriberID string, subscriber *SubscriberInput) (*Subscriber, error) {
	u := fmt.Sprintf("pages/%s/subscribers/%s", pageID, subscriberID)
	subscriberReq := &SubscriberRequest{Subscriber: subscriber}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, subscriberReq)
	if err != nil {
		return nil, err
	}

	updatedSubscriber := new(Subscriber)
	_, err = s.client.Do(ctx, req, updatedSubscriber)
	if err != nil {
		return nil, err
	}

	return updatedSubscriber, nil
}

func (s *SubscribersService) Delete(ctx context.Context, pageID, subscriberID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/subscribers/%s", pageID, subscriberID)
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

func (s *SubscribersService) Reactivate(ctx context.Context, pageID, subscriberID string) (*Subscriber, error) {
	u := fmt.Sprintf("pages/%s/subscribers/%s/reactivate", pageID, subscriberID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	subscriber := new(Subscriber)
	_, err = s.client.Do(ctx, req, subscriber)
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}

func (s *SubscribersService) Unsubscribe(ctx context.Context, pageID, subscriberID string) (*Subscriber, error) {
	u := fmt.Sprintf("pages/%s/subscribers/%s/unsubscribe", pageID, subscriberID)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	subscriber := new(Subscriber)
	_, err = s.client.Do(ctx, req, subscriber)
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}

func (s *SubscribersService) ResendConfirmation(ctx context.Context, pageID, subscriberID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/subscribers/%s/resend_confirmation", pageID, subscriberID)
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
