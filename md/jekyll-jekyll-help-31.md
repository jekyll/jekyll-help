# [Full story pages: How to skip posts that are of a certain category?](https://github.com/jekyll/jekyll-help/issues/31)

> state: **closed** opened by: **** on: ****

Hi,

### TL;DR:

On my full story pages (permalink pages), I&#x27;d like to setup the previous/next pagination so that it skips any posts that have a category of X.

### Long version:

Specifically, I have some posts that are categorized as &quot;external&quot;, and when these posts are clicked, the end user is taken to that external link.

Using the previous/next post pagination, I&#x27;d like to ignore all external posts and just paginate between actual posts.

I build my Jekyll site locally, so plugins aren&#x27;t an issue, but I was mostly curious if there&#x27;s built-in way to &quot;filter&quot; the pagination object based on post frontmatter (like category)?

Does any of that make sense?

Again, this is the pagination for the post pages … The permalink pages. Not the index page pagination.

For now, I don&#x27;t use the prev/next pagination on permalink pages because I couldn&#x27;t figure out how to ignore my &quot;external&quot; posts.

Thanks!

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42261367) on: **Monday, May 05, 2014**

Hi @mhulse, you could do something like:

&#x60;&#x60;&#x60;html
{% for post in site.posts %}
    {% unless post.categories contains &#x27;external&#x27; %}
        &lt;li&gt;&lt;a href=&quot;{{ post.url }}&quot;&gt;{{ post.title }}&lt;/a&gt;&lt;/li&gt;
    {% endif %}
{% endfor %}
&#x60;&#x60;&#x60;

I don&#x27;t know if this helps
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42261866) on: **Monday, May 05, 2014**

Ooh, unless is available to use now? Nice!!! I&#x27;ll give that a whirl and let you know how it goes. Thanks much @albertogg, I really appreciate it!
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42262015) on: **Monday, May 05, 2014**

Wait, &#x60;unless&#x60; has been available forever. What was I thinking. Lol. I must be thinking of some other filter. Anyway, I&#x27;ll be back shortly with results. 
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42262051) on: **Monday, May 05, 2014**

No problem!
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42266142) on: **Monday, May 05, 2014**

So, I think the problem is that I need to grab the next non-external post in order to make the permalink-page to permalink-page navigation smooth.

Let me clarify my question:

---

On my Jekyll permalink pages (the post detail view) I have previous and next buttons:

&#x60;&#x60;&#x60;html
{% if page.previous.url %} ... link to previous post ... {% endif %}
{% if page.next.url %} ... link to next post ... {% endif %}
&#x60;&#x60;&#x60;

That works great.

### My problem:

How do I skip a next/previous post if it&#x27;s of a certain type?

Specifically, I have some posts that are &quot;external&quot; links. They have a font matter value of &quot;layout: external&quot; (and a category of &quot;external&quot;).

### Goal:

I&#x27;d like to skip next/previous posts that are marked as external; in other words, if the next two posts were marked as &quot;external&quot;, my &quot;next&quot; button would skip to the third post.

I can&#x27;t think of an easy way to do this using liquid.

---

Maybe something like this (pseudo logic):

1. On permalink page …
1. … Get all posts as an object/array …
1. … Filter out posts marked as &quot;external&quot; (maybe this can be done in step above) …
1. … In object/array, find the current post …
1. … Assign &quot;prev&quot; and &quot;next&quot; variables to previous and next post in object/array (i.e. the next and previous post from current post position).

I hope that makes more sense.

---

@albertogg …

&gt; &#x60;&#x60;&#x60;
{% for post in site.posts %}
    {% unless post.categories contains &#x27;external&#x27; %}
        &lt;li&gt;&lt;a href=&quot;{{ post.url }}&quot;&gt;{{ post.title }}&lt;/a&gt;&lt;/li&gt;
    {% endif %}
{% endfor %}
&#x60;&#x60;&#x60;

After having thought about this, the above will spit out all posts not marked &quot;external&quot;. While that&#x27;s a good technique to filter posts, I don&#x27;t think it&#x27;s going to work like I need it to on the permalink pagination. :frowning: 


---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42268277) on: **Monday, May 05, 2014**

I still don&#x27;t understand quite well what you want to do. Do you have any existing code/repo I can check?
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42324784) on: **Tuesday, May 06, 2014**

