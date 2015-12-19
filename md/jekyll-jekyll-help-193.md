# [How do you deal with excess newlines?](https://github.com/jekyll/jekyll-help/issues/193)

> state: **closed** opened by: **** on: ****

Forloops, if-else conditionals and all that Liquid goodness leaves A LOT of trailer lines in my compiled code. Aside from running an HTML linter on top of the result, are there any good ways to deal with the cosmetic problem?

### Comments

---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/193#issuecomment-63063546) on: **Friday, November 14, 2014**

Are you concerned about the added whitespace, or the formatting? If it is the former, you could use a layout to minimise the html:

https://github.com/penibelst/jekyll-compress-html

If it is the latter then you can only really deal with that by formatting your pre-compiled code horribly, and I wouldn&#x27;t advise that personally.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/193#issuecomment-63072190) on: **Friday, November 14, 2014**

Post-processing (&#x60;minify&#x60;, etc.) is the simplest and most-straight-forward way to manage compiled/rendered source of course.

However, in the past, I have created liquid &quot;functions&quot; with liquid filters that provide for stripping newlines, etc. which accomplish more or less the same thing.

See the page, &quot;Liquid for Designers&quot;, and reference some of the stripping filters that remove the 
offending newlines, returns, whitespace, etc.

i.e. &#x60;strip_newlines - strip all newlines (\n) from string&#x60;

https://github.com/Shopify/liquid/wiki/Liquid-for-Designers


---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/193#issuecomment-63080558) on: **Friday, November 14, 2014**

So the idea would be to &#x60;{% capture %}&#x60; the resulting output and apply &#x60;strip_newlines&#x60; to it?
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/193#issuecomment-63087354) on: **Friday, November 14, 2014**

Yes; something like that... in some form or fashion, depending on your end goal.

I created routines/functions that handled doing that via an &#x60;include&#x60; I recall.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/193#issuecomment-63099574) on: **Friday, November 14, 2014**

Cool beans, sounds doable, especially within the include itself!
