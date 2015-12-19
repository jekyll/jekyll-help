# [Is it possible to iterate through site.posts and site.myCollection in dated order simultaneously?](https://github.com/jekyll/jekyll-help/issues/68)

> state: **closed** opened by: **** on: ****

Is there a good way to merge site.posts with a collection and iterate through them in dated order?

For example if my site.posts has:
-  Post A    24 May
-  Post B    30 May

And my site.myPhotoCollection
-  Photo A    1 June
-  Photo B    27 May

I&#x27;ve like to be able to create a listing page eg:

&#x60;&#x60;&#x60;
{% for each in merge(site.posts, site.myPhotoCollection) %}
    {{each.title}}
{% endfor %}
&#x60;&#x60;&#x60;&#x60;

and have the output be:
- Photo A
- Post B
- Photo B
- Post A



### Comments

---
> from: [**Daniel0524**](https://github.com/jekyll/jekyll-help/issues/68#issuecomment-45211574) on: **Thursday, June 05, 2014**

I&#x27;ve tried doing it with crazy sorting logic using liquid, but it&#x27;s such an enormous pain to do any sort of business logic with liquid tags. Any help would be great appreciated. Thanks!
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/68#issuecomment-45668600) on: **Tuesday, June 10, 2014**

I&#x27;m not sure about this and don&#x27;t know if it&#x27;s possible. Have you tried assign? something like &#x60;{% assign omg = site.posts %}&#x60;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/68#issuecomment-55626920) on: **Monday, September 15, 2014**

@Daniel0524 Any luck with this? The only thing I can think of trying is what @albertogg suggested.

This looks interesting, and I will try to hack something together on my own to see what can be done.
---
> from: [**Daniel0524**](https://github.com/jekyll/jekyll-help/issues/68#issuecomment-55628101) on: **Monday, September 15, 2014**

@troyswanson I just ended up using categories in lieu of a collections. 
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/68#issuecomment-55628410) on: **Monday, September 15, 2014**

Ah, sounds good! Glad you figured it out :smile: 
