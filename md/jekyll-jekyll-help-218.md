# [Can&#x27;t seem to get bundler to update Jekyll 1.5.1 to latest](https://github.com/jekyll/jekyll-help/issues/218)

> state: **closed** opened by: **** on: ****

I&#x27;m trying to update my [personal site](https://github.com/jglovier/jglovier.github.io) to use Jekyll 2.X.X locally, and cannot seem to get it to update for the life of me. :rage4: I&#x27;ve tried everything I can think of, and it probably stems from my lack of knowledge around bundler, so I&#x27;m reaching out ~~in despair~~ for help.

When I try simply running &#x60;bundle update&#x60;, bundler does it&#x27;s thing, but still ends up using Jekyll 1.5.1 at the end (maybe that command doesn&#x27;t do what I think it should be doing):

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle update
Fetching gem metadata from https://rubygems.org/........
Resolving dependencies..........
Using json 1.8.1
Using minitest 5.5.0
Using thread_safe 0.3.4
Using blankslate 2.1.2.4
Using i18n 0.7.0
Using fast-stemmer 1.0.2
Using colorator 0.1
Using highline 1.6.21
Using ffi 1.9.6
Using RedCloth 4.2.9
Using gemoji 1.5.0
Using liquid 2.5.5
Using rb-fsevent 0.9.4
Using maruku 0.7.0
Using posix-spawn 0.3.9
Using yajl-ruby 1.1.0
Using safe_yaml 1.0.4
Using mini_portile 0.6.1
Using redcarpet 2.3.0
Using kramdown 1.3.1
Using rdiscount 2.1.7
Using bundler 1.6.9
Using tzinfo 1.2.2
Using parslet 1.5.0
Using classifier 1.3.4
Using commander 4.1.6
Using rb-inotify 0.9.5
Using rb-kqueue 0.2.3
Using pygments.rb 0.5.4
Using activesupport 4.2.0
Using nokogiri 1.6.5
Using toml 0.1.2
Using listen 1.3.1
Using html-pipeline 1.5.0
Using jekyll 1.5.1
Using jekyll-mentions 0.0.9
Using jekyll-sitemap 0.3.0
Using jekyll-redirect-from 0.3.1
Using jemoji 0.1.0
Using github-pages 20
Your bundle is updated!
&#x60;&#x60;&#x60;

Also tried deleting &#x60;Gemfile.lock&#x60; and just running &#x60;bundle install&#x60;, but that appears to do the same thing, and the result is just a new &#x60;Gemfile.lock&#x60; with the same &#x60;Jekyll 1.5.1&#x60; listed. (Is there something wrong with the way my [Gemfile is setup](https://gist.github.com/jglovier/1477caab8209d8cbd27c)?)

So interestingly, when I run &#x60;jekyll -v&#x60; I get the following output, which I don&#x27;t understand at all:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ jekyll -v
WARN: Unresolved specs during Gem::Specification.reset:
      posix-spawn (~&gt; 0.3.6)
      redcarpet (~&gt; 3.1)
      listen (~&gt; 2.7)
      classifier-reborn (~&gt; 2.0)
WARN: Clearing out unresolved specs.
Please report a bug if this causes problems.
/opt/rubies/2.1.4-github1/lib/ruby/gems/2.1.0/gems/jekyll-2.5.3/bin/jekyll:21:in &#x60;block in &lt;top (required)&gt;&#x27;: cannot load such file -- jekyll/version (LoadError)
	from /opt/rubies/2.1.4-github1/lib/ruby/gems/2.1.0/gems/mercenary-0.3.5/lib/mercenary.rb:18:in &#x60;program&#x27;
	from /opt/rubies/2.1.4-github1/lib/ruby/gems/2.1.0/gems/jekyll-2.5.3/bin/jekyll:20:in &#x60;&lt;top (required)&gt;&#x27;
	from /opt/boxen/rbenv/versions/2.1.4-github/bin/jekyll:23:in &#x60;load&#x27;
	from /opt/boxen/rbenv/versions/2.1.4-github/bin/jekyll:23:in &#x60;&lt;main&gt;&#x27;
