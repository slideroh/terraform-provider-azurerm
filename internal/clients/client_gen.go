package clients

// NOTE: this file is generated - manual changes will be overwritten.

import (
	"context"

	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
	containers "github.com/hashicorp/terraform-provider-azurerm/internal/services/containers/client"
	devcenter "github.com/hashicorp/terraform-provider-azurerm/internal/services/devcenter/client"
	loadtestservice "github.com/hashicorp/terraform-provider-azurerm/internal/services/loadtestservice/client"
	managedidentity "github.com/hashicorp/terraform-provider-azurerm/internal/services/managedidentity/client"
)

type autoClient struct {
	ContainerService *containers.AutoClient
	DevCenter        *devcenter.AutoClient
	LoadTestService  *loadtestservice.AutoClient
	ManagedIdentity  *managedidentity.AutoClient
}

func buildAutoClients(client *autoClient, o *common.ClientOptions) (err error) {

	if client.ContainerService, err = containers.NewClient(o); err != nil {
		return fmt.Errorf("building client for ContainerService: %+v", err)
	}

	if client.DevCenter, err = devcenter.NewClient(o); err != nil {
		return fmt.Errorf("building client for DevCenter: %+v", err)
	}

	if client.LoadTestService, err = loadtestservice.NewClient(o); err != nil {
		return fmt.Errorf("building client for LoadTestService: %+v", err)
	}

	if client.ManagedIdentity, err = managedidentity.NewClient(o); err != nil {
		return fmt.Errorf("building client for ManagedIdentity: %+v", err)
	}

	return nil
}
