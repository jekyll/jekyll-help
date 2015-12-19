# [How to Change Your Homepage from Recent Posts to About?](https://github.com/jekyll/jekyll-help/issues/289)

> state: **open** opened by: **** on: ****

I have been messing around with the [Long-Haul](http://brianmaierjr.com/long-haul/) theme for Jekyll lately and I want to make &quot;About&quot; the homepage, not the standard &quot;Recent Posts&quot; page (which I want to instead put in a Blog section). Does anyone know how to do this? Any help is greatly appreciated. :)

The repository for Long-Haul is [here](https://github.com/brianmaierjr/long-haul).

### Comments

---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-91985116) on: **Saturday, April 11, 2015**

The index file in the root directory gets rendered as the homepage. Maybe just take the existing code out of that file and put it into a new page called &quot;Recent Posts.&quot; Then move the content from your About page into that index page (and delete the About page).
---
> from: [**DanielRoberts95**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92098494) on: **Sunday, April 12, 2015**

I&#x27;ve tried doing that but the problem is that there are two index.html files by default that are used to create the Recent Posts page. There&#x27;s one in the root directory and one in the _sites folder. The one in the _sites folder is what actually shows up when I run the command Jekyll serve, but the one in the root folder is used to take what you put in the _posts folder onto the page. I&#x27;m new to Jekyll so I&#x27;m not really sure why this is, but if you want you can check out the Long-Haul repository and see what I&#x27;m talking about [here](https://github.com/brianmaierjr/long-haul).
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92100247) on: **Sunday, April 12, 2015**

All the content in the _site folder is where Jekyll builds your site. You never edit those files, since that folder&#x27;s contents is deleted and regenerated each time you rebuild your site.

Looking in index.html, this is the code that is getting your recent posts (along with pagination):

