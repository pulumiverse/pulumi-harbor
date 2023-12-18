# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ConfigSecurityArgs', 'ConfigSecurity']

@pulumi.input_type
class ConfigSecurityArgs:
    def __init__(__self__, *,
                 cve_allowlists: pulumi.Input[Sequence[pulumi.Input[str]]],
                 expires_at: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ConfigSecurity resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cve_allowlists: System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        :param pulumi.Input[int] expires_at: The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
        """
        pulumi.set(__self__, "cve_allowlists", cve_allowlists)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)

    @property
    @pulumi.getter(name="cveAllowlists")
    def cve_allowlists(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        """
        return pulumi.get(self, "cve_allowlists")

    @cve_allowlists.setter
    def cve_allowlists(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "cve_allowlists", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[int]]:
        """
        The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expires_at", value)


@pulumi.input_type
class _ConfigSecurityState:
    def __init__(__self__, *,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 cve_allowlists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 expires_at: Optional[pulumi.Input[int]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConfigSecurity resources.
        :param pulumi.Input[str] creation_time: Time of creation of the list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cve_allowlists: System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        :param pulumi.Input[int] expires_at: The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
        :param pulumi.Input[str] update_time: Time of update of the list.
        """
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if cve_allowlists is not None:
            pulumi.set(__self__, "cve_allowlists", cve_allowlists)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of creation of the list.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter(name="cveAllowlists")
    def cve_allowlists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        """
        return pulumi.get(self, "cve_allowlists")

    @cve_allowlists.setter
    def cve_allowlists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cve_allowlists", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[int]]:
        """
        The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        Time of update of the list.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class ConfigSecurity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cve_allowlists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 expires_at: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        The list can be imported using the `id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/configSecurity:ConfigSecurity main "7" <break>```<break><break>` > Note that at this point of time Harbor doesn't has any api endpoint for deleting this list. Only updating the records.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cve_allowlists: System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        :param pulumi.Input[int] expires_at: The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConfigSecurityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        The list can be imported using the `id` eg, `<break><break>```sh<break> $ pulumi import harbor:index/configSecurity:ConfigSecurity main "7" <break>```<break><break>` > Note that at this point of time Harbor doesn't has any api endpoint for deleting this list. Only updating the records.

        :param str resource_name: The name of the resource.
        :param ConfigSecurityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfigSecurityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cve_allowlists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 expires_at: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConfigSecurityArgs.__new__(ConfigSecurityArgs)

            if cve_allowlists is None and not opts.urn:
                raise TypeError("Missing required property 'cve_allowlists'")
            __props__.__dict__["cve_allowlists"] = cve_allowlists
            __props__.__dict__["expires_at"] = expires_at
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["update_time"] = None
        super(ConfigSecurity, __self__).__init__(
            'harbor:index/configSecurity:ConfigSecurity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            cve_allowlists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            expires_at: Optional[pulumi.Input[int]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'ConfigSecurity':
        """
        Get an existing ConfigSecurity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_time: Time of creation of the list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cve_allowlists: System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        :param pulumi.Input[int] expires_at: The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
        :param pulumi.Input[str] update_time: Time of update of the list.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConfigSecurityState.__new__(_ConfigSecurityState)

        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["cve_allowlists"] = cve_allowlists
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["update_time"] = update_time
        return ConfigSecurity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        Time of creation of the list.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="cveAllowlists")
    def cve_allowlists(self) -> pulumi.Output[Sequence[str]]:
        """
        System allowlist. Vulnerabilities in this list will be ignored when pushing and pulling images. Should be in the format or `["CVE-123", "CVE-145"]` or `["CVE-123"]`
        """
        return pulumi.get(self, "cve_allowlists")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[Optional[int]]:
        """
        The time for expiration of the allowlist, in the form of seconds since epoch. This is an optional attribute, if it's not set the CVE allowlist does not expire.
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Time of update of the list.
        """
        return pulumi.get(self, "update_time")

