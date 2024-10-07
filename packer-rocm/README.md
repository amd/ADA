# packer-rocm

[MaaS](https://maas.io/)-enabled [Packer](https://www.packer.io/) images
with `amdgpu-dkms` and optional [ROCm](https://www.amd.com/en/products/software/rocm.html)
installed. Builds on the [canonical/packer-maas](https://github.com/canonical/packer-maas/)
project.


## Building

### Requirements

* [packer](https://developer.hashicorp.com/packer/docs/install)
* `ansible`: [pipx](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pipx) or [pip](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html#installing-and-upgrading-ansible-with-pip)
* `qemu`
* `rsync`
* `git`

### Playbook

```shell
ansible-galaxy collection install ansible.posix community.general
ansible-pull -U https://github.com/nod-ai/ADA.git packer-rocm/playbooks/build.yml
```

Variables noted in [I/O](#io) may be given like so: `ansible-pull ... -e 'var=value'`

### Manual

1. Clone repositories:

    ```shell
    git clone https://github.com/canonical/packer-maas.git
    git clone https://github.com/nod-ai/ADA.git
    ```

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
        -var kernel=linux-generic \
        -var rocm_releases="6.2.2,6.2.1" \
        -var rocm_extras="mesa-amdgpu-va-drivers,ansible" \
        -var rocm_builder_disk="70G" \
        -only=qemu.rocm .
    ```

### I/O

The artifact is named `ubuntu-rocm.dd.gz`. When building with `ansible-pull`, it may be here:  
`~/.ansible/pull/$HOSTNAME/packer-rocm/packer-maas/ubuntu`

| Variable | Description | Default |
|:----------:|-------------|:---------:|
| `rocm_releases` | One or more _ROCm_ releases to include in the image. Newest selects the `amdgpu` driver. | _6.2.2_ |
| `rocm_extras` | Packages to install _after_ `amdgpu-dkms` and _ROCm_. Accepts comma-separated strings. | _mesa-amdgpu-va-drivers_ |
| `rocm_installed` | If _ROCm_ multi-release packages should be installed | `False` |
| `rocm_builder_disk` | Space given to the builder VM; `amdgpu` and _ROCm_ releases compound quickly. | _70G_ |
| `headless` | If the VNC window for the builder VM is _hidden_ | `True` |
| `kernel` | _MaaS_ images normally do not include a kernel, set this to include one. | _Ansible:_ `linux-generic`<br />_Manual:_ None |

#### Proxy

If the build requires a proxy for downloading the ISO, updates, or ROCm... these _environment variables_ are respected:

* `http_proxy`
* `https_proxy`
* `no_proxy`