&#x60;&#x60;&#x60;

Of course, if I run &#x60;bundle exec jekyll -v&#x60; I get the expected:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle exec jekyll -v
jekyll 1.5.1
&#x60;&#x60;&#x60;

So I tried straight up doing &#x60;gem update jekyll&#x60;, and I get:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ gem update jekyll
Updating installed gems
Nothing to update
&#x60;&#x60;&#x60;

When I try &#x60;bundle exec gem update jekyll&#x60;, and I actually get:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle exec gem update jekyll
Updating installed gems
Updating jekyll
Successfully installed liquid-2.6.1
Successfully installed mercenary-0.3.5
Successfully installed pygments.rb-0.6.0
Building native extensions.  This could take a while...
Successfully installed redcarpet-3.2.2
Successfully installed jekyll-paginate-1.1.0
Successfully installed jekyll-gist-1.1.0
Successfully installed coffee-script-source-1.8.0
Successfully installed execjs-2.2.2
Successfully installed coffee-script-2.3.0
Successfully installed jekyll-coffeescript-1.0.1
Successfully installed sass-3.4.9
Successfully installed jekyll-sass-converter-1.3.0
Building native extensions.  This could take a while...
Successfully installed hitimes-1.2.2
Successfully installed timers-4.0.1
Successfully installed celluloid-0.16.0
Successfully installed listen-2.8.4
Successfully installed jekyll-watch-1.2.0
Successfully installed classifier-reborn-2.0.3
Successfully installed jekyll-2.5.3
Updating jekyll-mentions
Successfully installed classifier-reborn-2.0.3
Building native extensions.  This could take a while...
Successfully installed hitimes-1.2.2
Successfully installed timers-4.0.1
Successfully installed celluloid-0.16.0
Successfully installed listen-2.8.4
Successfully installed jekyll-watch-1.2.0
Successfully installed sass-3.4.9
Successfully installed jekyll-sass-converter-1.3.0
Successfully installed coffee-script-source-1.8.0
Successfully installed execjs-2.2.2
Successfully installed coffee-script-2.3.0
Successfully installed jekyll-coffeescript-1.0.1
Successfully installed jekyll-gist-1.1.0
Successfully installed jekyll-paginate-1.1.0
Building native extensions.  This could take a while...
Successfully installed redcarpet-3.2.2
Successfully installed pygments.rb-0.6.0
Successfully installed mercenary-0.3.5
Successfully installed liquid-2.6.1
Successfully installed jekyll-2.5.3
-------------------------------------------------
Thank you for installing html-pipeline!
You must bundle Filter gem dependencies.
See html-pipeline README.md for more details.
https://github.com/jch/html-pipeline#dependencies
-------------------------------------------------
Successfully installed html-pipeline-1.9.0
Successfully installed jekyll-mentions-0.2.1
Updating jekyll-redirect-from
Successfully installed classifier-reborn-2.0.3
Building native extensions.  This could take a while...
Successfully installed hitimes-1.2.2
Successfully installed timers-4.0.1
Successfully installed celluloid-0.16.0
Successfully installed listen-2.8.4
Successfully installed jekyll-watch-1.2.0
Successfully installed sass-3.4.9
Successfully installed jekyll-sass-converter-1.3.0
Successfully installed coffee-script-source-1.8.0
Successfully installed execjs-2.2.2
Successfully installed coffee-script-2.3.0
Successfully installed jekyll-coffeescript-1.0.1
Successfully installed jekyll-gist-1.1.0
Successfully installed jekyll-paginate-1.1.0
Building native extensions.  This could take a while...
Successfully installed redcarpet-3.2.2
Successfully installed pygments.rb-0.6.0
Successfully installed mercenary-0.3.5
Successfully installed liquid-2.6.1
Successfully installed jekyll-2.5.3
Successfully installed jekyll-redirect-from-0.6.2
Updating jekyll-sitemap
Successfully installed jekyll-sitemap-0.7.0
Gems updated: celluloid classifier-reborn coffee-script coffee-script-source execjs hitimes jekyll jekyll-coffeescript jekyll-gist jekyll-paginate jekyll-sass-converter jekyll-watch liquid listen mercenary pygments.rb redcarpet sass timers celluloid classifier-reborn coffee-script coffee-script-source execjs hitimes html-pipeline jekyll jekyll-coffeescript jekyll-gist jekyll-mentions jekyll-paginate jekyll-sass-converter jekyll-watch liquid listen mercenary pygments.rb redcarpet sass timers celluloid classifier-reborn coffee-script coffee-script-source execjs hitimes jekyll jekyll-coffeescript jekyll-gist jekyll-paginate jekyll-redirect-from jekyll-sass-converter jekyll-watch liquid listen mercenary pygments.rb redcarpet sass timers jekyll-sitemap
&#x60;&#x60;&#x60;

