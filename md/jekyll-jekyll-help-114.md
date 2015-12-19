# [Get Tags by Category](https://github.com/jekyll/jekyll-help/issues/114)

> state: **closed** opened by: **** on: ****

I am trying to make a page that will list Tags for a given Category - so I will specify the Category, and then want to get all Tags, with the corresponding posts.

I can get all Tags site wide with posts, but I can&#x27;t figure out a way to filter it by category.

I want to do ths same thing as this guy on SO:

http://stackoverflow.com/questions/23521931/list-tags-within-a-specific-category-in-jekyll




### Comments

---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/114#issuecomment-51693573) on: **Saturday, August 09, 2014**

I think I figured it out after about 1000 tries - seems to work for all kinds of filters like category or other front matter variables - like &quot;type&quot; so I can have type: article or type: video and this seems to get tags from just one of them if I put that in the where part.

    {% assign sorted_tags = site.tags | sort %}
    {% for tag in sorted_tags %}
    {% assign zz = tag[1] | where: &quot;category&quot;, &quot;Photoshop&quot; | sort %}
    {% if zz != empty %}

    &lt;li&gt;&lt;span class=&quot;tag&quot;&gt;{{ tag[0] }}&lt;/span&gt;
    &lt;ul&gt;
      {% for p in zz %}
      &lt;li&gt;&lt;a href=&quot;{{ p.url }}&quot;&gt;{{ p.title }}&lt;/a&gt;&lt;/li&gt;
      {% endfor %}
     &lt;/ul&gt;
     &lt;/li&gt;
     {% endif %}

     {% endfor %}

zz is just something to use to filter above the first tag[0] since all it seems to have is the tag itself, so you can&#x27;t filter anything else with it. tag[1] has all of the other stuff.

Initially I was using if zz != null or if zz != &quot;&quot; but neither of them worked.
