# [Jekyll-Paginate generates no pages](https://github.com/jekyll/jekyll-help/issues/149)

> state: **open** opened by: **** on: ****

I have a Jekyll instance with currently 4 blog posts. When I set paginate, the root gives me the number of posts per page (&#x60;paginate: 2&#x60;) but I always get the last 2 posts on index.html. The pagination link will show previous but the page files (&#x60;paginate_path&#x60;) are not generated. To test the pagination I have following settings.

    paginate: 2
    paginate_path: &quot;seite-:num&quot;

Even though I have a much more complex code with if statements to control pagination etc, the general idea of the template is as follows:

    {% for post in paginator.posts %}
    {% include article.html %}
    {% endfor %}

    &lt;p class=&quot;hidden&quot;&gt;Seite: {{ paginator.page }} von {{ paginator.total_pages }}&lt;/p&gt;
    &lt;ul class=&quot;pager&quot;&gt;
        {% if paginator.previous_page %}
        &lt;li class=&quot;previous&quot;&gt;
            &lt;a href=&quot;{{ paginator.previous_page_path | prepend: site.baseurl | replace: &#x27;//&#x27;, &#x27;/&#x27; }}&quot;&gt;&lt;i class=&quot;fa fa-chevron-left&quot;&gt;&lt;/i&gt; Neuer&lt;/a&gt;
        &lt;/li&gt;
        {% endif %}
        {% if paginator.next_page %}
        &lt;li class=&quot;next&quot;&gt;
            &lt;a href=&quot;{{ paginator.next_page_path | prepend: site.baseurl | replace: &#x27;//&#x27;, &#x27;/&#x27; }}&quot;&gt;&lt;i class=&quot;fa fa-chevron-right&quot;&gt;&lt;/i&gt; &amp;Auml;lter&lt;/a&gt;
        &lt;/li&gt;
        {% endif %}
    &lt;/ul&gt;

What do I do wrong?

Here are the versions I have on my system:

    $ gem list | grep jekyll
    jekyll (2.4.0, 2.3.0, 2.1.1, 2.0.3)
    jekyll-coffeescript (1.0.1, 1.0.0)
    jekyll-date-format (1.0.1, 1.0.0)
    jekyll-gist (1.1.0)
    jekyll-page-hooks (1.3.1, 1.2.0)
    jekyll-paginate (1.0.0)
    jekyll-sass-converter (1.2.1, 1.2.0, 1.0.0)
    jekyll-sitemap (0.6.0, 0.5.1, 0.4.1)
    jekyll-watch (1.1.1, 1.1.0, 1.0.0)

My code is on [bitbucket](https://bitbucket.org/emin/dokudo-blog-jekyll/).

### Comments

---
> from: [**emeQee**](https://github.com/jekyll/jekyll-help/issues/149#issuecomment-57054286) on: **Saturday, September 27, 2014**

Any insights on this one? I can&#x27;t make the pagination work with &#x60;relative_permalinks: false&#x60;
---
> from: [**corysimmons**](https://github.com/jekyll/jekyll-help/issues/149#issuecomment-77095939) on: **Tuesday, March 03, 2015**

bump?