...however when I run &#x60;bundle exec jekyll -v&#x60; it just shows:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle exec jekyll -v
jekyll 1.5.1
&#x60;&#x60;&#x60;

I even tried blowing away my entire local repo via &#x60;rm -rf&#x60;, and re-cloning from GitHub, and trying all this stuff again (in case it was just some odd bug), but no avail, and same results.

What am I doing wrong??? 

![img](http://gifs.joelglovier.com/accidents/fumble.gif)

### Comments

---
> from: [**mattr-**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68060057) on: **Wednesday, December 24, 2014**

The &#x60;github-pages&#x60; gem is what&#x27;s constraining your Jekyll version here. A &#x60;bundle update github-pages&#x60; should get you the upgraded version of Jekyll you&#x27;re looking for.
---
> from: [**mattr-**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68061468) on: **Wednesday, December 24, 2014**

You&#x27;ll also want to make sure that when you do anything with Jekyll locally, that you always run things through &#x60;bundle exec&#x60; as that&#x27;s the best way to ensure consistency.
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68066050) on: **Wednesday, December 24, 2014**

&gt;A &#x60;bundle update github-pages&#x60; should get you the upgraded version of Jekyll you&#x27;re looking for.

@mattr- thanks for the help! However, running &#x60;bundle update github-pages&#x60; doesn&#x27;t seem to affect anything. Still uses &#x60;1.5.1&#x60;:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle update github-pages
Fetching gem metadata from https://rubygems.org/........
Resolving dependencies...........
Using json 1.8.1
Using RedCloth 4.2.9
Using i18n 0.7.0
Using minitest 5.5.0
Using thread_safe 0.3.4
Using blankslate 2.1.2.4
Using fast-stemmer 1.0.2
Using colorator 0.1
Using highline 1.6.21
Using ffi 1.9.6
Using gemoji 1.5.0
Using liquid 2.5.5
Using rb-fsevent 0.9.4
Using maruku 0.7.0
Using posix-spawn 0.3.9
Using yajl-ruby 1.1.0
Using redcarpet 2.3.0
Using safe_yaml 1.0.4
Using mini_portile 0.6.1
Using kramdown 1.3.1
Using rdiscount 2.1.7
Using bundler 1.6.9
Using parslet 1.5.0
Using tzinfo 1.2.2
Using classifier 1.3.4
Using commander 4.1.6
Using rb-inotify 0.9.5
Using pygments.rb 0.5.4
Using rb-kqueue 0.2.3
Using toml 0.1.2
Using activesupport 4.2.0
Using nokogiri 1.6.5
Using listen 1.3.1
Using html-pipeline 1.5.0
Using jekyll 1.5.1
Using jekyll-redirect-from 0.3.1
Using jekyll-sitemap 0.3.0
Using jekyll-mentions 0.0.9
Using jemoji 0.1.0
Using github-pages 20
Your bundle is updated!

