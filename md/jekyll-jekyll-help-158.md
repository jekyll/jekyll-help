# [Cannot configure jekyll with github-pages](https://github.com/jekyll/jekyll-help/issues/158)

> state: **closed** opened by: **** on: ****

I installed github-pages, and attempted to run Jekyll with Bundler in a way that matches the github pages build server, but I keep getting this error.

![jekyllerror](https://cloud.githubusercontent.com/assets/6391122/4397892/7a3f3452-444a-11e4-8175-793d82b4d322.jpg)

I&#x27;ve seen a few issues that appear related, but nothing I&#x27;ve tried works. Ideas anyone?





### Comments

---
> from: [**bdgrider**](https://github.com/jekyll/jekyll-help/issues/158#issuecomment-56758854) on: **Wednesday, September 24, 2014**

The full output:

&#x60;&#x60;&#x60;
$ bundle exec jekyll serve
Configuration file: /cygdrive/c/jekyllTest/_config.yml
            Source: /cygdrive/c/jekyllTest
       Destination: /cygdrive/c/jekyllTest/_site
      Generating...
  Liquid Exception: cannot load such file -- rouge in _posts/2014-09-23-Random-photo-script.md/#excerpt
/usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/tags/highlight.rb:102:in &#x60;require&#x27;: cannot load such file -- rouge (LoadError)
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/tags/highlight.rb:102:in &#x60;render_rouge&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/tags/highlight.rb:53:in &#x60;render&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/liquid-2.6.1/lib/liquid/block.rb:109:in &#x60;block in render_all&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/liquid-2.6.1/lib/liquid/block.rb:96:in &#x60;each&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/liquid-2.6.1/lib/liquid/block.rb:96:in &#x60;render_all&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/liquid-2.6.1/lib/liquid/block.rb:82:in &#x60;render&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/liquid-2.6.1/lib/liquid/template.rb:128:in &#x60;render&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/liquid-2.6.1/lib/liquid/template.rb:138:in &#x60;render!&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/convertible.rb:105:in &#x60;render_liquid&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/convertible.rb:234:in &#x60;do_layout&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/post.rb:258:in &#x60;render&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/site.rb:263:in &#x60;block in render&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/site.rb:262:in &#x60;each&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/site.rb:262:in &#x60;render&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/site.rb:45:in &#x60;process&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/command.rb:28:in &#x60;process_site&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/commands/build.rb:55:in &#x60;build&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/commands/build.rb:33:in &#x60;process&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/lib/jekyll/commands/serve.rb:24:in &#x60;block (2 levels) in init_with_program&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;call&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;each&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/mercenary-0.3.4/lib/mercenary/program.rb:35:in &#x60;go&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/mercenary-0.3.4/lib/mercenary.rb:22:in &#x60;program&#x27;
        from /usr/lib/ruby/gems/1.9.1/gems/jekyll-2.3.0/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
        from /usr/bin/jekyll:23:in &#x60;load&#x27;
        from /usr/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;
---
> from: [**fabianrbz**](https://github.com/jekyll/jekyll-help/issues/158#issuecomment-57095415) on: **Sunday, September 28, 2014**

@bdgrider do you have rouge installed?
try &#x60;&#x60;&#x60;gem install rouge&#x60;&#x60;&#x60; and re-run jekyll
---
> from: [**bdgrider**](https://github.com/jekyll/jekyll-help/issues/158#issuecomment-57112350) on: **Sunday, September 28, 2014**

I had it installed. Ended up using pygments instead and got it to work. Thanks though.
