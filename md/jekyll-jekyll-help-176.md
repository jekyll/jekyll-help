# [Best mechanisms for single source in multiple locations](https://github.com/jekyll/jekyll-help/issues/176)

> state: **open** opened by: **** on: ****

Consider the requirement for a document site with one more more sections that have common topics. Each section has its own TOC and can navigate to the next/prev items in that TOC using prev-&gt;next buttons. The documents may have slightly different rendering in each section (and of course a different TOC/next/prev). There should be a single &quot;content&quot; document for the common content.

The obvious way to implement something like this is to have the &quot;common content&quot; source as an *_include*. Then have &quot;stub&quot; articles for each TOC item that imports it. This would work nicely and allow me to feed in &quot;version specific&quot; configuration in the front matter.

Are there any other ways that people have achieved this? For example, is there a way to automatically copy appropriate documents into specified locations at build time (from within Jekyll)?

Thank you!

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/176#issuecomment-61438594) on: **Sunday, November 02, 2014**

You might want to take a gander at [collections](http://jekyllrb.com/docs/collections/) for your shared content, but I might have misunderstood your situation. Can you give some more detail, and maybe an example of exactly what you&#x27;re after?
---
> from: [**hamishwillee**](https://github.com/jekyll/jekyll-help/issues/176#issuecomment-61439680) on: **Sunday, November 02, 2014**

Hi Troy,

Collections might be part of the solution to my problem, but not the whole problem.

Consider the imaginary case where I have a site supporting a number of manuals for using and assembling bicycles. Each section of the site has its own table of contents (and I might reasonably have these in collections). While some articles will be different for each bike, a few will be common (for example &quot;How to ring the bell&quot;).

My question was &quot;what is the BEST&quot; method for maintaining this common content so it renders in all places where it is used. I THINK that we need a separate stub for each instance of the article that includes a common document that we don&#x27;t render. The include markup allows this, but only for files in _include (which I guess is OK). 
 Is there any alternative? E.g. can we somehow declare in the file where it should build to as a first step. 

Hope that makes it clear!
