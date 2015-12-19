# [Creating static pages from markdown and evaluating liquid tags.](https://github.com/jekyll/jekyll-help/issues/51)

> state: **closed** opened by: **** on: ****

So I&#x27;d like to be able to create static pages from a markdown file in the same way I&#x27;d make a blog post from a markdown file, but I can&#x27;t figure out a way to get jekyll to render liquid content in any markdown file that isn&#x27;t in  _posts/

Is this even possible? I mean I suppose I could put my page in posts and handle the case with something hacky, but I&#x27;d really like to be able to generate my site/about page from liquid tags laden markdown.

So far I&#x27;ve got this. It evaluates the markdown content but totally ignores the liquid tags.

_includes/about.md
&#x60;&#x60;&#x60;
---
title: about
---
La La La
{%some liquid tag%}
&#x60;&#x60;&#x60;

and in about.html
&#x60;&#x60;&#x60;
---
layout: layout
---

{% capture about %}{% include about.md %}{% endcapture %}
{{ about  | markdownify }}
&#x60;&#x60;&#x60;

My output in the site/about page is just:

La La La

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-43755928) on: **Wednesday, May 21, 2014**

Make sure you include &#x60;{{ content }}&#x60; in your layout. That will output the converted and rendered version of any page which specifies it as its layout.
---
> from: [**Daniel0524**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-43757074) on: **Wednesday, May 21, 2014**

Well, that&#x27;s the issue, right?

I have an about.html file sitting in jekyll&#x27;s root directory which itself has a layout _layouts/about_layout 

I can include my about.md file in about.html but then when the layout evaluates it the liquid tags have been removed. I guess I don&#x27;t understand how I should be structuring this to achieve my goal. 

Right now I&#x27;ve got:
_includes/about.md being included in about.html (a page) whose layout is _layouts/about layout 
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-43759074) on: **Wednesday, May 21, 2014**

Front matter, in this case- &quot;layout: &quot;, in an included file, will not be parsed as you are supposing.

The main index file should possess the &quot;layout: &quot; specification, i.e.:

&gt; ./about/index.html

&#x60;&#x60;&#x60; bash
---
title: About
layout: about
---

## How about this?

{% include some-about-include.include %}
&#x60;&#x60;&#x60;

&gt;  ./_layouts/about.html

&#x60;&#x60;&#x60; bash
---
layout: sub-layout-if-desired
---

&lt;div class=&quot;about-class&quot;&gt;
... {{ content }}
&lt;/div&gt;
&#x60;&#x60;&#x60;

&gt; ./_includes/some-about-include.include

&#x60;&#x60;&#x60; bash
{% for bout in site.about %}
  {{ bout.title }}
{% endfor %}

## Some Markdown will work

- But you will need to filter through &#x60;mardownify&#x60;
  - (of course)
&#x60;&#x60;&#x60;

---
> from: [**Daniel0524**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-43759679) on: **Wednesday, May 21, 2014**

But that&#x27;s not the issue I&#x27;m having. My current setup works for evaluating markdown in the include but it does not evaluate any of the liquid tags in the markdown which is what I want.

The markdownify filter just throws away all my liquid tags.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-43761840) on: **Wednesday, May 21, 2014**

Do you wish for for the liquid code to be evaluated and produce output based on the evaluation?

Or, are you wishing for the liquid &#x60;code&#x60; itself to be (re)produced and output/shown?

Perhaps restate your goal, from the top, succinctly.

Liquid within included files can be, and is, evaluated and produces evaluated output, as expected.

Example:

&gt; ./_includes/posts.include

&#x60;&#x60;&#x60; bash
{% for post in site.posts %}
  {{ post.title }}
{% endfor %}
&#x60;&#x60;&#x60;

&gt; ./_layouts/default.html

&#x60;&#x60;&#x60; bash
&lt;html&gt;
  &lt;head&gt;...&lt;/head&gt;
  &lt;body&gt;
    {{ content }}
  &lt;/body&gt;
&lt;/html&gt;
&#x60;&#x60;&#x60;

&gt; ./about/index.html

&#x60;&#x60;&#x60; bash
---
title: &quot;My About Page&quot;
layout: default
---

## This is my About Page

- Here is a list of posts:

{% include posts.include %}

### Have a nice day
&#x60;&#x60;&#x60;

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-43787401) on: **Wednesday, May 21, 2014**

An example repository would also be very helpful.
---
> from: [**Daniel0524**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-43787884) on: **Wednesday, May 21, 2014**

Sorry for the lack of clarity.

I think I figured out what was going on. Thanks for the help and great examples!
---
> from: [**duanefields**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-76470267) on: **Friday, February 27, 2015**

Can you explain what the resolution was? I have having the same question....
---
> from: [**greenify**](https://github.com/jekyll/jekyll-help/issues/51#issuecomment-76543341) on: **Saturday, February 28, 2015**

Hi @duanefields,

I recently encountered a similar problem when I wanted to include projects which are independent jekyll pages on [a summary page](http://biojs.net/gsoc/2015/).

My ideas were
* &#x60;include_relative&#x60; (problem: one [can&#x27;t include pages with a YAML front matter](https://github.com/jekyll/jekyll/issues/3476))
* picking the page from all pages (problem: liquid tags are not evaluated)

As I needed to get this working for github pages (=&#x60;without plugins&#x60;), my workaround is a combination of both and looks horrible - I would be very happy if this gets fixed in a future release and there is an easier way to include a jekyll page in another jekyll page :)

Basically it loops over all pages, picks the one I want to include, uses their filePath to call &#x60;include_relative&#x60; and then strips the YAML front matter.

&#x60;&#x60;&#x60;
{% for p in include.pages %}
 {% if p.gsoc == include.gsocYear %}

 {% capture inpath %}{{page.gsocYear}}/projects/{{ p.name }}{% endcapture %}

 {% capture content %}{% include_relative {{inpath}} %}{% endcapture %}

&lt;!-- probably the ugliest way to remove the YAML front matter --&gt;
{% assign lines = content | newline_to_br | split: &quot;&lt;br /&gt;&quot; %}

{% assign newContent = &quot;&quot; %}
{% for l in lines %}
    {% if forloop.index &gt;= 6 %}
		{% assign newContent  = newContent | append: l %}
	{% endif %}
{% endfor %}
{{ newContent | markdownify }}
&#x60;&#x60;&#x60;

source: https://github.com/biojs/biojs.net/blob/gh-pages/gsoc/includes/ideas.md
page: http://biojs.net/gsoc/2015/
