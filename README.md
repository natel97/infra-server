# Service Deployment

**Note** Very early in the process so nothing probably works yet

A simpler way to manage your services

## Dream State - The Story

I want to host on my Raspberry Pi. But I want things to be updated seamlessly. No hassle, no fuss. I push code, code is shown on my website, or any website I purchase.

But wait, what if I decide to get into something deep in development? Now I need to create an API, connect databases, file buckets, Redis, etc. UGH this is a lot to connect. Especially if I want to keep to open source an my Pi. The point of having a Raspberry Pi is I only pay for electricity. Not cloud services.

In comes this - a way manage custom services. I can create a service (static website or kubernetes deployment), and manage my website domains that point to a folder in NGINX or proxy forward to an exposed kubernetes deployment. Or maybe I just want an internal API to keep in the shadows, only accessible to other services I've deployed.

Now we're in the weeds. That's a LOT to develop at once. Hence why this is broken into [different phases](docs/ITERATIVE_DEV_STRATEGY.md) (no need for thrashing :D)

## Getting Started

This requires go and nodejs

You must run a npm build first in the frontend directory (bundles frontend in prod, embed requires those files to exist)

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

### UI Sketches (incomplete) for the future

https://www.figma.com/file/aFc0Qj29k4HX9XD93pzXeH/Service-management-UI
