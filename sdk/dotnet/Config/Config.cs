// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumiverse.Harbor
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("harbor");

        private static readonly __Value<int?> _apiVersion = new __Value<int?>(() => __config.GetInt32("apiVersion") ?? 2);
        public static int? ApiVersion
        {
            get => _apiVersion.Get();
            set => _apiVersion.Set(value);
        }

        private static readonly __Value<bool?> _insecure = new __Value<bool?>(() => __config.GetBoolean("insecure") ?? Utilities.GetEnvBoolean("HARBOR_IGNORE_CERT") ?? true);
        public static bool? Insecure
        {
            get => _insecure.Get();
            set => _insecure.Set(value);
        }

        private static readonly __Value<string?> _password = new __Value<string?>(() => __config.Get("password") ?? Utilities.GetEnv("HARBOR_PASSWORD"));
        public static string? Password
        {
            get => _password.Get();
            set => _password.Set(value);
        }

        private static readonly __Value<string?> _url = new __Value<string?>(() => __config.Get("url") ?? Utilities.GetEnv("HARBOR_URL"));
        public static string? Url
        {
            get => _url.Get();
            set => _url.Set(value);
        }

        private static readonly __Value<string?> _username = new __Value<string?>(() => __config.Get("username") ?? Utilities.GetEnv("HARBOR_USERNAME"));
        public static string? Username
        {
            get => _username.Get();
            set => _username.Set(value);
        }

    }
}
