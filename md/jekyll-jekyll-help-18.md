# [THIS IS A PAIN IN THE ASS TO INSTALL](https://github.com/jekyll/jekyll-help/issues/18)

> state: **closed** opened by: **** on: ****

Can I get some instructions that are easier to install Jekyll????

### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/18#issuecomment-40535775) on: **Tuesday, April 15, 2014**

What instructions are you using currently and what operating system are you using? At it&#x27;s simplest it&#x27;s just (assuming you have ruby available on your system):

&#x60;&#x60;&#x60;
$ gem install jekyll
$ jekyll new my_new_website
$ cd my_new_website
$ jekyll build
&#x60;&#x60;&#x60;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/18#issuecomment-40535808) on: **Tuesday, April 15, 2014**

Hey. We need some more info in order to help you.

* What operating system are you running?
* Are you using a version manager for Ruby?
* What&#x27;s the error you&#x27;re getting?

Sorry you&#x27;re having trouble, but help us help you.
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/18#issuecomment-40536324) on: **Tuesday, April 15, 2014**

Yes, please provide more information to guide you in the best possible way. Have you checked the [documentation](http://jekyllrb.com/docs/installation/) about installation?
---
> from: [**iamgabeortiz**](https://github.com/jekyll/jekyll-help/issues/18#issuecomment-41510286) on: **Sunday, April 27, 2014**

For any windows customers, I&#x27;ve worked up a little [guide](https://github.com/iamgabeortiz/install-jekyll-win) that combines steps from a few sources into a working install for Windows 7.
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/18#issuecomment-41511258) on: **Sunday, April 27, 2014**

I have something similar also available at
http://matt.scharley.me/2012/03/10/windows-cygwin-and-jekyll.html for
jekyll via cygwin. I don&#x27;t actually use that setup any more and I keep
meaning to write up a guide without cygwin.
On 28/04/2014 7:43 AM, &quot;Gabe Ortiz&quot; &lt;notifications@github.com&gt; wrote:

&gt; For any windows customers, I&#x27;ve worked up a little guide&lt;https://github.com/iamgabeortiz/install-jekyll-win&gt;that combines steps from a few sources into a working install for Windows 7.
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/help/issues/18#issuecomment-41510286&gt;
&gt; .
&gt;
---
> from: [**acidflip**](https://github.com/jekyll/jekyll-help/issues/18#issuecomment-41511323) on: **Sunday, April 27, 2014**

I can install it (I think). I just don&#x27;t know what to do after I get the Jekyll files into the repository
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/18#issuecomment-41511843) on: **Sunday, April 27, 2014**

Try having a look at jekyllrb.com. There&#x27;s a bunch of documentation there
about layouts and posts and such. If you don&#x27;t use either, then jekyll will
essentially just translate files one to one from the source to destination
folders doing very little. If you don&#x27;t have a website folder at all yet,
try &#x60;jekyll new folder_name&#x60;.
On 28/04/2014 8:21 AM, &quot;Jason Kinetec&quot; &lt;notifications@github.com&gt; wrote:

&gt; I can install it (I think). I just don&#x27;t know what to do after I get the
&gt; Jekyll files into the repository
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/help/issues/18#issuecomment-41511323&gt;
&gt; .
&gt;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/18#issuecomment-41696959) on: **Tuesday, April 29, 2014**

Going to close this out.

@jsnkntc If you have any other questions, feel free to file a new issue and we&#x27;ll be happy to help you!
