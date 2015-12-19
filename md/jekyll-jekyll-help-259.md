# [Do folders that hold pages really need index.html files? ](https://github.com/jekyll/jekyll-help/issues/259)

> state: **open** opened by: **** on: ****

I saw this in the Jekyll docs [here](http://jekyllrb.com/docs/pages/):

&gt;Where you put HTML files for pages depends on how you want the pages to work. There are two main ways of creating pages:
&gt;Place named HTML files for each page in your site’s root folder.
Create a folder in the site’s root for each page, and place an index.html file in each page folder.

I have a folder in my directory called &quot;content.&quot; This directory includes lots of pages, but there is no index.html. The pages render just fine. Is this documentation outdated?

### Comments

---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/259#issuecomment-72156783) on: **Thursday, January 29, 2015**

Also, the section that follows ([Named folders containing index HTML files](http://jekyllrb.com/docs/pages/#named-folders-containing-index-html-files)) also seems outdated. This section talks about &quot;clean&quot; pages. I&#x27;m not implementing the technique, and my pages aren&#x27;t muddled with file extensions. They are clean by default.


---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/259#issuecomment-72164197) on: **Thursday, January 29, 2015**

There is a difference whether you have &#x60;/blablabla/mrfreeman.html&#x60; or &#x60;/blablabla/mrfreeman/index.html&#x60;. The last one can be accessed by navigating to &#x60;/blablabla/mrfreeman&#x60;, wheras the first needs the extension.
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/259#issuecomment-72164722) on: **Thursday, January 29, 2015**

I have the former setup (sans index.html) and I can access by navigating to &#x60;/blablabla/mrfreeman&#x60;.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/259#issuecomment-72165817) on: **Thursday, January 29, 2015**

Are you hosting on GitHub Pages or on something else? One can achieve extension-less permalinks via rewrite settings in a .htaccess file on the server for example.

For static pages without any further settings, you will need the extension or a directory with a index.html inside.
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/259#issuecomment-73106975) on: **Thursday, February 05, 2015**

&gt; Is this documentation outdated?

This documentation is opinionated. You are free to design your URLs as you want.


---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/259#issuecomment-73107126) on: **Thursday, February 05, 2015**

@tomjohnson1492 So you do *not* need &#x60;index.html&#x60; files to make Jekyll work.
