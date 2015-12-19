# [Non-post content image in RSS feed](https://github.com/jekyll/jekyll-help/issues/150)

> state: **closed** opened by: **** on: ****

I have a &#x27;feature&#x27; image on my site that isn&#x27;t part of the &#x27;main&#x27; post. That is to say, if you put &#x60;{{ post.content }}&#x60;, the feature image doesn&#x27;t come up. It&#x27;s called &#x60;{{ post.image.feature }}&#x60;

On my home page I have it come up like so:

    {% if post.image.feature %}
    &lt;a href=&quot;{{site.url}}{{post.url}}&quot;&gt; &lt;img src=&quot;{{site.url}}/images/{{ post.image.feature }}&quot;&gt;&lt;/a&gt;
    {% endif %}

This works great. It&#x27;s in an &#x60;if&#x60; statement because not every post has such an image.

I&#x27;m trying to get this to work in my &#x60;feed.xml&#x60; but a simple copy/paste isn&#x27;t doing the job - it just seems to break the feed to the point where it no longer renders posts.

This is my &#x60;feed.xml&#x60;:

    ---
    layout: none
    ---
    &lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;
    &lt;feed xmlns=&quot;http://www.w3.org/2005/Atom&quot; xml:lang=&quot;en&quot;&gt;
    &lt;title type=&quot;text&quot;&gt;{{ site.title }}&lt;/title&gt;
    &lt;subtitle type=&quot;text&quot;&gt;{{ site.tagline }}&lt;/subtitle&gt;
    &lt;generator uri=&quot;https://github.com/mojombo/jekyll&quot;&gt;Jekyll&lt;/generator&gt;
    &lt;link rel=&quot;self&quot; type=&quot;application/atom+xml&quot; href=&quot;{{ site.url }}/feed.xml&quot; /&gt;
    &lt;link rel=&quot;alternate&quot; type=&quot;text/html&quot; href=&quot;{{ site.url }}&quot; /&gt;
    &lt;updated&gt;{{ site.time | date_to_xmlschema }}&lt;/updated&gt;
    &lt;id&gt;{{ site.url }}/&lt;/id&gt;
    &lt;author&gt;
    &lt;name&gt;{{ site.owner.name }}&lt;/name&gt;
    &lt;uri&gt;{{ site.url }}/&lt;/uri&gt;
    &lt;email&gt;{{ site.owner.email }}&lt;/email&gt;
    &lt;/author&gt;
    {% for post in site.posts limit:20 %}

    &lt;entry&gt;
    &lt;title type=&quot;html&quot;&gt;&lt;![CDATA[{{ post.title | cdata_escape }}]]&gt;&lt;/title&gt;
    &lt;link rel=&quot;alternate&quot; type=&quot;text/html&quot; href=&quot;{{ site.url }}{{ post.url }}&quot; /&gt;
    &lt;id&gt;{{ site.url }}{{ post.id }}&lt;/id&gt;
    {% if post.modified %}&lt;updated&gt;{{ post.modified | to_xmlschema }}T00:00:00-00:00&lt;/updated&gt;
    &lt;published&gt;{{ post.date | date_to_xmlschema }}&lt;/published&gt;
    {% else %}&lt;published&gt;{{ post.date | date_to_xmlschema }}&lt;/published&gt;
    &lt;updated&gt;{{ post.date | date_to_xmlschema }}&lt;/updated&gt;{% endif %}
    &lt;author&gt;
    &lt;name&gt;{{ site.owner.name }}&lt;/name&gt;
    &lt;uri&gt;{{ site.url }}&lt;/uri&gt;
    &lt;email&gt;{{ site.owner.email }}&lt;/email&gt;
    &lt;/author&gt;
    &lt;content type=&quot;html&quot;&gt;{{ post.content | xml_escape }}
    &amp;lt;p&amp;gt;&amp;lt;a href=&amp;quot;{{ site.url }}{{ post.url }}&amp;quot;&amp;gt;{{ post.title }}&amp;lt;/a&amp;gt; was originally published by {{ site.owner.name }} at &amp;lt;a href=&amp;quot;{{ site.url }}&amp;quot;&amp;gt;{{ site.title }}&amp;lt;/a&amp;gt; on  {{ post.date | date: &quot;%B %d, %Y&quot; }}.&amp;lt;/p&amp;gt;&lt;/content&gt;
    &lt;/entry&gt;
    {% endfor %}
     &lt;/feed&gt;

I want the image below the title and date, but just above the contents. What do I need to do to make this happen and where do I need to put the code?

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/150#issuecomment-55621285) on: **Monday, September 15, 2014**

First of all, I&#x27;d recommend using [CDATA](http://www.w3schools.com/xml/xml_cdata.asp) in your &#x60;content&#x60; tag so that you don&#x27;t have to escape all of your &#x60;&lt;&gt;&quot;&#x60; characters.

Secondly, inside the CDATA area, you should be fine to add an &#x60;img&#x60; tag right above the &#x60;{{ post.content }}&#x60; stuff. You shouldn&#x27;t have to use the &#x60;xml_escape&#x60; filter on that either, as long as you&#x27;re using the CDATA enclosure.

PS: You might want to look at the &#x60;updated&#x60; tag, also. It doesn&#x27;t look like it will work properly if you have &#x60;modified&#x60; in your front-matter.
---
> from: [**benoliver999**](https://github.com/jekyll/jekyll-help/issues/150#issuecomment-55713346) on: **Tuesday, September 16, 2014**

Troy, this is exactly it. Thanks for your help.
