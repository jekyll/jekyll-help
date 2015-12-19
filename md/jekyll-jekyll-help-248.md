# [How can I limit the for loop to 2 items for pages (not posts)?](https://github.com/jekyll/jekyll-help/issues/248)

> state: **closed** opened by: **** on: ****

I have a for loop to get all pages with a certain tag:

&#x60;&#x60;&#x60;
 {% for page in site.pages %}
 {% for tag in page.tags %}
 {% if tag == &quot;layout&quot; %}
&lt;li&gt;&lt;a href=&quot;{{ page.permalink | prepend: site.baseurl }}&quot;&gt;{{page.title}}&lt;/a&gt;&lt;/li&gt;
 {% endif %}
 {% endfor %}
 {% endfor %}
&#x60;&#x60;&#x60;

How can I limit the results to just 2 pages? If I put the &#x60;limit&#x60; attribute in the first for statement, it limits the pages iterated through, not the number of pages with those tags iterated through. But if I put the limit on the second for loop, it doesn&#x27;t work.

### Comments

---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/248#issuecomment-71542189) on: **Monday, January 26, 2015**

Someone on [StackOverflow answered my question here](http://stackoverflow.com/questions/28156020/in-jekyll-how-can-i-get-all-pages-with-a-specific-tag-and-limit-the-results-to). 
