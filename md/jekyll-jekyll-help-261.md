# [Paginator items not displaying properly](https://github.com/jekyll/jekyll-help/issues/261)

> state: **open** opened by: **** on: ****

Hi

At the moment the paginator items on &#x60;index.html&#x60; does not display properly, it looks like this

![hyde-paginator](https://cloud.githubusercontent.com/assets/9358070/5983082/caec6ec0-a8c1-11e4-99cb-b4017aee0221.png)

I don&#x27;t know why it is stacking them vertically, the paginator code is unchanged from the hyde source https://github.com/poole/hyde (and so is the relevant CSS code in poole.css):

    &lt;div class=&quot;pagination&quot;&gt;
    {% if paginator.next_page %}
    &lt;a class=&quot;pagination-item older&quot; href=&quot;{{ site.baseurl }}page{{paginator.next_page}}&quot;&gt;Older&lt;/a&gt;
    {% else %}
    &lt;span class=&quot;pagination-item older&quot;&gt;Older&lt;/span&gt;
    {% endif %}
    {% if paginator.previous_page %}
      {% if paginator.page == 2 %}
        &lt;a class=&quot;pagination-item newer&quot; href=&quot;{{ site.baseurl }}&quot;&gt;Newer&lt;/a&gt;
      {% else %}
        &lt;a class=&quot;pagination-item newer&quot; href=&quot;{{ site.baseurl }}page{{paginator.previous_page}}&quot;&gt;Newer&lt;/a&gt;
      {% endif %}
    {% else %}
      &lt;span class=&quot;pagination-item newer&quot;&gt;Newer&lt;/span&gt;
    {% endif %}
    &lt;/div&gt;

### Comments

---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/261#issuecomment-72440807) on: **Monday, February 02, 2015**

This is clearly a question for the Hyde author, as neither that paginator code nor the associated css exists in Jekyll&#x27;s default template, and I see you&#x27;ve already created an issue there: https://github.com/poole/hyde/issues/80
