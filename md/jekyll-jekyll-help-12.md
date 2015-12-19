# [How do I create back/next pages like on jekyllrb.com?](https://github.com/jekyll/jekyll-help/issues/12)

> state: **closed** opened by: **** on: ****

I&#x27;d like to guide our readers from start to finish through our docs, so I need Back/Next links for **pages**, much like how on jekyllrb.com you start at [Home](http://jekyllrb.com/docs/home/) and you can &quot;Next&quot;-click all the way to [History](http://jekyllrb.com/docs/history/), where the page series ends.

All the guides I&#x27;ve found appear to be talking about pagination for posts, not *pages*. Could you give me some pointers?

As an aside, this personal preference is probably plugin territory, but I would love it if I could define something like this in front matter:

&#x60;&#x60;&#x60;
series: guide
order: 1
&#x60;&#x60;&#x60;
or
&#x60;guide-series: 1&#x60;

It would be in the same vein as using [custom taxonomies to achieve series](http://wordpress.org/plugins/series/) in WordPress.

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/12#issuecomment-39760325) on: **Monday, April 07, 2014**

I don&#x27;t really understand what you are trying to do, but have you looked into [jekyllrb section_nav](https://github.com/jekyll/jekyll/blob/gh-pages/_includes/section_nav.html) and [the front-matter of a page](https://github.com/jekyll/jekyll/blob/gh-pages/docs/installation.md)?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/12#issuecomment-39764826) on: **Monday, April 07, 2014**

@erlend-sh There are several ways to do this, and they&#x27;re all a bit involved. Check out @albertogg&#x27;s links for how it is done on [Jekyll&#x27;s site](http://jekyllrb.com) (for now, at least). I made a PR (jekyll/jekyll#1953) a while ago that uses Jekyll&#x27;s [data feature](http://jekyllrb.com/docs/datafiles/) to define the order for the documentation.
---
> from: [**afeld**](https://github.com/jekyll/jekyll-help/issues/12#issuecomment-57914785) on: **Saturday, October 04, 2014**

I&#x27;m looking for the same functionality. It&#x27;s a shame that, for the http://jekyllrb.com use case (at least), the order is essentially specified in three places:

* [In the data file](https://github.com/jekyll/jekyll/blob/592cee9596475c3fd85ce334d1832ee5339d8360/site/_data/docs.yml)
* As the &#x60;prev_section&#x60; in the page metadata
* As the &#x60;next_section&#x60; in the page metadata

Only using the data file would be nice because it has the ability to group the pages and you can see the order at a glance, but it would be messy to then create the [&#x60;section_nav&#x60;](https://github.com/jekyll/jekyll/blob/331247f96f1d885b1397c756dcd380caad76b5a7/_includes/section_nav.html) from the &#x60;site.data.docs&#x60;. [jekyll-page-collections](https://github.com/jeffkole/jekyll-page-collections) seems like it solves this problem, but of course, no Pages support :wink: 

/cc @jeffkole
---
> from: [**jeffkole**](https://github.com/jekyll/jekyll-help/issues/12#issuecomment-58043409) on: **Monday, October 06, 2014**

Have you looked into [Collections](http://jekyllrb.com/docs/collections/) at all? I can&#x27;t tell from the documentation if pagination is supported, but it might be a place to start.  When I was working on [jekyll-page-collections](https://github.com/jeffkole/jekyll-page-collections), pagination was hard-coded to be for posts only, but that was on Jekyll version 1.4.3.  It all may have changed since then.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/12#issuecomment-58171106) on: **Tuesday, October 07, 2014**

Collections currently don&#x27;t have pagination support afaik.
---
> from: [**mvaneijgen**](https://github.com/jekyll/jekyll-help/issues/12#issuecomment-66279770) on: **Tuesday, December 09, 2014**

Still not pagination for collections? 
---
> from: [**afeld**](https://github.com/jekyll/jekyll-help/issues/12#issuecomment-66304416) on: **Tuesday, December 09, 2014**

From @parkr&#x27;s [blog post](https://byparker.com/blog/2014/jekyll-3-the-road-ahead/) about the Jekyll 3 roadmap:

&gt; I’d like to focus on...Simplifying the internals and the abstractions (posts will become a collection)

Presumably this means that pagination would/could be easily rolled up to Collections.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/12#issuecomment-66364394) on: **Tuesday, December 09, 2014**

Realizing now this was originally posted quite some time ago, I have some Liquid markup that does what I think the original poster was looking for, that is, from a page in a collection automatically create a link to the previous and next page. A bit messy and hopefully this ability won&#x27;t get changed, but here it is: https://gist.github.com/budparr/3e637e575471401d01ec
