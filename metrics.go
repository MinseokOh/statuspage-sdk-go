package statuspage

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type MetricsService struct {
	client *Client
}

type MetricRequest struct {
	Metric *MetricInput `json:"metric"`
}

type MetricInput struct {
	Name          string  `json:"name,omitempty"`
	Suffix        string  `json:"suffix,omitempty"`
	YAxisMin      float64 `json:"y_axis_min,omitempty"`
	YAxisMax      float64 `json:"y_axis_max,omitempty"`
	YAxisHidden   *bool   `json:"y_axis_hidden,omitempty"`
	Transform     string  `json:"transform,omitempty"`
	DecimalPlaces int     `json:"decimal_places,omitempty"`
	Tooltip       string  `json:"tooltip,omitempty"`
	DisplayName   string  `json:"display_name,omitempty"`
}

type MetricDataRequest struct {
	Data *MetricDataInput `json:"data"`
}

type MetricDataInput struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

type MetricDataListOptions struct {
	From *time.Time `url:"from,omitempty"`
	To   *time.Time `url:"to,omitempty"`
}

func (s *MetricsService) List(ctx context.Context, pageID string) ([]*Metric, *Response, error) {
	u := fmt.Sprintf("pages/%s/metrics", pageID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var metrics []*Metric
	resp, err := s.client.Do(ctx, req, &metrics)
	if err != nil {
		return nil, resp, err
	}

	return metrics, resp, nil
}

func (s *MetricsService) Get(ctx context.Context, pageID, metricID string) (*Metric, *Response, error) {
	u := fmt.Sprintf("pages/%s/metrics/%s", pageID, metricID)
	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	metric := new(Metric)
	resp, err := s.client.Do(ctx, req, metric)
	if err != nil {
		return nil, resp, err
	}

	return metric, resp, nil
}

func (s *MetricsService) Create(ctx context.Context, pageID string, metric *MetricInput) (*Metric, *Response, error) {
	u := fmt.Sprintf("pages/%s/metrics", pageID)
	metricReq := &MetricRequest{Metric: metric}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, metricReq)
	if err != nil {
		return nil, nil, err
	}

	newMetric := new(Metric)
	resp, err := s.client.Do(ctx, req, newMetric)
	if err != nil {
		return nil, resp, err
	}

	return newMetric, resp, nil
}

func (s *MetricsService) Update(ctx context.Context, pageID, metricID string, metric *MetricInput) (*Metric, *Response, error) {
	u := fmt.Sprintf("pages/%s/metrics/%s", pageID, metricID)
	metricReq := &MetricRequest{Metric: metric}
	req, err := s.client.NewRequest(ctx, http.MethodPatch, u, metricReq)
	if err != nil {
		return nil, nil, err
	}

	updatedMetric := new(Metric)
	resp, err := s.client.Do(ctx, req, updatedMetric)
	if err != nil {
		return nil, resp, err
	}

	return updatedMetric, resp, nil
}

func (s *MetricsService) Delete(ctx context.Context, pageID, metricID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/metrics/%s", pageID, metricID)
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

func (s *MetricsService) AddData(ctx context.Context, pageID, metricID string, data *MetricDataInput) (*MetricData, *Response, error) {
	u := fmt.Sprintf("pages/%s/metrics/%s/data", pageID, metricID)
	dataReq := &MetricDataRequest{Data: data}
	req, err := s.client.NewRequest(ctx, http.MethodPost, u, dataReq)
	if err != nil {
		return nil, nil, err
	}

	metricData := new(MetricData)
	resp, err := s.client.Do(ctx, req, metricData)
	if err != nil {
		return nil, resp, err
	}

	return metricData, resp, nil
}

func (s *MetricsService) GetData(ctx context.Context, pageID, metricID string, opts *MetricDataListOptions) ([]*MetricData, *Response, error) {
	u := fmt.Sprintf("pages/%s/metrics/%s/data", pageID, metricID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var data []*MetricData
	resp, err := s.client.Do(ctx, req, &data)
	if err != nil {
		return nil, resp, err
	}

	return data, resp, nil
}

func (s *MetricsService) DeleteData(ctx context.Context, pageID, metricID string) (*Response, error) {
	u := fmt.Sprintf("pages/%s/metrics/%s/data", pageID, metricID)
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
