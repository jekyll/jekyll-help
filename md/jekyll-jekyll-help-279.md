# [site.tags does not list tags for pages](https://github.com/jekyll/jekyll-help/issues/279)

> state: **open** opened by: **** on: ****

If you add tags to a regular page (so outside _posts directory), they don&#x27;t get listed with site.tags. Any reason for this behavior?

Related source code: https://github.com/jekyll/jekyll/blob/v2.5.3/lib/jekyll/site.rb#L341

(And if you want the tags from the regular pages to be listed along posts tags, you can&#x27;t because you cannot join arrays in Liquid template: https://github.com/Shopify/liquid/issues/427)

### Comments

---
> from: [**tkrotoff**](https://github.com/jekyll/jekyll-help/issues/279#issuecomment-77065015) on: **Tuesday, March 03, 2015**

I manage to make this work by creating a [generator](http://jekyllrb.com/docs/plugins/#generators):

&#x60;&#x60;&#x60;Ruby
module Jekyll

  # The list of pages and posts for each tag
  class TagsGenerator &lt; Generator

    def generate(site)
      tags_html_page = site.pages.detect { |page| page.name == &#x27;tags.html&#x27; }
      tags_html_page.data[&#x27;all_tags&#x27;] = prepend_page_tags_to_post_tags(site)
    end

    # Returns a hash of pages and posts indexed by tags.
    #
    # Example:
    #   {
    #     &#x27;tech&#x27; =&gt; [&lt;Page A&gt;, &lt;Page B&gt;, &lt;Post A&gt;, &lt;Post B&gt;],
    #     &#x27;ruby&#x27; =&gt; [&lt;Post B&gt;]
    #   }
    #
    # By default Jekyll does not return the list of pages for each tag, only the posts
    def prepend_page_tags_to_post_tags(site)
      # Already given by Jekyll: the list of posts for each tag
      all_tags = site.post_attr_hash(&#x27;tags&#x27;)

      # Trick: loop over the pages in reverse order because they are prepend to the hash
      # =&gt; end up with the original order
      site.pages.reverse.each do |page|
        page_tags = page.data[&#x27;tags&#x27;]
        unless page_tags.nil?
          page_tags.each do |tag|
            # Pages in front of the list
            all_tags[tag].unshift(page)
          end
        end
      end

      return all_tags
    end
  end
end
&#x60;&#x60;&#x60;

The generator returns the list of pages and posts for each tag (instead of just the list of posts for each tag).
How to use: inside page &#x60;tags.html&#x60; (the page name matters), use &#x60;{{ page.all_tags }}&#x60;.

See commit https://github.com/tkrotoff/osteo15.com/commit/45f0301ee1bf7d7738d3ed27ea8160be38c0f76f

For my first attempt, I tried to extend [Jekyll::Site class](https://github.com/jekyll/jekyll/blob/v2.5.3/lib/jekyll/site.rb) by defining a new method &#x60;all_tags&#x60; without success. Could not find any resources on the subject beside [this unanswered question on Stack Overflow](http://stackoverflow.com/questions/26853529).
