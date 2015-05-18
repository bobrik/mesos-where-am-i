# Where am I on Mesos?

This docker container responds with `HOST` and `PORT` environment variables
and HTTP path to every HTTP request. It is only useful for demos and debugging.

Running without Mesos:

```
docker run --rm -it --net host -e HOST=127.0.0.1 -e PORT=9999 bobrik/mesos-where-am-i
```

Then check:

```
$ curl http://127.0.0.1:9999/aaa
Hello, "/aaa" from 127.0.0.1:9999
```
