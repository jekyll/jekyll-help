# [Access front matter with where filter](https://github.com/jekyll/jekyll-help/issues/95)

> state: **closed** opened by: **** on: ****

I want to list all posts with a certain value for a front matter key. Does the where filter allow the use of custom front matter key value pairs? If so, how do I get this to work?

&#x60;&#x60;&#x60;html
{% for post in site.posts | where:&quot;lang&quot;,&quot;de&quot; %}
     {% capture post_count %} {{ post_count | plus: 1 }} {% endcapture %}
{% endfor %}

&lt;!-- Currently, &#x60;post_count&#x60; matches all posts --&gt;
{{ post_count }}
&#x60;&#x60;&#x60;

### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/95#issuecomment-48808080) on: **Saturday, July 12, 2014**

Filters don’t work on &#x60;for&#x60; loops. Try this:

&#x60;&#x60;&#x60;
{{ site.posts | where:&quot;lang&quot;,&quot;de&quot; | size }}
&#x60;&#x60;&#x60;
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/95#issuecomment-48808144) on: **Saturday, July 12, 2014**

To get all de-posts, use &#x60;assign&#x60;

&#x60;&#x60;&#x60;
{% assign posts_de = site.posts | where:&quot;lang&quot;,&quot;de&quot; %}
&#x60;&#x60;&#x60;
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/95#issuecomment-48819271) on: **Saturday, July 12, 2014**

Thank you,

Changing the way my blog pages work:

&#x60;&#x60;&#x60;
{% assign posts_en = site.posts | where:&quot;lang&quot;,&quot;en&quot; %}
&lt;ul&gt;
{% for post in posts_en limit:10 %}
     &lt;li&gt;Post&lt;/li&gt;
{% endfor %}
&lt;/ul&gt;

{% if posts_en.size &gt; 10 %}
&lt;a href=&quot;/en/archive&quot; class=&quot;island&quot;&gt;Read more posts in the archive&lt;/a&gt;
{% endif %}
&#x60;&#x60;&#x60;