|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle exec jekyll -v
jekyll 1.5.1
&#x60;&#x60;&#x60;
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68067501) on: **Wednesday, December 24, 2014**

cc https://github.com/jekyll/jekyll/issues/3124 for another example case for that thread.
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68068186) on: **Wednesday, December 24, 2014**

&gt;A bundle update github-pages should get you the upgraded version of Jekyll you&#x27;re looking for.

@mattr- also, shouldn&#x27;t this have already been performed as part of running &#x60;bundle update&#x60;??
---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68070371) on: **Wednesday, December 24, 2014**

Perhaps what&#x27;s confusing it is that you have both &#x60;github-pages&#x60; and &#x60;jekyll&#x60; in the Gemfile? Maybe try removing &#x60;jekyll&#x60; and &#x60;bundle update&#x60; again.
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68075707) on: **Wednesday, December 24, 2014**

@alfredxing so I tried removing &#x60;jekyll&#x60; from the Gemfile, and ran &#x60;bundle update&#x60;, but it&#x27;s still using &#x60;jekyll 1.5.1&#x60; :frowning: 

I also tried running &#x60;gem cleanup&#x60;, which returned:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ gem cleanup
Cleaning up installed gems...
Attempting to uninstall github-pages-20
Successfully uninstalled github-pages-20
Attempting to uninstall jemoji-0.1.0
Successfully uninstalled jemoji-0.1.0
Attempting to uninstall jekyll-sitemap-0.3.0
Successfully uninstalled jekyll-sitemap-0.3.0
Attempting to uninstall jekyll-redirect-from-0.3.1
Successfully uninstalled jekyll-redirect-from-0.3.1
Attempting to uninstall jekyll-mentions-0.0.9
Successfully uninstalled jekyll-mentions-0.0.9
Attempting to uninstall jekyll-1.5.1
Successfully uninstalled jekyll-1.5.1
Attempting to uninstall html-pipeline-1.5.0
Successfully uninstalled html-pipeline-1.5.0
Attempting to uninstall gemoji-1.5.0
Successfully uninstalled gemoji-1.5.0
Attempting to uninstall redcarpet-2.3.0
Successfully uninstalled redcarpet-2.3.0
Attempting to uninstall pygments.rb-0.5.4
Successfully uninstalled pygments.rb-0.5.4
Attempting to uninstall listen-1.3.1
Successfully uninstalled listen-1.3.1
Attempting to uninstall liquid-2.5.5
Successfully uninstalled liquid-2.5.5
Clean Up Complete
&#x60;&#x60;&#x60;

But I ran &#x60;bundle update&#x60; again afterward and it&#x27;s still showing &#x60;jekyll 1.5.1&#x60;.
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68075738) on: **Wednesday, December 24, 2014**

How do you even computer??

![image](https://cloud.githubusercontent.com/assets/1319791/5551177/21e3fc5e-8b8d-11e4-8c5d-e1a071d69a6a.png)

---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68075778) on: **Wednesday, December 24, 2014**

Should I be deleting the &#x60;Gemfile.lock&#x60; manually? I&#x27;m assuming that supposed to get&#x27;s overwritten when I run &#x60;bundle update&#x60; and there is a new gem available. But is my assumption incorrect?
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68075959) on: **Wednesday, December 24, 2014**

