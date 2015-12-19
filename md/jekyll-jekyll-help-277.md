# [Access markup-free post and page content from Jekyll plugins](https://github.com/jekyll/jekyll-help/issues/277)

> state: **open** opened by: **** on: ****

[I’m working on a plugin to parse all posts and gather them into a JSON file to be consumed by a search mechanism.](https://github.com/tohuw/tohuw.net/blob/f6ff38247fb88f171ed2053631602427baac9ce9/_plugins/generate_searchcontent.rb) **How can I access just the text of the post from my plugin, with no markup applied?** I’m currently accessing &#x60;site.posts&#x60;, then e.g. &#x60;page.content&#x60; in loops. This returns the content of the post, but includes newline markers (&#x60;\n&#x60;) and Markdown syntax.

I saw another question in which someone wanted to [get Markdown processed content in a Jekyll tag plugin](https://stackoverflow.com/questions/19169849/how-to-get-markdown-processed-content-in-jekyll-tag-plugin), but my case is different: I don&#x27;t want any markup at all, just the plain text of the post, with no formatting applied.

Below is the key &#x60;def&#x60; from my current implementation.

    def generate(site)
      target = File.open(&#x27;js/searchcontent.js&#x27;, &#x27;w&#x27;)
      target.truncate(target.size)
      target.puts(&#x27;var tipuesearch = {&quot;pages&quot;: [&#x27;)

      all_but_last, last = site.posts[0..-2], site.posts.last

      # Process all posts but the last one
      all_but_last.each do |page|
        tp_page = TipuePage.new(
          page.data[&#x27;title&#x27;],
          &quot;#{page.data[&#x27;tags&#x27;]} #{page.data[&#x27;categories&#x27;]}&quot;,
          page.url,
          page.content
        )
        target.puts(tp_page.to_json + &#x27;,&#x27;)
      end

      # Do the last post
      tp_page = TipuePage.new(
        last.data[&#x27;title&#x27;],
        &quot;#{last.data[&#x27;tags&#x27;]} #{last.data[&#x27;categories&#x27;]}&quot;,
        last.url,
        last.content
      )
      target.puts(tp_page.to_json)

      target.puts(&#x27;]};&#x27;)
      target.close
    end

### Comments

