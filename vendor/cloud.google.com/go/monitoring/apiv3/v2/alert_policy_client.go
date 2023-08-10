// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package monitoring

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newAlertPolicyClientHook clientHook

// AlertPolicyCallOptions contains the retry settings for each method of AlertPolicyClient.
type AlertPolicyCallOptions struct {
	ListAlertPolicies []gax.CallOption
	GetAlertPolicy    []gax.CallOption
	CreateAlertPolicy []gax.CallOption
	DeleteAlertPolicy []gax.CallOption
	UpdateAlertPolicy []gax.CallOption
}

func defaultAlertPolicyGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("monitoring.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("monitoring.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://monitoring.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAlertPolicyCallOptions() *AlertPolicyCallOptions {
	return &AlertPolicyCallOptions{
		ListAlertPolicies: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetAlertPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateAlertPolicy: []gax.CallOption{},
		DeleteAlertPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        30000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateAlertPolicy: []gax.CallOption{},
	}
}

// internalAlertPolicyClient is an interface that defines the methods available from Cloud Monitoring API.
type internalAlertPolicyClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListAlertPolicies(context.Context, *monitoringpb.ListAlertPoliciesRequest, ...gax.CallOption) *AlertPolicyIterator
	GetAlertPolicy(context.Context, *monitoringpb.GetAlertPolicyRequest, ...gax.CallOption) (*monitoringpb.AlertPolicy, error)
	CreateAlertPolicy(context.Context, *monitoringpb.CreateAlertPolicyRequest, ...gax.CallOption) (*monitoringpb.AlertPolicy, error)
	DeleteAlertPolicy(context.Context, *monitoringpb.DeleteAlertPolicyRequest, ...gax.CallOption) error
	UpdateAlertPolicy(context.Context, *monitoringpb.UpdateAlertPolicyRequest, ...gax.CallOption) (*monitoringpb.AlertPolicy, error)
}

