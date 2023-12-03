# Iterative Dev Strategy

This is a lot of work! Everything can't be thought of at once, or progress would be very slow and this would be unusable for a long time. That's why we use a MVP to start. This is changable and more a draft to organize thoughts until this gets to a point where tasks are moved to a task board.

## Phase 1 - Minimal MVP, Pre-Release

What does this mean?

The UI won't be great, but we can deploy a website in a single use case. Some other code might be scaffolded, but not robust.

Definition of MVP:

1. We have a website bundled into the API for a single deployed executable
2. The main page is a list of existing deployments
3. You can create a deployment. Just a name, and for now only a static website
4. You can create environments. IE dev, staging, production
5. You can upload a ZIP folder to be deployed (IE a bundle after running `npm build`)
   - Bundle will be unzipped in the active environment directory and the previous bundle will be moved out and deleted on completion of deployment
6. This file will be stored indefinitely, and is referenced. Deleting is out of scope for now
7. You can create a URL to point to an environment
   - Create an NGINX config file
   - Generate a symlink from the config file to the sites-available directory of NGINX
   - Generate a symlink from the config file to the sites-enabled directory
   - Restart the NGINX server
   - Create a new URL on the DNS server
8. This data should exist within a SQLite database (for now)

### Intended use case

You can host this internally, but this should not be exposed to the public internet

Once main flow is complete, can create a tag/release

## Phase 2 - Securing, access controls, API keys & GitHub integration, API Docs

NOTE - Now that the app works, we can run releases as features are created. Phases can be major releases or minor. Still flushing out, less important right now haha

Validate and generate JWT tokens/API keys, store passwords with BCrypt, users can create API keys.

Access controls:

- create/delete/update services
- create/delete/update environments
- create/delete/update deployments from history
- create/delete/update new deployment (upload zip)
- create/delete/update users
- create/delete/update API keys

This should make it "safeish" to be visible on the web.

### Intended use case

You can connect to a GitHub/GitLab hook to send a zip/tarball to be decompressed on demand. And rollback as needed.
With better API docs, users can understand this somewhat better.

## Phase 2.5 - Admin

Tidy code some, open-source-ify; create tasks, UI designs, tag tasks for releases, look into licensing, etc.

## Phase 3 - External service control, logging

Update keys in the browser in a config screen (DNS, Load Balancer settings)

Create a UI settings page that can manage services. For now, we're only using Cloudflare. Maybe some other collaborator will want to add another DNS provider.. We can add auth for DNS (encrypted somehow), and manage load balancer settings.

Also - add UI logs for the process of releasing. We want details! Running deployment from date x on env y. Unzipping. DNS update. Done. Any errors are available there.

## Phase 4 - Kubernetes basic deployments, code cleanup

Allow a container to be deployed from Kubernetes.

User sets details on the settings page like for DNS.

For each deployment, user specifies a docker container and plaintext env variables (for now) to be set during runtime in that container. That container is then deployed. Users can specify a set of ports in the browser. Users can stop, start (if stopped), and restart a container. Only one pod per deployment (for now)

## Phase 5 - Kubernetes Services

Storage buckets, databases, vault environment variables
Scaling options for kubernetes services

To be refined

## Phase 6 - PWAify and UI Revamp

Hope to have a real name by now haha
Make it look good, work offline-ish, notify on successful deployment,
