# [blog creation issue using Jekyll](https://github.com/jekyll/jekyll-help/issues/181)

> state: **open** opened by: **** on: ****

 I went ahead and wanted to create my new blog with the name myblog, as per the instructions on the Jekyll website: http://jekyllrb.com/docs/quickstart/
///#
user@computer:~$ jekyll new myblog
/usr/lib/ruby/1.9.1/rubygems/custom_require.rb:36:in require&#x27;: iconv will be deprecated in the future, use String#encode instead. WARNING: Could not read configuration. Using defaults (and options). No such file or directory - new/_config.yml Building site: new -&gt; myblog /usr/lib/ruby/vendor_ruby/jekyll/site.rb:126:inchdir&#x27;: No such file or directory - /home/user/new/ (Errno::ENOENT)
from /usr/lib/ruby/vendor_ruby/jekyll/site.rb:126:in read_directories&#x27; from /usr/lib/ruby/vendor_ruby/jekyll/site.rb:98:inread&#x27;
from /usr/lib/ruby/vendor_ruby/jekyll/site.rb:38:in process&#x27; from /usr/bin/jekyll:250:in&#x27;
///

Please help me out.
Thanks

### Comments

---
> from: [**bettygcp**](https://github.com/jekyll/jekyll-help/issues/181#issuecomment-61152020) on: **Thursday, October 30, 2014**

I notice that someone closed this issue at jekyll/jekyll 3 hours ago - #3042 - basic usage: creating a new folder &#x27;myblog&#x27; issue. And correctly so.  But,  showing a big &#x27;closed button here might give people the wrong impression that the actual problem is solved. So, just to be clear, I reposted the issue here in the help section as per advice given in the #3042 issue and that&#x27;s what i did. If someone please could solve this unsolved issue, that&#x27;d be great. Thanks. ... Standing by on a solution, Betty.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/181#issuecomment-61438278) on: **Sunday, November 02, 2014**

Hey @bettygcp - Can you give some more detail on your setup? What version of Jekyll do you have, what operating system are you using, etc? I&#x27;m not having any trouble with the &#x60;jekyll new&#x60; command...

&#x60;&#x60;&#x60;bash
Troys-Air:Projects troy$ jekyll new myblog
New jekyll site installed in /Users/troy/Projects/myblog. 
Troys-Air:Projects troy$ cd myblog/
Troys-Air:myblog troy$ jekyll build
Configuration file: /Users/troy/Projects/myblog/_config.yml
            Source: /Users/troy/Projects/myblog
       Destination: /Users/troy/Projects/myblog/_site
      Generating... 
                    done.
 Auto-regeneration: disabled. Use --watch to enable.
&#x60;&#x60;&#x60;
---
> from: [**bettygcp**](https://github.com/jekyll/jekyll-help/issues/181#issuecomment-61454643) on: **Monday, November 03, 2014**

Hello @troyswanson , thanks for the reply!

Here&#x27;s my setup:

My system: ubuntu 14.04 LTS (trusty tahr) - OS type 64-bit

I installed the required packages for Jekyll to run:

ruby - version 1:1.9.3.4 which is the latest version (checked in synaptic package manager)

rubygems-integration package - version 1.5 (= latest version)

NodeJS - version 0.10.25-dfsg2-2ubuntu1 (= latest version)

And finally i did a gem install jekyll ( i actually had to do a sudo gem install jekyll, because otherwise it wouldn&#x27;t let me install) and got a jekyll folder in my home directory with what looks like a jekyll site 
(i also think stuff was installed in places where i don&#x27;t have permissions).

The installed version of Jekyll on my system is 0.11.2-1, which is the latest version

I wanted to try if by using the commands i found in the quickstart user guide, it would actually automatically create a folder in my home directory called &#x27;newblog&#x27;). For the command i tried see my first post in this issue #181 above.

Please note that i am fairly new to ubuntu systems and not all experienced in using the terminal.
I think what would be best is step by step instructions while explaining what each step actually does.
And then there&#x27;s the permissions issue too. Some commands will yield a response from the terminal that i don&#x27;t have permission to do that. At this time i can&#x27;t duplicate the commands which generated the &#x27;no permission&#x27; errors but i think the software installs jekyll in some of the protected system folders and what i believe is that everything should be installed inside my accessible home directory.

Thanks again for coming on board over this issue.
@bettygcp
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/181#issuecomment-61538876) on: **Monday, November 03, 2014**

Something is seriously amiss.

Jekyll 0.11.2 came out in December of 2011. The current version of Jekyll is 2.4.0.

A step-by-step tutorial will only really work on a clean installation of Ubuntu. It seems like you&#x27;ve got a lot of stuff that is kind of wonky (e.g. permissions). Are you using a cloud provider or any other system that will allow you to create a fresh instance?
---
> from: [**bettygcp**](https://github.com/jekyll/jekyll-help/issues/181#issuecomment-61564169) on: **Monday, November 03, 2014**

Well, i can say that i installed Jekyll and its required packages on my second ubuntu computer with a quasi fresh install. The only programs i had running on there prior to the Jekyll (and its required packages) install was gdebi, for installing deb packages, and Synaptic package Manager (for installing ubuntu programs and which i installed to see which versions i have actually running of Jekyll and its depency or required packages, after you asked me the Jekyll version.)

As for the installation of Jekyll, i did everything as per the Jeckyll instructions over at 
http://jekyllrb.com/docs/installation/ , the (hopefully) official Jekyll documentation site.
Firstly, the required packages ruby, rubygems and NodeJS are clearly mentioned on that Jekyll installation page as to be installed prior to installing Jeckyll, so that&#x27;s what i did.

Then, i tried  installing Jeckyl via rubygems, again, i used the exact command as specified on the Jekyll official installation page but this didn&#x27;t work.
Then i turned to the last option on the Jekyll installation page, which is to install the developer&#x27;s version of Jeckyll using the following commands, as specified on the installation page:

$ git clone git://github.com/jekyll/jekyll.git
$ cd jekyll
$ script/bootstrap
$ bundle exec rake build
$ ls pkg/*.gem | head -n 1 | xargs gem install -l

This worked up until to the last command.

It was only the last command which probably gave me a &#x27;no permissions&#x27; error. I do remember that that last command gave me an error but the Jeckyll folder was created in my home directory so i figured i&#x27;d be all right. But at first i didn&#x27;t realise that the Jeckyll folder in my directory actually just contains the Jeckyll site, as an example if you will. (that&#x27;s what i&#x27;m thinking)

Verifying the version in Synaptic, i ended up with the Jeckyll version 0.11.2-1, which Synaptic manager specifified is the latest version -Then, following the instructions in the quickart guide
at http://jekyllrb.com/docs/quickstart/ , using command ~ $ gem install jekyll didn&#x27;t work - the terminal gave an error and so when i tried again and put sudo in front of the command, that&#x27;s the only way it would actually install Jekyll, but again that&#x27;s probably where the &#x27;no permissions&#x27; error comes from. Installing without sudo didn&#x27;t work in the terminal. So something&#x27;s definitely wrong with that command in the official Jekyll installation page. I believe it actually gets installed in the /var or/usr directories or one of their subfolders. But like i said, it also created an accessible Jekyll folder in the home directory hte first time round.

After i had sudo gem installed Jekyll i then tried to create the &#x27;newblog&#x27; test blog like mentioned in the first post of this #181 issue. To no avail, see above.

I am not using a cloud provider. I don&#x27;t think i even know what a cloud provider is. But like i say it&#x27;s a fresh ubuntu install on a computer that&#x27;s not the computer i use for every day use so the only things really on there are the required packages for Jeckyll. 

A clean install after this won&#x27;t work because  like i said i first tried Jeckyll install command as per the instructions on their installation page and htat didn&#x27;t work so if i go through the trouble of re-installing ubuntu and then try to install Jekyll again as per the instructions it won&#x27;t work obviously, because i&#x27;ll end up trying the same method but expecting different results. And we all know what Albert Einstein said about that!

---
> from: [**sondr3**](https://github.com/jekyll/jekyll-help/issues/181#issuecomment-61680637) on: **Tuesday, November 04, 2014**

If you have the correct version of Ruby installed and have no problems with your permissions running &#x60;&#x60;&#x60;gem install jekyll&#x60;&#x60;&#x60; will give you a fully updated and ready Jekyll installed on your computer. Something is seriously amiss. Just a couple of days ago I installed Jekyll on Ubuntu locally and I had no issues whatsoever. If you install it via the gem the version should be 2.4.0, not 0.11.2.

I can&#x27;t help you much with synaptic however. When you run &#x60;&#x60;&#x60;ruby -v&#x60;&#x60;&#x60; what is the output? Also, you shouldn&#x27;t have to install the rubygem-integration thing, RubyGems should come installed with Ruby. If you run &#x60;&#x60;&#x60;gem -v&#x60;&#x60;&#x60; you should see the correct version. You are running packages that are a bit outdated, the latest NodeJS version is 0.10.33, the latest Ruby is 2.1.4 and so on.

If I were you I&#x27;d recommend ditching using the synaptic package manager and go with using the command line, as it is how you interact with these tools so getting used to it is better than via a package manager. You could follow this tutorial: https://www.digitalocean.com/community/tutorials/how-to-install-ruby-on-rails-on-ubuntu-12-04-lts-precise-pangolin-with-rvm and just replace the last step of install rails with jekyll instead. Good luck.
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/181#issuecomment-61681757) on: **Tuesday, November 04, 2014**

Thanks for jumping in, @sondr3. :clap: 
