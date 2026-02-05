# ADA
AMD Deployment Assistant


## Build a deployable image with [Packer](https://github.com/nod-ai/ADA/tree/main/packer-rocm) 

## Deploy the image with MAAS

### If you use NFS please update the /opt/rocm symlink to the NFS mount
```
ln -s /nfsshare/rocm-<version> /opt/rocm
```

## Deploy SLURM

## All Logs Collection Script smci_instinct_platform_amdgpu_alllogs_collection.py

Script generates and pulls the AllLogs from the target system.  As part of
operations the power supplies are checked.  If the UBB assembly cannot be
reached then, the script will attempt to power cycle the system.

### Execution on Windows

1.	Ensure Python is installed on the system by running python3.exe
    - If it is not installed, then please follow on screen directions to install it.
```
python3.exe
```
2.	Ensure the requests module is installed
```
python3.exe -m pip install --upgrade pip
python3.exe -m pip install requests
```
3.	Run the script to collect the logs
    - User will be prompted to provide the BMC username, password, and IP address
```
python3.exe smci_instinct_platform_amdgpu_alllogs_collection.py
```

### Example of running on Ubuntu

1.	Ensure Python with packages is installed on the system
```
sudo apt update && sudo apt install -y python3-full
```
3.	Run the script to collect the logs
    - User will be prompted to provide the BMC username, password, and IP address
```
python3 smci_instinct_platform_amdgpu_alllogs_collection.py
```

### Example of running on RHEL


1.	Ensure Python with packages is installed on the system
```
sudo yum update && sudo yum install -y python3-requests
```
3.	Run the script to collect the logs
    - User will be prompted to provide the BMC username, password, and IP address
```
python3 smci_instinct_platform_amdgpu_alllogs_collection.py
```

### Parameters

| Parameter      | Description                                                                          |  
|----------------|--------------------------------------------------------------------------------------|  
| --debug        | Adding the --debug option will provide debug information and show the BMC password.  |  
| --no-power-off | Do not power off the system even if power cycle required                             |  

### Batch use

To support batch use, user input can be replaced with environment variables set in the environment.

| Parameter    | Description                         |  
|--------------|-------------------------------------|  
| BMC_IP       | IPv4 address of the BMC.            |  
| BMC_USERNAME | Username for logging into the BMC.  |  
| BMC_PASSWORD | Password used to log into the BMC.  |  

## OAM Health Check Script smci_instinct_platform_amdgpu_health_check.py

Script checks the health status metrics from the OAMs and the HBM memory for
the OAMs.  If one or more OAMs are shown to have an issue then the AllLogs
will automatically be captured. As part of operations the power supplies are
checked.  If the UBB assembly cannot be reached, then the script will attempt
to power cycle the system.

### Execution on Windows

1.	Ensure Python is installed on the system by running python3.exe
    - If it is not installed, then please follow on screen directions to install it.
```
python3.exe
```
2.	Ensure the requests module is installed
```
python3.exe -m pip install --upgrade pip
python3.exe -m pip install requests
```
3.	Run the script to collect the logs
    - User will be prompted to provide the BMC username, password, and IP address
```
python3.exe smci_instinct_platform_amdgpu_health_check.py
```

### Example of running on Ubuntu

1.	Ensure Python with packages is installed on the system
```
sudo apt update && sudo apt install -y python3-full
```
3.	Run the script to collect the logs
    - User will be prompted to provide the BMC username, password, and IP address
```
python3 smci_instinct_platform_amdgpu_health_check.py
```

### Example of running on RHEL


1.	Ensure Python with packages is installed on the system
```
sudo yum update && sudo yum install -y python3-requests
```
3.	Run the script to collect the logs
    - User will be prompted to provide the BMC username, password, and IP address
```
python3 smci_instinct_platform_amdgpu_health_check.py
```

### Parameters

| Parameter      | Description                                                                          |  
|----------------|--------------------------------------------------------------------------------------|  
| --debug        | Adding the --debug option will provide debug information and show the BMC password.  |  
| --no-power-off | Do not power off the system even if power cycle required                             |  

### Batch use

To support batch use, user input can be replaced with environment variables set in the environment.

| Parameter    | Description                         |  
|--------------|-------------------------------------|  
| BMC_IP       | IPv4 address of the BMC.            |  
| BMC_USERNAME | Username for logging into the BMC.  |  
| BMC_PASSWORD | Password used to log into the BMC.  |  

