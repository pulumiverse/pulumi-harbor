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

__all__ = [
    'GetProjectsResult',
    'AwaitableGetProjectsResult',
    'get_projects',
    'get_projects_output',
]

@pulumi.output_type
class GetProjectsResult:
    """
    A collection of values returned by getProjects.
    """
    def __init__(__self__, id=None, name=None, projects=None, public=None, type=None, vulnerability_scanning=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if projects and not isinstance(projects, list):
            raise TypeError("Expected argument 'projects' to be a list")
        pulumi.set(__self__, "projects", projects)
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
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def projects(self) -> Sequence['outputs.GetProjectsProjectResult']:
        return pulumi.get(self, "projects")

    @property
    @pulumi.getter
    def public(self) -> Optional[bool]:
        return pulumi.get(self, "public")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vulnerabilityScanning")
    def vulnerability_scanning(self) -> Optional[bool]:
        return pulumi.get(self, "vulnerability_scanning")


class AwaitableGetProjectsResult(GetProjectsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectsResult(
            id=self.id,
            name=self.name,
            projects=self.projects,
            public=self.public,
            type=self.type,
            vulnerability_scanning=self.vulnerability_scanning)


def get_projects(name: Optional[str] = None,
                 public: Optional[bool] = None,
                 type: Optional[str] = None,
                 vulnerability_scanning: Optional[bool] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectsResult:
    """
    ## Example Usage
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['public'] = public
    __args__['type'] = type
    __args__['vulnerabilityScanning'] = vulnerability_scanning
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('harbor:index/getProjects:getProjects', __args__, opts=opts, typ=GetProjectsResult).value

    return AwaitableGetProjectsResult(
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        projects=pulumi.get(__ret__, 'projects'),
        public=pulumi.get(__ret__, 'public'),
        type=pulumi.get(__ret__, 'type'),
        vulnerability_scanning=pulumi.get(__ret__, 'vulnerability_scanning'))


@_utilities.lift_output_func(get_projects)
def get_projects_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                        public: Optional[pulumi.Input[Optional[bool]]] = None,
                        type: Optional[pulumi.Input[Optional[str]]] = None,
                        vulnerability_scanning: Optional[pulumi.Input[Optional[bool]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectsResult]:
    """
    ## Example Usage
    """
    ...
