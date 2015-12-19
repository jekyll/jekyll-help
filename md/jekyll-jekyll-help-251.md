# [How to create simple custom note tag?](https://github.com/jekyll/jekyll-help/issues/251)

> state: **closed** opened by: **** on: ****

I want to create a custom liquid tag that allows me to easily wrap a note in a specific class. 

For example, currently to create a note, I write this:

&#x60;&#x60;&#x60;
&lt;div class=&quot;alert alert-info&quot; role=&quot;alert&quot;&gt;This is a note. &lt;/div&gt;
&#x60;&#x60;&#x60;

I want to create a custom liquid tag that allows me to simply write this:

&#x60;&#x60;&#x60;
{% alert This is a note.%}
&#x60;&#x60;&#x60;

I looked at the [Jekyll docs](http://jekyllrb.com/docs/plugins/#available_plugins) for creating a custom liquid tags via plugins, and found this sample:

&#x60;&#x60;&#x60;ruby
module Jekyll
  class RenderTimeTag &lt; Liquid::Tag

    def initialize(tag_name, text, tokens)
      super
      @text = text
    end

    def render(context)
      &quot;#{@text} #{Time.now}&quot;
    end
  end
end

Liquid::Template.register_tag(&#x27;render_time&#x27;, Jekyll::RenderTimeTag)
&#x60;&#x60;&#x60;

I saved this code as a file and added it to my _plugins directory -- and it worked. So I tried to clone the code and modify it as follows:

&#x60;&#x60;&#x60;
module Jekyll
  class RenderNoteTag &lt; Liquid::Tag

    def initialize(tag_name, text, tokens)
      super
      @text = text
    end

    def render(context)
      &quot;#{@text} &lt;div class=&#x27;alert alert-info&#x27; role=&#x27;alert&#x27;&gt;my alert&lt;/div&gt;&quot;
    end
  end
end

Liquid::Template.register_tag(&#x27;render_note&#x27;, Jekyll::RenderNoteTag)
&#x60;&#x60;&#x60;

I then inserted it on a page using this:

&#x60;&#x60;&#x60;liquid
{% render_note %}
&#x60;&#x60;&#x60;

It works, but how do I add my note in the liquid tags? I want to add something like this:

&#x60;&#x60;&#x60;liquid
{% render_note this is my note %}
&#x60;&#x60;&#x60;

And have &quot;this is my note&quot; wrapped inside like this:

&#x60;&#x60;&#x60;html
&lt;div class=&quot;alert alert-info&quot; role=&quot;alert&quot;&gt;This is a note. &lt;/div&gt;
&#x60;&#x60;&#x60;

### Comments

---
> from: [**pootsbook**](https://github.com/jekyll/jekyll-help/issues/251#issuecomment-71469252) on: **Monday, January 26, 2015**

I believe you were almost there. This is worked for me (&#x60;jekyll 2.5.3&#x60;):

&#x60;&#x60;&#x60;ruby
# _plugins/render_note.rb
module Jekyll
  class RenderNoteTag &lt; Liquid::Tag

    def initialize(tag_name, text, tokens)
      super
      @text = text
    end

    def render(context)
      &quot;&lt;div class=\&quot;alert alert-info\&quot; role=\&quot;alert\&quot;&gt;#{@text}&lt;/div&gt;&quot;
    end
  end
end

Liquid::Template.register_tag(&#x27;render_note&#x27;, Jekyll::RenderNoteTag)
&#x60;&#x60;&#x60;
Then in a post,
&#x60;&#x60;&#x60;liquid
{% render_note This is a note. %}
&#x60;&#x60;&#x60;
resulted in:
&#x60;&#x60;&#x60;html
&lt;div class=&quot;alert alert-info&quot; role=&quot;alert&quot;&gt;This is a note. &lt;/div&gt;
&#x60;&#x60;&#x60;

Let me know how you get on.
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/251#issuecomment-71487338) on: **Monday, January 26, 2015**

Thanks for the help. However, the code you suggested didn&#x27;t work. It doesn&#x27;t render any surrounding div tags at all. I just get &#x60;&lt;p&gt;This is a note&lt;/p&gt;&#x60;. 

I&#x27;m not familiar with the liquid syntax here, so I&#x27;m not sure where &#x60;#{@text}&#x60; must appear. Is this ruby, or is it specific to Liquid?
---
> from: [**pootsbook**](https://github.com/jekyll/jekyll-help/issues/251#issuecomment-71515924) on: **Monday, January 26, 2015**

The syntax is Ruby String interpolation, e.g.

&#x60;&#x60;&#x60;ruby
@text = &quot;example&quot;
puts &quot;This is an #{@text}&quot;
#=&gt; &quot;This is an example&quot;
&#x60;&#x60;&#x60;
Note, that for the variable to be interpolated, the interpolation &#x60;#{}&#x60; must occur within a double quoted string.

Can you paste your code in so I can have a look for you?


---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/251#issuecomment-71517330) on: **Monday, January 26, 2015**

Hey, it&#x27;s working now. I didn&#x27;t restart the server earlier. When I restarted it, it&#x27;s all working. (The code is the same as what you pasted above.) Thanks so much for your help. This technique is going to be tremendously helpful.
---
> from: [**pootsbook**](https://github.com/jekyll/jekyll-help/issues/251#issuecomment-71517738) on: **Monday, January 26, 2015**

Great :+1: 
