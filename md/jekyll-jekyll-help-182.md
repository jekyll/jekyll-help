# [How does one list pages in a directory](https://github.com/jekyll/jekyll-help/issues/182)

> state: **open** opened by: **** on: ****

I have a couple of folders with documents.
  
    dir1\page1.md
    dir1\page2.md
    dir1\page3.md
    dir1\dir2\page4.md
    dir1\dir2\page5.md
    dir1\dir2\page6.md
    dir3\page7.md
    dir3\page8.md
    dir3\page9.md

For each folder, I&#x27;d created an &#x60;index.html&#x60; in each directory, including &#x60;dir1&#x60;, &#x60;dir1\dir2&#x60; and &#x60;dir3&#x60; with some custom content.  However, I&#x27;d like to generate list to each page inside the folder (but not subdirectories) similar to 

    &lt;h2&gt;dir1 specific header&lt;/h2&gt;
    &lt;p&gt;dir1specific content. bla bla bla.&lt;/p&gt;
    &lt;h3&gt;&lt;a href=&quot;page1.html&quot;Page1&lt;/a&gt;&lt;/h3&gt;
    &lt;p&gt;excerpt from Page1&lt;/p&gt;
    &lt;h3&gt;&lt;a href=&quot;page2.html&quot;&gt;Page2&lt;/a&gt;&lt;/h3&gt;
    &lt;p&gt;excerpt from Page2&lt;/p&gt;
    &lt;!-- more --&gt;

I added this to &#x60;index.html&#x60; in dir1

    &lt;h2&gt;dir1 specific header&lt;/h2&gt;
    &lt;p&gt;dir1specific content. bla bla bla.&lt;/p&gt;
    {% for page in site.pages %}
    &lt;h3&gt;&lt;a title=&quot;{{ page.title }}&quot; href=&quot;{{ page.url | prepend: site.baseurl }}&quot;&gt;{{ page.title }}&lt;/a&gt;&lt;/h3&gt;
    &lt;p&gt;{{page.excerpt}}&lt;/p&gt;
    {% endfor %}     

However, it shows all the pages from &#x60;page1.md&#x60; to &#x60;page9.md&#x60; rather than &#x60;page1&#x60; to &#x60;page3&#x60;. Is there a built-in way to limit the pages being listed to only have those in the same folder?


### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/182#issuecomment-61438016) on: **Sunday, November 02, 2014**

You might want to try breaking those pages into a spiffy Jekyll feature called :sparkles: [collections](http://jekyllrb.com/docs/collections/) :sparkles:

The source is closed, but for a website that I just redid, I broke out all of the legal pages into a &#x60;_legal&#x60; collection, and used a loop to iterate over them:

&#x60;&#x60;&#x60;html
&lt;ul&gt;
{% for page in site.collections.legal.docs %}
  &lt;li&gt;&lt;a href=&quot;{{ page.url }}&quot;&gt;{{ page.title }}&lt;/a&gt;&lt;/li&gt;
{% endfor %}
&lt;/ul&gt;
&#x60;&#x60;&#x60;

Without knowing much more about the relationship between your directories and pages, though, it will be tough to make specific recommendations for you. Can you give some more detail?
---
> from: [**bloudraak**](https://github.com/jekyll/jekyll-help/issues/182#issuecomment-61439199) on: **Sunday, November 02, 2014**

Thanks mate.

The pages do have references to each other. For example, page7.md may reference page4.md and vice versa. This could easily be written in ASP.NET, Python, Ruby or Java, but the content rarely changes (except for initial part).  The benefit of having a static website is that it can be uploaded to a CDN and I can get rather impressive performance. And then there is the thing that I don&#x27;t have to worry about security (aka administration site, CGI etc), or runtime requirements (OS and frameworks which all require patching and maintenance). 

For a moment, assume its a portfolio showing images with metadata.  So you&#x27;d have pages for each image, other pages for the location of the photos, and yet more pages for the keywords.  They are then interlinked such that when you view an image, you can select its location and see what photos where taken at that location.  

Or you may have a site describing events, people and places and link them together where it matters. 

You get the idea.  So each &quot;record kind&quot; is in its own directory, with references to pages in other directories.  What&#x27;s missing is the index pages. 




---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/182#issuecomment-61439521) on: **Sunday, November 02, 2014**

Dude, you didn&#x27;t include php! :wink: By the way, you don&#x27;t need to convince me the merits of Jekyllizing sites. The benefits are obvious. :+1: 

That said, it seems that you&#x27;ve got a pretty complicated setup from the lens of Jekyll, however. I don&#x27;t have enough time to demonstrate exactly what you ought to do (maybe I&#x27;ll create an example site like this one day...), but I&#x27;ll bet you can get by pretty far by using collections and including as much metadata in the front-matter as you can.

For instance, you could include the &quot;location&quot; front-matter variable, and then create an include module that will loop through other documents in the collection and display those that have the same value for the location. However, it might be a little tricky to create an index page for each location, based solely on its presence in the front-matter.

You could create an index of indexes, though, but you&#x27;d have to create a pretty smart looper since your collection won&#x27;t be organized by location. Lemme think on this one a bit. It&#x27;s a good problem to tackle!
---
> from: [**bloudraak**](https://github.com/jekyll/jekyll-help/issues/182#issuecomment-61444695) on: **Sunday, November 02, 2014**

Do you have an example perhaps of generating a site from JSON files which contains internal references? 
---
> from: [**MrHobbie**](https://github.com/jekyll/jekyll-help/issues/182#issuecomment-65709997) on: **Thursday, December 04, 2014**

The answer re collections was spot on for my issue.  That bit of code should be added to the documentation. Thanks!
