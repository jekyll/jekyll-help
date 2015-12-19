# [Front Matter defaults: Use categories in scope?](https://github.com/jekyll/jekyll-help/issues/45)

> state: **closed** opened by: **** on: ****

Is it possible to use categories as scope for Front Matter defaults?

&#x60;&#x60;&#x60;
  # all files with the category &#x60;portfolio&#x60; have the layout &#x60;post--portfolio&#x60;
  -
    scope:
      path: &quot;&quot;
      category: &quot;portfolio&quot; # Doesn&#x27;t work.
    values:
      layout: &quot;post--portfolio&quot;
&#x60;&#x60;&#x60;

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/45#issuecomment-43221789) on: **Thursday, May 15, 2014**

Actually, using &#x60;category: &quot;portfolio&quot;&#x60; works just as expected.

After going through the posts and layouts I was testing it with, I somehow figured it out. Can&#x27;t reproduce what I did wrong in the first place.
