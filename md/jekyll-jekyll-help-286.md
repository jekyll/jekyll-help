# [Installing Jekyll is out-dated](https://github.com/jekyll/jekyll-help/issues/286)

> state: **closed** opened by: **** on: ****

I&#x27;m trying to follow [instructions][1] to install Jekyll.
I&#x27;m using ruby 1.9.1
I got the following error:
&#x60;&#x60;&#x60;
ERROR:  Error installing github-pages:
        public_suffix requires Ruby version &gt;= 2.0.
&#x60;&#x60;&#x60;
This error happens if I run &#x60;bundler install&#x60; or &#x60;gem install github-pages&#x60;.

I think, it is a result of new version of public_suffix, released Mar. 25, 2015.
I&#x27;ve found the following workaround: &#x60;gem install public_suffix -v 1.4.6&#x60;

After this, command &#x60;gem install github-pages&#x60; succeeded.

[1]: https://help.github.com/articles/using-jekyll-with-pages/#installing-jekyll

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/286#issuecomment-89145503) on: **Thursday, April 02, 2015**

Hi @starius, I believe Jekyll will install with no issues under Ruby &gt;= 2.0... that said it seems that you did a work around and worked. Remember to update your Ruby to 2.0+ if you still have any issues. 
---
> from: [**starius**](https://github.com/jekyll/jekyll-help/issues/286#issuecomment-89218950) on: **Friday, April 03, 2015**

I believe that the issue is not fixed. [The instruction][1] says the following:

&gt; Your Ruby version should begin with 1.9.3 or 2.0.0.

This part is invalid, because the instruction doesn&#x27;t  work with Ruby 1.9.3, if you don&#x27;t apply the workaround, which isn&#x27;t mentioned in the instruction.

&gt; Remember to update your Ruby to 2.0+

Easier said than done. This would require me to upgrade gcc to 4.9, which means upgrading almost all packages from Debian Wheezy to Debian Testing. I use my workaround and want to share it with other users of Ruby 1.9.

Could you change the instruction, please? I think, it worth to mention about the problem with Ruby 1.9 and the workaround. Or you can remove &#x60;should begin with 1.9.3&#x60; part from the instruction.

[1]: https://help.github.com/articles/using-jekyll-with-pages/#installing-jekyll
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/286#issuecomment-89304997) on: **Friday, April 03, 2015**

@starius, current stable jekyll (version 2.5.x) does work under Ruby 1.9.3... but this [ruby has reached the end of life](https://www.ruby-lang.org/en/news/2015/02/23/support-for-ruby-1-9-3-has-ended/) and I will not advice installing a not supported ruby... That said, documentation for this version is correct. Remember that if you find any issues with the documentation you can submit a pull request to fix this. 
---
> from: [**KurtPfeifle**](https://github.com/jekyll/jekyll-help/issues/286#issuecomment-98888243) on: **Monday, May 04, 2015**

I currently have the same issue on OS X with...

1. ...first a Ruby v1.9.3 (installed via MacPorts)
1. ...second the system-provided Ruby v2.0.0 (installed by Apple)
1. ...third, a Ruby v2.0.0 (installed via MacPorts)
1. ...fourth, a Ruby v2.2.2 (installed via MacPorts)

I&#x27;ll now try the workaround...

---
> from: [**KurtPfeifle**](https://github.com/jekyll/jekyll-help/issues/286#issuecomment-98894269) on: **Monday, May 04, 2015**

Nope -- the workaround proposed by @starius doesn&#x27;t work here. For neither version of Ruby.

Because meanwhile I also tested said &quot;workaround&quot; in combination with a (reverted back to) Ruby v1.9 and an additionally installed Ruby v2.1.

    kp@mbp:&gt;  sudo gem install public_suffix -v 1.4.6
      Fetching: public_suffix-1.4.6.gem (100%)
      Successfully installed public_suffix-1.4.6
      1 gem installed
      Installing ri documentation for public_suffix-1.4.6...
      Building YARD (yri) index for public_suffix-1.4.6...
      [warn]: Could not find readme file: README.rdoc
      Installing RDoc documentation for public_suffix-1.4.6...

    kp@mbp:&gt;  ruby --version
      ruby 2.0.0p645 (2015-04-13 revision 50299) [x86_64-darwin13]

    kp@mbp:&gt;  gem list public_suffix
      *** LOCAL GEMS ***
      public_suffix (1.4.6)

    kp@mbp:&gt;  sudo gem install github-pages
      Fetching: github-pages-health-check-0.3.0.gem (100%)
      Fetching: terminal-table-1.4.5.gem (100%)
      Fetching: github-pages-34.gem (100%)
      ERROR:  Error installing github-pages:
              github-pages requires Ruby version &gt;= 2.0.0.

    kp@mbp:&gt;  sudo port -pf install ruby21
    kp@mbp:&gt;  sudo rm -rf /opt/local/bin/ruby
    kp@mbp:&gt;  sudo ln -s /opt/local/bin/ruby2.1 /opt/local/bin/ruby
    
    kp@mbp:&gt;  ruby --version
      ruby 2.1.6p336 (2015-04-13 revision 50298) [x86_64-darwin13]

    kp@mbp:&gt;  gem list public_suffix
      *** LOCAL GEMS ***
      public_suffix (1.4.6)

    kp@mbp:&gt;  sudo gem install github-pages
      ERROR:  Error installing github-pages:
	          github-pages requires Ruby version &gt;= 2.0.0.

    kp@mbp:&gt;  sudo rm -rf /opt/local/bin/ruby
    kp@mbp:&gt;  sudo ln -s /opt/local/bin/ruby1.9 /opt/local/bin/ruby
    
    kp@mbp:&gt;  ruby --version
      ruby 1.9.3p551 (2014-11-13 revision 48407) [x86_64-darwin13]

    kp@mbp:&gt;  gem list public_suffix
      *** LOCAL GEMS ***
      public_suffix (1.4.6)

    kp@mbp:&gt;  sudo gem install github-pages
      ERROR:  Error installing github-pages:
	          github-pages requires Ruby version &gt;= 2.0.0.

    kp@mbp:&gt;  port installed ruby*
      The following ports are currently installed:
        ruby @1.8.7-p374_0+dtrace+tk (active)
        ruby19 @1.9.3-p551_0+c_api_docs+doc+tk (active)
        ruby20 @2.0.0-p645_0+doc (active)
        ruby21 @2.1.6_0+doc (active)
        ruby22 @2.2.2_0+doc (active)
        ruby_select @1.0_0 (active)


---
> from: [**starius**](https://github.com/jekyll/jekyll-help/issues/286#issuecomment-98974591) on: **Monday, May 04, 2015**

Hi @KurtPfeifle,

&gt;   ERROR:  Error installing github-pages:
&gt;          github-pages requires Ruby version &gt;= 2.0.0.

This requirement changed in [github-pages version 34][34], released April 8, 2015, after my comments were written (March 28, 2015 and April 3, 2015).

Try to set github-pages&#x27;s [version 33][33]:

&#x60;&#x60;&#x60;bash
$ sudo gem install github-pages -v 33
&#x60;&#x60;&#x60;

By the way make sure that &#x60;gem&#x60; tool is also updated, when you update &#x60;ruby&#x60;. As I remember, they can be of different versions... I don&#x27;t have Mac to check it though. 

[33]: https://rubygems.org/gems/github-pages/versions/33
[34]: https://rubygems.org/gems/github-pages/versions/34
---
> from: [**snarkmaster**](https://github.com/jekyll/jekyll-help/issues/286#issuecomment-100722100) on: **Sunday, May 10, 2015**

Ubuntu 14.04.1 LTS packages Ruby 1.9.3, so this problem will continue well into 2019 (https://wiki.ubuntu.com/LTS). So much for &quot;end of life&quot;, huh?
