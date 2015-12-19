# [Can&#x27;t start server with jekyll 2.0.3 on windows.](https://github.com/jekyll/jekyll-help/issues/64)

> state: **closed** opened by: **** on: ****

When I type:

    D:\Fernando\prog\fernandobasso.github.io&gt;jekyll server --watch
    Configuration file: D:/Fernando/prog/fernandobasso.github.io/_config.yml
                Source: D:/Fernando/prog/fernandobasso.github.io
           Destination: D:/Fernando/prog/fernandobasso.github.io/_site
          Generating...
                        done.
    jekyll 2.0.3 | Error:  undefined method &#x60;_log&#x27; for Listen::Adapter::Windows:Class

Google didn&#x27;t help this time.

ruby 1.9.3p545 (2014-02-24) [i386-mingw32]

gem 1.8.28

### Comments

---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-44989775) on: **Tuesday, June 03, 2014**

What happens if you type &#x60;serve&#x60;, without the &#x60;r&#x60; at the end?
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-44991418) on: **Tuesday, June 03, 2014**

Same stuff. If in _config.yml I comment out the line

    highlighter: true

then it works. So, that seems to be the problem. The error message gives no clue about this, though.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-44991913) on: **Tuesday, June 03, 2014**

&#x60;highlighter&#x60; should be set to either &#x60;pygments&#x60; or &#x60;rouge&#x60;, not to anything else.
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-44992910) on: **Tuesday, June 03, 2014**

The problem is the &#x60;--watch&#x60; option. That is where the error about the &#x60;_log&#x60; stuff comes from.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-44998813) on: **Tuesday, June 03, 2014**

That probably has something to do with https://github.com/jekyll/jekyll/issues/2320.
Try upgrading to Listen v2.7.6 and trying again.
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45001007) on: **Tuesday, June 03, 2014**

&#x60;&#x60;&#x60;text
D:\Fernando\prog\fernandobasso.github.io&gt;gem install listen -v 2.7.6
Fetching: listen-2.7.6.gem (100%)
Successfully installed listen-2.7.6
1 gem installed
Installing ri documentation for listen-2.7.6...
Installing RDoc documentation for listen-2.7.6...

D:\Fernando\prog\fernandobasso.github.io&gt;jekyll serve --watch
Configuration file: D:/Fernando/prog/fernandobasso.github.io/_config.yml
            Source: D:/Fernando/prog/fernandobasso.github.io
       Destination: D:/Fernando/prog/fernandobasso.github.io/_site
      Generating...
C:/Ruby193/lib/ruby/gems/1.9.1/gems/posix-spawn-0.3.8/lib/posix/spawn.rb:162: warning: cannot close fd before spawn
&#x27;which&#x27; não é reconhecido como um comando interno
ou externo, um programa operável ou um arquivo em lotes.
←[31m  Liquid Exception: No such file or directory - python C:/Ruby193/lib/ruby/gems/1.9.1/gems/pygments.rb-0.5.4/lib/pygments/mentos.py in _posts/2014-03-10-welcome-to-jekyll.markdown←[0m
                    done.
jekyll 2.0.3 | Error:  undefined method &#x60;_log&#x27; for Listen::Adapter::Windows:Class
&#x60;&#x60;&#x60;
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45001369) on: **Tuesday, June 03, 2014**

Looks like you need to install Python and Pygments in order to run the site. Set &#x60;highlighter: rouge&#x60; in your &#x60;_config.yml&#x60;.

Also, uninstall all other versions of Listen you have installed.
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45001608) on: **Tuesday, June 03, 2014**

Okay, but my _config.yml has just this:

    name: Your New Jekyll Site
    markdown: redcarpet
    #highlighter: pygments

Shouldn&#x27;t jekyll completely ignore highlighter then?

