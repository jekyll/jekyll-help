# [Running multiple Jekyll Servers?](https://github.com/jekyll/jekyll-help/issues/85)

> state: **closed** opened by: **** on: ****

Hello, I&#x27;m extremely pleased with running Jekyll on my own server so far. I was thinking about setting up a few sites for my co-workers to play around with and was wondering what suggestions there are on hosting multiple Jekyll sites?

Ideally, I&#x27;d want to use the same EC2 box to run 3 or 4 Jekyll instances. I&#x27;d assume I just tell it to run on separate ports and handle the port mapping with virtual hosts?

I couldn&#x27;t find any info on this and figured this is the best way to answer, thanks for the pointers :-)

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/85#issuecomment-47137460) on: **Wednesday, June 25, 2014**

Hey Andy,

I&#x27;m glad you&#x27;ve been pleased running Jekyll on your own server so far! What I&#x27;d suggest is that you use a stock server (like nginx or apache) and use Jekyll to build your static site to a different directory for each site. On my server, I have a git post-receive hook which takes the repo, clones it to &#x60;/tmp&#x60;, and builds it to my &#x60;/var/www&#x60; directory for that site. That way, deploying my site is as easy as &#x60;git push&#x60;.

## Example &#x60;post-receive&#x60; hook

&#x60;&#x60;&#x60;bash
source /home/git/.bash_profile # the git user has jekyll installed

REPO_NAME=example.org
GIT_REPO=$HOME/repositories/$REPO_NAME.git # i use gitolite
TMP_GIT_CLONE=/tmp/$REPO_NAME
PUBLIC_WWW=/var/www/$REPO_NAME

git clone $GIT_REPO $TMP_GIT_CLONE

cd $TMP_GIT_CLONE &amp;&amp; jekyll build -s $TMP_GIT_CLONE -d $PUBLIC_WWW --trace
rm -Rf $TMP_GIT_CLONE
exit
&#x60;&#x60;&#x60;

## Example &#x60;nginx&#x60; config (goes in &#x60;sites-enabled&#x60;)

&#x60;&#x60;&#x60;nginx
server {
  listen 80;
  server_name example.org;
  root /var/www/example.org;
}
&#x60;&#x60;&#x60;

If you set up one of these post-receive hooks and nginx configurations for each repository, you should be able to use just one aws ec2 instance and one server (nginx) to manage all your jekyll sites.
---
> from: [**andysanchez**](https://github.com/jekyll/jekyll-help/issues/85#issuecomment-47142357) on: **Wednesday, June 25, 2014**

Thanks Parker, BTW I checked out your site, the Berlin map looks great!!

What you said above makes sense, however, rather than pulling from Git Hub I&#x27;ll be storing the posts on our server, which is why I initially planned on just running them on different ports. BUT running one sounds much better and having ngix do the heavy lifting.

I&#x27;m assuming I can look around and track down a bash that keeps an eye out for changes to particular directories and have it build jekyll accordingly?

I just cloned my directory and fired up two sites with &#x60;&#x60;&#x60;jekyll serve&#x60;&#x60;&#x60; in their root directory. Immediately I get the warning there is no config file in the directory, because it&#x27;s in the site folders of that directory. I&#x27;m assuming I need to update the path in each sites yml, but along those lines Is there a yml i can create to specify where to find each site?


Edit: I apparently just learned how to spell and my post reflects that.

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/85#issuecomment-47247229) on: **Thursday, June 26, 2014**

&gt; Is there a yml i can create to specify where to find each site?

@andysanchez There are a couple options:

1. Include a &#x60;source&#x60; flag in the &#x60;jekyll serve&#x60; command
2. Include a &#x60;source&#x60; option in &#x60;_config.yaml&#x60;

The documentation for this can be found on the [Configuration](http://jekyllrb.com/docs/configuration/) page of the Jekyll product site.

PS: I don&#x27;t think you should not be using &#x60;jekyll serve&#x60; to host a production website. Ideally, you should use &#x60;jekyll build&#x60; along with a &#x60;destination&#x60; flag/option to send the processed files to a directory that is being hosted by Apache, nginx, or some other dedicated HTTP service.
---
> from: [**andysanchez**](https://github.com/jekyll/jekyll-help/issues/85#issuecomment-47249437) on: **Thursday, June 26, 2014**

Thanks Troy, the source command did it!

I set up my virtual host file to serve the appropriate domain and used &#x60;build&#x60; &amp; &#x60;destination&#x60; to get it in the correct Apache directory. 

Now I suppose only thing I have left is to make a bash script that looks for new posts and builds the correct Jekyll site and processes to correct Apache destination.

...I will say though, after the learning curve of getting this up and running I am so much happier with this than the mess that the database driven CMS (namely the PHP ones) creates.

If anyone has an example of a bash script to handle re-building after a new post is uploaded to a directory that&#x27;d be a huge help. Thanks for the responses Parker &amp; Tommy, I got this covered!

*presses close and does dance*
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/85#issuecomment-47249916) on: **Thursday, June 26, 2014**

:metal: ROCK ON!

I would *highly* encourage you to write about your exploits. This is a tremendously important workflow that *a lot* of people will use, so the more lessons and walkthroughs that are out in the wild, the better. :+1: 
---
> from: [**andysanchez**](https://github.com/jekyll/jekyll-help/issues/85#issuecomment-47250882) on: **Thursday, June 26, 2014**

I absolutely will! I wound up here after years of headaches with Drupal &amp; ExpressionEngine, luckily I managed to stay away from Wordpress. I was attracted to Jekyll because I&#x27;m very good at building Shopify themes, so my comfort with Liquid was a major draw. Plus after testing the same page, with each respective platform, I found that Jekyll&#x27;s time to paint was around 60 ms, the same Wordpress page was close to 2.5 seconds.

As soon as I saw that the same page, on the same server, was loading over 40 times as fast I was immediately sold. I know as I add complexities to it that&#x27;ll go down, but it is a much more logical workflow.

I&#x27;ll still need to figure out a solution for more in depth contact forms, but I&#x27;m very excited to get my blog running and seeing what I can contribute to the community, namely along the lines of easing people into Liquid.

Thanks again!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/85#issuecomment-47251530) on: **Thursday, June 26, 2014**

&gt; I&#x27;ll still need to figure out a solution for more in depth contact forms

@andysanchez Sounds like a great topic for a [new issue](https://github.com/jekyll/jekyll-help/issues/new)!
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/85#issuecomment-47601136) on: **Monday, June 30, 2014**

&gt; rather than pulling from Git Hub I&#x27;ll be storing the posts on our server, which is why I initially planned on just running them on different ports

I don&#x27;t pull from GitHub for this, I push to my own server with gitolite installed (though that detail is unnecessary). I use the SSH protocol (&#x60;git@myserver.com:repo.git&#x60;) and the &#x60;post-receive&#x60; hook is executed when the git server has received the entire push contents. It&#x27;s super convenient and totally customizable.

&gt; I&#x27;m assuming I can look around and track down a bash that keeps an eye out for changes to particular directories and have it build jekyll accordingly?

You could do an upstart process to run &#x60;jekyll build -w&#x60;, maybe? I&#x27;ve always relied on the system I described before to build whenever the git repo receives an update.

&gt; I&#x27;m assuming I need to update the path in each sites yml, but along those lines Is there a yml i can create to specify where to find each site?

I&#x27;d suggest running &#x60;jekyll build&#x60; or &#x60;jekyll serve&#x60; with the &#x60;--source&#x60; and &#x60;--destination&#x60; switches for this!
