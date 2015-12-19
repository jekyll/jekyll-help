# [Custom liquid tags: what’s the preferred method of providing arguments containing quotes](https://github.com/jekyll/jekyll-help/issues/239)

> state: **open** opened by: **** on: ****

(Cross-posted from https://github.com/Shopify/liquid/issues/507 – as received no answers, hoping to find some help here!)

* * *

**TL;DR: How should I provide/parse arguments that may contain a single or double quote(s) in a custom tag?**

* * *

Not sure if this is the correct place to ask a question, but couldn&#x27;t find documentation elsewhere…

I’ve created a custom liquid tag block, to render a figure with optional caption. Here’s an example:

&#x60;&#x60;&#x60;liquid
{% figure class:&#x27;class1 class2&#x27;, caption:&#x27;&#x27;Cloud Gate&#x27; by Anish Kapoor&#x27; %}
![](/assets/images/cloudgate.jpg)
{% endfigure %}
&#x60;&#x60;&#x60;

Firstly, is this the correct method of supplying arguments to a custom block? I see that pipe delimited arguments are for filters, the = symbol for assignments.

I&#x27;m using &#x60;Liquid::TagAttributes&#x60;, to parse each argument. I’m then having to remove (with &#x60;gsub&#x60;) the first and last &#x60;&#x27;&#x60; or &#x60;&quot;&#x60; else these get output in the HTML. But the generated content is malformed if strings contain quotes. I’ve documented some examples here: 

  https://gist.github.com/paulrobertlloyd/dff85e8af0f039255760

This is the code for my custom tag:

&#x60;&#x60;&#x60;ruby
module Jekyll
  class FigureTag &lt; Liquid::Block
    def initialize(tag_name, markup, tokens)
      super
      @attributes = {}
 
      markup.scan(Liquid::TagAttributes) do |key, value|
        @attributes[key] = value.gsub(/^&#x27;|&quot;/, &#x27;&#x27;).gsub(/&#x27;|&quot;$/, &#x27;&#x27;)
      end
    end
 
    def render(context)
      site = context.registers[:site]
      converter = site.getConverterImpl(::Jekyll::Converters::Markdown)
 
      figure_classes = @attributes[&#x27;class&#x27;].to_s
      figure_main = converter.convert(super(context))
      figure_caption = converter.convert(@attributes[&#x27;caption&#x27;].to_s)
 
      # Render &lt;figure&gt;
      if @attributes[&#x27;class&#x27;]
        source = &quot;&lt;figure class=\&quot;figure #{figure_classes}\&quot;&gt;&quot;
      else
        source = &quot;&lt;figure class=\&quot;figure\&quot;&gt;&quot;
      end
  
      source += &quot;&lt;div class=\&quot;figure__main\&quot;&gt;#{figure_main}&lt;/div&gt;&quot;
 
      if @attributes[&#x27;caption&#x27;]
        source += &quot;&lt;figcaption class=\&quot;figure__caption\&quot;&gt;#{figure_caption}&lt;/figcaption&gt;&quot;
        source += &quot;&lt;/figure&gt;&quot;
      else
        source += &quot;&lt;/figure&gt;&quot;
      end
    end
  end
end
 
Liquid::Template.register_tag(&#x27;figure&#x27;, Jekyll::FigureTag)
&#x60;&#x60;&#x60;

Hope this makes sense!

### Comments