// AlertPolicyClient is a client for interacting with Cloud Monitoring API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The AlertPolicyService API is used to manage (list, create, delete,
// edit) alert policies in Cloud Monitoring. An alerting policy is
// a description of the conditions under which some aspect of your
// system is considered to be “unhealthy” and the ways to notify
// people or services about this state. In addition to using this API, alert
// policies can also be managed through
// Cloud Monitoring (at https://cloud.google.com/monitoring/docs/),
// which can be reached by clicking the “Monitoring” tab in
// Cloud console (at https://console.cloud.google.com/).
type AlertPolicyClient struct {
	// The internal transport-dependent client.
	internalClient internalAlertPolicyClient

	// The call options for this service.
	CallOptions *AlertPolicyCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AlertPolicyClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AlertPolicyClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AlertPolicyClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListAlertPolicies lists the existing alerting policies for the workspace.
func (c *AlertPolicyClient) ListAlertPolicies(ctx context.Context, req *monitoringpb.ListAlertPoliciesRequest, opts ...gax.CallOption) *AlertPolicyIterator {
	return c.internalClient.ListAlertPolicies(ctx, req, opts...)
}

// GetAlertPolicy gets a single alerting policy.
func (c *AlertPolicyClient) GetAlertPolicy(ctx context.Context, req *monitoringpb.GetAlertPolicyRequest, opts ...gax.CallOption) (*monitoringpb.AlertPolicy, error) {
	return c.internalClient.GetAlertPolicy(ctx, req, opts...)
}

// CreateAlertPolicy creates a new alerting policy.
func (c *AlertPolicyClient) CreateAlertPolicy(ctx context.Context, req *monitoringpb.CreateAlertPolicyRequest, opts ...gax.CallOption) (*monitoringpb.AlertPolicy, error) {
	return c.internalClient.CreateAlertPolicy(ctx, req, opts...)
}

// DeleteAlertPolicy deletes an alerting policy.
func (c *AlertPolicyClient) DeleteAlertPolicy(ctx context.Context, req *monitoringpb.DeleteAlertPolicyRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteAlertPolicy(ctx, req, opts...)
}

// UpdateAlertPolicy updates an alerting policy. You can either replace the entire policy with
// a new one or replace only certain fields in the current alerting policy by
// specifying the fields to be updated via updateMask. Returns the
// updated alerting policy.
func (c *AlertPolicyClient) UpdateAlertPolicy(ctx context.Context, req *monitoringpb.UpdateAlertPolicyRequest, opts ...gax.CallOption) (*monitoringpb.AlertPolicy, error) {
	return c.internalClient.UpdateAlertPolicy(ctx, req, opts...)
}

// alertPolicyGRPCClient is a client for interacting with Cloud Monitoring API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type alertPolicyGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing AlertPolicyClient
	CallOptions **AlertPolicyCallOptions

	// The gRPC API client.
	alertPolicyClient monitoringpb.AlertPolicyServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAlertPolicyClient creates a new alert policy service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The AlertPolicyService API is used to manage (list, create, delete,
// edit) alert policies in Cloud Monitoring. An alerting policy is
// a description of the conditions under which some aspect of your
// system is considered to be “unhealthy” and the ways to notify
// people or services about this state. In addition to using this API, alert
// policies can also be managed through
// Cloud Monitoring (at https://cloud.google.com/monitoring/docs/),
// which can be reached by clicking the “Monitoring” tab in
// Cloud console (at https://console.cloud.google.com/).
func NewAlertPolicyClient(ctx context.Context, opts ...option.ClientOption) (*AlertPolicyClient, error) {
	clientOpts := defaultAlertPolicyGRPCClientOptions()
	if newAlertPolicyClientHook != nil {
		hookOpts, err := newAlertPolicyClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AlertPolicyClient{CallOptions: defaultAlertPolicyCallOptions()}

	c := &alertPolicyGRPCClient{
		connPool:          connPool,
		disableDeadlines:  disableDeadlines,
		alertPolicyClient: monitoringpb.NewAlertPolicyServiceClient(connPool),
		CallOptions:       &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *alertPolicyGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *alertPolicyGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *alertPolicyGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *alertPolicyGRPCClient) ListAlertPolicies(ctx context.Context, req *monitoringpb.ListAlertPoliciesRequest, opts ...gax.CallOption) *AlertPolicyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListAlertPolicies[0:len((*c.CallOptions).ListAlertPolicies):len((*c.CallOptions).ListAlertPolicies)], opts...)
	it := &AlertPolicyIterator{}
	req = proto.Clone(req).(*monitoringpb.ListAlertPoliciesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*monitoringpb.AlertPolicy, string, error) {
		resp := &monitoringpb.ListAlertPoliciesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.alertPolicyClient.ListAlertPolicies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetAlertPolicies(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *alertPolicyGRPCClient) GetAlertPolicy(ctx context.Context, req *monitoringpb.GetAlertPolicyRequest, opts ...gax.CallOption) (*monitoringpb.AlertPolicy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAlertPolicy[0:len((*c.CallOptions).GetAlertPolicy):len((*c.CallOptions).GetAlertPolicy)], opts...)
	var resp *monitoringpb.AlertPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.alertPolicyClient.GetAlertPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *alertPolicyGRPCClient) CreateAlertPolicy(ctx context.Context, req *monitoringpb.CreateAlertPolicyRequest, opts ...gax.CallOption) (*monitoringpb.AlertPolicy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateAlertPolicy[0:len((*c.CallOptions).CreateAlertPolicy):len((*c.CallOptions).CreateAlertPolicy)], opts...)
	var resp *monitoringpb.AlertPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.alertPolicyClient.CreateAlertPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *alertPolicyGRPCClient) DeleteAlertPolicy(ctx context.Context, req *monitoringpb.DeleteAlertPolicyRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteAlertPolicy[0:len((*c.CallOptions).DeleteAlertPolicy):len((*c.CallOptions).DeleteAlertPolicy)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.alertPolicyClient.DeleteAlertPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *alertPolicyGRPCClient) UpdateAlertPolicy(ctx context.Context, req *monitoringpb.UpdateAlertPolicyRequest, opts ...gax.CallOption) (*monitoringpb.AlertPolicy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "alert_policy.name", url.QueryEscape(req.GetAlertPolicy().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateAlertPolicy[0:len((*c.CallOptions).UpdateAlertPolicy):len((*c.CallOptions).UpdateAlertPolicy)], opts...)
	var resp *monitoringpb.AlertPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.alertPolicyClient.UpdateAlertPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AlertPolicyIterator manages a stream of *monitoringpb.AlertPolicy.
type AlertPolicyIterator struct {
	items    []*monitoringpb.AlertPolicy
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*monitoringpb.AlertPolicy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AlertPolicyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AlertPolicyIterator) Next() (*monitoringpb.AlertPolicy, error) {
	var item *monitoringpb.AlertPolicy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AlertPolicyIterator) bufLen() int {
	return len(it.items)
}

func (it *AlertPolicyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
