# sos

This role installs [sosreport](https://github.com/sosreport/sos),
[xsos](https://github.com/ryran/xsos),
and `extras.d` entries. Report _generation/collection_ is left as an admin activity.

## Variables

1. `sos_extras`: custom commands or files in `sos` reports.
Default: see the [example playbook](#example)
2. `sos_xsos`: controls [xsos](https://github.com/ryran/xsos) installation.
Default: `true`
3. `sos_xsos_url`: `xsos` installation URL.
[Default](https://github.com/ryran/xsos/raw/master/xsos)

## Example

```yaml
---
- name: "'sos' role"
  hosts: all
  roles:
    - name: Configure 'sos', extras, and 'xsos'
      role: sos
      vars:
        sos_xsos_url: 'https://raw.githubusercontent.com/ryran/xsos/v0.7.33/xsos'
        sos_extras:
          amdgpu:
            - 'rocm-smi -a'
          yours:
            - ':/some/file/to/read'
```
