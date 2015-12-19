# [LoadError](https://github.com/jekyll/jekyll-help/issues/89)

> state: **closed** opened by: **** on: ****

I have ruby-del and rubygems install on my fedora 20, but when I used gem command to install jekyll, I got the following error.

&#x60;&#x60;&#x60;
Successfully installed jekyll-2.1.0
/usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;: cannot load such file -- json/pure (LoadError)
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/share/gems/gems/json-1.7.7/lib/json.rb:60:in &#x60;rescue in &lt;module:JSON&gt;&#x27;
	from /usr/share/gems/gems/json-1.7.7/lib/json.rb:57:in &#x60;&lt;module:JSON&gt;&#x27;
	from /usr/share/gems/gems/json-1.7.7/lib/json.rb:54:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/text.rb:16:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/code_object.rb:28:in &#x60;&lt;class:CodeObject&gt;&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/code_object.rb:26:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/generator/markup.rb:59:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/generator/darkfish.rb:6:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/rdoc.rb:565:in &#x60;&lt;top (required)&gt;&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/core_ext/kernel_require.rb:55:in &#x60;require&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/rubygems_hook.rb:73:in &#x60;load_rdoc&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/rubygems_hook.rb:238:in &#x60;setup&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/rubygems_hook.rb:151:in &#x60;generate&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/rubygems_hook.rb:56:in &#x60;block in generation_hook&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/rubygems_hook.rb:55:in &#x60;each&#x27;
	from /usr/share/gems/gems/rdoc-4.0.1/lib/rdoc/rubygems_hook.rb:55:in &#x60;generation_hook&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/request_set.rb:178:in &#x60;call&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/request_set.rb:178:in &#x60;block in install&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/request_set.rb:177:in &#x60;each&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/request_set.rb:177:in &#x60;install&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/commands/install_command.rb:249:in &#x60;install_gem&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/commands/install_command.rb:291:in &#x60;block in install_gems&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/commands/install_command.rb:287:in &#x60;each&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/commands/install_command.rb:287:in &#x60;install_gems&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/commands/install_command.rb:202:in &#x60;execute&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/command.rb:307:in &#x60;invoke_with_build_args&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/command_manager.rb:167:in &#x60;process_args&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/command_manager.rb:137:in &#x60;run&#x27;
	from /usr/local/share/ruby/site_ruby/rubygems/gem_runner.rb:54:in &#x60;run&#x27;
	from /bin/gem:21:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/89#issuecomment-48094742) on: **Saturday, July 05, 2014**

Try &#x60;gem install json&#x60; or &#x60;gem install json_pure&#x60;. I&#x27;ve never seen this before.
---
> from: [**AJ-Acevedo**](https://github.com/jekyll/jekyll-help/issues/89#issuecomment-48097554) on: **Saturday, July 05, 2014**

What is the output of &#x60;ruby -v; gem -v&#x60; on the command line?

---
> from: [**oblakeerickson**](https://github.com/jekyll/jekyll-help/issues/89#issuecomment-49157204) on: **Wednesday, July 16, 2014**

@imkaywu looks like you got your [site](https://github.com/imkaywu/imkaywu.github.io/commits?author=imkaywu) setup, just wanted to confirm that you got passed this error?
---
> from: [**oblakeerickson**](https://github.com/jekyll/jekyll-help/issues/89#issuecomment-49246990) on: **Wednesday, July 16, 2014**

If this is still an issue @imkaywu we can open it back up again.
