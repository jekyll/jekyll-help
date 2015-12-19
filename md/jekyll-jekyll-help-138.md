# [category in post YAML mixed with automatic categories](https://github.com/jekyll/jekyll-help/issues/138)

> state: **closed** opened by: **** on: ****

Recently I found that the categories of my site become in a mass. 

Before. Category is: 
![8](https://cloud.githubusercontent.com/assets/3502393/4028533/ae7fc8b8-2c3e-11e4-8335-49a4c7930a35.png)
Now. It become:
![9](https://cloud.githubusercontent.com/assets/3502393/4028705/6b1bd60e-2c41-11e4-8457-e41483a27944.png)

The code for this part is:

    &lt;div class=&quot;sidebar_unit&quot;&gt;
      &lt;div class=&quot;unit_container&quot;&gt;
          &lt;div class=&quot;aside_title&quot;&gt;&lt;i class=&quot;icon_sprite icon_categories&quot;&gt;&lt;/i&gt;&lt;em&gt;文章分类&lt;/em&gt;&lt;span&gt;Categories&lt;/span&gt;&lt;/div&gt;
            &lt;ul class=&quot;category_list clearfix&quot;&gt;
                {% for category in site.categories  %}
                &lt;li class=&quot;category_item&quot;&gt;&lt;a class=&quot;category_link&quot; href=&quot;{{ BASE_PATH }}{{ site.JB.categories_path }}#{{ category | first }}-ref&quot; title=&quot;&quot;&gt;{{ category | first }}&lt;/a&gt;&lt;span class=&quot;post_count&quot;&gt;({{ category | last | size }})&lt;/span&gt;&lt;/li&gt;
                {% endfor  %}
            &lt;/ul&gt;
        &lt;/div&gt;
    &lt;/div&gt;

As posts in my folders are like:
**_posts/css/2013-03-13-sticky-footer.md** (YAML - category: css)
**_posts/essay/2014-05-24-new-acgtofe-with-responsive-design.md** (YAML - category: 记录随笔)

So. I guess jekyll built category information refer to the subfolders of **_posts**. In each of my post, I wrote a YAML category. Maybe jekyll massed them up? How can I cancel jekyll&#x27;s &quot;automatic category build&quot;?

When I run my site locally, It&#x27;s OK. It just happens on github pages.




### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/138#issuecomment-53254040) on: **Monday, August 25, 2014**

As long as posts lay inside a subdirectory like &#x60;[dir]/_posts/[post].md&#x60;, &#x60;[dir]&#x60; is treated as a category as well. It&#x27;s currently not possible to disable this behavior as far as I am aware.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/138#issuecomment-53255323) on: **Monday, August 25, 2014**

@kenanpengyou, the current problem is as follows:

- The current version of GitHub&#x27; Jekyll is &#x60;2.2.0&#x60;.
  - https://pages.github.com/versions/
- The current version of Jekyll is &#x60;2.3.0&#x60;.
  - http://jekyllrb.com/news/

Version 2.2.0 of Jekyll changed the historic behavior of folders within _posts&#x27; directories as follows:

&gt; Categories in a post’s path are now respected (i.e. folders in _posts will now work properly).
(* http://jekyllrb.com/news/2014/07/29/jekyll-2-2-0-released/)

Version 2.3.0 of Jekyll changed that behavior *back* because of confusion, conflicts, and feedback such as what you described:

&gt; One change deserves special note. In #2633, subfolders inside a _posts folder were processed and added as categories to the posts. It turns out, this behaviour was unwanted by a large number of individuals, as it is a handy way to organize posts. Ultimately, we decided to revert this change in #2705, because it was a change in behaviour that was already well-established (at least since Jekyll v0.7.0), and was convenient.
(* http://jekyllrb.com/news/2014/08/11/jekyll-2-3-0-released/)

Once GitHub Pages upgrades to Jekyll 2.3.0, the behavior of Jekyll with regards to folders nested within _posts&#x27; folder will revert to its historic personality.

Until then, you could- as follows:

- Ignore and put up with the added categorization until GitHub Pages updates to Jekyll 2.3.0.
- Do not use subdirectories nested within _posts folders, to avoid the added categorization, now.

Jekyll will continue to provide and support categorization based upon folders leading up to _posts/ directories. i.e. &#x60;/llamas/sassy/_posts/why-llamas-should-not-be-trusted.md.&#x60;

---
> from: [**kenanpengyou**](https://github.com/jekyll/jekyll-help/issues/138#issuecomment-53256538) on: **Monday, August 25, 2014**

@jaybe-jekyll , Thank you very much! That&#x27;s it. Maybe I will adapt to this behavior temporarily and wait for github updating its jekyll version.

I did know the method to use paths like &#x60;[dir]/_posts/[post].md&#x60; as @kleinfreund mentioned to make &quot;automatic categories&quot;. While I need Chinese characters, which cannot be used in path, I just treated subfolders of **_posts** as a way to keep files ordered.

I didn&#x27;t expect that jekyll 2.2.0 does this and change it back in 2.3.0. It seems that I&#x27;m one of those &quot;individuals&quot; who really think it unwanted. :)   
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/138#issuecomment-53257396) on: **Monday, August 25, 2014**

You are welcome.

More information, for reference:

- https://github.com/jekyll/jekyll/pull/2633
- https://github.com/jekyll/jekyll/issues/2691

