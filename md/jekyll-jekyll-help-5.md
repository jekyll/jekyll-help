# [How do I create &quot;Edit this page&quot; links in Jekyll?](https://github.com/jekyll/jekyll-help/issues/5)

> state: **closed** opened by: **** on: ****

A few Jekyll docs pages I&#x27;ve visited have a handy &quot;Edit&quot; link that takes me to that page&#x27;s file on GitHub, so I can edit it directly or make a pull request.

Example 1: Top right, &quot;Improve this doc&quot;
http://ionicframework.com/docs/api/directive/ionTabs/

Example 2: Top right, &quot;Edit&quot;
http://sendgrid.com/docs/API_Reference/index.html

Is there a default way of achieving this in Jekyll?

### Comments

---
> from: [**akoeplinger**](https://github.com/jekyll/jekyll-help/issues/5#issuecomment-39033862) on: **Sunday, March 30, 2014**

If you&#x27;re hosting the site on GitHub Pages, they inject the source repository into a variable called site.github.repository_url which you can use in combination with page.path to wire up the URL on github.com.

Example: &#x60;&#x60;&#x60;&lt;a href=&quot;{{site.github.repository_url}}/blob/gh-pages/{{page.path}}&quot;&gt;Edit this page&lt;/a&gt;&#x60;&#x60;&#x60;

(Note: on a collection document, you&#x27;ll need to use &#x60;{{page.relative_path}}&#x60; instead of &#x60;{{page.path}}&#x60; to get the desired behavior, see http://jekyllrb.com/docs/collections/#documents)
---
> from: [**erlend-sh**](https://github.com/jekyll/jekyll-help/issues/5#issuecomment-39038242) on: **Sunday, March 30, 2014**

Example and everything, thanks a lot!
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/5#issuecomment-39040524) on: **Sunday, March 30, 2014**

@erlend-sh if it fits your use-case you send the link directly to the [prose](https://github.com/prose/prose) content editor. I put my repo info in my config file, so for me the link looks like this:
&#x60;&#x60;&#x60;
http://prose.io/#{{site.repo}}/edit/{{site.branch}}/{{ page.path }}
&#x60;&#x60;&#x60;



---
> from: [**garethredfern**](https://github.com/jekyll/jekyll-help/issues/5#issuecomment-64114129) on: **Sunday, November 23, 2014**

Hi there I wonder if you can help me, I am trying to use @akoeplinger example but when using &#x60;{{ page.path }}&#x60; the word &#x60;repo&#x60; gets prepended to the path do I need to set the repo name in the config and then do something like:

&#x60;{{site.github.repository_url}}{{page.path}}&#x60;

Thanks in advance.
---
> from: [**akoeplinger**](https://github.com/jekyll/jekyll-help/issues/5#issuecomment-64119011) on: **Sunday, November 23, 2014**

@garethredfern I had a look at your site, and it appears that collection pages (or documents as they are called in Jekyll docs) are handled a bit different than regular pages.

E.g. &#x60;{{page.path}}&#x60; in a collection results in &#x60;&lt;a href=&quot;/home/alexander/dev/statamicthemes.github.io/_articles/2014-07-31-creating-a-simple-contact-form-with-raven.md&quot; ...&#x60; (-&gt; the full path to the file), while in a regular page it only results in the relative path from the site root.

For a collection document, you&#x27;ll need to use &#x60;{{page.relative_path}}&#x60; to get the desired behavior (see http://jekyllrb.com/docs/collections/#documents)
---
> from: [**garethredfern**](https://github.com/jekyll/jekyll-help/issues/5#issuecomment-64119403) on: **Sunday, November 23, 2014**

Ahh thank you @akoeplinger that worked perfect :+1: for anyone else that might want the full code I used here you go:

&#x60;&lt;a href=&quot;{{ site.github.repository_url }}/tree/master/{{ page.relative_path }}&quot;&gt;Edit&lt;/a&gt;&#x60;
