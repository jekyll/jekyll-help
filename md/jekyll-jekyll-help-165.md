# [sort or sort_by filter does not work](https://github.com/jekyll/jekyll-help/issues/165)

> state: **closed** opened by: **** on: ****

It seems like this is a bug in &#x60;sort&#x60;. How is sorting supposed to work? I&#x27;m getting errors when I use either the &#x60;sort&#x60; or the &#x60;sort_by&#x60; filter in my Liquid template. [Liquid docs](http://docs.shopify.com/themes/liquid-documentation/filters/array-filters#sort) say it&#x27;s called &#x60;sort&#x60;, but [others](http://cheat.markdunkley.com/) say it&#x27;s now called &#x60;sort_by&#x60;. Which am I supposed to use here?
&#x60;&#x60;&#x60;html
{% assign sorted_articles = site.data.articles  | sort:&quot;julian&quot; %}
&#x60;&#x60;&#x60;
I have a yaml file containing references to magazine articles. I&#x27;m trying to sort them by date. I&#x27;ve tried expressing dates in three different formats. Liquid fails to parse &#x60;YYYY-MM-DD&#x60; as a date, so &#x60;sort: &quot;date&quot;&#x60; sorts alphabetically instead of chronologically. Here is an example (&#x60;julian&#x60; is the same value as &#x60;date&#x60; and &#x60;iso_date&#x60;, just expressed as a Julian date integer):

    - title: &#x27;Measuring Information Work Productivity&#x27;
      desc: &#x27;Computerworld&#x27;
      url: &#x27;measuring-info.html&#x27;
      date: &#x27;2004-3&#x27;
      iso_date: &#x27;2004-3-1 00:00:01&#x27;
      julian: 731641

 &#x60;sort&#x60; works fine on my local machine with &#x60;jekyll serve&#x60;, but it gives an error (see below) with &#x60;jekyll serve --safe&#x60; and it fails to build on github-pages.

 &#x60;sort_by&#x60; builds without error on my local machine and builds successfully on github-pages. But the results are useless. (1) On my local machine, the articles aren&#x27;t actually sorted, they appear to be in random order. (2) With &#x60;--safe&#x60; or on gh-pages, the sorted array appears to be empty instead of being sorted.

Example of error when using &#x60;sort&#x60; (fails with &#x60;--safe&#x60; and fails to build on github-pages):
&#x60;&#x60;&#x60;
$ jekyll serve --watch
    Configuration file: /Users/sstrassmann/src/jekhub/pas/_config.yml
            Source: /Users/sstrassmann/src/jekhub/pas
       Destination: /Users/sstrassmann/src/jekhub/pas/_site
      Generating... 
                    done.
     Auto-regeneration: enabled for &#x27;/Users/sstrassmann/src/jekhub/pas&#x27;
Configuration file: /Users/sstrassmann/src/jekhub/pas/_config.yml
    Server address: http://0.0.0.0:4000/pas/
  Server running... press ctrl-c to stop.
 ^C
$ jekyll serve --watch --safe
Configuration file: /Users/sstrassmann/src/jekhub/pas/_config.yml
            Source: /Users/sstrassmann/src/jekhub/pas
       Destination: /Users/sstrassmann/src/jekhub/pas/_site
      Generating...     
  Liquid Exception: undefined method &#x60;sort&#x27; for nil:NilClass in articles.html
jekyll 2.4.0 | Error:  undefined method &#x60;sort&#x27; for nil:NilClass
&#x60;&#x60;&#x60;
So, what is the preferred way to sort a yaml array?

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/165#issuecomment-61440088) on: **Sunday, November 02, 2014**

Looks like it as simple as adding some parens around a part of your first &#x60;assign&#x60; block:

&#x60;&#x60;&#x60;html
{% assign sorted_articles = (site.data.articles | sort: &#x27;julian&#x27;) %}
&#x60;&#x60;&#x60;

Stole from [this article](http://www.leveluplunch.com/blog/2014/04/03/sort-pages-by-title-filter-array-by-layout-jekyllrb/).
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/165#issuecomment-61504549) on: **Monday, November 03, 2014**

Here&#x27;s an example that I cooked up. Works like a champion.

https://github.com/troyswanson/troyswanson.github.io/commit/6785844d56e62a4c51f19956b1ba901138947525
---
> from: [**straz**](https://github.com/jekyll/jekyll-help/issues/165#issuecomment-61514504) on: **Monday, November 03, 2014**

Cool, thank you, Troy!

I wish the provided example of an &#x60;assign&#x60; block in [the documentation](http://docs.shopify.com/themes/liquid-documentation/filters/array-filters#sort) used parens in the same way.

Does this mean that the new [&#x60;sort_by&#x60;](http://cheat.markdunkley.com/) tag is now deprecated?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/165#issuecomment-61552304) on: **Monday, November 03, 2014**

&#x60;¯\_(ツ)_/¯&#x60;

Hopefully this issue will show up in a search for folks who are having the same issue. Glad I could help!
