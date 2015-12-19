# [Get excerpt from specific element](https://github.com/jekyll/jekyll-help/issues/145)

> state: **closed** opened by: **** on: ****

Howdy peeps!

I&#x27;d like to ask if it is possible to get an excerpt from a specific element?

For example, let&#x27;s say I have a blog post and in it is:

&#x60;&#x60;&#x60;
&lt;div id=&quot;excerpt&quot;&gt;
  &lt;!-- Excerpt --&gt;
&lt;/div&gt;
&#x60;&#x60;&#x60;

I then want the content, HTML and all, within said &#x60;&lt;div&gt;&#x60; to be used as the excerpt in post list.

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/145#issuecomment-54933634) on: **Tuesday, September 09, 2014**

That is not possible at the moment. However, you can set &#x60;excerpt&#x60; in the front matter and put the content there.

&#x60;&#x60;&#x60;yaml
---
excerpt: &quot;Whatever you want it to be…&quot;
---
&#x60;&#x60;&#x60;
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/145#issuecomment-54972113) on: **Tuesday, September 09, 2014**

I need the HTML within the &#x60;&lt;div&gt;&#x60; to be parsed as part of the excerpt as well.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/145#issuecomment-54973157) on: **Tuesday, September 09, 2014**

&#x60;&#x60;&#x60;yaml
---
excerpt: &quot;&lt;div class=&#x27;note&#x27;&gt;&lt;p&gt;Whatever you want it to be…&lt;/p&gt;&lt;/div&gt;&quot;
---
&#x60;&#x60;&#x60;

…should do the trick. Be careful with &#x60;&quot;&#x60; and &#x60;&#x27;&#x60;.
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/145#issuecomment-54974746) on: **Tuesday, September 09, 2014**

Yes I could do that but sometimes the HTML gets quite complicated and it creates a mess at the beginning at the document.

It doesn&#x27;t really matter too much. I was hoping that the functionality would exist but it isn&#x27;t crucial that from me. I&#x27;ll find a different way around it.

Thanks for the help @kleinfreund!
