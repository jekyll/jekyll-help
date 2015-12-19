# [How can I include a transformed to HTML markdown file?](https://github.com/jekyll/jekyll-help/issues/173)

> state: **closed** opened by: **** on: ****

I use {% include something.md %} and jekyll does it literally without transforming md to html. How do I enable the transformation?

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/173#issuecomment-59679221) on: **Sunday, October 19, 2014**

You may capture and filter through Markdown.

{% capture something %]
  {% include something.md %}
{% endcapture %}

{{ something | markdownify }}

http://jekyllrb.com/docs/templates/#filters

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/173#issuecomment-61438727) on: **Sunday, November 02, 2014**

Usually, includes are already HTML (think of them as modules for the layout). However, @jaybe-jekyll&#x27;s solution should work as advertised.
