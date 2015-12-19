# [How to get offline documentation with the command docs?](https://github.com/jekyll/jekyll-help/issues/166)

> state: **open** opened by: **** on: ****

Hi.

As I understood the command &#x60;docs&#x60; is for viewing *jekyll* documentation offline, but I don&#x27;t know how to use it.

This is what I get for &#x60;jekyll docs&#x60;:

                Source: /var/lib/gems/2.1.0/gems/jekyll-2.4.0/site
       Destination: /var/lib/gems/2.1.0/gems/jekyll-2.4.0/site/_site
      Generating... 
    jekyll 2.4.0 | Error:  No such file or directory @ dir_chdir - /var/lib/gems/2.1.0/gems/jekyll-2.4.0/site/

Note that I&#x27;m new to ruby and gem&#x27;s system, so maybe I&#x27;ve missed something.

Configuration:
- Jekyll 2.4.0
- Xubuntu 14.04





### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/166#issuecomment-58755711) on: **Saturday, October 11, 2014**

Getting the the same on Windows 7:

&#x60;&#x60;&#x60;
$ jekyll docs
Configuration file: none
            Source: c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.4.0/site
       Destination: c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.4.0/site/_site
      Generating...
jekyll 2.4.0 | Error:  No such file or directory - c:/Ruby200-x64/lib/ruby/gems/2.0.0/gems/jekyll-2.4.0/site/
&#x60;&#x60;&#x60;

The &#x60;site&#x60; directory is missing in the specified path. Seems to be broken, huh?
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/166#issuecomment-59082210) on: **Tuesday, October 14, 2014**

@kleinfreund just tried this and also getting this error. It seems that is broken. :cry: 

&#x60;&#x60;&#x60;bash
$ be jekyll docs --port 3000 --host 0.0.0.0 --trace
Configuration file: none
            Source: /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/site
       Destination: /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/site/_site
      Generating...
/Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:136:in &#x60;chdir&#x27;: No such file or directory @ dir_chdir - /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/site/ (Errno::ENOENT)
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:136:in &#x60;read_directories&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:122:in &#x60;read&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:44:in &#x60;process&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/command.rb:28:in &#x60;process_site&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:55:in &#x60;build&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:33:in &#x60;process&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/lib/jekyll/commands/docs.rb:20:in &#x60;block (2 levels) in init_with_program&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;call&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;each&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/mercenary-0.3.4/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/mercenary-0.3.4/lib/mercenary/program.rb:35:in &#x60;go&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/mercenary-0.3.4/lib/mercenary.rb:22:in &#x60;program&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/lib/ruby/gems/2.1.0/gems/jekyll-2.4.0/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/bin/jekyll:23:in &#x60;load&#x27;
	from /Users/albertogg/.rbenv/versions/2.1.2/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;
---
> from: [**redknitin**](https://github.com/jekyll/jekyll-help/issues/166#issuecomment-65911236) on: **Saturday, December 06, 2014**

This seems to be broken in the Jekyll release (I&#x27;ve got 2.5.2 installed through &#x60;gem install jekyll&#x60;) but when running &#x60;jekyll docs&#x60; in the bin directory of the source code that I cloned from git, it works.
