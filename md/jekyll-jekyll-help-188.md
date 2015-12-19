# [Stop Jekyll Wrapping Images in P Tags](https://github.com/jekyll/jekyll-help/issues/188)

> state: **closed** opened by: **** on: ****

Can somebody advise on how it is possible to stop jekyll wrapping p tags around image tags using markdown

Regards

Mark

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/188#issuecomment-62296020) on: **Sunday, November 09, 2014**

That is a Markdown thing and pretty sane if you ask me. What do you need it for, though? 
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/188#issuecomment-62310373) on: **Sunday, November 09, 2014**

&#x60;&lt;img&gt;&#x60; is an inline element, so it should *always* exist within a block level element. So I would agree, the markdown converters behaviour is entirely correct, in my opinion.

This behaviour is also detailed in the CommonMark Spec, the most detailed Markdown specification that currently exists:

http://spec.commonmark.org/0.10/#images 

That said, if you don&#x27;t want your images inside &#x60;&lt;p&gt;&#x60; you could place them inside another suitable block-level element like &#x60;&lt;div&gt;&#x60; or &#x60;&lt;figure&gt;&#x60; but, from a semantics standpoint, I don&#x27;t see anything wrong with &#x60;&lt;p&gt;&#x60; unless &#x60;&lt;figure&gt;&#x60; is a better fit.
---
> from: [**juicymark**](https://github.com/jekyll/jekyll-help/issues/188#issuecomment-62311295) on: **Sunday, November 09, 2014**

Hi nternetinspired, Thanks for the reply.

I needed to stop the img tag being wrapped in a p tag since my paragraphs contain a margin on the bottom so its messing up my styling, since its a HTML5 document i have decided to go with using the &lt;code&gt;&amp;lt;figure&amp;gt;&lt;/code&gt; tag around any images.

Thanks
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/188#issuecomment-62312009) on: **Sunday, November 09, 2014**

Why is using the figure tag around images better? Are all your images figures? 
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/188#issuecomment-62353039) on: **Monday, November 10, 2014**

&gt; Why is using the figure tag around images better? Are all your images figures?

This. Images are not always figures, and if they aren&#x27;t you&#x27;d be better not using that element. No semantics are better than incorrect semantics.
