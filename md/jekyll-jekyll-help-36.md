# [static_comments plugin no longer working with Jekyll v2. Why?](https://github.com/jekyll/jekyll-help/issues/36)

> state: **closed** opened by: **** on: ****

Hello,

I’m currently rebuilding my site with Jekyll. I’ve retained the comments that were posted on my previous Movable Type-powered website (using the Movable Type import script) and the static_comments.rb plugin, found here: https://github.com/shigeya/jekyll-static-comments/

This was working _perfectly_ up until v1.5.1, but since the update to v2.0.x, I get the following error (using --trace): 

&#x60;&#x60;&#x60;
Pinova:paulrobertlloyd.dev paulrobertlloyd$ jekyll build --trace
Configuration file: /Users/paulrobertlloyd/Sites/paulrobertlloyd.dev/_config.yml
            Source: /Users/paulrobertlloyd/Sites/paulrobertlloyd.dev
       Destination: /Users/paulrobertlloyd/Sites/paulrobertlloyd.dev/_site
      Generating... 
  Liquid Exception: undefined local variable or method &#x60;relative_path&#x27; for #&lt;StaticComments::StaticComment:0x007f877c1ee3c0&gt; in _includes/sections/comments.html, included in _layouts/post.html
/Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/tags/include.rb:114:in &#x60;rescue in render&#x27;: undefined local variable or method &#x60;relative_path&#x27; for #&lt;StaticComments::StaticComment:0x007f877c1ee3c0&gt; (Jekyll::Tags::IncludeTagError)
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/tags/include.rb:106:in &#x60;render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/block.rb:106:in &#x60;block in render_all&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/block.rb:93:in &#x60;each&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/block.rb:93:in &#x60;render_all&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/tags/if.rb:40:in &#x60;block (2 levels) in render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/tags/if.rb:38:in &#x60;each&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/tags/if.rb:38:in &#x60;block in render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/context.rb:105:in &#x60;stack&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/tags/if.rb:37:in &#x60;render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/block.rb:106:in &#x60;block in render_all&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/block.rb:93:in &#x60;each&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/block.rb:93:in &#x60;render_all&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/block.rb:82:in &#x60;render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/template.rb:124:in &#x60;render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/liquid-2.5.5/lib/liquid/template.rb:132:in &#x60;render!&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/convertible.rb:96:in &#x60;render_liquid&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/convertible.rb:142:in &#x60;render_all_layouts&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/convertible.rb:176:in &#x60;do_layout&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/post.rb:266:in &#x60;render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/site.rb:245:in &#x60;block in render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/site.rb:244:in &#x60;each&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/site.rb:244:in &#x60;render&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/site.rb:43:in &#x60;process&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/command.rb:43:in &#x60;process_site&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/commands/build.rb:46:in &#x60;build&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/commands/build.rb:30:in &#x60;process&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/lib/jekyll/commands/build.rb:17:in &#x60;block (2 levels) in init_with_program&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;call&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;each&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.3/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.3/lib/mercenary/program.rb:35:in &#x60;go&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/mercenary-0.3.3/lib/mercenary.rb:22:in &#x60;program&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/gems/jekyll-2.0.2/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/bin/jekyll:23:in &#x60;load&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/bin/ruby_executable_hooks:15:in &#x60;eval&#x27;
	from /Users/paulrobertlloyd/.rvm/gems/ruby-2.1.0/bin/ruby_executable_hooks:15:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

My source code is available here: https://github.com/paulrobertlloyd/paulrobertlloyd.com 

_(Note that comments are not available on my GitHub repo as they contain user’s email addresses. However, there is a folder called _comments in the root of the project, which contains a folder for each post with comments, and a .markdown file for each comment.)_

What has changed in v2.0 that would cause this issue?

Thanks,

Paul

### Comments

---
> from: [**shigeya**](https://github.com/jekyll/jekyll-help/issues/36#issuecomment-42758896) on: **Saturday, May 10, 2014**

Hi, as I informed to @paulrobertlloyd, it&#x27;s an incompatibility  of my patch ( [the plugin](https://github.com/shigeya/jekyll-static-comments/) ) with Jekyll 2.0. it&#x27;s fixed on the branch, https://github.com/shigeya/jekyll-static-comments/tree/mt_static_comments . Please close this issue.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/36#issuecomment-42758986) on: **Saturday, May 10, 2014**

Thanks @shigeya!
