# [How do I enable Redcarpet&#x27;s &quot;smart&quot; extension?](https://github.com/jekyll/jekyll-help/issues/56)

> state: **closed** opened by: **** on: ****

The [Redcarpet configuration](http://jekyllrb.com/docs/configuration/#redcarpet) documentation is a little unclear on how to specifically enable the &#x60;smart&#x60; extension (smart quotes, em dashes, etc.).

I found another site that used the following snippet, but it&#x27;s not working for me. I made sure to stop/start Jekyll and I&#x27;m using Jekyll 2.x.

&#x60;&#x60;&#x60;
markdown: redcarpet
redcarpet:
  extension: [&quot;smart&quot;]
&#x60;&#x60;&#x60;

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/56#issuecomment-44567426) on: **Thursday, May 29, 2014**

Hi @astanush, It seems that there is a typo, it should be &#x60;extensions&#x60; plural. If you change that it should work! that code hasn&#x27;t been [changed in over a year](https://github.com/jekyll/jekyll/blob/master/lib/jekyll/converters/markdown/redcarpet_parser.rb#L94). Try it out! and let us know.
---
> from: [**astanush**](https://github.com/jekyll/jekyll-help/issues/56#issuecomment-44568395) on: **Thursday, May 29, 2014**

Doh! That works, thanks! :+1: 
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/56#issuecomment-44568513) on: **Thursday, May 29, 2014**

No problem :smiley:! glad I could help!
---
> from: [**astanush**](https://github.com/jekyll/jekyll-help/issues/56#issuecomment-44581895) on: **Thursday, May 29, 2014**

Related: If I recall correctly, this only applies to markdown files. Is there a way to apply Smartypants for &#x60;.html&#x60; Jekyll pages?
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/56#issuecomment-44582280) on: **Thursday, May 29, 2014**

I really don&#x27;t know. Sorry about that.—
Alberto Grespan

On Thu, May 29, 2014 at 4:07 PM, Aaron Stanush &lt;notifications@github.com&gt;
wrote:

&gt; Related: If I recall correctly, this only applies to markdown files. Is there a way to apply Smartypants for &#x60;.html&#x60; Jekyll pages?
&gt; ---
&gt; Reply to this email directly or view it on GitHub:
&gt; https://github.com/jekyll/jekyll-help/issues/56#issuecomment-44581895
---
> from: [**mattr-**](https://github.com/jekyll/jekyll-help/issues/56#issuecomment-44586019) on: **Thursday, May 29, 2014**

There’s no way to apply SmartyPants for .html files from within Jekyll.
