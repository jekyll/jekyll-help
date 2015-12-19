# [liquid templating not evaluated in collection when parsing all collection items together on one page](https://github.com/jekyll/jekyll-help/issues/118)

> state: **closed** opened by: **** on: ****

I have added a collection called &quot;sections&quot; to my jekyll site. When I try to display all the sections (markdown pages) on the index.html page, the liquid template directives are not evaluated. 

E.g. in 01-test.md:  &#x60;{{ site.url }}&#x60; is displayed as {{ site.url }}

When I slightly change the site and use _posts instead of _sections, the liquid tags are correctly evaluated. 

Example: https://github.com/enketo/odk-xform-spec (which is live [here](http://enketo.github.io/odk-xform-spec/))

Is there a way to get liquid tags of pages inside collections to be evaluated?

### Comments

---
> from: [**MartijnR**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-51843914) on: **Monday, August 11, 2014**

Just noticed markdown is also not rendered. Hmmm, I must be missing something elementary here.

update:
Using this inside the site.sections loop solves the markdown issue but not the liquid issue:
&#x60;&#x60;&#x60;
{% capture content %}{{ section.content }}{% endcapture %}
{{ content | markdownify }}
&#x60;&#x60;&#x60; 
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-51845691) on: **Monday, August 11, 2014**

Can you double check that you&#x27;re running a version of Jekyll that supports collections?
---
> from: [**MartijnR**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-51846895) on: **Monday, August 11, 2014**

Confirmed. I&#x27;m using version 2.2.0
---
> from: [**ivantsepp**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-52883992) on: **Wednesday, August 20, 2014**

You need to set &#x60;output: true&#x60; to render your collection files ([doc](http://jekyllrb.com/docs/collections/#step-3-optionally-render-your-collections-documents-into-independent-files)).

So, in your repo:
&#x60;&#x60;&#x60;yaml
collections:
  sections:
    output: true
&#x60;&#x60;&#x60;
---
> from: [**MartijnR**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-52926203) on: **Thursday, August 21, 2014**

Thanks @ivantsepp. I have tried this (and now also updated that in the repo linked above), but it does not resolve the issue.
---
> from: [**ivantsepp**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-52926782) on: **Thursday, August 21, 2014**

http://enketo.github.io/odk-xform-spec/sections/01-test.html, does that look right to you?
---
> from: [**MartijnR**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-52928065) on: **Thursday, August 21, 2014**

Yes, it does.  Sorry, I should rewrite that the issue I&#x27;m experiencing occurs when trying to parse all the sections together on the home page: https://github.com/enketo/odk-xform-spec/blob/gh-pages/index.html. 

I&#x27;ll change the description. 

I won&#x27;t actually require each section to be published on a separate html page.
---
> from: [**MartijnR**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-52930532) on: **Thursday, August 21, 2014**

This is where I feel I&#x27;m missing something to get liquid tags to be processed properly when iterating through a collection: https://github.com/enketo/odk-xform-spec/blob/gh-pages/index.html#L18
---
> from: [**ivantsepp**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-52930948) on: **Thursday, August 21, 2014**

Ah, sorry for misunderstanding your issue in the first place. In that case, &#x60;output: true&#x60; isn&#x27;t needed. Can you try using &#x60;{{section.output}}&#x60; instead?
---
> from: [**MartijnR**](https://github.com/jekyll/jekyll-help/issues/118#issuecomment-52932214) on: **Thursday, August 21, 2014**

Ouch, I missed that in the documentation!

Thanks so much! Works like a charm now!
