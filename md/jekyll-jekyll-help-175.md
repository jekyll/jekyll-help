# [Final site problems](https://github.com/jekyll/jekyll-help/issues/175)

> state: **closed** opened by: **** on: ****

After running jekyll new in a directory, jekyll builds an example site. This site displays completely differently when using jekyll serve and when opening the html files on their own. This is a problem as the instructions say the the content of the _site folder is the final product but this does not appear to be so. 
When using jekyll serve, the site is properly formatted and styled. When opening the html files, there is no styling and the icons are very large. The CSS files have been generated though.
Is there any way to fix this?
Thanks in advance!

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/175#issuecomment-60136405) on: **Wednesday, October 22, 2014**

This is most likely an issue of how the CSS file is referenced from the default.html layout.

Inside &#x60;_includes/head.html&#x60; this line is causing you troubles: &#x60;&lt;link rel=&quot;stylesheet&quot; href=&quot;{{ &quot;/css/main.css&quot; | prepend: site.baseurl }}&quot;&gt;&#x60;.

Without the first slash, it would also work locally, e.g. &#x60;css/main.css&#x60;. However if you set &#x60;baseurl&#x60; in your config, you end up with referencing &#x60;site.tld/blogcss/main.css&#x60;.

By the way, when I run &#x60;jekyll new&#x60; I get &#x60;jekyll 2.4.0 | Error:  You must specify a path.&#x60;. Do I need to &#x60;mkdir test &amp;&amp; jekyll new test&#x60; or something? Seems so.
---
> from: [**sarvagnan**](https://github.com/jekyll/jekyll-help/issues/175#issuecomment-60153551) on: **Wednesday, October 22, 2014**

That worked!! Thank you!!

You have to give jekyll new a location to run in. &quot;jekyll new test&quot; will create a folder called test and put all its stuff in there, so &quot;mkdir test&quot; isn&#x27;t necessary.
