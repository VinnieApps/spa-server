# SPA Server

Simplest server for single page applications.

# What does it do?

Serve static files with a simple rule:

- Is it a root request (e.g.: `/` or `/requests`)? -> Return `/index.html`
- Is it a directory? -> return `index.html`
- File doesn't exist? -> Return `index.html`
- Otherwise, serve file

# How do I use it?

Create a Dockerfile with the following content:

```Dockerfile
FROM scratch

RUN mkdir -p /opt/app

# Copy spa-server
COPY spa-server /opt/app/spa-server
# Copy your SPA content
COPY dist /opt/app

# Start server
WORKDIR /opt/app
ENTRYPOINT ["/spa-server"]
```

Then download the latest version and build your Docker image:

```
$ curl -L 'https://github.com/VinnieApps/spa-server/releases/latest/download/spa-server_linux_amd64' > spa-server
$ chmod +x spa-server
$ docker build -t my-image .
```

Or do it all from inside Docker using multistate Dockerfiles:

```Dockerfile
FROM ubuntu AS downloader
RUN apt-get update && apt-get install -y curl

RUN mkdir -p /tmp/build/opt/app

RUN curl -L 'https://github.com/VinnieApps/spa-server/releases/latest/download/spa-server_linux_amd64' > /tmp/build/opt/app/spa-server
RUN chmod +x /tmp/build/opt/app/spa-server

FROM scratch

COPY --from=downloader /tmp/build /
COPY dist /opt/app

WORKDIR /opt/app
ENTRYPOINT ["./spa-server"]
```
