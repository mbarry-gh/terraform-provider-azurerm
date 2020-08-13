package parse

import "fmt"

type BackendPoolId struct {
	ResourceGroup string
	FrontDoorName string
	Name          string
}

func NewBackendPoolID(id FrontDoorId, name string) BackendPoolId {
	return BackendPoolId{
		ResourceGroup: id.ResourceGroup,
		FrontDoorName: id.Name,
		Name:          name,
	}
}

func (id BackendPoolId) ID(subscriptionId string) string {
	base := NewFrontDoorID(id.ResourceGroup, id.FrontDoorName).ID(subscriptionId)
	return fmt.Sprintf("%s/backendPools/%s", base, id.Name)
}

func BackendPoolID(input string) (*BackendPoolId, error) {
	frontDoorId, id, err := parseFrontDoorChildResourceId(input)
	if err != nil {
		return nil, fmt.Errorf("parsing Backend Pool ID %q: %+v", input, err)
	}

	poolId := BackendPoolId{
		ResourceGroup: frontDoorId.ResourceGroup,
		FrontDoorName: frontDoorId.Name,
	}

	// TODO: handle this being case-insensitive
	// https://github.com/Azure/azure-sdk-for-go/issues/6762
	if poolId.Name, err = id.PopSegment("backendPools"); err != nil {
		return nil, err
	}

	if err := id.ValidateNoEmptySegments(input); err != nil {
		return nil, err
	}

	return &poolId, nil
}
