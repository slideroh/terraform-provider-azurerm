package vaults

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *ActionsRequired                        `json:"actionsRequired,omitempty"`
	Description     *string                                 `json:"description,omitempty"`
	Status          *PrivateEndpointServiceConnectionStatus `json:"status,omitempty"`
}
