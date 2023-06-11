# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ConfigSystemArgs', 'ConfigSystem']

@pulumi.input_type
class ConfigSystemArgs:
    def __init__(__self__, *,
                 project_creation_restriction: Optional[pulumi.Input[str]] = None,
                 read_only: Optional[pulumi.Input[bool]] = None,
                 robot_name_prefix: Optional[pulumi.Input[str]] = None,
                 robot_token_expiration: Optional[pulumi.Input[int]] = None,
                 scanner_skip_update_pulltime: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ConfigSystem resource.
        """
        if project_creation_restriction is not None:
            pulumi.set(__self__, "project_creation_restriction", project_creation_restriction)
        if read_only is not None:
            pulumi.set(__self__, "read_only", read_only)
        if robot_name_prefix is not None:
            pulumi.set(__self__, "robot_name_prefix", robot_name_prefix)
        if robot_token_expiration is not None:
            pulumi.set(__self__, "robot_token_expiration", robot_token_expiration)
        if scanner_skip_update_pulltime is not None:
            pulumi.set(__self__, "scanner_skip_update_pulltime", scanner_skip_update_pulltime)

    @property
    @pulumi.getter(name="projectCreationRestriction")
    def project_creation_restriction(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_creation_restriction")

    @project_creation_restriction.setter
    def project_creation_restriction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_creation_restriction", value)

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "read_only")

    @read_only.setter
    def read_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read_only", value)

    @property
    @pulumi.getter(name="robotNamePrefix")
    def robot_name_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "robot_name_prefix")

    @robot_name_prefix.setter
    def robot_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "robot_name_prefix", value)

    @property
    @pulumi.getter(name="robotTokenExpiration")
    def robot_token_expiration(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "robot_token_expiration")

    @robot_token_expiration.setter
    def robot_token_expiration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "robot_token_expiration", value)

    @property
    @pulumi.getter(name="scannerSkipUpdatePulltime")
    def scanner_skip_update_pulltime(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "scanner_skip_update_pulltime")

    @scanner_skip_update_pulltime.setter
    def scanner_skip_update_pulltime(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "scanner_skip_update_pulltime", value)


@pulumi.input_type
class _ConfigSystemState:
    def __init__(__self__, *,
                 project_creation_restriction: Optional[pulumi.Input[str]] = None,
                 read_only: Optional[pulumi.Input[bool]] = None,
                 robot_name_prefix: Optional[pulumi.Input[str]] = None,
                 robot_token_expiration: Optional[pulumi.Input[int]] = None,
                 scanner_skip_update_pulltime: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering ConfigSystem resources.
        """
        if project_creation_restriction is not None:
            pulumi.set(__self__, "project_creation_restriction", project_creation_restriction)
        if read_only is not None:
            pulumi.set(__self__, "read_only", read_only)
        if robot_name_prefix is not None:
            pulumi.set(__self__, "robot_name_prefix", robot_name_prefix)
        if robot_token_expiration is not None:
            pulumi.set(__self__, "robot_token_expiration", robot_token_expiration)
        if scanner_skip_update_pulltime is not None:
            pulumi.set(__self__, "scanner_skip_update_pulltime", scanner_skip_update_pulltime)

    @property
    @pulumi.getter(name="projectCreationRestriction")
    def project_creation_restriction(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project_creation_restriction")

    @project_creation_restriction.setter
    def project_creation_restriction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_creation_restriction", value)

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "read_only")

    @read_only.setter
    def read_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read_only", value)

    @property
    @pulumi.getter(name="robotNamePrefix")
    def robot_name_prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "robot_name_prefix")

    @robot_name_prefix.setter
    def robot_name_prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "robot_name_prefix", value)

    @property
    @pulumi.getter(name="robotTokenExpiration")
    def robot_token_expiration(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "robot_token_expiration")

    @robot_token_expiration.setter
    def robot_token_expiration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "robot_token_expiration", value)

    @property
    @pulumi.getter(name="scannerSkipUpdatePulltime")
    def scanner_skip_update_pulltime(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "scanner_skip_update_pulltime")

    @scanner_skip_update_pulltime.setter
    def scanner_skip_update_pulltime(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "scanner_skip_update_pulltime", value)


class ConfigSystem(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_creation_restriction: Optional[pulumi.Input[str]] = None,
                 read_only: Optional[pulumi.Input[bool]] = None,
                 robot_name_prefix: Optional[pulumi.Input[str]] = None,
                 robot_token_expiration: Optional[pulumi.Input[int]] = None,
                 scanner_skip_update_pulltime: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_harbor as harbor

        main = harbor.ConfigSystem("main",
            project_creation_restriction="adminonly",
            robot_name_prefix="harbor@",
            robot_token_expiration=30)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ConfigSystemArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_harbor as harbor

        main = harbor.ConfigSystem("main",
            project_creation_restriction="adminonly",
            robot_name_prefix="harbor@",
            robot_token_expiration=30)
        ```

        :param str resource_name: The name of the resource.
        :param ConfigSystemArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigSystemArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 project_creation_restriction: Optional[pulumi.Input[str]] = None,
                 read_only: Optional[pulumi.Input[bool]] = None,
                 robot_name_prefix: Optional[pulumi.Input[str]] = None,
                 robot_token_expiration: Optional[pulumi.Input[int]] = None,
                 scanner_skip_update_pulltime: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigSystemArgs.__new__(ConfigSystemArgs)

            __props__.__dict__["project_creation_restriction"] = project_creation_restriction
            __props__.__dict__["read_only"] = read_only
            __props__.__dict__["robot_name_prefix"] = robot_name_prefix
            __props__.__dict__["robot_token_expiration"] = robot_token_expiration
            __props__.__dict__["scanner_skip_update_pulltime"] = scanner_skip_update_pulltime
        super(ConfigSystem, __self__).__init__(
            'harbor:index/configSystem:ConfigSystem',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            project_creation_restriction: Optional[pulumi.Input[str]] = None,
            read_only: Optional[pulumi.Input[bool]] = None,
            robot_name_prefix: Optional[pulumi.Input[str]] = None,
            robot_token_expiration: Optional[pulumi.Input[int]] = None,
            scanner_skip_update_pulltime: Optional[pulumi.Input[bool]] = None) -> 'ConfigSystem':
        """
        Get an existing ConfigSystem resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigSystemState.__new__(_ConfigSystemState)

        __props__.__dict__["project_creation_restriction"] = project_creation_restriction
        __props__.__dict__["read_only"] = read_only
        __props__.__dict__["robot_name_prefix"] = robot_name_prefix
        __props__.__dict__["robot_token_expiration"] = robot_token_expiration
        __props__.__dict__["scanner_skip_update_pulltime"] = scanner_skip_update_pulltime
        return ConfigSystem(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="projectCreationRestriction")
    def project_creation_restriction(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "project_creation_restriction")

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "read_only")

    @property
    @pulumi.getter(name="robotNamePrefix")
    def robot_name_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "robot_name_prefix")

    @property
    @pulumi.getter(name="robotTokenExpiration")
    def robot_token_expiration(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "robot_token_expiration")

    @property
    @pulumi.getter(name="scannerSkipUpdatePulltime")
    def scanner_skip_update_pulltime(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "scanner_skip_update_pulltime")

