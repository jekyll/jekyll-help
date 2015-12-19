# [Frontmatter in page should override Frontmatter in layout](https://github.com/jekyll/jekyll-help/issues/59)

> state: **closed** opened by: **** on: ****

For a layout abc with:
&#x60;&#x60;&#x60;
---
mykey: &quot;myvalue&quot;
---
My Key = {{mykey}}
&#x60;&#x60;&#x60;

and a page with
&#x60;&#x60;&#x60;
layout: abc
mykey: &quot;bettervalue&quot;
&#x60;&#x60;&#x60;

currently myvalue is displayed. This really should be bettervalue to allow easy overriding of defaults.

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/59#issuecomment-44751491) on: **Saturday, May 31, 2014**

An approach to accomplishing the desired logic is as follows:

&#x60;./a-page.md&#x60;

&#x60;&#x60;&#x60;
---
layout: layout-a
mykey: &quot;my custom non-default value&quot;
---

...
&#x60;&#x60;&#x60;

&#x60;./_layouts/layout-a.html&#x60;

&#x60;&#x60;&#x60;
---
layout: default
---

&lt;p&gt;
  mykey:  {% if page.mykey %}{{ page.mykey }}{% else %}MY *DEFAULT* MYKEY!{% endif %}
  mkey again, another way:  {% if page.mykey %}{{ page.mykey }}{% else %}{{ site.default.mykey }}{% endif %}
&lt;/p&gt;

&#x60;&#x60;&#x60;

&gt; Logic translation:

If page.mkey is defined/exists via a page, then use it, else use either my static declared default, and or secondly, as above, a global variable.

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/59#issuecomment-44755310) on: **Saturday, May 31, 2014**

Yes, use &#x60;page.mykey&#x60;. We don&#x27;t define any top-level Liquid vars other than &#x60;page&#x60;, and &#x60;site&#x60;.
---
> from: [**axelfontaine**](https://github.com/jekyll/jekyll-help/issues/59#issuecomment-44823355) on: **Monday, June 02, 2014**

Thanks for the suggestion. With the proposed syntax, I think we can all agree this qualifies as a workaround at best.

I hope you&#x27;ll consider this change of overriding precedence for a future version. After all the whole point of a template is the be a base to be customized by the individual pages building upon it.

Cheers and thanks for the fantastic product,
Axel
