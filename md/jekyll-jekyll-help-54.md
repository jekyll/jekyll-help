# [Example use of &#x60;where&#x60; and &#x60;group_by&#x60; filters?](https://github.com/jekyll/jekyll-help/issues/54)

> state: **closed** opened by: **** on: ****

I&#x27;ve been looking through the docs and testing out the &#x60;where&#x60; and &#x60;group_by&#x60; filters. They don&#x27;t seem to be returning the values that I want. Could someone show me an example of how to properly use these?

### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/54#issuecomment-45052191) on: **Tuesday, June 03, 2014**

&gt; They don&#x27;t seem to be returning the values that I want.

These filters are available since version 2.0.0. What version are you using? The [current version on Github Pages](https://pages.github.com/versions/) is older than 2.0.0.
---
> from: [**JesseHerrick**](https://github.com/jekyll/jekyll-help/issues/54#issuecomment-45052433) on: **Tuesday, June 03, 2014**

Version 2.0.3 @penibelst
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/54#issuecomment-45057760) on: **Wednesday, June 04, 2014**

&#x60;&#x60;&#x60;
{{ site.pages | where: &#x27;title&#x27;, &#x27;Home&#x27; }}
&#x60;&#x60;&#x60;

Shows the content of the page with &#x60;title&#x60; property set to “Home”.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/54#issuecomment-48135528) on: **Sunday, July 06, 2014**

@penibelst :clap: Perfect. I&#x27;m going to close this issue, since it&#x27;s pretty old, however...

@JesseHerrick If you&#x27;d like, feel free to post some code snippets with what you&#x27;re expecting to see and what you&#x27;re actually seeing; we&#x27;d be happy to help you troubleshoot!
