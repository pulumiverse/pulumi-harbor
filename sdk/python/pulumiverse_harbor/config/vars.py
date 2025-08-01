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
from .. import _utilities

import types

__config__ = pulumi.Config('harbor')


class _ExportableConfig(types.ModuleType):
    @property
    def api_version(self) -> int:
        return __config__.get_int('apiVersion') or 2

    @property
    def bearer_token(self) -> Optional[str]:
        return __config__.get('bearerToken')

    @property
    def insecure(self) -> bool:
        return __config__.get_bool('insecure') or (_utilities.get_env_bool('HARBOR_IGNORE_CERT') or True)

    @property
    def password(self) -> Optional[str]:
        return __config__.get('password') or _utilities.get_env('HARBOR_PASSWORD')

    @property
    def robot_prefix(self) -> Optional[str]:
        return __config__.get('robotPrefix')

    @property
    def session_id(self) -> Optional[str]:
        return __config__.get('sessionId')

    @property
    def url(self) -> Optional[str]:
        return __config__.get('url') or _utilities.get_env('HARBOR_URL')

    @property
    def username(self) -> Optional[str]:
        return __config__.get('username') or _utilities.get_env('HARBOR_USERNAME')

