# vra-sdk-go

## Overview

This Go language SDK is for the VMware vRealize Automation products.

## Building and updating the SDK

This SDK is auto-generated from the swagger files produced from the
vRealize Automation product. As such, it is suggested to file an issue
for the maintainers to update the SDK rather than a PR.

To update, run the following commands:
```
# Download new base swagger files
make update

# Process the swagger files to generate the SDK
make swaggeer
```

Note: the process of creating the SDK will validate the swagger for
correctness as well as run some fixup scripts to normalizes the name
mangling and adds some missing filter parameters.

Tools used for the SDK generation:
- [go-swagger](https://github.com/go-swagger/go-swagger)

## Contributing

The vra-sdk-go project team welcomes contributions from the community.
If you wish to contribute code and you have not signed our contributor
license agreement (CLA), our bot will update the issue when you open a
Pull Request. For any questions about the CLA process, please refer to
our [FAQ](https://cla.vmware.com/faq). For more detailed information,
refer to [CONTRIBUTING.md](CONTRIBUTING.md).

## License

vra-sdk-go is available under the [Apache 2 license](LICENSE).
