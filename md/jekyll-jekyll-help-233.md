# [Disqus comments are not showing up](https://github.com/jekyll/jekyll-help/issues/233)

> state: **closed** opened by: **** on: ****

Hi

I&#x27;m having problems with the include files for the Disqus comments (and the counter) on my posts.  I&#x27;ve enabled comments in &#x60;_config.yml&#x60; as follows:

    # Comments
     comments:
       provider: disqus
     disqus: 
       short_name: codefree

My &#x60;post.html&#x60; layout file, with the two &#x60;include&#x60; statements at the end, looks like this:
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
          &lt;a href=&quot;{{site.url}}{{ post.url }}#disqus_thread&quot;&gt;
            {{ post.title }}
            &lt;small&gt;{{ post.date | date_to_string }}&lt;/small&gt;
          &lt;/a&gt;
        &lt;/h3&gt;
      &lt;/li&gt;
    {% endfor %}
    &lt;/ul&gt;
    &lt;/div&gt;
    {% include disqus_comments.html %}
    {% include disqus_comments_counter.html %}
&#x60;&#x60;&#x60;
These Disqus code in the include files is standard.  I am using the Hyde theme.  I&#x27;ve checked 
What could be going wrong here?  At the moment these are the only two problems I need to fix before adding more content.  Your help would be much appreciated.

Thanks


### Comments

---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/233#issuecomment-69178984) on: **Thursday, January 08, 2015**

Have you restarted Jekyll? Changes in config.yaml are only picked up on start.
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/233#issuecomment-69179257) on: **Thursday, January 08, 2015**

Beyond that, perhaps you can include a link to your repo? 

Neither Jekyll nor the Hyde theme include Disqus comments, so you are not really giving us much to work on ;) 
---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/233#issuecomment-69186242) on: **Thursday, January 08, 2015**

Hi

Sorry about the info., I tried again and it worked.  I think it was the &#x60;disqus.short_name&#x60; field that I didn&#x27;t set the correct value for, but it does work now.  Here&#x27;s the link to the Disqus thread for one of my posts

https://sr-murthy.github.io/os/x,/shells/2015/01/08/Bashed/#disqus_thread

Thanks.

P. S. THe 
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/233#issuecomment-69186440) on: **Thursday, January 08, 2015**

:cool: 

Thanks for confirming.
