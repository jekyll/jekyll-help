# [&#x60;jekyll serve -w&#x60; is much slower than just restarting jekyll](https://github.com/jekyll/jekyll-help/issues/103)

> state: **closed** opened by: **** on: ****

If I run &#x60;jekyll serve --watch&#x60; and leave it running, jekyll takes 20 seconds between detecting the change and finishing the build. (i.e., it prints &#x60;Regenerating: 1 files&#x60; instantly, but it takes 20 seconds until &#x60;done&#x60; gets printed)

However, if I hit &lt;kbd&gt;Ctrl+C&lt;/kbd&gt; and run &#x60;jekyll serve -w&#x60; again, or run &#x60;jekyll build&#x60;, it takes just 3 seconds to rebuild everything.

How can I debug this issue? How can I make jekyll fast?
I know nothing about Ruby, so I don&#x27;t even know what to look for.

---

My system:
* Ubuntu 13.10 64-bit
* Ruby 1.9.3p194
* gem list:
&#x60;&#x60;&#x60; 
activesupport (4.1.4, 4.1.1, 3.2.18)
addressable (2.3.6)
blankslate (3.1.3, 3.1.2, 2.1.2.4)
builder (3.2.2)
bundler (1.6.4, 1.6.2)
celluloid (0.15.2)
classifier (1.3.4)
coderay (1.1.0)
coffee-script (2.3.0, 2.2.0)
coffee-script-source (1.7.1, 1.7.0)
colorator (0.1)
commander (4.2.0, 4.1.6)
cucumber (1.3.15, 1.3.11)
diff-lcs (1.2.5)
docile (1.1.5, 1.1.4)
execjs (2.2.1, 2.2.0, 2.1.0, 2.0.2)
fast-stemmer (1.0.2)
fastercsv (1.5.5)
ffi (1.9.3)
gemoji (1.5.0)
gherkin (2.12.2)
github-pages (20, 19)
highline (1.6.21)
hitimes (1.2.2, 1.2.1)
hpricot (0.8.6)
html-pipeline (1.9.0, 1.8.0, 1.5.0)
i18n (0.6.11, 0.6.9)
jekyll (2.1.1, 2.0.3, 1.5.1)
jekyll-coffeescript (1.0.0)
jekyll-gist (1.1.0, 1.0.0)
jekyll-import (0.4.0, 0.3.0, 0.2.0)
jekyll-mentions (0.1.2, 0.1.0, 0.0.9, 0.0.6)
jekyll-paginate (1.0.0)
jekyll-redirect-from (0.4.0, 0.3.1)
jekyll-sass-converter (1.0.0)
jekyll-sitemap (0.5.0, 0.4.1, 0.3.0)
jekyll-watch (1.0.0)
jekyll_test_plugin (0.1.0)
jekyll_test_plugin_malicious (0.1.0)
jemoji (0.2.0, 0.1.0)
json (1.8.1)
kramdown (1.4.0, 1.3.3, 1.3.1)
launchy (2.4.2)
liquid (2.6.1, 2.5.5)
listen (2.7.9, 2.7.8, 2.7.6, 2.7.5, 1.3.1)
maruku (0.7.2, 0.7.0)
mercenary (0.3.4, 0.3.3)
mime-types (2.3, 1.25.1)
mini_portile (0.6.0, 0.5.3)
minitest (5.4.0, 5.3.4)
multi_json (1.10.1)
multi_test (0.1.1)
nokogiri (1.6.3, 1.6.2.1, 1.6.1)
parslet (1.6.1, 1.5.0)
posix-spawn (0.3.8)
pygments.rb (0.6.0, 0.5.4)
rake (10.3.2)
rb-fsevent (0.9.4)
rb-inotify (0.9.5, 0.9.4)
rb-kqueue (0.2.3, 0.2.2)
rdiscount (2.1.7.1, 2.1.7, 1.6.8)
rdoc (4.1.1, 3.12.2)
redcarpet (3.1.2, 2.3.0)
RedCloth (4.2.9)
redgreen (1.2.2)
rouge (1.5.1, 1.4.0)
rr (1.1.2)
safe_yaml (1.0.3)
sass (3.3.10, 3.3.8, 3.3.7)
sequel (4.12.0, 4.11.0, 4.10.0)
shoulda (3.5.0)
shoulda-context (1.2.1)
shoulda-matchers (2.6.2, 2.6.1)
simplecov (0.9.0, 0.8.2)
simplecov-gem-adapter (1.0.1)
simplecov-html (0.8.0)
thread_safe (0.3.4)
timers (3.0.1, 2.0.0, 1.1.0)
toml (0.1.1)
tzinfo (1.2.1, 1.2.0)
yajl-ruby (1.2.1, 1.2.0, 1.1.0)
&#x60;&#x60;&#x60;
* _config.yml
&#x60;&#x60;&#x60;
encoding: utf-8
timezone: America/Sao_Paulo
name: Denilson Sá
markdown: redcarpet
redcarpet:
  extensions:
    - no_intra_emphasis
    - fenced_code_blocks
    - autolink
    - strikethrough
    - tables
    - smart
    - with_toc_data
