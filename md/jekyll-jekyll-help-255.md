# [Simple way to do case-insensitive sort?](https://github.com/jekyll/jekyll-help/issues/255)

> state: **open** opened by: **** on: ****

Hello-

I&#x27;m trying to figure out if there is a straightforward way to sort a list of tags in **case-insensitive** order.

Right now my tags post looks like this:

&#x60;&#x60;&#x60;
&lt;ul&gt;
  {% assign sorted_tags = site.tags | sort %}
  {% for tag in sorted_tags %}
    &lt;h4 id=&quot;{{tag[0]}}&quot; class=&quot;meta&quot;&gt;{{tag[0] | upcase}}&lt;/h4&gt;
    &lt;br&gt;
    {% for tagged_post in tag[1] %}
      &lt;a href=&quot;{{site.baseurl}}{{tagged_post.url}}&quot;&gt;{{tagged_post.title}}&lt;/a&gt;
      &lt;br&gt;
    {% endfor %}
    &lt;hr&gt;
  {% endfor %}
&lt;/ul&gt;
&#x60;&#x60;&#x60;

This returns the capitalized Tag names first, followed by lowercase names, in alphabetical order. I&#x27;d rather just upcase/downcase everything before sorting so the capitalization of the tag name doesn&#x27;t matter.

Is this a place to use the &#x60;map&#x60; filter? There is not very much documentation on how to use this on either the Jekyll or Shopify docs, so I&#x27;m sort of at a loss there.

Thanks for reading!

Edit: the repo in question is [here](https://github.com/gettypubs/beyond-the-printed-page) for reference.


### Comments

---
> from: [**tkrotoff**](https://github.com/jekyll/jekyll-help/issues/255#issuecomment-75932053) on: **Wednesday, February 25, 2015**

I have the same problem. I&#x27;ve opened an issue at https://github.com/Shopify/liquid/issues/529 asking for adding case-insensitive sort.
---
> from: [**tkrotoff**](https://github.com/jekyll/jekyll-help/issues/255#issuecomment-76960444) on: **Tuesday, March 03, 2015**

@egardner I&#x27;ve created a filter sort_casecmp that does what you want: https://github.com/tkrotoff/osteo15.com/commit/c7a25082b5d4c83c4d6c75db3ede3395c0277b80


