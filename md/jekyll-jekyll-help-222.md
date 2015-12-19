# [Can I create a association between two posts (e.g. for multilingual posts)?](https://github.com/jekyll/jekyll-help/issues/222)

> state: **closed** opened by: **** on: ****

I have a bilingual site. For each post, I like to link to its counterpart. Currently, there is no logical connection between the two posts.

The German posts (default) reside in &#x60;/_posts&#x60; while the English posts are located in &#x60;/en/_posts&#x60;. The category &#x60;en&#x60; which is automatically created by having the posts inside &#x60;/en&#x60; is used for permalinks.

__My question is:__ Is there a way I can _entangle_ two posts (German and English) and create the links to the respective counterpart without writing them directly in my posts? What would be the strategy for that?

_(I let Pages do the build, so plugins won&#x27;t work. Also, I don&#x27;t want working code, just a small entry point.)_

### Comments

---
> from: [**mattr-**](https://github.com/jekyll/jekyll-help/issues/222#issuecomment-68529203) on: **Friday, January 02, 2015**

There&#x27;s not really any way to do this currently. It&#x27;s up to the site creator to determine how they want to handle multilingual sites and how they link their posts together.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/222#issuecomment-68659958) on: **Sunday, January 04, 2015**

As @mattr- said, there isn&#x27;t anything built in, but you could create a loop in your posts layout which looks for the translated counterpart and outputs some HTML if it finds something.

Hope this helps.
