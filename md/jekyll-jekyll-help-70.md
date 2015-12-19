# [configuration documentation](https://github.com/jekyll/jekyll-help/issues/70)

> state: **closed** opened by: **** on: ****

Where can I find a documented list of the configuration options?
[The one on the main website](http://jekyllrb.com/docs/configuration/) doesn&#x27;t contain them all, and the explanations are not very clear.

I&#x27;ve got the feeling that some of those pages are not really up to date, e.g. the [upgrade notes](http://jekyllrb.com/docs/upgrading/), that talk about version &#x60;1.0&#x60;.


At the moment I&#x27;m trying to figure out what&#x27;s the purpose of the &#x60;url&#x60; option, which doesn&#x27;t seem to have any effect.
Is it just some kind of meta-data to be referenced in the other pages?

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/70#issuecomment-45443742) on: **Sunday, June 08, 2014**

Yes, &#x60;url&#x60; has no meaning – it&#x27;s just a piece of metadata. Sorry that wasn&#x27;t clear!
---
> from: [**tompave**](https://github.com/jekyll/jekyll-help/issues/70#issuecomment-45444685) on: **Sunday, June 08, 2014**

Thank you for confirming this!
Does this mean that the main website is still the most _authoritative_ source of documentation?  

I&#x27;m still in the process of understanding how the configuration options work.

In the case of &#x60;url&#x60; is this a correct way to use it?

&#x60;&#x60;&#x60;html
&lt;a href=&quot;{{ post.url | prepend: site.baseurl | prepend: site.url }}&quot;&gt;
  {{ post.title }}
&lt;/a&gt;
&#x60;&#x60;&#x60;
Os is there a more straightforward way to reference it?

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/70#issuecomment-45445775) on: **Sunday, June 08, 2014**

&gt; Does this mean that the main website is still the most authoritative source of documentation?

Yes yes yes! What the line you posted does is output the post URL, then prepend it with two metadata pieces: the baseurl (i.e. &#x60;/hi&#x60; in &#x60;http://example.org/hi&#x60;) and the site&#x27;s URL, (i.e. &#x60;http://example.org&#x60; in &#x60;http://example.org/hi&#x60;). You can set these in the &#x60;_config.yml&#x60; file, but both are optional. They&#x27;re there so that if your site requires a base url or the full URL, you don&#x27;t have to change the templates, just the configuration file.
