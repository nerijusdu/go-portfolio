# VESA

VESA (Very Simple Sys Admin) is a tool to help you deploy and manage docker containers.

## What can it do?

### Containers and Networks

You can see a list of currently running containers, stop them, delete them or create new ones.
You can also see the basic details for the containers like logs, environment variables and etc.

VESA also alows you to manage docker networks and connect them to containers.

### Templates

This feature is the reason VESA exists, it allows you to create templates of containers and save them for later uses, simmilar to docker-compose.
While docker-compose focuses on making multiple containers work with each other, VESA currently focuses on single container apps, although you can still connect them with networks.
Another advantage (or disadvantage to some) over docker-compose is that VESA is controlled with a web GUI, so no more SSHing into your server to add an environment variable.

### Deployments

VESA provides a github action that connects to VESA API and deploy your changes. All you have to do is have a dockerfile, copy paste the github action files and add required github secrets.


### Routing

VESA uses [Traefik](https://doc.traefik.io/traefik/) as a reverse proxy to route traffic to containers or host services. It also allows you to setup SSL certificates for your domains very easily.

## Why does this exist?

I have quite a few side project and it got really annoying to deploy new ones. It was a lot of copy pasting, a lot of SSHing into the server and a lot of looking into previous projects to remember how I did this before.
I really liked how, for example in Azure, you can just login to the control panel and create a web app or spin up a database. Thats what I tried to do - some kind of control panel for your server.

## Tech details

Web UI is built with my standard frontend stack:

- Typescript
- React
- Chakra UI
- React Query
- React Hook Form

Backend is build with Go and is using [Docker Engine SDK](https://docs.docker.com/engine/api/sdk).
I'm not using any database, just saving everything to json files, because there's not a lot of data and it's easier to share or duplicate your setup.

