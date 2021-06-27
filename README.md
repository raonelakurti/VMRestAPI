# VMRestAPI to Perform GET, CREATE, DELETE and Update operations.

Run Locally:
I have used go mod to maintain dependencies.

$ go build
$ ./VMRestAPI

use below to perform do list, create, update and destroy

Get Existing VM's info: http://localhost:9000/vmlist
Create a new VM: http://localhost:9000/createvm
Update VM based out of UUID: http://localhost:9000/updatevm
Delete a VM based on the UUID: http://localhost:9000/deletevm

Usage:
POST Operation:
POST /createvm

```
{
    "adminUsername": "azureUser",
    "password": "Welcome!23",
    "vmName": "AKS-VM1",
    "vmSize": "Standard_D2_v3",
    "region": "us-chicago",
    "osImage": "linux"
  }
```

Output: 
```
{
"id": "0e3d75cf7e444b7c9a2357199f640cb3",
"adminUsername": "azureUser",
"password": "Welcome!23",
"vmName": "AKS-VM1",
"vmSize": "Standard_D2_v3",
"region": "us-chicago",
"osImage": "linux"
}
```

Update VM password:
PUT /updatevm

```
{
"id": "0e3d75cf7e444b7c9a2357199f640cb3",
"adminUsername": "azureUser",
"password": "Welcome@23",
"vmName": "AKS-VM1",
"vmSize": "Standard_D2_v3",
"region": "us-chicago",
"osImage": "linux"
}
```

Output:

```
{
"id": "0e3d75cf7e444b7c9a2357199f640cb3",
"adminUsername": "azureUser",
"password": "Welcome@23",
"vmName": "AKS-VM1",
"vmSize": "Standard_D2_v3",
"region": "us-chicago",
"osImage": "linux"
}
```

GET Operation to list VMs:
GET /vmlist

OUTPUT:

```
{
"id": "0e3d75cf7e444b7c9a2357199f640cb3",
"adminUsername": "azureUser",
"password": "Welcome@23",
"vmName": "AKS-VM1",
"vmSize": "Standard_D2_v3",
"region": "us-chicago",
"osImage": "linux"
}
```

DELETE VM Based out of ID:

DELETE /deletevm
```
{
    "id": "0e3d75cf7e444b7c9a2357199f640cb3"
}
```

Output: 200 Response
