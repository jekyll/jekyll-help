# [How can I programmatically modify permalinks?](https://github.com/jekyll/jekyll-help/issues/278)

> state: **open** opened by: **** on: ****

[I asked this question on Stack Overflow](http://stackoverflow.com/questions/28681575/in-jekyll-how-can-i-programmatically-modify-permalinks-for-pages) but thought it would be good to ask as well, sorry if some of you already read this.

I want to tweak the permalink of each page programatically, using [permalink templates](http://jekyllrb.com/docs/permalinks/) is not enough. I tried using a Generator plugin, something like:

&#x60;&#x60;&#x60;ruby
module MySite
  class MySiteGenerator &lt; Jekyll::Generator
    def generate(site)
      site.pages.each do |page|
        page.data[&#x27;permalink&#x27;] = &#x27;/foo&#x27; + page.url
        # real world manipulation of course more complicated
      end
    end
  end
end
&#x60;&#x60;&#x60;
But although this would run and set the &#x60;page.data[&#x27;permalink&#x27;]&#x60; field, the output was still the same.

On Stack Overflow I got the suggestion to instead override the &#x60;Page&#x60; class, like:

&#x60;&#x60;&#x60;ruby
module Jekyll
  class Page
    alias orig_permalink permalink
    def permalink
      permalink    = orig_permalink
      newPermalink = &quot;foo/#{permalink}&quot;
    end
  end
end
&#x60;&#x60;&#x60;
But this didn&#x27;t work either, only if a page had a manually configured permalink in the front matter did it work, otherwise the pages were not generated in &#x60;_site&#x60; at all. [Here&#x27;s a test repo.](https://github.com/skagedal/permalink_mod_test)

Is there a way to get what I want?  Thanks!

### Comments

