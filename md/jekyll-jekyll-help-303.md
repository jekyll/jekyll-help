# [{{ content }} Liquid tag not showing blog posts](https://github.com/jekyll/jekyll-help/issues/303)

> state: **open** opened by: **** on: ****

Link: https://github.com/AlvSovereign/AlvSovereign.github.io

Hi all

I am a beginner using jekyll and new to web development, so please be patient with me.

I am coding my portfolio site, with a link to my blog. The fully processed portfolio site sits in _site/index.html which contains a link to the blog (fully processed link exists in _site/blog.html.

Blog.html at the root of the repo has in the front matter a layout of &quot;bl&quot;, which is defined in my layouts folder under &quot;bl.html&quot;.

&quot;bl.html&quot; is the layout I want for my blog page, which contains my includes etc.... It also contains the {{ content }} liquid tag.

If I am thinking about this the right way, my posts are being parsed (I think this is the right terminology) into &quot;post.html&quot; - which has the front matter &quot;layout: bl&quot;,  which then parsed into &quot;bl.html&quot;, and that this is so as they both have {{ content }} in each files.

Now all my posts are showing correctly (i think) within &quot;_site/(year)/(month) etc...&quot;. However, they are not ending being parsed through my &quot;blog.html&quot; file, and then visible on my webpage. How I know this is that the fully processed &quot;blog.html&quot; file does not have any of the posts in them.

What will I need to do to solve this issue?

If it helps, I use Prepros for LIveReload, and using LANYON template, http://lanyon.getpoole.com/

Thank you for your time and I really do appreciate the help you can give.

Sorry for the confusing detail.

Alvin Sovereign



### Comments

