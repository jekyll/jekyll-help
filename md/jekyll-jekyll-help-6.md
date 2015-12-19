# [Implementing shorturls?](https://github.com/jekyll/jekyll-help/issues/6)

> state: **closed** opened by: **** on: ****

What is the best and/or easiest way to create base16 shorturls (short urls) mapping from your blog posts to your own domain without relying on services like bit.ly—which don’t protect your links, should the company fail or get acquired?

How have you implemented them on your site?

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39100977) on: **Monday, March 31, 2014**

Hey @ndarville - your question has several avenues that end in different results, depending on which path you choose to travel.

The *easiest* way is pretty obvious: use a system that already exists ([bitly](http://bit.ly), [tinyurl](http://tinyurl.com), [is.gd](http://is.gd), etc.) to create the short URLs. These services - like anything else on the internet - have the possibility of going under, which would cause those shortened URLs to disappear.

Another way would be to host your own URL shortening service. It&#x27;s way more involved, but obviously gives you total control over the system. The popular self-hosted URL shorteners ([YOURS](http://yourls.org/) and [phurl](https://code.google.com/p/phurl/)) are PHP scripts that live on some server that you are in charge of.

Here&#x27;s another idea: if you are using GitHub Pages and have a server that you&#x27;d like to use to help automate the creation of these short URLs, take a gander at the &#x60;page_build&#x60; [webhook](http://developer.github.com/webhooks/). This hook fires when GitHub publishes your site to their Pages platform, giving you the ability to create a script that crawls the site as soon as it is built and dynamically create shortened URLs as it come across new pages.

Which route you decide to take depends on the level of effort you wish to trade off for the amount of control you want to have over the system.

Personally, I think URL shorteners are very rarely necessary. The only time I need to shorten a link is when I&#x27;m sending a tweet, but Twitter does that automatically through its [t.co](http://t.co) service.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39111724) on: **Monday, March 31, 2014**

@troyswanson @ndarville what about something with &#x60;jekyll-redirect-from&#x60; ? I&#x27;ve never used that gem but from my understanding it does redirects from different URLs. I may also be completely wrong.  
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39112499) on: **Monday, March 31, 2014**

Same, I&#x27;ve never used it, but I believe it deals with creating redirects within your site, which is only for your own domain. For instance, setting up a &quot;vanity URL&quot; like &#x60;http://example.com/a-new-hope&#x60; which would redirect to &#x60;http://example.com/2014/4/1/star-wars-a-new-hope-review/&#x60;.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39114090) on: **Monday, March 31, 2014**

@troyswanson yes, but it may work for what @ndarville wants; and as you said, that be, e.g. &#x60;http://example.com/XdFdE&#x60; would redirect to &#x60;http://example.com/2014/4/1/star-wars-a-new-hope-review/&#x60; he will just need a rake task or whatever script to generate a draft with the &#x60;redirect_from:&#x60; and the base64.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39116510) on: **Monday, March 31, 2014**

It would be awesome to have some special keyword built into https://github.com/jekyll/jekyll-redirect-from ... :)
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39238627) on: **Tuesday, April 01, 2014**

@parkr On second thought, maybe not, huh?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39239953) on: **Tuesday, April 01, 2014**

Yurrrrp. 
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39240200) on: **Tuesday, April 01, 2014**

So the parting advice is to not involve Jekyll in the process at all?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39242260) on: **Tuesday, April 01, 2014**

My parting advice would be to use jekyll-redirect-from and use more semantic short URL&#x27;s, rather than the base 64 encoded gobbledy-gook we love so much. ;) E.g. My post &quot;The Crime of Net Neutrality and What We Can Do to Fix It&quot;, I&#x27;d probably shorten that to &#x60;redirect_from: /net-neutrality&#x60; or something. Or, if you wish, you can even make up a lil set of nonsense chars: &#x60;redirect_from: /b32JakI&#x60;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/6#issuecomment-39242373) on: **Tuesday, April 01, 2014**

Bingo. :raised_hands: 
