# [jekyll 2.5.2 | Error:  undefined method &#x60;in_dest_dir&#x27; for nil:NilClass](https://github.com/jekyll/jekyll-help/issues/208)

> state: **closed** opened by: **** on: ****

Our website is built on Jekyll, originally version 1.x.

I just got a new laptop (MacBook Pro with Yosemite), so installed Jekyll from fresh. Unfortunately when attempting to build our site, I get the following error:

    jekyll 2.5.2 | Error:  undefined method &#x60;in_dest_dir&#x27; for nil:NilClass

Any suggestions? Is there something I need to do to migrate, or something I can do to diagnose the problem?

Many thanks!

### Comments

---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/208#issuecomment-66232995) on: **Monday, December 08, 2014**

What does it spit out when you run Jekyll with the &#x60;&#x60;&#x60;--trace&#x60;&#x60;&#x60; command?
---
> from: [**joethephish**](https://github.com/jekyll/jekyll-help/issues/208#issuecomment-66252727) on: **Tuesday, December 09, 2014**

      Generating... 
    /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/static_file.rb:40:in &#x60;destination&#x27;: undefined method &#x60;in_dest_dir&#x27; for nil:NilClass (NoMethodError)
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/cleaner.rb:43:in &#x60;block in new_files&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/site.rb:477:in &#x60;block (2 levels) in each_site_file&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/site.rb:476:in &#x60;each&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/site.rb:476:in &#x60;block in each_site_file&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/site.rb:475:in &#x60;each&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/site.rb:475:in &#x60;each_site_file&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/cleaner.rb:43:in &#x60;new_files&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/cleaner.rb:24:in &#x60;obsolete_files&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/cleaner.rb:15:in &#x60;cleanup!&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/site.rb:308:in &#x60;cleanup&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/site.rb:52:in &#x60;process&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/command.rb:28:in &#x60;process_site&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/commands/build.rb:56:in &#x60;build&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/commands/build.rb:34:in &#x60;process&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/lib/jekyll/commands/serve.rb:26:in &#x60;block (2 levels) in init_with_program&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;call&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;each&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/mercenary-0.3.5/lib/mercenary/program.rb:42:in &#x60;go&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/mercenary-0.3.5/lib/mercenary.rb:19:in &#x60;program&#x27;
    from /Library/Ruby/Gems/2.0.0/gems/jekyll-2.5.2/bin/jekyll:20:in &#x60;&lt;top (required)&gt;&#x27;
    from /usr/bin/jekyll:23:in &#x60;load&#x27;
    from /usr/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
---
> from: [**joethephish**](https://github.com/jekyll/jekyll-help/issues/208#issuecomment-66277340) on: **Tuesday, December 09, 2014**

Ah, I think I may know what it is. I forgot that I had hacked together a plugin (warning, it was basically my first foray into Ruby). Here is the plugin... perhaps it&#x27;s now incompatible with Jekyll 2.x, or perhaps I didn&#x27;t write it well in the first place? :)

&lt;https://gist.github.com/joethephish/09d12d80bd0a98c16670&gt;
---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/208#issuecomment-66367506) on: **Tuesday, December 09, 2014**

Huh, yeah, I&#x27;m a bit of a loss as to what gives this error. Did you try without your plugin? Since 2.x is a breaking change from 1.x it might be. How did you upgrade Jekyll by the way/do you have the source available anywhere so that I could take a look myself?
---
> from: [**joethephish**](https://github.com/jekyll/jekyll-help/issues/208#issuecomment-66429854) on: **Wednesday, December 10, 2014**

Thanks for your help! Yeah, it was definitely my plugin. Looks like it was something to do with the fact that StaticFile objects now require the &#x60;site&#x60; property to be set in their constructor? So, when creating my StaticFile subclass, I passed it the site, and it was happy again.
---
> from: [**johnpflyer74**](https://github.com/jekyll/jekyll-help/issues/208#issuecomment-71947125) on: **Wednesday, January 28, 2015**

I&#x27;m new to **GitHub** and **jekyll**, so please forgive any rookie behavior on my part. 

I am testing out a plugin ( [&#x60;jekyll-thumbnailify&#x60;](https://github.com/10io/jekyll-thumbnailify) ) and I am receiving the exact same error as **joethephish**. The only difference is that I have slightly newer versions of **gems** and **jekyll**:

&#x60;&#x60;&#x60;
Generating...
/usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/static_file.rb:40:in &#x60;destination&#x27;: undefined method &#x60;in_dest_dir&#x27; for nil:NilClass (NoMethodError)
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/cleaner.rb:43:in &#x60;block in new_files&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/site.rb:477:in &#x60;block (2 levels) in each_site_file&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/site.rb:476:in &#x60;each&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/site.rb:476:in &#x60;block in each_site_file&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/site.rb:475:in &#x60;each&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/site.rb:475:in &#x60;each_site_file&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/cleaner.rb:43:in &#x60;new_files&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/cleaner.rb:24:in &#x60;obsolete_files&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/cleaner.rb:15:in &#x60;cleanup!&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/site.rb:308:in &#x60;cleanup&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/site.rb:52:in &#x60;process&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/command.rb:28:in &#x60;process_site&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/commands/build.rb:56:in &#x60;build&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/commands/build.rb:34:in &#x60;process&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/lib/jekyll/commands/serve.rb:26:in &#x60;block (2 levels) in init_with_program&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;call&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;each&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary/program.rb:42:in &#x60;go&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/mercenary-0.3.5/lib/mercenary.rb:19:in &#x60;program&#x27;
	from /usr/local/lib/ruby/gems/2.2.0/gems/jekyll-2.5.3/bin/jekyll:20:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/local/bin/jekyll:23:in &#x60;load&#x27;
	from /usr/local/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

I see that **joethefish** came to the conclusion that &quot;...  it was something to do with the fact that StaticFile objects now require the &#x60;site&#x60; property to be set in their constructor.&quot; I poked around to try and figure out how to do that, but I haven&#x27;t had any luck.

Could one of you guys help me understand how to pass the &#x60;site&#x60; to the StaticFile subclass of the plugin?

I&#x27;m not even sure if my team is going to use &#x60;jekyll-thumbnailify&#x60;, but I&#x27;ve spent so much time trying to figure out how to fix the issue that I need to see it through to the end.

Oh, and I don&#x27;t know if this is relevant, but I used **Home Brew** to install &#x60;imagemagick&#x60; today, so I believe I have the latest version.
---
> from: [**joethephish**](https://github.com/jekyll/jekyll-help/issues/208#issuecomment-71992282) on: **Thursday, January 29, 2015**

Here&#x27;s what I did (albeit in my own hacked plugin that is now a bit different from &#x60;jekyll-thumbnailify&#x60;):

Line 65: Add site as first parameter:

    static_file = ImageStaticFile.new(site, thumbnail, image_folder, thumbnail_name)

Line 122: Take this new parameter to the &#x60;ImageStaticFile&#x60; subclass (and also &#x60;@base&#x60;, not sure what that&#x27;s for):

    def initialize(site, image_src, image_dest_dir, image_dest_name)
        @site = site
        @base = image_src

**Warning: I have near-zero experience with Ruby, and I do not understand much of the way Jekyll works, so this may be a terrible awful hack, and I wouldn&#x27;t be surprised if it didn&#x27;t work ;-) Use at your own risk!**

That said, I hope it helps you!
---
> from: [**johnpflyer74**](https://github.com/jekyll/jekyll-help/issues/208#issuecomment-72039477) on: **Thursday, January 29, 2015**

That worked! Thanks for the help.

Hack or not, it did the trick. What I&#x27;m learning is that it is more important to understand the principals of coding and less so to understand the language. 

Cheers.
