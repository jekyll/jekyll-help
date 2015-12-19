# [Best way to make an &quot;about $author&quot; partial](https://github.com/jekyll/jekyll-help/issues/65)

> state: **closed** opened by: **** on: ****

Hello.

I&#x27;m converting a blog to Jekyll and I want there to be an &quot;About Author&quot; box at the end of each post, containing the author&#x27;s name, picture and description and I want this box to be consistent for each author. The easiest way to approach this is to provide the author&#x27;s information in every posts written by them, but then there will be several copies of the same information across the articles written by that author.

I, then, figured I could create a plugin to create those partials, but, as far as I can see, Generator plugins are only able to create full pages, not just partials.

What, then, would be the best approach to do that?

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/65#issuecomment-44993598) on: **Tuesday, June 03, 2014**

You can add authors to a &#x60;&#x60;&#x60;_data/authors.yaml&#x60;&#x60;&#x60; file and reference them in the front matter of each post. Might look like this: 
&#x60;&#x60;&#x60;
{% assign article_author = site.data.authors | where: &#x27;author&#x27;, &#x27;author&#x27; %}
{% for author in article_author  %}
    {{author.name}}
    {{author.bio}}
{% endfor %}
&#x60;&#x60;&#x60;
edit: you&#x27;d probably need to capture the post&#x27;s author beforehand with capture page.author. sorry, I left that out initially
---
> from: [**romariorios**](https://github.com/jekyll/jekyll-help/issues/65#issuecomment-44996195) on: **Tuesday, June 03, 2014**

I suppose I should replace that first &#x27;author&#x27; in the first line by the name of the author, right?
How can I find further documentation on that filter? I can&#x27;t find anything in the liquid docs.

Thanks.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/65#issuecomment-44998023) on: **Tuesday, June 03, 2014**

authors.yaml:
&#x60;&#x60;&#x60;
-
 name: Author Name
 email: email@gmail.com
 author_id: author_identifier
&#x60;&#x60;&#x60;
post.md
&#x60;&#x60;&#x60;
---
title: Post
layout: article
author: author_identifier
---
copy
&#x60;&#x60;&#x60;

template:
&#x60;&#x60;&#x60;

{% assign the_author = site.data.authors | where: &#x27;author_id&#x27;, {{page.author}} %}
{% for author in the_author %}
 name: {{author.name}}&lt;br /&gt;
 email: {{author.email}}
{% endfor %}
&#x60;&#x60;&#x60;

I tend to use twitter as the author_id, and I tend to capture the page author before I use it in assign, but this will work.
---
> from: [**romariorios**](https://github.com/jekyll/jekyll-help/issues/65#issuecomment-45001108) on: **Tuesday, June 03, 2014**

It worked nicely. Thanks.

Just one more (noob) question: is that &quot;the_author&quot; var available in all templates including that part, or it&#x27;s just available locally? I need the author data in more than one place and don&#x27;t want to write all that boilerplate code everywhere I use that data, so I thought it would be easier if I could create a partial that grabs the data and saves them in other variables.

EDIT: Nevermind, it works.

Thanks for all the help.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/65#issuecomment-45003849) on: **Tuesday, June 03, 2014**

Glad you got it working. You can put that in a partial and call the partial with the author as a parameter: http://jekyllrb.com/docs/templates/#includes
