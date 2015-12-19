# [After migration, how can I redirect old URLs to new ones?](https://github.com/jekyll/jekyll-help/issues/282)

> state: **open** opened by: **** on: ****

Hi,
I&#x27;m migrating my blog from Wordpress to Jekyll, but I&#x27;m concerned about preserving old URLs from 404.

The ideal would be to redirect old URLs, which are not many (about 20), to the new ones with a 304.
How can I do this?

**Example:**
* Old URL: http://www.andreaverlicchi.eu/advanced-sprite-generation-using-compass-sass/
* New URL: http://www.andreaverlicchi.eu/development/2015/03/06/advanced-sprite-generation-using-compass-sass.html
...

I don&#x27;t expect to do the &quot;old URL - new URL&quot; association automatically... no problem to do it automatically, but I&#x27;d like to know what is the best way to do the redirect.

**Additional info:**
My site is hosted on a linux server supporting PHP.

### Comments

---
> from: [**willnorris**](https://github.com/jekyll/jekyll-help/issues/282#issuecomment-77776572) on: **Sunday, March 08, 2015**

It won&#x27;t give you the &#x60;30x&#x60; response codes, but https://github.com/jekyll/jekyll-redirect-from provides a pretty simple way to redirect from old URLs.  It does a client-side redirect, which I&#x27;m fairly certain Google still treats like a 301 redirect, so you still get the SEO benefits.

If you still really want the &#x60;30x&#x60; response code, you&#x27;ll need to do it in your web server itself.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/282#issuecomment-78794638) on: **Thursday, March 12, 2015**

Yes, +1 to @willnorris, this is may be easiest way to do it, you should give it a try and see if you like it or not... if you&#x27;ll be hosting your own jekyll blog you can also use apache htaccess or nginx to do this. Hope this helps.
