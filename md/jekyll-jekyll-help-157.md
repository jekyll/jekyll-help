# [Problem with array &quot;sort:&quot; liquid filter on a data file (CSV)](https://github.com/jekyll/jekyll-help/issues/157)

> state: **closed** opened by: **** on: ****

I am having *all kinds* of problems with using &#x27;sort&#x27; on data files (a CSV in this case).

Any ideas on why this isn&#x27;t doing what I&#x27;m expecting it to do? https://gist.github.com/phillipadsmith/0fdb339c911c21c56b09

Also included here for convenience:

&#x60;&#x60;&#x60;
{% assign candidates = site.data.toronto_council | sort: &#x27;sortby&#x27; %}
{% for candidate in candidates %}
{% if candidate.ward == page.wid %}
&lt;li&gt;&lt;a title=&quot;Information on Toronto ward {{ page.wid }} city council candidate {{ candidate.name_full }}&quot; href=&quot;/toronto-city-council/{{ candidate.slug }}&quot;&gt;{{ candidate.name_full }}&lt;/a&gt; {% if candidate.incumbent == &#x27;yes&#x27; %}(Incumbent){% endif %}&lt;/li&gt;
{% endif %}
{% endfor %}
&#x60;&#x60;&#x60;
It&#x27;s very odd because if I dump the json for candidates, the array of objects is sorted properly (or appears to be). But then the {% for candidate in candidates %} loop seems to loose the sort order.

(Obviously, if I sort the CSV itself, the sort is unnecessary).

The CSV is here: https://github.com/phillipadsmith/everycandidate.org/blob/master/_data/toronto_council.csv


### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/157#issuecomment-56623374) on: **Tuesday, September 23, 2014**

Ensure you are formatting and therefore defining the &#x60;page.wid&#x60; front matter properly to be interpreted later within the &#x60;if&#x60; conditional.

i.e. &#x60;wid: &#x27;44&#x27;&#x60;

i.e. not &#x60;wid: 44&#x60;


---
> from: [**phillipadsmith**](https://github.com/jekyll/jekyll-help/issues/157#issuecomment-56668818) on: **Wednesday, September 24, 2014**

Hi @jaybe-jekyll Thanks for this. That&#x27;s odd, as the if *is* working, but the sort seems off. I&#x27;ll fix up the front matter anyway to see if that helps. 

I&#x27;m not able to grab that CSV you messaged me last night (I was AFK at the time). Can you re-send or post here?

Appreciate your help. 
---
> from: [**phillipadsmith**](https://github.com/jekyll/jekyll-help/issues/157#issuecomment-56825107) on: **Thursday, September 25, 2014**

Just for those reading along... 

The issue was actually that the column I was using to sort was labelled &#x27;sortby&#x27; and that seems to be a reserved word in Liquid or something. 

When changed to &#x27;name_last_lower&#x27; it worked as expected.
