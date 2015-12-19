# [How can I limit results to one match when two conditions match?](https://github.com/jekyll/jekyll-help/issues/247)

> state: **open** opened by: **** on: ****

I want to cycle through my pages and get all pages that have either a tag called &quot;release_notes&quot; or a tag called &quot;news.&quot; Here&#x27;s my code:

&#x60;&#x60;&#x60;
{% for page in site.pages limit: 5 %}
{% for tag in page.tags %}
{% if tag == &quot;release_notes&quot; or tag == &quot;news&quot; %}
&lt;li&gt;&lt;a href=&quot;{{ page.permalink | prepend: site.baseurl }}&quot;&gt;{{page.title}}&lt;/a&gt;&lt;/li&gt;
&lt;div class=&quot;summary&quot;&gt;{{page.summary}}&lt;/div&gt;
{% endif %}
{% endfor %}
{% endfor %}
&#x60;&#x60;&#x60;

One of my pages has front matter as follows:

&#x60;&#x60;&#x60;
---
layout: page
title: 2.0 Release Notes
permalink: /release_notes_2_0/
tags: 
- release_notes
- news
---
&#x60;&#x60;&#x60;

This page appears twice in the &#x60;for&#x60; loop. How can I make it so that the page only appears once? 

(I realize the simple solution would be to restrict the tags to one per page, or restrict the for loop to one tag per loop, but maybe it&#x27;s easier to do this than I thought.) 

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/247#issuecomment-71337214) on: **Saturday, January 24, 2015**

Hello @tomjohnson1492, could you try something like:

&#x60;&#x60;&#x60;html
    {% assign pages = site.pages limit:5 %}
    {% for p in pages  %}{% unless p.tags == release_notes || news %}
      &lt;li&gt;&lt;a href=&quot;{{ page.permalink | prepend: site.baseurl }}&quot;&gt;{{page.title}}&lt;/a&gt;&lt;/li&gt;
      &lt;div class=&quot;summary&quot;&gt;{{page.summary}}&lt;/div&gt;
    {% endunless %}{% endfor %}
&#x60;&#x60;&#x60;

and see if it works? I&#x27;m not sure, tho.
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/247#issuecomment-71416104) on: **Sunday, January 25, 2015**

Thanks for responding. I tried your suggestion but it didn&#x27;t work. I got a long list of bullet items, without any content (text) with each bullet.
