# [Jekyll suddenly not regenerating on two machines ...](https://github.com/jekyll/jekyll-help/issues/216)

> state: **closed** opened by: **** on: ****

SEE COMMENT BELOW. Probable bug in regeneration logic.

Jekyll has generally automatically regenerated after I run &quot;jekyll serve&quot;. This weekend, it stopped. More interestingly, it stopped on both my MacBook Air and my iMac. I am running the same site on both, pulling from my rj-com repo. 

So, I&#x27;m imagining that there could be something wrong with my site, perhaps in a plugin? I&#x27;ve got a few, including a couple that my colleague and I wrote. The site generates OK the first time and after that nothing happens. Advise me, please.

### Comments

---
> from: [**RonJeffries**](https://github.com/jekyll/jekyll-help/issues/216#issuecomment-67916139) on: **Monday, December 22, 2014**

Update: issue was caused by string &#x27;spec&#x27; in exclude: in _config.yml:

exclude:
- spec
- README.MD

this causes all folders whose names include &#x27;spec&quot; to ignore saves of files inside them for regeneration. when regeneration is triggered manually (typing jekyll serve again) the files inside are not ignored for the build and display correctly on the site.

I filed an issue on the main jekyll project, or tried to anyway.

Thanks!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/216#issuecomment-68660717) on: **Sunday, January 04, 2015**

Alrighty then. Let us know if you need any more help!
