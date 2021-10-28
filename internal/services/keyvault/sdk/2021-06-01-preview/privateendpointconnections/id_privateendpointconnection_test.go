package privateendpointconnections

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = PrivateEndpointConnectionId{}

func TestPrivateEndpointConnectionIDFormatter(t *testing.T) {
	actual := NewPrivateEndpointConnectionID("{subscriptionId}", "{resourceGroupName}", "{vaultName}", "{privateEndpointConnectionName}").ID()
	expected := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if actual != expected {
		t.Fatalf("Expected %q but got %q", expected, actual)
	}
}

func TestParsePrivateEndpointConnectionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PrivateEndpointConnectionId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/{subscriptionId}/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/{subscriptionId}/resourceGroups/",
			Error: true,
		},

		{
			// missing VaultName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/",
			Error: true,
		},

		{
			// missing value for VaultName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/",
			Error: true,
		},

		{
			// missing Name
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/",
			Error: true,
		},

		{
			// missing value for Name
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}",
			Expected: &PrivateEndpointConnectionId{
				SubscriptionId: "{subscriptionId}",
				ResourceGroup:  "{resourceGroupName}",
				VaultName:      "{vaultName}",
				Name:           "{privateEndpointConnectionName}",
			},
		},

		{
			// upper-cased
			Input: "/SUBSCRIPTIONS/{SUBSCRIPTIONID}/RESOURCEGROUPS/{RESOURCEGROUPNAME}/PROVIDERS/MICROSOFT.KEYVAULT/VAULTS/{VAULTNAME}/PRIVATEENDPOINTCONNECTIONS/{PRIVATEENDPOINTCONNECTIONNAME}",
			Error: true,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePrivateEndpointConnectionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.VaultName != v.Expected.VaultName {
			t.Fatalf("Expected %q but got %q for VaultName", v.Expected.VaultName, actual.VaultName)
		}
		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
	}
}

func TestParsePrivateEndpointConnectionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *PrivateEndpointConnectionId
	}{

		{
			// empty
			Input: "",
			Error: true,
		},

		{
			// missing SubscriptionId
			Input: "/",
			Error: true,
		},

		{
			// missing value for SubscriptionId
			Input: "/subscriptions/",
			Error: true,
		},

		{
			// missing ResourceGroup
			Input: "/subscriptions/{subscriptionId}/",
			Error: true,
		},

		{
			// missing value for ResourceGroup
			Input: "/subscriptions/{subscriptionId}/resourceGroups/",
			Error: true,
		},

		{
			// missing VaultName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/",
			Error: true,
		},

		{
			// missing value for VaultName
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/",
			Error: true,
		},

		{
			// missing Name
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/",
			Error: true,
		},

		{
			// missing value for Name
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/",
			Error: true,
		},

		{
			// valid
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateEndpointConnections/{privateEndpointConnectionName}",
			Expected: &PrivateEndpointConnectionId{
				SubscriptionId: "{subscriptionId}",
				ResourceGroup:  "{resourceGroupName}",
				VaultName:      "{vaultName}",
				Name:           "{privateEndpointConnectionName}",
			},
		},

		{
			// lower-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/privateendpointconnections/{privateEndpointConnectionName}",
			Expected: &PrivateEndpointConnectionId{
				SubscriptionId: "{subscriptionId}",
				ResourceGroup:  "{resourceGroupName}",
				VaultName:      "{vaultName}",
				Name:           "{privateEndpointConnectionName}",
			},
		},

		{
			// upper-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/VAULTS/{vaultName}/PRIVATEENDPOINTCONNECTIONS/{privateEndpointConnectionName}",
			Expected: &PrivateEndpointConnectionId{
				SubscriptionId: "{subscriptionId}",
				ResourceGroup:  "{resourceGroupName}",
				VaultName:      "{vaultName}",
				Name:           "{privateEndpointConnectionName}",
			},
		},

		{
			// mixed-cased segment names
			Input: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/VaUlTs/{vaultName}/PrIvAtEeNdPoInTcOnNeCtIoNs/{privateEndpointConnectionName}",
			Expected: &PrivateEndpointConnectionId{
				SubscriptionId: "{subscriptionId}",
				ResourceGroup:  "{resourceGroupName}",
				VaultName:      "{vaultName}",
				Name:           "{privateEndpointConnectionName}",
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParsePrivateEndpointConnectionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %s", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.SubscriptionId != v.Expected.SubscriptionId {
			t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
		}
		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
		if actual.VaultName != v.Expected.VaultName {
			t.Fatalf("Expected %q but got %q for VaultName", v.Expected.VaultName, actual.VaultName)
		}
		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
	}
}
