# [Jekyll Liquid interface](https://github.com/jekyll/jekyll-help/issues/300)

> state: **open** opened by: **** on: ****

Dear Jekyll-help: 

I have spent frustrating hours trying to do what I think is simple. I have a YFM variable called &#x60;navbar_links&#x60; which is a list of strings. Each of these strings is a key in a file in &#x60;_data&#x60; named &#x60;navbarlinks.yml&#x60;. I would like to retrieve the value of each of the key in the list &#x60;navbar_links&#x60;. 

Any help would be appreciated. 

Thanks.  

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/300#issuecomment-100599616) on: **Sunday, May 10, 2015**

How is the front matter (your navbar links)  formatted? 
---
> from: [**knsam**](https://github.com/jekyll/jekyll-help/issues/300#issuecomment-100601432) on: **Sunday, May 10, 2015**

Do you mean the file &#x60;navbarlinks.yml&#x60;?

~~~
- name: &quot;Home&quot;
  url: &quot;/&quot;

- name: &quot;About&quot;
  url: &quot;/about/&quot;
~~~

and I include a &#x60;nav.html&#x60; which is a partial with the following code: 

~~~
&lt;span id=&quot;nav&quot;&gt; 
&lt;ul&gt;
  {% for navbar_link in page.navbar_links %}
      &lt;span&gt;&lt;a href={{ site.data.navbarlinks | where:&quot;name&quot;,&quot;navbar_link&quot; }}&gt;{{ navbar_link }}&lt;/a&gt;&lt;/span&gt;
  {% endfor %}
&lt;/ul&gt;
&lt;/span&gt; 
~~~

I will just say that I have tried many variants that might make sense for the yml file and the above html file. 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/300#issuecomment-100604865) on: **Sunday, May 10, 2015**

So you have a &#x60;navbarlinks.yml&#x60; inside your &#x60;_data&#x60; directory?

The procedure with that would look like this:

&#x60;&#x60;&#x60;liquid
{% for navbar_link in site.data.navbarlinks %}
    &lt;a href=&quot;{{ navbar_link.url }}&quot;&gt;{{ navbar_link.name }}&lt;/a&gt;
{% endfor %}
&#x60;&#x60;&#x60;
---
> from: [**knsam**](https://github.com/jekyll/jekyll-help/issues/300#issuecomment-100604937) on: **Sunday, May 10, 2015**

@kleinfreund But I don&#x27;t want do that for all the entries in navbarlinks.yml, but only for those specified in the YFM variable &#x60;navbar_links&#x60;. This variable takes a list of strings at the moment. 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/300#issuecomment-100617040) on: **Sunday, May 10, 2015**

How is the front matter for this variable structured? It always helps to include a minimal example.
---
> from: [**knsam**](https://github.com/jekyll/jekyll-help/issues/300#issuecomment-100675741) on: **Sunday, May 10, 2015**

@kleinfreund I pointed out in my OP that it was a list of strings like so: 

    ---
    navbar_links: [&quot;Home&quot;]
    ---
