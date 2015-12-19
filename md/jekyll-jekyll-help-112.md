# [&quot;Build Warning: Layout &#x27;nil&#x27; requested in sitemap.xml does not exist.&quot;](https://github.com/jekyll/jekyll-help/issues/112)

> state: **closed** opened by: **** on: ****

Hello! I should warn you all, I&#x27;m still pretty new to everything Jekyll/GitHub so if I need some obvious help just let me know.

I&#x27;m trying to start a blog with Jekyll, and I found a theme that I like (https://github.com/hmfaysal/hmfaysal-omega-theme). I forked the repo, cloned it to desktop, and without modifying anything, ran &#x60;&#x60;&#x60;jekyll serve --watch&#x60;&#x60;&#x60;. I&#x27;ve done a few simple Jekyll tutorials and I can at least get my web page to load most of the time on localhost:4000. But this one gave me the error message in the title. And for what it&#x27;s worth, there&#x27;s definitely stuff in sitemap.xml so I don&#x27;t know what it&#x27;s talking about. Here&#x27;s my fork (https://github.com/dwana49/hmfaysal-omega-theme). Can anyone help me? :confused: 

Dan

### Comments

---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/112#issuecomment-51134267) on: **Monday, August 04, 2014**

This warning was introduced recently in jekyll/jekyll#2620. Officially, in YAML, &#x60;layout: null&#x60; is the proper way to specify no layout, but there&#x27;s [an issue](https://github.com/jekyll/jekyll/pull/2652) open discussing usage of &#x60;nil&#x60;. For now, I&#x27;d recommend changing &#x60;layout: nil&#x60; at the the top of your &#x60;sitemap.xml&#x60; file to &#x60;layout: null&#x60;.
---
> from: [**dwana49**](https://github.com/jekyll/jekyll-help/issues/112#issuecomment-51137504) on: **Monday, August 04, 2014**

I think that did the trick. Thanks for your help!
