package sdk

import (
	"context"
)

var _ NetworkRules = (*networkRules)(nil)

type networkRules struct {
	client *Client
}

func (v *networkRules) Create(ctx context.Context, request *CreateNetworkRuleRequest) error {
	opts := request.toOpts()
	return validateAndExec(v.client, ctx, opts)
}

func (r *CreateNetworkRuleRequest) toOpts() *CreateNetworkRuleOptions {
	opts := &CreateNetworkRuleOptions{
		OrReplace:             r.OrReplace,
		name:                  r.name,
		NetworkIdentifierType: r.NetworkIdentifierType,

		NetworkRuleMode: r.NetworkRuleMode,
		Comment:         r.Comment,
	}
	if r.NetworkIdentifiers != nil {
		s := make([]NetworkIdentifier, len(r.NetworkIdentifiers))
		for i, v := range r.NetworkIdentifiers {
			s[i] = NetworkIdentifier{
				Value: v.Value,
			}
		}
		opts.NetworkIdentifiers = s
	}
	return opts
}
