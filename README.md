# kaar

Kubernetes Application Archive.

Bundle up a Kubernetes application 📦 into a single static OCI compliant archive. 

 - Search for valid Kubernetes manifests (YAML)
 - Identify references to container images (OCI)
 - Create a single OCI compliant artifact that contains all the application data, and container image data

## Runtime

`kaar` works just like Linux `tar`. 

```
app/
├── deploy.yaml   # References a container image
├── Dockerfile
├── main.go
└── service.yaml


kaar [flags] [archive] [path]
kaar -cf myapp.kaar ./app        Create an archive with container images referenced in deploy.yaml
kaar -xf myapp.karr ./app        Extract an archive with container images referenced in deploy.yaml

 -x Extract
 -f File
 -z Zip
 -c Create
```

### How it works

`kaar` will recursively iterate through every file in the `path` and search for valid Kubernetes YAML.
Next `kaar` will identify all container images referenced from the YAML.
Finally `kaar` will archive the container images (local first, remote next) as well as the YAML from the local directory.
The resulting archive will be saved as an OCI compliant container image that can be uploaded to any container registry.


### .kaar 

Within each `kaar` archive there is a special directory `.kaar` which is used to store raw container image data, and meta information for each archive.
