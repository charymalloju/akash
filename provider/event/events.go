package event

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ovrclk/akash/manifest"
	dquery "github.com/ovrclk/akash/x/deployment/query"
	mtypes "github.com/ovrclk/akash/x/market/types"
)

// LeaseWon is the data structure that includes leaseID, group and price
type LeaseWon struct {
	LeaseID mtypes.LeaseID
	Group   *dquery.Group
	Price   sdk.Coin
}

// ManifestReceived stores leaseID, manifest received, deployment and group details
type ManifestReceived struct {
	LeaseID    mtypes.LeaseID
	Manifest   *manifest.Manifest
	Deployment *dquery.Deployment
	Group      *dquery.Group
}

// ManifestGroup returns group if present in manifest or nil
func (ev ManifestReceived) ManifestGroup() *manifest.Group {
	for _, mgroup := range *ev.Manifest {
		if mgroup.Name == ev.Group.Name {
			return &mgroup
		}
	}
	return nil
}

// ClusterDeploymentStatus represents status of the cluster deployment
type ClusterDeploymentStatus string

const (
	// ClusterDeploymentPending is used when cluster deployment status is pending
	ClusterDeploymentPending ClusterDeploymentStatus = "pending"
	// ClusterDeploymentDeployed is used when cluster deployment status is deployed
	ClusterDeploymentDeployed ClusterDeploymentStatus = "deployed"
)

// ClusterDeployment stores leaseID, group details and deployment status
type ClusterDeployment struct {
	LeaseID mtypes.LeaseID
	Group   *manifest.Group
	Status  ClusterDeploymentStatus
}
