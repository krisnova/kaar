# Kubernetes Application Archive

```
                                                                                
.------------------------------------------------- ------- -- -  -         hX!  
|  ______. ____       __________         __________         ____________        
: |      |/   /     _/     _    \___   _/     _    \___   _/     _      \       
| |      /   /____  \      \X       \  \      \X       \  \      \\     /__     
`-|      T\       \--\_      \       \--\_      \       \--\_     /_       \--- 
  |______| \       \  /_______\       \  /_______\       \  /_______\       \   
            \_______\\         \_______\\         \_______\\         \_______\\ 

```

Bundle up a Kubernetes application into a single static OCI compliant archive. 

## Flags

``` 
 -x Extract
 -f File
 -z Zip (gzip)
 -c Create
```

### Archive

`kaar` will recursively iterate through every file in the `path` and search for valid Kubernetes YAML.
Next `kaar` will identify all container images referenced from the YAML.
Finally `kaar` will archive the container images (local first, remote next) as well as the YAML from the local directory.
The resulting archive will be saved as an OCI compliant container image that can be uploaded to any container registry.

```bash 
kaar [flags] [archive] [path]
kaar -cf myapp.kaar ./app/data
```

### Unarchive

`kaar` will preserve state by default.
`kaar` can unarchive a `kaar` file into a local path.

```bash 
kaar [flags] [archive] [path]
kaar -xf myapp.karr ./app/data
```

### .kaar 

Within each `kaar` archive there is a special directory `.kaar` which is used to store raw container image data, and meta information for each archive.