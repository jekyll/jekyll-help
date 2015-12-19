# [How do I auto-start jekyll?](https://github.com/jekyll/jekyll-help/issues/212)

> state: **open** opened by: **** on: ****

First I have to commend the jekyll team on a great product! Installation and usage is extremely easy, for the most part. Thanks for all of your hard work!

That said, after nearly two-hours of very-frustrated searching, I can&#x27;t find any tips on how to auto-start Jekyll. It&#x27;s obviously not reasonable for me to have to type into terminal &quot;jekyll serve --source /var/www/my-awesome-site&quot; every single time my server needs to reboot. So how the heck do I auto-start jekyll? I&#x27;m new to upstart but over an hour of scripting hasn&#x27;t yielded any results.

Can someone please breathe some life into this process before I have to switch to another blogging solution, simply because I can&#x27;t get it to auto-start? Thanks!!!!

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/212#issuecomment-67527836) on: **Thursday, December 18, 2014**

Glad you are enjoying Jekyll! It is indeed wonderful.

Regarding serving and auto-starting... Jekyll leverages Ruby&#x27; Webrick web server for local development serving. The Ruby&#x27; Webrick server should not be used for serving public/production. It is designed to provide a simple, local development server.

If your goal is to always have Jekyll running (--watch) and watching for changes to source files in order to re-render the [updated] site, that can be accomplished, i.e. &#x60;jekyll serve -w&#x60;.

For automatically starting unix processes at boot time, that is a system administration detail and task unrelated to Jekyll or any other server-side technology.

With that said, one could use &#x60;cron&#x60;, or &#x60;init.d&#x60;, or now, &#x60;systemd&#x60;, or the myriad other runtime management solutions.

Another popular way to dynamically build sites when they change is to use &#x60;git&#x60; along with &#x60;git hooks&#x60; that triggers &#x60;jekyll&#x60; to build and output to the served web directory. Note GitHub does this for you.

---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/212#issuecomment-67528097) on: **Thursday, December 18, 2014**

If you&#x27;re planning to host the generated site for any prolonged period of time or to the public, I recommend serving the files with an actual web server like nginx or Apache. @jaybe-jekyll beat me to it, but I&#x27;ll just reiterate that &#x60;jekyll serve&#x60; was meant to be used for development, not production.

For your hosting needs, I also recommend [GitHub Pages](https://pages.github.com), which can build Jekyll sites without plugins.
