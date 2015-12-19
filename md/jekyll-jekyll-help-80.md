# [post.html not doing anything other than displaying post content (can&#x27;t include header, footer, etc)](https://github.com/jekyll/jekyll-help/issues/80)

> state: **closed** opened by: **** on: ****

So I am using https://github.com/centresource/generator-playbook to kickstart my jekyll blog and everytihng is fine except for the fact that my post.html doesn&#x27;t seem to be applying the layout it should be (default.html).

In fact, anything i add to post.html gets ignored, all it does is display the post {{content}} and nothing else. 
Am I doing something wrong ? There also seems to be no css file that works on it, yet all of this works everywhere else.

Thank you,

Here is the markup : 

&#x60;&#x60;&#x60;html
---
layout: default
---

&lt;div class=&quot;post&quot;&gt;
  &lt;h1 class=&quot;post-title&quot;&gt;{{ page.title }}&lt;/h1&gt;
  &lt;span class=&quot;post-date&quot;&gt;{{ page.date | date_to_string }}&lt;/span&gt;
  {{ content }}
&lt;/div&gt;

&lt;div class=&quot;related&quot;&gt;
  &lt;h2&gt;Related Posts&lt;/h2&gt;
  &lt;ul class=&quot;related-posts&quot;&gt;
    {% for post in site.related_posts limit:3 %}
      &lt;li&gt;
        &lt;h3&gt;
          &lt;a href=&quot;{{ site.baseurl }}{{ post.url }}&quot;&gt;
            {{ post.title }}
            &lt;small&gt;{{ post.date | date_to_string }}&lt;/small&gt;
          &lt;/a&gt;
        &lt;/h3&gt;
      &lt;/li&gt;
    {% endfor %}
  &lt;/ul&gt;
&lt;/div&gt;
&#x60;&#x60;&#x60;

### Comments

---
> from: [**HarrisRobin**](https://github.com/jekyll/jekyll-help/issues/80#issuecomment-46758335) on: **Saturday, June 21, 2014**

Apparently, I have to include the layout: default in every markdown post i make. I assumed it would automatically apply if i put it in post.html.

didn&#x27;t know, sorry.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/80#issuecomment-46758781) on: **Saturday, June 21, 2014**

Yep, Jekyll tries to stay away from being too &quot;automagical&quot; so as not to confuse. Glad you got it figured out!

If you&#x27;re using Jekyll 2.x, check out [Front-Matter Defaults](http://jekyllrb.com/docs/configuration/#frontmatter-defaults).
