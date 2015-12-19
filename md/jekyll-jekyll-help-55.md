# [ERROR when installing jekyll](https://github.com/jekyll/jekyll-help/issues/55)

> state: **closed** opened by: **** on: ****



When i trying to install jekyll with &lt;code&gt;$ gem install jekyll&lt;/code&gt; i got this:
&#x60;&#x60;&#x60;
ERROR:  Loading command: install (LoadError)
    cannot load such file -- zlib
ERROR:  While executing gem ... (NoMethodError)
    undefined method &#x60;invoke_with_build_args&#x27; for nil:NilClass
&#x60;&#x60;&#x60;

I have already installed this package:
&#x60;&#x60;&#x60;

$ ruby -v
ruby 2.1.2p95 (2014-05-08 revision 45877) [i686-linux]
$ gem -v
2.2.2
$ sudo update-alternatives --config ruby
There is only one alternative in link group ruby: /usr/bin/ruby1.9.1
Nothing to configure.
$ sudo update-alternatives --config gem
There is only one alternative in link group gem: /usr/bin/gem1.9.1
Nothing to configure.
&#x60;&#x60;&#x60;



### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-44146722) on: **Sunday, May 25, 2014**

Hi @Vercryger, It seems that &#x60;zlib&#x60; is missing... If you are using Ubuntu I believe its name will be &#x60;zlib1g-dev&#x60;; try installing it and see if it solves your problem.
---
> from: [**Vercryger**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-44146804) on: **Sunday, May 25, 2014**

Hi @albertogg, &lt;code&gt;zlib1g-dev&lt;/code&gt; is already installed in the newest version, i don&#x27;t understand :/

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-44146826) on: **Sunday, May 25, 2014**

Did you install it before installing your Rubies?
---
> from: [**Vercryger**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-44146885) on: **Sunday, May 25, 2014**

Mmm i think no,  should i have to?
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-44146992) on: **Sunday, May 25, 2014**

I&#x27;m not completely sure of, but I always do, you can try that.
---
> from: [**Vercryger**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-44147886) on: **Sunday, May 25, 2014**

Great! It&#x27;s all working! I followed this steps:
First i removed all previously installed rubies on my OS.
And next this:

&#x60;&#x60;&#x60;
sudo apt-get update
sudo apt-get install git-core curl zlib1g-dev build-essential libssl-dev libreadline-dev libyaml-dev libsqlite3-dev sqlite3 libxml2-dev libxslt1-dev libcurl4-openssl-dev python-software-properties

sudo apt-get install libgdbm-dev libncurses5-dev automake libtool bison libffi-dev
curl -L https://get.rvm.io | bash -s stable
source ~/.rvm/scripts/rvm
echo &quot;source ~/.rvm/scripts/rvm&quot; &gt;&gt; ~/.bashrc
rvm install 2.1.2
rvm use 2.1.2 --default
&#x60;&#x60;&#x60;
and finally
&lt;code&gt;gem install jekyll&lt;/code&gt;
and Bingo!

Thanks a lot! This thread can be closed.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-44147911) on: **Sunday, May 25, 2014**

No problem! glad I could help.
---
> from: [**thomasdola**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-87967108) on: **Monday, March 30, 2015**

how did y remove all previously installed rubies on your OS???
need help with that... thank y...
---
> from: [**Vercryger**](https://github.com/jekyll/jekyll-help/issues/55#issuecomment-88182022) on: **Tuesday, March 31, 2015**

In my case, i did that just removing from my ubuntu synaptic any gem and ruby too. Then i deleted all my gems from RVM, i followed [this guide](http://stackoverflow.com/questions/4907668/removing-all-installed-gems-and-starting-over).
Remember, you have to have Node previously installed for Jekyll.
