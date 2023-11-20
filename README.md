# Service Deployment

**Note** Very early in the process so nothing probably works yet

A simpler way to manage your services

## Supported Services

This is compatible with the following:

### Load balancers

- NGINX
  - Static Website
  - Proxy'd Port Forward

## Getting Started

This requires go and nodejs

You must run a npm build first in the frontend directory (bundles frontend in prod)

```sh
# ./frontend
npm run build
```

then

```bash
# ./
go run server.go
```

Then navigate to [http://localhost:5173/](http://localhost:5173/). API requests are proxied to the backend on port 8080.

### Production builds

The server fowards a static directory (react build) bundled with the application

Some misc config notes

```bash
sudo useradd system-control
sudo passwd system-control
# Enter password to create
mkdir /home/system-control
sudo chown -hR system-control /home/system-control

# Grant access to newly created user to manage services
echo 'system-control ALL = NOPASSWD: /etc/init.d/nginx' | sudo tee -a /etc/sudoers.d/nginx

# Get go command location
GOEXEC=$(which go)
su - system-control $VALUE run server.go
# Enter newly created password


```

## ENV

Override environment variables with local.env
