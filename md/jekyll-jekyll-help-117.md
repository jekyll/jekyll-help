# [Why does Jekyll regenerate when I save css or add an image?](https://github.com/jekyll/jekyll-help/issues/117)

> state: **closed** opened by: **** on: ****

My site is getting larger and the regeneration times are getting close to a minute whenever I edit something (locally). I only get annoyed when it does this after I save a css file or add an image. Couldn&#x27;t there be a filter where if a non transformable file like css or jpg is changed, it just gets copied over by itself rather than regenerating the entire site?

I suppose this could be a problem for sass or whatever (I still have a normal css file). Maybe there could be a config option to allow certain file types to be copied over and skip regeneration?

So far I have about 300 posts. The posts themselves don&#x27;t really slow it down, it is all the tag and category lists that I have I think.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/117#issuecomment-51823948) on: **Monday, August 11, 2014**

Jekyll isn&#x27;t too smart in this regard. It just looks for changes in the source directory and regenerates the entire site when it sees them.

@mattr- has been working on this in jekyll/jekyll#1761
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/117#issuecomment-51830726) on: **Monday, August 11, 2014**

yes, that sounds like what I am looking for. Thanks for the info.
