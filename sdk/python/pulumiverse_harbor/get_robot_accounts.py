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
from . import outputs

__all__ = [
    'GetRobotAccountsResult',
    'AwaitableGetRobotAccountsResult',
    'get_robot_accounts',
    'get_robot_accounts_output',
]

@pulumi.output_type
class GetRobotAccountsResult:
    """
    A collection of values returned by getRobotAccounts.
    """
    def __init__(__self__, id=None, level=None, name=None, project_id=None, robot_accounts=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if level and not isinstance(level, str):
            raise TypeError("Expected argument 'level' to be a str")
        pulumi.set(__self__, "level", level)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project_id and not isinstance(project_id, int):
            raise TypeError("Expected argument 'project_id' to be a int")
        pulumi.set(__self__, "project_id", project_id)
        if robot_accounts and not isinstance(robot_accounts, list):
            raise TypeError("Expected argument 'robot_accounts' to be a list")
        pulumi.set(__self__, "robot_accounts", robot_accounts)

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def level(self) -> Optional[builtins.str]:
        """
        Level of the robot account, currently either `system` or `project`. Default is `system`.
        """
        return pulumi.get(self, "level")

    @property
    @pulumi.getter
    def name(self) -> Optional[builtins.str]:
        """
        The name of the robot account to filter by.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[builtins.int]:
        """
        The id of the project within harbor.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="robotAccounts")
    def robot_accounts(self) -> Sequence['outputs.GetRobotAccountsRobotAccountResult']:
        return pulumi.get(self, "robot_accounts")


class AwaitableGetRobotAccountsResult(GetRobotAccountsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRobotAccountsResult(
            id=self.id,
            level=self.level,
            name=self.name,
            project_id=self.project_id,
            robot_accounts=self.robot_accounts)


def get_robot_accounts(level: Optional[builtins.str] = None,
                       name: Optional[builtins.str] = None,
                       project_id: Optional[builtins.int] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRobotAccountsResult:
    """
    ## Example Usage


    :param builtins.str level: Level of the robot account, currently either `system` or `project`. Default is `system`.
    :param builtins.str name: The name of the robot account to filter by.
    :param builtins.int project_id: The id of the project within harbor.
    """
    __args__ = dict()
    __args__['level'] = level
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('harbor:index/getRobotAccounts:getRobotAccounts', __args__, opts=opts, typ=GetRobotAccountsResult).value

    return AwaitableGetRobotAccountsResult(
        id=pulumi.get(__ret__, 'id'),
        level=pulumi.get(__ret__, 'level'),
        name=pulumi.get(__ret__, 'name'),
        project_id=pulumi.get(__ret__, 'project_id'),
        robot_accounts=pulumi.get(__ret__, 'robot_accounts'))
def get_robot_accounts_output(level: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                              project_id: Optional[pulumi.Input[Optional[builtins.int]]] = None,
                              opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRobotAccountsResult]:
    """
    ## Example Usage


    :param builtins.str level: Level of the robot account, currently either `system` or `project`. Default is `system`.
    :param builtins.str name: The name of the robot account to filter by.
    :param builtins.int project_id: The id of the project within harbor.
    """
    __args__ = dict()
    __args__['level'] = level
    __args__['name'] = name
    __args__['projectId'] = project_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('harbor:index/getRobotAccounts:getRobotAccounts', __args__, opts=opts, typ=GetRobotAccountsResult)
    return __ret__.apply(lambda __response__: GetRobotAccountsResult(
        id=pulumi.get(__response__, 'id'),
        level=pulumi.get(__response__, 'level'),
        name=pulumi.get(__response__, 'name'),
        project_id=pulumi.get(__response__, 'project_id'),
        robot_accounts=pulumi.get(__response__, 'robot_accounts')))
