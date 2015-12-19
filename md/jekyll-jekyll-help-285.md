# [Integrate Premailer with Jekyll](https://github.com/jekyll/jekyll-help/issues/285)

> state: **closed** opened by: **** on: ****

Hi, I&#x27;m working on a project that has both static web pages and also HTML Email templates. 

As you probably know, HTML Emails require all CSS to be inlined which is a huge pain to manage. Most people use Premailer to handle this automatically - https://github.com/premailer/premailer

How would I go about using premailer with Jekyll for a certain layout?

Would it be possible to use premailer via a plugin?

### Comments

---
> from: [**mattr-**](https://github.com/jekyll/jekyll-help/issues/285#issuecomment-86303020) on: **Wednesday, March 25, 2015**

I can&#x27;t answer your question about using premailer with Jekyll for a
certain layout, but I would imagine that it would be possible to use
premailer via a plugin. However, a lot of that will depend on how premailer
works, so you&#x27;d have to write the plugin and see for yourself how the two
would work together.

---
> from: [**alexphelps**](https://github.com/jekyll/jekyll-help/issues/285#issuecomment-86303966) on: **Wednesday, March 25, 2015**

Thanks Matt. I got a little help on stackoverflow here - http://stackoverflow.com/questions/29246991/premailer-with-jekyll

I&#x27;m testing rewriting the code as a premailer plugin like below, which is working properly. Hopefully this will help anyone else trying to accomplish the same thing.

&#x60;&#x60;&#x60;ruby

require &#x27;premailer&#x27;

module Jekyll
  class Site

    #declare this plugin as safe 
    safe true
    # create an alias for the overriden site::write method
    alias orig_write write

    # we override the site::write method
    def write

      # first call the overriden write method,
      # in order to be sure to find generated css
      orig_write

      # layout name we are looking for in pages/post front matter
      # this can come from site config
      @layout   = &#x27;email-base-test&#x27;

      # this the css that will be inlined
      # this can come from site config
      @css_path = &#x27;assets/css/email.css&#x27;

      # look for generated css file content
      @cssPages = pages.select{ |page|
        page.destination(dest).include? @css_path
      }

      # look again in pages and posts
      # to generate inlinedcontent with premailer
      inlinedcontent = [posts, pages].flatten.select{ |page_or_post|
        page_or_post.data[&#x27;layout&#x27;] == @layout
      }

      inlinedcontent.each do |inlinedcontent|

        inlinedcontent.output = Premailer.new(
            inlinedcontent.output,
            {
                # declare that we pass page html as a string not an url
                :with_html_string =&gt; true,
                # also pass css content as a string
                :css_string       =&gt; @cssPages.join,
            }
        ).to_inline_css;

        # rewrite the newsletter with inlined css
        inlinedcontent.write(dest)

      end
    end
  end
end

&#x60;&#x60;&#x60;
