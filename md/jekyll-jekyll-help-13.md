# [Jekyll safe_yaml libyaml warning](https://github.com/jekyll/jekyll-help/issues/13)

> state: **closed** opened by: **** on: ****

I&#x27;m getting this safe_yaml warning any time I run a Jekyll command.

&#x60;&#x60;&#x60;
$ jekyll help

  SafeYAML Warning
  ----------------

  You appear to have an outdated version of libyaml (0.1.5) installed on your system.

  Prior to 0.1.6, libyaml is vulnerable to a heap overflow exploit from malicious YAML payloads.

  For more info, see:
  https://www.ruby-lang.org/en/news/2014/03/29/heap-overflow-in-yaml-uri-escape-parsing-cve-2014-2525/

  The easiest thing to do right now is probably to update Psych to the latest version and enable
  the &#x27;bundled-libyaml&#x27; option, which will install a vendored libyaml with the vulnerability patched:

  gem install psych -- --enable-bundled-libyaml
&#x60;&#x60;&#x60;

There are a lot of folks who solved this by either upgrading their libyaml through brew, which I have done, or installing the psych gem with the lib built in. I tried both with no results either way. I&#x27;d prefer to use libyaml through brew.

I believe the problem is in rvm, but I&#x27;m not sure how to troubleshoot this. Maybe I need to make sure Ruby is using the libyaml from brew when it installs.

Here&#x27;s some info.

&#x60;&#x60;&#x60;
$ brew info libyaml
libyaml: stable 0.1.6 (bottled)
http://pyyaml.org/wiki/LibYAML
/usr/local/Cellar/libyaml/0.1.6 (7 files, 348K) *
  Poured from bottle
From: https://github.com/Homebrew/homebrew/commits/master/Library/Formula/libyaml.rb
==&gt; Options
--universal
	Build a universal binary
&#x60;&#x60;&#x60;

One weird thing is that I get the message saying I&#x27;m using libyaml 0.1.5 but I can&#x27;t find any evidence that it&#x27;s installed.

Help is appreciated.

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/13#issuecomment-40009709) on: **Wednesday, April 09, 2014**

My guess is that Ruby (which is executing Jekyll) doesn&#x27;t know about your brew installation.

You may want to check out jekyll/jekyll#2207. One of the suggested fixes is to try &#x60;rvm reinstall all --force&#x60;, but definitely give that thread a gander and see if anything pops out at you.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/13#issuecomment-40019153) on: **Wednesday, April 09, 2014**

Yes, as @troyswanson said, it may be that Ruby doesn&#x27;t know about the brew installation! You can check the version of your Rubies libyaml, using with this command &#x60;ruby -r psych -e &quot;puts Psych.libyaml_version.join(&#x27;.&#x27;)&quot;&#x60; then I will suggest the same... reinstall.
---
> from: [**ryanburnette**](https://github.com/jekyll/jekyll-help/issues/13#issuecomment-40021437) on: **Wednesday, April 09, 2014**

@albertogg 

The &#x60;ruby -r psych -e &quot;puts Psych.libyaml_version.join(&#x27;.&#x27;)&quot;&#x60; was what I was looking for it indicated that libyaml is at version 0.1.5. I&#x27;ll report back when I find the fix that works for me.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/13#issuecomment-40021471) on: **Wednesday, April 09, 2014**

:clap: Thanks!
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/13#issuecomment-40021603) on: **Wednesday, April 09, 2014**

Ok, your Ruby was compiled using 0.1.5 and it&#x27;s not looking for the brew installed version. I&#x27;ll suggest reinstall as @troyswanson said.
---
> from: [**ryanburnette**](https://github.com/jekyll/jekyll-help/issues/13#issuecomment-40022570) on: **Wednesday, April 09, 2014**

Yep. That&#x27;s weird because I&#x27;m using rvm and it usually plays nice using brew to fulfill requirements.
---
> from: [**ryanburnette**](https://github.com/jekyll/jekyll-help/issues/13#issuecomment-40022858) on: **Wednesday, April 09, 2014**

I think the issue was that my .zshrc file (I&#x27;m using oh-my-zsh) had a path line that didn&#x27;t re-insert the $PATH variable. I think that might be in oh-my-zsh&#x27;s .zshrc template. Not sure. Anyway... That was breaking rvm and keeping it from using the proper libyaml during compile. Thanks for the help guys!
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/13#issuecomment-40022947) on: **Wednesday, April 09, 2014**

No problem! :wink: :tada: 
