# [Accessing un-published posts in a template?](https://github.com/jekyll/jekyll-help/issues/15)

> state: **closed** opened by: **** on: ****

Hello,

New to Jekyll, but recently ported my blog to Octopress. And sorry if this should be posted over at Octopress instead, but it felt Jekyll related.

I&#x27;m working on a &#x60;series&#x60; about porting a blog from BlogSpot/Blogger to Octopress and found [this post](http://realjenius.com/2012/11/03/jekyll-series-list/) about creating a series template to live in &#x60;_includes&#x60;.

I&#x27;ve since [modified the template](https://gist.github.com/staxmanade/10601886) and want access to un-published posts within the template. (or if this can be done through drafts - maybe I just need a tip on how to get them).

My end goal would be to get the following template to work.

[Series Gist Template](https://gist.github.com/staxmanade/10601886)

&gt; Notice how I&#x27;m trying to use &#x60;post.published == false&#x60; to output &#x60;(coming soon)&#x60;...

This works when I use &#x60;rake generate; rake preview&#x60; but the (coming soon) posts don&#x27;t show up immediately. There&#x27;s some sort of delay before this gets generated, and I presume this is related to the &#x60;preview&#x60; rake task in Octopress (but haven&#x27;t dug in enough). After some time if I refresh it will display what I want, but I&#x27;m pretty sure &#x60;rake generate; rake deploy&#x60; won&#x27;t produce the same thing.

### Is it possible to access un-published posts or possibly drafts - to accomplish something like this?

--------

This article is **Part 2** of **5** in a series about **something awesome.**

- Part 1 - something awesome - [Introduction](http://example.com)
- **Part 2 - something awesome - (current) Step 1**
- Part 3 - something awesome - [Step 2](http://example.com)
- Part 4 - something awesome - (coming soon) Step 3
- Part 5 - something awesome - (coming soon) Wrapping up


Over on Jekyll/Jekyll#2226 @parkr

&gt; Presently, only posts which are published are added to the site.posts array, so this isn&#x27;t possible.

Was hoping anyone with experience could light the way to an alternative approach to accomplish the desires output.


### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/15#issuecomment-40328899) on: **Sunday, April 13, 2014**

Consider using a different YAML front-matter item:

## In your post:

&#x60;&#x60;&#x60;yaml
---
title: My upcoming post
written: false
permalink: THIS-CRAZY-UNGUESSABLE-LINK-HAHAHAHOIAENFOEA
---

This post is unwritten
&#x60;&#x60;&#x60;

## Now list the posts, but showing only those you want to show:

&#x60;&#x60;&#x60;html
&lt;ul&gt;
{% for post in site.posts %}
  &lt;li&gt;
  {% if post.written %}
    &lt;a href=&quot;{{ post.url }}&quot;&gt;{{ post.title }}&lt;/a&gt;
  {% else %}
    {{ post.title }} (coming soon)
  {% endif %}
  &lt;/li&gt;
{% endfor %}
&lt;/ul&gt;
&#x60;&#x60;&#x60;
---
> from: [**staxmanade**](https://github.com/jekyll/jekyll-help/issues/15#issuecomment-40330463) on: **Sunday, April 13, 2014**

Thanks for the info - that could work at the post level, but I have the following issues with it.

1. The series component allows me to just include &#x60;{% include series.html %}&#x60; anywhere and it generates the series
2. I would rather not duplicate code like this within blog-post content if possible. Would love to find a way to abstract out the &#x60;series&#x60; concept from the posts...
3. I&#x27;m not sure how to get what I want with that example and not have the future posts arrive in the &#x60;rss&#x60; feed and the RecentPosts section on the right of the page.


---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/15#issuecomment-40330501) on: **Sunday, April 13, 2014**

You&#x27;ll like collections: https://github.com/jekyll/jekyll/pull/2199
---
> from: [**staxmanade**](https://github.com/jekyll/jekyll-help/issues/15#issuecomment-40332841) on: **Sunday, April 13, 2014**

Thanks, I&#x27;ll be patient and wait to take a look at this.

For now I may just go with what I have and not worry about showing the up-and-coming posts. Since the tool will automatically re-generate the series section anyway.

Thanks!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/15#issuecomment-41727603) on: **Tuesday, April 29, 2014**

Closing this out. I think this is a fantastic use case for collections, though. We should write a tutorial on how to do this once 2.0 is finally released.
