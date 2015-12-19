# [Jekyll + Docker = Awesome?](https://github.com/jekyll/jekyll-help/issues/141)

> state: **closed** opened by: **** on: ****

Hey everyone,

I just wanted to start a little discussion about using Jekyll in a Docker container. Has anyone done this? What kind of pitfalls should I expect? Should I use a dedicated web server? Which makes more sense: Nginx, or Apache? Will &#x60;jekyll serve&#x60; running WEBrick be good enough?

I&#x27;m approaching this as a guy who has never really run any Docker stuff in production, so any extra help on this would be awesome and will probably be helpful in the future. Thanks!

### Comments

---
> from: [**oblakeerickson**](https://github.com/jekyll/jekyll-help/issues/141#issuecomment-53893171) on: **Friday, August 29, 2014**

I&#x27;m just getting started with docker myself, but I would be interested in this as well. I currently have a discourse server which uses Docker containers and then I have a completely separate server just for my jekyll blog (discourse is hosting the comments). It might be good to put them on the same server though and just put jekyll in its own container.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/141#issuecomment-53919223) on: **Friday, August 29, 2014**

Here&#x27;s how I see it working, but I don&#x27;t know nearly enough about Docker&#x27;s best practices to know if this makes sense or flies in the face of everything that is good and true about the system:

1. Write a Dockerfile that:
   1. Starts from a standard issue base image (like Ubuntu)
   2. Uses &#x60;apt-get&#x60; to update itself and install Nginx
   3. Configures Nginx to host a static site (cache control, etc)
   4. Installs Ruby and the necessary Gems
   5. Pulls down the source files for the Jekyll site
   6. Runs &#x60;jekyll build&#x60; on the source
   7. Copies the destination &#x60;_site&#x60; folder to the folder that Nginx expects
3. Commit the image
4. Pull the updated image on whatever server it&#x27;s running on

The problems that I see with this have to do with networking. I don&#x27;t really know how to link up port 80 on the host machine with whatever port Nginx would be running on. Also, I know nothing about DNS stuff, so I&#x27;m kind of lost there, too.
---
> from: [**cpuguy83**](https://github.com/jekyll/jekyll-help/issues/141#issuecomment-53920569) on: **Friday, August 29, 2014**

Please don&#x27;t use webrick for anything, ever... it&#x27;s horribly slow, single-threaded, not evented, etc. :)
I&#x27;d be happy to talk about docker-ish things in #docker on freenode.

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/141#issuecomment-54923067) on: **Monday, September 08, 2014**

Well, this worked out really awesome. Here&#x27;s some stuff that I made so far. Eventually, I&#x27;ll write a blog post about how I got this stuff up and running. In the meantime, if anyone has any questions, feel free to hit me up.

**Base Image**

* https://registry.hub.docker.com/u/troyswanson/jekyll
* https://github.com/troyswanson/docker-jekyll

**Dockerfile for my Jekyll site**

* https://github.com/troyswanson/troyswanson.github.io/blob/master/Dockerfile
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/141#issuecomment-56446807) on: **Monday, September 22, 2014**

I wrote a blog post about this topic last night. Will write a second one with even more info later!

Check it out if you&#x27;re interested: http://troy.im/2014/09/21/dockerizing-jekyll/
---
> from: [**jhabdas**](https://github.com/jekyll/jekyll-help/issues/141#issuecomment-98557750) on: **Sunday, May 03, 2015**

@troyswanson I think I&#x27;ve got what you&#x27;re looking for: [Simple websites with Jekyll and Docker](http://habd.as/simple-websites-jekyll-docker). In this post I explain how to use OS X or Windows to build a Jekyll-based docker site locally. Includes Dockerfile and instructions for deploying the image to a remote server via the command line.
