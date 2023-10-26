# My server setup

If you're like me, you probably have a lot of unused domains and a lot of unfinished side projects. So today, I thought I would share how I host my side projects, why I use a VPS instead of cloud providers, and why I reinvented docker-compose. I have quite a few side projects, so [check them out if you haven't](/projects).

## VPS vs Cloud

To host my side projects, I rent a VPS instead of using cloud hosting mainly because of the costs. Currently, all my side projects make me 0 money and they have about the same amount of users. It doesn't really make sense to have infinitely scalable serverless quantum edge functions to serve my portfolio for 3 monthly views (2 of which is me checking if my SSL cert is not expired).

![Serverless is the future meme](/data/images/serverless.jpg)

But, Nerijus, I hear you ask, most cloud providers give you a free tier for your projects, why not use that? Well, because they are extremely slow! I know I just said I don't have any users using my projects, but it's so slow even when I'm the only person using it, it's almost useless if more users try to use it.

I pay around 40€/year for my VPS, it has 2 cores, 2GB of RAM, 40GB SSD space, and 2TB bandwidth.
It is running these things (mostly in docker containers) with no trouble:

- 3 Go apps/sites
- 2 .NET sites
- 1 Next.JS site
- Ghost blog
- Postgres database
- PgAdmin
- Docker registry
- Plausible analytics server
- Grafana

Another advantage of using a VPS over cloud hosting is the learning experience that comes with it. By using a VPS you will learn about server configuration, security measures, networking and so much more. This experience can be incredibly valuable.

## VPS Setup

Now I will assume that I convinced you that VPSes are great and tell you about my setup.

I run almost everything in docker except nginx (could probably move this to docker too at some point). I have my own docker registry that has all my containers and I used to have docker-compose files for every project. 

![Containers everywhere meem](/data/images/containers-everywhere.jpg)

For deployments I used to have github actions that build and push docker image to my registry, then SSH into the VPS and pull the latest image.

This is obviously not ideal and it's pretty cumbersome to add new projects. So I built [VESA (Very Simple Sys Admin)](/projects/vesa) that would ease my pain. Now I can add new projects through a web interface, copy paste a github action and update nginx config (will try to include this in VESA too).

## Conclusion

With all that said, I'm not trying to make you ditch the cloud, that's just the setup that works for me and my projects. With VESA I'm trying to get the convenience of the cloud with the benefits of a VPS.
