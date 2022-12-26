# Pulumi Harbor Provider Example

Attention: Don't forget to set following pulumi config values, before you attempt to run the example:

```
pulumi config set harbor:username <your harbor username>
pulumi config set harbor:password <your harbor password> --secret
pulumi config set harbor:url <your harbor url>
```

or via config yaml:

````yaml
config:
  harbor:password: Harbor12345
  harbor:url: http://core.harbor.domain
  harbor:username: admin
````
