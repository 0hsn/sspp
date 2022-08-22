## Command reference

```bash
$ sspp -j # json
$ sspp -y # yaml
$ sspp -t # toml
$ sspp -h # hcl
$ sspp -x # xml
$ sspp -w # jwt
$ sspp -i # ini
$ sspp -p # .properties
```

```bash
$ '{"data": [1, {"name": "some"}, true]}' | sspp -j='.data.1.name' --or='nil'
```

```bash
$ sspp -j='.data.1.name' --or='nil' --data='{"data": [1, {"name": "some"}, true]}'
```
