# [Jekyll Build site doesn&#x27;t seems to be right](https://github.com/jekyll/jekyll-help/issues/267)

> state: **open** opened by: **** on: ****

I&#x27;m testing Jekyll, created a site, it seems to have some sample content and it looks great when I run jekyll serve.

![screen shot 2015-02-05 at 7 42 34 pm](https://cloud.githubusercontent.com/assets/1278556/6069897/4851bfd0-ad6f-11e4-9dc9-dd624c1fae4e.png)

Then, when I do jekyll build and I open _site/index.html, the site looks terrible this (no errors on build):

![screen shot 2015-02-05 at 7 42 10 pm](https://cloud.githubusercontent.com/assets/1278556/6069904/4f79f3f4-ad6f-11e4-9883-1f573d28a2bc.png)
![screen shot 2015-02-05 at 7 42 51 pm](https://cloud.githubusercontent.com/assets/1278556/6069906/524c8f1a-ad6f-11e4-88f5-b08937290caa.png)


Any ideas why this is happening?



### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/267#issuecomment-73226423) on: **Friday, February 06, 2015**

Yeah, check the URLs in the source. They are root relative, not file relative. Thus previewing the site template with the local files without a local server doesn&#x27;t work. 
---
> from: [**kyrsideris**](https://github.com/jekyll/jekyll-help/issues/267#issuecomment-88921582) on: **Thursday, April 02, 2015**

I had exactly the same problem and &#x60;site.url&#x60; or &#x60;site.baseurl&#x60; inside &#x60;_config.yml&#x60; were looking sane. What I did to overcome the problem was, as @kleinfreund suggested, to &#x60;prepend: site.url&#x60; or/and &#x60;site.baseurl&#x60; whenever was missing.
The main damage to your css comes from &#x60;_includes/head.html&#x60; which should be changed from:
&#x60;&#x60;&#x60;html
&lt;link rel=&quot;stylesheet&quot; href=&quot;{{ &quot;/css/main.css&quot; | prepend: site.baseurl }}&quot;&gt;
&#x60;&#x60;&#x60;
to:
&#x60;&#x60;&#x60;html
&lt;link rel=&quot;stylesheet&quot; href=&quot;{{ &quot;/css/main.css&quot; | prepend: site.baseurl | prepend: site.url }}&quot;&gt;
&#x60;&#x60;&#x60;
Hope that helps!
