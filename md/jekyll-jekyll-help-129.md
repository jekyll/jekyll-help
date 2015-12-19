# [category with 2 words url problems](https://github.com/jekyll/jekyll-help/issues/129)

> state: **open** opened by: **** on: ****

I am defining:
category: abc xyz
in the post front matter.

Then I have the permalink in the _config file as:
permalink: /:categories/:title/

when the site generates I get a url that has spaces in it.

Is it possible to have the space in the category replaced with a dash in the url?

This is breaking when I upload to Amazon S3.

### Comments

---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52447915) on: **Sunday, August 17, 2014**

&#x60;categories: abc xyz&#x60;
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52448983) on: **Sunday, August 17, 2014**

Not sure I understand. It is one category with 2 words - like My Category - not two separate one word categories.

Are you saying the proper way to reference a category in the front matter is with categories rather than category?

I am pretty confused by the difference (if there is one).

I will try it out on a sample site - the actual site already has hundreds of posts - not as simple to test.
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52449302) on: **Sunday, August 17, 2014**

it looks like 
&#x60;&#x60;&#x60;categories: abc xyz&#x60;&#x60;&#x60; 
gives me abc/xyz/title/ 

and 
&#x60;&#x60;&#x60;category: abc xyz&#x60;&#x60;&#x60; 
gives me /abc xyz/title/ - which is what I want but with a dash instead of a space.

What I want is /abc-xyz/title/

I&#x27;m thinking having a category with more than one word is a problem.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52458769) on: **Monday, August 18, 2014**

&#x60;&#x60;&#x60;yaml
categories:
  - This is one category
  - This is another category
&#x60;&#x60;&#x60;
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52485234) on: **Monday, August 18, 2014**

And or:

&#x60;categories: [&#x27;The is a category&#x27;, &#x27;This is another category&#x27;]&#x60;

Documentation:  http://jekyllrb.com/docs/frontmatter/#predefined-global-variables
* see Category/Categories

Documentation for YAML lists:  http://en.wikipedia.org/wiki/YAML#Lists


---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52485919) on: **Monday, August 18, 2014**

&gt; Documentation for YAML lists: http://en.wikipedia.org/wiki/YAML#Lists

Wow. Since when is Wikipedia the place where you find the indept docs? Good to know.
---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52509515) on: **Monday, August 18, 2014**

I understand how all of those examples work, but none of them generate a proper url with a dash instead of a space - sure, they are multi-word categories, but the resulting url to the post has spaces between the words which breaks on amazon s3.

I don&#x27;t understand who would want a space in the post url.

I don&#x27;t think this is solvable as is (other than by taking category out of the url), so I made all of my categories one word :(.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52517713) on: **Monday, August 18, 2014**

Oh wow, I haven’t realized this. This is not good. Generated URL&#x27;s should have dashes instead of spaces, in my opinion. I think we should enforce this Jekyll’s permalinks.
---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52599946) on: **Tuesday, August 19, 2014**

If the &#x60;category: abc xyz&#x60; is actually supported, I’d be sure to write it as &#x60;category: &quot;abc&quot; &quot;xyz&quot;&#x60; just to remove any ambiguous interpretation of the two elements.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-52600395) on: **Tuesday, August 19, 2014**

@ndarville If you use the singular keyword, you only will get one category. However you can use the plural key with only one category. Thus always using &#x60;categories: […]&#x60; is fine and leaves open all possibilities.
---
> from: [**JonathanReeve**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-55535973) on: **Sunday, September 14, 2014**

+1 on this. I have a new jekyll blog (at http://jonreeve.com) that has a category called &quot;digital humanities.&quot; I had to choose a permalink that didn&#x27;t include the category, since a permalink with the form &#x60;http://domain.com/category/year/month/title&#x60; gave me urls like &#x60;http://jonreeve.com/digital humanities/2014/04/post.html&#x60;. If anyone knows of a plugin that can convert categories containing spaces into urls with dashes, please let me know. 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-55537284) on: **Sunday, September 14, 2014**

&gt; If anyone knows of a plugin that can convert categories containing spaces into urls with dashes, please let me know.

If this is really not the case, I consider this a bug which should be fixed. Categories with spaces should result in a URL with dashes.
---
> from: [**binoyte**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-57279943) on: **Tuesday, September 30, 2014**

I have a &quot;black and white&quot; category. How can I filter like :

    {% for post in site.categories.&#x27;black and white&#x27; %}

I think I have to replace blanks by underscores and use &#x60; | replace: &#x27;_&#x27;,&#x27;  &#x27;&#x60; for clean ouputs.

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-57282355) on: **Tuesday, September 30, 2014**

&#x60;&#x60;&#x60;liquid
{% for post in site.posts | where: &#x27;category&#x27;,&#x27;black and white&#x27; %}
&#x60;&#x60;&#x60;

Can you try something like this? I haven&#x27;t tested that. I&#x27;m not sure, because when you have multiple categories you usually have &#x60;categories&#x60; instead of &#x60;category&#x60;.
---
> from: [**jrabef**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-58727520) on: **Friday, October 10, 2014**

Has this been looked into? I&#x27;m considering migrating from wordpress to Jekyll and I want this &quot;feature&quot; added. :)
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-58742957) on: **Saturday, October 11, 2014**

Multiple word categories are available.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-59679478) on: **Sunday, October 19, 2014**

