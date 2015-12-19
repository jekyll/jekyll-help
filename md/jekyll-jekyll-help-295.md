# [Sorting yaml front matter variable with heterogeneous data](https://github.com/jekyll/jekyll-help/issues/295)

> state: **open** opened by: **** on: ****

I am creating a simple song site with Jekyll and my songs are represented as posts with the following front matter:
&#x60;&#x60;&#x60;
---
layout: post
title: &quot;My Song&quot;
date: 2015-05-05T13:40:02-05:00
song: 3
---
&#x60;&#x60;&#x60;
I am porting over old paper bound books to this site, and unfortunately the songs are ordered by both number and letters (The song above could be 3 or it could be &quot;A&quot;). My attempts to sorting them are failing.

When doing this:
&#x60;&#x60;&#x60;{% assign posts = site.posts | sort:&#x27;song&#x27; %}&#x60;&#x60;&#x60;

It throws this error:
&#x60;&#x60;&#x60;Error: comparison of Jekyll::Post with Jekyll::Post failed&#x60;&#x60;&#x60;

Any way I can work around this?

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/295#issuecomment-99210160) on: **Tuesday, May 05, 2015**

Try putting quote marks around the song value, like &#x60;&#x60;&#x60;song: &#x27;3&#x27;&#x60;&#x60;&#x60;
---
> from: [**DonnieWest**](https://github.com/jekyll/jekyll-help/issues/295#issuecomment-99278658) on: **Tuesday, May 05, 2015**

Duh. Thank you!

On Tue, May 5, 2015, 3:25 PM Bud Parr &lt;notifications@github.com&gt; wrote:

&gt; Try putting quote marks around the song value, like song: &#x27;3&#x27;
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub.
&gt;

