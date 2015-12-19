# [Page build failure](https://github.com/jekyll/jekyll-help/issues/185)

> state: **closed** opened by: **** on: ****

I&#x27;m sorry if this is a duplicate but I&#x27;m new to SCSS and keep getting emailed with this when trying to commit any changes:

&gt; The page build failed with the following error:
&gt; The file &#x60;blog/css/main.scss&#x60; contains syntax errors. For more information, see 
&gt; https://help.github.com/articles/page-build-failed-markdown-errors.
&gt; If you have any questions please contact us at https://github.com/contact.

My compiler is telling me the following:
&gt; Invalid CSS after &quot;-&quot;: expected number or function, was &quot;--&quot;

### Comments

---
> from: [**tjmapes**](https://github.com/jekyll/jekyll-help/issues/185#issuecomment-61747526) on: **Tuesday, November 04, 2014**

Note, I&#x27;m brand new to Jekyll and am just trying to wrap my head around it.  A bit obsessed ATM and really want to get things working smoothly to learn as much as possible about Jekyll.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/185#issuecomment-61781862) on: **Wednesday, November 05, 2014**

Would be helpful to see that SCSS. Can you link to it in that repo?
---
> from: [**tjmapes**](https://github.com/jekyll/jekyll-help/issues/185#issuecomment-61817650) on: **Wednesday, November 05, 2014**

Sure thing @kleinfreund thank you.
https://github.com/tjmapes/tjmapes.github.io/blob/946691f7460831727975baf829e4ebe379039e5b/blog/css/main.scss
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/185#issuecomment-61818980) on: **Wednesday, November 05, 2014**

I&#x27;m not entirely sure about this, but I think you (now) need to add &#x60;&#x60;&#x60;layout: null&#x60;&#x60;&#x60; in your front matter

so, instead of 
&#x60;&#x60;&#x60;
---
# Only the main Sass file needs front matter (the dashes are enough)
---
&#x60;&#x60;&#x60;
try this:
&#x60;&#x60;&#x60;
---
layout: null
# Only the main Sass file needs front matter (the dashes are enough)
---
&#x60;&#x60;&#x60;
---
> from: [**tjmapes**](https://github.com/jekyll/jekyll-help/issues/185#issuecomment-61821189) on: **Wednesday, November 05, 2014**

Thanks @budparr.
