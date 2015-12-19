# [Collection files don&#x27;t display without YAML front matter](https://github.com/jekyll/jekyll-help/issues/223)

> state: **closed** opened by: **** on: ****

I&#x27;ve been playing with collections since they first came out and I built this with it: http://style.maban.co.uk/ (code here: https://github.com/maban/styleguide) – each item on the page is pulled from a collection which is a series of pure HTML files that don&#x27;t contain any front matter, and I generated it like this:

&#x60;&#x60;&#x60;
{% for pattern in site.patterns %}
    {{ pattern.content }}
{% endfor %}
&#x60;&#x60;&#x60;
and in my config file I&#x27;ve got
&#x60;&#x60;&#x60;
collections:
    patterns:
        output: true
&#x60;&#x60;&#x60;
It was working great until recently, and I hadn&#x27;t touched the code until someone said it wasn&#x27;t working for them and I tried it out again locally. Now when I build the site, none of the files in the collection get generated on the page. The only way they&#x27;ll display at all is if I insert an empty front matter to each file in the collection.

[In the documentation](http://jekyllrb.com/docs/collections/) for the content variable of collections, it says &quot;If no YAML Front Matter is provided, this is the entirety of the file contents.&quot;, so it should still work, right?

Has something changed in the codebase that basically says &quot;don&#x27;t output this collection file if it&#x27;s not got front matter&quot;? Or am I doing something wrong – is there a way to force it to generate it anyway?

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-68660085) on: **Sunday, January 04, 2015**

As far as I know, Jekyll will only process a file through it&#x27;s generator system if it contains front-matter. I honestly don&#x27;t know the history of that behavior, but I feel like it&#x27;s been like that since Collections was introduced. Obviously there could be nuiances that I&#x27;m not aware of, but the idea is generally: &quot;Add front-matter to files that you want Jekyll to process.&quot;
---
> from: [**maban**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-68677588) on: **Sunday, January 04, 2015**

It did used to work. If that&#x27;s how it is now, would it be worth changing the wording of the documentation? This had me thinking it would process a file without front-matter:

&gt; Create a corresponding folder (e.g. &lt;source&gt;/_my_collection) and add documents. YAML Front Matter is read in as data if it exists, if not, then everything is just stuck in the Document’s content attribute.
http://jekyllrb.com/docs/collections/#step-2-add-your-content
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-68685195) on: **Monday, January 05, 2015**

@maban That&#x27;s a very inspirational use of Jekyll, I love it! 

:bow: 
---
> from: [**maban**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-68686229) on: **Monday, January 05, 2015**

Thank you, @nternetinspired! I&#x27;m also building a jekyll-based collaborative bookmarking tool. Here are my very early workings: https://github.com/maban/bookmarks
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-68687790) on: **Monday, January 05, 2015**

Looks great @maban, I&#x27;ll certainly keep an eye on that. I&#x27;m really interested in how Jekyll can be used in new ways, like cutting production time of the less glamorous, but essential, parts of the web design process; and I think style guides are a really important part of that.

I&#x27;ve been dabbling with using Jekyll to replace traditional web design tools in rapidly creating things like element collages, style tiles, and type samples in addition to just wireframing: https://github.com/nternetinspired/jekyll-design 

A fully fleshed out pattern library is next on my hit list… 
---
> from: [**maban**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-68687964) on: **Monday, January 05, 2015**

Oooh, looks really cool!
---
> from: [**paulrobertlloyd**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-68757540) on: **Monday, January 05, 2015**

Hey @nternetinspired, your project looks fantastic! I love what people are doing with Jekyll… next on my list is Jekyll-ising my Barebones project (https://github.com/paulrobertlloyd/barebones).
---
> from: [**nternetinspired**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-68991895) on: **Wednesday, January 07, 2015**

Thanks for the kind words @paulrobertlloyd :blush: 

I really need to give it some more time and refine things a bit further because I&#x27;m happy with how it works as a POC. I&#x27;ll be keeping an eye on Barebones for inspiration ;) 
---
> from: [**pmackay**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-69164077) on: **Thursday, January 08, 2015**

I&#x27;d agree here that the docs need a fix, or the behaviour needs changing so files without front matter are processed.
---
> from: [**maban**](https://github.com/jekyll/jekyll-help/issues/223#issuecomment-69164616) on: **Thursday, January 08, 2015**

Oh, the docs are on GitHub! I&#x27;ll try and write something that might help clarify things, and do a pull request after work.