What&#x27;s weird here is when I run &#x60;bundle install&#x60; or &#x60;bundle update&#x60; again, I see version 20 of the GH pages gem listed:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle update github-pages
Fetching gem metadata from https://rubygems.org/.........
Resolving dependencies...........
Using json 1.8.1
Using i18n 0.7.0
Using RedCloth 4.2.9
Using minitest 5.5.0
Using blankslate 2.1.2.4
Using thread_safe 0.3.4
Using fast-stemmer 1.0.2
Using colorator 0.1
Using ffi 1.9.6
Using highline 1.6.21
Using gemoji 1.5.0
Using liquid 2.5.5
Using rb-fsevent 0.9.4
Using maruku 0.7.0
Using posix-spawn 0.3.9
Using yajl-ruby 1.1.0
Using redcarpet 2.3.0
Using safe_yaml 1.0.4
Using mini_portile 0.6.1
Using kramdown 1.3.1
Using bundler 1.6.9
Using rdiscount 2.1.7
Using tzinfo 1.2.2
Using classifier 1.3.4
Using parslet 1.5.0
Using commander 4.1.6
Using rb-inotify 0.9.5
Using rb-kqueue 0.2.3
Using pygments.rb 0.5.4
Using nokogiri 1.6.5
Using toml 0.1.2
Using activesupport 4.2.0
Using listen 1.3.1
Using html-pipeline 1.5.0
Using jekyll 1.5.1
Using jekyll-redirect-from 0.3.1
Using jekyll-sitemap 0.3.0
Using jekyll-mentions 0.0.9
Using jemoji 0.1.0
Using github-pages 20
Your bundle is updated!
&#x60;&#x60;&#x60;

However when I actually ask which version is installed, it tell sme version 31:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ github-pages -v
github-pages 31
&#x60;&#x60;&#x60;
---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68080550) on: **Wednesday, December 24, 2014**

Maybe try removing &#x60;Gemfile.lock&#x60;, then &#x60;gem uninstall github-pages jekyll&#x60;, &#x60;gem install github-pages&#x60;, &#x60;bundle install&#x60;?
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68084377) on: **Wednesday, December 24, 2014**

Ugh - looked like it worked, but does not appear to have:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ gem uninstall github-pages jekyll

Select gem to uninstall:
 1. jekyll-1.5.1
 2. jekyll-2.4.0
 3. jekyll-2.5.3
 4. All versions
&gt; 1.

You have requested to uninstall the gem:
	jekyll-1.5.1

github-pages-20 depends on jekyll (= 1.5.1)
jekyll-mentions-0.0.9 depends on jekyll (~&gt; 1.4)
jekyll-redirect-from-0.3.1 depends on jekyll (~&gt; 1.4)
jekyll-sitemap-0.3.0 depends on jekyll (~&gt; 1.4)
jemoji-0.1.0 depends on jekyll (~&gt; 1.4)
If you remove this gem, these dependencies will not be met.
Continue with Uninstall? [yN]  y
Successfully uninstalled jekyll-1.5.1

Select gem to uninstall:
 1. github-pages-20
 2. github-pages-31
 3. All versions
&gt; 3.
Successfully uninstalled github-pages-20
Remove executables:
	github-pages

in addition to the gem? [Yn]  Y
Removing github-pages
Successfully uninstalled github-pages-31


|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ gem uninstall jekyll

Select gem to uninstall:
 1. jekyll-2.4.0
 2. jekyll-2.5.3
 3. All versions
&gt; 3.
Successfully uninstalled jekyll-2.4.0

You have requested to uninstall the gem:
	jekyll-2.5.3

jekyll-coffeescript-1.0.1 depends on jekyll (~&gt; 2.0, development)
jekyll-gist-1.1.0 depends on jekyll (~&gt; 2.0, development)
jekyll-mentions-0.2.1 depends on jekyll (~&gt; 2.0)
jekyll-paginate-1.1.0 depends on jekyll (~&gt; 2.0, development)
jekyll-redirect-from-0.6.2 depends on jekyll (~&gt; 2.0)
jekyll-sass-converter-1.3.0 depends on jekyll (~&gt; 2.0, development)
jekyll-sass-converter-1.2.0 depends on jekyll (~&gt; 2.0, development)
jekyll-sitemap-0.7.0 depends on jekyll (~&gt; 2.0, development)
jekyll-sitemap-0.6.3 depends on jekyll (~&gt; 2.0, development)
jekyll-watch-1.2.0 depends on jekyll (~&gt; 2.0, development)
jekyll_test_plugin-0.1.0 depends on jekyll (&gt;= 0, development)
jekyll_test_plugin_malicious-0.1.0 depends on jekyll (&gt;= 0, development)
jemoji-0.4.0 depends on jekyll (~&gt; 2.0)
If you remove this gem, these dependencies will not be met.
Continue with Uninstall? [yN]  y
Remove executables:
	jekyll

