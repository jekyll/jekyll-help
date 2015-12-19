# [Please document how to invoke liquid filters from plugins](https://github.com/jekyll/jekyll-help/issues/97)

> state: **closed** opened by: **** on: ****

There is a general issue here regarding documentation which would (I think) resolve my specific issue.

Basically it is not obvious as to how to invoke liquid filters from jekyll plugins. In some cases you can look through the source code to find the specific filter. Some live in liquid itself, in &#x60;standardfilters.rb&#x60;. There are others which live in jekyll, in &#x60;filters.rb&#x60;.

Looking in the source code is apparently required, in order to identify the ruby module for the relevant function, and of course the function name itself. The [documentation](http://docs.shopify.com/themes/liquid-documentation/filters) does not contain this information.

However, apparently grepping the installed gems for the filter name is not sufficient to locate the source for some filters. Specifically, I can&#x27;t locate the source for the &#x60;handleize&#x60; filter, and hence cannot invoke it from a plugin. Maybe it&#x27;s just me, or some weird quirk of my install.

I think all that would be required is to add module::function declaration for each filter in the documentation. This would help me at least. Thanks.

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49188513) on: **Wednesday, July 16, 2014**

This is a very interesting issue. I don&#x27;t think filters were meant to be called from plug-ins which is why the documentation does not exist. What is your usecase?
---
> from: [**randomphrase**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49343031) on: **Thursday, July 17, 2014**

OK, so I have to ask why this would be the case? I&#x27;m going to take a guess and say that filters are intended to be simple wrappers around ruby library calls, and hence if you need to call the filter from a plugin you should in fact be calling the underlying ruby library instead. This sounds reasonable, but again kind of exposes a bit of a documentation problem... (What is the underlying ruby call in each case?)

Anyway my specific use case is as follows. I&#x27;m creating a per-category archive page. I&#x27;m copying the sample plugin exactly, but I found that my category names need to be converted into something that can be put into a URL (and ideally match that of my original wordpress blog). For example &quot;foo bar baz&quot; category needs to be converted to &quot;foo-bar-baz&quot;. This needs to be done from within liquid (so that each blog post can link to the relevant category archive) and from the plugin (so that the category archive page can be created).

So &#x60;handleize&#x60; does what I want, at least on the liquid side of things, and that&#x27;s what I&#x27;m using. On the plugin side it&#x27;s not so easy. Now I could easily re-implement the &#x60;handleize&#x60; functionality in ruby, at least to the extent that it works for my categories, but this seems wrong. &quot;It should be easy to call the handleize function directly!&quot; I thought, apparently incorrectly. I guess I can do this if there&#x27;s no better solution, but again it seems a bit hacky.

Appreciate any further assistance, thanks.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49349191) on: **Thursday, July 17, 2014**

&gt; I have to ask why this would be the case?

Liquid Filters (which is what these are) are for use in Liquid templates. That&#x27;s why they exist – they aren&#x27;t meant to be used in any Ruby code other than Liquid&#x27;s internals. Why wouldn&#x27;t you just do &#x60;{{ categories | handleize }}&#x60; from your category archive page liquid template?

At any rate, it&#x27;s luckily quite straightforward to use these filters in your ruby code. They&#x27;re all modules, so you can include them:

&#x60;&#x60;&#x60;ruby
class CategoryArchivePage
  include Filters # this will change depending upon the module &#x60;handleize&#x60; belongs to
  def handle
    handleize(categories)
  end
end
&#x60;&#x60;&#x60;
---
> from: [**randomphrase**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49357115) on: **Thursday, July 17, 2014**

OK so looking into it a bit closer, looks like &#x60;handleize&#x60; doesn&#x27;t work at all on my system. It just silently passes through whatever was passed to it. The example given in the documentation &#x60;{{ &#x27;100% M &amp; Ms!!!&#x27; | handleize }}&#x60; outputs &quot;100% M &amp; Ms!!!&quot; not the expected &quot;100-m-ms&quot;. How can I resolve this?

Is it possible I am missing some gem or other? Again: where in the code is this &#x60;handleize&#x60; function actually defined, because I can&#x27;t see it. Is this some kind of shopify-specific function?

Thanks for the example code, but knowing the exact module name to include is exactly the problem I raised this ticket for. How am I to determine this other than asking, or just knowing it?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49360687) on: **Thursday, July 17, 2014**

@randomphrase Sorry for the confusion. If you can&#x27;t find the code for it in either the Liquid or Jekyll source, it isn&#x27;t available to Liquid.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49360744) on: **Thursday, July 17, 2014**

One option is to use &#x60;stringex&#x60; package and write your own handleize filter. Where are you seeing this handleize filter, anyway?
---
> from: [**randomphrase**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49446243) on: **Friday, July 18, 2014**

I think this is most likely a documentation issue. The [Jekyll documentation](http://jekyllrb.com/docs/templates/) says that &quot;all of the standard Liquid filters are available&quot; and provides a link to the shopify wiki, which in turn [describes the handelize filter](http://docs.shopify.com/themes/liquid-documentation/filters/string-filters#handle)

So I guess one solution would be to remove the link to the shopify wiki, and add a description of (only) the standard liquid filters which are supported. Alternatively the shopify wiki could be updated to describe which were standard liquid filters and which were extensions.

Is it a bug that using an unknown filter doesn&#x27;t seem to generate an error, or even a warning?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49448493) on: **Friday, July 18, 2014**

@randomphrase Yeah, we should remove the shopify wiki link. It has proprietary filters, because it&#x27;s for shopify.com templates (recently expanded beyond core liquid). Ultimately, we should point to docs for Liquid core, which are currently sparse as ever: http://shopify.github.io/liquid/

@fw42, do you have a designer I can work with to rewrite the shopify.github.io/liquid/ docs? We could version them alongside the liquid code so you can get docs &amp; new code all in one PR.

&gt; Is it a bug that using an unknown filter doesn&#x27;t seem to generate an error, or even a warning?

I&#x27;d certainly prefer Liquid to error out, but it&#x27;s the Liquid core team&#x27;s decision there. We should use the strict parser if there is one.
---
> from: [**fw42**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49459202) on: **Friday, July 18, 2014**

@Tetsuro is one of the maintainers of the Shopify Documentation for Liquid. I talked with him and we will try to come up with something. We know that the Liquid core docs are terribly bad :-(
---
> from: [**Tetsuro**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49459511) on: **Friday, July 18, 2014**

Hey guys! Tets here, designer at Shopify. I recently re-did the [Liquid Documentation for Shopify](http://docs.shopify.com/themes). 

I totally understand where you guys are coming from. I don&#x27;t have an immediate solution for this, but @fw42 and I will start brainstorming some solutions.  
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49462313) on: **Friday, July 18, 2014**

@Tetsuro @fw42 Thanks, guys! I am more than happy to help where I can. Just ping me in an issue/pr and I will jump in!
---
> from: [**randomphrase**](https://github.com/jekyll/jekyll-help/issues/97#issuecomment-49462631) on: **Friday, July 18, 2014**

Thanks a lot guys - I really appreciate the help
