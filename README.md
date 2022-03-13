# Redhat NVidia gpu init container

## Why
The GPU operator needs to be initialized before workloads can leverage GPU capabilities.
This small container will deploy CR's that will trigger the GPU operator reconciliation.

//TODO: Add more docs + code examples.

## Generate new gpu operator image bundle for addon
Prerequisites:
 - [Skipper](https://github.com/Stratoscale/skipper)
 - Podman / Docker
 - Modify `skipper.yaml` by adding a new volume mount
    ```(yaml)
    /local/path/to/managed-tenants-bundles: "/managed-tenants-bundles"
    ```

Execute:
```(shell)
skipper run "hack/gpu_operator_new_version.py -mP /managed-tenants-bundles -c <channel(stable/beta/...)> -v v<version> -pv <previuose-version>"

```

