# [How can I allow people to read my Jekyll site offline?](https://github.com/jekyll/jekyll-help/issues/241)

> state: **open** opened by: **** on: ****

To avoid having to give offline readers a PDF, I want to simply deliver the _site files to readers. They would simply open the index.html file and load the content in a browser. However, right now it doesn&#x27;t work. I have this in my _config.yml file:

&#x60;&#x60;&#x60;
baseurl: &quot;/acme&quot;
&#x60;&#x60;&#x60;

What would I need to change in the baseurl or elsewhere in order for people to be able to simply open index.html from _site without having jekyll installed and be able to view the site?

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/241#issuecomment-70563072) on: **Monday, January 19, 2015**

You need to make sure, that all your links (everything inside &#x60;href=&quot;&quot;&#x60; and &#x60;src=&quot;&quot;&#x60;) are __relative to the file__. That means they must never start with a &#x60;/&#x60; or a protocol &#x60;http://&#x60;, etc. You need to reference ressources like this: &#x60;src=&quot;../img/image.pnG&quot;&#x60;. Than, if your site is static, it should work out of the box.
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/241#issuecomment-70563327) on: **Monday, January 19, 2015**

Great. Thanks!
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/241#issuecomment-70565148) on: **Monday, January 19, 2015**

Of course if another file is in the same directory, you just need &#x60;href=&quot;about.html&quot;&#x60;, etc.
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/241#issuecomment-70571634) on: **Monday, January 19, 2015**

for me the trick will be figuring out how to create a toggle in my config file that says whether to create an offline site or not. the main audience will read the content online, but a small percentage will need it offline. if online, then the baseurl needs to be appended; if offline, it doesn&#x27;t -- that kind of thing.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/241#issuecomment-70571991) on: **Monday, January 19, 2015**

The tradeoff with what I suggested above is maintainability. Writing out links relatively to files is a _harder_ task than using the root (with &#x60;/&#x60;). However, the above works both online and offline. That&#x27;s the whole point of it.
---
> from: [**tomjohnson1492**](https://github.com/jekyll/jekyll-help/issues/241#issuecomment-71107028) on: **Thursday, January 22, 2015**

I&#x27;m still having trouble getting this to work. I installed a vanilla jekyll site. I added a page called cats.html. On my index page, I linked to it like this: &#x60;&lt;a href=&quot;/cats.html&quot;&gt;cats&lt;/a&gt;&#x60;. The link works fine. However, if I launch the index file in the _site root directory, instead of using the server mode, I see this: http://www.screencast.com/t/eCD9yNjyP

I tried commenting out baseurl and url in the _config.yml file, but I still get the same result when I launch the index.html file.

I know it may seem an odd request, but basically I&#x27;m using Jekyll for documentation. Sometimes people request PDFs due to lack of access to the site. I want to send them a zip file containing the static site, and tell them to just click index.html to open it. 

What am I doing wrong here? Thanks,

Tom


