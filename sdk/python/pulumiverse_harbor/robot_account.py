# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from . import outputs
from ._inputs import *

__all__ = ['RobotAccountArgs', 'RobotAccount']

@pulumi.input_type
class RobotAccountArgs:
    def __init__(__self__, *,
                 level: pulumi.Input[str],
                 permissions: pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 disable: Optional[pulumi.Input[bool]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RobotAccount resource.
        :param pulumi.Input[str] level: Level of the robot account, currently either `system` or `project`.
        :param pulumi.Input[str] description: The description of the robot account will be displayed in harbor.
        :param pulumi.Input[bool] disable: Disables the robot account when set to `true`.
        :param pulumi.Input[int] duration: By default, the robot account will not expire. Set it to the amount of days until the account should expire.
        :param pulumi.Input[str] name: The name of the project that will be created in harbor.
        :param pulumi.Input[str] secret: The secret of the robot account used for authentication. Defaults to random generated string from Harbor.
        """
        pulumi.set(__self__, "level", level)
        pulumi.set(__self__, "permissions", permissions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable is not None:
            pulumi.set(__self__, "disable", disable)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter
    def level(self) -> pulumi.Input[str]:
        """
        Level of the robot account, currently either `system` or `project`.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: pulumi.Input[str]):
        pulumi.set(self, "level", value)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionArgs']]]:
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionArgs']]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the robot account will be displayed in harbor.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disable(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables the robot account when set to `true`.
        """
        return pulumi.get(self, "disable")

    @disable.setter
    def disable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        By default, the robot account will not expire. Set it to the amount of days until the account should expire.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the project that will be created in harbor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The secret of the robot account used for authentication. Defaults to random generated string from Harbor.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


@pulumi.input_type
class _RobotAccountState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 disable: Optional[pulumi.Input[bool]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionArgs']]]] = None,
                 robot_id: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RobotAccount resources.
        :param pulumi.Input[str] description: The description of the robot account will be displayed in harbor.
        :param pulumi.Input[bool] disable: Disables the robot account when set to `true`.
        :param pulumi.Input[int] duration: By default, the robot account will not expire. Set it to the amount of days until the account should expire.
        :param pulumi.Input[str] level: Level of the robot account, currently either `system` or `project`.
        :param pulumi.Input[str] name: The name of the project that will be created in harbor.
        :param pulumi.Input[str] secret: The secret of the robot account used for authentication. Defaults to random generated string from Harbor.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable is not None:
            pulumi.set(__self__, "disable", disable)
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if full_name is not None:
            pulumi.set(__self__, "full_name", full_name)
        if level is not None:
            pulumi.set(__self__, "level", level)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if robot_id is not None:
            pulumi.set(__self__, "robot_id", robot_id)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the robot account will be displayed in harbor.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def disable(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables the robot account when set to `true`.
        """
        return pulumi.get(self, "disable")

    @disable.setter
    def disable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable", value)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        By default, the robot account will not expire. Set it to the amount of days until the account should expire.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "full_name")

    @full_name.setter
    def full_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_name", value)

    @property
    @pulumi.getter
    def level(self) -> Optional[pulumi.Input[str]]:
        """
        Level of the robot account, currently either `system` or `project`.
        """
        return pulumi.get(self, "level")

    @level.setter
    def level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "level", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the project that will be created in harbor.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionArgs']]]]:
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RobotAccountPermissionArgs']]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="robotId")
    def robot_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "robot_id")

    @robot_id.setter
    def robot_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "robot_id", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The secret of the robot account used for authentication. Defaults to random generated string from Harbor.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


class RobotAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable: Optional[pulumi.Input[bool]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RobotAccountPermissionArgs', 'RobotAccountPermissionArgsDict']]]]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### System Level
        Introduced in harbor 2.2.0, system level robot accounts can have basically [all available permissions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go) in harbor and are not dependent on a single project.

        ### Global

        The above example, creates a system level robot account with permissions to
        - permission to create labels on system level
        - pull repository across all projects
        - push repository to project "my-project-name"

        ### Project

        Other than system level robot accounts, project level robot accounts can interact on project level only.
        The [available permissions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go) are mostly the same as for system level robots.

        The above example creates a project level robot account with permissions to
        - pull repository on project "main"
        - push repository on project "main"

        ## Import

        ```sh
        $ pulumi import harbor:index/robotAccount:RobotAccount system /robots/123
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the robot account will be displayed in harbor.
        :param pulumi.Input[bool] disable: Disables the robot account when set to `true`.
        :param pulumi.Input[int] duration: By default, the robot account will not expire. Set it to the amount of days until the account should expire.
        :param pulumi.Input[str] level: Level of the robot account, currently either `system` or `project`.
        :param pulumi.Input[str] name: The name of the project that will be created in harbor.
        :param pulumi.Input[str] secret: The secret of the robot account used for authentication. Defaults to random generated string from Harbor.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RobotAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### System Level
        Introduced in harbor 2.2.0, system level robot accounts can have basically [all available permissions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go) in harbor and are not dependent on a single project.

        ### Global

        The above example, creates a system level robot account with permissions to
        - permission to create labels on system level
        - pull repository across all projects
        - push repository to project "my-project-name"

        ### Project

        Other than system level robot accounts, project level robot accounts can interact on project level only.
        The [available permissions](https://github.com/goharbor/harbor/blob/-/src/common/rbac/const.go) are mostly the same as for system level robots.

        The above example creates a project level robot account with permissions to
        - pull repository on project "main"
        - push repository on project "main"

        ## Import

        ```sh
        $ pulumi import harbor:index/robotAccount:RobotAccount system /robots/123
        ```

        :param str resource_name: The name of the resource.
        :param RobotAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RobotAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable: Optional[pulumi.Input[bool]] = None,
                 duration: Optional[pulumi.Input[int]] = None,
                 level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RobotAccountPermissionArgs', 'RobotAccountPermissionArgsDict']]]]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RobotAccountArgs.__new__(RobotAccountArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["disable"] = disable
            __props__.__dict__["duration"] = duration
            if level is None and not opts.urn:
                raise TypeError("Missing required property 'level'")
            __props__.__dict__["level"] = level
            __props__.__dict__["name"] = name
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["secret"] = None if secret is None else pulumi.Output.secret(secret)
            __props__.__dict__["full_name"] = None
            __props__.__dict__["robot_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(RobotAccount, __self__).__init__(
            'harbor:index/robotAccount:RobotAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable: Optional[pulumi.Input[bool]] = None,
            duration: Optional[pulumi.Input[int]] = None,
            full_name: Optional[pulumi.Input[str]] = None,
            level: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['RobotAccountPermissionArgs', 'RobotAccountPermissionArgsDict']]]]] = None,
            robot_id: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None) -> 'RobotAccount':
        """
        Get an existing RobotAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the robot account will be displayed in harbor.
        :param pulumi.Input[bool] disable: Disables the robot account when set to `true`.
        :param pulumi.Input[int] duration: By default, the robot account will not expire. Set it to the amount of days until the account should expire.
        :param pulumi.Input[str] level: Level of the robot account, currently either `system` or `project`.
        :param pulumi.Input[str] name: The name of the project that will be created in harbor.
        :param pulumi.Input[str] secret: The secret of the robot account used for authentication. Defaults to random generated string from Harbor.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RobotAccountState.__new__(_RobotAccountState)

        __props__.__dict__["description"] = description
        __props__.__dict__["disable"] = disable
        __props__.__dict__["duration"] = duration
        __props__.__dict__["full_name"] = full_name
        __props__.__dict__["level"] = level
        __props__.__dict__["name"] = name
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["robot_id"] = robot_id
        __props__.__dict__["secret"] = secret
        return RobotAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the robot account will be displayed in harbor.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def disable(self) -> pulumi.Output[Optional[bool]]:
        """
        Disables the robot account when set to `true`.
        """
        return pulumi.get(self, "disable")

    @property
    @pulumi.getter
    def duration(self) -> pulumi.Output[Optional[int]]:
        """
        By default, the robot account will not expire. Set it to the amount of days until the account should expire.
        """
        return pulumi.get(self, "duration")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter
    def level(self) -> pulumi.Output[str]:
        """
        Level of the robot account, currently either `system` or `project`.
        """
        return pulumi.get(self, "level")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the project that will be created in harbor.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence['outputs.RobotAccountPermission']]:
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="robotId")
    def robot_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "robot_id")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        The secret of the robot account used for authentication. Defaults to random generated string from Harbor.
        """
        return pulumi.get(self, "secret")

