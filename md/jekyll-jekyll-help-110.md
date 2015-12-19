# [Upgrading jekyll-pages from 20 to 21 broke Markdown parsing](https://github.com/jekyll/jekyll-help/issues/110)

> state: **closed** opened by: **** on: ****

Going from 20 to 21 made this:

![screen shot 2014-08-02 at 00 36 16](https://cloud.githubusercontent.com/assets/198682/3785950/f4d9a36c-19cc-11e4-85e0-351814dadaf8.png)

look like this:

![screen shot 2014-08-02 at 00 36 40](https://cloud.githubusercontent.com/assets/198682/3785946/ed2d3408-19cc-11e4-83fe-87174a86b7da.png)

The offending code is this:

&#x60;&#x60;&#x60;md
## Further Reading
* Kjærsgaard decries public intolerance and harassment by Muslims in [her weekly party newsletter from November 24 2003][noerrebro].

    &gt;In the interim, we back home will work to get back [the capital district] Nørrebro.
    &gt;
    &gt;So that once again, tolerance and free-mindedness can come to the north of Dronning Louise&#x27;s Bridge.
* [2014 report][fra-report] by the European Agency for Fundamental Rights on violence againt women.
&#x60;&#x60;&#x60;

I’d imagine it has to do with Kramdown, but from what I can tell, the version remains 1.3.1 between upgrades, so I don’t really know what’s up.

The whole project is available [here](http://github.com/hafniatimes/hafniatimes.github.io).

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/110#issuecomment-50944287) on: **Friday, August 01, 2014**

Maybe https://github.com/jekyll/jekyll/issues/2676 ?
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/110#issuecomment-51027143) on: **Monday, August 04, 2014**

It looks like upgrading from 20 to 21 broke the Markdown on my **local** set-up, but from what I can tell, the parsing doesn’t work on *either* in **production**.

So while there is a general problem, I’m gonna hold off on blaming 21 for the time being, until I can figure out how to use something like &#x60;rbenv&#x60; to perform a full &#x60;git bisect&#x60; diagnosis.

([Speaking of](https://github.com/jekyll/jekyll/issues/1997).) :)
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/110#issuecomment-55643927) on: **Monday, September 15, 2014**

@ndarville Are you still having issues with this? Going to close for now since this is a month old, but if it&#x27;s still giving you trouble, feel free to re-open it.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/110#issuecomment-55659022) on: **Monday, September 15, 2014**

Yep. Not sure what to do about it, but I’ve just written it in plain HTML for the sake of my own sanity.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/110#issuecomment-55660083) on: **Monday, September 15, 2014**

Oh, I see what&#x27;s going on. Looks like Kramdown is trying to format that as a code block, since it&#x27;s indented. But you&#x27;ve only indented it because it&#x27;s part of a list.

This is exactly the same problem as jekyll/jekyll#2676, which doesn&#x27;t look like it was fully resolved.

Maybe create an issue with the [Kramdown](https://github.com/gettalong/kramdown) folks?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/110#issuecomment-55973451) on: **Wednesday, September 17, 2014**

Hey @ndarville, I did some experimentation with this. Check it out:

&#x60;&#x60;&#x60;markdown
* List item 1

  &gt; some quote
  
      some code

* List item 2
* List item 3
&#x60;&#x60;&#x60;

renders as...

&#x60;&#x60;&#x60;html
&lt;ul&gt;
  &lt;li&gt;
    &lt;p&gt;List item 1&lt;/p&gt;

    &lt;blockquote&gt;
      &lt;p&gt;some quote&lt;/p&gt;
    &lt;/blockquote&gt;

    &lt;pre&gt;&lt;code&gt;some code
&lt;/code&gt;&lt;/pre&gt;
  &lt;/li&gt;
  &lt;li&gt;List item 2&lt;/li&gt;
  &lt;li&gt;List item 3&lt;/li&gt;
&lt;/ul&gt;
&#x60;&#x60;&#x60;

Hope this helps!
