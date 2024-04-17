// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.Id = ManagedHSMRoleDefinitionId{}

type ManagedHSMRoleDefinitionId struct {
	VaultBaseUrl string
	Scope        string
	Name         string
}

func NewManagedHSMRoleDefinitionID(hsmBaseUrl, scope string, name string) (*ManagedHSMRoleDefinitionId, error) {
	keyVaultUrl, err := url.Parse(hsmBaseUrl)
	if err != nil || hsmBaseUrl == "" {
		return nil, fmt.Errorf("parsing managedHSM nested itemID %q: %+v", hsmBaseUrl, err)
	}

	return &ManagedHSMRoleDefinitionId{
		VaultBaseUrl: keyVaultUrl.String(),
		Scope:        scope,
		Name:         name,
	}, nil
}

func (n ManagedHSMRoleDefinitionId) ID() string {
	// example: https://tharvey-keyvault.managedhsm.azure.net///uuid-idshifds-fks
	segments := []string{
		strings.TrimSuffix(n.VaultBaseUrl, "/"),
		n.Scope,
		"RoleDefinition",
		n.Name,
	}
	return strings.TrimSuffix(strings.Join(segments, "/"), "/")
}

func (n ManagedHSMRoleDefinitionId) String() string {
	return n.ID()
}

func ManagedHSMRoleDefinitionID(input string) (*ManagedHSMRoleDefinitionId, error) {
	idURL, err := url.ParseRequestURI(input)
	if err != nil {
		return nil, fmt.Errorf("cannot parse Azure KeyVault Child Id: %s", err)
	}

	path := idURL.Path

	path = strings.TrimPrefix(path, "/")
	path = strings.TrimSuffix(path, "/")

	nameSep := strings.LastIndex(path, "/")
	if nameSep <= 0 {
		return nil, fmt.Errorf("no name speparate exist in %s", input)
	}
	scope, name := path[:nameSep], path[nameSep+1:]

	typeSep := strings.LastIndex(scope, "/")
	if typeSep <= 0 {
		return nil, fmt.Errorf("no type speparate exist in %s", input)
	}
	scope, typ := scope[:typeSep], scope[typeSep+1:]
	if typ != "RoleDefinition" {
		return nil, fmt.Errorf("invalid type %s, must be 'RoleDefinition'", typ)
	}

	childId := ManagedHSMRoleDefinitionId{
		VaultBaseUrl: fmt.Sprintf("%s://%s/", idURL.Scheme, idURL.Host),
		Scope:        scope,
		Name:         name,
	}

	return &childId, nil
}
