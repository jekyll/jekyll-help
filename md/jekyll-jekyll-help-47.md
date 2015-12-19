# [&quot;Could not find blankslate-2.1.2.4 in any of the sources&quot;](https://github.com/jekyll/jekyll-help/issues/47)

> state: **closed** opened by: **** on: ****

Hi folks,

I&#x27;m playing with Jekyll 2 in a fresh dir, in preparation for upgrading a few of our 1.5 sites to it. (thanks for all your work on the 2.0 release!)

The problem I&#x27;m running into, possibly a local env issue, is:

&#x60;&#x60;&#x60;console
$ bundle exec jekyll --version
Could not find blankslate-2.1.2.4 in any of the sources
Run &#x60;bundle install&#x60; to install missing gems.
&#x60;&#x60;&#x60;

My test Gemfile&#x27;s contents are simply:

&#x60;&#x60;&#x60;ruby
source &quot;https://rubygems.org&quot;

gem &#x27;jekyll&#x27;
&#x60;&#x60;&#x60;

And this is the output I&#x27;m seeing when I install the gems:

&#x60;&#x60;&#x60;ruby
$ bundle install --path vendor/bundle
Fetching gem metadata from https://rubygems.org/.......
Fetching additional metadata from https://rubygems.org/..
Installing blankslate 2.1.2.4
Installing timers 1.1.0
Installing celluloid 0.15.2
Installing nio4r 1.0.0
Installing celluloid-io 0.15.0
Installing fast-stemmer 1.0.2
Installing classifier 1.3.4
Installing coffee-script-source 1.7.0
Installing execjs 2.0.2
Installing coffee-script 2.2.0
Installing colorator 0.1
Installing ffi 1.9.3
Installing jekyll-coffeescript 1.0.0
Installing sass 3.3.7
Installing jekyll-sass-converter 1.0.0
Installing kramdown 1.3.3
Installing liquid 2.5.5
Installing rb-fsevent 0.9.4
Installing rb-inotify 0.9.4
Installing listen 2.7.4
Installing mercenary 0.3.3
Installing posix-spawn 0.3.8
Installing yajl-ruby 1.1.0
Installing pygments.rb 0.5.4
Installing redcarpet 3.1.2
Installing safe_yaml 1.0.3
Installing parslet 1.5.0
Installing toml 0.1.1
Installing jekyll 2.0.3
Using bundler 1.6.2
Your bundle is complete!
It was installed into ./vendor/bundle
&#x60;&#x60;&#x60;

My global Ruby (via rbenv) is in &#x60;~/.rbenv/shims/ruby&#x60; — &#x60;ruby 2.1.1p76 (2014-02-24 revision 45161) [x86_64-darwin13.0]&#x60;

Any ideas why this might be happening?

Thanks in advance,
-Eric

### Comments

---
> from: [**case**](https://github.com/jekyll/jekyll-help/issues/47#issuecomment-43116279) on: **Wednesday, May 14, 2014**

Update: some combination of incantations from [this StackOverflow answer](http://stackoverflow.com/a/11146496) addressed the issue. 
