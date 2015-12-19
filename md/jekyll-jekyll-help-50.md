# [Liquid Exception: cannot load such file -- yajl/2.0/yajl](https://github.com/jekyll/jekyll-help/issues/50)

> state: **closed** opened by: **** on: ****

Hi,
When I run jekyll build command on windows I get this error.
:o(


&#x60;&#x60;&#x60;text
C:\Users\rdquintas\Desktop\jek\zrq&gt;jekyll build
Configuration file: C:/Users/rdquintas/Desktop/jek/zrq/_config.yml
            Source: C:/Users/rdquintas/Desktop/jek/zrq
       Destination: C:/Users/rdquintas/Desktop/jek/zrq/_site
      Generating...
  Liquid Exception: cannot load such file -- yajl/2.0/yajl in _posts/2014-05-21-
welcome-to-jekyll.markdown
C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;: c
annot load such file -- yajl/2.0/yajl (LoadError)
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in
 &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/yajl-ruby-1.1.0-x86-mingw32/lib
/yajl/yajl.rb:2:in &#x60;&lt;top (required)&gt;&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in
 &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in
 &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/yajl-ruby-1.1.0-x86-mingw32/lib
/yajl.rb:1:in &#x60;&lt;top (required)&gt;&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:135:i
n &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:135:i
n &#x60;rescue in require&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:144:i
n &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/pygments.rb-0.5.4/lib/pygments/
popen.rb:3:in &#x60;&lt;top (required)&gt;&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in
 &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in
 &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/pygments.rb-0.5.4/lib/pygments.
rb:1:in &#x60;&lt;top (required)&gt;&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in
 &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in
 &#x60;require&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/tags/hi
ghlight.rb:53:in &#x60;render_pygments&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/tags/hi
ghlight.rb:41:in &#x60;render&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/liquid-2.5.5/lib/liquid/block.r
b:106:in &#x60;block in render_all&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/liquid-2.5.5/lib/liquid/block.r
b:93:in &#x60;each&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/liquid-2.5.5/lib/liquid/block.r
b:93:in &#x60;render_all&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/liquid-2.5.5/lib/liquid/block.r
b:82:in &#x60;render&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/liquid-2.5.5/lib/liquid/templat
e.rb:124:in &#x60;render&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/liquid-2.5.5/lib/liquid/templat
e.rb:132:in &#x60;render!&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/convert
ible.rb:96:in &#x60;render_liquid&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/convert
ible.rb:170:in &#x60;do_layout&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/post.rb
:266:in &#x60;render&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb
:245:in &#x60;block in render&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb
:244:in &#x60;each&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb
:244:in &#x60;render&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/site.rb
:43:in &#x60;process&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/command
.rb:43:in &#x60;process_site&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/command
s/build.rb:46:in &#x60;build&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/command
s/build.rb:30:in &#x60;process&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/command
s/build.rb:17:in &#x60;block (2 levels) in init_with_program&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/c
ommand.rb:220:in &#x60;call&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/c
ommand.rb:220:in &#x60;block in execute&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/c
ommand.rb:220:in &#x60;each&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/c
ommand.rb:220:in &#x60;execute&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/p
rogram.rb:35:in &#x60;go&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary.r
b:22:in &#x60;program&#x27;
        from C:/Ruby200/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/bin/jekyll:18:in &#x60;
&lt;top (required)&gt;&#x27;
        from C:/Ruby200/bin/jekyll:23:in &#x60;load&#x27;
        from C:/Ruby200/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;

C:\Users\rdquintas\Desktop\jek\zrq&gt;
&#x60;&#x60;&#x60;
























### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-43702805) on: **Tuesday, May 20, 2014**

Could you please put the error in [Fenced code blocks](https://help.github.com/articles/github-flavored-markdown#fenced-code-blocks) so it&#x27;s easier to read! *add highlight if posible*.
---
> from: [**rdquintas**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-44036625) on: **Friday, May 23, 2014**

Hi Albert,
thank you for pointing that out.
I was not aware of that feature.

Nevertheless, tried to do as you suggested, but markdown preview gave similar graphic output.

Neverhtess, the error is related to this line only:

&#x60;&#x60;&#x60;
  Liquid Exception: cannot load such file -- yajl/2.0/yajl in _posts/2014-05-21-
welcome-to-jekyll.markdown
C:/Ruby200/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;: c
annot load such file -- yajl/2.0/yajl (LoadError)
&#x60;&#x60;&#x60;
Guess this is related with the fact that I&#x27;m running these versions on windows:
**Compass 1.0.0.alpha.19**
**ruby 2.0.0p481 (2014-05-08) [i386-mingw32]**

If I downgrade Ruby it seems to work.
But I do need to have ruby running on 2.0, I&#x27;m afraid.

Cheers.

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-44038260) on: **Friday, May 23, 2014**

It seems to be a problem with &#x60;yajl&#x60; not being found anywhere, it may be a path problem, do you have a &#x60;Gemfile&#x60;? Could you try to see if &#x60;yajl-ruby&#x60; is installed on your system?
---
> from: [**shellphon**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-44536961) on: **Thursday, May 29, 2014**

@rdquintas hello, i have the same error in my machine. I guess maybe your local python is 3.x (the same to mine). And the pygments.rb needs to run with python 2.x (says the ReadMe.md of pygments.rb) , I&#x27;m afraid.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-44578917) on: **Thursday, May 29, 2014**

@shellphon if you think thats the case can you try using Rouge? &#x60;gem install rouge&#x60; and set it up [in your _config.yml](https://github.com/albertogg/albertogg.github.com/blob/master/_config.yml#L31)
---
> from: [**winghouchan**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-47430823) on: **Saturday, June 28, 2014**

Hello all!

I started experiencing the same problem once I upgraded to Ruby 2.0.0 . I upgraded following a suggestion that it would help solve the original problem of this Liquid Exception:

&#x60;&#x60;&#x60;
Liquid Exception: No such file or directory - python C:/Ruby193/lib/ruby/gems/1.9.1/gems/pygments.rb-0.5.0/lib/pygments/mentos.py in _posts/2014/06/28-Name-of-post.md
&#x60;&#x60;&#x60;

Details of that issue are here: https://github.com/jekyll/jekyll/issues/2551

I also did a search of the Liquid Exception referenced in the title &#x60;Liquid Exception: cannot load such file -- yajl/2.0/yajl&#x60; and found this issue brianmario/yajl-ruby#116, which you may want to read.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-55625569) on: **Monday, September 15, 2014**

Going to close this as it seems to be resolved.
---
> from: [**gazeldx**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-64940168) on: **Friday, November 28, 2014**

    gem install rouge
works for me. Thanks.
---
> from: [**XhmikosR**](https://github.com/jekyll/jekyll-help/issues/50#issuecomment-68726896) on: **Monday, January 05, 2015**

This is definitely an issue and I&#x27;m still trying to find a sensible fix.

https://ci.appveyor.com/project/XhmikosR/pages-gem/build/28/job/j75kggkac3fh2l2k