If I run the server without the &#x60;--watch&#x60; option, the server runs:

    D:\Fernando\prog\fernandobasso.github.io&gt;jekyll serve
    Configuration file: D:/Fernando/prog/fernandobasso.github.io/_config.yml
                Source: D:/Fernando/prog/fernandobasso.github.io
           Destination: D:/Fernando/prog/fernandobasso.github.io/_site
          Generating...
    C:/Ruby193/lib/ruby/gems/1.9.1/gems/posix-spawn-0.3.8/lib/posix/spawn.rb:162: warning: cannot close fd before spawn
    &#x27;which&#x27; não é reconhecido como um comando interno
    ou externo, um programa operável ou um arquivo em lotes.
    ←[31m  Liquid Exception: No such file or directory - python C:/Ruby193/lib/ruby/gems/1.9.1/gems/pygments.rb-0.5.4/lib/pygments/mentos.py in _posts/2014-03-10-welcome-to-jekyll.markdown←[0m
                        done.
    Configuration file: D:/Fernando/prog/fernandobasso.github.io/_config.yml
        Server address: http://0.0.0.0:4000/
      Server running... press ctrl-c to stop.

But when I try &#x60;http://127.0.0.1:4000&#x27; I only see a blank page, even though I do have an index.html with some stuff in it.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45008468) on: **Tuesday, June 03, 2014**

&gt; Shouldn&#x27;t jekyll completely ignore highlighter then?

No, Jekyll defaults &#x60;highlighter&#x60; to &#x60;pygments&#x60;. Set it to &#x60;rouge&#x60; or &#x60;null&#x60;.

&gt; I only see a blank page, even though I do have an index.html with some stuff in it.

What is the contents of &#x60;_site/index.html&#x60;?
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45008476) on: **Tuesday, June 03, 2014**

@FernandoBasso It still is using pygments. try what @parkr said.. install rouge &#x60;gem install rouge&#x60; if you haven&#x27;t and test that.
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45009372) on: **Tuesday, June 03, 2014**

With &#x60;rouge&#x60; it works. Still, I had to drop the &#x60;--watch&#x60; option for the server to run.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45010002) on: **Tuesday, June 03, 2014**

@FernandoBasso great. Check that you are using Listen version 2.7.6.
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45010339) on: **Tuesday, June 03, 2014**

Yes.

    gem list | findstr listen
    listen (2.7.6)
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45011349) on: **Tuesday, June 03, 2014**

Can you post the complete stack trace? I don&#x27;t know if it would work but, &#x60;jekyll serve --watch --trace&#x60;
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45011636) on: **Tuesday, June 03, 2014**

EDIT: This is what SOLVED the problem:

Uninstalling &#x60;listen 2.7.6&#x60; and installing &#x60;listen 2.7.5&#x60; seems to have solved the issue.

This [stackoverflow post](http://stackoverflow.com/questions/23998719/running-guard-throwing-error) lead me to give it a try.
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45012164) on: **Tuesday, June 03, 2014**

I uninstalled listen 2.7.5 and installed 2.7.6 to get you the stack trace. Here it is:

&#x60;&#x60;&#x60;
Removing listen
Successfully uninstalled listen-2.7.5

D:\Fernando\prog\fernandobasso.github.io&gt;gem install listen
Fetching: listen-2.7.6.gem (100%)
Successfully installed listen-2.7.6
Parsing documentation for listen-2.7.6
Installing ri documentation for listen-2.7.6
1 gem installed

D:\Fernando\prog\fernandobasso.github.io&gt;jekyll serve --watch --trace
Configuration file: D:/Fernando/prog/fernandobasso.github.io/_config.yml
            Source: D:/Fernando/prog/fernandobasso.github.io
       Destination: D:/Fernando/prog/fernandobasso.github.io/_site
      Generating...
                    done.
