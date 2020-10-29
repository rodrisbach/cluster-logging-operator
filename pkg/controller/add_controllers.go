package controller

import (
	"github.com/rodrisbach/cluster-logging-operator/pkg/controller/clusterlogging"
	"github.com/rodrisbach/cluster-logging-operator/pkg/controller/collector"
	"github.com/rodrisbach/cluster-logging-operator/pkg/controller/forwarding"
	"github.com/rodrisbach/cluster-logging-operator/pkg/controller/proxyconfig"
	"github.com/rodrisbach/cluster-logging-operator/pkg/controller/trustedcabundle"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs,
		clusterlogging.Add,
		forwarding.Add,
		collector.Add,
		proxyconfig.Add,
		trustedcabundle.Add)
}
