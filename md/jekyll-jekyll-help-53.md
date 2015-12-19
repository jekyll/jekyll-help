# [Documentation matching github pages](https://github.com/jekyll/jekyll-help/issues/53)

> state: **closed** opened by: **** on: ****

Where will I find documentation that matches the version of Jekyll running on github pages?   Or otherwise how can I know what features are and are not supported?

I tried setting the &quot;front matter defaults&quot; according to the documentation only to find after some digging that this feature became available in jekyll 2 but GHP is running 1.5.1.    Will I have to cross-reference every feature in the documentation with the changelog or can I see the docs for 1.5.1 somewhere?

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/53#issuecomment-44146533) on: **Sunday, May 25, 2014**

The easiest way is to match what your development jekyll with GitHub Pages jekyll is to use the &#x60;github-pages&#x60; gem. AFAIK the documentation is not versioned so searching only the 1.5.1 docs will not be posible, although I&#x27;m not sure. You can take a look at the changelog [here](http://jekyllrb.com/docs/history/#section-4) 

&#x60;&#x60;&#x60;sh
$ gem install github-pages
&#x60;&#x60;&#x60;

or add it to your &#x60;Gemfile&#x60;, the [latest version](http://rubygems.org/gems/github-pages) is &#x60;19&#x60;.

&#x60;&#x60;&#x60;sh
gem &#x27;github-pages&#x27;, &#x27;~&gt; 19&#x27;
&#x60;&#x60;&#x60;
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/53#issuecomment-45025057) on: **Tuesday, June 03, 2014**

https://pages.github.com/versions/
---
> from: [**avdd**](https://github.com/jekyll/jekyll-help/issues/53#issuecomment-45043492) on: **Tuesday, June 03, 2014**

Neither of those comments address my issue which is that the gh-pages *docs* link to jekyll *docs* that don&#x27;t match the production version.  Namely, that the docs are misleading.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/53#issuecomment-45048102) on: **Tuesday, June 03, 2014**

@avdd Don&#x27;t be so quick to judge – @albertogg&#x27;s comment was just about there. Run &#x60;jekyll docs&#x60; on the command line with the Jekyll version &#x60;github-pages&#x60; installs. This will launch a server (much the same as &#x60;jekyll serve&#x60; will for your local site) with the Jekyll documentation for that version on &#x60;http://localhost:4000&#x60;. This is the correct documentation for this version of Jekyll.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/53#issuecomment-45048189) on: **Tuesday, June 03, 2014**

Additionally, front-matter defaults are a Jekyll 2.0 feature, and by going to https://pages.github.com/versions/, you would see that GitHub Pages (and the &#x60;github-pages&#x60; gem) have not been updated to Jekyll 2.0. This means front-matter defaults are not available on GitHub Pages (yet).
---
> from: [**avdd**](https://github.com/jekyll/jekyll-help/issues/53#issuecomment-45054066) on: **Tuesday, June 03, 2014**

Thank you.  Perhaps mention &#x60;jekyll docs&#x60; [here](https://help.github.com/articles/using-jekyll-with-pages) so it&#x27;s clearer what features are supported.