in addition to the gem? [Yn]  Y
Removing jekyll
Successfully uninstalled jekyll-2.5.3


|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ gem install github-pages
Fetching: jekyll-2.4.0.gem (100%)
Successfully installed jekyll-2.4.0
Fetching: github-pages-31.gem (100%)
Successfully installed github-pages-31
2 gems installed


|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle install
Fetching gem metadata from https://rubygems.org/.........
Resolving dependencies...........
Using i18n 0.7.0
Using json 1.8.1
Using RedCloth 4.2.9
Using minitest 5.5.0
Using thread_safe 0.3.4
Using blankslate 2.1.2.4
Using fast-stemmer 1.0.2
Using colorator 0.1
Using highline 1.6.21
Using ffi 1.9.6
Using gemoji 1.5.0
Using liquid 2.5.5
Using rb-fsevent 0.9.4
Using maruku 0.7.0
Using posix-spawn 0.3.9
Using yajl-ruby 1.1.0
Using redcarpet 2.3.0
Using safe_yaml 1.0.4
Using mini_portile 0.6.1
Using kramdown 1.3.1
Using rdiscount 2.1.7
Using bundler 1.6.9
Using parslet 1.5.0
Using tzinfo 1.2.2
Using commander 4.1.6
Using classifier 1.3.4
Using rb-inotify 0.9.5
Using rb-kqueue 0.2.3
Using pygments.rb 0.5.4
Using nokogiri 1.6.5
Using toml 0.1.2
Using activesupport 4.2.0
Using listen 1.3.1
Using html-pipeline 1.5.0
Installing jekyll 1.5.1
Using jekyll-sitemap 0.3.0
Using jekyll-mentions 0.0.9
Using jemoji 0.1.0
Using jekyll-redirect-from 0.3.1
Installing github-pages 20
Your bundle is complete!
Use &#x60;bundle show [gemname]&#x60; to see where a bundled gem is installed.


|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ jekyll -v
jekyll 2.4.0


|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle exec jekyll -v
jekyll 1.5.1
&#x60;&#x60;&#x60;
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68084391) on: **Wednesday, December 24, 2014**

@alfredxing btw thank you for your help!!
---
> from: [**mattr-**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68086592) on: **Wednesday, December 24, 2014**

@jglovier I&#x27;ve [updated your Gemfile][here]. Setting the explicit version for the &#x60;github-pages&#x60; gem and then running &#x60;bundle update&#x60; gets you where you want to go. Once the update is complete, you can remove that version specifier from your Gemfile.

I&#x27;m not sure why running &#x60;bundle update&#x60; doesn&#x27;t take care of this for you, because this is exactly what it should do according to the docs. It would appear we&#x27;ve found a bug in Bundler.

[here]: https://gist.github.com/mattr-/410afef45f3baac50189
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68088452) on: **Wednesday, December 24, 2014**

&gt;Setting the explicit version for the github-pages gem and then running bundle update gets you where you want to go.

Whoa - interesting.

&gt;Once the update is complete, you can remove that version specifier from your Gemfile.

So, interestingly, once I removed the version and ran &#x60;bundle update&#x60; again, it actually reverted back to version 20:

