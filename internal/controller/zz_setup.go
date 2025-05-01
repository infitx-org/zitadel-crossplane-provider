// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	oidc "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/application/oidc"
	user "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/human/user"
	member "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/instance/member"
	usermachine "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/machine/user"
	memberorg "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/org/member"
	role "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/project/role"
	providerconfig "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/providerconfig"
	actions "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/trigger/actions"
	grant "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/user/grant"
	action "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/zitadel/action"
	org "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/zitadel/org"
	project "github.com/infitx-org/zitadel-crossplane-provider/internal/controller/zitadel/project"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		oidc.Setup,
		user.Setup,
		member.Setup,
		usermachine.Setup,
		memberorg.Setup,
		role.Setup,
		providerconfig.Setup,
		actions.Setup,
		grant.Setup,
		action.Setup,
		org.Setup,
		project.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
