# [Confused by YAML Matter Defaults](https://github.com/jekyll/jekyll-help/issues/83)

> state: **closed** opened by: **** on: ****

I don&#x27;t understand how one is supposed to use YAML matter defaults. In the example given (reproduced below), the scope is listed as &#x60;path: &quot;about/blog&quot;&#x60; and &#x60;type: &quot;post&quot;&#x60;. But all posts are stored in one place &#x60;_posts&#x60;. So is &#x60;about/blog&#x60; the *destination* somehow? But how would a blog post end up in &#x60;about/blog&#x60; when they almost certainly end up in &#x60;:year/:month/:day/...&#x60; ?

&#x60;&#x60;&#x60;
defaults:
  -
    scope:
      path: &quot;&quot; # empty string for all files
    values:
      layout: &quot;my-site&quot;
  -
    scope:
      path: &quot;about/blog&quot;
      type: &quot;post&quot;
    values:
      layout: &quot;meta-blog&quot; # overrides previous default layout
      author: &quot;Dr. Hyde&quot;
&#x60;&#x60;&#x60;


### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/83#issuecomment-46888882) on: **Monday, June 23, 2014**

Posts can be located in other locations, such as within subdirectories.

For example:

&gt; &#x60;./about/blog/_posts/&#x60;

Another, likely-more-appropriate example:

&gt; &#x60;./technology/electronics/_posts/&#x60;

Defaults specified with respect to the examples would apply to each example.
---
> from: [**trans**](https://github.com/jekyll/jekyll-help/issues/83#issuecomment-46902777) on: **Monday, June 23, 2014**

I see. Okay thanks.
