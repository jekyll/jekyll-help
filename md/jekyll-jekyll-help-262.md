# [Controlling the length/size of post excerpts on index.html](https://github.com/jekyll/jekyll-help/issues/262)

> state: **closed** opened by: **** on: ****

Hi

At the moment I use post excerpts on &#x60;index.html&#x60; on my Jekyll site http://sandeepmurthy.is, but I am not completely happy with the apparently arbitrary way in the lengths of the excerpts seem to vary.  For some posts, like one with an inline image in the first paragraph, the excerpts seem to be much longer than others where the first paragraph is pure text.  For posts where there is a single large image, the excerpt includes all of that, rather than just showing part of the image.

Is there a way to control the size of an excerpt?

Sandeep

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/262#issuecomment-72311624) on: **Saturday, January 31, 2015**

Yes, there is. The default excerpt seperator are two newlines (&#x60;\n\n&#x60;). You can change this in your config.
---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/262#issuecomment-72318477) on: **Saturday, January 31, 2015**

Thanks, it works.
---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/262#issuecomment-72319096) on: **Saturday, January 31, 2015**

Hi,

just a final question, is there a similar property for controlling the space between posts.  At the moment I&#x27;d like my posts to be spaced a bit more closely in &#x60;index.html&#x60; but there seems to be no way of controlling this via the code:

    &lt;div class=&quot;posts&quot;&gt;
    {% for post in paginator.posts %}
        &lt;div class=&quot;post&quot;&gt;
        &lt;h1 class=&quot;post-title&quot;&gt;
          &lt;a href=&quot;{{ post.url }}&quot;&gt;{{post.title}}&lt;/a&gt;
        &lt;/h1&gt;
        &lt;span class=&quot;post-date&quot;&gt;{{ post.date | date_to_string }}&lt;/span&gt;
        {{ post.excerpt }}&lt;a href=&quot;{{ site.url }}{{ post.url }}&quot;&gt;....&lt;/a&gt;
        &lt;/div&gt;
    {% endfor %}
    &lt;/div&gt;

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/262#issuecomment-72321095) on: **Saturday, January 31, 2015**

CSS.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/262#issuecomment-72322266) on: **Saturday, January 31, 2015**

&gt; CSS.

These things are controlled via CSS, correct. For example, you could change the &#x60;line-height&#x60; on the &#x60;html&#x60; element. I usually prefer more room between my lines and have something like &#x60;line-height: 1.5;&#x60;.
---
> from: [**sr-murthy**](https://github.com/jekyll/jekyll-help/issues/262#issuecomment-72322285) on: **Saturday, January 31, 2015**

OK, thanks.
