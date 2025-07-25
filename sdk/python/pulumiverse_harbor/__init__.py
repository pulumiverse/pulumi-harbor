# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
from . import _utilities
import typing
# Export this package's modules as members:
from .config_auth import *
from .config_security import *
from .config_system import *
from .garbage_collection import *
from .get_groups import *
from .get_project import *
from .get_project_member_groups import *
from .get_project_member_users import *
from .get_projects import *
from .get_registry import *
from .get_robot_accounts import *
from .get_users import *
from .group import *
from .immutable_tag_rule import *
from .interrogation_services import *
from .label import *
from .preheat_instance import *
from .project import *
from .project_member_group import *
from .project_member_user import *
from .project_webhook import *
from .provider import *
from .purge_audit_log import *
from .registry import *
from .replication import *
from .retention_policy import *
from .robot_account import *
from .tasks import *
from .user import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumiverse_harbor.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumiverse_harbor.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "harbor",
  "mod": "index/configAuth",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/configAuth:ConfigAuth": "ConfigAuth"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/configSecurity",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/configSecurity:ConfigSecurity": "ConfigSecurity"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/configSystem",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/configSystem:ConfigSystem": "ConfigSystem"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/garbageCollection",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/garbageCollection:GarbageCollection": "GarbageCollection"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/group",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/group:Group": "Group"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/immutableTagRule",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/immutableTagRule:ImmutableTagRule": "ImmutableTagRule"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/interrogationServices",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/interrogationServices:InterrogationServices": "InterrogationServices"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/label",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/label:Label": "Label"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/preheatInstance",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/preheatInstance:PreheatInstance": "PreheatInstance"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/project",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/project:Project": "Project"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/projectMemberGroup",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/projectMemberGroup:ProjectMemberGroup": "ProjectMemberGroup"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/projectMemberUser",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/projectMemberUser:ProjectMemberUser": "ProjectMemberUser"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/projectWebhook",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/projectWebhook:ProjectWebhook": "ProjectWebhook"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/purgeAuditLog",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/purgeAuditLog:PurgeAuditLog": "PurgeAuditLog"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/registry",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/registry:Registry": "Registry"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/replication",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/replication:Replication": "Replication"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/retentionPolicy",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/retentionPolicy:RetentionPolicy": "RetentionPolicy"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/robotAccount",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/robotAccount:RobotAccount": "RobotAccount"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/tasks",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/tasks:Tasks": "Tasks"
  }
 },
 {
  "pkg": "harbor",
  "mod": "index/user",
  "fqn": "pulumiverse_harbor",
  "classes": {
   "harbor:index/user:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "harbor",
  "token": "pulumi:providers:harbor",
  "fqn": "pulumiverse_harbor",
  "class": "Provider"
 }
]
"""
)
