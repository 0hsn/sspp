### Usage

[See command reference](cmdref.md)

**JSON** can be parsed in following ways:

```bash
$ curl -s https://httpbin.org/get | sspp --json='headers.Host'

# you can pass data directly with --data flag
$ sspp --json='data.email' --data='{"data": {"email": "sample@example.org"}}'

# It is possible to pass default data directly with --or flag
$ sspp --json='data.nonExistent' --data='{"data": {"email": "sample@example.org"}}' --or="nil"
$ curl -s https://httpbin.org/get | sspp --json='headers.nonExistent' --or="nil"
```
---

**XML** can be parsed in following ways:

```bash
$ curl -s https://httpbin.org/xml | sspp --xml='slideshow.slide.0.title'

# you can pass data directly with --data flag
$ sspp --xml='numbers.number.0' --data='<numbers><number>1</number><number>2</number></numbers>'
```
---

**YAML** can be parsed in following ways:

```bash
$ curl -s https://raw.githubusercontent.com/istio/istio/release-1.3/samples/httpbin/httpbin.yaml | sspp --yaml='spec.ports.0.port'
```

Another example of yaml configuration file

```yaml
# controllers/nginx-deployment.yaml

apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```
Then in command-line

```bash
$ cat controllers/nginx-deployment.yaml | sspp -Y='spec.template.spec.containers.0.image'
```
