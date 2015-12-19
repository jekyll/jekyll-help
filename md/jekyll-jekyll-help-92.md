# [Jekyll 2.1.0 search plugin that supports Collections](https://github.com/jekyll/jekyll-help/issues/92)

> state: **open** opened by: **** on: ****

Has anyone found a good search plugin for Jekyll 2.1.0 that will support Collections?

Looks like [jekyll-lunr-js-search](https://github.com/slashdotdash/jekyll-lunr-js-search/issues/28) wont be supporting Collections.



### Comments

---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48097836) on: **Saturday, July 05, 2014**

I don&#x27;t like tooting my own horn, but …

Not sure if it will be of any help, but I [wrote this](https://github.com/mhulse/jquery-bigglesworth) for my Jekyll site. It uses a [customizable JSON file](https://github.com/mhulse/jquery-bigglesworth/blob/gh-pages/demo/search.json) for the search results.

For example, here&#x27;s a basic Jekyll JSON file:

&#x60;/search.json&#x60;  
&#x60;&#x60;&#x60;text
---
---

[
  {% for post in site.posts %}
    {% include post.json %}{% if forloop.last %}{% else %},{% endif %}
  {% endfor %}
]
&#x60;&#x60;&#x60;

&#x60;/includes/post.json&#x60;  
&#x60;&#x60;&#x60;text
{
  &quot;title&quot;    : &quot;{% if post.external_site %}{{ post.external_site | xml_escape }}: {% endif %}{{ post.title | xml_escape }}&quot;,
  &quot;category&quot; : [{% for category in post.categories %}&quot;{{ category }}&quot;{% if forloop.last %}{% else %}, {% endif %}{% endfor %}],
  &quot;tags&quot;     : [{% for tag in post.tags %}&quot;{{ tag }}&quot;{% if forloop.last %}{% else %}, {% endif %}{% endfor %}],
  &quot;uri&quot;     : &quot;{% if post.external_url %}{{ post.external_url }}{% else %}{{ post.url }}{% endif %}&quot;,
  &quot;date&quot;     : {
    &quot;day&quot;   : &quot;{{ post.date | date: &quot;%d&quot; }}&quot;,
    &quot;month&quot; : &quot;{{ post.date | date: &quot;%B&quot; }}&quot;,
    &quot;year&quot;  : &quot;{{ post.date | date: &quot;%Y&quot; }}&quot;
  }
}
&#x60;&#x60;&#x60;

If interested, I could provide a more complete Jekyll example.

I setup the JS to recurse through nested keys, but I can&#x27;t guarantee how speedy it will be if you go crazy in that department (i.e., collections … which, I have yet to use/test myself, so no guarantees).

I wrote the above code for personal use. I&#x27;m sure there&#x27;s way better stuff out there. Personally, I could not find anything that had the features I wanted (nested keys) plus was simple to setup (lunr is awesome, but had more moving parts than I personally cared to maintain).
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48098070) on: **Saturday, July 05, 2014**

This one is pretty simple but it works: http://mathayward.com/jekyll-search/
And if you need something more like filtering with a search component I&#x27;ve used [list.js](http://listjs.com/) to do that and it&#x27;s really simple. 
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48101223) on: **Saturday, July 05, 2014**

@budparr I just tested [mathaywarduk/jekyll-search](https://github.com/mathaywarduk/jekyll-search/). It currently doesn&#x27;t support Collections.

@mhulse &quot;I&#x27;m sure there&#x27;s way better stuff out there.&quot; If your jekyll plugin could be updated to support Collections it would be the only one out there. :)
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48101628) on: **Saturday, July 05, 2014**

@AJ-Acevedo I haven&#x27;t needed collections so I&#x27;m not sure what it would take.

With that said, seems like you should be able to generate a JSON file however you want. Aren&#x27;t collections just objects to be iterated over? If you can create a JSON file for your collection, then it should work. I guess I don&#x27;t see the limitation. :confused: 

Eventually I may use collections … I just haven&#x27;t had a need to use &#x27;em to date. :+1:

I&#x27;d be curious to know what the limitation is? What&#x27;s so complicated about collections that none of these plugins work?
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48101874) on: **Saturday, July 05, 2014**

@mhulse All of the plugins search pages and posts. The pages from collections are not being searched.

Here is an example project using collections: [JekyllKB](https://github.com/AJAlabs/jekyllkb) the _kb directory is a collection.
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48102184) on: **Saturday, July 05, 2014**

@AJ-Acevedo I think I see.

