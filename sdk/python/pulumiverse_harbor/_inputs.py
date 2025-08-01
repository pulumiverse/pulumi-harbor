# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'ConfigSystemBannerMessageArgs',
    'ConfigSystemBannerMessageArgsDict',
    'ReplicationFilterArgs',
    'ReplicationFilterArgsDict',
    'RetentionPolicyRuleArgs',
    'RetentionPolicyRuleArgsDict',
    'RobotAccountPermissionArgs',
    'RobotAccountPermissionArgsDict',
    'RobotAccountPermissionAccessArgs',
    'RobotAccountPermissionAccessArgsDict',
]

MYPY = False

if not MYPY:
    class ConfigSystemBannerMessageArgsDict(TypedDict):
        message: pulumi.Input[builtins.str]
        """
        The message to display in the banner.
        """
        closable: NotRequired[pulumi.Input[builtins.bool]]
        """
        Whether or not the banner message is closable.
        """
        from_date: NotRequired[pulumi.Input[builtins.str]]
        """
        The date the banner message will start displaying. (Format: `MM/DD/YYYY`)
        """
        to_date: NotRequired[pulumi.Input[builtins.str]]
        """
        The date the banner message will stop displaying. (Format: `MM/DD/YYYY`)
        """
        type: NotRequired[pulumi.Input[builtins.str]]
        """
        The type of banner message. Can be `"info"`, `"warning"`, `"success"` or `"danger"`.
        """
