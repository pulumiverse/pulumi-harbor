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
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    """
    A collection of values returned by getProject.
    """
    def __init__(__self__, id=None, name=None, project_id=None, public=None, type=None, vulnerability_scanning=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if public and not isinstance(public, bool):
            raise TypeError("Expected argument 'public' to be a bool")
        pulumi.set(__self__, "public", public)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vulnerability_scanning and not isinstance(vulnerability_scanning, bool):
            raise TypeError("Expected argument 'vulnerability_scanning' to be a bool")
        pulumi.set(__self__, "vulnerability_scanning", vulnerability_scanning)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the project.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> builtins.int:
        """
        The id of the project within harbor.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def public(self) -> builtins.bool:
        """
        If the project has public accessibility.
        """
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        The type of the project : Project or ProxyCache.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vulnerabilityScanning")
    def vulnerability_scanning(self) -> builtins.bool:
        """
        If the images is scanned for vulnerabilities when push to harbor.
        """
        return pulumi.get(self, "vulnerability_scanning")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            id=self.id,
            name=self.name,
            project_id=self.project_id,
            public=self.public,
            type=self.type,
            vulnerability_scanning=self.vulnerability_scanning)


def get_project(name: Optional[builtins.str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    ## Example Usage


    :param builtins.str name: The name of the project.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('harbor:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        public=pulumi.get(__ret__, 'public'),
        type=pulumi.get(__ret__, 'type'),
        vulnerability_scanning=pulumi.get(__ret__, 'vulnerability_scanning'))
def get_project_output(name: Optional[pulumi.Input[builtins.str]] = None,
                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetProjectResult]:
    """
    ## Example Usage


    :param builtins.str name: The name of the project.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('harbor:index/getProject:getProject', __args__, opts=opts, typ=GetProjectResult)
    return __ret__.apply(lambda __response__: GetProjectResult(
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        public=pulumi.get(__response__, 'public'),
        type=pulumi.get(__response__, 'type'),
        vulnerability_scanning=pulumi.get(__response__, 'vulnerability_scanning')))
