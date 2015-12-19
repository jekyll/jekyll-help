# [New pushes aren&#x27;t appearing in blog](https://github.com/jekyll/jekyll-help/issues/93)

> state: **closed** opened by: **** on: ****

Hi!

So I&#x27;ve been pushing to my gh-pages for my blog and I can see the changes/added posts take affect within my Github repo under gh-pages, but when I go to my blog the changes don&#x27;t appear. It&#x27;s been a few days now since I&#x27;ve noticed this happening and it still hasn&#x27;t updated. Any ideas on what could be going wrong?

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/93#issuecomment-48222131) on: **Monday, July 07, 2014**

Is this your repo? https://github.com/trevordjones/trevordjones.github.io

If so, there isn&#x27;t actually a gh-pages branch. Are you sure you&#x27;re pushing to the right remote repo?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/93#issuecomment-48223002) on: **Monday, July 07, 2014**

@kleinfreund For user pages (like the one you linked to), the &#x60;master&#x60; branch is used.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/93#issuecomment-48223207) on: **Monday, July 07, 2014**

@trevordjones Use &#x60;gh-pages&#x60; branch for project pages or &#x60;master&#x60; for your user pages repo (your trevordjones.github.io repo). Be sure to check your email – failed builds will trigger an email notification with information as to why it failed.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/93#issuecomment-48223821) on: **Monday, July 07, 2014**

@parkr Yeah, sure, I&#x27;m rockin&#x27; the same. I asked, because I sometime didn&#x27;t point to an actual remote repo as well.
---
> from: [**trevordjones**](https://github.com/jekyll/jekyll-help/issues/93#issuecomment-48233060) on: **Monday, July 07, 2014**

@kleinfreund it&#x27;s my blog repo at https://github.com/trevordjones/blog/ - I am getting a build error in my email (google filtered it out) but it didn&#x27;t give me a reason why. I have to keep digging.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/93#issuecomment-48235553) on: **Monday, July 07, 2014**

@trevordjones I notice your gemfile points to Jekyll, which is at &#x60;&#x60;&#x60;2.1&#x60;&#x60;&#x60;, but gh-pages is at &#x60;&#x60;&#x60;1.5.1&#x60;&#x60;&#x60; so you need &#x60;&#x60;&#x60;gem &#x27;github-pages&#x27;&#x60;&#x60;&#x60; as long as you&#x27;re not using any 2.x features.
---
> from: [**trevordjones**](https://github.com/jekyll/jekyll-help/issues/93#issuecomment-48252883) on: **Monday, July 07, 2014**

Thanks @budparr - turns out it wasn&#x27;t accepting my post_url links. I feel like that should be documented somewhere at http://jekyllrb.com/docs/templates/#post_url - but alas, they don&#x27;t mention it anywhere.
