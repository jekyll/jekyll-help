# [Trouble with permalinks - :title and :categories not being expanded correctly](https://github.com/jekyll/jekyll-help/issues/66)

> state: **closed** opened by: **** on: ****

&#x60;&#x60;&#x60;
--&gt; jekyll build --trace
Configuration file: D:/Fernando/prog/fernandobasso.github.io/_config.yml
            Source: D:/Fernando/prog/fernandobasso.github.io
       Destination: D:/Fernando/prog/fernandobasso.github.io/_site
      Generating...
C:/Ruby200-x64/lib/ruby/2.0.0/fileutils.rb:245:in &#x60;mkdir&#x27;: Invalid argument - D:/Fernando/prog/fernandobasso.github.io/_site/:title (Errno::EINVAL)
        from C:/Ruby200-x64/lib/ruby/2.0.0/fileutils.rb:245:in &#x60;fu_mkdir&#x27;
        from C:/Ruby200-x64/lib/ruby/2.0.0/fileutils.rb:219:in &#x60;block (2 levels) in mkdir_p&#x27;
        from C:/Ruby200-x64/lib/ruby/2.0.0/fileutils.rb:217:in &#x60;reverse_each&#x27;
        from C:/Ruby200-x64/lib/ruby/2.0.0/fileutils.rb:217:in &#x60;block in mkdir_p&#x27;
        from C:/Ruby200-x64/lib/ruby/2.0.0/fileutils.rb:203:in &#x60;each&#x27;
        from C:/Ruby200-x64/lib/ruby/2.0.0/fileutils.rb:203:in &#x60;mkdir_p&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/convertible.rb:186:in &#x60;write&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb:262:in &#x60;block in write&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb:420:in &#x60;block (2 levels) in each_site_file&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb:419:in &#x60;each&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb:419:in &#x60;block in each_site_file&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb:418:in &#x60;each&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb:418:in &#x60;each_site_file&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb:262:in &#x60;write&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb:45:in &#x60;process&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/command.rb:43:in &#x60;process_site&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/commands/build.rb:46:in &#x60;build&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/commands/build.rb:30:in &#x60;process&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/commands/build.rb:17:in &#x60;block (2 levels) in init_with_program&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;call&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;each&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/program.rb:35:in &#x60;go&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary.rb:22:in &#x60;program&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
        from C:/Ruby200-x64/bin/jekyll:23:in &#x60;load&#x27;
        from C:/Ruby200-x64/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

That happens when I try to set permalinks with &#x60;:title&#x60; or &#x60;:categories&#x60; for instance. It seems jekyll is trying to literally run &#x60;mkdir&#x60; on those literal strings.

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/66#issuecomment-45121143) on: **Wednesday, June 04, 2014**

I&#x27;m clueless here. How did you installed your Ruby?
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/66#issuecomment-45133480) on: **Wednesday, June 04, 2014**

Neither for [&#x60;Page&#x60;](https://github.com/jekyll/jekyll/blob/master/lib/jekyll/page.rb#L92-L98), [&#x60;Post&#x60;](https://github.com/jekyll/jekyll/blob/master/lib/jekyll/post.rb#L218-L232), nor [&#x60;Document&#x60;](https://github.com/jekyll/jekyll/blob/master/lib/jekyll/document.rb#L110-L116) does &#x60;:title&#x60; exist. Only &#x60;Post&#x60;s can have categories.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/issues/66#issuecomment-45166169) on: **Wednesday, June 04, 2014**

Via IRC discussion(s), https://github.com/jekyll/jekyll/issues/2475 became relevant.

This thread can probably be closed. See https://github.com/jekyll/jekyll/issues/2475.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/66#issuecomment-55625943) on: **Monday, September 15, 2014**

:boom: 