&#x60;&#x60;&#x60;
 &lt;ul class=&quot;posts noList&quot;&gt;
          {% for post in paginator.posts %}
            &lt;li&gt;
            	&lt;span class=&quot;date&quot;&gt;{{ post.date | date: &#x27;%B %d, %Y&#x27; }}&lt;/span&gt;
            	&lt;h3&gt;&lt;a href=&quot;{{ site.url }}{{ post.url }}&quot;&gt;{{ post.title }}&lt;/a&gt;&lt;/h3&gt;
            	&lt;p class=&quot;description&quot;&gt;{% if post.description %}{{ post.description  | strip_html | strip_newlines | truncate: 250 }}{% else %}{{ post.content | strip_html | strip_newlines | truncate: 250 }}{% endif %}&lt;/p&gt;
            &lt;/li&gt;
          {% endfor %}
        &lt;/ul&gt;
        &lt;!-- Pagination links --&gt;
        &lt;div class=&quot;pagination&quot;&gt;
          {% if paginator.previous_page %}
            &lt;a href=&quot;{{ site.url }}{{ paginator.previous_page_path }}&quot; class=&quot;previous button__outline&quot;&gt;Newer Posts&lt;/a&gt;  
          {% endif %} 
          {% if paginator.next_page %}
            &lt;a href=&quot;{{ site.url }}{{ paginator.next_page_path }}&quot; class=&quot;next button__outline&quot;&gt;Older Posts&lt;/a&gt;
          {% endif %}
        &lt;/div&gt;
&#x60;&#x60;&#x60;
Just remove this content and replace it with your About information. Create a new page and stick this content in there. See how it works. 
---
> from: [**DanielRoberts95**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92115404) on: **Sunday, April 12, 2015**

While that did work for putting the About page content on the homepage (thank you!) this is what the Blog page looks like for some reason.
![atfirst](https://cloud.githubusercontent.com/assets/11403206/7107236/073aa3b2-e114-11e4-9e31-96d64443ed9b.png)

This is what the index.html code looked like originally in the blog.md
![originalcode](https://cloud.githubusercontent.com/assets/11403206/7107254/bb3f06be-e114-11e4-93c3-486ced21fb3a.png)

I then changed the indentation of the code (by just pressing backspace in front of a lot of the code) which changed it to this:
![fixedindentation](https://cloud.githubusercontent.com/assets/11403206/7107258/d125ba72-e114-11e4-82f2-75a1c6a862a4.png)

Doing so changed the page to this, just making the entirety of the Blog page blank:
![after](https://cloud.githubusercontent.com/assets/11403206/7107267/2c9f0ed0-e115-11e4-97ff-df6d9029f425.png)

Thank you so much for all of your help so far by the way @tomjohnson1492 :)
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92119029) on: **Sunday, April 12, 2015**

Markdown files process spacing as syntax (two spaces to the left sets off text as code), so that&#x27;s why pasting the code into an md file initially probably set it off as code. you have to remove all the spacing to the left at the start of the first element and the last element. I believe you did that, so I&#x27;m not sure why it&#x27;s not rendering. I would play around with the spacing a bit more. 

Might be easier to change the file extension from .md to .html for that recent posts file.

~~(I assume you registered the collection (paginator) in your config file.)~~
---
> from: [**DanielRoberts95**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92121808) on: **Sunday, April 12, 2015**

I messed around with the spacing a bit more, here&#x27;s what the blog.md file looks like now:
![blogmd](https://cloud.githubusercontent.com/assets/11403206/7107378/ae75be10-e118-11e4-8b55-d2e77dad26b9.png)
From what I see the only difference is that it registered the HTML comment as a comment, unfortunately it did not change anything and the Blog page is still blank.

I tried changing the file extension from .md to .html in the blog file but the page is still blank.

Here&#x27;s a screenshot of my config.yml file if you&#x27;re interested.
![config](https://cloud.githubusercontent.com/assets/11403206/7107397/4b9bf88a-e119-11e4-9089-1f41764c3b1c.png)

---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92125478) on: **Sunday, April 12, 2015**

Sorry, I haven&#x27;t actually worked with pagination much, so I incorrectly assumed something about a collection that wasn&#x27;t the case at all. The paginator reference is a feature of Jekyll for paginating blog posts (e.g., showing 10 per page or something). See [http://jekyllrb.com/docs/pagination/](http://jekyllrb.com/docs/pagination/). See esp. the Enable Pagination section. 

I am not sure how to reconfigure pagination on a page other than the homepage. I would assume it&#x27;s possible -- I just don&#x27;t know how. Maybe look at another theme to see how it&#x27;s done?

As a short term fix, you could just change the following line in the recent posts code: 

&#x60;&#x60;&#x60;
{% for post in paginator.posts %}
&#x60;&#x60;&#x60;
to this:

&#x60;&#x60;&#x60;
{% for post in site.posts %}
&#x60;&#x60;&#x60;

Then it works. Not sure how you get pagination, though. On my site, for the blog archive, I just point to a page that will list a bunch of old posts:

&#x60;&#x60;&#x60;
&lt;div class=&quot;home-read-more&quot;&gt;
  &lt;a href=&quot;{{ &quot;/archive&quot; | prepend:site.baseurl }}&quot; class=&quot;btn btn-primary btn-lg&quot;&gt;View All {{ site.posts | size }} Articles →&lt;/a&gt;
&lt;/div&gt;
&#x60;&#x60;&#x60;
As I said, I am primarily working with pages to build out a doc theme, not using Jekyll for a blog (not yet anyway).

Tom
---
> from: [**DanielRoberts95**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92135346) on: **Sunday, April 12, 2015**

It works, thanks! Yeah I&#x27;m not sure how to get pagination either, but I will try my best to figure it out. Thank you so much for all of your help! :)
---
> from: [**jamesonzimmer**](https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92188218) on: **Sunday, April 12, 2015**

FYI to my knowledge, pagination only works with files in your &#x27;post&#x27; folder
via a &#x27;for&#x27; liquid loop. Probably you already know this though :-).

I found the Poole theme (http://getpoole.com) really useful for figuring
out proper ways to use pagination. Also you can get endless scroll hooked
up pretty easy: https://github.com/tobiasahlin/infinite-jekyll

On Mon, Apr 13, 2015 at 5:21 AM, DanielRoberts95 &lt;notifications@github.com&gt;
wrote:

&gt; It works, thanks! Yeah I&#x27;m not sure how to get pagination either, but I
&gt; will try my best to figure it out. Thank you so much for all of your help!
&gt; :)
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub
&gt; &lt;https://github.com/jekyll/jekyll-help/issues/289#issuecomment-92135346&gt;.
&gt;

