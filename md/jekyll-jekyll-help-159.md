# [Custom sitemap overwritten in production](https://github.com/jekyll/jekyll-help/issues/159)

> state: **open** opened by: **** on: ****

(Via hafniatimes/issues#12.)

I have a [manual sitemap file](https://github.com/hafniatimes/hafniatimes.github.io/blob/master/sitemap.xml) I wish to use for my site in production. For some reason, though, I get [a different output](http://hafniatimes.com/sitemap.xml) in production than in my local &#x60;_site&#x60; folder.

I’ll post the content and output in a separate comment.

### Comments

---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/159#issuecomment-56801661) on: **Thursday, September 25, 2014**

Locally:

&#x60;&#x60;&#x60;xml
&lt;?xml version=&quot;1.0&quot; encoding=&quot;UTF-8&quot;?&gt;
&lt;urlset xmlns=&quot;http://www.sitemaps.org/schemas/sitemap/0.9&quot; xmlns:xhtml=&quot;http://www.w3.org/1999/xhtml&quot;&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/da/articles/2014/09/14/danish-polls.html&lt;/loc&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;da&quot; href=&quot;http://hafniatimes.com/da/articles/2014/09/14/danish-polls.html&quot; /&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;en&quot; href=&quot;http://hafniatimes.com/articles/2014/09/14nish-polls.html&quot; /&gt;
        &lt;lastmod&gt;2014-09-14T20:00:00+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/articles/2014/09/14/danish-polls.html&lt;/loc&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;en&quot; href=&quot;http://hafniatimes.com/articles/2014/09/14/danish-polls.html&quot; /&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;da&quot; href=&quot;http://hafniatimes.com/da/articles/2014/09/14/danish-polls.html&quot; /&gt;
        &lt;lastmod&gt;2014-09-14T20:00:00+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/articles/2014/08/29/antiterror-manhunt.html&lt;/loc&gt;
        &lt;lastmod&gt;2014-08-29T22:00:00+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/articles/2014/08/29/muslims-vs-jews.html&lt;/loc&gt;
        &lt;lastmod&gt;2014-08-29T20:00:00+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/articles/2014/08/02/a-sideorder-of-sexism.html&lt;/loc&gt;
        &lt;lastmod&gt;2014-08-02T14:00:00+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/articles/2014/07/25/danish-peoples-party-and-headscarves.html&lt;/loc&gt;
        &lt;lastmod&gt;2014-07-25T17:00:00+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/da/articles/2014/06/06/intro.html&lt;/loc&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;da&quot; href=&quot;http://hafniatimes.com/da/articles/2014/06/06/intro.html&quot; /&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;en&quot; href=&quot;http://hafniatimes.com/articles/2014/06/06/intro.html&quot; /&gt;
        &lt;lastmod&gt;2014-06-06T22:00:00+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/articles/2014/06/06/intro.html&lt;/loc&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;en&quot; href=&quot;http://hafniatimes.com/articles/2014/06/06/intro.html&quot; /&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;da&quot; href=&quot;http://hafniatimes.com/da/articles/2014/06/06/intro.html&quot; /&gt;
        &lt;lastmod&gt;2014-06-06T22:00:00+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/da/&lt;/loc&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;da&quot; href=&quot;http://hafniatimes.com/da/&quot; /&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;en&quot; href=&quot;http://hafniatimes.com/&quot; /&gt;
        &lt;lastmod&gt;2014-09-22T15:37:56+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/&lt;/loc&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;en&quot; href=&quot;http://hafniatimes.com/&quot; /&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;da&quot; href=&quot;http://hafniatimes.com/da/&quot; /&gt;
        &lt;lastmod&gt;2014-09-22T15:37:56+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
    &lt;url&gt;
        &lt;loc&gt;http://hafniatimes.com/404/&lt;/loc&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;&quot; href=&quot;http://hafniatimes.com/404/&quot; /&gt;
        &lt;xhtml:link rel=&quot;alternate&quot; hreflang=&quot;da&quot; href=&quot;http://hafniatimes.com/404/&quot; /&gt;
        &lt;lastmod&gt;2014-09-22T15:37:56+02:00&lt;/lastmod&gt;
    &lt;/url&gt;
&lt;/urlset&gt;
&#x60;&#x60;&#x60;

[Production](http://hafniatimes.com/sitemap.xml):

&#x60;&#x60;&#x60;
http://hafniatimes.com/articles/2014/09/14/danish-polls.html 2014-09-14T18:00:00+00:00 
http://hafniatimes.com/da/articles/2014/09/14/danish-polls.html 2014-09-14T18:00:00+00:00 
http://hafniatimes.com/articles/2014/08/29/antiterror-manhunt.html 2014-08-29T22:00:00+00:00 
http://hafniatimes.com/articles/2014/08/29/muslims-vs-jews.html 2014-08-29T18:00:00+00:00 
http://hafniatimes.com/articles/2014/08/02/a-sideorder-of-sexism.html 2014-08-02T12:00:00+00:00 
http://hafniatimes.com/articles/2014/07/25/danish-peoples-party-and-headscarves.html 2014-07-25T15:00:00+00:00 http://hafniatimes.com/articles/2014/06/06/intro.html 2014-06-06T20:00:00+00:00 
http://hafniatimes.com/da/articles/2014/06/06/intro.html 2014-06-06T20:00:00+00:00 
http://hafniatimes.com/ 2014-09-22T13:39:09+00:00 http://hafniatimes.com/da/ 2014-09-22T13:39:09+00:00 http://hafniatimes.com/404/ 2014-09-22T13:39:09+00:00
&#x60;&#x60;&#x60;
---
> from: [**mattr-**](https://github.com/jekyll/jekyll-help/issues/159#issuecomment-56809855) on: **Thursday, September 25, 2014**

Is the sitemap generated by a plugin? If so, it&#x27;s possible that the plugin isn&#x27;t running correctly in production. 

Also, could the meaning of &#x27;production&#x27; be clarified? Is it GitHub Pages?
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/159#issuecomment-56810396) on: **Thursday, September 25, 2014**

The sitemap is just generated from the sitemap.xml file—with a (--- Jekyll
YAML header). The (production) site is built and hosted with GitHub Pages,
yes.

I dont’t know whether the jekyll-sitemap gem has to be disabled explicitly
in my config somehow, if that’s what causing it. Something is definitely
superceding my own sitemap.xml file.

On 25 September 2014 14:14, Matt Rogers &lt;notifications@github.com&gt; wrote:

&gt; Is the sitemap generated by a plugin? If so, it&#x27;s possible that the plugin
&gt; isn&#x27;t running correctly in production.
&gt;
&gt; Also, could the meaning of &#x27;production&#x27; be clarified? Is it GitHub Pages?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub
&gt; &lt;https://github.com/jekyll/jekyll-help/issues/159#issuecomment-56809855&gt;.
&gt;
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/159#issuecomment-56810570) on: **Thursday, September 25, 2014**

No, &#x60;jekyll-sitemap&#x60; is only activated when explicitly stated.
