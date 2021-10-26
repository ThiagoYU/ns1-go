package account

// APIKey wraps an NS1 /account/apikeys resource
type APIKey struct {
	// Read-only fields
	ID         string `json:"id,omitempty"`
	Key        string `json:"key,omitempty"`
	LastAccess int    `json:"last_access,omitempty"`

	Name              string          `json:"name"`
	TeamIDs           []string        `json:"teams,omitempty"`
	Permissions       *PermissionsMap `json:"permissions,omitempty"`
	IPWhitelist       []string        `json:"ip_whitelist,omitempty"`
	IPWhitelistStrict bool            `json:"ip_whitelist_strict"`
}