C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/adapter/windows.rb:21:in &#x60;rescue in usable?&#x27;: undefined method &#x60;_log&#x27; for Listen::Adapter::Windows:Class (NoMethodError)
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/adapter/windows.rb:17:in &#x60;usable?&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/adapter.rb:32:in &#x60;each&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/adapter.rb:32:in &#x60;detect&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/adapter.rb:32:in &#x60;_usable_adapter_class&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/adapter.rb:20:in &#x60;select&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/listener.rb:252:in &#x60;_adapter_class&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/listener.rb:208:in &#x60;_init_actors&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/listener.rb:72:in &#x60;block in &lt;class:Listener&gt;&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/celluloid-0.15.2/lib/celluloid/fsm.rb:175:in &#x60;instance_eval&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/celluloid-0.15.2/lib/celluloid/fsm.rb:175:in &#x60;call&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/celluloid-0.15.2/lib/celluloid/fsm.rb:127:in &#x60;transition_with_callbacks!&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/celluloid-0.15.2/lib/celluloid/fsm.rb:95:in &#x60;transition&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/listen-2.7.6/lib/listen/listener.rb:85:in &#x60;start&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/commands/build.rb:87:in &#x60;watch&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/commands/build.rb:31:in &#x60;process&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/lib/jekyll/commands/serve.rb:23:in &#x60;block (2 levels) in init_with_program&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;call&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;each&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary/program.rb:35:in &#x60;go&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/mercenary-0.3.3/lib/mercenary.rb:22:in &#x60;program&#x27;
        from C:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.0.3/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
        from C:/Ruby200-x64/bin/jekyll:23:in &#x60;load&#x27;
        from C:/Ruby200-x64/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;

D:\Fernando\prog\fernandobasso.github.io&gt;
&#x60;&#x60;&#x60;
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45080971) on: **Wednesday, June 04, 2014**

Well, I am stuck on this windows machine at work anyway. Thanks for all the help and patience. I really appreciate it.

PS: Should I close this?
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45120208) on: **Wednesday, June 04, 2014**

No problem, at least you got a solution. I will reference this issue in the guard/listen repo to let them know of this problem. It seems that they touched some Windows stuff in the 2.7.6 which may be causing this issue.
---
> from: [**seanpoulter**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45179079) on: **Wednesday, June 04, 2014**

Unfortunately I can confirm I get the same stack trace as @FernandoBasso [above](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45012164) using _ruby 2.0.0p481 (2014-05-08) [i386-mingw32]_, _jekyll-2.0.3_, and _listen (2.7.6)_.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45937269) on: **Thursday, June 12, 2014**

Hey, @seanpoulter and @FernandoBasso could you try running jekyll with listen *(2.7.7)* and see if it stills errors out?
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-45937461) on: **Thursday, June 12, 2014**

Not for me. I still get that error:

    jekyll 2.0.3 | Error:  undefined method &#x60;_log&#x27; for Listen::Adapter::Windows:Class
---
> from: [**rohaan911**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-46084648) on: **Saturday, June 14, 2014**

Running &#x60;gem update listen&#x60; to update my &#x27;listen&#x27; ruby gem to version 2.7.8 (released on 12th June, 2014) seemed to fix this bug.

Also, to run Jekyll server in watch mode for versions 1.0+ , I use &#x60;jekyll serve --watch&#x60;. According to the documentation, the &#x60;--server&#x60; command now seems to be obsolete.

Hope that helps.
---
> from: [**FernandoBasso**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-46229822) on: **Monday, June 16, 2014**

Updated to &#x60;listen 2.7.8&#x60; and now I am able to start jekyll server with the &#x60;--watch&#x60; option. Everything is working fine now. Thanks to everybody for the support.
---
> from: [**alcidesqueiroz**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-46522056) on: **Wednesday, June 18, 2014**

I&#x27;ve just updated &#x60;listen&#x60; to 2.7.8 and that did the job! Now, &#x60;jekyll serve --watch&#x60; is working perfectly!
---
> from: [**rsb12php**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-86538784) on: **Thursday, March 26, 2015**

Well, i tried all! But, fail :(
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-86629984) on: **Thursday, March 26, 2015**

@rsb12php You&#x27;re using &#x60;bundle exec jekyll serve&#x60; on the latest version of Jekyll?
---
> from: [**rsb12php**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-87773957) on: **Monday, March 30, 2015**

@parkr it&#x27;s all right now, but thak you
---
> from: [**rsb12php**](https://github.com/jekyll/jekyll-help/issues/64#issuecomment-87776361) on: **Monday, March 30, 2015**

@parkr I using Windows 8
jekyll - Auto regeneration is not working!
Do you can help me?
