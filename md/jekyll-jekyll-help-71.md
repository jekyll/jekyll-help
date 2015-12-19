# [Correct way to add a collection to Site](https://github.com/jekyll/jekyll-help/issues/71)

> state: **closed** opened by: **** on: ****

I am writing a Jekyll plugin that will fetch your latest entries from your social feeds for display on a page.  For an example look at this blog [here](http://www.codersblock.com)

I have my superfeed.rb file in the plugins directory and I am attempting to add a feed_items property to the Site object.
&#x60;&#x60;&#x60;Ruby
module Jekyll
#Monkey patching site
  class Site
    attr_accessor :feed

    def feed_items
      @feed
    end
  end
#Here I generate the social feed
  class FeedGenerator &lt; Generator
    def generate(site)
      #Lots of things happened here to get all_items generated
      site.feed = all_items
    end
  end
#This is the type of class contained in all_items above 
  class FeedItem
    extend Forwardable
    attr_accessor :icon, :source_name, :published, :title, :title_link, :snippet, :image_preview_uri

    def initialize(options = {})
      @options = OpenStruct.new(options)
      self.class.instance_eval do
        def_delegators :@options, *options.keys
      end
    end

    def to_liquid
      {
        &#x27;icon&#x27;              =&gt;  icon,
        &#x27;source_name&#x27;       =&gt;  source_name,
        &#x27;published&#x27;         =&gt;  published,
        &#x27;title&#x27;             =&gt;  title,
        &#x27;title_link&#x27;        =&gt;  title_link,
        &#x27;snippet&#x27;           =&gt;  snippet,
        &#x27;image_preview_uri&#x27; =&gt;  image_preview_uri
      }
    end
  end
&#x60;&#x60;&#x60;
I am then trying to consume this work in one of my pages like so.
&#x60;&#x60;&#x60;HTML
{% for item in site.feed_items %}
&lt;hr/&gt;
&lt;div class=&quot;row&quot;&gt;
    &lt;div class=&quot;small-2 columns&quot;&gt;
        &lt;p&gt;
            &lt;a href=&quot;/blog&quot;&gt;&lt;img src=&quot;{{ item.icon }}&quot;&gt;&lt;/a&gt;
        &lt;/p&gt;
        {{ item.published | date_to_string }}
     &lt;/div&gt;
    &lt;div class=&quot;small-10 columns&quot;&gt;
        &lt;p&gt;
            &lt;strong&gt;&lt;a href=&quot;{{ item.title_url }}&quot;&gt;{{ item.title }}&lt;/a&gt;:&lt;/strong&gt;
            &lt;br&gt;
            {{ item.snippet | strip_html }}
            &lt;a href=&quot;{{ item.title_url }}&quot;&gt;Read full post &amp;raquo;&lt;/a&gt;
        &lt;/p&gt;
    &lt;/div&gt;
&lt;/div&gt;
{% endfor %}
&#x60;&#x60;&#x60;
The problem I am having is that nothing outputs...it seems as if site.feed_items is returning nothing.  Inside of my generator I set a breakpoint and I can see that at the end of my generate method my site object has the collection.  So my question is, is this the right way to do this?  Should I be patching Site?  I am pretty new to Jekyll and I&#x27;ve been having a hard time finding examples doing something like this.

### Comments

---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/71#issuecomment-48097751) on: **Saturday, July 05, 2014**

Not sure if this will help, but here is how I did it. Note that my _config.yml file has the [collections output set to true](http://jekyllrb.com/docs/collections/#step-3-optionally-render-your-collections-documents-into-independent-files).

_config.yml
&#x60;&#x60;&#x60;
collections:
  kb:
    output: true
&#x60;&#x60;&#x60;

KB Page
&#x60;&#x60;&#x60;
{% for each in site.kb %}
  &lt;li&gt;&lt;a href=&quot;{{ each.url }}&quot;&gt;{{ each.title }}&lt;/a&gt;&lt;/li&gt;
{% endfor %}
&#x60;&#x60;&#x60;

Reference: [JekyllKB](https://github.com/AJAlabs/jekyllkb/blob/master/_kb/index.html)
