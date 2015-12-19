# [Where are the group_by docs? :-)](https://github.com/jekyll/jekyll-help/issues/42)

> state: **closed** opened by: **** on: ****

I tried to group my posts by year by doing something like this:

&#x60;&#x60;&#x60;&#x60;
{% for year in site.posts | group_by:&quot;{{ date | date:&#x27;%Y&#x27; }}&quot; %}

&#x60;&#x60;&#x60;&#x60;

but it doesn&#x27;t work, and I&#x27;ve been unable to find docs on what you can group by or how you should write it. And: 

&#x60;&#x60;&#x60;&#x60;
{% for year in {{ site.posts | group_by:&quot;{{ date | date:&#x27;%Y&#x27; }}&quot; }} %}
## {{ year.name }}
{% for post in year.items %}
- [{{ post.title | markdownify | remove:&#x27;&lt;p&gt;&#x27; | remove:&#x27;&lt;/p&gt;&#x27; }}]({{ post.url }}) (&lt;time&gt;{{ post.date | date:&#x27;%F&#x27; }}&lt;/time&gt;){% endfor %}
{% endfor %}
&#x60;&#x60;&#x60;&#x60;

doesn&#x27;t work either, like:

&#x60;&#x60;&#x60;&#x60;
{% for year in site.posts | group_by:&quot;date | date:&#x27;%Y&#x27;&quot; %}
&#x60;&#x60;&#x60;&#x60;

and 

&#x60;&#x60;&#x60;&#x60;
{% for year in {{ site.posts | group_by:&quot;date | date:&#x27;%Y&#x27;&quot; }} %}
&#x60;&#x60;&#x60;&#x60;



### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/42#issuecomment-42772421) on: **Sunday, May 11, 2014**

The documentation is here: http://jekyllrb.com/docs/templates/

{% for year in site.posts | group_by:&quot;date&quot; %}

Which probably won&#x27;t do what you want. But group_by doesn&#x27;t take
expressions like what you&#x27;re trying to do, unfortunately.


On 11 May 2014 23:43, Giuseppe C &lt;notifications@github.com&gt; wrote:

&gt; I tried to group my posts by year by doing something like this:
&gt;
&gt; {% for year in site.posts | group_by:&quot;{{ date | date:&#x27;%Y&#x27; }}&quot; %}
&gt;
&gt;
&gt; but it doesn&#x27;t work, and I&#x27;ve been unable to find docs on what you can
&gt; group by or how you should write it. And:
&gt;
&gt; {% for year in {{ site.posts | group_by:&quot;{{ date | date:&#x27;%Y&#x27; }}&quot; }} %}
&gt; ## {{ year.name }}
&gt; {% for post in year.items %}
&gt; - [{{ post.title | markdownify | remove:&#x27;&lt;p&gt;&#x27; | remove:&#x27;&lt;/p&gt;&#x27; }}]({{ post.url }}) (&lt;time&gt;{{ post.date | date:&#x27;%F&#x27; }}&lt;/time&gt;){% endfor %}
&gt; {% endfor %}
&gt;
&gt; doesn&#x27;t work either, like:
&gt;
&gt; {% for year in site.posts | group_by:&quot;date | date:&#x27;%Y&#x27;&quot; %}
&gt;
&gt; and
&gt;
&gt; {% for year in {{ site.posts | group_by:&quot;date | date:&#x27;%Y&#x27;&quot; }} %}
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/42&gt;
&gt; .
&gt;
---
> from: [**9peppe**](https://github.com/jekyll/jekyll-help/issues/42#issuecomment-42773114) on: **Sunday, May 11, 2014**

&gt; The documentation is here: http://jekyllrb.com/docs/templates/

I saw that, I didn&#x27;t find it very clear on what you can group by. Issue jekyll/jekyll#1735 makes it somewhat clearer. 

&gt; {% for year in site.posts | group_by:&quot;date&quot; %}
&gt; 
&gt; Which probably won&#x27;t do what you want. But group_by doesn&#x27;t take
&gt; expressions like what you&#x27;re trying to do, unfortunately.

Yeah, that probably means a feature request. :/
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/42#issuecomment-42773345) on: **Sunday, May 11, 2014**

It&#x27;s clear, it&#x27;s just unfortunately succinct and technical. Not really sure
what can be done on either front, the site design doesn&#x27;t really allow for
great amounts of verbosity there as there&#x27;s not much room.


On 12 May 2014 01:10, Giuseppe C &lt;notifications@github.com&gt; wrote:

&gt; The documentation is here: http://jekyllrb.com/docs/templates/
&gt;
&gt; I saw that, I didn&#x27;t find it very clear on what you can group by. Issue
&gt; jekyll/jekyll#1735 &lt;https://github.com/jekyll/jekyll/issues/1735&gt; makes
&gt; it somewhat clearer.
&gt;
&gt; {% for year in site.posts | group_by:&quot;date&quot; %}
&gt;
&gt; Which probably won&#x27;t do what you want. But group_by doesn&#x27;t take
&gt; expressions like what you&#x27;re trying to do, unfortunately.
&gt;
&gt; Yeah, that probably means a feature request. :/
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/42#issuecomment-42773114&gt;
&gt; .
&gt;