## Clear Logs Script smci_instinct_platform_amdgpu_clear_all_logs.py

###  Overview

The script deletes logs from the Instict DCGPU assembly followed by a power cycle.  This does not
delete the BMC logs.  Running this script is strongly suggested after updating the BKC (Instict
firmware) or replace an OAM module on the assembly.  Script can run from Windows or Linux.

### Execution on Windows

1.	Ensure Python is installed on the system by running python3.exe
    - If it is not installed, then please follow on screen directions to install it.
    - If it is installed, then please type quit to exit python3.exe.
```
python3.exe
```
2.	Ensure the requests module is installed
```
python3.exe -m pip install --upgrade pip
python3.exe -m pip install requests
```
3.	Run the script to clear the logs
    - User will be prompted to provide the BMC username, password, and IP address.
```
python3.exe smci_instinct_platform_amdgpu_clear_all_logs.py
```

### Example of running on Ubuntu

1.	Ensure Python with packages is installed on the system
```
sudo apt update && sudo apt install -y python3-full
```
2.	Run the script to clear the logs
    - User will be prompted to provide the BMC username, password, and IP address.
```
python3 smci_instinct_platform_amdgpu_clear_all_logs.py
```

### Example of running on RHEL

1.	Ensure Python with packages is installed on the system
```
sudo yum update && sudo yum install -y python3-requests
```
2.	Run the script to clear the logs
    - User will be prompted to provide the BMC username, password, and IP address.
```
python3 smci_instinct_platform_amdgpu_clear_all_logs.py
```

### Parameters

| Parameter      | Description                                                                          |  
|----------------|--------------------------------------------------------------------------------------|  
| --debug        | Adding the --debug option will provide debug information and show the BMC password.  |  
| --no-power-off | Do not power off the system even if power cycle required                             |  

### Batch use

To support batch use, user input can be replaced with environment variables set in the environment.

| Parameter    | Description                         |  
|--------------|-------------------------------------|  
| BMC_IP       | IPv4 address of the BMC.            |  
| BMC_USERNAME | Username for logging into the BMC.  |  
| BMC_PASSWORD | Password used to log into the BMC.  |  

## BKC Update Script smci_instinct_platform_amdgpu_update_bkc.py

Script updates the BKC on the DCGPU on the target system.  As part of
operations the power supplies are checked.  If the UBB assembly cannot be
reached, then the script will attempt to power cycle the system.

### Execution on Windows

1.	Ensure Python is installed on the system by running python3.exe
    - If it is not installed, then please follow on screen directions to install it.
```
python3.exe
```
2.	Ensure the requests module is installed
```
python3.exe -m pip install --upgrade pip
python3.exe -m pip install requests
```
3.	Run the script to update the BKC on the assebly
    - User will be prompted to provide the BMC username, password, IP address, BKC file name,
      and target BKC version.
```
python3.exe smci_instinct_platform_amdgpu_update_bkc.py
```

### Example of running on Ubuntu

1.	Ensure Python with packages is installed on the system
```
sudo apt update && sudo apt install -y python3-full
```
2.	Run the script to update the BKC on the assebly
    - User will be prompted to provide the BMC username, password, IP address, BKC file name,
      and target BKC version.
```
python3 smci_instinct_platform_amdgpu_update_bkc.py
```

### Example of running on RHEL

1.	Ensure Python with packages is installed on the system
```
sudo yum update && sudo yum install -y python3-requests
```
2.	Run the script to update the BKC on the assebly
    - User will be prompted to provide the BMC username, password, IP address, BKC file name,
      and target BKC version.
```
python3 smci_instinct_platform_amdgpu_update_bkc.py
```
### Parameters

| Parameter      | Description                                                                          |  
|----------------|--------------------------------------------------------------------------------------|  
| --debug        | Adding the --debug option will provide debug information and show the BMC password.  |  
| --no-power-off | Do not power off the system even if power cycle required                             |  

### Batch use

To support batch use, user input can be replaced with environment variables set in the environment.

| Parameter          | Description                                                  |  
|--------------------|--------------------------------------------------------------|  
| BMC_IP             | IPv4 address of the BMC.                                     |  
| BMC_USERNAME       | Username for logging into the BMC.                           |  
| BMC_PASSWORD       | Password used to log into the BMC.                           |  
| BKC_PLDM           | Name of the PLDM file for the BKC.                           |  
| BKC_TARGET_VERSION | BKC version the system will update to, example 01.25.03.04.  |  

