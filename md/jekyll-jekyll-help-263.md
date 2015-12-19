# [more information about gemfiles?](https://github.com/jekyll/jekyll-help/issues/263)

> state: **open** opened by: **** on: ****

I wish the Jekyll documentation had more information about gemfiles. I find this aspect of jekyll really confusing. As I understand it, the gemfile is where you can reference gems specific to your project. Github pages requires a certain gem in order to use github as the repository, and they describe it here: [https://help.github.com/articles/using-jekyll-with-pages/](https://help.github.com/articles/using-jekyll-with-pages/). 

The basic setup, I think, is to do this:

1. Create a file called Gemfile (no extension) that includes this:
    
    source &#x27;https://rubygems.org&#x27;
    gem &#x27;github-pages&#x27;
    
2. Save it in your project directory. 
3. Run &#x60;bundle install&#x60;. 
    A lot of things get installed...

&#x60;&#x60;&#x60;
Fetching gem metadata from https://rubygems.org/..........
Resolving dependencies...
Installing RedCloth 4.2.9
Installing i18n 0.7.0
Installing json 1.8.2
Installing minitest 5.5.1
Installing thread_safe 0.3.4
Installing tzinfo 1.2.2
Installing activesupport 4.2.0
Using blankslate 2.1.2.4
Using hitimes 1.2.2
Using timers 4.0.1
Using celluloid 0.16.0
Using fast-stemmer 1.0.2
Using classifier-reborn 2.0.3
Using coffee-script-source 1.9.0
Using execjs 2.2.2
Using coffee-script 2.3.0
Using colorator 0.1
Using ffi 1.9.6
Installing gemoji 2.1.0
Installing net-dns 0.8.0
Installing public_suffix 1.4.6
Installing github-pages-health-check 0.2.1
Using jekyll-coffeescript 1.0.1
Using jekyll-gist 1.1.0
Using jekyll-paginate 1.1.0
Using sass 3.4.11
Installing jekyll-sass-converter 1.2.0
Using rb-fsevent 0.9.4
Using rb-inotify 0.9.5
Using listen 2.8.5
Using jekyll-watch 1.2.1
Using kramdown 1.5.0
Installing liquid 2.6.1
Using mercenary 0.3.5
Using posix-spawn 0.3.9
Using yajl-ruby 1.2.1
Installing pygments.rb 0.6.1
Installing redcarpet 3.1.2
Using safe_yaml 1.0.4
Using parslet 1.5.0
Using toml 0.1.2
Installing jekyll 2.4.0
Installing mini_portile 0.6.2
Installing nokogiri 1.6.6.2
Installing html-pipeline 1.9.0
Installing jekyll-mentions 0.2.1
Installing jekyll-redirect-from 0.6.2
Installing jekyll-sitemap 0.6.3
Installing jemoji 0.4.0
Installing maruku 0.7.0
Installing rdiscount 2.1.7
Installing terminal-table 1.4.5
Installing github-pages 33
Using bundler 1.7.12
Your bundle is complete!
It was installed into ./_vendor/bundle
Post-install message from html-pipeline:
-------------------------------------------------
Thank you for installing html-pipeline!
You must bundle Filter gem dependencies.
See html-pipeline README.md for more details.
https://github.com/jch/html-pipeline#dependencies
&#x60;&#x60;&#x60;
    
4. Then I run &#x60;jekyll serve&#x60; and get the following errors: 

&#x60;&#x60;&#x60;
tjohnson-mbpr13:resolvev3 tjohnson$ jekyll serve
Configuration file: /Users/tjohnson/projects/resolvev3/_config.yml
/Users/tjohnson/projects/resolvev3/_plugins/jekyll-lunr-js-search.rb:1:in &#x60;require&#x27;: cannot load such file -- jekyll_lunr_js_search/version (LoadError)
	from /Users/tjohnson/projects/resolvev3/_plugins/jekyll-lunr-js-search.rb:1:in &#x60;&lt;top (required)&gt;&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/plugin_manager.rb:58:in &#x60;require&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/plugin_manager.rb:58:in &#x60;block (2 levels) in require_plugin_files&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/plugin_manager.rb:57:in &#x60;each&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/plugin_manager.rb:57:in &#x60;block in require_plugin_files&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/plugin_manager.rb:56:in &#x60;each&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/plugin_manager.rb:56:in &#x60;require_plugin_files&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/plugin_manager.rb:18:in &#x60;conscientious_require&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:74:in &#x60;setup&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/site.rb:36:in &#x60;initialize&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:28:in &#x60;new&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/commands/build.rb:28:in &#x60;process&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/lib/jekyll/commands/serve.rb:25:in &#x60;block (2 levels) in init_with_program&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;call&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;block in execute&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;each&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/mercenary-0.3.5/lib/mercenary/command.rb:220:in &#x60;execute&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/mercenary-0.3.5/lib/mercenary/program.rb:42:in &#x60;go&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/mercenary-0.3.5/lib/mercenary.rb:19:in &#x60;program&#x27;
	from /Users/tjohnson/projects/resolvev3/_vendor/bundle/ruby/2.0.0/gems/jekyll-2.4.0/bin/jekyll:18:in &#x60;&lt;top (required)&gt;&#x27;
	from /Users/tjohnson/.rvm/gems/ruby-2.0.0-p481/bin/jekyll:23:in &#x60;load&#x27;
	from /Users/tjohnson/.rvm/gems/ruby-2.0.0-p481/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
	from /Users/tjohnson/.rvm/gems/ruby-2.0.0-p481/bin/ruby_executable_hooks:15:in &#x60;eval&#x27;
	from /Users/tjohnson/.rvm/gems/ruby-2.0.0-p481/bin/ruby_executable_hooks:15:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

So I remove the Gemfile and Gemfile.lock, and rebuild Jekyll, and then it works. 

Why am I getting so many errors with this Gemfile?

I do have ruby installed:

    ruby 2.0.0p481 (2014-05-08 revision 45883) [x86_64-darwin13.1.0]

And XCode. 

Besides the confusion with the errors, I&#x27;m also confused about these points:

- where do all of the bundle install utilities get installed? I don&#x27;t see any change inside Gemfile. I can&#x27;t open Gemfile.lock.

- are Gemfiles required, or only required for Github Pages solutions?

- what&#x27;s the difference between installing something myself (&#x60;gem install github-pages&#x60;) versus adding a gem in my Gemfile and then running bundle install? 

- are gems the way that people get around the &quot;no plugins allowed&quot; restrcction on github?




### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/263#issuecomment-72746777) on: **Tuesday, February 03, 2015**

&gt; Github pages requires a certain gem in order to use github as the repository, and they describe it here

No. You do *not* need a Gemfile to run Jekyll at Github Pages.
