# [Jekyll seems to half work](https://github.com/jekyll/jekyll-help/issues/226)

> state: **closed** opened by: **** on: ****

Repo: https://github.com/gitjohndoe/mediator

Github Page: http://gitjohndoe.github.io/mediator/

Not sure what is wrong here but I have been troubleshooting all day with no results. Can anyone help and tell me what I am doing wrong?

### Comments

---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68653788) on: **Sunday, January 04, 2015**

your css is not loading - in your head section try

 /mediator/css/main.css

---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68653914) on: **Sunday, January 04, 2015**

that goes for any of your other urls to - you need the /mediator/ part. Like your logo.
---
> from: [**gitjohndoe**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68654219) on: **Sunday, January 04, 2015**

Yeah so /mediator will do the trick for css, js and images.
Any idea why the posts are not working? 
http://gitjohndoe.github.io/mediator/feature/2014/11/30/mediator_features.html the first post can not be found.
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68654961) on: **Sunday, January 04, 2015**

i think because in your _config.yml file you need your base url to include the /mediator .

Which I think you just figured out.
---
> from: [**gitjohndoe**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68655094) on: **Sunday, January 04, 2015**

baseurl: &quot;/mediator&quot;
This does fix the css and js but does not fix the posts issue. Still confused here. Really weird.
Thanks for the help by the way
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68655142) on: **Sunday, January 04, 2015**

&#x60;&#x60;&#x60;
url: &quot;http://gitjohndoe.github.io/mediator&quot; # the base hostname &amp; protocol for your site
&#x60;&#x60;&#x60;
I don&#x27;t use github, and am not 100% sure on the useage of url vs baseurl. But all your problems stem from the  fact that your site is at /mediator not just github.io. 

i think i would try just using URL and leave Baseurl blank?
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68655282) on: **Sunday, January 04, 2015**

the url to the features article is really:
&#x60;&#x60;&#x60;
http://gitjohndoe.github.io/mediator/mediator/feature/2014/11/30/mediator_features.html
&#x60;&#x60;&#x60;
the mediator bit is very confusing because it is the name of your repo and the name of the skin you are using. The url above has mediator/mediator/feature because your repo is at /mediator - so it will always be there, then the next mediator is because the Category set for that articel in the frontmatter is &#x27;mediator feature&#x27; and jekyll is doing the post url as each word in category is another &#x27;/&#x27; if that makes sense. 

---
> from: [**gitjohndoe**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68655720) on: **Sunday, January 04, 2015**

Alright, I will play around with it. At least I know now the posts are being created. Greatly appreciate the help.
---
> from: [**gitjohndoe**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68656571) on: **Sunday, January 04, 2015**

http://jekyllrb.com/docs/github-pages/#project-pages
For others to reference. Again thanks for the help.
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68657100) on: **Sunday, January 04, 2015**

no problem, good luck!

I think you need url: &quot;http://gitjohndoe.github.io/mediator&quot;
and baseurl: &quot;/mediator&quot;

but not 100% sure - I think baseurl is for local, url is for what it really is when live.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/226#issuecomment-68659540) on: **Sunday, January 04, 2015**

Gents, I would highly encourage you to not use &#x60;baseurl&#x60; and instead use [GitHub Pages Metadata](https://help.github.com/articles/repository-metadata-on-github-pages/).

A quick example: Assuming you have a CSS file called &#x60;/css/styles.css&#x60; relative to the root of your repo, you can use something like &#x60;&lt;link rel=&quot;stylesheet&quot; href=&quot;{{ site.github.url }}/css/styles.css&quot;&gt;&#x60;. The &#x60;{{ site.github.url }}&#x60; value will be replaced by the URL for your repo&#x27;s site.