&#x60;&#x60;&#x60;
|===jglovier===| ~/code/github-jglovier/jglovier.github.io $ bundle update
Fetching gem metadata from https://rubygems.org/.........
Resolving dependencies...........
Using RedCloth 4.2.9
Using minitest 5.5.0
Using thread_safe 0.3.4
Using blankslate 2.1.2.4
Using fast-stemmer 1.0.2
Using colorator 0.1
Using highline 1.6.21
Using i18n 0.7.0
Using ffi 1.9.6
Using liquid 2.5.5 (was 2.6.1)
Using gemoji 1.5.0 (was 2.1.0)
Using rb-fsevent 0.9.4
Using maruku 0.7.0
Using posix-spawn 0.3.9
Using yajl-ruby 1.1.0
Using redcarpet 2.3.0 (was 3.1.2)
Using safe_yaml 1.0.4
Using json 1.8.1
Using mini_portile 0.6.1
Using kramdown 1.3.1
Using rdiscount 2.1.7
Using bundler 1.6.9
Using tzinfo 1.2.2
Using parslet 1.5.0
Using commander 4.1.6
Using classifier 1.3.4
Using rb-inotify 0.9.5
Using rb-kqueue 0.2.3
Using pygments.rb 0.5.4 (was 0.6.0)
Using nokogiri 1.6.5
Using activesupport 4.2.0
Using toml 0.1.2
Using listen 1.3.1 (was 2.8.4)
Using html-pipeline 1.5.0 (was 1.9.0)
Using jekyll 1.5.1 (was 2.4.0)
Using jekyll-mentions 0.0.9 (was 0.2.1)
Using jekyll-redirect-from 0.3.1 (was 0.6.2)
Using jekyll-sitemap 0.3.0 (was 0.6.3)
Using jemoji 0.1.0 (was 0.4.0)
Using github-pages 20 (was 31)
Your bundle is updated!
&#x60;&#x60;&#x60;

@mattr- thanks for your help here!!!

And Merry Christmas to all. :christmas_tree: :santa: 
---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68091572) on: **Wednesday, December 24, 2014**

I just set up the same scenario and got the same result. Interestingly though, if I removed the &#x60;gem &#x27;jekyll-redirect-from&#x27;&#x60; line in the Gemfile and did &#x60;bundle update&#x60;, it allowed for Pages version 31 and Jekyll 2.4...

Can you try that and see if you get the same result?
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68153463) on: **Friday, December 26, 2014**

&gt;Can you try that and see if you get the same result?

Yup, good catch. When I remove &#x60;jekyll-redirect-from&#x60; bundler updates to the latest versions! Should I open an issue on &#x60;jekyll-redirect-from&#x60; for that then? cc @parkr 
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68556394) on: **Friday, January 02, 2015**

&gt; Yup, good catch. When I remove jekyll-redirect-from bundler updates to the latest versions! Should I open an issue on jekyll-redirect-from for that then?

I don&#x27;t think so – it&#x27;s the way RubyGems resolves dependencies. If you specify &#x60;gem &#x27;jekyll-redirect-from&#x27;&#x60; and in your &#x60;Gemfile.lock&#x60; it&#x27;s locked to &#x60;0.3.1&#x60;, then any time you run &#x60;bundle install&#x60;, it&#x27;ll make sure all the other dependencies allow for this, by taking, for example, a version of &#x60;github-pages&#x60; which contained &#x60;jekyll-redirect-from v0.3.1&#x60;. If you delete your Gemfile.lock, or manually update your &#x60;gem &#x27;jekyll-redirect-from&#x27;&#x60; to &#x60;gem &#x27;jekyll-redirect-from&#x27;, &#x27;~&gt; 0.6&#x27;&#x60;, and run &#x60;bundle install&#x60;, then you should be good to go. :)
---
> from: [**jglovier**](https://github.com/jekyll/jekyll-help/issues/218#issuecomment-68956264) on: **Tuesday, January 06, 2015**

Ahh :cool:. @parkr as always, thanks for the education!
