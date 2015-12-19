# [Problems with installing](https://github.com/jekyll/jekyll-help/issues/43)

> state: **closed** opened by: **** on: ****

Tried to install jekyll today on Ubuntu 12.04 LTS.
Installed ruby1.9.1 and rubygems using     
&#x60;sudo apt-get install ruby1.9.1 rubygems&#x60;    
After problems installing jekyll via termina also installed ruby1.9.1-dev using   
&#x60;sudo apt-get install ruby1.9.1-dev&#x60;    
But I still couldn&#x27;t install jekyll using    
&#x60;gem install jekyll&#x60;    
Output was    
&#x60;ERROR:  While executing gem ... (Errno::EACCES)
    Permission denied - /var/lib/gems&#x60;    
Used    
&#x60;sudo gem install jekyll&#x60; instead.
The result was following    
&#x60;Fetching: liquid-2.5.5.gem (100%)
Fetching: fast-stemmer-1.0.2.gem (100%)
Building native extensions.  This could take a while...
Fetching: classifier-1.3.4.gem (100%)
Fetching: timers-1.1.0.gem (100%)
Fetching: celluloid-0.15.2.gem (100%)
ERROR:  Error installing jekyll:
	celluloid requires Ruby version &gt;= 1.9.2.&#x60;    
Installed ruby1.9.3 using    
&#x60;sudo apt-get install ruby1.9.3&#x60;.

After that I tried again    
&#x60;sudo gem install jekyll&#x60;.
Result was again    
&#x60;ERROR:  Error installing jekyll:
	celluloid requires Ruby version &gt;= 1.9.2.&#x60;

What&#x27;s wrong? Installation experience could be better.

Thanks in advance for reply.

Magnus

### Comments

---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/43#issuecomment-42775262) on: **Sunday, May 11, 2014**

@magnusmetz Remove all Ruby versions previous to 1.9.3. You should see:

&#x60;&#x60;&#x60;bash
$ ruby -v
ruby 1.9.3p0 (2011-10-30 revision 33570) [i686-linux]
&#x60;&#x60;&#x60;
---
> from: [**magnusmetz**](https://github.com/jekyll/jekyll-help/issues/43#issuecomment-42776922) on: **Sunday, May 11, 2014**

thanks. followed your instructions resulting in following output:
~~~~~
$ ruby -v
ruby 1.9.3p0 (2011-10-30 revision 33570) [x86_64-linux]
~~~~~

Tried to install jekyll again with
~~~
gem install jekyll
~~~
but resulted in following error message:

~~~
Fetching: liquid-2.5.5.gem (100%)
ERROR:  While executing gem ... (Errno::EACCES)
    Permission denied - /var/lib/gems/1.9.1
~~~
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/43#issuecomment-42777358) on: **Sunday, May 11, 2014**

&#x60;sudo gem install&#x60; if you&#x27;re using system ruby.


On 12 May 2014 03:23, Magnus Metz &lt;notifications@github.com&gt; wrote:

&gt; thanks. followed your instructions resulting in following output:
&gt;
&gt; $ ruby -v
&gt; ruby 1.9.3p0 (2011-10-30 revision 33570) [x86_64-linux]
&gt;
&gt; Tried to install jekyll again with
&gt;
&gt; gem install jekyll
&gt;
&gt; but resulted in following error message:
&gt;
&gt; Fetching: liquid-2.5.5.gem (100%)
&gt; ERROR:  While executing gem ... (Errno::EACCES)
&gt;     Permission denied - /var/lib/gems/1.9.1
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/43#issuecomment-42776922&gt;
&gt; .
&gt;
---
> from: [**magnusmetz**](https://github.com/jekyll/jekyll-help/issues/43#issuecomment-42777573) on: **Sunday, May 11, 2014**

thanks. now I was able to install it and I can continue with my first steps in jekyll.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/43#issuecomment-42777938) on: **Sunday, May 11, 2014**

Quickest way to fix it is to &#x60;gem install therubyracer&#x60;, but the part of
jekyll causing this issue is being pulled out into it&#x27;s own gem &quot;Soon&quot;[TM]


On 12 May 2014 03:58, Magnus Metz &lt;notifications@github.com&gt; wrote:

&gt; Now I have a new problem while trying:
&gt;
&gt; jekyll new Dynamic Documents
&gt;
&gt; I get following output:
&gt;
&gt; /var/lib/gems/1.9.1/gems/execjs-2.0.2/lib/execjs/runtimes.rb:51:in &#x60;autodetect&#x27;: Could not find a JavaScript runtime. See https://github.com/sstephenson/execjs for a list of available runtimes. (ExecJS::RuntimeUnavailable)
&gt;     from /var/lib/gems/1.9.1/gems/execjs-2.0.2/lib/execjs.rb:5:in &#x60;&lt;module:ExecJS&gt;&#x27;
&gt;     from /var/lib/gems/1.9.1/gems/execjs-2.0.2/lib/execjs.rb:4:in &#x60;&lt;top (required)&gt;&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /var/lib/gems/1.9.1/gems/coffee-script-2.2.0/lib/coffee_script.rb:1:in &#x60;&lt;top (required)&gt;&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /var/lib/gems/1.9.1/gems/coffee-script-2.2.0/lib/coffee-script.rb:1:in &#x60;&lt;top (required)&gt;&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /var/lib/gems/1.9.1/gems/jekyll-coffeescript-1.0.0/lib/jekyll-coffeescript.rb:2:in &#x60;&lt;top (required)&gt;&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /var/lib/gems/1.9.1/gems/jekyll-2.0.3/lib/jekyll.rb:73:in &#x60;&lt;top (required)&gt;&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in &#x60;require&#x27;
&gt;     from /var/lib/gems/1.9.1/gems/jekyll-2.0.3/bin/jekyll:6:in &#x60;&lt;top (required)&gt;&#x27;
&gt;     from /usr/local/bin/jekyll:19:in &#x60;load&#x27;
&gt;     from /usr/local/bin/jekyll:19:in &#x60;&lt;main&gt;&#x27;
&gt;
&gt; I already tried
&gt;
&gt; sudo gem install execjs
&gt;
&gt; but didn&#x27;t solve anything.
&gt;
&gt; What&#x27;s to do? Thanks for further help...
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/jekyll-help/issues/43#issuecomment-42777872&gt;
&gt; .
&gt;
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/43#issuecomment-42779351) on: **Sunday, May 11, 2014**

@magnusmetz Can you please close the dupe https://github.com/jekyll/jekyll/issues/2392?
