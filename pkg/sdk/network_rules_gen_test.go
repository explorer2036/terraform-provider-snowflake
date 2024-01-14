package sdk

import "testing"

func TestNetworkRules_Create(t *testing.T) {
	id := RandomSchemaObjectIdentifier()

	defaultOpts := func() *CreateNetworkRuleOptions {
		return &CreateNetworkRuleOptions{
			name: id,
		}
	}

	t.Run("validation: nil options", func(t *testing.T) {
		opts := (*CreateForJavaFunctionOptions)(nil)
		assertOptsInvalidJoinedErrors(t, opts, ErrNilOptions)
	})

	t.Run("validation: incorrect identifier", func(t *testing.T) {
		opts := defaultOpts()
		opts.name = NewSchemaObjectIdentifier("", "", "")
		assertOptsInvalidJoinedErrors(t, opts, ErrInvalidObjectIdentifier)
	})

	t.Run("no value list", func(t *testing.T) {
	})

	t.Run("all options", func(t *testing.T) {
		opts := defaultOpts()
		opts.OrReplace = Bool(true)
		opts.NetworkIdentifierType = NetworkIdentifierTypePointer(NetworkIdentifierTypeIpv4)
		opts.NetworkIdentifiers = []NetworkIdentifier{
			{
				Value: "47.88.25.32/27",
			},
		}
		opts.NetworkRuleMode = NetworkRuleModePointer(NetworkRuleModeEgress)
		opts.Comment = String("comment")
		assertOptsValidAndSQLEquals(t, opts, `CREATE OR REPLACE NETWORK RULE %s TYPE = IPV4 VALUE_LIST = ('47.88.25.32/27') MODE = EGRESS COMMENT = 'comment'`, id.FullyQualifiedName())
	})
}
