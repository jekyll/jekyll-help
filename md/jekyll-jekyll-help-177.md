# [Non-home pagination either doesn&#x27;t work or messes up navigation](https://github.com/jekyll/jekyll-help/issues/177)

> state: **open** opened by: **** on: ****

Consider the following url structure:

&#x60;domain/&#x60; for the homepage (static)
&#x60;domain/news/&#x60; for the list of posts (paginated)
&#x60;domain/news/page-n/&#x60; for page n of posts

To my understanding, there are two ways I can structure the &#x60;news&#x60; page: either by using &#x60;news.html&#x60; or by using &#x60;news/index.html&#x60;. The first method will render the pagination in the home layout - and return empty on the &#x60;news.html&#x60; loop.

No problem, let&#x27;s try the &#x60;news/index.html&#x60; method. Pagination now works, but every paginated page now shows up in the navigation - making it look like &#x60;news | news | news | news | news&#x60;. Not good either.

So, my question is: is there a way to provide pagination on a page other than home, without messing up the navigation?

### Comments

---
> from: [**creativedutchmen**](https://github.com/jekyll/jekyll-help/issues/177#issuecomment-61363281) on: **Saturday, November 01, 2014**

Any ideas on what could be the problem here? Thanks!
---
> from: [**TheCodeDestroyer**](https://github.com/jekyll/jekyll-help/issues/177#issuecomment-75371747) on: **Saturday, February 21, 2015**

I have the same issue... Any news on this?
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/177#issuecomment-75379433) on: **Saturday, February 21, 2015**

Pagination will only work on index.html files and with no filtered group of content like categories, tags, etc. I&#x27;m not sure what&#x27;s messing up your last example, though.
