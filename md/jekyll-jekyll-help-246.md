# [Paths to other pages versus paths to images seems inconsistent](https://github.com/jekyll/jekyll-help/issues/246)

> state: **open** opened by: **** on: ****

I&#x27;m a little confused on subdirectory paths with with links to pages versus links to images. In my root directory, I have a folder called &quot;content.&quot; In &quot;content&quot; I have a subdirectory called &quot;web&quot;. If I want to link to a page in content/web, I just reference the permalink, like this: 

&#x60;&#x60;&#x60;
[Tutorial]({{ &quot;/tutorial&quot; | prepend: site.baseurl }})
&#x60;&#x60;&#x60;
In my root directory, I have another folder called &quot;images.&quot; Inside images, I have a subdirectory called web. If I have an image in there and want to link to it, I would use this:

&#x60;&#x60;&#x60;
![Sample Wiki Image]({{ &quot;/images/web/samplewiki.png&quot; | prepend: site.baseurl }})
&#x60;&#x60;&#x60;

It seems that with pages, which have a permalink attribute in the front matter, I don&#x27;t need to specify the path to the page -- I just reference the page&#x27;s permalink. 

But with images, it&#x27;s not the case. Images require me to use the subdirectory path. 

Is this correct? 


### Comments

