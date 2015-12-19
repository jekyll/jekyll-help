# [Lots of white space -- will it affect performance?](https://github.com/jekyll/jekyll-help/issues/269)

> state: **closed** opened by: **** on: ****

I have some code like this:

&#x60;&#x60;&#x60;
&lt;li&gt;&lt;a href=&quot;#&quot;&gt;Scenarios&lt;/a&gt;
&lt;ul&gt;
{% for page in site.pages | sort: &quot;weight&quot; %}
{% for tag in page.tags %}
{% if tag == &quot;Scenarios&quot; %}
&lt;li&gt;&lt;a class=&quot;post-link&quot; href=&quot;{{ page.url | prepend: site.baseurl }}&quot;&gt;{{ page.title }}&lt;/a&gt;&lt;/li&gt;
{% endif %}
{% endfor %}
{% endfor %}
&lt;/ul&gt;
&lt;/li&gt;
&#x60;&#x60;&#x60;

When I view the source of the generated HTML, there&#x27;s a TON of blank lines. Is that normal? Doesn&#x27;t that affect performance?

### Comments

---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/269#issuecomment-73797181) on: **Tuesday, February 10, 2015**

It is normal, since Liquid repeats newlines and whitespace in loops.

From what I know, the only downside to increased page size from whitespace is increased loading (downloading) time. From experience, this increase in time is extremely small (especially if gzip is enabled), and shouldn&#x27;t be an issue at all.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/269#issuecomment-73803092) on: **Tuesday, February 10, 2015**

I think @penibelst has done some work on solving this issue: https://github.com/penibelst/jekyll-compress-html
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/269#issuecomment-73804310) on: **Tuesday, February 10, 2015**

As @budparr mentioned you can remove whitespaces and other optional things by using my Compress Layout.

@alfredxing The download performance affect is small, but performance is more than downloading: gzipping itself, rendering, and caching, especially on small devices, depend on the *raw* file size.
---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/269#issuecomment-73842115) on: **Tuesday, February 10, 2015**

@penibelst True, thanks for the comment!
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/269#issuecomment-73845840) on: **Wednesday, February 11, 2015**

It&#x27;s worth mentioning the case of hosting on GitHub Pages: Content will be gzipped automatically.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/269#issuecomment-73980051) on: **Wednesday, February 11, 2015**

Thoroughly answered, gents. Thanks!
