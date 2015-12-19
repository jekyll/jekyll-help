# [Template has page.excerpt, not post.excerpt.](https://github.com/jekyll/jekyll-help/issues/180)

> state: **closed** opened by: **** on: ****

In my head-template I tried to use the following HTML:
&#x60;&#x60;&#x60;html
&lt;meta name=&quot;description&quot; content=&quot;{% if post.excerpt %}{{ post.excerpt | strip_html }}{% else %}{{ site.description }}{% endif %}&quot;&gt;
&#x60;&#x60;&#x60;

This always displayed the site description, because post.excerpt is empty on a post. Changing it to the following code worked:
&#x60;&#x60;&#x60;html
&lt;meta name=&quot;description&quot; content=&quot;{% if page.excerpt %}{{ page.excerpt | strip_html }}{% else %}{{ site.description }}{% endif %}&quot;&gt;
&#x60;&#x60;&#x60;

The excerpt is stored on the page, not the post. Is that intended behavior? If so, I&#x27;ll create a pull request to update the documentation.

I&#x27;m using Jekyll v2.4.0 on Windows 8.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/180#issuecomment-61438383) on: **Sunday, November 02, 2014**

&#x60;page&#x60; refers to the current page being displayed, &#x60;post&#x60; doesn&#x27;t really refer to anything, unless you&#x27;re using &#x60;post&#x60; as your iterating variable in a loop, like &#x60;{% for post in site.posts %}&#x60;.

I believe your change is correct; feel free to make a PR. Thanks! :smile: 
