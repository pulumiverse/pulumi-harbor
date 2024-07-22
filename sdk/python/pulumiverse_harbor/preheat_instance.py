# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PreheatInstanceArgs', 'PreheatInstance']

@pulumi.input_type
class PreheatInstanceArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input[str],
                 vendor: pulumi.Input[str],
                 auth_mode: Optional[pulumi.Input[str]] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PreheatInstance resource.
        :param pulumi.Input[str] endpoint: The endpoint of the preheat instance.
        :param pulumi.Input[str] vendor: The vendor of the preheat instance. Must be either "dragonfly" or "kraken".
        :param pulumi.Input[str] auth_mode: The authentication mode for the preheat instance. Must be either "NONE", "BASIC", or "OAUTH". Defaults to "NONE".
        :param pulumi.Input[bool] default: Whether the preheat instance is the default instance. Defaults to false.
        :param pulumi.Input[str] description: The description of the preheat instance. Defaults to an empty string.
        :param pulumi.Input[bool] enabled: Whether the preheat instance is enabled. Defaults to true.
        :param pulumi.Input[bool] insecure: Whether to allow insecure connections to the preheat instance. Defaults to false.
        :param pulumi.Input[str] name: The name of the preheat instance.
        :param pulumi.Input[str] password: The password for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        :param pulumi.Input[str] token: The token for the preheat instance. Required if `auth_mode` is "OAUTH". Defaults to an empty string.
        :param pulumi.Input[str] username: The username for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "vendor", vendor)
        if auth_mode is not None:
            pulumi.set(__self__, "auth_mode", auth_mode)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        The endpoint of the preheat instance.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def vendor(self) -> pulumi.Input[str]:
        """
        The vendor of the preheat instance. Must be either "dragonfly" or "kraken".
        """
        return pulumi.get(self, "vendor")

    @vendor.setter
    def vendor(self, value: pulumi.Input[str]):
        pulumi.set(self, "vendor", value)

    @property
    @pulumi.getter(name="authMode")
    def auth_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication mode for the preheat instance. Must be either "NONE", "BASIC", or "OAUTH". Defaults to "NONE".
        """
        return pulumi.get(self, "auth_mode")

    @auth_mode.setter
    def auth_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_mode", value)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the preheat instance is the default instance. Defaults to false.
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the preheat instance. Defaults to an empty string.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the preheat instance is enabled. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to allow insecure connections to the preheat instance. Defaults to false.
        """
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the preheat instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token for the preheat instance. Required if `auth_mode` is "OAUTH". Defaults to an empty string.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _PreheatInstanceState:
    def __init__(__self__, *,
                 auth_mode: Optional[pulumi.Input[str]] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PreheatInstance resources.
        :param pulumi.Input[str] auth_mode: The authentication mode for the preheat instance. Must be either "NONE", "BASIC", or "OAUTH". Defaults to "NONE".
        :param pulumi.Input[bool] default: Whether the preheat instance is the default instance. Defaults to false.
        :param pulumi.Input[str] description: The description of the preheat instance. Defaults to an empty string.
        :param pulumi.Input[bool] enabled: Whether the preheat instance is enabled. Defaults to true.
        :param pulumi.Input[str] endpoint: The endpoint of the preheat instance.
        :param pulumi.Input[bool] insecure: Whether to allow insecure connections to the preheat instance. Defaults to false.
        :param pulumi.Input[str] name: The name of the preheat instance.
        :param pulumi.Input[str] password: The password for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        :param pulumi.Input[str] token: The token for the preheat instance. Required if `auth_mode` is "OAUTH". Defaults to an empty string.
        :param pulumi.Input[str] username: The username for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        :param pulumi.Input[str] vendor: The vendor of the preheat instance. Must be either "dragonfly" or "kraken".
        """
        if auth_mode is not None:
            pulumi.set(__self__, "auth_mode", auth_mode)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if vendor is not None:
            pulumi.set(__self__, "vendor", vendor)

    @property
    @pulumi.getter(name="authMode")
    def auth_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication mode for the preheat instance. Must be either "NONE", "BASIC", or "OAUTH". Defaults to "NONE".
        """
        return pulumi.get(self, "auth_mode")

    @auth_mode.setter
    def auth_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auth_mode", value)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the preheat instance is the default instance. Defaults to false.
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the preheat instance. Defaults to an empty string.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the preheat instance is enabled. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint of the preheat instance.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to allow insecure connections to the preheat instance. Defaults to false.
        """
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the preheat instance.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        The password for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token for the preheat instance. Required if `auth_mode` is "OAUTH". Defaults to an empty string.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def vendor(self) -> Optional[pulumi.Input[str]]:
        """
        The vendor of the preheat instance. Must be either "dragonfly" or "kraken".
        """
        return pulumi.get(self, "vendor")

    @vendor.setter
    def vendor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vendor", value)


class PreheatInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_mode: Optional[pulumi.Input[str]] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Basic Usage

        ### Usage with Authentication

        ## Import

        The `harbor_preheat_instance` resource can be imported using the preheat instance ID.

        ```sh
        $ pulumi import harbor:index/preheatInstance:PreheatInstance example /p2p/preheat/instances/example-preheat-instance
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_mode: The authentication mode for the preheat instance. Must be either "NONE", "BASIC", or "OAUTH". Defaults to "NONE".
        :param pulumi.Input[bool] default: Whether the preheat instance is the default instance. Defaults to false.
        :param pulumi.Input[str] description: The description of the preheat instance. Defaults to an empty string.
        :param pulumi.Input[bool] enabled: Whether the preheat instance is enabled. Defaults to true.
        :param pulumi.Input[str] endpoint: The endpoint of the preheat instance.
        :param pulumi.Input[bool] insecure: Whether to allow insecure connections to the preheat instance. Defaults to false.
        :param pulumi.Input[str] name: The name of the preheat instance.
        :param pulumi.Input[str] password: The password for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        :param pulumi.Input[str] token: The token for the preheat instance. Required if `auth_mode` is "OAUTH". Defaults to an empty string.
        :param pulumi.Input[str] username: The username for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        :param pulumi.Input[str] vendor: The vendor of the preheat instance. Must be either "dragonfly" or "kraken".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PreheatInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Basic Usage

        ### Usage with Authentication

        ## Import

        The `harbor_preheat_instance` resource can be imported using the preheat instance ID.

        ```sh
        $ pulumi import harbor:index/preheatInstance:PreheatInstance example /p2p/preheat/instances/example-preheat-instance
        ```

        :param str resource_name: The name of the resource.
        :param PreheatInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PreheatInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_mode: Optional[pulumi.Input[str]] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 vendor: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PreheatInstanceArgs.__new__(PreheatInstanceArgs)

            __props__.__dict__["auth_mode"] = auth_mode
            __props__.__dict__["default"] = default
            __props__.__dict__["description"] = description
            __props__.__dict__["enabled"] = enabled
            if endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint'")
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["insecure"] = insecure
            __props__.__dict__["name"] = name
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
            __props__.__dict__["username"] = username
            if vendor is None and not opts.urn:
                raise TypeError("Missing required property 'vendor'")
            __props__.__dict__["vendor"] = vendor
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password", "token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(PreheatInstance, __self__).__init__(
            'harbor:index/preheatInstance:PreheatInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_mode: Optional[pulumi.Input[str]] = None,
            default: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            endpoint: Optional[pulumi.Input[str]] = None,
            insecure: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None,
            vendor: Optional[pulumi.Input[str]] = None) -> 'PreheatInstance':
        """
        Get an existing PreheatInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_mode: The authentication mode for the preheat instance. Must be either "NONE", "BASIC", or "OAUTH". Defaults to "NONE".
        :param pulumi.Input[bool] default: Whether the preheat instance is the default instance. Defaults to false.
        :param pulumi.Input[str] description: The description of the preheat instance. Defaults to an empty string.
        :param pulumi.Input[bool] enabled: Whether the preheat instance is enabled. Defaults to true.
        :param pulumi.Input[str] endpoint: The endpoint of the preheat instance.
        :param pulumi.Input[bool] insecure: Whether to allow insecure connections to the preheat instance. Defaults to false.
        :param pulumi.Input[str] name: The name of the preheat instance.
        :param pulumi.Input[str] password: The password for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        :param pulumi.Input[str] token: The token for the preheat instance. Required if `auth_mode` is "OAUTH". Defaults to an empty string.
        :param pulumi.Input[str] username: The username for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        :param pulumi.Input[str] vendor: The vendor of the preheat instance. Must be either "dragonfly" or "kraken".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PreheatInstanceState.__new__(_PreheatInstanceState)

        __props__.__dict__["auth_mode"] = auth_mode
        __props__.__dict__["default"] = default
        __props__.__dict__["description"] = description
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["endpoint"] = endpoint
        __props__.__dict__["insecure"] = insecure
        __props__.__dict__["name"] = name
        __props__.__dict__["password"] = password
        __props__.__dict__["token"] = token
        __props__.__dict__["username"] = username
        __props__.__dict__["vendor"] = vendor
        return PreheatInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authMode")
    def auth_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The authentication mode for the preheat instance. Must be either "NONE", "BASIC", or "OAUTH". Defaults to "NONE".
        """
        return pulumi.get(self, "auth_mode")

    @property
    @pulumi.getter
    def default(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the preheat instance is the default instance. Defaults to false.
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the preheat instance. Defaults to an empty string.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the preheat instance is enabled. Defaults to true.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The endpoint of the preheat instance.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def insecure(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to allow insecure connections to the preheat instance. Defaults to false.
        """
        return pulumi.get(self, "insecure")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the preheat instance.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        """
        The password for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional[str]]:
        """
        The token for the preheat instance. Required if `auth_mode` is "OAUTH". Defaults to an empty string.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        """
        The username for the preheat instance. Required if `auth_mode` is "BASIC". Defaults to an empty string.
        """
        return pulumi.get(self, "username")

    @property
    @pulumi.getter
    def vendor(self) -> pulumi.Output[str]:
        """
        The vendor of the preheat instance. Must be either "dragonfly" or "kraken".
        """
        return pulumi.get(self, "vendor")
