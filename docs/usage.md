### Usage

[See command reference](cmdref.md)

**JSON** can be parsed in following ways:

```bash
curl -s https://httpbin.org/get | sspp --json='headers.Host'

# you can pass data directly with --data flag
sspp --json='data.email' --data='{"data": {"email": "sample@example.org"}}'

# It is possible to pass default data directly with --or flag
sspp --json='data.nonExistent' --data='{"data": {"email": "sample@example.org"}}' --or="nil"
curl -s https://httpbin.org/get | sspp --json='headers.nonExistent' --or="nil"
```

---

**XML** can be parsed in following ways:

```bash
curl -s https://httpbin.org/xml | sspp --xml='slideshow.slide.0.title'

# you can pass data directly with --data flag
sspp --xml='numbers.number.0' --data='<numbers><number>1</number><number>2</number></numbers>'
```

---

**YAML** can be parsed in following ways:

```bash
curl -s https://raw.githubusercontent.com/istio/istio/release-1.3/samples/httpbin/httpbin.yaml | sspp --yaml='spec.ports.0.port'
```

Another example of `.yaml` or `.yml` configuration file parsing

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
cat controllers/nginx-deployment.yaml | sspp -Y='spec.template.spec.containers.0.image'
```

---

**TOML** can be parsed in following ways:

```bash
curl -s https://raw.githubusercontent.com/Praqma/helmsman/master/examples/example.toml | sspp --toml="metadata.org" --or='nil'
```

Another example of `.toml` configuration file parsing

```toml
# praqma/helmsman.toml
...

# define your environments and their k8s namespaces
# syntax:
# [namespaces.<your namespace>] -- whitespace before this entry does not matter, use whatever indentation style you like
# protected = <true or false> -- default to false
[namespaces]
  [namespaces.production]
    protected = true
    [[namespaces.production.limits]]
      type = "Container"
      [namespaces.production.limits.default]
        cpu = "300m"
        memory = "200Mi"
      [namespaces.production.limits.defaultRequest]
        cpu = "200m"
        memory = "100Mi"
    [[namespaces.production.limits]]
      type = "Pod"
      [namespaces.production.limits.max]
        memory = "300Mi"
  [namespaces.staging]
    protected = false
    [namespaces.staging.labels]
      env = "staging"
    [namespaces.staging.quotas]
       "limits.cpu" = "10"
       "limits.memory" = "30Gi"
       pods = "25"
       "requests.cpu" = "10"
       "requests.memory" = "30Gi"
       [[namespaces.staging.quotas.customQuotas]]
         name = "requests.nvidia.com/gpu"
         value = "2"
...
```

Then in command-line

```bash
cat praqma/helmsman.toml | sspp -T='namespaces.production.limits.0.default.cpu'
```

---

**INI** file can be parsed in following way:

```bash
curl -s https://raw.githubusercontent.com/emoncms/emoncms/master/example.settings.ini | sspp -I 'redis.enabled' --or='nil'
```

Another example of `.ini` configuration file parsing

```ini
# dist/db_config.ini

;debug=1
;default_action=home
;google_translate_url="http://weblite-dns2.com/proxy.php"
google_translate_url="http://ec2-75-101-244-123.compute-1.amazonaws.com/proxy.php"
title="Web Lite Translate"
default_price_per_word=0.15

;;Configuration settings for application
title="translation_weblite_ca"
scriptUrl="http://translation.weblite.ca/index.php"
multilingual_content=1

[_database]
    host="localhost"
    name="mydb"
    user="mydbuser"
    password="foo"

[_tables]
    webpage_sites="Websites"
    translations = "Translations"
    packages="Packages"
    users=Users
    proof_jobs="Jobs"
    webpage_status="Webpages Status"

[_auth]
    users_table=users
    username_column=username
    password_column=password
    secret_code="ljkasdfjkldsafliasdoiudsfoi"
    allow_register=1
    session_timeout=999999999
```

Then in command-line

```bash
cat dist/db_config.ini | sspp -I='_auth.username_column'
```
