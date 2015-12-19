# [How can I modify slugify utility so that I can use underscores in URLs, not hyphens.](https://github.com/jekyll/jekyll-help/issues/254)

> state: **open** opened by: **** on: ****

I’m moving my old site (with legacy URLs) over to Jekyll, and that site generated URLs with underscores (which I actually prefer). I&#x27;ve subclassed slugify (&#x60;Jekyll::Utils&#x60;) to achieve this on my new Jekyll-powered site, but I&#x27;d like to release this as a plugin with a configurable value. This is what I’ve got so far:

&#x60;&#x60;&#x60;ruby
module Jekyll

  class SlugifyCharacterGenerator &lt; Jekyll::Generator
    def generate(site)
      @site = site
      @slugify_character = @site.config[&#x27;slugify_character&#x27;]

      if @slugify_character.nil?
        @slugify_character = &#x27;-&#x27;
      else
        @slugify_character = @slugify_character
      end
    end
  end

  module Utils
    # Slugify a filename or title.
    #
    # string - the filename or title to slugify
    # mode - how string is slugified
    #
    # When mode is &quot;none&quot;, return the given string in lowercase.
    #
    # When mode is &quot;raw&quot;, return the given string in lowercase,
    # with every sequence of spaces characters replaced with a hyphen.
    #
    # When mode is &quot;default&quot; or nil, non-alphabetic characters are
    # replaced with a hyphen too.
    #
    # When mode is &quot;pretty&quot;, some non-alphabetic characters (._~!$&amp;&#x27;()+,;=@)
    # are not replaced with hyphen.
    #
    # Examples:
    #   slugify(&quot;The _config.yml file&quot;)
    #   # =&gt; &quot;the-config-yml-file&quot;
    #
    #   slugify(&quot;The _config.yml file&quot;, &quot;pretty&quot;)
    #   # =&gt; &quot;the-_config.yml-file&quot;
    #
    # Returns the slugified string.
    def slugify(string, mode=nil)
      mode ||= &#x27;default&#x27;
      return nil if string.nil?
      return string.downcase unless SLUGIFY_MODES.include?(mode)

      # Replace each character sequence with a hyphen
      re = case mode
      when &#x27;raw&#x27;
        SLUGIFY_RAW_REGEXP
      when &#x27;default&#x27;
        SLUGIFY_DEFAULT_REGEXP
      when &#x27;pretty&#x27;
        # &quot;._~!$&amp;&#x27;()+,;=@&quot; is human readable (not URI-escaped) in URL
        # and is allowed in both extN and NTFS.
        SLUGIFY_PRETTY_REGEXP
      end

      string.
        # Strip according to the mode
        gsub(re, @slugify_character).
        # Remove leading/trailing hyphen
        gsub(/^\-|\-$/i, &#x27;&#x27;).
        # Downcase
        downcase
    end

  end
end
&#x60;&#x60;&#x60;

@slugify_character isn&#x27;t output when placed inside the subclassed slugify. I am a total Ruby n00b — how can I make this work? Secondly, do I need to replicate the entirety of this class, or is there a simpler means of achieving the same… I bet there is!

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/254#issuecomment-71501375) on: **Monday, January 26, 2015**

I don&#x27;t know Ruby very well myself, but hopefully someone will be able to help out.

However, I would be remiss if I didn&#x27;t point out that you really should use hyphens instead of underscores in your URL, though. Google specifically recommends this in a [support article](https://support.google.com/webmasters/answer/76329?hl=en):

&gt; We recommend that you use hyphens (-) instead of underscores (_) in your URLs.
---
> from: [**paulrobertlloyd**](https://github.com/jekyll/jekyll-help/issues/254#issuecomment-71501570) on: **Monday, January 26, 2015**

Google recommending hyphens is actually one of the reasons I prefer to use underscores ;-)
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/254#issuecomment-71506392) on: **Monday, January 26, 2015**

&gt; Google recommending hyphens is actually one of the reasons I prefer to use underscores ;-)

Why?
---
> from: [**paulrobertlloyd**](https://github.com/jekyll/jekyll-help/issues/254#issuecomment-71507974) on: **Monday, January 26, 2015**

@kleinfreund:

&gt; Make pages primarily for users, not for search engines.
— https://support.google.com/webmasters/answer/35769?hl=en
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/254#issuecomment-71508219) on: **Monday, January 26, 2015**

@paulrobertlloyd @kleinfreund :disappointed: Let&#x27;s not get into this here, please.

If anyone wants to help out with @paulrobertlloyd&#x27;s plugin, feel free to chime in.
