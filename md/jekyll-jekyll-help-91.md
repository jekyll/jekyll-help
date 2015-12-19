# [Iterating through multiple _data items](https://github.com/jekyll/jekyll-help/issues/91)

> state: **closed** opened by: **** on: ****

If I have the following in a document&#x27;s front matter:

&#x60;&#x60;&#x60;
authors: &quot;author1, author5, author7&quot;
&#x60;&#x60;&#x60;

and a _data file &#x60;&#x60;&#x60;authors.yaml&#x60;&#x60;&#x60; with, say 20 authors:

&#x60;&#x60;&#x60;
- 
 identifier: author1
 name: Author Name
 bio: &quot;Author Bio&quot;
- 
 identifier: author2
 name: Author2 Name
 bio: &quot;Author2 Bio&quot;

etc....
&#x60;&#x60;&#x60;
I&#x27;d like pull each author&#x27;s name and bio in my template. What&#x27;s the best way to go about this?



### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/91#issuecomment-48088269) on: **Saturday, July 05, 2014**

The following in an approach to accomplishing:
*Based upon latest &#x60;jekyll new testsite&#x60; (v.2.1.0)*

&gt; ./_data/authors.yml

&#x60;&#x60;&#x60; yaml
- id: &#x27;author1&#x27;
  name: &#x27;Author Name&#x27;
  bio: &#x27;Author Bio&#x27;

- id: &#x27;author2&#x27;
  name: &#x27;Author2 Name&#x27;
  bio: &#x27;Author2 Bio&#x27;

- id: &#x27;author3&#x27;
  name: &#x27;Author Name&#x27;
  bio: &#x27;Author Bio&#x27;

- id: &#x27;author4&#x27;
  name: &#x27;Author2 Name&#x27;
  bio: &#x27;Author2 Bio&#x27;
&#x60;&#x60;&#x60;

&gt; ./_layouts/post.html

&#x60;&#x60;&#x60; html
---
layout: default
---
&lt;div class=&quot;post&quot;&gt;

  &lt;header class=&quot;post-header&quot;&gt;
    &lt;h1&gt;{{ page.title }}&lt;/h1&gt;
    &lt;p class=&quot;meta&quot;&gt;{{ page.date | date: &quot;%b %-d, %Y&quot; }}{% if page.author %} • {{ page.author }}{% endif %}{% if page.meta %} • {{ page.meta }}{% endif %}&lt;/p&gt;
    &lt;p class=&quot;meta&quot;&gt;{% include authors.html %}&lt;/p&gt;
  &lt;/header&gt;

  &lt;article class=&quot;post-content&quot;&gt;
  {{ content }}
  &lt;/article&gt;

&lt;/div&gt;
&#x60;&#x60;&#x60;

*Key addition is: &#x60;&lt;p class=&quot;meta&quot;&gt;{% include authors.html %}&lt;/p&gt;&#x60;*

&gt; ./_posts/2014-07-05-welcome-to-jekyll.markdown

&#x60;&#x60;&#x60;
---
layout: post
title:  &quot;Welcome to Jekyll!&quot;
date:   2014-07-05 09:07:48
categories: jekyll update
authors: 
  - author2
  - author4
---

You&#x27;ll find this post in your &#x60;_posts&#x60; directory - edit this post and re-build (or run with the &#x60;-w&#x60; switch) to see your changes!
To add new posts, simply add a file in the &#x60;_posts&#x60; directory that follows the convention: YYYY-MM-DD-name-of-post.ext.

Jekyll also offers powerful support for code snippets:

{% highlight ruby %}
def print_hi(name)
  puts &quot;Hi, #{name}&quot;
end
print_hi(&#x27;Tom&#x27;)
#=&gt; prints &#x27;Hi, Tom&#x27; to STDOUT.
{% endhighlight %}

Check out the [Jekyll docs][jekyll] for more info on how to get the most out of Jekyll. File all bugs/feature requests at [Jekyll&#x27;s GitHub repo][jekyll-gh].

[jekyll-gh]: https://github.com/jekyll/jekyll
[jekyll]:    http://jekyllrb.com
&#x60;&#x60;&#x60;

*Key additions:*
&#x60;&#x60;&#x60; yaml
authors: 
  - author2
  - author4
&#x60;&#x60;&#x60;

&gt; ./_includes/authors.html

&#x60;&#x60;&#x60; liquid
{% for author in page.authors %}
  {% for siteauthor in site.data.authors %}
    {% if author == siteauthor.id %}
      &lt;p&gt;Author: {{ siteauthor }}&lt;/p&gt;
    {% endif %}
  {% endfor %}
{% endfor %}
&#x60;&#x60;&#x60;

&gt; Expected/Example Output:

&#x60;&#x60;&#x60;
...

Welcome to Jekyll!
Jul 5, 2014
Author: {&quot;id&quot;=&gt;&quot;author2&quot;, &quot;name&quot;=&gt;&quot;Author2 Name&quot;, &quot;bio&quot;=&gt;&quot;Author2 Bio&quot;}
Author: {&quot;id&quot;=&gt;&quot;author4&quot;, &quot;name&quot;=&gt;&quot;Author2 Name&quot;, &quot;bio&quot;=&gt;&quot;Author2 Bio&quot;}
You’ll find this post in your _posts directory - edit this post and re-build (or run with the -w switch) to see your changes! To add new posts, simply add a file in the _posts directory that follows the convention: YYYY-MM-DD-name-of-post.ext.

...
&#x60;&#x60;&#x60;
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/91#issuecomment-48093049) on: **Saturday, July 05, 2014**

Ah, of course! So simple to nest the loop, now that you say it. Much appreciated, @jaybe-jekyll. 

Here&#x27;s what I ended up with for anyone else who finds this post:

&#x60;&#x60;&#x60;
{% for article_authors in page.authors %}
  {% assign author_data = site.data.authors | where: &#x27;twitter&#x27;, article_authors %}
  {% for author in author_data %}
      {{ author.name }}
  {% endfor %}
{% endfor %}
&#x60;&#x60;&#x60;