Hi @albertogg, thanks again for the help! :+1: 

I&#x27;m not at a machine where I can slap together demo code, but when I am, I&#x27;ll put together a simple demo to show you the problem I&#x27;m trying to solve.

The most simple way I can describe it:

On a post detail page (aka, permalink page) I have the Jekyll tags to get previous and next posts:

&#x60;&#x60;&#x60;html
{% if page.previous.url %} ... link to previous post ... {% endif %}
{% if page.next.url %} ... link to next post ... {% endif %}
&#x60;&#x60;&#x60;
When clicking on the above links, I want to ignore next or previous pages that are &quot;external&quot; (i.e., isn&#x27;t meant to have a permalink page as it takes the user to the external website).

Optimally, I&#x27;d like to have my next/previous page navigation to skip &quot;external&quot; links and just page through pages/posts that are not marked external.

I hope that helps to explain things better, though, I&#x27;m sure an example will make things clearer. :laughing: 

With that said, this is kinda like what I want:

http://ajclarkson.co.uk/blog/jekyll-category-post-navigation/
https://gist.github.com/stravid/4078840

Except, I want to ignore a category (or, custom frontmatter key/value).

I probably should spend some more time coding up a plugin … I started this code:

https://github.com/mhulse/mhulse.github.io/issues/31#issuecomment-35181904

… but I don&#x27;t remember if I ever got it working. I kinda gave up on the issue.

Since this repo was created, I thought I&#x27;d try again; I&#x27;m mostly curious if there&#x27;s a way to do this using Liquid only. With Jekyll 2.0 coming out, maybe there&#x27;s some new code that could make this happen using only template tags?

Anyway, I&#x27;ll setup a demo site with minimal code as an example.

Thanks for listening to me ramble. :octocat: 
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42393737) on: **Tuesday, May 06, 2014**

Oooh, the 2.0 &#x60;where&#x60; and &#x60;group_by&#x60; filters might be of some help here:

http://jekyllrb.com/docs/templates/#filters
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42449002) on: **Wednesday, May 07, 2014**

Oh, I&#x27;m sorry I haven&#x27;t been very helpful... hope jekyll 2.0 will make it easier.
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42453147) on: **Wednesday, May 07, 2014**

&gt; Oh, I&#x27;m sorry I haven&#x27;t been very helpful... hope jekyll 2.0 will make it easier.

@albertogg Oh, no, you&#x27;ve been very helpful! I&#x27;m sorry that I have not been describing my problem very well.

I&#x27;m not sure if the &#x60;where&#x60; or &#x60;group_by&#x60; will really help; I have yet to do any experimentation with Jekyll 2.0 … There might be some new liquid bits that may be of some help … I&#x27;m really not sure though.

I plan on setting up a simple demo this weekend if you are still interested in seeing a working example. I think once you see a demo of my problem then all will be clear. :+1: 

Thanks for all of your help on this! Again, I apologize if I&#x27;ve been too wordy and/or not clear in my exact problem.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42454475) on: **Wednesday, May 07, 2014**

Yes, I&#x27;m interested on a demo to see if we can solve this. Don&#x27;t worry about it, I know it&#x27;s hard to explain and also english is not my first language.
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42454889) on: **Wednesday, May 07, 2014**

&gt; Yes, I&#x27;m interested on a demo to see if we can solve this.

Sounds good! I&#x27;ll have a demo up this weekend. Whether or not it&#x27;s solvable is one thing though … I may be trying to do something that&#x27;s just not possible. :laughing: 

&gt; Don&#x27;t worry about it, I know it&#x27;s hard to explain and also english is not my first language.

Wow, your English is very good! I would have never guessed. English is my only language, and as you can see, I&#x27;m not very good at it. Hehe.

I&#x27;ll post a link as soon as I have a demo up. Thanks again Alberto!
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42798898) on: **Sunday, May 11, 2014**

There is currently no way to filter the posts in &#x60;paginator.posts&#x60;. Thus, you only can build a pagination based on all of your posts, not just a couple of them.
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/31#issuecomment-42800325) on: **Sunday, May 11, 2014**

Ok, thanks for the help folks, I really appreciate it. I&#x27;m going to live without pagination on permalink pages. Also, I was afk all weekend, so I did not have a chance to create a demo. Oh well, looks like the demo would be moot anyway. Thanks again all!!!
