# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RetentionPolicyArgs', 'RetentionPolicy']

@pulumi.input_type
class RetentionPolicyArgs:
    def __init__(__self__, *,
                 rules: pulumi.Input[Sequence[pulumi.Input['RetentionPolicyRuleArgs']]],
                 scope: pulumi.Input[str],
                 schedule: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RetentionPolicy resource.
        :param pulumi.Input[str] scope: The project id of which you would like to apply this policy.
        :param pulumi.Input[str] schedule: The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        """
        pulumi.set(__self__, "rules", rules)
        pulumi.set(__self__, "scope", scope)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Input[Sequence[pulumi.Input['RetentionPolicyRuleArgs']]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: pulumi.Input[Sequence[pulumi.Input['RetentionPolicyRuleArgs']]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input[str]:
        """
        The project id of which you would like to apply this policy.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)


@pulumi.input_type
class _RetentionPolicyState:
    def __init__(__self__, *,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['RetentionPolicyRuleArgs']]]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RetentionPolicy resources.
        :param pulumi.Input[str] schedule: The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        :param pulumi.Input[str] scope: The project id of which you would like to apply this policy.
        """
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if schedule is not None:
            pulumi.set(__self__, "schedule", schedule)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RetentionPolicyRuleArgs']]]]:
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RetentionPolicyRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def schedule(self) -> Optional[pulumi.Input[str]]:
        """
        The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        """
        return pulumi.get(self, "schedule")

    @schedule.setter
    def schedule(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "schedule", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input[str]]:
        """
        The project id of which you would like to apply this policy.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scope", value)


class RetentionPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RetentionPolicyRuleArgs']]]]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import harbor:index/retentionPolicy:RetentionPolicy main /retentions/10
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] schedule: The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        :param pulumi.Input[str] scope: The project id of which you would like to apply this policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RetentionPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import harbor:index/retentionPolicy:RetentionPolicy main /retentions/10
        ```

        :param str resource_name: The name of the resource.
        :param RetentionPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RetentionPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RetentionPolicyRuleArgs']]]]] = None,
                 schedule: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RetentionPolicyArgs.__new__(RetentionPolicyArgs)

            if rules is None and not opts.urn:
                raise TypeError("Missing required property 'rules'")
            __props__.__dict__["rules"] = rules
            __props__.__dict__["schedule"] = schedule
            if scope is None and not opts.urn:
                raise TypeError("Missing required property 'scope'")
            __props__.__dict__["scope"] = scope
        super(RetentionPolicy, __self__).__init__(
            'harbor:index/retentionPolicy:RetentionPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RetentionPolicyRuleArgs']]]]] = None,
            schedule: Optional[pulumi.Input[str]] = None,
            scope: Optional[pulumi.Input[str]] = None) -> 'RetentionPolicy':
        """
        Get an existing RetentionPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] schedule: The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        :param pulumi.Input[str] scope: The project id of which you would like to apply this policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RetentionPolicyState.__new__(_RetentionPolicyState)

        __props__.__dict__["rules"] = rules
        __props__.__dict__["schedule"] = schedule
        __props__.__dict__["scope"] = scope
        return RetentionPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Sequence['outputs.RetentionPolicyRule']]:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    def schedule(self) -> pulumi.Output[Optional[str]]:
        """
        The schedule of when you would like the policy to run. This can be `Hourly`, `Daily`, `Weekly` or can be a custom cron string.
        """
        return pulumi.get(self, "schedule")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        The project id of which you would like to apply this policy.
        """
        return pulumi.get(self, "scope")

