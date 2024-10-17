# packer-rocm

[MaaS](https://maas.io/)-enabled [Packer](https://www.packer.io/) images
with `amdgpu-dkms` and _(optional)_ [ROCm](https://www.amd.com/en/products/software/rocm.html). Builds on the [canonical/packer-maas](https://github.com/canonical/packer-maas/)
project.


## Building

### Prerequisites

* [packer](https://developer.hashicorp.com/packer/docs/install)
* `ansible`: [pipx](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx) or [pip](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pip)
* `qemu`, `qemu-nbd`, `nbd-fuse`, `fuse`, `nbdkit`
* `nbdkit-{nbd-plugin,plugin-nbd}` _(Fedora/Debian, respectively)_
* `xorriso` or `genisoimage`
* `rsync`
* `git`

### Playbook

```shell
ansible-galaxy collection install ansible.posix community.general
ansible-pull -U https://github.com/nod-ai/ADA.git packer-rocm/playbooks/build.yml \
    -e rocm_releases=6.2.2,6.2.1 \
    -e rocm_kernel=linux-image-generic-hwe-22.04 \
    -e rocm_extras=linux-headers-generic-hwe-22.04,mesa-amdgpu-va-drivers \
    -e rocm_builder_cpus=16 \
    -e rocm_builder_disk=70G
```

All of these variables are _optional_. Please see [I/O](#io) for more.

### Manual

1. Clone repositories:

    ```shell
    git clone https://github.com/canonical/packer-maas.git
    git clone https://github.com/nod-ai/ADA.git
    ```

    Place any `.deb` packages to include with the image in `ADA/packer-rocm/ubuntu/packages/`

2. Copy assets from _ADA_ `packer-rocm` to the _Canonical_ `packer-maas` source:

    ```shell
    # Repeat '--exclude' with shell expansion, slashes are significant for 'rsync'
    rsync -avP --exclude={'*.md','LICENSE','NOTICE'} ADA/packer-rocm/ packer-maas/
    ```

3. Install plugins:

    ```shell
    cd packer-maas/ubuntu
    packer init .
    ```

4. Build

    ```shell
    # Change working directory to the prepared sources
    cd packer-maas/ubuntu

    # Build
    PACKER_LOG=1 packer build \
        -var rocm_releases="6.2.2,6.2.1" \
        -var rocm_kernel="linux-image-generic-hwe-22.04" \
        -var rocm_extras="linux-headers-generic-hwe-22.04,mesa-amdgpu-va-drivers" \
        -var rocm_builder_disk="70G" \
        -only=qemu.rocm .
    ```

### I/O

| Variable | Description | Default |
|:----------:|-------------|:---------:|
| `rocm_releases` | One or more versions to include _[comma-separated string]_.<br/>Newest selects the `amdgpu` driver. | `6.2.2` |
| `rocm_kernel` | The kernel package with an optional release specifier. | `linux-image-generic-hwe-22.04` |
| `rocm_extras` | Packages to install _before_ `amdgpu-dkms` and _ROCm_.<br/>May also specify releases with `=x.y.z` or globbed. | `linux-headers-generic-hwe-22.04,mesa-amdgpu-va-drivers` |
| `rocm_filename` | The name of the output file/artifact _(tarball)_ | `ubuntu-rocm.tar.gz` |
| `rocm_installed` | If _ROCm_ packages are installed. The `amdgpu` _driver/extras_ are, always. | `False` |
| `rocm_builder_cpus` | Number of virtual CPUs given to the builder VM. | _4_ |
| `rocm_builder_disk` | Space given to the builder; releases compound quickly. | _70G_ |
| `rocm_builder_memory` | Megabytes of memory given to the builder.<br/>Reduction may cause out-of-memory conditions when compiling. | _4096_ |
| `niccli_wanted` | If [niccli](https://techdocs.broadcom.com/us/en/storage-and-ethernet-connectivity/ethernet-nic-controllers/bcm957xxx/adapters/Configuration-adapter/nic-cli-configuration-utility.html) is included in the image. | `True` |
| `niccli_url` | The URL for the _Broadcom_ `niccli` installation archive. | [Link](https://docs.broadcom.com/docs-and-downloads/ethernet-network-adapters/NXE/Thor2/GCA1/bcm5760x_230.2.52.0a.zip) |
| `niccli_sum` | Optional, string. Checksum to use for validating `niccli_url`.<br/>Example: `sha256:abcd1234` | _Undefined_ |
| `niccli_driver` | If the `bnxt_{en,re}` NIC drivers are included. | `True` |
| `headless` | If the VNC window for the VM is _hidden_ during build. | `True` |

#### MaaS

This image is built _(and uploaded)_ in `tgz` format to respect customized disk layouts:

```shell
maas admin boot-resources create \
     name='custom/packer-rocm-ubuntu-22.04.5' \
     title='packer-rocm (Ubuntu 22.04.5)' \
     architecture='amd64/generic' \
     filetype='tgz' \
     content@=ubuntu-rocm.tar.gz
```

The artifact is named `ubuntu-rocm.tar.gz`. When building with `ansible-pull`, it may be here:  
`~/.ansible/pull/$HOSTNAME/packer-rocm/packer-maas/ubuntu`

#### Proxy

These _environment variables_ are respected:

* `http_proxy`
* `https_proxy`
* `no_proxy`
