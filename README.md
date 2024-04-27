# OCI container with Bazel

This small repo is meant to demonstrate an HTTP server container which is super easy to build and deploy.

Please read the [accompanying article](http://popovicu.com/posts/containers-bazel-one-command) on my website for more details.

## Building

```
bazel build //container
```

The output is a tarball that can further be loaded into something like Docker and executed.

```
...
INFO: Found 1 target...
Target //container:container up-to-date:
  bazel-bin/container/container/tarball.tar
...
```

For Docker, load like this:

```
docker load < tarball.tar
```

Run the container with port forwarding to `12345`:

```
docker run -p 12345:12345 --name myfoo popovicu.com/foo:latest
```

Test with `curl`:

```
curl http://localhost:12345/hello/world
```

Output:

```
Hello world
```