Is the problem stated as follows?

&gt; Multi-word categor[y|ies] separated with white space results in a permalink with an explicit space versus a slugged version (hyphenated).

-- 

&gt; Slugify
Convert a string into a lowercase URL &quot;slug&quot; by replacing every sequence of spaces and non-alphanumeric characters with a hyphen.

http://jekyllrb.com/docs/templates/#filters

Should we be slugify&#x27;ing the permalink but are not?

---
> from: [**rdyar**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-59680327) on: **Sunday, October 19, 2014**

yes, that is my problem. I gave up and made all my multi word categories into one word categories. 

Multi word categories work fine except for the space in the url is not slugged.

I have not tested this since August but assume it has not been fixed yet.
---
> from: [**JonathanReeve**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-59752586) on: **Monday, October 20, 2014**

@kleinfreund, yes, multiple word categories are available, if you don&#x27;t mind having malformed URLs. But people probably don&#x27;t want spaces in their URLs. @jaybe-jekyll: you got it. That&#x27;s the problem. 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-59809269) on: **Monday, October 20, 2014**

That is an issue which should be fixed. We can&#x27;t support multiple word categories without them being slugified like single word categories are.

I consider that a bug.
---
> from: [**djacquel**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-61255284) on: **Friday, October 31, 2014**

This needs to change the way Jekyll generates urls for posts. This can be done with a plugin.

    # _plugins/post.rb
    module Jekyll

      class Post

        # override post method in order to return categories names as slug
        # instead of strings
        #
        # An url for a post with category &quot;category with space&quot; will be in
        # slugified form : /category-with-space
        # instead of url encoded form : /category%20with%20space
        #
        # @see utils.slugify
        def url_placeholders
          {
              :year        =&gt; date.strftime(&quot;%Y&quot;),
              :month       =&gt; date.strftime(&quot;%m&quot;),
              :day         =&gt; date.strftime(&quot;%d&quot;),
              :title       =&gt; slug,
              :i_day       =&gt; date.strftime(&quot;%-d&quot;),
              :i_month     =&gt; date.strftime(&quot;%-m&quot;),
              :categories  =&gt; (categories || []).map { |c| Utils.slugify(c) }.join(&#x27;/&#x27;),
              :short_month =&gt; date.strftime(&quot;%b&quot;),
              :short_year  =&gt; date.strftime(&quot;%y&quot;),
              :y_day       =&gt; date.strftime(&quot;%j&quot;),
              :output_ext  =&gt; output_ext
          }
        end

      end

    end

@kleinfreund : The actual way Jekyll does urls generation is [RFC 1738](http://www.ietf.org/rfc/rfc1738.txt) compliant. There is no bug in here.


---
> from: [**djacquel**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-61261006) on: **Friday, October 31, 2014**

@xirux-nefer : Beware of Jekyll Categories !

This is valid : 

    categories: [mutli word, toto]
    or
    category: mutli word

But this doesn&#x27;t have the same effect on post object.
When declaring &#x60;category&#x60;, post.category (string) and post.categories (array) are set.
When declaring &#x60;categories&#x60;, only post.categories is set.

The test &#x60;{% if post.category == &quot;Something with several words&quot; %}&#x60; might fail if, for example you have default categories set in config. 

It&#x27;s better to test on &#x60;post.categories&#x60; with &#x60;{% if post.categories contains &quot;Something with several words&quot; %}&#x60;

I have checked you tests on a fresh Jekyll and :

    {% for post in site.posts | where: &#x27;category&#x27;,&#x27;Something with several words&#x27; %}

Will not work, because you cannot put a filter on a loop. You have to capture first, then loop.

    {% capture myposts %}
    {{ site.posts where: &#x27;category&#x27;,&#x27;Something with several words&#x27; %}
    {%endcapture%}
    {% for post in myposts %}
    ...

**problem** : if your capture return only one post, you cannot loop through it. The loop will only work if you have two or more posts returned.

With &#x60;category: &quot;Something with several words&quot;&#x60; set, this test works for me.

    {% for post in site.posts %}
        {% if post.category == &quot;Something with several words&quot; %}
  

    






---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-64135413) on: **Sunday, November 23, 2014**

If we slugify these categories, we will likely break a lot of sites. Your web server should be able to handle spaces in URL&#x27;s.
---
> from: [**pamgriffith**](https://github.com/jekyll/jekyll-help/issues/129#issuecomment-89047759) on: **Thursday, April 02, 2015**

Is there any progress on this? I, too, would like to have multi-word categories and no spaces in my urls. I will try the plugin, but this seems like it should be a core feature, as I believe spaces in URLs are supposed to be percent encoded (both by the RFC referenced above, which specifically calls them out as &quot;unsafe&quot; and [rfc3986](http://tools.ietf.org/html/rfc3986), which updates it).
