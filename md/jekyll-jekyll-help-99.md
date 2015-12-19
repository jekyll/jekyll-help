# [How to create a listing of categories](https://github.com/jekyll/jekyll-help/issues/99)

> state: **closed** opened by: **** on: ****

Hello,

I&#x27;m slowly getting my website together, it&#x27;s live, it&#x27;s healthy but I want to add bits and pieces. Right now I&#x27;m struggling with categories. 
I want to create a listing for my categories to keep them as a sidebar. I&#x27;ve been searching for a while how to do this and I can&#x27;t find an answer.

I&#x27;m not sure if Jekyll generates a folder for each category or how it works. I know I have my categories defined in my _config.yml file and each post has its category in the front matter.

Can somebody give me some clues or point me a resource?

Thanks!


### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/99#issuecomment-49442358) on: **Friday, July 18, 2014**

Have a look at my [categories.html](https://github.com/kleinfreund/kleinfreund.github.io/blob/dev/en/categories.html):

&#x60;&#x60;&#x60;
{% for category in site.categories %}
    {% assign posts = category[1] | where:&quot;lang&quot;,&quot;en&quot; %}

    {% if posts.size &gt; 0 %}
    &lt;h2 id=&quot;{{ category[0] }}&quot;&gt;{{ category[0] }}&lt;/h2&gt;
    &lt;ul&gt;
        {% for post in posts %}
        &lt;li&gt;
            &lt;a href=&quot;{{ post.url }}&quot;&gt;{{ post.title }}&lt;/a&gt;
        &lt;/li&gt;
        {% endfor %}
    &lt;/ul&gt;
    {% endif %}

    {% assign posts = nil %}
{% endfor %}
&#x60;&#x60;&#x60;

It lists all categories and English posts in them. However you just need a smaller bit:

&#x60;&#x60;&#x60;
&lt;ul class=&quot;category-list&quot;&gt;
{% for category in site.categories %}
    &lt;li&gt;{{ category[0] }}&lt;/li&gt;
{% endfor %}
&lt;/ul&gt;
&#x60;&#x60;&#x60;
---
> from: [**bitcolorine**](https://github.com/jekyll/jekyll-help/issues/99#issuecomment-49854484) on: **Wednesday, July 23, 2014**

Thank you very much.
I will try this. 
Also, your site comes in handy for study because I also need a bilingual site!



---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/99#issuecomment-49855311) on: **Wednesday, July 23, 2014**

@bitcolorine I&#x27;m currently changing a few bits around. When 2.1 hits on GHP, I go back to pushing the uncompiled Jekyll files. In the meantime I push my current uncompiled state to the dev branch. …if you&#x27;re interested.
---
> from: [**bitcolorine**](https://github.com/jekyll/jekyll-help/issues/99#issuecomment-49855633) on: **Wednesday, July 23, 2014**

Definitely interested!

I had been searching for Jekyll sites to study, but many have been abandoned. I&#x27;m building my first so I need to look at a lot of examples.

Thanks again.
