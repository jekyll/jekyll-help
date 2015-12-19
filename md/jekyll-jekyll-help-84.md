# [Jekyll Serve crashing?](https://github.com/jekyll/jekyll-help/issues/84)

> state: **closed** opened by: **** on: ****

Hello, I was playing around with Jekyll and got it installed and functioning properly; however, after I run the command &#x27;jekyll serve&#x27; and the site is accessible, when I close my laptop the site goes down because the IDE loses connection.

I&#x27;m using Nitrous.io

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/84#issuecomment-46981053) on: **Tuesday, June 24, 2014**

I am going to assume you do not expect the computer/site to continue to be accessible when the laptop is closed, sleeping, or turned off.

Are you suggesting/asking, upon reopening/awaking/starting your laptop, that the site is no longer continuing to be served as it was prior to sleep/shutdown?
---
> from: [**andysanchez**](https://github.com/jekyll/jekyll-help/issues/84#issuecomment-46981950) on: **Tuesday, June 24, 2014**

It&#x27;s installed fully on Nitrous.io - here&#x27;s the address and it is currently working (address below). Since nitrous has a built in Web command line it should be running when I close out of it; however, it&#x27;s not- it stops the process completely and the site is kicked off line.

In regards to my laptop being closed or asleep, the site isn&#x27;t being hosted from it, it&#x27;s a remote server.

http://blogsanchez-124746.use1.nitrousbox.com:4000/

![nitrous io 2014-06-24 10-53-37 2014-06-24 10-54-25](https://cloud.githubusercontent.com/assets/7904813/3372993/d3199308-fbaf-11e3-8f97-3aba594b6de1.jpg)

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/84#issuecomment-46986981) on: **Tuesday, June 24, 2014**

Thanks for the additional information.

I am not familiar with the service you refer to. However, I can conclude that if you are hosting remotely within a shell or shell-like environment, you likely simply need to focus on executing long-running commands within the appropriate context.

For example, within *nix and *nix-like environments, there is &#x60;screen&#x60; and &#x60;nohup&#x60;, etc. that put commands into the background while continuing to run... *even after disconnection*.

Short of that, I would get with the third party service that is providing your hosting platform and request more information about how to accomplish what you desire.
---
> from: [**andysanchez**](https://github.com/jekyll/jekyll-help/issues/84#issuecomment-46987500) on: **Tuesday, June 24, 2014**

Ahhh... I just noticed something in another thread based off your response and just figured it out, woohoo thank you!! Nitrous is a newer platform that there wasn&#x27;t any documentation really. 

To fix this and keep your jekyll server running, switch to the directory of your app and instead of just booting the server you have to tell it to run in the background by doing this:
&#x60;&#x60;&#x60;
jekyll serve --detach
&#x60;&#x60;&#x60;

And you&#x27;ll be up and running instantly. You will need to kill all to stop it, or just make note of the process id.

Thanks for the help!
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/84#issuecomment-46988348) on: **Tuesday, June 24, 2014**

Absolutely; you&#x27;re welcome. Glad you got things figured.
