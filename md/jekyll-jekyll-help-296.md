# [problem with installing jekyll on windows](https://github.com/jekyll/jekyll-help/issues/296)

> state: **closed** opened by: **** on: ****

C:\Users\amysidra&gt;gem install jekyll
ERROR:  While executing gem ... (Gem::RemoteFetcher::UnknownHostError)
    no such name (https://api.rubygems.org/quick/Marshal.4.8/jekyll-2.5.3.gemspe
c.rz)

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/296#issuecomment-100771219) on: **Sunday, May 10, 2015**

Sometimes there&#x27;s a problem getting to RubyGems from Southeast Asia: https://github.com/jekyll/jekyll/issues/1409

&#x60;&#x60;&#x60;text
$ gem sources --remove http://rubygems.org/
$ gem sources -a http://ruby.taobao.org/
$ gem sources -l
*** CURRENT SOURCES ***
http://ruby.taobao.org    
$ gem install rack 
&#x60;&#x60;&#x60;
