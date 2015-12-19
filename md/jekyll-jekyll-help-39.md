# [Thoughts on a Content Folder?](https://github.com/jekyll/jekyll-help/issues/39)

> state: **closed** opened by: **** on: ****

**(tl;dr) Seems to me that it&#x27;d be great to put all content, whether that&#x27;s collections, pages, or posts into one content folder.**


Jumping in to collections I&#x27;m creating a site now that has a bunch of folders of varying content, plus regular pages and posts. This is getting unruly because in one directory I&#x27;ve got all those content files and folders and config files, CNAME file, etc. etc.

I&#x27;ve long put my pages in their own folder (_pages) and created an include directive for that folder in my config file and just set up permalinks in front-matter. 

I don&#x27;t think there&#x27;s currently a way to do what I&#x27;m suggesting aside from altering the permalink in each and every document in the front-matter, but I don&#x27;t think we should need to put permalinks inside of every document in a collection because that could be cumbersome and error prone for content creators.

I think someone on here had suggested essentially the opposite, putting layouts and includes in a theme folder, but at any rate, I&#x27;d like to be able to organize things a bit better. 

I&#x27;d be interested in hearing any thoughts you guys have on this or any way you&#x27;ve handled this yourself.

### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42601525) on: **Thursday, May 08, 2014**

That is precisely what the source option is for. See for instance, https://bitbucket.org/matthew-scharley/matt.scharley.me/src
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42602312) on: **Thursday, May 08, 2014**

Fabulous! Can you just put content in &quot;source&quot;? 
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42603081) on: **Thursday, May 08, 2014**

Basically it moves where jekyll looks for just about everything instead of
&#x27;./&#x27; by default. Not sure if you don&#x27;t specify plugins if that will then
default to source + &#x27;/_plugins&#x27; or not though. Check the docs for what
other options are there. Definitely sounds like source is the option you&#x27;re
looking for though.


On 9 May 2014 06:30, Bud Parr &lt;notifications@github.com&gt; wrote:

&gt; Fabulous! Can you just put content in &quot;source&quot;?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42602312&gt;
&gt; .
&gt;
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42603343) on: **Thursday, May 08, 2014**

Yeah, the docs just say &quot;Change the directory where Jekyll will read files&quot; so that&#x27;s for everything, not just content. It&#x27;s close though.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42603653) on: **Thursday, May 08, 2014**

Personally, I&#x27;d argue that source should be the option you are looking for
and there should be more options to move the special folders like _includes
individually. Probably just me arguing semantics though.


On 9 May 2014 06:40, Bud Parr &lt;notifications@github.com&gt; wrote:

&gt; Yeah, the docs just say &quot;Change the directory where Jekyll will read
&gt; files&quot; so that&#x27;s for everything, not just content. It&#x27;s close though.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42603343&gt;
&gt; .
&gt;
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42603978) on: **Thursday, May 08, 2014**

Either way would be fine if they accomplish the goal of keeping content orderly.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42604147) on: **Thursday, May 08, 2014**

Personally, I don&#x27;t have an issue with having jekyll&#x27;s special folders in
with the content, it was mostly just that I didn&#x27;t want to exclude all the
other project files all manually. Each to their own. @parkr: Opinions?


On 9 May 2014 06:45, Bud Parr &lt;notifications@github.com&gt; wrote:

&gt; Either way would be fine if they accomplish the goal of keeping content
&gt; orderly.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42603978&gt;
&gt; .
&gt;
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/39#issuecomment-42604402) on: **Thursday, May 08, 2014**

I want to say to a content creator &quot;your content is all in this folder, ignore everything else.&quot; 
