// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harbor

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "harbor:index/configAuth:ConfigAuth":
		r = &ConfigAuth{}
	case "harbor:index/configEmail:ConfigEmail":
		r = &ConfigEmail{}
	case "harbor:index/configSecurity:ConfigSecurity":
		r = &ConfigSecurity{}
	case "harbor:index/configSystem:ConfigSystem":
		r = &ConfigSystem{}
	case "harbor:index/garbageCollection:GarbageCollection":
		r = &GarbageCollection{}
	case "harbor:index/group:Group":
		r = &Group{}
	case "harbor:index/immutableTagRule:ImmutableTagRule":
		r = &ImmutableTagRule{}
	case "harbor:index/interrogationServices:InterrogationServices":
		r = &InterrogationServices{}
	case "harbor:index/label:Label":
		r = &Label{}
	case "harbor:index/project:Project":
		r = &Project{}
	case "harbor:index/projectMemberGroup:ProjectMemberGroup":
		r = &ProjectMemberGroup{}
	case "harbor:index/projectMemberUser:ProjectMemberUser":
		r = &ProjectMemberUser{}
	case "harbor:index/projectWebhook:ProjectWebhook":
		r = &ProjectWebhook{}
	case "harbor:index/purgeAuditLog:PurgeAuditLog":
		r = &PurgeAuditLog{}
	case "harbor:index/registry:Registry":
		r = &Registry{}
	case "harbor:index/replication:Replication":
		r = &Replication{}
	case "harbor:index/retentionPolicy:RetentionPolicy":
		r = &RetentionPolicy{}
	case "harbor:index/robotAccount:RobotAccount":
		r = &RobotAccount{}
	case "harbor:index/tasks:Tasks":
		r = &Tasks{}
	case "harbor:index/user:User":
		r = &User{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:harbor" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"harbor",
		"index/configAuth",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/configEmail",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/configSecurity",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/configSystem",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/garbageCollection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/immutableTagRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/interrogationServices",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/label",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/projectMemberGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/projectMemberUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/projectWebhook",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/purgeAuditLog",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/registry",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/replication",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/retentionPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/robotAccount",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/tasks",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harbor",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"harbor",
		&pkg{version},
	)
}
