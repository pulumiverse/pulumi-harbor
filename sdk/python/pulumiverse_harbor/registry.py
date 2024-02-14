# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RegistryArgs', 'Registry']

@pulumi.input_type
class RegistryArgs:
    def __init__(__self__, *,
                 endpoint_url: pulumi.Input[str],
                 provider_name: pulumi.Input[str],
                 access_id: Optional[pulumi.Input[str]] = None,
                 access_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Registry resource.
        """
        pulumi.set(__self__, "endpoint_url", endpoint_url)
        pulumi.set(__self__, "provider_name", provider_name)
        if access_id is not None:
            pulumi.set(__self__, "access_id", access_id)
        if access_secret is not None:
            pulumi.set(__self__, "access_secret", access_secret)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="endpointUrl")
    def endpoint_url(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endpoint_url")

    @endpoint_url.setter
    def endpoint_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_url", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter(name="accessId")
    def access_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_id")

    @access_id.setter
    def access_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_id", value)

    @property
    @pulumi.getter(name="accessSecret")
    def access_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_secret")

    @access_secret.setter
    def access_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_secret", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RegistryState:
    def __init__(__self__, *,
                 access_id: Optional[pulumi.Input[str]] = None,
                 access_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint_url: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 registry_id: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Registry resources.
        """
        if access_id is not None:
            pulumi.set(__self__, "access_id", access_id)
        if access_secret is not None:
            pulumi.set(__self__, "access_secret", access_secret)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if endpoint_url is not None:
            pulumi.set(__self__, "endpoint_url", endpoint_url)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if provider_name is not None:
            pulumi.set(__self__, "provider_name", provider_name)
        if registry_id is not None:
            pulumi.set(__self__, "registry_id", registry_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accessId")
    def access_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_id")

    @access_id.setter
    def access_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_id", value)

    @property
    @pulumi.getter(name="accessSecret")
    def access_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "access_secret")

    @access_secret.setter
    def access_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_secret", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="endpointUrl")
    def endpoint_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "endpoint_url")

    @endpoint_url.setter
    def endpoint_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_url", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "provider_name")

    @provider_name.setter
    def provider_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_name", value)

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "registry_id")

    @registry_id.setter
    def registry_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "registry_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class Registry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_id: Optional[pulumi.Input[str]] = None,
                 access_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint_url: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        Harbor project can be imported using the `registry id` eg,<break><break> ` <break><break> ```sh<break> $ pulumi import harbor:index/registry:Registry main /registries/7 <break>```<break><break>  `<break><break>

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegistryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        Harbor project can be imported using the `registry id` eg,<break><break> ` <break><break> ```sh<break> $ pulumi import harbor:index/registry:Registry main /registries/7 <break>```<break><break>  `<break><break>

        :param str resource_name: The name of the resource.
        :param RegistryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegistryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_id: Optional[pulumi.Input[str]] = None,
                 access_secret: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 endpoint_url: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegistryArgs.__new__(RegistryArgs)

            __props__.__dict__["access_id"] = access_id
            __props__.__dict__["access_secret"] = None if access_secret is None else pulumi.Output.secret(access_secret)
            __props__.__dict__["description"] = description
            if endpoint_url is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_url'")
            __props__.__dict__["endpoint_url"] = endpoint_url
            __props__.__dict__["insecure"] = insecure
            __props__.__dict__["name"] = name
            if provider_name is None and not opts.urn:
                raise TypeError("Missing required property 'provider_name'")
            __props__.__dict__["provider_name"] = provider_name
            __props__.__dict__["registry_id"] = None
            __props__.__dict__["status"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessSecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Registry, __self__).__init__(
            'harbor:index/registry:Registry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_id: Optional[pulumi.Input[str]] = None,
            access_secret: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            endpoint_url: Optional[pulumi.Input[str]] = None,
            insecure: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            provider_name: Optional[pulumi.Input[str]] = None,
            registry_id: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'Registry':
        """
        Get an existing Registry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RegistryState.__new__(_RegistryState)

        __props__.__dict__["access_id"] = access_id
        __props__.__dict__["access_secret"] = access_secret
        __props__.__dict__["description"] = description
        __props__.__dict__["endpoint_url"] = endpoint_url
        __props__.__dict__["insecure"] = insecure
        __props__.__dict__["name"] = name
        __props__.__dict__["provider_name"] = provider_name
        __props__.__dict__["registry_id"] = registry_id
        __props__.__dict__["status"] = status
        return Registry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessId")
    def access_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "access_id")

    @property
    @pulumi.getter(name="accessSecret")
    def access_secret(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "access_secret")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endpointUrl")
    def endpoint_url(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint_url")

    @property
    @pulumi.getter
    def insecure(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "insecure")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="providerName")
    def provider_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "provider_name")

    @property
    @pulumi.getter(name="registryId")
    def registry_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "registry_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

