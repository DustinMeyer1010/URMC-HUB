First I will have to get the Go package for building it

```shell
go install github.com/akavel/rsrc@latest
```

Then I will have to put the icon in the package

```shell
rsrc -manifest app.manifest -ico URMC.ico -o rsrc.syso
```

Then put the syso file in the main package