# [Linked tag and linked category pages](https://github.com/jekyll/jekyll-help/issues/179)

> state: **closed** opened by: **** on: ****

I&#x27;m trying to transition from a simple WordPress install to Jekyll. I get how to do the individual posts, and I sorta have the landing index.html page, but so far I haven&#x27;t found any working methods (for me) to generate a page with linked tags, or a page with linked categories.

My goal is to have these pages generate automatically, one with authors and one with topics (this is a site for interesting quotes), so that if you want to see all quotes by Thomas Jefferson, you can go to authors.html and click on Jefferson. If you want to see all quotes that I&#x27;ve categorized as dealing with taxes (for example), you go to categories.html and click on Taxes.

I find myself wishing that the documentation had a simple Jekyll site available for download, so neophytes such as myself can grasp what goes where.

Any help is appreciated. Thanks in advance.

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/179#issuecomment-60771572) on: **Tuesday, October 28, 2014**

The source for jekyllrb.com is available within the ./site/ folder of the Jekyll source repository. It has examples of categorization (News) and parsing.

To *automatically* build category and tag pages, etc. you could leverage a plugin to perform the task.

To [better] understand how templates, Jekyll, and Liquid work together... you could fairly easily create your own simple template that would show all posts within a particular category/categories:

&#x60;&#x60;&#x60; html+jinja
&lt;ul&gt;
  {% for post in site.categories.taxes %}
    &lt;li&gt;
      &lt;a href=&quot;{{ post.url }}&quot;&gt;{{ post.title }}&lt;/a&gt;
    &lt;/li&gt;
  {% endfor%}
&lt;/ul&gt;
&#x60;&#x60;&#x60;

There are other ways to accomplish this pragmatically as well, but that is a simple, straight-forward starter to understanding the logic.

https://github.com/shopify/liquid/wiki/liquid-for-designers

https://github.com/jekyll/jekyll/tree/master/site

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/179#issuecomment-60789002) on: **Tuesday, October 28, 2014**

Please note that pages like for category _xyz_ are not automatically available like you know from Wordpress. Also having pagination for filtered content (categories, tags, ...) is not possible without plugins.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/179#issuecomment-60789870) on: **Tuesday, October 28, 2014**

Good points, @kleinfreund.

I would like to see pagination available more than once and amongst categories, tags, collections, etc. @parkr, comment on such support within 3x roadmap?

Additional idea/thought-- native support in Jekyll to optionally specify providing a category/tag index. i.e.:

&gt; _config.yml or _data/example.yml or _collection/...

category:
  \- taxes:
    pagination: true
  \- llamas:
    pagination: true  # default: [undefined]  
    layout: fancyindex

Results in category index being created for URI /taxes/[index.html] and /llamas/[index.html]. # Note actual paths would be &#x60;permalink: &#x60;-driven.

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/179#issuecomment-60799757) on: **Tuesday, October 28, 2014**

&gt; I would like to see pagination available more than once and amongst categories, tags, collections, etc. @parkr, comment on such support within 3x roadmap?

I&#x27;m not able to devote time to it. We could add it tomorrow if we had people who wanted to work on the code here: https://github.com/jekyll/jekyll-paginate Just need pull requests and excitement.
