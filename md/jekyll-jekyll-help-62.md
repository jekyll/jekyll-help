# [Examples of pages (not posts) on multilingual Jekyll sites ](https://github.com/jekyll/jekyll-help/issues/62)

> state: **closed** opened by: **** on: ****

I can see how you can make a site with multilingual posts using templates like [this](https://github.com/tristen/multilingual) (I’m less keen on how it works using tags like in [this](http://developmentseed.org/blog/multilingual-jekyll-sites/) example), but can someone give me an example of what a project would look like, if you also had to do multi-lingual pages with long-form text—on top of the posts?

Just a mock dir for starters.

### Comments

---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/62#issuecomment-44824468) on: **Monday, June 02, 2014**

Healthcare.gov provided a great example of a complex multi-lingual Jekyll site, IMO, it&#x27;s a real shame they back-pedalled on their commitment to open source and took down the repo.

You&#x27;ll find my fork of the original repo here: https://github.com/nternetinspired/healthcare.gov

Please fork it to help try and keep this excellent teaching resource alive. Thanks. 
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/62#issuecomment-44830105) on: **Monday, June 02, 2014**

Ah, so they did go with a folder for each language in root. Hmm.

Thanks for the link, @nternetinspired. And yes, it’s a pain they decided to axe the open source. That and squashing commits from contributors drives me bonkers, too.
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/62#issuecomment-44837053) on: **Monday, June 02, 2014**

My pleasure @ndarville,

That repo has been a treasure-trove of Jekyll tricks for me so I want to keep it alive so that others can learn from it. Long live open source! 
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/62#issuecomment-44841826) on: **Monday, June 02, 2014**

@ndarville I recently did a multi-lingual site. However, my goal was just to &quot;make it work&quot; and I can&#x27;t say if it&#x27;s any sort of a model of how it should be done. All I did was create direct linkages to a page&#x27;s other language mirror and then some conditional statements in the templates. Probably can be done more efficiently.
[source](https://github.com/sonnetmedia/linnullmann)

[site](https://github.com/sonnetmedia/linnullmann)

@nternetinspired  - funny I keep a fork of that site too for review/ideas.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/62#issuecomment-45061761) on: **Wednesday, June 04, 2014**

What do you do with URLs, and what do you recommend? Do you preserve the English URLs as in

&#x60;&#x60;&#x60;
   /about/
/es/about/
&#x60;&#x60;&#x60;

Or do you use a new translated folder directory instead? Or both. :)
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/62#issuecomment-45080676) on: **Wednesday, June 04, 2014**

Yes, I just used a folder for /en/ posts (for my site EN was the second language) and set the page permalinks that way. In my _data file I set a langauge key.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/62#issuecomment-45450991) on: **Sunday, June 08, 2014**

The healthcare.gov source really helped—because how the heck is anyone supposed to know that you had to put the translation in &#x60;_config.yml&#x60; without an example—so I’ve set up [a project](https://github.com/hafniatimes/hafniatimes.github.io) set to be multilingual.

I still have yet to figure out how to go about displaying dates; although [some guides](http://sylvain.durand.tf/multilingual-website-with-jekyll/) are available, it always ends up incredibly messy, which isn’t something I fancy. :)
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/62#issuecomment-55625836) on: **Monday, September 15, 2014**

@ndarville Thanks for the links. I&#x27;m going to close this for now, but if you&#x27;d like to open a new issue to discuss displaying links in different locale contexts, go right ahead! :smile: 
