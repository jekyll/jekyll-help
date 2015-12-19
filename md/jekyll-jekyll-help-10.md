# [Output different on GH-Pages vs local build](https://github.com/jekyll/jekyll-help/issues/10)

> state: **closed** opened by: **** on: ****

Wondering if any of you have seen this sort of thing?

I&#x27;m running jekyll (1.5.1), liquid (2.5.5), which are the same version GH Pages is running (https://pages.github.com/versions/), yet, I&#x27;m getting an odd difference between what I&#x27;m seeing on my local build and the Pages version. 

Specifically, the order of items in my navigation is different locally than it is on Pages. This appears to be the only difference.

Jekyll  sorts pages by the filename, so I&#x27;ve given my pages files name beginning with numbers, e.g. &#x27;01-home&#x27;, &#x27;02-about&#x27;. This order is respected on my local build, but not on the Pages build. In fact, there doesn&#x27;t appear to be a specific order at all on the Pages build. Hopefully I&#x27;m just overlooking something!

Here&#x27;s the site: http://linnullmann.pubspringdev1.com/
https://github.com/sonnetmedia/linnullmann

Here&#x27;s my code (it&#x27;s a bi-lingual site):

&#x60;&#x60;&#x60;
 {% if page.url contains &#x27;/en/&#x27; %}
      &lt;a href=&quot;/en/&quot;&gt;&lt;li class=&quot;sidebar-nav-item&quot;&gt;home&lt;/li&gt;&lt;/a&gt;
      {% for page in site.pages %}
        {% if page.permalink contains &#x27;/en/&#x27;  %}
          {% if page.show_in_menu == true  %}
            &lt;a href=&quot;{{page.url}}&quot;&gt;
              &lt;li id=&quot;{{page.url}}&quot; class=&quot;sidebar-nav-item{% if page.url contains /page.url/  %} active{% endif %}&quot;&gt;{{page.title}}&lt;/li&gt;
            &lt;/a&gt;  
          {% endif %}
        {% endif %}
      {% endfor %}
    {% else %}{% comment %}If EN isn&#x27;t in URL serve up the &#x27;unless&#x27; case{% endcomment %}
      &lt;a href=&quot;/&quot;&gt;&lt;li class=&quot;sidebar-nav-item&quot;&gt;hjem&lt;/li&gt;&lt;/a&gt;
      {% for page in site.pages %}
        {% unless page.permalink contains &#x27;/en/&#x27;  %}
          {% if page.show_in_menu == true  %}
            &lt;a href=&quot;{{page.url}}&quot;&gt;
              &lt;li id=&quot;{{page.url}}&quot; class=&quot;sidebar-nav-item{% if page.url contains /page.url/  %} active{% endif %}&quot;&gt;{{page.title}}&lt;/li&gt;
            &lt;/a&gt;  
          {% endif %}
        {% endunless %}
      {% endfor %}
    {% endif %}
&#x60;&#x60;&#x60;

LOCAL
![screen shot 2014-04-03 at 9 16 49 am](https://cloud.githubusercontent.com/assets/115347/2603706/54932c3c-bb32-11e3-8d17-74c367a2770a.jpg)

PAGES
![screen shot 2014-04-03 at 9 18 23 am](https://cloud.githubusercontent.com/assets/115347/2603721/7bc4e0b6-bb32-11e3-8097-9b4e1cb6ad09.jpg)



### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39462425) on: **Thursday, April 03, 2014**

Very odd. I&#x27;m seeing a different order myself!

![screen shot 2014-04-03 at 9 46 15 am](https://cloud.githubusercontent.com/assets/1420926/2604870/cb992a14-bb3e-11e3-9103-096ed31fc907.png) 
![screen shot 2014-04-03 at 9 50 43 am](https://cloud.githubusercontent.com/assets/1420926/2604937/77bb5d9e-bb3f-11e3-8126-3a7d3e47dac2.png)

For the record, the order should be:

1. home / hjem
2. about / om forfatteren
3. contact / kontact
4. books / bøker
5. interviews / intervjuer

I don&#x27;t know what could be causing this weird sorting issue. @benbalter, maybe something with GitHub&#x27;s flavor of Ruby?

In the meantime, there are a couple of things I noticed that you should probably take a look at:

Your &#x60;&lt;li&gt;&#x60; tags are living inside your &#x60;&lt;a&gt;&#x60; tags. &#x60;li&#x60;s are block elements, whereas &#x60;a&#x60;s are span elements. The rules of HTML say the &#x60;a&#x60; tag should live inside the &#x60;li&#x60;s.

Also, you are creating a loop and using &#x60;page&#x60; as the iterating variable. I believe this would overwrite the &#x60;page&#x60; variable set by Jekyll (which contains all of the data for the page like &#x60;url&#x60;, &#x60;permalink&#x60;, &#x60;title&#x60;, etc). Not sure if it&#x27;s actually causing a problem, just something to be aware of.

Finally, all of your nav elements are showing up with a class &#x60;active&#x60;.

It looks like you&#x27;ve got some good stuff set up in the &#x60;_data&#x60;, maybe set up another data file for nav and include permalinks for English and Norwegian within each element of the nav? This would probably alleviate the sorting issue, too, since you won&#x27;t be relying on file system traversal to do your sorting.

Just some thoughts. Let me know what you think.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39464142) on: **Thursday, April 03, 2014**

Thanks, Troy. Much appreciated.

In HTML5 you can indeed wrap links around block level elements and it makes it easier to add padding for fat-finger on devices.

Good point about using page as the iterating variable. I&#x27;ve used that in other places so that I could easily reuse partials, and hadn&#x27;t noticed a problem, but will keep a close eye on it. 

I think I will resort to using a data file for the nav. Will at least keep things cleaner. So, with that—from the perspective of overcoming the issue—I&#x27;m good, but it sure is an odd one. Should I close the issue?



---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39466214) on: **Thursday, April 03, 2014**

Negative, we should keep it open. It still seems like a weird thing that should be investigated.

Not to side track but I&#x27;d still nest the &#x60;&lt;a&gt;&#x60; tag inside the &#x60;&lt;li&gt;&#x60;. According to the [validator](http://validator.w3.org/check?uri=http%3A%2F%2Flinnullmann.pubspringdev1.com%2F&amp;charset=%28detect+automatically%29&amp;doctype=Inline&amp;group=0), it isn&#x27;t quite kosher. You might be right that you can wrap &#x60;&lt;a&gt;&#x60; tags around block level elements, but in this context, an &#x60;&lt;a&gt;&#x60; tag cannot be the direct child of a &#x60;&lt;ul&gt;&#x60; tag.

Use &#x60;display:block;&#x60; on the &#x60;&lt;a&gt;&#x60; to give it the full width and to retain the touch target that you&#x27;re after. The reason I&#x27;m bringing it up is because when you inspect the element with developer tools (in Chrome), it&#x27;s not drawing the box around &#x60;&lt;a&gt;&#x60; element properly. This might cause an actual usability problem in other browsers.

Here&#x27;s a screenshot of what I mean:
![screen shot 2014-04-03 at 10 27 10 am](https://cloud.githubusercontent.com/assets/1420926/2605473/b400692a-bb44-11e3-8229-35ed5c90aa39.png)
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39466721) on: **Thursday, April 03, 2014**

You mean you&#x27;d have to wrap it around the entire block (the UL) to be effective. Good point. 
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39467425) on: **Thursday, April 03, 2014**

&#x60;&#x60;&#x60;html
&lt;ul class=&quot;nav&quot;&gt;
  &lt;li class=&quot;sidebar-nav-item&quot;&gt;&lt;a href=&quot;/&quot;&gt;Home&lt;/a&gt;&lt;/li&gt;
  &lt;!-- more li elements here --&gt;
&lt;/ul&gt;
&#x60;&#x60;&#x60;

&#x60;&#x60;&#x60;css
ul.nav li {
  text-align: center;
  text-transform: lowercase;
  padding: 0.35em 0;
}
ul.nav li a {
  display: block;
}
&#x60;&#x60;&#x60;

Something like this oughta handle everything just fine. :beers: 
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39477602) on: **Thursday, April 03, 2014**

I switched out to using a data file to feed my navigation, but let me know if there&#x27;s something I can help to troubleshoot the original issue. Thanks @troyswanson !
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39485482) on: **Thursday, April 03, 2014**

Any time, @budparr.

@parkr - Any thoughts on what might be causing the original issue?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39487427) on: **Thursday, April 03, 2014**

Do we explicitly call &#x60;sort&#x60; on that list? If not, then it&#x27;s up the FS and the way &#x60;Dir.glob&#x60; works on each system.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39487812) on: **Thursday, April 03, 2014**

Isn&#x27;t the default to sort by file name? These are pages and I didn&#x27;t know there was, practically speaking, any other way to sort them. But to answer your question, no.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39488674) on: **Thursday, April 03, 2014**

As you all know, I&#x27;m no Ruby developer. I think I found the [code](https://github.com/jekyll/jekyll/blob/df8458275de4dc3d0f9b92c5247ff20832d1cc8e/lib/jekyll/site.rb#L131-L161) that grabs the files and converts them to entries (pages or posts), though.

It looks like there is a &#x60;sort_by&#x60; method happening on the &#x60;pages&#x60; variable at the very end.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39489680) on: **Thursday, April 03, 2014**

That&#x27;s funny because there&#x27;s no date object or anything else pages could sort by*. I&#x27;ve seen some tricks out there, but I&#x27;ve not had any luck with anything other than filename. But maybe tonight I&#x27;ll fork that repo and see if using sort_by changes anything.

*Collections, I&#x27;m hoping, will fix that - I&#x27;m looking forward to the alpha release.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39495789) on: **Thursday, April 03, 2014**

I’ve previously [experienced a mismatch](https://github.com/ndarville/ndarville.github.io/issues/7), because my installed gems differed between development and production. Is there a way for you to create an isolated Ruby environment to see whether that would change the output?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39496061) on: **Thursday, April 03, 2014**

ProTip :tm: Always be using the &#x60;github-pages&#x60; gem when developing locally for Pages.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39496084) on: **Thursday, April 03, 2014**

@ndarville I do believe that there is something. I can&#x27;t remember if @parkr did that.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39496413) on: **Thursday, April 03, 2014**

Oh, there it you go! &#x60;github-pages&#x60; gem :tada: 
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39500540) on: **Thursday, April 03, 2014**

I am using the &#x60;&#x60;&#x60;github-pages&#x60;&#x60;&#x60; gem, so that wouldn&#x27;t explain this problem.
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39583383) on: **Friday, April 04, 2014**

@budparr Add an &#x60;order&#x60; integer variable to the front-matter of the pages you want to sort. Then sort them by &#x60;order&#x60;.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-39585531) on: **Friday, April 04, 2014**

Thanks @penibelst - for some reason I thought that wouldn&#x27;t work, but I&#x27;ll try it today. Much appreciated.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/10#issuecomment-40401738) on: **Monday, April 14, 2014**

This issue seems to be throughly discussed and resolved. Closing.
