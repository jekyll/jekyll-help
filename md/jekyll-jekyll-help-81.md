# [is it possible to do post.image ? ](https://github.com/jekyll/jekyll-help/issues/81)

> state: **closed** opened by: **** on: ****

So I get how to add images to my content, however I would like to add an image in my index that is specific to every post. exactly like post.excerpt, but for images.

How would I go about doing this ?


Thanks,

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/81#issuecomment-46759059) on: **Saturday, June 21, 2014**

Add it to your front-matter and use &#x60;page.image&#x60;:

&#x60;&#x60;&#x60;html
---
title: A page
image: /path/to/my/image.png
---

&lt;img src=&quot;{{ page.image }}&quot; /&gt;
&#x60;&#x60;&#x60;
---
> from: [**HarrisRobin**](https://github.com/jekyll/jekyll-help/issues/81#issuecomment-46759069) on: **Saturday, June 21, 2014**

Ok that was way simpler than I thought. Thanks @parkr . Much appreciated. I now understand how to generate data i need in my posts.
