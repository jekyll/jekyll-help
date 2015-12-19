# [Include tags not rendering when called from plugin](https://github.com/jekyll/jekyll-help/issues/284)

> state: **open** opened by: **** on: ****

I&#x27;m trying to write a basic generator plugin that creates a JSON version of each page in my site.
I&#x27;ve got this working, with one exception: &#x60;{% include %}&#x60; tags in the content of the pages are not being rendered when the JSON version of a page is created.

These tags **are** being rendered when Jekyll creates the standard HTML version of the pages, however, so I don&#x27;t understand what I&#x27;m missing. 

Here&#x27;s the content of my plugin file (adapted from [here](https://github.com/jgarber623/sixtwothree.org/blob/master/src/_plugins/json_page_generator.rb):

&#x60;&#x60;&#x60;ruby
module Jekyll
  class JSONPage &lt; Page
    def initialize(site, base, dir, name, content)
      @site = site
      @base = base
      @dir  = dir
      @name = name

      self.data = {}
      self.content = content

      process(@name)
    end

    def read_yaml(*)
      # Do nothing
    end

    def render_with_liquid?
      false
    end
  end

  class JSONPostGenerator &lt; Generator
    safe true

    def generate(site)
      
      site.pages.each do |page|
        if page.html?

          path = page.destination(&quot;&quot;).sub &quot;.html&quot;, &quot;.json&quot;

          output = page.to_liquid

          output[&#x27;content&#x27;] = page.transform

          site.pages &lt;&lt; JSONPage.new(site, site.source, File.dirname(path), File.basename(path), output.to_json)
        end
      end

    end

  end
end
&#x60;&#x60;&#x60;
In each of my pages, I&#x27;m calling one or more &#x60;include&#x60; tags in the markdown content of the page and passing params. For example:

&#x60;&#x60;&#x60;liquid
{% include compare.html first_figure=&quot;1a&quot; second_figure=&quot;1b&quot; %}
&#x60;&#x60;&#x60;

The HTML output of &#x60;jekyll build&#x60; renders that line as the following HTML:
&#x60;&#x60;&#x60;html
&lt;figure class=&quot;compare&quot;&gt;
  &lt;img class=&quot;lazyload&quot; data-src=&quot;images/001a.jpg&quot; alt=&quot;Figure1a&quot; id=&quot;fig1a&quot; /&gt;
  &lt;img class=&quot;lazyload&quot; data-src=&quot;images/001b.jpg&quot; alt=&quot;Figure1b&quot; id=&quot;fig1b&quot; /&gt;
  &lt;figcaption&gt;
    &lt;span&gt;&lt;strong&gt;Figure 1a.&lt;/strong&gt; Riace Bronze A. Greek, mid-5th century B.C. H: 1.98 m (ca. 6½ ft.).Reggio Calabria, Museo Nazionale.&lt;/span&gt;
    &lt;span&gt;&lt;strong&gt;Figure 1b.&lt;/strong&gt; Riace Bronze B. Greek, mid-5th century B.C. H: 1.97 m (ca. 6½ ft.).Reggio Calabria, Museo Nazionale.&lt;/span&gt; 
  &lt;/figcaption&gt;
&lt;/figure&gt;
&#x60;&#x60;&#x60;

This is all working as it should. However, the JSON output of the same page (which is otherwise correct, complete with front-matter properly formatted) still has the raw &#x60;{% include compare.html %}&#x60; in its content; it&#x27;s not being rendered by Jekyll.

I would assume that calling &#x60;page.transform&#x60; in the plugin would &quot;transform&quot; this particular content, but that&#x27;s not happening. Any ideas?

Sample (abbreviated) JSON output below – HTML content is otherwise rendering, escaping, etc. – but Jekyll-specific &#x60;{%include%}&#x60; syntax is untouched...

&#x60;&#x60;&#x60;json
{
  &quot;id&quot;: 0,
  &quot;title&quot;: &quot;Foreward&quot;,
  &quot;cover_image&quot;: &quot;023.jpg&quot;,
  &quot;content&quot;: &quot;&lt;p&gt;Fact and fiction mingle in the classical myth of Talos, a Cretan giant made of bronze... &lt;/p&gt;\n\n&lt;p&gt;{% include compare.html first_figure=\u201d1a\u201d second_figure=\u201d1b\u201d %}&lt;/p&gt;&quot;
}
&#x60;&#x60;&#x60;

### Comments

