# [Problem getting the &quot;truncate&quot; filter to work](https://github.com/jekyll/jekyll-help/issues/72)

> state: **closed** opened by: **** on: ****

jekyll (2.0.3), liquid (2.5.5)

I&#x27;m having a problem with the Liquid filter [&quot;truncate&quot;](https://github.com/Shopify/liquid/wiki/Liquid-for-Designers#standard-filters)
&#x60;&#x60;&#x60;
{% if item.front_matter_key %}
  {{ item.front_matter_key  | truncate: 1 }}
{% endif %}
&#x60;&#x60;&#x60;
For me, this returns nothing but ellipses (&#x27;...&#x27;), when if the front matter value was &#x60;&#x60;&#x60;foo&#x60;&#x60;&#x60; I&#x27;d expect the filter to return &#x60;&#x60;&#x60;f&#x60;&#x60;&#x60;

I&#x27;ve tried using capture and adding other filters to no avail. I&#x27;m assuming that anything on the Liquid for Designers page should work for us in Jekyll.

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/72#issuecomment-46127836) on: **Sunday, June 15, 2014**

Figured this out. Liquid apparently, obviously, adds the ellipses as a result of the truncate filter, but what I didn&#x27;t understand is that to get the desired result (of the value truncated to one character) you have to include the ellipses, so &quot;truncate: 4&quot; works, whereas whereas &quot;truncate: 1&quot; does not.

So, this seems strictly a liquid issue, which I will take up over there.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/72#issuecomment-46128460) on: **Sunday, June 15, 2014**

Issue [here](https://github.com/Shopify/liquid/issues/366) if anyone is intrested.
