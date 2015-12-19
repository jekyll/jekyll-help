# [Copy files verbatim to destination, without processing](https://github.com/jekyll/jekyll-help/issues/214)

> state: **open** opened by: **** on: ****

There should be an option that tells Jekyll to copy files verbatim, without processing. At present, to my knowledge, there is no way that one can host a markdown file verbatim. This [stackoverflow post](http://stackoverflow.com/questions/20509912/how-to-post-an-md-file-on-jekyll-website) suggests that one could host markdown files that don&#x27;t have YAML front-matter, but that doesn&#x27;t help when I want to host files that actually do contain such front-matter.

A simple configuration option like this:
&#x60;verbatim: [DIR, FILE, ...]&#x60;
modeled after &#x60;exclude&#x60; and &#x60;include&#x60;, would do the trick.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/214#issuecomment-68661317) on: **Sunday, January 04, 2015**

I haven&#x27;t tested this, so I don&#x27;t know if this would actually work, but try using a &#x60;txt&#x60; file extension instead of &#x60;md&#x60; or &#x60;markdown&#x60;. I believe Jekyll converts files based on their file extension, so if you give it a &#x60;txt&#x60;, it won&#x27;t have a converter for it and just dump it as it sees it.

Let me know if that works, please.
---
> from: [**clauswilke**](https://github.com/jekyll/jekyll-help/issues/214#issuecomment-68984935) on: **Tuesday, January 06, 2015**

File extensions don&#x27;t matter, as far as I can tell. I have files ending in .Rmd, which should be ignored, but they are not. 
