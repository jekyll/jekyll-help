# [Loop trough subfolder of _posts?](https://github.com/jekyll/jekyll-help/issues/293)

> state: **open** opened by: **** on: ****

I can&#x27;t seem to figur this out. Is it possible to loop trough posts and pages in a subfolder of &quot;_posts&quot;?

Like so: 

&#x60; {% for page in posts.staticpages %} {{ page.title }} {% endfor %} &#x60;

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/293#issuecomment-97617771) on: **Wednesday, April 29, 2015**

I believe you&#x27;re looking for &#x60;&#x60;&#x60;site.categories.category_name&#x60;&#x60;&#x60; I know that works for where you have:
&#x60;&#x60;&#x60;
category_name
  _posts
    2015-04-29-post-name.md
&#x60;&#x60;&#x60;
but not if the folder is within the _posts folder
---
> from: [**baasdesign**](https://github.com/jekyll/jekyll-help/issues/293#issuecomment-97692277) on: **Thursday, April 30, 2015**

That&#x27;s the thing, I want to be able to loop trough markdown files in a subfolder of the posts folder. 
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/293#issuecomment-97780144) on: **Thursday, April 30, 2015**

Well it&#x27;s a clunky solution, but if you have or make something in your front matter that&#x27;s consistent you can use the &#x27;where&#x27; filter. Sorry I don&#x27;t have a better solution!
---
> from: [**jamesonzimmer**](https://github.com/jekyll/jekyll-help/issues/293#issuecomment-98011782) on: **Thursday, April 30, 2015**

You could put the categories folders within the posts folder, like:

_posts
   category-a
      2015-01-01-a-post.md
   category-b
      2016-01-02-another-post.md

Maybe see https://github.com/mmistakes/so-simple-theme for example of using
categories like this...  Still clunky but at least keeps stuff within posts
folder

Well it&#x27;s a clunky solution, but if you have or make something in your
front matter that&#x27;s consistent you can use the &#x27;where&#x27; filter. Sorry I
don&#x27;t have a better solution!

—
Reply to this email directly or view it on GitHub
&lt;https://github.com/jekyll/jekyll-help/issues/293#issuecomment-97780144&gt;.

---
> from: [**baasdesign**](https://github.com/jekyll/jekyll-help/issues/293#issuecomment-98220247) on: **Friday, May 01, 2015**

That&#x27;s how I solved it now yes :) 
But its good to see that most people solve it like this, atleast I know there isn&#x27;t a better solution. 

Thanks
