# [Is the option relative_permalinks deprecated?](https://github.com/jekyll/jekyll-help/issues/26)

> state: **closed** opened by: **** on: ****

If I set the variable &#x60;page.permalink&#x60; to a relative permalink, I get a deprecation message.

&gt; Deprecation: Starting in 2.0, permalinks for pages in subfolders must be relative to the site source directory, not the parent directory. Check http://jekyllrb.com/docs/upgrading/ for more info.

If I explicitly set 

&#x60;&#x60;&#x60;yaml
relative_permalinks: true
&#x60;&#x60;&#x60;
I still see the deprecation message. Does it mean, that &#x60;site.relative_permalinks&#x60; will completely disappear in 2.0.0?

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/26#issuecomment-41628339) on: **Monday, April 28, 2014**

It will default to &#x60;false&#x60; starting in 2.0. Not sure about complete removal. It was a relic of the pre-1.0 days that made its way into 1.0. Ref: https://github.com/jekyll/jekyll/pull/1081

/cc @benbalter
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/26#issuecomment-41650824) on: **Tuesday, April 29, 2014**

Thank you.