highlighter: pygments
relative_permalinks: false
permalink: /blog/:year-:month-:day/:title
excerpt_separator: &quot;\n\n\n&quot;
lsi: false
gems:
  - jekyll-sitemap
  - jekyll-redirect-from
&#x60;&#x60;&#x60;

### Comments

---
> from: [**phillipadsmith**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-50364528) on: **Monday, July 28, 2014**

/subscribe (experiencing same)
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-55629304) on: **Monday, September 15, 2014**

I&#x27;m at a loss, here. Does anyone have any ideas? :confused: 
---
> from: [**alfredxing**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-55642076) on: **Monday, September 15, 2014**

Well, there&#x27;s [only one command](https://github.com/jekyll/jekyll-watch/blob/440d70b59e47dae4f86a8e8e6f60cf9777494f71/lib/jekyll/watcher.rb#L39-L42) between the &quot;Regenerating&quot; message and the &quot;done&quot; message, and that&#x27;s &#x60;Jekyll::Command.process_site&#x60;...

Maybe if anyone having this issue wants to add some more debugging messages throughout that chunk of code (including in &#x60;process_site&#x60;), it might become clearer what the issue is. The weird thing is, &#x60;jekyll build&#x60; calls &#x60;process_site&#x60; also.
---
> from: [**tommai78101**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-61356054) on: **Friday, October 31, 2014**

Is this fixed? Or there&#x27;s no update on this?
---
> from: [**sheymann**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-72070320) on: **Thursday, January 29, 2015**

:+1: compilation time goes from 5 to 20s and it&#x27;s like random.
---
> from: [**phillipadsmith**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-72582953) on: **Monday, February 02, 2015**

Hoping to try Jekyll 3.0 to see if its incremental builds fixes this...
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-75172188) on: **Thursday, February 19, 2015**

I&#x27;m going to close this since there isn&#x27;t any movement on it.

The only thing I can think of that would cause this kind of thing is either faulty hardware (bad HDD, memory, etc) or something else on the machine eating up huge cycles.
---
> from: [**phillipadsmith**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-75172611) on: **Thursday, February 19, 2015**

I&#x27;m not seeing this anymore with 2.4.0. Re-building is fairly quick. Might have been a project-specific issue.
---
> from: [**denilsonsa**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-75180786) on: **Thursday, February 19, 2015**

Here, &#x60;jekyll build&#x60; is now taking 10 seconds, while regeneration in &#x60;jekyll serve -w&#x60; takes 21 seconds.

Same system as my original post, but now with jekyll 2.5.3 (after running &#x60;gem update&#x60;).

Nothing else in the machine is eating cycles (I watched using &#x60;htop&#x60;, ruby19 process is eating the CPU). It also cannot be faulty hardware, because other people experience the same issue.

EDIT: should this issue be moved to jekyll/jekyll repository?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/103#issuecomment-75314506) on: **Friday, February 20, 2015**

&gt; should this issue be moved to jekyll/jekyll repository?

I think so... I&#x27;m still completely at a loss as to why this might be happening.