I guess my thinking is: If collections can be iterated over, like pages/posts, then you should be able to build a JSON file, using liquid tags, so that a plugin can use it for searching.

Then again, I suppose I need to play with collections before I make any assumptions! :laughing: 

Good luck in your quest!
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48102277) on: **Saturday, July 05, 2014**

@mhulse is right. Just change the JSON file to whatever you want to search through, including collections. It works, I&#x27;ve done it.
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48103077) on: **Saturday, July 05, 2014**

I took a stab at hacking on the [mathaywarduk/jekyll-search json file](https://github.com/mathaywarduk/jekyll-search/blob/master/feeds/feed.json)

Couldn&#x27;t get it working after a few attempts. Any ideas? Here is what I added

&#x60;&#x60;&#x60;json
{% for page in site.collections %}
    {% if page.layout != &#x27;none&#x27; and page.layout != &#x27;none&#x27; and page.title != null and page.title != empty and page.search_omit != true %}
        {% if forloop.index &gt; 1 %},{%endif%}{
            &quot;title&quot;: {{page.title | json }},
            &quot;content&quot;: {{page.content | strip_html | strip_newlines | json}},
            &quot;link&quot;: &quot;{{ site.url }}{{ page.url | replace: &#x27;index.html&#x27;, &#x27;&#x27; }}&quot;,
            &quot;date&quot;: {{ page.date | json }},
            &quot;excerpt&quot;: {{ page.description | json }}
        }
    {%endif%}
{% endfor %}
&#x60;&#x60;&#x60;
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-48109820) on: **Sunday, July 06, 2014**

Try checking your JSON with this: http://jsonlint.com if it passes as valid it should work.
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-51695666) on: **Saturday, August 09, 2014**

Anyone come across any working solutions? 
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-53511779) on: **Tuesday, August 26, 2014**

There is now a bounty up for this feature:

[![Bountysource](https://www.bountysource.com/badge/issue?issue_id=2074457)](https://www.bountysource.com/issues/2074457-new-jekyll-collections?utm_source=2074457&amp;utm_medium=shield&amp;utm_campaign=ISSUE_BADGE)
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-53512330) on: **Tuesday, August 26, 2014**

Correct me if I&#x27;m wrong, but isn&#x27;t your goal to have a &quot;Search&quot; plugin recurse nested keys?

That&#x27;s all that collections are, right? Just multi-dimensional arrays.

That was the problem I found with the Jekyll search plugins when I needed a &quot;search mechanism&quot; on my site several months ago (and the Lunr one was just a bit too many moving parts, so I did not check if it could recurse nested keys).

That&#x27;s why I wrote my own (simple) search plugin. I wanted to itterate/search across nested json objects.

@AJ-Acevedo, that&#x27;s the trick though. You need to find a JS search plugin that supports nested data structures. Most of the ones I&#x27;ve seen only look at the top level (not sure about Lunr search).
---
> from: [**iloveip**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-68846414) on: **Tuesday, January 06, 2015**

@AJ-Acevedo Did you find any solution for this? I&#x27;m also looking for a search plugin which powers Jekyll Collections now.
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-69507701) on: **Sunday, January 11, 2015**

Nope there is still a $65 bounty for this

[![Bountysource](https://www.bountysource.com/badge/issue?issue_id=2074457)](https://www.bountysource.com/issues/2074457-new-jekyll-collections?utm_source=2074457&amp;utm_medium=shield&amp;utm_campaign=ISSUE_BADGE)
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-69511829) on: **Sunday, January 11, 2015**

@iloveip and @AJ-Acevedo, why don&#x27;t one of you post an example of the JSON output you want to search? Again, most search plugins I&#x27;ve seen just loop over json to search, they don&#x27;t care about Jekyll&#x27;s &quot;Collections&quot;.

**To clarify: It&#x27;s all about the structure of your JSON and whether or not the search plugin supports nested JSON structure.**

**Suggestion:** Rather than worry about the search plugin supporting &quot;Collections&quot;, worry about the plugin supporting nested structures.

With that said, I personally would like to see an example JSON output that you would like to have searched. Can one of you post an example? **I want to see the output file … that means, no Jekyll tags.** I&#x27;m positive I could get my plugin to work if I had an example JSON file to loop over.

This is not rocket science.
---
> from: [**kojoru**](https://github.com/jekyll/jekyll-help/issues/92#issuecomment-81555994) on: **Monday, March 16, 2015**

Just letting you know that we&#x27;ve successfully merged collection support into &#x60;jekyll-junr-js-search&#x60;. Try it and see.
