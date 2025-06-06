/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/infitx-org/zitadel-crossplane-provider/config/action"
	"github.com/infitx-org/zitadel-crossplane-provider/config/applicationoidc"
	"github.com/infitx-org/zitadel-crossplane-provider/config/humanuser"
	"github.com/infitx-org/zitadel-crossplane-provider/config/instancemember"
	"github.com/infitx-org/zitadel-crossplane-provider/config/machineuser"
	"github.com/infitx-org/zitadel-crossplane-provider/config/org"
	"github.com/infitx-org/zitadel-crossplane-provider/config/orgmember"
	"github.com/infitx-org/zitadel-crossplane-provider/config/project"
	"github.com/infitx-org/zitadel-crossplane-provider/config/projectrole"
	"github.com/infitx-org/zitadel-crossplane-provider/config/triggeractions"
	"github.com/infitx-org/zitadel-crossplane-provider/config/usergrant"
)

const (
	resourcePrefix = "zitadel"
	modulePath     = "github.com/infitx-org/zitadel-crossplane-provider"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("zitadel.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		humanuser.Configure,
		project.Configure,
		applicationoidc.Configure,
		usergrant.Configure,
		projectrole.Configure,
		machineuser.Configure,
		org.Configure,
		orgmember.Configure,
		instancemember.Configure,
		triggeractions.Configure,
		action.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
