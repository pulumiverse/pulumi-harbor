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

__all__ = ['ImmutableTagRuleArgs', 'ImmutableTagRule']

@pulumi.input_type
class ImmutableTagRuleArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[builtins.str],
                 disabled: Optional[pulumi.Input[builtins.bool]] = None,
                 repo_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 repo_matching: Optional[pulumi.Input[builtins.str]] = None,
                 tag_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 tag_matching: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ImmutableTagRule resource.
        :param pulumi.Input[builtins.str] project_id: The project id of which you would like to apply this policy.
        :param pulumi.Input[builtins.bool] disabled: Specify if the rule is disable or not. Defaults to `false`
        :param pulumi.Input[builtins.str] repo_excluding: For the repositories excluding.
        :param pulumi.Input[builtins.str] repo_matching: For the repositories matching.
        :param pulumi.Input[builtins.str] tag_excluding: For the tag excluding.
        :param pulumi.Input[builtins.str] tag_matching: For the tag matching.
        """
        pulumi.set(__self__, "project_id", project_id)
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if repo_excluding is not None:
            pulumi.set(__self__, "repo_excluding", repo_excluding)
        if repo_matching is not None:
            pulumi.set(__self__, "repo_matching", repo_matching)
        if tag_excluding is not None:
            pulumi.set(__self__, "tag_excluding", tag_excluding)
        if tag_matching is not None:
            pulumi.set(__self__, "tag_matching", tag_matching)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[builtins.str]:
        """
        The project id of which you would like to apply this policy.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specify if the rule is disable or not. Defaults to `false`
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="repoExcluding")
    def repo_excluding(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For the repositories excluding.
        """
        return pulumi.get(self, "repo_excluding")

    @repo_excluding.setter
    def repo_excluding(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repo_excluding", value)

    @property
    @pulumi.getter(name="repoMatching")
    def repo_matching(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For the repositories matching.
        """
        return pulumi.get(self, "repo_matching")

    @repo_matching.setter
    def repo_matching(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repo_matching", value)

    @property
    @pulumi.getter(name="tagExcluding")
    def tag_excluding(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For the tag excluding.
        """
        return pulumi.get(self, "tag_excluding")

    @tag_excluding.setter
    def tag_excluding(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tag_excluding", value)

    @property
    @pulumi.getter(name="tagMatching")
    def tag_matching(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For the tag matching.
        """
        return pulumi.get(self, "tag_matching")

    @tag_matching.setter
    def tag_matching(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tag_matching", value)


@pulumi.input_type
class _ImmutableTagRuleState:
    def __init__(__self__, *,
                 disabled: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 repo_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 repo_matching: Optional[pulumi.Input[builtins.str]] = None,
                 tag_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 tag_matching: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ImmutableTagRule resources.
        :param pulumi.Input[builtins.bool] disabled: Specify if the rule is disable or not. Defaults to `false`
        :param pulumi.Input[builtins.str] project_id: The project id of which you would like to apply this policy.
        :param pulumi.Input[builtins.str] repo_excluding: For the repositories excluding.
        :param pulumi.Input[builtins.str] repo_matching: For the repositories matching.
        :param pulumi.Input[builtins.str] tag_excluding: For the tag excluding.
        :param pulumi.Input[builtins.str] tag_matching: For the tag matching.
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if repo_excluding is not None:
            pulumi.set(__self__, "repo_excluding", repo_excluding)
        if repo_matching is not None:
            pulumi.set(__self__, "repo_matching", repo_matching)
        if tag_excluding is not None:
            pulumi.set(__self__, "tag_excluding", tag_excluding)
        if tag_matching is not None:
            pulumi.set(__self__, "tag_matching", tag_matching)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specify if the rule is disable or not. Defaults to `false`
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The project id of which you would like to apply this policy.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="repoExcluding")
    def repo_excluding(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For the repositories excluding.
        """
        return pulumi.get(self, "repo_excluding")

    @repo_excluding.setter
    def repo_excluding(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repo_excluding", value)

    @property
    @pulumi.getter(name="repoMatching")
    def repo_matching(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For the repositories matching.
        """
        return pulumi.get(self, "repo_matching")

    @repo_matching.setter
    def repo_matching(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "repo_matching", value)

    @property
    @pulumi.getter(name="tagExcluding")
    def tag_excluding(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For the tag excluding.
        """
        return pulumi.get(self, "tag_excluding")

    @tag_excluding.setter
    def tag_excluding(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tag_excluding", value)

    @property
    @pulumi.getter(name="tagMatching")
    def tag_matching(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        For the tag matching.
        """
        return pulumi.get(self, "tag_matching")

    @tag_matching.setter
    def tag_matching(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "tag_matching", value)


@pulumi.type_token("harbor:index/immutableTagRule:ImmutableTagRule")
class ImmutableTagRule(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 repo_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 repo_matching: Optional[pulumi.Input[builtins.str]] = None,
                 tag_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 tag_matching: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import harbor:index/immutableTagRule:ImmutableTagRule main /projects/4/immutabletagrules/25
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] disabled: Specify if the rule is disable or not. Defaults to `false`
        :param pulumi.Input[builtins.str] project_id: The project id of which you would like to apply this policy.
        :param pulumi.Input[builtins.str] repo_excluding: For the repositories excluding.
        :param pulumi.Input[builtins.str] repo_matching: For the repositories matching.
        :param pulumi.Input[builtins.str] tag_excluding: For the tag excluding.
        :param pulumi.Input[builtins.str] tag_matching: For the tag matching.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ImmutableTagRuleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import harbor:index/immutableTagRule:ImmutableTagRule main /projects/4/immutabletagrules/25
        ```

        :param str resource_name: The name of the resource.
        :param ImmutableTagRuleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ImmutableTagRuleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[builtins.bool]] = None,
                 project_id: Optional[pulumi.Input[builtins.str]] = None,
                 repo_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 repo_matching: Optional[pulumi.Input[builtins.str]] = None,
                 tag_excluding: Optional[pulumi.Input[builtins.str]] = None,
                 tag_matching: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ImmutableTagRuleArgs.__new__(ImmutableTagRuleArgs)

            __props__.__dict__["disabled"] = disabled
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["repo_excluding"] = repo_excluding
            __props__.__dict__["repo_matching"] = repo_matching
            __props__.__dict__["tag_excluding"] = tag_excluding
            __props__.__dict__["tag_matching"] = tag_matching
        super(ImmutableTagRule, __self__).__init__(
            'harbor:index/immutableTagRule:ImmutableTagRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disabled: Optional[pulumi.Input[builtins.bool]] = None,
            project_id: Optional[pulumi.Input[builtins.str]] = None,
            repo_excluding: Optional[pulumi.Input[builtins.str]] = None,
            repo_matching: Optional[pulumi.Input[builtins.str]] = None,
            tag_excluding: Optional[pulumi.Input[builtins.str]] = None,
            tag_matching: Optional[pulumi.Input[builtins.str]] = None) -> 'ImmutableTagRule':
        """
        Get an existing ImmutableTagRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] disabled: Specify if the rule is disable or not. Defaults to `false`
        :param pulumi.Input[builtins.str] project_id: The project id of which you would like to apply this policy.
        :param pulumi.Input[builtins.str] repo_excluding: For the repositories excluding.
        :param pulumi.Input[builtins.str] repo_matching: For the repositories matching.
        :param pulumi.Input[builtins.str] tag_excluding: For the tag excluding.
        :param pulumi.Input[builtins.str] tag_matching: For the tag matching.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ImmutableTagRuleState.__new__(_ImmutableTagRuleState)

        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["repo_excluding"] = repo_excluding
        __props__.__dict__["repo_matching"] = repo_matching
        __props__.__dict__["tag_excluding"] = tag_excluding
        __props__.__dict__["tag_matching"] = tag_matching
        return ImmutableTagRule(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specify if the rule is disable or not. Defaults to `false`
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[builtins.str]:
        """
        The project id of which you would like to apply this policy.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="repoExcluding")
    def repo_excluding(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        For the repositories excluding.
        """
        return pulumi.get(self, "repo_excluding")

    @property
    @pulumi.getter(name="repoMatching")
    def repo_matching(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        For the repositories matching.
        """
        return pulumi.get(self, "repo_matching")

    @property
    @pulumi.getter(name="tagExcluding")
    def tag_excluding(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        For the tag excluding.
        """
        return pulumi.get(self, "tag_excluding")

    @property
    @pulumi.getter(name="tagMatching")
    def tag_matching(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        For the tag matching.
        """
        return pulumi.get(self, "tag_matching")

