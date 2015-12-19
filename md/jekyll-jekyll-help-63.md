# [Extra Empty Nav Elements being created?](https://github.com/jekyll/jekyll-help/issues/63)

> state: **closed** opened by: **** on: ****

Hello,

I&#x27;m developing my first Jekyll site, a design portfolio. I&#x27;ve got the basic setup going, by following the installation guide on the Jekyll site, and I&#x27;ve made a few sample pages duplicating the about.md example, and posts duplicating the example in the _post directory.

It&#x27;s great that it automatically adds the nav when a new markdown page is created, but one thing I&#x27;ve noticed is that there&#x27;s an extra nav element being created. It appears it goes to the feed.xml page when collapsed to the hamburger menu, and one is completely empty when fully expanded. I can&#x27;t figure out what or why these extra spaces are being added into the main nav. 

Any ideas as to what could this be and how can I avoid it? 

Thanks in advance for any help,
Jason


### Comments

---
> from: [**jbishop228**](https://github.com/jekyll/jekyll-help/issues/63#issuecomment-44891822) on: **Monday, June 02, 2014**

Here&#x27;s my Jekyll directory structure:

_archive
_config.yml
_includes
_layouts
_plugins
_posts
_site
.git
.gitignore
about.md
contact.md
css
index.html
portfolio.md
README.md
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/63#issuecomment-45052531) on: **Tuesday, June 03, 2014**

&gt; I&#x27;ve noticed is that there&#x27;s an extra nav element being created

The page with the extra nav contains a layout name:

&#x60;&#x60;&#x60;html
---
layout: foo
---
&lt;!-- content --&gt;
&#x60;&#x60;&#x60;

Go to &#x60;_layouts/foo.html&#x60;, the nav must be there. If not, look in the parent layout. It’s the layout, that may be referenced in the &#x60;foo.html&#x60;:

&#x60;&#x60;&#x60;html
---
layout: bar
---
&lt;!-- content --&gt;
&#x60;&#x60;&#x60;

---
> from: [**RussEby**](https://github.com/jekyll/jekyll-help/issues/63#issuecomment-46096173) on: **Saturday, June 14, 2014**

I noticed the same thing. I looked at the default site that Jekyll creates when you just type in new. The extra elements were there as well.

What I&#x27;m doing:
I&#x27;m using the &#x27;for page in site.pages&#x27; to create a li element for the nav section.  

What it&#x27;s doing:
It&#x27;s creating a li element for every file including the files without a page.title, thus causing an empty li element.

How I solved it:
I did a test and it only made the li element if there was a page.title.

    {% for page in site.pages %}
    {% if page.title %}
    &lt;li&gt;
	 &lt;a href=&quot;{{ page.url | prepend: site.baseurl }}&quot;&gt;{{ page.title }}&lt;/a&gt;
    &lt;/li&gt;
    {% endif %}
    {% endfor %}

Hope this helps.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/63#issuecomment-48134920) on: **Sunday, July 06, 2014**

@RussEby :metal: 

![That&#x27;s a bingo!](http://i.imgur.com/fypAmty.gif)

You could also add something into the front-matter of the files that you want to show up in the nav, like &#x60;nav: yes&#x60; or something like that. Be aware of YAML boolean values, though - they can be a little tricky. Check out https://github.com/jekyll/jekyll/issues/2508#issuecomment-45940721 for some more detail.
