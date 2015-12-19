# [Error while trying to install pre-release](https://github.com/jekyll/jekyll-help/issues/9)

> state: **closed** opened by: **** on: ****

While running &#x60;&#x60;&#x60;sudo gem install jekyll --pre&#x60;&#x60;&#x60; on mac, I get the following:

&#x60;&#x60;&#x60;
Building native extensions.  This could take a while...
ERROR:  Error installing jekyll:
	ERROR: Failed to build gem native extension.

    /System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/bin/ruby extconf.rb
checking for rb_thread_blocking_region()... yes
checking for sys/select.h... yes
checking for poll.h... yes
checking for sys/epoll.h... no
checking for sys/event.h... yes
checking for sys/queue.h... yes
checking for port.h... no
checking for sys/resource.h... yes
creating Makefile

make &quot;DESTDIR=&quot;
compiling monitor.c
In file included from monitor.c:6:
In file included from ./nio4r.h:10:
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.9.sdk/System/Library/Frameworks/Ruby.framework/Versions/2.0/usr/include/ruby-2.0.0/ruby/backward/rubyio.h:2:2: warning: use &quot;ruby/io.h&quot; instead of &quot;rubyio.h&quot; [-W#warnings]
#warning use &quot;ruby/io.h&quot; instead of &quot;rubyio.h&quot;
&#x60;&#x60;&#x60;

Here is the full output https://gist.github.com/justinmusgrove/9946007. I am pretty new to ruby so if you need more details or want me to try something, let me know. Thanks!

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/9#issuecomment-39402701) on: **Wednesday, April 02, 2014**

It&#x27;s highly encouraged to use a Ruby version manager like [rbenv](https://github.com/sstephenson/rbenv) to run Ruby projects on your machine in lieu of &quot;System&quot; Ruby (the one that comes with your Mac). See if you can get rbenv set up and try installing Jekyll again. If you need any more help, feel free to let us know!
---
> from: [**justinmusgrove**](https://github.com/jekyll/jekyll-help/issues/9#issuecomment-39406179) on: **Wednesday, April 02, 2014**

Thanks @troyswanson 

K was able to install rbenv, running jekyll 1.5 worked but after installing jekyll-2.0.0.alpha.2 received the following:
&#x60;&#x60;&#x60;
jekyll serve
/Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;: cannot load such file -- rouge (LoadError)
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll/converters/markdown/redcarpet_parser.rb:42:in &#x60;&lt;module:WithRouge&gt;&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll/converters/markdown/redcarpet_parser.rb:41:in &#x60;&lt;class:RedcarpetParser&gt;&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll/converters/markdown/redcarpet_parser.rb:4:in &#x60;&lt;class:Markdown&gt;&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll/converters/markdown/redcarpet_parser.rb:3:in &#x60;&lt;module:Converters&gt;&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll/converters/markdown/redcarpet_parser.rb:2:in &#x60;&lt;module:Jekyll&gt;&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll/converters/markdown/redcarpet_parser.rb:1:in &#x60;&lt;top (required)&gt;&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll.rb:11:in &#x60;block in require_all&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll.rb:10:in &#x60;each&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll.rb:10:in &#x60;require_all&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/lib/jekyll.rb:61:in &#x60;&lt;top (required)&gt;&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/2.0.0/rubygems/core_ext/kernel_require.rb:45:in &#x60;require&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/lib/ruby/gems/2.0.0/gems/jekyll-2.0.0.alpha.2/bin/jekyll:6:in &#x60;&lt;top (required)&gt;&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/bin/jekyll:23:in &#x60;load&#x27;
	from /Users/justinm/.rbenv/versions/2.0.0-p247/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

Thoughts?


---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/9#issuecomment-39406489) on: **Wednesday, April 02, 2014**

@justinmusgrove it seems like &#x60;rouge&#x60; is missing. try installing it &#x60;gem install rouge&#x60;, or add it to your Gemfile.
---
> from: [**justinmusgrove**](https://github.com/jekyll/jekyll-help/issues/9#issuecomment-39406621) on: **Wednesday, April 02, 2014**

That worked, rookie mistake - thanks for your help...
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/9#issuecomment-39406693) on: **Wednesday, April 02, 2014**

I&#x27;m glad! we are here to help!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/9#issuecomment-39407713) on: **Wednesday, April 02, 2014**

Hey man, there&#x27;s nothing rookie about that mistake. jekyll/jekyll#2055 :wink:
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/9#issuecomment-39407790) on: **Wednesday, April 02, 2014**

PS: This issue should be resolved with the release of alpha 3.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/9#issuecomment-39455617) on: **Thursday, April 03, 2014**

Yep, fixed with the merge of https://github.com/jekyll/jekyll/pull/2189
