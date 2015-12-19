# [Problems installing Jekyll on Ubuntu](https://github.com/jekyll/jekyll-help/issues/8)

> state: **closed** opened by: **** on: ****

When I run
&#x60;&#x60;&#x60;
&gt;&gt; sudo gem install jekyll
&#x60;&#x60;&#x60;

I get the error

&#x60;&#x60;&#x60;
ERROR:  While generating documentation for jekyll-1.5.1
... MESSAGE:   Unhandled special: Special: type=17, text=&quot;&lt;!-- more --&gt;&quot;
... RDOC args: --ri --op /var/lib/gems/1.8/doc/jekyll-1.5.1/ri --charset=UTF-8 lib README.markdown LICENSE --title jekyll-1.5.1 Documentation --quiet
&#x60;&#x60;&#x60;

Any ideas?

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39393037) on: **Wednesday, April 02, 2014**

Hello, In what platform (windows, mac, Linux) are you installing Jekyll? Also I&#x27;m  not sure if Jekyll supports ruby 1.8.x any more, and it seems you are using that version. Can you confirm that?—
Alberto Grespan

On Wed, Apr 2, 2014 at 6:02 PM, rowoflo &lt;notifications@github.com&gt; wrote:

&gt; When I run
&gt; &#x60;&#x60;&#x60;
&gt;&gt;&gt; sudo gem install jekyll
&gt; &#x60;&#x60;&#x60;
&gt; I get the error
&gt; &#x60;&#x60;&#x60;
&gt; ERROR:  While generating documentation for jekyll-1.5.1
&gt; ... MESSAGE:   Unhandled special: Special: type=17, text=&quot;&lt;!-- more --&gt;&quot;
&gt; ... RDOC args: --ri --op /var/lib/gems/1.8/doc/jekyll-1.5.1/ri --charset=UTF-8 lib README.markdown LICENSE --title jekyll-1.5.1 Documentation --quiet
&gt; &#x60;&#x60;&#x60;
&gt; Any ideas?
&gt; ---
&gt; Reply to this email directly or view it on GitHub:
&gt; https://github.com/jekyll/help/issues/8
---
> from: [**rowoflo**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39393443) on: **Wednesday, April 02, 2014**

I am using Ubuntu 12.04 in a virtual machine. I am using ruby 1.87 and I installed gem with
&#x60;&#x60;&#x60;
sudo apt-get install rubygems
&#x60;&#x60;&#x60;

Sorry, I accidentally hit the close button before. 

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39395389) on: **Wednesday, April 02, 2014**

Ok, I just tried that on a VM as well, and I&#x27;m getting the same error while generating the &#x60;ri-documentation&#x60; for jekyll. If you don&#x27;t care about the local gem documentation and just want to install the gem try this, &#x60;sudo gem install jekyll --no-ri --no-rdoc&#x60;. We will check for this error.

Hope this helps.
---
> from: [**rowoflo**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39395845) on: **Wednesday, April 02, 2014**

That allowed me to install jekyll with no errors, but when I try &#x60;&#x60;&#x60;jekyll new my-site&#x60;&#x60;&#x60; I get this error
&#x60;&#x60;&#x60;
$ jekyll new my-site
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
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39396234) on: **Wednesday, April 02, 2014**

Yes, I&#x27;m getting this error too. Are you forced to use Ruby 1.8.7?
---
> from: [**rowoflo**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39396403) on: **Wednesday, April 02, 2014**

Not at all. Like I said before I just did &#x60;&#x60;&#x60;sudo apt-get install rubygems&#x60;&#x60;&#x60;.

What is the best way to install Ruby on Ubuntu? I usually work on a Mac, so not quite up to speed on the best way to do stuff on the Linux side things.
---
> from: [**fabianrbz**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39397515) on: **Wednesday, April 02, 2014**

the support for ruby 1.8.7 was removed in [this](https://github.com/jekyll/jekyll/pull/1780/) pull request, and it&#x27;s not part of v1.5.1.

@rowoflo in that pull request there are some comments explaining how to install different versions of ruby in ubuntu
---
> from: [**rowoflo**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39399559) on: **Wednesday, April 02, 2014**

Ok, got it working! Just for the record...

I followed the steps as stated [here](http://leonard.io/blog/2011/12/installing-ruby-1-9-2-on-ubuntu-11-10-oneric-ocelot-without-using-rvm/) to remove Ruby 1.8. However, my gem version was not updating even though my ruby version was updating.
&#x60;&#x60;&#x60;
$ ruby --version
ruby 1.9.3p0 (2011-10-30 revision 33570) [x86_64-linux]
$ gem --version
1.8.11
&#x60;&#x60;&#x60;
But gem 1.9.1 was installed in /usr/bin. So, I just installed Jekyll using gem 1.9.1 using
&#x60;&#x60;&#x60;
sudo /usr/bin/gem1.9.1 install jekyll
&#x60;&#x60;&#x60;
After that I could run Jekyll commands as usual
&#x60;&#x60;&#x60;
jekyll new website
cd website
jekyll build
jekyll serve
&#x60;&#x60;&#x60;
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39400990) on: **Wednesday, April 02, 2014**

The &#x60;gem --version&#x60; will not necessarily be the same version as the package version (1.9.1). &#x60;sudo update-alternatives --config gem&#x60; should have done the symlink, check if &#x60;/usr/bin/gem&#x60; is symlinked or hardlinked to &#x60;/usr/bin/gem1.9.1&#x60; or &#x60;/etc/alternatives/gem&#x60;. You can use this command &#x60;ls -la /usr/bin | grep gem&#x60;. If that&#x27;s the case you are fine and can use the gem command normally.
---
> from: [**rowoflo**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39401417) on: **Wednesday, April 02, 2014**

Yep, you are correct. I checked that gem is actually symlinked to &#x60;&#x60;&#x60;/usr/bin/gem1.9.1&#x60;&#x60;&#x60;. Just to make sure, I uninstalled Jekyll and reinstall with just the gem command everything still worked fine.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39401665) on: **Wednesday, April 02, 2014**

I&#x27;m glad! Let&#x27;s close this issue!
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-39402810) on: **Wednesday, April 02, 2014**

:+1: Good job all!
---
> from: [**yourmoonlight**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-40186691) on: **Friday, April 11, 2014**

many thanks!!
---
> from: [**rowoflo**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-40200526) on: **Friday, April 11, 2014**

Thank you for all the help!
---
> from: [**doriel**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-46497632) on: **Wednesday, June 18, 2014**

@rowoflo and @albertogg  Thanks this info really helped me a lot :+1:  
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/8#issuecomment-46504367) on: **Wednesday, June 18, 2014**

@doriel glad it did!
