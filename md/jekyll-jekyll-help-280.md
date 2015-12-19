# [How can I increase site regeneration time given 50 capture tags?](https://github.com/jekyll/jekyll-help/issues/280)

> state: **closed** opened by: **** on: ****

I have around 50 capture tags that I&#x27;m using to populate tooltips on another site using the method I described here: [Create a help API](http://tomjohnson1492.github.io/documentation-theme-jekyll/help_api/). 

I also want to embed the tooltips inside my jekyll site. I have a page that looks like this: 

&#x60;&#x60;&#x60;
{% for page in site.tooltips %}

{% if page.id == &quot;sample1&quot; %} {% capture sample1 %} {{page.content}} {% endcapture %}{% endif %}
{% if page.id == &quot;sample2&quot; %} {% capture sample2 %} {{page.content}} {% endcapture %}{% endif %}

// repeat this about 50 times with different page.id&#x27;s...

{% endfor%}
&#x60;&#x60;&#x60;
When I implemented this code, my site&#x27;s regeneration time jumped from 1.5 seconds to 11 seconds. How can I assign all these capture tags more efficiently?



### Comments

---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/280#issuecomment-77523624) on: **Friday, March 06, 2015**

I realized that the for loops are requiring the long site build time. I moved everything into data files and now the site loads &lt; 2 seconds. 
