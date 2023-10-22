# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetGroupsResult',
    'AwaitableGetGroupsResult',
    'get_groups',
    'get_groups_output',
]

@pulumi.output_type
class GetGroupsResult:
    """
    A collection of values returned by getGroups.
    """
    def __init__(__self__, group_name=None, groups=None, id=None, ldap_group_dn=None):
        if group_name and not isinstance(group_name, str):
            raise TypeError("Expected argument 'group_name' to be a str")
        pulumi.set(__self__, "group_name", group_name)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ldap_group_dn and not isinstance(ldap_group_dn, str):
            raise TypeError("Expected argument 'ldap_group_dn' to be a str")
        pulumi.set(__self__, "ldap_group_dn", ldap_group_dn)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[str]:
        """
        The name of the group.
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetGroupsGroupResult']:
        """
        (Computed) A list of groups matching the previous arguments. Each `group` object provides the attributes documented below.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ldapGroupDn")
    def ldap_group_dn(self) -> Optional[str]:
        """
        The LDAP group DN of the group.
        """
        return pulumi.get(self, "ldap_group_dn")


class AwaitableGetGroupsResult(GetGroupsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupsResult(
            group_name=self.group_name,
            groups=self.groups,
            id=self.id,
            ldap_group_dn=self.ldap_group_dn)


def get_groups(group_name: Optional[str] = None,
               ldap_group_dn: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupsResult:
    """
    ## Example Usage


    :param str group_name: The name of the group to filter by.
    :param str ldap_group_dn: The LDAP group DN to filter by.
    """
    __args__ = dict()
    __args__['groupName'] = group_name
    __args__['ldapGroupDn'] = ldap_group_dn
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('harbor:index/getGroups:getGroups', __args__, opts=opts, typ=GetGroupsResult).value

    return AwaitableGetGroupsResult(
        group_name=pulumi.get(__ret__, 'group_name'),
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        ldap_group_dn=pulumi.get(__ret__, 'ldap_group_dn'))


@_utilities.lift_output_func(get_groups)
def get_groups_output(group_name: Optional[pulumi.Input[Optional[str]]] = None,
                      ldap_group_dn: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupsResult]:
    """
    ## Example Usage


    :param str group_name: The name of the group to filter by.
    :param str ldap_group_dn: The LDAP group DN to filter by.
    """
    ...
