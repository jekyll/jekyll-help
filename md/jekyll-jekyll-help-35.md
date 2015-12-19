# [Problem with Upgrade](https://github.com/jekyll/jekyll-help/issues/35)

> state: **closed** opened by: **** on: ****

Hello, first thanks for Jekyll, and congrats on the anniversary! I was just playing around over the last few days on 1.5.1, then saw the update for 2.0.2 today and upgraded. However, if I change into a directory where I had a site built before it&#x27;s still running 1.5.1 (and presently wont serve or build). I&#x27;m new to Ruby so not really sure where this all lives on my Mac (10.9) but I must have something configured incorrectly...

Here is my terminal history... is there a way to switch/force to use the new Jekyll? I suspect this may be a conflict with bundler, homebrew or other package manager? I&#x27;ll admit I didn&#x27;t know what the hell I was doing when I installed 1.5.1 last week :/ 

&#x60;&#x60;&#x60;
mbp:~ zachlyon$ jekyll -v
jekyll 2.0.2
mbp:~ zachlyon$ cd Sites/mysite.jekyll/
mbp:mysite.jekyll zachlyon$ jekyll -v
jekyll 1.5.1
mbp:mysite.jekyll zachlyon$ gem update jekyll
Updating installed gems
Nothing to update
mbp:mysite.jekyll zachlyon$ gem uninstall jekyll

Select gem to uninstall:
 1. jekyll-1.5.1
 2. jekyll-2.0.2
 3. All versions
&gt; 1

You have requested to uninstall the gem:
	jekyll-1.5.1

jekyll-coffeescript-1.0.0 depends on jekyll (~&gt; 1.0, development)
jekyll-sitemap-0.3.0 depends on jekyll (~&gt; 1.4)
If you remove this gem, these dependencies will not be met.
Continue with Uninstall? [yN]  y
Successfully uninstalled jekyll-1.5.1
mbp:mysite.jekyll zachlyon$ jekyll -v
Could not find proper version of jekyll (1.5.1) in any of the sources
Run &#x60;bundle install&#x60; to install missing gems.
mbp:mysite.jekyll zachlyon$ 
&#x60;&#x60;&#x60;

The bundle install just grabs another 1.5.1. Any ideas?

Thanks!
Zach


### Comments

---
> from: [**ndarville**](https://github.com/jekyll/jekyll-help/issues/35#issuecomment-42639048) on: **Friday, May 09, 2014**

You closed your issue. Was this intentional?
---
> from: [**zachlyon**](https://github.com/jekyll/jekyll-help/issues/35#issuecomment-42685207) on: **Friday, May 09, 2014**

Yep, I figure it was my bad. I basically just fired up a new jekyll site in a different directory and copied my files over to it and started working from there. I couldn&#x27;t figure out how one site/directory was &quot;stuck&quot; on 1.5.1 and every new site or any other dir on my system was 2.0.2 according to jekyll -v.



---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/35#issuecomment-42691073) on: **Friday, May 09, 2014**

@zachlyon for future reference install **bundler** (&#x60;gem install bundler&#x60;), create a &#x60;Gemfile&#x60; file and add your project dependencies there and then type &#x60;bundle install&#x60;. Run your project prefixing it with &#x60;bundle exec&#x60;, e.g. &#x60;bundle exec jekyll serve -w&#x60; that will help you avoid this type of problems. I know it&#x27;s a bunch of stuff but you can check [my jekyll site](https://github.com/albertogg/albertogg.github.com) if you want.
