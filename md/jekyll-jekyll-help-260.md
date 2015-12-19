# [Do any generators work with pages or collections?](https://github.com/jekyll/jekyll-help/issues/260)

> state: **open** opened by: **** on: ****

Are there any generators that work with pages or collections? I want to auto-generate a page that lists all pages that have a specific tag (like a tag archive but for pages instead of posts). Alternatively, I would like to auto-generate a collection archive (a page showing all pages in the collection). I know how to manually create these archive pages, but is there a way to auto-generate this kind of archive page on build? Ideally, it would work like this [categories archive plugin](https://github.com/recurser/jekyll-plugins) but for pages. 

### Comments

---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73415802) on: **Sunday, February 08, 2015**

https://github.com/jekyll/jekyll-archives
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73417980) on: **Sunday, February 08, 2015**

Thanks, but does this work for pages, or just for posts? When I implement it, it seems to only create archives for posts.
---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73418062) on: **Sunday, February 08, 2015**

Oh, you wanted for pages too? Not to sure that this is possible right now with Jekyll as pages and posts are two different things when generated, this might change in 3.0 but I&#x27;m not sure. 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73420111) on: **Sunday, February 08, 2015**

@tomjohnson1492 May I ask what kind of pages you have and why you want to why you want to put them together in an archive?
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73422523) on: **Sunday, February 08, 2015**

I&#x27;m using jekyll for a documentation site (e.g., user guide). There are various sections in the user guide, such as Getting Started, Implementation, Configuration, Troubleshooting. I want to create different lists for pages in these sections. So I&#x27;m using tags with pages. I&#x27;m autopopulating my navigation sidebar like this:

&#x60;&#x60;&#x60;
&lt;ul&gt;
  {% assign counter = &#x27;0&#x27; %}
                {% assign pages = site.pages | sort: &quot;weight&quot;  %}
                {% for page in pages %}
                {% for tag in page.tags %}
                {% if tag == &quot;Implementation&quot; and counter &lt; &#x27;9&#x27; %}
                {% capture counter %}{{ counter | plus:&#x27;1&#x27; }}{% endcapture %}
                  &lt;li class=&quot;list-group-item&quot;&gt;&lt;a href=&quot;{{ page.permalink | prepend: site.baseurl }}&quot;&gt;{{page.title}}&lt;/a&gt;&lt;/li&gt;
                {% endif %}
                {% endfor %}
                {% endfor %}
&lt;/ul&gt;
&#x60;&#x60;&#x60;

When there are more than 9 topics, I want to say something like &quot;See all&quot; and direct people to an archive view that contains all topics. 

Am I going about this the wrong way? Should I be using posts instead? I just don&#x27;t want to have to put dates in all my post file names, but maybe I can live with it if posts give me more functionality.
---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73422934) on: **Sunday, February 08, 2015**

Personally I think this would pretty easily be solved by having the pages be posts and have the permalink setting to be /:category/:name/ or something similar. Otherwise I think you might have to create separate pages manually and populate them yourself. You could possibly use a similar solution to how Jekyll organizes it&#x27;s sidenav in the documentation on their website.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73423350) on: **Sunday, February 08, 2015**

What soundr3 said sound5 logical. I did a very basic archive on my blog with my posts as well. Linking to the archive when there is 10 posts already.
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73423426) on: **Sunday, February 08, 2015**

How are pages and posts changing in v3 of Jekyll? I&#x27;m not sure I totally understand the difference between the two. What I know is that posts are meant for chronological sorting, they have dates in file names, and they occupy a different namespace (posts instead of pages). Pages are meant to stand alone, they don&#x27;t sort by filename, and they have a posts namespace. 

But it seems like more features are built around posts (like this plugin) since jekyll was initially conceived as a blogging platform, not a website/user-guide platform. Right?
---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73423629) on: **Sunday, February 08, 2015**

It&#x27;s mostly a backend change as far as I know and I haven&#x27;t really read much up on the specifics, but you can see the main issue here: jekyll/jekyll#3169. 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73423722) on: **Sunday, February 08, 2015**

@tomjohnson1492 Yeah, I see your point. As far as I understand it, pages are meant to not have any relation between one and another. That&#x27;s what collections _should_ be for.
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73424585) on: **Sunday, February 08, 2015**

I understand how collections work, but I&#x27;m looking to develop a simple
template that doesn&#x27;t require other writers to add collection types to a
configuration file. I think categories or tags on posts will work just fine
for me.

---------------------
blog: idratherbewriting.com
twitter: tomjohnson
408-540-8562

On Sun, Feb 8, 2015 at 10:22 AM, Philipp Rudloff &lt;notifications@github.com&gt;
wrote:

&gt; @tomjohnson1492 &lt;https://github.com/tomjohnson1492&gt; Yeah, I see your
&gt; point. As far as I understand it, pages are meant to not have any relation
&gt; between one and another. That&#x27;s what collections *should* be for.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub
&gt; &lt;https://github.com/jekyll/jekyll-help/issues/260#issuecomment-73423722&gt;.
&gt;

