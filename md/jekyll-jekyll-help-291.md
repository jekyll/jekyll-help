# [problems with &#x27;bundle exec jekyll serve&#x27;](https://github.com/jekyll/jekyll-help/issues/291)

> state: **closed** opened by: **** on: ****

Hi, I&#x27;ve been following the guide to using Jekyll with bundler (https://help.github.com/articles/using-jekyll-with-pages/) and I installed ruby/devkit/bundler and updated the gems however when I try to run the command &#x27;bundle exec jekyll serve&#x27; I get this error:

&#x60;&#x60;&#x60;
D:\Coding\delekru.github.io&gt;bundle exec jekyll serve
Configuration file: none
            Source: D:/Coding/delekru.github.io
       Destination: D:/Coding/delekru.github.io/_site
      Generating...
                    done.
D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/hitimes-1.2.2-x86-mingw32/lib/hitimes.rb:37:in &#x60;require&#x27;: cannot load such file -- hitimes/hitimes (LoadError)
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/hitimes-1.2.2-x86-mingw32/lib/hitimes.rb:37:in &#x60;rescue in &lt;top (required)&gt;&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/hitimes-1.2.2-x86-mingw32/lib/hitimes.rb:32:in &#x60;&lt;top (required)&gt;&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/timers-4.0.1/lib/timers/group.rb:4:in &#x60;require&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/timers-4.0.1/lib/timers/group.rb:4:in &#x60;&lt;top (required)&gt;&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/timers-4.0.1/lib/timers.rb:4:in &#x60;require&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/timers-4.0.1/lib/timers.rb:4:in &#x60;&lt;top (required)&gt;&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/celluloid-0.16.0/lib/celluloid/receivers.rb:3:in &#x60;require&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/celluloid-0.16.0/lib/celluloid/receivers.rb:3:in &#x60;&lt;top (required)&gt;&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/celluloid-0.16.0/lib/celluloid.rb:475:in &#x60;require&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/celluloid-0.16.0/lib/celluloid.rb:475:in &#x60;&lt;top (required)&gt;&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/listen-2.10.0/lib/listen.rb:1:in &#x60;require&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/listen-2.10.0/lib/listen.rb:1:in &#x60;&lt;top (required)&gt;&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/jekyll-watch-1.2.1/lib/jekyll/watcher.rb:26:in &#x60;require&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/jekyll-watch-1.2.1/lib/jekyll/watcher.rb:26:in &#x60;build_listener&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/jekyll-watch-1.2.1/lib/jekyll/watcher.rb:7:in &#x60;watch&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:67:in &#x60;watch&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:37:in &#x60;process&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/jekyll-2.4.0/lib/jekyll/commands/serve.rb:25:in &#x60;block (2 levels) in init_with_program&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;call&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;each&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/program.rb:42:in &#x60;go&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary.rb:19:in &#x60;program&#x27;
        from D:/Coding/Ruby22/lib/ruby/gems/2.2.0/gems/jekyll-2.4.0/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
        from D:/Coding/Ruby22/bin/jekyll:23:in &#x60;load&#x27;
        from D:/Coding/Ruby22/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;
I&#x27;ve tried looking at previous posts and &#x60;gem cleanup&#x60; and &#x60;bundle update&#x60; but I can&#x27;t figure it out.

### Comments