elif False:
    ConfigSystemBannerMessageArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ConfigSystemBannerMessageArgs:
    def __init__(__self__, *,
                 message: pulumi.Input[builtins.str],
                 closable: Optional[pulumi.Input[builtins.bool]] = None,
                 from_date: Optional[pulumi.Input[builtins.str]] = None,
                 to_date: Optional[pulumi.Input[builtins.str]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] message: The message to display in the banner.
        :param pulumi.Input[builtins.bool] closable: Whether or not the banner message is closable.
        :param pulumi.Input[builtins.str] from_date: The date the banner message will start displaying. (Format: `MM/DD/YYYY`)
        :param pulumi.Input[builtins.str] to_date: The date the banner message will stop displaying. (Format: `MM/DD/YYYY`)
        :param pulumi.Input[builtins.str] type: The type of banner message. Can be `"info"`, `"warning"`, `"success"` or `"danger"`.
        """
        pulumi.set(__self__, "message", message)
        if closable is not None:
            pulumi.set(__self__, "closable", closable)
        if from_date is not None:
            pulumi.set(__self__, "from_date", from_date)
        if to_date is not None:
            pulumi.set(__self__, "to_date", to_date)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def message(self) -> pulumi.Input[builtins.str]:
        """
        The message to display in the banner.
        """
        return pulumi.get(self, "message")

    @message.setter
    def message(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "message", value)

    @property
    @pulumi.getter
    def closable(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether or not the banner message is closable.
        """
        return pulumi.get(self, "closable")

    @closable.setter
    def closable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "closable", value)

    @property
    @pulumi.getter(name="fromDate")
    def from_date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date the banner message will start displaying. (Format: `MM/DD/YYYY`)
        """
        return pulumi.get(self, "from_date")

    @from_date.setter
    def from_date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "from_date", value)

    @property
    @pulumi.getter(name="toDate")
    def to_date(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date the banner message will stop displaying. (Format: `MM/DD/YYYY`)
        """
        return pulumi.get(self, "to_date")

    @to_date.setter
    def to_date(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "to_date", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The type of banner message. Can be `"info"`, `"warning"`, `"success"` or `"danger"`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


if not MYPY:
    class ReplicationFilterArgsDict(TypedDict):
        decoration: NotRequired[pulumi.Input[builtins.str]]
        """
        Matches or excludes the result. Can be one of the following. `matches`, `excludes`
        """
        labels: NotRequired[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]
        """
        Filter on the resource according to labels.
        """
        name: NotRequired[pulumi.Input[builtins.str]]
        """
        Filter on the name of the resource.
        """
        resource: NotRequired[pulumi.Input[builtins.str]]
        """
        Filter on the resource type. Can be one of the following types. `chart`, `artifact`
        """
        tag: NotRequired[pulumi.Input[builtins.str]]
        """
        Filter on the tag/version of the resource.
        """
elif False:
    ReplicationFilterArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class ReplicationFilterArgs:
    def __init__(__self__, *,
                 decoration: Optional[pulumi.Input[builtins.str]] = None,
                 labels: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 resource: Optional[pulumi.Input[builtins.str]] = None,
                 tag: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] decoration: Matches or excludes the result. Can be one of the following. `matches`, `excludes`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] labels: Filter on the resource according to labels.
        :param pulumi.Input[builtins.str] name: Filter on the name of the resource.
        :param pulumi.Input[builtins.str] resource: Filter on the resource type. Can be one of the following types. `chart`, `artifact`
        :param pulumi.Input[builtins.str] tag: Filter on the tag/version of the resource.
        """
        if decoration is not None:
            pulumi.set(__self__, "decoration", decoration)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if resource is not None:
            pulumi.set(__self__, "resource", resource)
        if tag is not None:
            pulumi.set(__self__, "tag", tag)

    @property
    @pulumi.getter
    def decoration(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Matches or excludes the result. Can be one of the following. `matches`, `excludes`
        """
        return pulumi.get(self, "decoration")

    @decoration.setter
    def decoration(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "decoration", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Filter on the resource according to labels.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Filter on the name of the resource.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def resource(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Filter on the resource type. Can be one of the following types. `chart`, `artifact`
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter
    def tag(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Filter on the tag/version of the resource.
        """
        return pulumi.get(self, "tag")

    @tag.setter
    def tag(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tag", value)


if not MYPY:
    class RetentionPolicyRuleArgsDict(TypedDict):
        always_retain: NotRequired[pulumi.Input[builtins.bool]]
        disabled: NotRequired[pulumi.Input[builtins.bool]]
        most_recently_pulled: NotRequired[pulumi.Input[builtins.int]]
        most_recently_pushed: NotRequired[pulumi.Input[builtins.int]]
        n_days_since_last_pull: NotRequired[pulumi.Input[builtins.int]]
        n_days_since_last_push: NotRequired[pulumi.Input[builtins.int]]
        repo_excluding: NotRequired[pulumi.Input[builtins.str]]
        repo_matching: NotRequired[pulumi.Input[builtins.str]]
        tag_excluding: NotRequired[pulumi.Input[builtins.str]]
        tag_matching: NotRequired[pulumi.Input[builtins.str]]
        untagged_artifacts: NotRequired[pulumi.Input[builtins.bool]]
elif False:
    RetentionPolicyRuleArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class RetentionPolicyRuleArgs:
    def __init__(__self__, *,
                 always_retain: Optional[pulumi.Input[builtins.bool]] = None,
                 disabled: Optional[pulumi.Input[builtins.bool]] = None,
                 most_recently_pulled: Optional[pulumi.Input[builtins.int]] = None,
                 most_recently_pushed: Optional[pulumi.Input[builtins.int]] = None,
                 n_days_since_last_pull: Optional[pulumi.Input[builtins.int]] = None,
                 n_days_since_last_push: Optional[pulumi.Input[builtins.int]] = None,
                 repo_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 repo_matching: Optional[pulumi.Input[builtins.str]] = None,
                 tag_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 tag_matching: Optional[pulumi.Input[builtins.str]] = None,
                 untagged_artifacts: Optional[pulumi.Input[builtins.bool]] = None):
        if always_retain is not None:
            pulumi.set(__self__, "always_retain", always_retain)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if most_recently_pulled is not None:
            pulumi.set(__self__, "most_recently_pulled", most_recently_pulled)
        if most_recently_pushed is not None:
            pulumi.set(__self__, "most_recently_pushed", most_recently_pushed)
        if n_days_since_last_pull is not None:
            pulumi.set(__self__, "n_days_since_last_pull", n_days_since_last_pull)
        if n_days_since_last_push is not None:
            pulumi.set(__self__, "n_days_since_last_push", n_days_since_last_push)
        if repo_excluding is not None:
            pulumi.set(__self__, "repo_excluding", repo_excluding)
        if repo_matching is not None:
            pulumi.set(__self__, "repo_matching", repo_matching)
        if tag_excluding is not None:
            pulumi.set(__self__, "tag_excluding", tag_excluding)
        if tag_matching is not None:
            pulumi.set(__self__, "tag_matching", tag_matching)
        if untagged_artifacts is not None:
            pulumi.set(__self__, "untagged_artifacts", untagged_artifacts)

    @property
    @pulumi.getter(name="alwaysRetain")
    def always_retain(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "always_retain")

    @always_retain.setter
    def always_retain(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "always_retain", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="mostRecentlyPulled")
    def most_recently_pulled(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "most_recently_pulled")

    @most_recently_pulled.setter
    def most_recently_pulled(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "most_recently_pulled", value)

    @property
    @pulumi.getter(name="mostRecentlyPushed")
    def most_recently_pushed(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "most_recently_pushed")

    @most_recently_pushed.setter
    def most_recently_pushed(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "most_recently_pushed", value)

    @property
    @pulumi.getter(name="nDaysSinceLastPull")
    def n_days_since_last_pull(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "n_days_since_last_pull")

    @n_days_since_last_pull.setter
    def n_days_since_last_pull(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "n_days_since_last_pull", value)

    @property
    @pulumi.getter(name="nDaysSinceLastPush")
    def n_days_since_last_push(self) -> Optional[pulumi.Input[builtins.int]]:
        return pulumi.get(self, "n_days_since_last_push")

    @n_days_since_last_push.setter
    def n_days_since_last_push(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "n_days_since_last_push", value)

    @property
    @pulumi.getter(name="repoExcluding")
    def repo_excluding(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "repo_excluding")

    @repo_excluding.setter
    def repo_excluding(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repo_excluding", value)

    @property
    @pulumi.getter(name="repoMatching")
    def repo_matching(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "repo_matching")

    @repo_matching.setter
    def repo_matching(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repo_matching", value)

    @property
    @pulumi.getter(name="tagExcluding")
    def tag_excluding(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "tag_excluding")

    @tag_excluding.setter
    def tag_excluding(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tag_excluding", value)

    @property
    @pulumi.getter(name="tagMatching")
    def tag_matching(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "tag_matching")

    @tag_matching.setter
    def tag_matching(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tag_matching", value)

    @property
    @pulumi.getter(name="untaggedArtifacts")
    def untagged_artifacts(self) -> Optional[pulumi.Input[builtins.bool]]:
        return pulumi.get(self, "untagged_artifacts")

    @untagged_artifacts.setter
    def untagged_artifacts(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "untagged_artifacts", value)


if not MYPY:
    class RobotAccountPermissionArgsDict(TypedDict):
        accesses: pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionAccessArgsDict']]]
        kind: pulumi.Input[builtins.str]
        """
        Either `system` or `project`.
        """
        namespace: pulumi.Input[builtins.str]
        """
        namespace is the name of your project. For kind `system` permissions, always use `/` as namespace. Use `*` to match all projects.
        """
elif False:
    RobotAccountPermissionArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class RobotAccountPermissionArgs:
    def __init__(__self__, *,
                 accesses: pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionAccessArgs']]],
                 kind: pulumi.Input[builtins.str],
                 namespace: pulumi.Input[builtins.str]):
        """
        :param pulumi.Input[builtins.str] kind: Either `system` or `project`.
        :param pulumi.Input[builtins.str] namespace: namespace is the name of your project. For kind `system` permissions, always use `/` as namespace. Use `*` to match all projects.
        """
        pulumi.set(__self__, "accesses", accesses)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def accesses(self) -> pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionAccessArgs']]]:
        return pulumi.get(self, "accesses")

    @accesses.setter
    def accesses(self, value: pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionAccessArgs']]]):
        pulumi.set(self, "accesses", value)

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Input[builtins.str]:
        """
        Either `system` or `project`.
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[builtins.str]:
        """
        namespace is the name of your project. For kind `system` permissions, always use `/` as namespace. Use `*` to match all projects.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "namespace", value)


if not MYPY:
    class RobotAccountPermissionAccessArgsDict(TypedDict):
        action: pulumi.Input[builtins.str]
        """
        Eg. `push`, `pull`, `read`, etc. Check [available actions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
        """
        resource: pulumi.Input[builtins.str]
        """
        Eg. `repository`, `labels`, etc. Check [available resources](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
        """
        effect: NotRequired[pulumi.Input[builtins.str]]
        """
        Either `allow` or `deny`. Defaults to `allow`.
        """
elif False:
    RobotAccountPermissionAccessArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class RobotAccountPermissionAccessArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[builtins.str],
                 resource: pulumi.Input[builtins.str],
                 effect: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] action: Eg. `push`, `pull`, `read`, etc. Check [available actions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
        :param pulumi.Input[builtins.str] resource: Eg. `repository`, `labels`, etc. Check [available resources](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
        :param pulumi.Input[builtins.str] effect: Either `allow` or `deny`. Defaults to `allow`.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "resource", resource)
        if effect is not None:
            pulumi.set(__self__, "effect", effect)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[builtins.str]:
        """
        Eg. `push`, `pull`, `read`, etc. Check [available actions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Input[builtins.str]:
        """
        Eg. `repository`, `labels`, etc. Check [available resources](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go).
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter
    def effect(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Either `allow` or `deny`. Defaults to `allow`.
        """
        return pulumi.get(self, "effect")

    @effect.setter
    def effect(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "effect", value)


