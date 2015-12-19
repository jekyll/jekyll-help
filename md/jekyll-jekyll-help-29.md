# [Once again I have another issue: running &#x27;jekyll new myblog&#x27;](https://github.com/jekyll/jekyll-help/issues/29)

> state: **closed** opened by: **** on: ****

I got jekyll installed via &#x27;sudo gem install jekyll&#x27; and now I get errors running &#x27;jekyll new myblog&#x27; from quick start guide:

&#x60;&#x60;&#x60;text
ubuntu:~$ jekyll new myblog
/var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/runner.rb:385:in &#x60;require_program&#x27;: program version required (Commander::Runner::CommandError)
	from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/runner.rb:384:in &#x60;each&#x27;
	from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/runner.rb:384:in &#x60;require_program&#x27;
	from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/runner.rb:52:in &#x60;run!&#x27;
	from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/delegates.rb:8:in &#x60;run!&#x27;
	from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/import.rb:10
	from /usr/local/bin/jekyll:19
/usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in &#x60;gem_original_require&#x27;: no such file to load -- json (LoadError)
	from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in &#x60;require&#x27;
	from /var/lib/gems/1.8/gems/jekyll-1.5.1/bin/../lib/jekyll/filters.rb:2
	from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in &#x60;gem_original_require&#x27;
	from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in &#x60;require&#x27;
	from /var/lib/gems/1.8/gems/jekyll-1.5.1/bin/../lib/jekyll.rb:44
	from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in &#x60;gem_original_require&#x27;
	from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in &#x60;require&#x27;
	from /var/lib/gems/1.8/gems/jekyll-1.5.1/bin/jekyll:7
	from /usr/local/bin/jekyll:19:in &#x60;load&#x27;
	from /usr/local/bin/jekyll:19
&#x60;&#x60;&#x60;

I really don&#x27;t know what any of this means... Can someone help please?

### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/29#issuecomment-42144563) on: **Sunday, May 04, 2014**

You&#x27;re running ruby 1.8. This isn&#x27;t advised, but should you should be able
to run jekyll just fine if you also install the json gem manually.


On 5 May 2014 06:50, Max &lt;notifications@github.com&gt; wrote:

&gt; I got jekyll installed via &#x27;sudo gem install jekyll&#x27; and now I get errors
&gt; running &#x27;jekyll new myblog&#x27; from quick start guide:
&gt;
&gt; ubuntu:~$ jekyll new myblog
&gt; /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/runner.rb:385:in require_program&#x27;:
&gt; program version required (Commander::Runner::CommandError)
&gt; from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/runner.rb:384:in
&gt; each&#x27;
&gt; from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/runner.rb:384:in
&gt; require_program&#x27;
&gt; from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/runner.rb:52:in
&gt; run!&#x27;
&gt; from
&gt; /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/delegates.rb:8:in
&gt; run!&#x27;
&gt; from /var/lib/gems/1.8/gems/commander-4.1.6/lib/commander/import.rb:10
&gt; from /usr/local/bin/jekyll:19
&gt; /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:ingem_original_require&#x27;:
&gt; no such file to load -- json (LoadError)
&gt; from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in
&gt; require&#x27;
&gt; from /var/lib/gems/1.8/gems/jekyll-1.5.1/bin/../lib/jekyll/filters.rb:2
&gt; from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in
&gt; gem_original_require&#x27;
&gt; from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in
&gt; require&#x27;
&gt; from /var/lib/gems/1.8/gems/jekyll-1.5.1/bin/../lib/jekyll.rb:44
&gt; from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in
&gt; gem_original_require&#x27;
&gt; from /usr/lib/ruby/vendor_ruby/1.8/rubygems/custom_require.rb:36:in
&gt; require&#x27;
&gt; from /var/lib/gems/1.8/gems/jekyll-1.5.1/bin/jekyll:7
&gt; from /usr/local/bin/jekyll:19:inload&#x27;
&gt; from /usr/local/bin/jekyll:19
&gt;
&gt; I really don&#x27;t know what any of this means... Can someone help please?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/29&gt;
&gt; .
&gt;
---
> from: [**UbuntuTheOS**](https://github.com/jekyll/jekyll-help/issues/29#issuecomment-42144603) on: **Sunday, May 04, 2014**

So do I install the newest version of ruby?
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/29#issuecomment-42144633) on: **Sunday, May 04, 2014**

Personally, I&#x27;d recommend Ruby 2.0+, ruby 1.8 has been end of life for a
while now. As far as Jekyll goes, no, you don&#x27;t need to, you just need to
install json manually (for now at least).


On 5 May 2014 06:52, Max &lt;notifications@github.com&gt; wrote:

&gt; So do I install the newest version of ruby?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/29#issuecomment-42144603&gt;
&gt; .
&gt;
---
> from: [**UbuntuTheOS**](https://github.com/jekyll/jekyll-help/issues/29#issuecomment-42144776) on: **Sunday, May 04, 2014**

Ok thank you!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/29#issuecomment-42191846) on: **Monday, May 05, 2014**

Going to close this one out. Looks like it has been resolved. Good job! :+1: 
