### Usage

**JSON** can parsed on following ways:

```bash
$ curl -s https://httpbin.org/get | sspp --json='headers.Host'

# you can pass data directly with --data flag
$ sspp --json='data.email' --data='{"data": {"email": "sample@example.org"}}'

# It is possible to pass default data directly with --or flag
$ sspp --json='data.nonExistent' --data='{"data": {"email": "sample@example.org"}}' --or="nil"
$ curl -s https://httpbin.org/get | sspp --json='headers.nonExistent' --or="nil"
```
---

**XML** can parsed on following ways:

```bash
$ curl -s https://httpbin.org/xml | sspp --xml='slideshow.slide.0.title'

# you can pass data directly with --data flag
$ sspp --xml='numbers.number.0' --data='<numbers><number>1</number><number>2</number></numbers>'
```
