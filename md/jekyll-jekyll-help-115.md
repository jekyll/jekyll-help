# [Use &#x27;Where&#x27; filter with &#x27;or&#x27; ](https://github.com/jekyll/jekyll-help/issues/115)

> state: **closed** opened by: **** on: ****

is it possible to use the &#x27;where&#x27; filter with more than one comparison? with &#x27;or&#x27; or &#x27;and&#x27;?
I would like to do:

&#x60;{% assign mylist = tag[1] | sort | where: &quot;type&quot;, &quot;Videos&quot; or &quot;type&quot;,&quot;Articles&quot; %}&#x60;

But I can&#x27;t figure it out.

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/115#issuecomment-51705182) on: **Saturday, August 09, 2014**

This is not available. You&#x27;ll have to use the &#x60;where&#x60; filter to extract each set and merge the two resulting arrays. I can&#x27;t think of a good way to support this in our native &#x60;where&#x60; filter.
